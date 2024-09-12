package hook

import "github.com/hard-soft-ware/mpost/enum"

////////////////////////////////////

var eventHandlers = make(map[enum.EventType]func(int), enum.Event_End)

func Add(ev enum.EventType, h func(int)) {
	eventHandlers[ev] = h
}

func Clean() {
	eventHandlers = make(map[enum.EventType]func(int), enum.Event_End)

	Connected = false
	Escrow = false
	PUPEscrow = false
	Stacked = false
	Returned = false
	Rejected = false
	Cheated = false
	StackerFull = false
	CalibrateStart = false
	CalibrateProgress = false
	CalibrateFinish = false
	DownloadStart = false
	DownloadRestart = false
	DownloadProgress = false
	DownloadFinish = false
	PauseDetected = false
	PauseCleared = false
	StallDetected = false
	StallCleared = false
	JamDetected = false
	JamCleared = false
	PowerUp = false
	InvalidCommand = false
	CashBoxAttached = false
	CashBoxRemoved = false
	Disconnected = false
}
