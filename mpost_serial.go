package mpost

import (
	"github.com/hard-soft-ware/mpost/acceptor"
	"github.com/hard-soft-ware/mpost/consts"
	"github.com/hard-soft-ware/mpost/enum"
	"github.com/hard-soft-ware/mpost/hook"
	"github.com/hard-soft-ware/mpost/serial"
)

////////////////////////////////////

func (a *MpostObj) Open(portName string, powerUp enum.PowerUpType) error {
	if acceptor.Connected {
		a.Log.Msg("already connected")
		return nil
	}

	acceptor.Device.PowerUp = powerUp

	port, err := serial.Open(portName, &acceptor.Connected)
	if err != nil {
		a.Log.Err("failed to open serial port", err)
		return err
	}
	a.port = port
	a.Log.Msg("Serial Open")

	go a.messageLoopThread()
	go a.openThread()

	return nil
}

func (a *MpostObj) Close() {
	hook.Raise.Disconnected()
	a.port.Close()

	defer a.CtxCancel()

	if !acceptor.Connected {
		acceptor.StopOpenThread = true
		return
	}

	if acceptor.Enable.Acceptance {
		acceptor.Enable.Acceptance = false
	}

	acceptor.StopWorkerThread = true

	a.port = nil
	acceptor.Connected = false
	a.Log.Msg("Close")
}

////

func (a *MpostObj) queryDeviceCapabilities() {
	if !acceptor.IsQueryDeviceCapabilitiesSupported {
		return
	}

	payload := []byte{consts.CmdAuxiliary.Byte(), 0x00, 0x00, consts.CmdAuxDeviceCapabilities.Byte()}
	reply, err := a.SendSynchronousCommand(payload)

	if len(reply) < 4 {
		a.Log.Err("Reply too short, unable to process.", err)
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
