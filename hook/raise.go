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
	r.run(enum.EventInvalidCommand, 0)
	InvalidCommand = false
}

//

func (r RaiseObj) Connected() {
	if Connected {
		return
	}

	r.run(enum.EventConnected, 0)
	Disconnected = true
	Connected = false
}

func (r RaiseObj) Disconnected() {
	if !Disconnected {
		return
	}

	r.run(enum.EventDisconnected, 0)
	Disconnected = false
	Connected = true
}

func (r RaiseObj) PowerUp() {
	r.run(enum.EventPowerUp, 0)
	PowerUp = false
}

func (r RaiseObj) Returned() {
	r.run(enum.EventReturned, 0)
	Returned = false
}

func (r RaiseObj) Rejected() {
	r.run(enum.EventRejected, 0)
	Rejected = false
}

func (r RaiseObj) Cheated() {
	r.run(enum.EventCheated, 0)
	Cheated = false
}

//

func (r RaiseObj) Stacked() {
	r.run(enum.EventStacked, 0)
	Stacked = false
}

func (r RaiseObj) StackerFull() {
	r.run(enum.EventStackerFull, 0)
	StackerFull = false
}

//

func (r RaiseObj) PUPEscrow() {
	r.run(enum.EventPUPEscrow, 0)
	PUPEscrow = false
}

func (r RaiseObj) Escrow() {
	r.run(enum.EventEscrow, 0)
	Escrow = false
}
