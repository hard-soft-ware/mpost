package mpost

import (
	"errors"
	"github.com/hard-soft-ware/mpost/acceptor"
	"github.com/hard-soft-ware/mpost/consts"
)

////////////////////////////////////

func (a *MpostObj) GetCapCashBoxTotal() bool {
	a.Log.Method("GetCapCashBoxTotal", nil)
	return acceptor.Cap.CashBoxTotal
}

func (a *MpostObj) GetCashBoxAttached() bool {
	a.Log.Method("GetCashBoxAttached", nil)
	return acceptor.Cash.BoxAttached
}

func (a *MpostObj) GetCashBoxFull() bool {
	a.Log.Method("GetCashBoxFull", nil)
	return acceptor.Cash.BoxFull
}

func (a *MpostObj) GetCashBoxTotal() int {
	a.Log.Method("GetCashBoxTotal", nil)

	err := acceptor.Verify(acceptor.Cap.CashBoxTotal, "GetCashBoxTotal")
	if err != nil {
		a.Log.Err("GetCashBoxTotal", err)
		return 0
	}

	payload := []byte{consts.CmdOmnibus.Byte(), 0x7F, 0x3C, 0x02}

	reply, err := a.SendSynchronousCommand(payload)
	if err != nil {
		a.Log.Err("GetCashBoxTotal", err)
		return 0
	}

	if len(reply) < 9 {
		return 0
	}

	total := int(reply[3]&0x0F)<<20 |
		int(reply[4]&0x0F)<<16 |
		int(reply[5]&0x0F)<<12 |
		int(reply[6]&0x0F)<<8 |
		int(reply[7]&0x0F)<<4 |
		int(reply[8]&0x0F)

	return total
}

func (a *MpostObj) ClearCashBoxTotal() (err error) {
	a.Log.Method("ClearCashBoxTotal", nil)

	if !acceptor.Connected {
		err = errors.New("ClearCashBoxTotal called when not connected")
		a.Log.Err("ClearCashBoxTotal", err)
		return
	}

	payload := []byte{consts.CmdCalibrate.Byte(), 0x00, 0x00, consts.CmdAuxCashBoxTotal.Byte()}

	reply, err := a.SendSynchronousCommand(payload)
	if err != nil {
		a.Log.Err("ClearCashBoxTotal", err)
		return
	}

	a.dataLinkLayer.ProcessReply(reply)
	return
}
