package mpost

////////////////////////////////////

type CouponObj struct {
	OwnerID int
	Value   float64
}

func newCoupon(ownerID int, value float64) *CouponObj {
	return &CouponObj{
		OwnerID: ownerID,
		Value:   value,
	}
}
