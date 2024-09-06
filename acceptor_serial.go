package mpost

import (
	"fmt"
	"github.com/hard-soft-ware/mpost/consts"
	"github.com/hard-soft-ware/mpost/enum"
	"go.bug.st/serial"
	"strings"
	"time"
)

func byteSliceToString(b []byte) string {
	var sb strings.Builder
	for i, byteVal := range b {
		if i > 0 {
			sb.WriteString(" ")
		}
		fmt.Fprintf(&sb, "%02X", byteVal)
	}
	return sb.String()
}

////////////////////////////////////

func (a *CAcceptor) Open(portName string, powerUp enum.PowerUpType) error {
	lg := a.log.NewLog("OpenSerial")

	if a.connected {
		lg.Debug().Msg("already connected")
		return nil
	}

	a.devicePortName = portName
	a.devicePowerUp = powerUp

	err := a.OpenPort(lg)
	if err != nil {
		lg.Debug().Err(err).Msg("failed to open serial port")
		return err
	}

	go a.MessageLoopThread()
	go a.OpenThread()

	return nil
}

func (a *CAcceptor) OpenPort(lg *LogGlobalStruct) error {
	mode := &serial.Mode{
		BaudRate: 9600,
		DataBits: 7,
		Parity:   serial.EvenParity,
		StopBits: serial.OneStopBit,
	}

	port, err := serial.Open(a.devicePortName, mode)
	if err != nil {
		lg.Debug().Err(err).Msg("failed to open serial port")
		return err
	}

	port.SetReadTimeout(100 * time.Millisecond)
	port.ResetInputBuffer()

	port.SetDTR(false)
	port.SetRTS(true)
	time.Sleep(100 * time.Millisecond)

	port.SetDTR(true)
	port.SetRTS(false)
	time.Sleep(5 * time.Millisecond)

	port.ResetInputBuffer()
	a.port = port

	a.connected = true
	lg.Debug().Msg("Connected")
	return nil
}

func (a *CAcceptor) Close() {

	if !a.connected {
		a.stopOpenThread = true
		return
	}

	if a.enableAcceptance {
		a.enableAcceptance = false
	}

	if a.dataLinkLayer != nil {
		a.dataLinkLayer.FlushIdenticalTransactionsToLog()
	}

	a.stopWorkerThread = true
	a.workerThread.Wait()

	a.port.Close()
	a.port = nil
	a.connected = false
}

////

func (a *CAcceptor) QueryDeviceCapabilities(lg *LogGlobalStruct) {
	if !a.isQueryDeviceCapabilitiesSupported {
		return
	}

	payload := []byte{consts.CmdAuxiliary.Byte(), 0x00, 0x00, consts.CmdAuxDeviceCapabilities.Byte()}
	reply, err := a.SendSynchronousCommand(payload)

	if len(reply) < 4 {
		lg.Debug().Err(err).Msg("Reply too short, unable to process.")
		return
	}

	if reply[3]&0x01 != 0 {
		a.capPupExt = true
	}
	if reply[3]&0x02 != 0 {
		a.capOrientationExt = true
	}
	if reply[3]&0x04 != 0 {
		a.capApplicationID = true
		a.capVariantID = true
	}
	if reply[3]&0x08 != 0 {
		a.capBNFStatus = true
	}
	if reply[3]&0x10 != 0 {
		a.capTestDoc = true
	}
}
