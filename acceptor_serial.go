package mpost

import (
	"fmt"
	"github.com/hard-soft-ware/mpost/acceptor"
	"github.com/hard-soft-ware/mpost/consts"
	"github.com/hard-soft-ware/mpost/enum"
	"github.com/hard-soft-ware/mpost/serial"
)

////////////////////////////////////

func (a *CAcceptor) Open(portName string, powerUp enum.PowerUpType) error {
	if acceptor.Connected {
		a.log.Msg("already connected")
		return nil
	}

	acceptor.Device.PowerUp = powerUp

	port, err := serial.Open(portName, &acceptor.Connected)
	if err != nil {
		a.log.Err("failed to open serial port", err)
		return err
	}
	a.port = port
	a.log.Msg("Serial Open")

	go a.MessageLoopThread()
	go a.OpenThread()

	return nil
}

func (a *CAcceptor) Close() {

	if !acceptor.Connected {
		acceptor.StopOpenThread = true
		return
	}

	if acceptor.Enable.Acceptance {
		acceptor.Enable.Acceptance = false
	}

	if a.dataLinkLayer != nil {
		a.log.Msg(fmt.Sprintf("IdenticalCommandAndReplyCount: %d", a.dataLinkLayer.IdenticalCommandAndReplyCount))
	}

	acceptor.StopWorkerThread = true
	a.workerThread.Wait()

	a.port.Close()
	a.port = nil
	acceptor.Connected = false
	a.log.Msg("Close")
}

////

func (a *CAcceptor) QueryDeviceCapabilities() {
	if !acceptor.IsQueryDeviceCapabilitiesSupported {
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
