package ibapi

import (
	"errors"
	"net"

	"log"
	"strings"
)

const (
	DefaultGateway = "127.0.0.1:4001"
)

var (
	ErrBadServerHandshake = errors.New("Invalid handshake response")
	ErrClosed             = errors.New("Session closed")
)

type Engine struct {
	conn         net.Conn
	replyReader  *replyReader
	requestBytes *requestBytes
}

func NewEngine(gateway string, clientId int64) (*Engine, error) {
	e := new(Engine)

	// dial conn
	conn, err := net.Dial("tcp", gateway)
	if err != nil {
		return nil, err
	}

	e.conn = conn
	e.replyReader = NewReplyReader(e.conn)
	e.requestBytes = NewRequestBytes()

	// exchange handshaking info: send client info, receive server info

	c := &MsgOutClientHandshake{VerClient, clientId}
	log.Println("Writing Client Handshake")
	e.requestBytes.writeStruct(c)
	err = e.writeRequestBytes()
	if err != nil {
		return nil, err
	}
	log.Println("Wrote Client Handshake")

	log.Println("Reading Server Handshake")
	rep := e.replyReader.readStruct(new(MsgInServerHandshake), 0)
	log.Println("Read Server Handshake")
	s, ok := rep.(*MsgInServerHandshake)
	if !ok {
		return nil, ErrBadServerHandshake
	}

	e.requestBytes.verServer = s.Version
	log.Println("Server Time ", s.Time)

	return e, nil
}

func (e *Engine) ReadReply() (interface{}, error) {
	if e.conn == nil {
		return nil, ErrClosed
	}

	return e.replyReader.Read()
}

func (e *Engine) WriteRequest(req interface{}) error {
	if e.conn == nil {
		return ErrClosed
	}

	// load to bytes then write to conn
	err := e.requestBytes.Write(req)
	if err != nil {
		return err
	}

	return e.writeRequestBytes()
}

func (e *Engine) writeRequestBytes() error {
	log.Println("Writing bytes, %v", strings.Replace(e.requestBytes.String(), "\000", "-", -1))
	_, err := e.conn.Write(e.requestBytes.Bytes())
	return err
}

func (e *Engine) Stop() error {
	err := e.conn.Close()
	e.conn = nil
	return err
}
