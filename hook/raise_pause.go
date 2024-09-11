package hook

import "github.com/hard-soft-ware/mpost/enum"

////////////////////////////////////

type RaisePauseObj struct {
	r *RaiseObj
}

////////

func (r RaisePauseObj) Detected() {
	r.r.run(enum.EventPauseDetected, 0)
	PauseDetected = false
}

func (r RaisePauseObj) Cleared() {
	r.r.run(enum.EventPauseCleared, 0)
	PauseCleared = false
}
