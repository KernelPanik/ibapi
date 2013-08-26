package ibapi

import (
	"time"
)

type MsgInServerHandshake struct {
	Version int64
	Time    time.Time
}

type MsgOutClientHandshake struct {
	Version int64
	Id      int64
}

type MsgInHeader struct {
	Code    int64
	Version int64
}

type MsgInError struct {
	Id      int64
	Code    int64
	Message string
}

type MsgInCurrentTime struct {
	Time int64
}
