package mpost

import (
	"context"
	"github.com/hard-soft-ware/mpost/enum"
	"github.com/hard-soft-ware/mpost/serial"
	"github.com/rs/zerolog"
	"sync"
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

	workerThread        sync.WaitGroup
	openThread          bool
	flashDownloadThread bool
	dataLinkLayer       *CDataLinkLayer

	messageQueue chan *CMessage
	replyQueue   chan []byte

	eventHandlers map[enum.EventType]EventHandler

	log LogStruct

	Ctx       context.Context
	CtxCancel context.CancelFunc
	ss        bool
}

var DefAcceptor = &CAcceptor{
	eventHandlers: make(map[enum.EventType]EventHandler, enum.Event_End),

	messageQueue: make(chan *CMessage, 1),
	replyQueue:   make(chan []byte, 1),
}

func init() {
	DefAcceptor.Ctx, DefAcceptor.CtxCancel = context.WithCancel(context.Background())
}

//

func (a *CAcceptor) AddLog(log zerolog.Logger, root string) {
	a.log = NewLog(log, root)
}

func (a *CAcceptor) AddHook(ev enum.EventType, h EventHandler) {
	a.eventHandlers[ev] = h
}
