package mpost

import (
	"errors"
	"time"
)

////////////////////////////////////

func (a *CAcceptor) ConstructOmnibusCommand(payload []byte, controlCode byte, data0Index int) {
	payload[0] = controlCode

	if a.enableBookmarks && a.enableAcceptance && a.deviceState != Calibrating {
		payload[0] |= 0x20
	}

	data0 := byte(0)

	if a.enableAcceptance && a.deviceState != Calibrating {
		if a.expandedNoteReporting {
			data0 |= 0x7F
		} else {
			if len(a.billTypeEnables) == 0 {
				data0 |= 0x7F
			} else {
				for i, enable := range a.billTypeEnables {
					if enable {
						data0 |= 1 << i
					}
				}
			}
		}
	}

	data1 := byte(0) // Ignore bit 0 since we are not supporting special interrupt mode.

	if a.highSecurity {
		data1 |= 0x02
	}

	switch a.orientationCtl {
	case TwoWay:
		data1 |= 0x04
	case FourWay:
		data1 |= 0x0C
	}

	data1 |= 0x10 // Always enable escrow mode.

	data2 := byte(0)

	if a.enableNoPush {
		data2 |= 0x01
	}

	if a.enableBarCodes && a.enableAcceptance && a.deviceState != Calibrating {
		data2 |= 0x02
	}

	switch a.devicePowerUp {
	case B:
		data2 |= 0x04
	case C:
		data2 |= 0x0C
	}

	if a.expandedNoteReporting {
		data2 |= 0x10
	}

	if a.enableCouponExt && a.capCouponExt {
		data2 |= 0x20
	}

	payload[data0Index] = data0
	payload[data0Index+1] = data1
	payload[data0Index+2] = data2
}

////

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
