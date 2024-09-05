package mpost

////////////////////////////////////

type BNFStatus byte

const (
	Unknown BNFStatus = iota
	OK
	NotAttached
	Error
)

////

type DocumentType byte

const (
	None DocumentType = iota
	NoValue
	Bill
	Barcode
	Coupon
)

////

type Orientation byte

const (
	RightUp Orientation = iota
	RightDown
	LeftUp
	LeftDown
	UnknownOrientation
)

////

type OrientationControl byte

const (
	FourWay OrientationControl = iota
	TwoWay
	OneWay
)

////

type PowerUp byte

const (
	A PowerUp = iota
	B
	C
	E
)

////

type PupExt byte

const (
	Return PupExt = iota
	OutOfService
	StackNoCredit
	Stack
	WaitNoCredit
	Wait
)

////

type State byte

const (
	Escrow State = iota
	Stacked
	Returned
	Rejected
	Stalled
	Accepting
	CalibrateStart
	Calibrating
	Connecting
	Disconnected
	Downloading
	DownloadRestart
	DownloadStart
	Failed
	Idling
	PupEscrow
	Returning
	Stacking
)

////

type Bezel byte

const (
	Standard Bezel = iota
	Platform
	Diagnostic
)

////

type Event byte

type EventHandler func(*CAcceptor, int)

const (
	Events_Begin Event = iota
	ConnectedEvent
	EscrowEvent
	PUPEscrowEvent
	StackedEvent
	ReturnedEvent
	RejectedEvent
	CheatedEvent
	StackerFullEvent
	CalibrateStartEvent
	CalibrateProgressEvent
	CalibrateFinishEvent
	DownloadStartEvent
	DownloadRestartEvent
	DownloadProgressEvent
	DownloadFinishEvent
	PauseDetectedEvent
	PauseClearedEvent
	StallDetectedEvent
	StallClearedEvent
	JamDetectedEvent
	JamClearedEvent
	PowerUpEvent
	InvalidCommandEvent
	CashBoxAttachedEvent
	CashBoxRemovedEvent
	DisconnectedEvent
	Events_End
)

////
