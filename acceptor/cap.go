package acceptor

////////////////////////////////////

type CapStruct struct {
	ApplicationID      bool
	ApplicationPN      bool
	AssetNumber        bool
	Audit              bool
	BarCodes           bool
	BarCodesExt        bool
	BNFStatus          bool
	Bookmark           bool
	BootPN             bool
	Calibrate          bool
	CashBoxTotal       bool
	CouponExt          bool
	DevicePaused       bool
	DeviceSoftReset    bool
	DeviceType         bool
	DeviceResets       bool
	DeviceSerialNumber bool
	EscrowTimeout      bool
	FlashDownload      bool
	NoPush             bool
	OrientationExt     bool
	PupExt             bool
	TestDoc            bool
	VariantID          bool
	VariantPN          bool
}

var Cap CapStruct

////
