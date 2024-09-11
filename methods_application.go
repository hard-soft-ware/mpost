package mpost

import (
	"github.com/hard-soft-ware/mpost/acceptor"
	"github.com/hard-soft-ware/mpost/consts"
)

////////////////////////////////////

func (a *MpostObj) GetApplicationID() string {
	a.Log.Method("GetApplicationID", nil)

	err := acceptor.Verify(acceptor.Cap.ApplicationID, "GetApplicationID")
	if err != nil {
		a.Log.Err("GetApplicationID", err)
		return ""
	}

	payload := []byte{consts.CmdAuxiliary.Byte(), 0, 0, consts.CmdAuxAcceptorApplicationID.Byte()}

	reply, err := a.SendSynchronousCommand(payload)
	if err != nil {
		a.Log.Err("GetApplicationID", err)
		return ""
	}

	if len(reply) == 14 {
		s := string(reply[3:12])
		return s
	}

	return ""
}

func (a *MpostObj) GetApplicationPN() string {
	a.Log.Method("GetApplicationPN", nil)

	err := acceptor.Verify(acceptor.Cap.ApplicationPN, "ApplicationPN")
	if err != nil {
		a.Log.Err("ApplicationPN", err)
		return ""
	}

	payload := []byte{consts.CmdAuxiliary.Byte(), 0, 0, consts.CmdAuxAcceptorApplicationPartNumber.Byte()}

	reply, err := a.SendSynchronousCommand(payload)
	if err != nil {
		a.Log.Err("GetApplicationPN", err)
		return ""
	}

	if len(reply) == 14 {
		s := string(reply[3:12])
		return s
	}

	return ""
}

func (a *MpostObj) GetCapApplicationID() bool {
	a.Log.Method("GetCapApplicationID", nil)
	return acceptor.Cap.ApplicationID
}

func (a *MpostObj) GetCapApplicationPN() bool {
	a.Log.Method("GetCapApplicationPN", nil)
	return acceptor.Cap.ApplicationPN
}
