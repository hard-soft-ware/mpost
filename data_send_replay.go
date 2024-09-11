package mpost

import (
	"bufio"
	"errors"
	"github.com/hard-soft-ware/mpost/acceptor"
	"github.com/hard-soft-ware/mpost/command"
	"github.com/hard-soft-ware/mpost/consts"
	"github.com/hard-soft-ware/mpost/enum"
	"github.com/hard-soft-ware/mpost/hook"
	"io"
)

////////////////////////////////////

func (dl *CDataLinkLayer) SendPacket(payload []byte) {
	send := command.CreateMsg(payload)

	dl.CurrentCommand = send
	dl.EchoDetect = send

	dl.Acceptor.Log.SerialSend(command.Parse(send))
	n, err := dl.Acceptor.port.Write(send)
	if err != nil || n == 0 {
		dl.Acceptor.Log.Err("Failed to write to port", err)
		err = dl.Acceptor.port.Restart()
		if err != nil {
			dl.Acceptor.Log.Err("Failed restart to port", err)
		}
	}
}

//

func (dl *CDataLinkLayer) ReceiveReply() ([]byte, error) {
	reply := []byte{}

	timeout := acceptor.Timeout.Transaction
	if acceptor.Device.State == enum.StateDownloadStart || acceptor.Device.State == enum.StateDownloading {
		timeout = acceptor.Timeout.Download
	}

	dl.Acceptor.port.SetTimeout(timeout)

	reader := bufio.NewReader(dl.Acceptor.port.Port())
	stxAndLength := make([]byte, 2)
	bytesRead, err := io.ReadFull(reader, stxAndLength)
	if err != nil {
		return nil, err
	}
	if bytesRead < 2 {
		return nil, errors.New("received insufficient bytes")
	}
	if stxAndLength[0] != consts.DataSTX.Byte() {
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

	dl.Acceptor.Log.SerialRead(command.Parse(reply))
	return reply, nil
}

////

func (dl *CDataLinkLayer) ProcessReply(reply []byte) {
	if len(reply) < 3 {
		return
	}

	ctl := reply[2]

	if (ctl & 0x70) == 0x20 {
		dl.ProcessStandardOmnibusReply(reply)
	}

	if (ctl & 0x70) == 0x50 {
		acceptor.Device.State = enum.StateDownloadRestart
	}

	if (ctl & 0x70) == 0x70 {
		subType := reply[3]
		switch subType {
		case 0x01:
			dl.ProcessExtendedOmnibusBarCodeReply(reply)
		case 0x02:
			dl.ProcessExtendedOmnibusExpandedNoteReply(reply)
			if acceptor.Device.State == enum.StateEscrow || (acceptor.Device.State == enum.StateStacked && !acceptor.WasDocTypeSetOnEscrow) {
				if acceptor.Cap.OrientationExt {
					switch acceptor.OrientationCtlExt {
					case enum.OrientationControlOneWay:
						if acceptor.EscrowOrientation != enum.OrientationRightUp {
							dl.EscrowReturn()
						}
					case enum.OrientationControlTwoWay:
						if acceptor.EscrowOrientation != enum.OrientationRightUp && acceptor.EscrowOrientation != enum.OrientationLeftUp {
							dl.EscrowReturn()
						}
					case enum.OrientationControlFourWay:
						// Accept all orientations.
					}
				}
			}
		case 0x04:
			dl.ProcessExtendedOmnibusExpandedCouponReply(reply)
		}
		dl.RaiseEvents()
	}

	if acceptor.Device.State == enum.StateEscrow && acceptor.AutoStack {
		dl.EscrowStack()
		hook.Escrow = false
	}

	if acceptor.Device.State != enum.StateEscrow && acceptor.Device.State != enum.StateStacking {
		acceptor.WasDocTypeSetOnEscrow = false
	}
}
