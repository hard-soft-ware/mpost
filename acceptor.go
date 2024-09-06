package mpost

import (
	"github.com/hard-soft-ware/mpost/acceptor"
	"github.com/hard-soft-ware/mpost/enum"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"go.bug.st/serial"
	"os"
	"sync"
	"time"
)

////////////////////////////////////

type EventHandler func(*CAcceptor, int)

type CAcceptor struct {
	port                serial.Port
	auditLifeTimeTotals []int
	auditPerformance    []int
	auditQP             []int

	coupon *CCoupon

	docType enum.DocumentType

	workerThread        sync.WaitGroup
	openThread          chan bool
	flashDownloadThread chan bool
	dataLinkLayer       *CDataLinkLayer

	messageQueue chan *CMessage
	replyQueue   chan []byte

	eventHandlers map[enum.EventType]EventHandler

	log LogStruct
}

func NewCAcceptor(transactionTimeout, downloadTimeout time.Duration) *CAcceptor {

	acceptor.Timeout.Transaction = transactionTimeout
	acceptor.Timeout.Download = downloadTimeout

	a := &CAcceptor{
		eventHandlers: make(map[enum.EventType]EventHandler, enum.Event_End),

		messageQueue:        make(chan *CMessage, 1),
		replyQueue:          make(chan []byte, 1),
		flashDownloadThread: make(chan bool, 1),
		openThread:          make(chan bool, 1),

		log: NewLog(
			log.Output(zerolog.ConsoleWriter{
				Out:        os.Stdout,
				NoColor:    false,
				TimeFormat: "15:04:05",
			}),
			"Acceptor",
		),
	}

	return a
}

//

func (a *CAcceptor) getTickCount() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

func (a *CAcceptor) SetEventHandler(event enum.EventType, eventHandler func(*CAcceptor, int)) {
	a.eventHandlers[event] = eventHandler
}
