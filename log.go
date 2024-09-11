package mpost

import (
	"github.com/hard-soft-ware/mpost/consts"
	"github.com/hard-soft-ware/mpost/enum"
)

////////////////////////////////////

type LogObj struct {
	Event  func(enum.EventType, int)
	Method func(string, any)

	Msg func(string)
	Err func(string, error)

	SerialSend func(consts.CmdType, []byte)
	SerialRead func(consts.CmdType, []byte)
}

func newLog() *LogObj {
	obj := LogObj{}

	obj.Event = func(eventType enum.EventType, i int) {}
	obj.Method = func(s string, a any) {}

	obj.Msg = func(s string) {}
	obj.Err = func(s string, err error) {}

	obj.SerialSend = func(cmd consts.CmdType, data []byte) {}
	obj.SerialRead = func(cmd consts.CmdType, data []byte) {}

	return &obj
}
