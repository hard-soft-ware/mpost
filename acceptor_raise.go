package mpost

import (
	"github.com/hard-soft-ware/mpost/acceptor"
	"github.com/hard-soft-ware/mpost/enum"
)

////////////////////////////////////

func (a *CAcceptor) raiseXX(e enum.EventType, b int) {
	a.log.Event(e)
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
	if acceptor.ShouldRaise.DisconnectedEvent {
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
}

func (a *CAcceptor) RaiseCalibrateFinishEvent() {
	a.raiseXX(enum.EventCalibrateFinish, 0)
}
