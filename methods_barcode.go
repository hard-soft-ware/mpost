package mpost

import "github.com/hard-soft-ware/mpost/acceptor"

////////////////////////////////////

type MethodsBarCodeObj struct {
	a   *MpostObj
	Get MethodBarCodeGetObj
	Set MethodBarCodeSetObj
}

type MethodBarCodeGetObj struct{ a *MpostObj }

type MethodBarCodeSetObj struct{ a *MpostObj }

func (m *MethodsObj) newBarCode() *MethodsBarCodeObj {
	obj := MethodsBarCodeObj{}

	obj.a = m.a
	obj.Get.a = m.a
	obj.Set.a = m.a

	return &obj
}

////////////////

func (m *MethodBarCodeGetObj) Self() string {
	m.a.Log.Method("GetBarCode", nil)
	return acceptor.BarCode
}

func (m *MethodBarCodeGetObj) Enable() bool {
	m.a.Log.Method("GetEnableBarCodes", nil)
	return acceptor.Enable.BarCodes
}

//

func (m *MethodBarCodeGetObj) Cap() bool {
	m.a.Log.Method("GetCapBarCodes", nil)
	return acceptor.Cap.BarCodes
}

func (m *MethodBarCodeGetObj) CapExt() bool {
	m.a.Log.Method("GetCapBarCodesExt", nil)
	return acceptor.Cap.BarCodesExt
}

////

func (m *MethodBarCodeSetObj) Enable(v bool) {
	m.a.Log.Method("SetEnableBarCodes", nil)
	acceptor.Enable.BarCodes = v
}
