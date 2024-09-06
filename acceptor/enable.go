package acceptor

////////////////////////////////////

type EnableStruct struct {
	Acceptance bool
	BarCodes   bool
	Bookmarks  bool
	CouponExt  bool
	NoPush     bool
}

var Enable EnableStruct
