package ibapi

import (
	"bufio"
	"bytes"
	"reflect"
	"strconv"
	"testing"
	"time"
)

func TestWriteString(t *testing.T) {
	b := NewRequestBytes()
	b.writeString("foobar")
	expected := "foobar\000"
	if b.String() != expected {
		t.Fatalf("writeString('foobar') = %s, want %s", b.String(), expected)
	}
}

func TestWriteInt(t *testing.T) {
	b := NewRequestBytes()
	b.writeInt(int64(57))
	expected := "57\000"
	if b.String() != expected {
		t.Fatalf("writeInt(57) = %s, want %s", b.String(), expected)
	}
}

func TestWriteTime(t *testing.T) {
	b := NewRequestBytes()
	ts := time.Now()
	b.writeTime(ts)
	expected := ts.Format(ibTimeFormat) + "\000"
	if b.String() != expected {
		t.Fatalf("writeTime(%s) = %s, want %s", ts, b.String(), expected)
	}
}

func TestWriteFloat(t *testing.T) {
	f := 0.535
	b := NewRequestBytes()
	b.writeFloat(f)
	expected := strconv.FormatFloat(f, 'g', 10, 64) + "\000"
	if b.String() != expected {
		t.Fatalf("writeFloat(%g) = %s, want %s", f, b.String(), expected)
	}
}

func TestReadString(t *testing.T) {
	x := "foobar"
	b := NewRequestBytes()

	b.writeString(x)
	r := NewReplyReader(bufio.NewReader(bytes.NewReader(b.Bytes())))
	y := r.readString()

	if x != y {
		t.Fatalf("expected %d but got %d", x, y)
	}
}

func TestReadInt(t *testing.T) {
	x := int64(57)
	b := NewRequestBytes()

	b.writeInt(x)
	r := NewReplyReader(bufio.NewReader(bytes.NewReader(b.Bytes())))
	y := r.readInt()

	if x != y {
		t.Fatalf("expected %d but got %d", x, y)
	}
}

func TestReadTime(t *testing.T) {
	x := time.Now()
	x = x.Add(time.Duration(-1 * x.Nanosecond()))
	b := NewRequestBytes()

	b.writeTime(x)
	r := NewReplyReader(bufio.NewReader(bytes.NewReader(b.Bytes())))
	y := r.readTime()

	if x != y {
		t.Fatalf("expected %v but got %v", x, y)
	}
}

func TestReadFloat(t *testing.T) {
	x := 0.545
	b := NewRequestBytes()

	b.writeFloat(x)
	r := NewReplyReader(bufio.NewReader(bytes.NewReader(b.Bytes())))
	y := r.readFloat()

	if x != y {
		t.Fatalf("expected %v but got %v", x, y)
	}
}

func TestReadStruct(t *testing.T) {
	x := &MsgInTickPrice{}
	x.TickerId = 1
	x.Type = TickBid
	x.Price = 1351.1353
	x.Size = 1234
	x.CanAutoExecute = false

	s := "1\00010\0001\0001\0001351.1353\0001234\0000\000"

	b := NewRequestBytes()
	b.Write(x)
	r := NewReplyReader(bufio.NewReader(bytes.NewReader([]byte(s))))
	y, err := r.Read()

	if err != nil {
		t.Fatalf("failed to read %v, from %v, got error %v", x, b.String(), err)
	}
	if !reflect.DeepEqual(x, y) {
		t.Fatalf("expected %v but got %v", x, y)
	}
}

func TestWriteStruct(t *testing.T) {
	x := &MsgOutCxlMktData{1}

	s := "2\0001\0001\000"

	b := NewRequestBytes()
	err := b.Write(x)

	if err != nil {
		t.Fatalf("failed to write %v, got error %v", x, err)
	}
	if s != b.String() {
		t.Fatalf("expected %v but got %v", s, b.String())
	}
}

type foo struct {
	Codes []testStruct
}

type testStruct struct {
	Code string
	Desc string
}

func TestReadSlice(t *testing.T) {
	x := &foo{
		Codes: []testStruct{
			{"Code1", "Desc1"},
			{"Code2", "Desc2"},
		},
	}

	b := NewRequestBytes()
	b.writeStruct(x)
	r := NewReplyReader(bufio.NewReader(bytes.NewReader(b.Bytes())))
	y := r.readStruct(&foo{}, 0)

	if !reflect.DeepEqual(x, y) {
		t.Fatalf("expected %d but got %d", x, y)
	}
}
