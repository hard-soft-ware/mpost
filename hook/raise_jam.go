package hook

import "github.com/hard-soft-ware/mpost/enum"

////////////////////////////////////

type RaiseJamObj struct {
	r *RaiseObj
}

////////

func (r RaiseJamObj) Detected() {
	r.r.run(enum.EventJamDetected, 0)
	JamDetected = false
}

func (r RaiseJamObj) Cleared() {
	r.r.run(enum.EventJamCleared, 0)
	JamCleared = false
}
