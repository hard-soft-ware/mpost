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

	dl.Acceptor.SendAsynchronousCommand(payload)
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
		dl.Acceptor.RaisePowerUpEvent()
	}

	if acceptor.IsVeryFirstPoll {
		acceptor.IsVeryFirstPoll = false
		return
	}

	switch acceptor.Device.State {
	case enum.StateEscrow:
		if acceptor.IsPoweredUp && acceptor.ShouldRaise.PUPEscrowEvent {
			dl.Acceptor.RaisePUPEscrowEvent()
		} else if acceptor.ShouldRaise.EscrowEvent {
			dl.Acceptor.RaiseEscrowEvent()
		}
	case enum.StateStacked:
		if acceptor.ShouldRaise.StackedEvent {
			dl.Acceptor.RaiseStackedEvent()
		}
	case enum.StateReturned:
		if acceptor.ShouldRaise.ReturnedEvent {
			dl.Acceptor.RaiseReturnedEvent()
		}
	case enum.StateRejected:
		if acceptor.ShouldRaise.RejectedEvent {
			dl.Acceptor.RaiseRejectedEvent()
		}
	case enum.StateStalled:
		if acceptor.ShouldRaise.StallDetectedEvent {
			dl.Acceptor.RaiseStallDetectedEvent()
		}
	}

	if acceptor.Device.State != enum.StateStalled && acceptor.ShouldRaise.StallClearedEvent {
		dl.Acceptor.RaiseStallClearedEvent()
	}

	if acceptor.Cash.BoxFull && acceptor.ShouldRaise.StackerFullEvent {
		dl.Acceptor.RaiseStackerFullEvent()
	}

	if acceptor.IsCheated && acceptor.ShouldRaise.CheatedEvent {
		dl.Acceptor.RaiseCheatedEvent()
	}

	if acceptor.Cash.BoxAttached && acceptor.ShouldRaise.CashBoxAttachedEvent {
		dl.Acceptor.RaiseCashBoxAttachedEvent()
	}

	if !acceptor.Cash.BoxAttached && acceptor.ShouldRaise.CashBoxRemovedEvent {
		dl.Acceptor.RaiseCashBoxRemovedEvent()
	}

	if acceptor.Device.Paused && acceptor.ShouldRaise.PauseDetectedEvent {
		dl.Acceptor.RaisePauseDetectedEvent()
	}

	if !acceptor.Device.Paused && acceptor.ShouldRaise.PauseClearedEvent {
		dl.Acceptor.RaisePauseClearedEvent()
	}

	if acceptor.Device.Jammed && acceptor.ShouldRaise.JamDetectedEvent {
		dl.Acceptor.RaiseJamDetectedEvent()
	}

	if !acceptor.Device.Jammed && acceptor.ShouldRaise.JamClearedEvent {
		dl.Acceptor.RaiseJamClearedEvent()
	}

	if acceptor.IsInvalidCommand && acceptor.ShouldRaise.InvalidCommandEvent {
		dl.Acceptor.RaiseInvalidCommandEvent()
	}

	if acceptor.ShouldRaise.CalibrateFinishEvent {
		dl.Acceptor.RaiseCalibrateFinishEvent()
	}
}
