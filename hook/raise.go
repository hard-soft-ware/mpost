package hook

import "github.com/hard-soft-ware/mpost/enum"

////////////////////////////////////

type RaiseObj struct {
	Log func(enum.EventType, int)

	Download  RaiseDownloadObj
	Calibrate RaiseCalibrateObj
	Stall     RaiseStallObj
	CashBox   RaiseCashBoxObj
	Pause     RaisePauseObj
	Jam       RaiseJamObj
}

var Raise RaiseObj

func init() {
	Raise.Log = func(enum.EventType, int) {}

	Raise.Download.r = &Raise
	Raise.Calibrate.r = &Raise
	Raise.Stall.r = &Raise
	Raise.CashBox.r = &Raise
	Raise.Pause.r = &Raise
	Raise.Jam.r = &Raise
}

////

func (r RaiseObj) run(e enum.EventType, b int) {
	r.Log(e, b)

	if handler, exists := eventHandlers[e]; exists && handler != nil {
		handler(b)
	}
}

////////

func (r RaiseObj) InvalidCommand() {
	InvalidCommand = false
	r.run(enum.EventInvalidCommand, 0)
}

//

func (r RaiseObj) Connected() {
	if Connected {
		return
	}

	Disconnected = true
	Connected = false
	r.run(enum.EventConnected, 0)
}

func (r RaiseObj) Disconnected() {
	if !Disconnected {
		return
	}

	Disconnected = false
	Connected = true
	r.run(enum.EventDisconnected, 0)
}

func (r RaiseObj) PowerUp() {
	PowerUp = false
	r.run(enum.EventPowerUp, 0)
}

func (r RaiseObj) Returned() {
	Returned = false
	r.run(enum.EventReturned, 0)
}

func (r RaiseObj) Rejected() {
	Rejected = false
	r.run(enum.EventRejected, 0)
}

func (r RaiseObj) Cheated() {
	Cheated = false
	r.run(enum.EventCheated, 0)
}

//

func (r RaiseObj) Stacked() {
	Stacked = false
	r.run(enum.EventStacked, 0)
}

func (r RaiseObj) StackerFull() {
	StackerFull = false
	r.run(enum.EventStackerFull, 0)
}

//

func (r RaiseObj) PUPEscrow() {
	PUPEscrow = false
	r.run(enum.EventPUPEscrow, 0)
}

func (r RaiseObj) Escrow() {
	Escrow = false
	r.run(enum.EventEscrow, 0)
}
