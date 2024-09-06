package acceptor

////////////////////////////////////

type ShouldRaiseStruct struct {
	ConnectedEvent         bool
	EscrowEvent            bool
	PUPEscrowEvent         bool
	StackedEvent           bool
	ReturnedEvent          bool
	RejectedEvent          bool
	CheatedEvent           bool
	StackerFullEvent       bool
	CalibrateStartEvent    bool
	CalibrateProgressEvent bool
	CalibrateFinishEvent   bool
	DownloadStartEvent     bool
	DownloadRestartEvent   bool
	DownloadProgressEvent  bool
	DownloadFinishEvent    bool
	PauseDetectedEvent     bool
	PauseClearedEvent      bool
	StallDetectedEvent     bool
	StallClearedEvent      bool
	JamDetectedEvent       bool
	JamClearedEvent        bool
	PowerUpEvent           bool
	InvalidCommandEvent    bool
	CashBoxAttachedEvent   bool
	CashBoxRemovedEvent    bool
	DisconnectedEvent      bool
}

var ShouldRaise ShouldRaiseStruct
