package ibapi

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"reflect"
	"runtime"
	"strconv"
	"time"

	"log"
)

const ibTimeFormat = "20060102 15:04:05 MST"

var (
	ErrUnknownReplyCode   = errors.New("Unknown Reply Code")
	ErrUnknownRequestType = errors.New("Unknown Request Type")
)

func unpanic() {
	if r := recover(); r != nil {
		log.Printf("unpanic, %v", r)
		if _, ok := r.(runtime.Error); ok {
			panic(r)
		}
	}
}

// Decode

type replyReader struct {
	r *bufio.Reader
}

type ReplyDecoder interface {
	ReplyDecoder(m *replyReader, version int64)
}

func NewReplyReader(r io.Reader) *replyReader {
	return &replyReader{bufio.NewReader(r)}
}

func (m *replyReader) readString() string {
	data, err := m.r.ReadString(0)

	if err != nil {
		panic(err)
	}

	return string(data[:len(data)-1])
}

func (m *replyReader) readInt() int64 {
	i, err := strconv.ParseInt(m.readString(), 10, 64)

	if err != nil {
		panic(err)
	}

	return i
}

func (m *replyReader) readFloat() float64 {
	f, err := strconv.ParseFloat(m.readString(), 64)

	if err != nil {
		panic(err)
	}

	return f
}

func (m *replyReader) readBool() bool {
	return (m.readInt() > 0)
}

func (m *replyReader) readTime() time.Time {
	t, err := time.Parse(ibTimeFormat, m.readString())

	if err != nil {
		panic(err)
	}

	return t
}

func (m *replyReader) readField(i interface{}, ver, minVer int64) interface{} {
	if ver < minVer {
		return nil
	}

	switch i.(type) {
	case string:
		return m.readString()
	case int64:
		return m.readInt()
	case float64:
		return m.readFloat()
	case bool:
		return m.readBool()
	case time.Time:
		return m.readTime()
	}

	t := reflect.TypeOf(i)
	switch t.Kind() {
	case reflect.Slice:
		size := int(m.readInt())
		// create new array
		arrayValue := reflect.MakeSlice(t, size, size)
		for j := 0; j < size; j++ {
			newElem := reflect.New(t.Elem()).Interface()
			elem := m.readField(newElem, ver, minVer)
			arrayValue.Index(j).Set(reflect.Indirect(reflect.ValueOf(elem)))
		}
		return arrayValue.Interface()
	case reflect.Ptr:
		if reflect.TypeOf(reflect.Indirect(reflect.ValueOf(i))).Kind() == reflect.Struct {
			return m.readStruct(i, ver)
		}
	}

	return nil
}

func (m *replyReader) defaultReplyDecode(rep interface{}, ver int64) {
	// write serialized fields
	val := reflect.ValueOf(rep).Elem()
	for i := 0; i < val.NumField(); i++ {
		valueField := val.Field(i)
		typeField := val.Type().Field(i)
		tag := typeField.Tag

		var minVer int64 = 0
		if ver, err := strconv.ParseInt(tag.Get("minVer"), 10, 64); err != nil {
			minVer = ver
		}
		val := m.readField(valueField.Interface(), ver, minVer)
		if val != nil {
			valueField.Set(reflect.ValueOf(val))
		}
	}
}

func (m *replyReader) readStruct(s interface{}, version int64) interface{} {
	if sDec, ok := s.(ReplyDecoder); ok {
		sDec.ReplyDecoder(m, version)
	} else {
		m.defaultReplyDecode(s, version)
	}

	return s
}

func (m *replyReader) Read() (rep interface{}, err error) {
	defer unpanic()

	// read code
	code := m.readInt()
	// read version
	version := m.readInt()

	// read payload
	rep = repStruct(code)
	if rep == nil {
		log.Println("Bad Reply With Code ", code, " version ", version)
		return nil, ErrUnknownReplyCode
	}

	return m.readStruct(rep, version), nil
}

// Encode
type RequestEncoder interface {
	RequestEncode(m *requestBytes)
}

type requestBytes struct {
	buf       *bytes.Buffer
	verServer int64
}

func NewRequestBytes() *requestBytes {
	return &requestBytes{bytes.NewBuffer(make([]byte, 0, 4096)), 0}
}

func (m *requestBytes) writeString(v string) {
	if _, err := m.buf.WriteString(v + "\000"); err != nil {
		panic(err)
	}
}

