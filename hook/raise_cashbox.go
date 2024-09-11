package hook

import "github.com/hard-soft-ware/mpost/enum"

////////////////////////////////////

type RaiseCashBoxObj struct {
	r *RaiseObj
}

////////

func (r RaiseCashBoxObj) Attached() {
	r.r.run(enum.EventCashBoxAttached, 0)
	CashBoxAttached = false
	CashBoxRemoved = true
}

func (r RaiseCashBoxObj) Removed() {
	r.r.run(enum.EventCashBoxRemoved, 0)
	CashBoxRemoved = false
	CashBoxAttached = true
}
