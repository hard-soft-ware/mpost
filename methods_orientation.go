package mpost

import (
	"github.com/hard-soft-ware/mpost/acceptor"
	"github.com/hard-soft-ware/mpost/enum"
)

////////////////////////////////////

func (m *MethodsObj) GetCapOrientationExt() bool {
	m.a.Log.Method("GetCapOrientationExt", nil)
	return acceptor.Cap.OrientationExt
}

func (m *MethodsObj) GetEscrowOrientation() enum.OrientationType {
	m.a.Log.Method("GetEscrowOrientation", nil)
	if acceptor.Cap.OrientationExt {
		return acceptor.EscrowOrientation
	}
	return enum.OrientationUnknownOrientation
}

func (m *MethodsObj) GetOrientationControl() enum.OrientationControlType {
	m.a.Log.Method("GetOrientationControl", nil)
	return acceptor.OrientationCtl
}

func (m *MethodsObj) SetOrientationControl(v enum.OrientationControlType) {
	m.a.Log.Method("SetOrientationControl", nil)
	acceptor.OrientationCtl = v
}

func (m *MethodsObj) GetOrientationCtlExt() enum.OrientationControlType {
	m.a.Log.Method("GetOrientationCtlExt", nil)
	return acceptor.OrientationCtlExt
}

func (m *MethodsObj) SetOrientationCtlExt(v enum.OrientationControlType) {
	m.a.Log.Method("SetOrientationCtlExt", nil)
	acceptor.OrientationCtlExt = v
}
