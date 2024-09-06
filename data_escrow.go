package mpost

import (
	"github.com/hard-soft-ware/mpost/acceptor"
	"github.com/hard-soft-ware/mpost/consts"
	"github.com/hard-soft-ware/mpost/enum"
)

////////////////////////////////////

func (dl *CDataLinkLayer) escrowXX(b byte) {
	if !acceptor.Connected {
		dl.log.Debug().Msg("serial not connected")
		return
	}

	payload := make([]byte, 4)
	acceptor.ConstructOmnibusCommand(payload, consts.CmdOmnibus, 1, dl.Acceptor.billTypeEnables)

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
	if dl.Acceptor.isPoweredUp && acceptor.ShouldRaise.PowerUpEvent {
		dl.log.Debug().Msg("Power Up Event Raised")
		acceptor.ShouldRaise.PowerUpEvent = false
	}

	if dl.Acceptor.isVeryFirstPoll {
		dl.Acceptor.isVeryFirstPoll = false
		return
	}

	switch acceptor.Device.State {
	case enum.StateEscrow:
		if dl.Acceptor.isPoweredUp && acceptor.ShouldRaise.PUPEscrowEvent {
			dl.log.Debug().Msg("PUP Escrow Event Raised")
			acceptor.ShouldRaise.PUPEscrowEvent = false
		} else if acceptor.ShouldRaise.EscrowEvent {
			dl.log.Debug().Msg("Escrow Event Raised")
			acceptor.ShouldRaise.EscrowEvent = false
		}
	case enum.StateStacked:
		if acceptor.ShouldRaise.StackedEvent {
			dl.log.Debug().Msg("Stacked Event Raised")
			acceptor.ShouldRaise.StackedEvent = false
		}
	case enum.StateReturned:
		if acceptor.ShouldRaise.ReturnedEvent {
			dl.log.Debug().Msg("Returned Event Raised")
			acceptor.ShouldRaise.ReturnedEvent = false
		}
	case enum.StateRejected:
		if acceptor.ShouldRaise.RejectedEvent {
			dl.log.Debug().Msg("Rejected Event Raised")
			acceptor.ShouldRaise.RejectedEvent = false
		}
	case enum.StateStalled:
		if acceptor.ShouldRaise.StallDetectedEvent {
			dl.log.Debug().Msg("Stall Detected Event Raised")
			acceptor.ShouldRaise.StallDetectedEvent = false
		}
	}

	if acceptor.Device.State != enum.StateStalled && acceptor.ShouldRaise.StallClearedEvent {
		dl.log.Debug().Msg("Stall Cleared Event Raised")
		acceptor.ShouldRaise.StallClearedEvent = false
	}

	if acceptor.Cash.BoxFull && acceptor.ShouldRaise.StackerFullEvent {
		dl.log.Debug().Msg("Stacker Full Event Raised")
		acceptor.ShouldRaise.StackerFullEvent = false
	}

	if dl.Acceptor.isCheated && acceptor.ShouldRaise.CheatedEvent {
		dl.log.Debug().Msg("Cheated Event Raised")
		acceptor.ShouldRaise.CheatedEvent = false
	}

	if acceptor.Cash.BoxAttached && acceptor.ShouldRaise.CashBoxAttachedEvent {
		dl.log.Debug().Msg("Cash Box Attached Event Raised")
		acceptor.ShouldRaise.CashBoxAttachedEvent = false
		acceptor.ShouldRaise.CashBoxRemovedEvent = true
	}

	if !acceptor.Cash.BoxAttached && acceptor.ShouldRaise.CashBoxRemovedEvent {
		dl.log.Debug().Msg("Cash Box Removed Event Raised")
		acceptor.ShouldRaise.CashBoxRemovedEvent = false
		acceptor.ShouldRaise.CashBoxAttachedEvent = true
	}

	if acceptor.Device.Paused && acceptor.ShouldRaise.PauseDetectedEvent {
		dl.log.Debug().Msg("Pause Detected Event Raised")
		acceptor.ShouldRaise.PauseDetectedEvent = false
	}

	if !acceptor.Device.Paused && acceptor.ShouldRaise.PauseClearedEvent {
		dl.log.Debug().Msg("Pause Cleared Event Raised")
		acceptor.ShouldRaise.PauseClearedEvent = false
	}

	if dl.Acceptor.isDeviceJammed && acceptor.ShouldRaise.JamDetectedEvent {
		dl.log.Debug().Msg("Jam Detected Event Raised")
		acceptor.ShouldRaise.JamDetectedEvent = false
	}

	if !dl.Acceptor.isDeviceJammed && acceptor.ShouldRaise.JamClearedEvent {
		dl.log.Debug().Msg("Jam Cleared Event Raised")
		acceptor.ShouldRaise.JamClearedEvent = false
	}

	if dl.Acceptor.isInvalidCommand && acceptor.ShouldRaise.InvalidCommandEvent {
		dl.log.Debug().Msg("Invalid Command Event Raised")
		acceptor.ShouldRaise.InvalidCommandEvent = false
	}

	if acceptor.ShouldRaise.CalibrateFinishEvent {
		dl.log.Debug().Msg("Calibrate Finish Event Raised")
		acceptor.ShouldRaise.CalibrateFinishEvent = false
	}
}
