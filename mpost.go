package mpost

import (
	"context"
	"github.com/hard-soft-ware/mpost/acceptor"
	"github.com/hard-soft-ware/mpost/command"
	"github.com/hard-soft-ware/mpost/enum"
	"github.com/hard-soft-ware/mpost/hook"
	"github.com/hard-soft-ware/mpost/serial"
)

////////////////////////////////////

type MpostObj struct {
	port          *serial.SerialStruct
	dataLinkLayer *dataObj
	coupon        *CouponObj

	messageQueue chan *messageObj
	replyQueue   chan []byte

	DocType enum.DocumentType
	Log     *LogObj
	Method  *MethodsObj

	Ctx       context.Context
	CtxCancel context.CancelFunc
}

func New() *MpostObj {
	obj := MpostObj{
		messageQueue: make(chan *messageObj, 1),
		replyQueue:   make(chan []byte, 1),
		Log:          newLog(),
	}

	acceptor.Clean()
	command.Clean()
	hook.Clean()

	obj.Method = obj.newMethods()
	obj.Ctx, obj.CtxCancel = context.WithCancel(context.Background())

	hook.Raise.Log = func(eventType enum.EventType, i int) {
		obj.Log.Event(eventType, i)
	}

	return &obj
}

//

func (a *MpostObj) AddHook(ev enum.EventType, h func(int)) {
	hook.Add(ev, h)
}
