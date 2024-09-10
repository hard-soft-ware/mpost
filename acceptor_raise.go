package mpost

import (
	"github.com/hard-soft-ware/mpost/acceptor"
	"github.com/hard-soft-ware/mpost/enum"
)

////////////////////////////////////

func (a *CAcceptor) raiseXX(e enum.EventType, b int) {
	if e != enum.EventJamCleared {
		a.log.Event(e)
	}

	if handler, exists := a.eventHandlers[e]; exists && handler != nil {
		handler(a, b)
	}
}

////

func (a *CAcceptor) RaiseConnectedEvent() {
	if acceptor.ShouldRaise.ConnectedEvent {
		return
	}

	a.raiseXX(enum.EventConnected, 0)
	acceptor.ShouldRaise.DisconnectedEvent = true
	acceptor.ShouldRaise.ConnectedEvent = false
}

func (a *CAcceptor) RaiseDisconnectedEvent() {
	if !acceptor.ShouldRaise.DisconnectedEvent {
		return
	}

	a.raiseXX(enum.EventDisconnected, 0)
	acceptor.ShouldRaise.DisconnectedEvent = false
	acceptor.ShouldRaise.ConnectedEvent = true
}

func (a *CAcceptor) RaiseDownloadRestartEvent() {
	a.raiseXX(enum.EventDownloadRestart, 0)
	acceptor.ShouldRaise.DownloadRestartEvent = false
}

//

func (a *CAcceptor) RaiseCalibrateProgressEvent() {
	a.raiseXX(enum.EventCalibrateProgress, 0)
	acceptor.ShouldRaise.CalibrateProgressEvent = false
}

func (a *CAcceptor) RaiseCalibrateStartEvent() {
	a.raiseXX(enum.EventCalibrateStart, 0)
	acceptor.ShouldRaise.CalibrateStartEvent = false
}

func (a *CAcceptor) RaiseCalibrateFinishEvent() {
	a.raiseXX(enum.EventCalibrateFinish, 0)
	acceptor.ShouldRaise.CalibrateFinishEvent = false
}

//

func (a *CAcceptor) RaiseDownloadFinishEvent(st bool) {
	v := 0
	if st {
		v = 1
	}
	a.raiseXX(enum.EventDownloadFinish, v)
	acceptor.ShouldRaise.DownloadFinishEvent = false
}

func (a *CAcceptor) RaiseDownloadStartEvent(v int) {
	a.raiseXX(enum.EventDownloadStart, v)

	acceptor.ShouldRaise.DownloadStartEvent = false
	acceptor.ShouldRaise.DownloadProgressEvent = true
}

func (a *CAcceptor) RaiseDownloadProgressEvent(v int) {
	a.raiseXX(enum.EventDownloadProgress, v)
}

//

func (a *CAcceptor) RaisePowerUpEvent() {
	a.raiseXX(enum.EventPowerUp, 0)
	acceptor.ShouldRaise.PowerUpEvent = false
}

func (a *CAcceptor) RaisePUPEscrowEvent() {
	a.raiseXX(enum.EventPUPEscrow, 0)
	acceptor.ShouldRaise.PUPEscrowEvent = false
}

func (a *CAcceptor) RaiseEscrowEvent() {
	a.raiseXX(enum.EventEscrow, 0)
	acceptor.ShouldRaise.EscrowEvent = false
}

func (a *CAcceptor) RaiseStackedEvent() {
	a.raiseXX(enum.EventStacked, 0)
	acceptor.ShouldRaise.StackedEvent = false
}

func (a *CAcceptor) RaiseReturnedEvent() {
	a.raiseXX(enum.EventReturned, 0)
	acceptor.ShouldRaise.ReturnedEvent = false
}

func (a *CAcceptor) RaiseRejectedEvent() {
	a.raiseXX(enum.EventRejected, 0)
	acceptor.ShouldRaise.RejectedEvent = false
}

func (a *CAcceptor) RaiseStallDetectedEvent() {
	a.raiseXX(enum.EventStallDetected, 0)
	acceptor.ShouldRaise.StallDetectedEvent = false
}

func (a *CAcceptor) RaiseStallClearedEvent() {
	a.raiseXX(enum.EventStallCleared, 0)
	acceptor.ShouldRaise.StallClearedEvent = false
}

func (a *CAcceptor) RaiseStackerFullEvent() {
	a.raiseXX(enum.EventStackerFull, 0)
	acceptor.ShouldRaise.StackerFullEvent = false
}

func (a *CAcceptor) RaiseCheatedEvent() {
	a.raiseXX(enum.EventCheated, 0)
	acceptor.ShouldRaise.CheatedEvent = false
}

func (a *CAcceptor) RaiseCashBoxAttachedEvent() {
	a.raiseXX(enum.EventCashBoxAttached, 0)
	acceptor.ShouldRaise.CashBoxAttachedEvent = false
	acceptor.ShouldRaise.CashBoxRemovedEvent = true
}

func (a *CAcceptor) RaiseCashBoxRemovedEvent() {
	a.raiseXX(enum.EventCashBoxRemoved, 0)
	acceptor.ShouldRaise.CashBoxRemovedEvent = false
	acceptor.ShouldRaise.CashBoxAttachedEvent = true
}

func (a *CAcceptor) RaisePauseDetectedEvent() {
	a.raiseXX(enum.EventPauseDetected, 0)
	acceptor.ShouldRaise.PauseDetectedEvent = false
}

func (a *CAcceptor) RaisePauseClearedEvent() {
	a.raiseXX(enum.EventPauseCleared, 0)
	acceptor.ShouldRaise.PauseClearedEvent = false
}

func (a *CAcceptor) RaiseJamDetectedEvent() {
	a.raiseXX(enum.EventJamDetected, 0)
	acceptor.ShouldRaise.JamDetectedEvent = false
}

func (a *CAcceptor) RaiseJamClearedEvent() {
	a.raiseXX(enum.EventJamCleared, 0)
	acceptor.ShouldRaise.JamClearedEvent = false
}

func (a *CAcceptor) RaiseInvalidCommandEvent() {
	a.raiseXX(enum.EventInvalidCommand, 0)
	acceptor.ShouldRaise.InvalidCommandEvent = false
}
