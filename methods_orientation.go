package mpost

import (
	"github.com/hard-soft-ware/mpost/acceptor"
	"github.com/hard-soft-ware/mpost/enum"
)

////////////////////////////////////

func (a *MpostObj) GetCapOrientationExt() bool {
	a.Log.Method("GetCapOrientationExt", nil)
	return acceptor.Cap.OrientationExt
}

func (a *MpostObj) GetEscrowOrientation() enum.OrientationType {
	a.Log.Method("GetEscrowOrientation", nil)
	if acceptor.Cap.OrientationExt {
		return acceptor.EscrowOrientation
	}
	return enum.OrientationUnknownOrientation
}

func (a *MpostObj) GetOrientationControl() enum.OrientationControlType {
	a.Log.Method("GetOrientationControl", nil)
	return acceptor.OrientationCtl
}

func (a *MpostObj) SetOrientationControl(v enum.OrientationControlType) {
	a.Log.Method("SetOrientationControl", nil)
	acceptor.OrientationCtl = v
}

func (a *MpostObj) GetOrientationCtlExt() enum.OrientationControlType {
	a.Log.Method("GetOrientationCtlExt", nil)
	return acceptor.OrientationCtlExt
}

func (a *MpostObj) SetOrientationCtlExt(v enum.OrientationControlType) {
	a.Log.Method("SetOrientationCtlExt", nil)
	acceptor.OrientationCtlExt = v
}
