package mpost

import (
	"errors"
	"time"
)

////////////////////////////////////

func (a *CAcceptor) SendSynchronousCommand(payload []byte) ([]byte, error) {
	a.messageQueue <- NewCMessage(payload, true)

	select {
	case reply := <-a.replyQueue:
		return reply, nil
	case <-time.After(30 * time.Second):
		return nil, errors.New("timeout waiting for response")
	}

	return nil, errors.New("invalid response")
}
