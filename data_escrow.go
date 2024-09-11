package mpost

import (
	"github.com/hard-soft-ware/mpost/acceptor"
	"github.com/hard-soft-ware/mpost/enum"
	"github.com/hard-soft-ware/mpost/hook"
)

////////////////////////////////////

func (dl *dataObj) EscrowReturn() {
	dl.escrowXX(0x40)
}

func (dl *dataObj) EscrowStack() {
	dl.escrowXX(0x20)
}

////

func (dl *dataObj) RaiseEvents() {
	if acceptor.IsPoweredUp && hook.PowerUp {
		hook.Raise.PowerUp()
	}

	if acceptor.IsVeryFirstPoll {
		acceptor.IsVeryFirstPoll = false
		return
	}

	switch acceptor.Device.State {
	case enum.StateEscrow:
		if acceptor.IsPoweredUp && hook.PUPEscrow {
			hook.Raise.PUPEscrow()
		} else if hook.Escrow {
			hook.Raise.Escrow()
		}
	case enum.StateStacked:
		if hook.Stacked {
			hook.Raise.Stacked()
		}
	case enum.StateReturned:
		if hook.Returned {
			hook.Raise.Returned()
		}
	case enum.StateRejected:
		if hook.Rejected {
			hook.Raise.Rejected()
		}
	case enum.StateStalled:
		if hook.StallDetected {
			hook.Raise.Stall.Detected()
		}
	}

	if acceptor.Device.State != enum.StateStalled && hook.StallCleared {
		hook.Raise.Stall.Cleared()
	}

	if acceptor.Cash.BoxFull && hook.StackerFull {
		hook.Raise.StackerFull()
	}

	if acceptor.IsCheated && hook.Cheated {
		hook.Raise.Cheated()
	}

	if acceptor.Cash.BoxAttached && hook.CashBoxAttached {
		hook.Raise.CashBox.Attached()
	}

	if !acceptor.Cash.BoxAttached && hook.CashBoxRemoved {
		hook.Raise.CashBox.Removed()
	}

	if acceptor.Device.Paused && hook.PauseDetected {
		hook.Raise.Pause.Detected()
	}

	if !acceptor.Device.Paused && hook.PauseCleared {
		hook.Raise.Pause.Cleared()
	}

	if acceptor.Device.Jammed && hook.JamDetected {
		hook.Raise.Jam.Detected()
	}

	if !acceptor.Device.Jammed && hook.JamCleared {
		hook.Raise.Jam.Cleared()
	}

	if acceptor.IsInvalidCommand && hook.InvalidCommand {
		hook.Raise.InvalidCommand()
	}

	if hook.CalibrateFinish {
		hook.Raise.Calibrate.Finish()
	}
}
