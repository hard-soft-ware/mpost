package hook

import "github.com/hard-soft-ware/mpost/enum"

////////////////////////////////////

type RaiseCalibrateObj struct {
	r *RaiseObj
}

////////

func (r RaiseCalibrateObj) Progress() {
	r.r.run(enum.EventCalibrateProgress, 0)
	CalibrateProgress = false
}

func (r RaiseCalibrateObj) Start() {
	r.r.run(enum.EventCalibrateStart, 0)
	CalibrateStart = false
}

func (r RaiseCalibrateObj) Finish() {
	r.r.run(enum.EventCalibrateFinish, 0)
	CalibrateFinish = false
}
