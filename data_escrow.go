package mpost

import (
	"github.com/hard-soft-ware/mpost/acceptor"
	"github.com/hard-soft-ware/mpost/bill"
	"github.com/hard-soft-ware/mpost/consts"
	"github.com/hard-soft-ware/mpost/enum"
)

////////////////////////////////////

func (dl *CDataLinkLayer) escrowXX(b byte) {
	if !acceptor.Connected {
		dl.log.Msg("serial not connected")
		return
	}

	payload := make([]byte, 4)
	acceptor.ConstructOmnibusCommand(payload, consts.CmdOmnibus, 1, bill.TypeEnables)

	payload[2] |= b

	dl.Acceptor.messageQueue <- NewCMessage(payload, false)
}

func (dl *CDataLinkLayer) EscrowReturn() {
	dl.escrowXX(0x40)
}

func (dl *CDataLinkLayer) EscrowStack() {
	dl.escrowXX(0x20)
}

////

func (dl *CDataLinkLayer) RaiseEvents() {
	if acceptor.IsPoweredUp && acceptor.ShouldRaise.PowerUpEvent {
		dl.log.Msg("Power Up Event Raised")
		acceptor.ShouldRaise.PowerUpEvent = false
	}

	if acceptor.IsVeryFirstPoll {
		acceptor.IsVeryFirstPoll = false
		return
	}

	switch acceptor.Device.State {
	case enum.StateEscrow:
		if acceptor.IsPoweredUp && acceptor.ShouldRaise.PUPEscrowEvent {
			dl.log.Msg("PUP Escrow Event Raised")
			acceptor.ShouldRaise.PUPEscrowEvent = false
		} else if acceptor.ShouldRaise.EscrowEvent {
			dl.log.Msg("Escrow Event Raised")
			acceptor.ShouldRaise.EscrowEvent = false
		}
	case enum.StateStacked:
		if acceptor.ShouldRaise.StackedEvent {
			dl.log.Msg("Stacked Event Raised")
			acceptor.ShouldRaise.StackedEvent = false
		}
	case enum.StateReturned:
		if acceptor.ShouldRaise.ReturnedEvent {
			dl.log.Msg("Returned Event Raised")
			acceptor.ShouldRaise.ReturnedEvent = false
		}
	case enum.StateRejected:
		if acceptor.ShouldRaise.RejectedEvent {
			dl.log.Msg("Rejected Event Raised")
			acceptor.ShouldRaise.RejectedEvent = false
		}
	case enum.StateStalled:
		if acceptor.ShouldRaise.StallDetectedEvent {
			dl.log.Msg("Stall Detected Event Raised")
			acceptor.ShouldRaise.StallDetectedEvent = false
		}
	}

	if acceptor.Device.State != enum.StateStalled && acceptor.ShouldRaise.StallClearedEvent {
		dl.log.Msg("Stall Cleared Event Raised")
		acceptor.ShouldRaise.StallClearedEvent = false
	}

	if acceptor.Cash.BoxFull && acceptor.ShouldRaise.StackerFullEvent {
		dl.log.Msg("Stacker Full Event Raised")
		acceptor.ShouldRaise.StackerFullEvent = false
	}

	if acceptor.IsCheated && acceptor.ShouldRaise.CheatedEvent {
		dl.log.Msg("Cheated Event Raised")
		acceptor.ShouldRaise.CheatedEvent = false
	}

	if acceptor.Cash.BoxAttached && acceptor.ShouldRaise.CashBoxAttachedEvent {
		dl.log.Msg("Cash Box Attached Event Raised")
		acceptor.ShouldRaise.CashBoxAttachedEvent = false
		acceptor.ShouldRaise.CashBoxRemovedEvent = true
	}

	if !acceptor.Cash.BoxAttached && acceptor.ShouldRaise.CashBoxRemovedEvent {
		dl.log.Msg("Cash Box Removed Event Raised")
		acceptor.ShouldRaise.CashBoxRemovedEvent = false
		acceptor.ShouldRaise.CashBoxAttachedEvent = true
	}

	if acceptor.Device.Paused && acceptor.ShouldRaise.PauseDetectedEvent {
		dl.log.Msg("Pause Detected Event Raised")
		acceptor.ShouldRaise.PauseDetectedEvent = false
	}

	if !acceptor.Device.Paused && acceptor.ShouldRaise.PauseClearedEvent {
		dl.log.Msg("Pause Cleared Event Raised")
		acceptor.ShouldRaise.PauseClearedEvent = false
	}

	if acceptor.IsDeviceJammed && acceptor.ShouldRaise.JamDetectedEvent {
		dl.log.Msg("Jam Detected Event Raised")
		acceptor.ShouldRaise.JamDetectedEvent = false
	}

	if !acceptor.IsDeviceJammed && acceptor.ShouldRaise.JamClearedEvent {
		dl.log.Msg("Jam Cleared Event Raised")
		acceptor.ShouldRaise.JamClearedEvent = false
	}

	if acceptor.IsInvalidCommand && acceptor.ShouldRaise.InvalidCommandEvent {
		dl.log.Msg("Invalid Command Event Raised")
		acceptor.ShouldRaise.InvalidCommandEvent = false
	}

	if acceptor.ShouldRaise.CalibrateFinishEvent {
		dl.log.Msg("Calibrate Finish Event Raised")
		acceptor.ShouldRaise.CalibrateFinishEvent = false
	}
}