func (m *requestBytes) writeInt(v int64) {
	m.writeString(strconv.FormatInt(v, 10))
}

func (m *requestBytes) writeFloat(v float64) {
	m.writeString(strconv.FormatFloat(v, 'g', 10, 64))
}

func (m *requestBytes) writeBool(v bool) {
	if v {
		m.writeString("1")
	} else {
		m.writeString("0")
	}
}

func (m *requestBytes) writeTime(v time.Time) {
	m.writeString(v.Format(ibTimeFormat))
}

func (m *requestBytes) writeField(i interface{}, minVer int64) {
	if m.verServer < minVer {
		return
	}

	switch v := i.(type) {
	case string:
		m.writeString(v)
	case int64:
		m.writeInt(v)
	case float64:
		m.writeFloat(v)
	case bool:
		m.writeBool(v)
	case time.Time:
		m.writeTime(v)
	}

	t := reflect.TypeOf(i)
	v := reflect.ValueOf(i)
	switch t.Kind() {
	case reflect.Slice:
		// send size and iterate
		size := v.Len()
		m.writeInt(int64(size))
		for j := 0; j < size; j++ {
			elem := v.Index(j).Interface()
			m.writeField(elem, 0)
		}
	case reflect.Ptr:
		if reflect.TypeOf(reflect.Indirect(v)).Kind() == reflect.Struct {
			m.writeStruct(i)
		}
	case reflect.Struct:
		m.writeStruct(i)
	}
}

func (m *requestBytes) writeStruct(req interface{}) {
	if reqEnc, ok := req.(RequestEncoder); ok {
		reqEnc.RequestEncode(m)
	} else {
		m.defaultRequestEncode(req)
	}
}

func (m *requestBytes) Write(req interface{}) error {
	defer unpanic()

	m.buf.Reset()

	code := reqCode(req)
	if code < 0 {
		return ErrUnknownRequestType
	}

	// write MSG code
	m.writeInt(code)
	// write MSG version
	m.writeInt(reqVersion(req))
	// write struct fields
	m.writeStruct(req)

	return nil
}

func (m *requestBytes) Bytes() []byte {
	return m.buf.Bytes()
}

func (m *requestBytes) String() string {
	return m.buf.String()
}

func (m *requestBytes) defaultRequestEncode(req interface{}) {
	// write serialized fields
	val := reflect.ValueOf(req)
	if val.Type().Kind() != reflect.Struct {
		val = val.Elem()
	}
	for i := 0; i < val.NumField(); i++ {
		valueField := val.Field(i)
		typeField := val.Type().Field(i)
		tag := typeField.Tag

		var minVer int64 = 0
		if ver, err := strconv.ParseInt(tag.Get("minVer"), 10, 64); err != nil {
			minVer = ver
		}
		m.writeField(valueField.Interface(), minVer)
	}
}

func reqCode(i interface{}) int64 {
	switch i.(type) {
	case *MsgOutReqMktData:
		return mOutRequestMarketData
	case *MsgOutCxlMktData:
		return mOutCancelMarketData
	case *MsgOutReqContractData:
		return mOutRequestContractData
	case *MsgOutReqMktDataType:
		return mOutRequestMarketDataType
	}
	return -1
}

func reqVersion(i interface{}) int64 {
	switch i.(type) {
	case *MsgOutReqMktData:
		return 9
	case *MsgOutReqContractData:
		return 6
	}
	return 1
}

func repStruct(code int64) interface{} {
	switch code {
	case mInTickPrice:
		return &MsgInTickPrice{}
	case mInTickSize:
		return &MsgInTickSize{}
	case mInTickOptionComputation:
		return &MsgInTickOptionComputation{}
	case mInTickGeneric:
		return &MsgInTickGeneric{}
	case mInTickString:
		return &MsgInTickString{}
	case mInTickEFP:
		return &MsgInTickEFP{}
	case mInTickSnapshotEnd:
		return &MsgInTickSnapshotEnd{}
	case mInMarketDataType:
		return &MsgInMarketDataType{}
	case mInContractData:
		return &MsgInContractData{}
	case mInContractDataEnd:
		return &MsgInContractDataEnd{}
	case mInManagedAccounts:
		return &MsgInManagedAccounts{}
	case mInNextValidId:
		return &MsgInNextValidId{}
	case mInErrorMessage:
		return &MsgInError{}
	case mInCurrentTime:
		return &MsgInCurrentTime{}
	default:
		return nil
	}
}
