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

func (a *MpostObj) SendSynchronousCommand(payload []byte) ([]byte, error) {

	a.messageQueue <- NewCMessage(payload, true)

	select {
	case <-a.Ctx.Done():
		a.Close()
		return nil, errors.New("close from context")
	case reply := <-a.replyQueue:
		return reply, nil
	case <-time.After(30 * time.Second):
		return nil, errors.New("timeout waiting for response")
	}

	return nil, errors.New("invalid response")
}

func (a *MpostObj) SendAsynchronousCommand(payload []byte) {
	a.messageQueue <- NewCMessage(payload, false)
}
