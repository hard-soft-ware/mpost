package mpost

import "github.com/hard-soft-ware/mpost/acceptor"

////////////////////////////////////

func (a *MpostObj) GetCapCouponExt() bool {
	a.Log.Method("GetCapCouponExt", nil)
	return acceptor.Cap.CouponExt
}

func (a *MpostObj) GetCoupon() *CouponObj {
	a.Log.Method("GetCoupon", nil)
	return a.coupon
}

func (a *MpostObj) GetEnableCouponExt() bool {
	a.Log.Method("GetEnableCouponExt", nil)
	return acceptor.Enable.CouponExt
}

func (a *MpostObj) SetEnableCouponExt(v bool) {
	a.Log.Method("SetEnableCouponExt", nil)
	acceptor.Enable.CouponExt = v
}
