package mpost

import "github.com/hard-soft-ware/mpost/acceptor"

////////////////////////////////////

func (m *MethodsObj) GetCoupon() *CouponObj {
	m.a.Log.Method("GetCoupon", nil)
	return m.a.coupon
}

func (m *MethodsObj) GetEnableCouponExt() bool {
	m.a.Log.Method("GetEnableCouponExt", nil)
	return acceptor.Enable.CouponExt
}

//

func (m *MethodsObj) GetCapCouponExt() bool {
	m.a.Log.Method("GetCapCouponExt", nil)
	return acceptor.Cap.CouponExt
}

////

func (m *MethodsObj) SetEnableCouponExt(v bool) {
	m.a.Log.Method("SetEnableCouponExt", v)
	acceptor.Enable.CouponExt = v
}
