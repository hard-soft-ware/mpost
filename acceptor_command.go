package mpost

import (
	"errors"
	"time"
)

////////////////////////////////////

func (a *CAcceptor) SendSynchronousCommand(payload []byte) ([]byte, error) {
	a.log.Debug().Bytes("payload", payload).Msg("SendCommand")
	a.messageQueue <- NewCMessage(payload, true)

	select {
	case reply := <-a.replyQueue:
		a.log.Debug().Bytes("payload", reply).Msg("Reply queued")
		return reply, nil
	case <-time.After(30 * time.Second):
		return nil, errors.New("timeout waiting for response")
	}
}
