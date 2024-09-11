package mpost

import (
	"context"
	"github.com/hard-soft-ware/mpost/enum"
	"github.com/hard-soft-ware/mpost/hook"
	"github.com/hard-soft-ware/mpost/serial"
)

////////////////////////////////////

type EventHandler func(*CAcceptor, int)

type CAcceptor struct {
	port                *serial.SerialStruct
	auditLifeTimeTotals []int
	auditPerformance    []int
	auditQP             []int

	coupon *CCoupon

	docType enum.DocumentType

	openThread          bool
	flashDownloadThread bool
	dataLinkLayer       *CDataLinkLayer

	messageQueue chan *CMessage
	replyQueue   chan []byte

	eventHandlers map[enum.EventType]EventHandler

	Log *LogObj

	Ctx       context.Context
	CtxCancel context.CancelFunc
	ss        bool
}

var DefAcceptor = &CAcceptor{
	eventHandlers: make(map[enum.EventType]EventHandler, enum.Event_End),

	messageQueue: make(chan *CMessage, 1),
	replyQueue:   make(chan []byte, 1),
	Log:          newLog(),
}

func init() {
	DefAcceptor.Ctx, DefAcceptor.CtxCancel = context.WithCancel(context.Background())

	hook.Raise.Log = func(eventType enum.EventType, i int) {
		DefAcceptor.Log.Event(eventType, i)
	}
}

//

func (a *CAcceptor) AddHook(ev enum.EventType, h EventHandler) {
	a.eventHandlers[ev] = h
}
