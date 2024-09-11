package mpost

import (
	"github.com/hard-soft-ware/mpost/acceptor"
	"github.com/hard-soft-ware/mpost/enum"
)

////////////////////////////////////

type MethodsOrientationObj struct {
	a   *MpostObj
	Get MethodsOrientationGetObj
	Set MethodsOrientationSetObj
}

type MethodsOrientationGetObj struct{ a *MpostObj }
type MethodsOrientationSetObj struct{ a *MpostObj }

func (m *MethodsObj) newOrientation() *MethodsOrientationObj {
	obj := MethodsOrientationObj{}

	obj.a = m.a
	obj.Get.a = m.a
	obj.Set.a = m.a

	return &obj
}

////////////////

func (m *MethodsOrientationGetObj) Escrow() enum.OrientationType {
	m.a.Log.Method("GetEscrowOrientation", nil)
	if acceptor.Cap.OrientationExt {
		return acceptor.EscrowOrientation
	}
	return enum.OrientationUnknownOrientation
}

func (m *MethodsOrientationGetObj) Control() enum.OrientationControlType {
	m.a.Log.Method("GetOrientationControl", nil)
	return acceptor.OrientationCtl
}

func (m *MethodsOrientationGetObj) ControlEscrow() enum.OrientationControlType {
	m.a.Log.Method("GetOrientationCtlExt", nil)
	return acceptor.OrientationCtlExt
}

//

func (m *MethodsOrientationGetObj) CapEscrow() bool {
	m.a.Log.Method("GetCapOrientationExt", nil)
	return acceptor.Cap.OrientationExt
}

////

func (m *MethodsOrientationSetObj) Control(v enum.OrientationControlType) {
	m.a.Log.Method("SetOrientationControl", nil)
	acceptor.OrientationCtl = v
}

func (m *MethodsOrientationSetObj) ControlEscrow(v enum.OrientationControlType) {
	m.a.Log.Method("SetOrientationCtlExt", nil)
	acceptor.OrientationCtlExt = v
}
