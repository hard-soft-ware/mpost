package mpost

import "github.com/hard-soft-ware/mpost/acceptor"

////////////////////////////////////

type MethodsEnableObj struct {
	a   *MpostObj
	Get MethodsEnableGetObj
	Set MethodsEnableSetObj
}

type MethodsEnableGetObj struct{ a *MpostObj }
type MethodsEnableSetObj struct{ a *MpostObj }

func (m *MethodsObj) newEnable() *MethodsEnableObj {
	obj := MethodsEnableObj{}

	obj.a = m.a
	obj.Get.a = m.a
	obj.Set.a = m.a

	return &obj
}

////////////////

func (m *MethodsEnableGetObj) Acceptance() bool {
	m.a.Log.Method("GetEnableAcceptance", nil)
	return acceptor.Enable.Acceptance
}

func (m *MethodsEnableGetObj) Bookmarks() bool {
	m.a.Log.Method("GetEnableBookmarks", nil)
	return acceptor.Enable.Bookmarks
}

func (m *MethodsEnableGetObj) NoPush() bool {
	m.a.Log.Method("GetEnableNoPush", nil)
	return acceptor.Enable.NoPush
}

//

func (m *MethodsEnableSetObj) Acceptance(v bool) {
	m.a.Log.Method("SetEnableAcceptance", nil)
	acceptor.Enable.Acceptance = v
}

func (m *MethodsEnableSetObj) Bookmarks(v bool) {
	m.a.Log.Method("SetEnableBookmarks", nil)
	acceptor.Enable.Bookmarks = v
}

func (m *MethodsEnableSetObj) NoPush(v bool) {
	m.a.Log.Method("SetEnableNoPush", nil)
	acceptor.Enable.NoPush = v
}

/**/

func (m *MethodsEnableGetObj) BarCode() bool {
	return m.a.Method.BarCode.Get.Enable()
}

func (m *MethodsEnableGetObj) BillType() []bool {
	return m.a.Method.Bill.Get.EnablesType()
}

func (m *MethodsEnableGetObj) BillValue() []bool {
	return m.a.Method.Bill.Get.EnablesValue()
}

func (m *MethodsEnableGetObj) Coupon() bool {
	return m.a.Method.Coupon.Get.Enable()
}

//

func (m *MethodsEnableSetObj) BarCode(v bool) {
	m.a.Method.BarCode.Set.Enable(v)
}

func (m *MethodsBillSetObj) BillType(v []bool) {
	m.a.Method.Bill.Set.EnablesType(v)
}

func (m *MethodsBillSetObj) BillValue(v []bool) {
	m.a.Method.Bill.Set.EnablesValue(v)
}

func (m *MethodCouponSetObj) Coupon(v bool) {
	m.a.Method.Coupon.Set.Enable(v)
}
