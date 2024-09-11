package mpost

import (
	"github.com/hard-soft-ware/mpost/acceptor"
	"github.com/hard-soft-ware/mpost/consts"
)

////////////////////////////////////

type MethodsApplicationObj struct {
	a   *MpostObj
	Get MethodApplicationGetObj
}

type MethodApplicationGetObj struct{ a *MpostObj }

func (m *MethodsObj) newApplication() *MethodsApplicationObj {
	obj := MethodsApplicationObj{}

	obj.a = m.a
	obj.Get.a = m.a

	return &obj
}

////////////////

func (m *MethodApplicationGetObj) ID() string {
	m.a.Log.Method("GetApplicationID", nil)

	err := acceptor.Verify(acceptor.Cap.ApplicationID, "GetApplicationID")
	if err != nil {
		m.a.Log.Err("GetApplicationID", err)
		return ""
	}

	payload := []byte{consts.CmdAuxiliary.Byte(), 0, 0, consts.CmdAuxAcceptorApplicationID.Byte()}

	reply, err := m.a.SendSynchronousCommand(payload)
	if err != nil {
		m.a.Log.Err("GetApplicationID", err)
		return ""
	}

	if len(reply) == 14 {
		s := string(reply[3:12])
		return s
	}

	return ""
}

func (m *MethodApplicationGetObj) PN() string {
	m.a.Log.Method("GetApplicationPN", nil)

	err := acceptor.Verify(acceptor.Cap.ApplicationPN, "ApplicationPN")
	if err != nil {
		m.a.Log.Err("ApplicationPN", err)
		return ""
	}

	payload := []byte{consts.CmdAuxiliary.Byte(), 0, 0, consts.CmdAuxAcceptorApplicationPartNumber.Byte()}

	reply, err := m.a.SendSynchronousCommand(payload)
	if err != nil {
		m.a.Log.Err("GetApplicationPN", err)
		return ""
	}

	if len(reply) == 14 {
		s := string(reply[3:12])
		return s
	}

	return ""
}

//

func (m *MethodApplicationGetObj) CapID() bool {
	m.a.Log.Method("GetCapApplicationID", nil)
	return acceptor.Cap.ApplicationID
}

func (m *MethodApplicationGetObj) CapPN() bool {
	m.a.Log.Method("GetCapApplicationPN", nil)
	return acceptor.Cap.ApplicationPN
}
