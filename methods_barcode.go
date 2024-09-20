package mpost

import "github.com/hard-soft-ware/mpost/acceptor"

////////////////////////////////////

func (m *MethodsObj) GetBarCode() string {
	m.a.Log.Method("GetBarCode", nil)
	return acceptor.BarCode
}

func (m *MethodsObj) GetEnableBarCodes() bool {
	m.a.Log.Method("GetEnableBarCodes", nil)
	return acceptor.Enable.BarCodes
}

//

func (m *MethodsObj) GetCapBarCodes() bool {
	m.a.Log.Method("GetCapBarCodes", nil)
	return acceptor.Cap.BarCodes
}

func (m *MethodsObj) GetCapBarCodesExt() bool {
	m.a.Log.Method("GetCapBarCodesExt", nil)
	return acceptor.Cap.BarCodesExt
}

////

func (m *MethodsObj) SetEnableBarCodes(v bool) {
	m.a.Log.Method("SetEnableBarCodes", v)
	acceptor.Enable.BarCodes = v
}
