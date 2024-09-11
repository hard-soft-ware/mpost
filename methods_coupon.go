package mpost

import "github.com/hard-soft-ware/mpost/acceptor"

////////////////////////////////////

type MethodsCouponObj struct {
	a   *MpostObj
	Get MethodCouponGetObj
	Set MethodCouponSetObj
}

type MethodCouponGetObj struct{ a *MpostObj }
type MethodCouponSetObj struct{ a *MpostObj }

func (m *MethodsObj) newCoupon() *MethodsCouponObj {
	obj := MethodsCouponObj{}

	obj.a = m.a
	obj.Get.a = m.a
	obj.Set.a = m.a

	return &obj
}

////////////////

func (m *MethodCouponGetObj) Self() *CouponObj {
	m.a.Log.Method("GetCoupon", nil)
	return m.a.coupon
}

func (m *MethodCouponGetObj) Enable() bool {
	m.a.Log.Method("GetEnableCouponExt", nil)
	return acceptor.Enable.CouponExt
}

//

func (m *MethodCouponGetObj) Cap() bool {
	m.a.Log.Method("GetCapCouponExt", nil)
	return acceptor.Cap.CouponExt
}

////

func (m *MethodCouponSetObj) Enable(v bool) {
	m.a.Log.Method("SetEnableCouponExt", nil)
	acceptor.Enable.CouponExt = v
}
