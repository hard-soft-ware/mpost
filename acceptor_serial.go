package mpost

import (
	"fmt"
	"github.com/hard-soft-ware/mpost/acceptor"
	"github.com/hard-soft-ware/mpost/consts"
	"github.com/hard-soft-ware/mpost/enum"
	"go.bug.st/serial"
	"time"
)

////////////////////////////////////

func (a *CAcceptor) Open(portName string, powerUp enum.PowerUpType) error {
	lg := a.log.New("OpenSerial")

	if acceptor.Connected {
		lg.Msg("already connected")
		return nil
	}

	acceptor.Device.PortName = portName
	acceptor.Device.PowerUp = powerUp

	err := a.OpenPort(lg)
	if err != nil {
		a.log.Err("failed to open serial port", err)
		return err
	}

	go a.MessageLoopThread()
	go a.OpenThread()

	return nil
}

func (a *CAcceptor) OpenPort(lg *LogStruct) error {
	mode := &serial.Mode{
		BaudRate: 9600,
		DataBits: 7,
		Parity:   serial.EvenParity,
		StopBits: serial.OneStopBit,
	}

	port, err := serial.Open(acceptor.Device.PortName, mode)
	if err != nil {
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

	acceptor.Connected = true
	lg.Msg("Connected")
	return nil
}

func (a *CAcceptor) Close() {

	if !acceptor.Connected {
		a.stopOpenThread = true
		return
	}

	if acceptor.Enable.Acceptance {
		acceptor.Enable.Acceptance = false
	}

	if a.dataLinkLayer != nil {
		a.log.Msg(fmt.Sprintf("IdenticalCommandAndReplyCount: %d", a.dataLinkLayer.IdenticalCommandAndReplyCount))
	}

	a.stopWorkerThread = true
	a.workerThread.Wait()

	a.port.Close()
	a.port = nil
	acceptor.Connected = false
}

////

func (a *CAcceptor) QueryDeviceCapabilities(lg *LogStruct) {
	if !a.isQueryDeviceCapabilitiesSupported {
		return
	}

	payload := []byte{consts.CmdAuxiliary.Byte(), 0x00, 0x00, consts.CmdAuxDeviceCapabilities.Byte()}
	reply, err := a.SendSynchronousCommand(payload)

	if len(reply) < 4 {
		a.log.Err("Reply too short, unable to process.", err)
		return
	}

	if reply[3]&0x01 != 0 {
		acceptor.Cap.PupExt = true
	}
	if reply[3]&0x02 != 0 {
		acceptor.Cap.OrientationExt = true
	}
	if reply[3]&0x04 != 0 {
		acceptor.Cap.ApplicationID = true
		acceptor.Cap.VariantID = true
	}
	if reply[3]&0x08 != 0 {
		acceptor.Cap.BNFStatus = true
	}
	if reply[3]&0x10 != 0 {
		acceptor.Cap.TestDoc = true
	}
}
