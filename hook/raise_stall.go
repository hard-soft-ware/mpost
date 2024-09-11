package hook

import "github.com/hard-soft-ware/mpost/enum"

////////////////////////////////////

type RaiseStallObj struct {
	r *RaiseObj
}

////////

func (r RaiseStallObj) Detected() {
	r.r.run(enum.EventStallDetected, 0)
	StallDetected = false
}

func (r RaiseStallObj) Cleared() {
	r.r.run(enum.EventStallCleared, 0)
	StallCleared = false
}
