package mpost

////////////////////////////////////

type MethodsObj struct {
	a *MpostObj

	Enable      *MethodsEnableObj
	Application *MethodsApplicationObj
	Audit       *MethodsAuditObj
	BarCode     *MethodsBarCodeObj
	Bill        *MethodsBillObj
	CashBox     *MethodsCashBoxObj
	Coupon      *MethodsCouponObj
	Device      *MethodsDeviceObj
	Orientation *MethodsOrientationObj
	Variant     *MethodsVariantObj
	Timeout     *MethodsTimeoutObj
	BNF         *MethodsBNFObj

	Other *MethodsOtherObj
}

////

func (a *MpostObj) newMethods() *MethodsObj {
	obj := MethodsObj{a: a}

	obj.Enable = obj.newEnable()
	obj.Application = obj.newApplication()
	obj.Audit = obj.newAudit()
	obj.BarCode = obj.newBarCode()
	obj.Bill = obj.newBill()
	obj.CashBox = obj.newCashBox()
	obj.Coupon = obj.newCoupon()
	obj.Device = obj.newDevice()
	obj.Orientation = obj.newOrientation()
	obj.Variant = obj.newVariant()
	obj.Timeout = obj.newTimeout()
	obj.BNF = obj.newBNF()

	obj.Other = obj.newOther()

	return &obj
}

////////////////
