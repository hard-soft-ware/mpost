package mpost

////////////////////////////////////

type CCoupon struct {
	OwnerID int
	Value   float64
}

func NewCCoupon(ownerID int, value float64) *CCoupon {
	return &CCoupon{
		OwnerID: ownerID,
		Value:   value,
	}
}

////

func (c *CCoupon) GetOwnerID() int {
	return c.OwnerID
}

func (c *CCoupon) GetValue() float64 {
	return c.Value
}

func (c *CCoupon) SetOwnerID(ownerID int) {
	c.OwnerID = ownerID
}

func (c *CCoupon) SetValue(value float64) {
	c.Value = value
}
