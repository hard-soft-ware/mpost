package mpost

import "github.com/hard-soft-ware/mpost/acceptor"

////////////////////////////////////

func (a *MpostObj) GetBarCode() string {
	a.Log.Method("GetBarCode", nil)
	return acceptor.BarCode
}

func (a *MpostObj) GetCapBarCodes() bool {
	a.Log.Method("GetCapBarCodes", nil)
	return acceptor.Cap.BarCodes
}

func (a *MpostObj) GetCapBarCodesExt() bool {
	a.Log.Method("GetCapBarCodesExt", nil)
	return acceptor.Cap.BarCodesExt
}

func (a *MpostObj) GetEnableBarCodes() bool {
	a.Log.Method("GetEnableBarCodes", nil)
	return acceptor.Enable.BarCodes
}

func (a *MpostObj) SetEnableBarCodes(v bool) {
	a.Log.Method("SetEnableBarCodes", nil)
	acceptor.Enable.BarCodes = v
}
