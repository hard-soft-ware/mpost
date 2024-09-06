package mpost

import (
	"github.com/hard-soft-ware/mpost/acceptor"
	"github.com/hard-soft-ware/mpost/enum"
)

////////////////////////////////////

func (a *CAcceptor) raiseXX(e enum.EventType) {
	a.log.Event(e)
	if handler, exists := a.eventHandlers[e]; exists && handler != nil {
		handler(a, 0)
	}
}

////

func (a *CAcceptor) RaiseConnectedEvent() {
	a.raiseXX(enum.EventConnected)
	acceptor.ShouldRaise.DisconnectedEvent = true
	acceptor.ShouldRaise.ConnectedEvent = false
}

func (a *CAcceptor) RaiseDisconnectedEvent() {
	a.raiseXX(enum.EventDisconnected)
	acceptor.ShouldRaise.DisconnectedEvent = false
	acceptor.ShouldRaise.ConnectedEvent = true
}

func (a *CAcceptor) RaiseDownloadRestartEvent() {
	a.raiseXX(enum.EventDownloadRestart)
	acceptor.ShouldRaise.DownloadRestartEvent = false
}

func (a *CAcceptor) RaiseCalibrateProgressEvent() {
	a.raiseXX(enum.EventCalibrateProgress)
	acceptor.ShouldRaise.CalibrateProgressEvent = false
}
