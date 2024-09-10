package mpost

import (
	"errors"
	"time"
)

////////////////////////////////////

type CMessage struct {
	Payload       []byte
	PayloadLength int
	IsSynchronous bool
}

func NewCMessage(payload []byte, isSynchronous bool) *CMessage {
	return &CMessage{
		Payload:       payload,
		PayloadLength: len(payload),
		IsSynchronous: isSynchronous,
	}
}

////

func (a *CAcceptor) SendSynchronousCommand(payload []byte) ([]byte, error) {
	if !a.ss {
		a.ss = true
	} else {
		a.ss = false
		payload[0] += 1
	}

	a.messageQueue <- NewCMessage(payload, true)

	select {
	case <-a.ctx.Done():
		a.Close()
		return nil, errors.New("close from context")
	case reply := <-a.replyQueue:
		return reply, nil
	case <-time.After(30 * time.Second):
		return nil, errors.New("timeout waiting for response")
	}

	return nil, errors.New("invalid response")
}

func (a *CAcceptor) SendAsynchronousCommand(payload []byte) {
	if !a.ss {
		a.ss = true
	} else {
		a.ss = false
		payload[0] += 1
	}

	a.messageQueue <- NewCMessage(payload, false)
}
