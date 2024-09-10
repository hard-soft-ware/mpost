package mpost

import (
	"github.com/hard-soft-ware/mpost/acceptor"
	"github.com/hard-soft-ware/mpost/enum"
	"github.com/hard-soft-ware/mpost/serial"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"sync"
	"time"
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
	openThread          chan bool
	flashDownloadThread chan bool
	dataLinkLayer       *CDataLinkLayer

	messageQueue chan *CMessage
	replyQueue   chan []byte

	eventHandlers map[enum.EventType]EventHandler

	log LogStruct

	ss bool
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

func (a *CAcceptor) AddHook(ev enum.EventType, h EventHandler) {
	a.eventHandlers[ev] = h
}
