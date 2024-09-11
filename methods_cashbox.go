package mpost

import (
	"errors"
	"github.com/hard-soft-ware/mpost/acceptor"
	"github.com/hard-soft-ware/mpost/consts"
)

////////////////////////////////////

func (m *MethodsObj) GetCapCashBoxTotal() bool {
	m.a.Log.Method("GetCapCashBoxTotal", nil)
	return acceptor.Cap.CashBoxTotal
}

func (m *MethodsObj) GetCashBoxAttached() bool {
	m.a.Log.Method("GetCashBoxAttached", nil)
	return acceptor.Cash.BoxAttached
}

func (m *MethodsObj) GetCashBoxFull() bool {
	m.a.Log.Method("GetCashBoxFull", nil)
	return acceptor.Cash.BoxFull
}

func (m *MethodsObj) GetCashBoxTotal() int {
	m.a.Log.Method("GetCashBoxTotal", nil)

	err := acceptor.Verify(acceptor.Cap.CashBoxTotal, "GetCashBoxTotal")
	if err != nil {
		m.a.Log.Err("GetCashBoxTotal", err)
		return 0
	}

	payload := []byte{consts.CmdOmnibus.Byte(), 0x7F, 0x3C, 0x02}

	reply, err := m.a.SendSynchronousCommand(payload)
	if err != nil {
		m.a.Log.Err("GetCashBoxTotal", err)
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

func (m *MethodsObj) ClearCashBoxTotal() (err error) {
	m.a.Log.Method("ClearCashBoxTotal", nil)

	if !acceptor.Connected {
		err = errors.New("ClearCashBoxTotal called when not connected")
		m.a.Log.Err("ClearCashBoxTotal", err)
		return
	}

	payload := []byte{consts.CmdCalibrate.Byte(), 0x00, 0x00, consts.CmdAuxCashBoxTotal.Byte()}

	reply, err := m.a.SendSynchronousCommand(payload)
	if err != nil {
		m.a.Log.Err("ClearCashBoxTotal", err)
		return
	}

	m.a.dataLinkLayer.ProcessReply(reply)
	return
}
