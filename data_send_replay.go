package mpost

import (
	"bufio"
	"errors"
	"io"
	"time"
)

////////////////////////////////////

func (dl *CDataLinkLayer) SendPacket(payload []byte) {
	payloadLength := len(payload)
	commandLength := payloadLength + 4 // STX + Length char + ETX + Checksum

	command := make([]byte, 0, commandLength)
	command = append(command, STX)
	command = append(command, byte(commandLength))

	command = append(command, payload...)
	command[2] |= dl.AckToggleBit

	command = append(command, ETX)
	command = append(command, dl.ComputeCheckSum(command))

	dl.CurrentCommand = command
	dl.EchoDetect = command

	dl.log.Debug().Str("data", byteSliceToString(command)).Msg("SERIAL SEND")
	n, err := dl.Acceptor.port.Write(command)
	if err != nil || n == 0 {
		dl.log.Debug().Err(err).Msg("Failed to write to port")

		dl.Acceptor.port.Close()
		dl.Acceptor.OpenPort(dl.log)
	}
}

func (dl *CDataLinkLayer) WaitForQuiet() {
	for {
		buf := make([]byte, 1)
		timeout := 20 * time.Millisecond

		dl.Acceptor.port.SetReadTimeout(timeout)

		_, err := dl.Acceptor.port.Read(buf)
		if err != nil {
			return
		}
	}
}

//

func (dl *CDataLinkLayer) ReceiveReply() ([]byte, error) {
	reply := []byte{}

	timeout := dl.Acceptor.transactionTimeout
	if dl.Acceptor.deviceState == DownloadStart || dl.Acceptor.deviceState == Downloading {
		timeout = dl.Acceptor.downloadTimeout
	}

	dl.Acceptor.port.SetReadTimeout(timeout)

	reader := bufio.NewReader(dl.Acceptor.port)
	stxAndLength := make([]byte, 2)
	bytesRead, err := io.ReadFull(reader, stxAndLength)
	if err != nil {
		return nil, err
	}
	if bytesRead < 2 {
		return nil, errors.New("received insufficient bytes")
	}
	if stxAndLength[0] != STX {
		return nil, errors.New("invalid STX received")
	}

	reply = append(reply, stxAndLength...)

	length := int(stxAndLength[1])
	checkByte := stxAndLength[1]

	bytesRemaining := length - 2
	payloadAndETXBuffer := make([]byte, 128)

	for bytesRemaining > 0 {
		n, err := io.ReadFull(reader, payloadAndETXBuffer[:bytesRemaining])
		if err != nil {
			return nil, err
		}

		bytesRemaining -= n
		reply = append(reply, payloadAndETXBuffer[:n]...)

		for i := 0; i < n; i++ {
			if len(reply) < length-1 {
				checkByte ^= payloadAndETXBuffer[i]
			}
		}
	}

	dl.log.Debug().Str("data", byteSliceToString(reply)).Msg("SERIAL READ")
	return reply, nil
}

////

func (dl *CDataLinkLayer) ReplyAcked(reply []byte) bool {
	if len(reply) < 3 {
		return false
	}

	if (reply[2] & ACKMask) == dl.AckToggleBit {
		dl.AckToggleBit ^= 0x01 // Переключаем бит подтверждения

		dl.NakCount = 0

		return true
	} else {
		dl.NakCount++

		// Если получено 8 последовательных NAK, принудительно переключаем бит
		if dl.NakCount == 8 {
			dl.AckToggleBit ^= 0x01
			dl.NakCount = 0
		}

		return false
	}
}

func (dl *CDataLinkLayer) ProcessReply(reply []byte) {
	if len(reply) < 3 {
		return
	}

	ctl := reply[2]

	if (ctl & 0x70) == 0x20 {
		dl.ProcessStandardOmnibusReply(reply)
	}

	if (ctl & 0x70) == 0x50 {
		dl.Acceptor.deviceState = DownloadRestart
	}

	if (ctl & 0x70) == 0x70 {
		subType := reply[3]
		switch subType {
		case 0x01:
			dl.ProcessExtendedOmnibusBarCodeReply(reply)
		case 0x02:
			dl.ProcessExtendedOmnibusExpandedNoteReply(reply)
			if dl.Acceptor.deviceState == Escrow || (dl.Acceptor.deviceState == Stacked && !dl.Acceptor.wasDocTypeSetOnEscrow) {
				if dl.Acceptor.capOrientationExt {
					switch dl.Acceptor.orientationCtlExt {
					case OneWay:
						if dl.Acceptor.escrowOrientation != RightUp {
							dl.EscrowReturn()
						}
					case TwoWay:
						if dl.Acceptor.escrowOrientation != RightUp && dl.Acceptor.escrowOrientation != LeftUp {
							dl.EscrowReturn()
						}
					case FourWay:
						// Accept all orientations.
					}
				}
			}
		case 0x04:
			dl.ProcessExtendedOmnibusExpandedCouponReply(reply)
		}
		dl.RaiseEvents()
	}

	if dl.Acceptor.deviceState == Escrow && dl.Acceptor.autoStack {
		dl.EscrowStack()
		dl.Acceptor.shouldRaiseEscrowEvent = false
	}

	if dl.Acceptor.deviceState != Escrow && dl.Acceptor.deviceState != Stacking {
		dl.Acceptor.wasDocTypeSetOnEscrow = false
	}
}
