package mpost

import (
	"errors"
	"time"
)

////////////////////////////////////

func (a *CAcceptor) SendSynchronousCommand(payload []byte) ([]byte, error) {

	//todo неведомая но очень нужная херня
	if !a.ss {
		a.ss = true
	} else {
		a.ss = false
		payload[0] += 1
	}

	a.messageQueue <- NewCMessage(payload, true)

	select {
	case reply := <-a.replyQueue:
		return reply, nil
	case <-time.After(30 * time.Second):
		return nil, errors.New("timeout waiting for response")
	}

	return nil, errors.New("invalid response")
}
