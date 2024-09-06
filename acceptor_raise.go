package mpost

import (
	"github.com/hard-soft-ware/mpost/acceptor"
	"github.com/hard-soft-ware/mpost/enum"
)

////////////////////////////////////

func (a *CAcceptor) raiseXX(e enum.EventType) {
	if handler, exists := a.eventHandlers[e]; exists && handler != nil {
		handler(a, 0)
	}
}

////

func (a *CAcceptor) RaiseConnectedEvent() {
	a.log.Debug().Str("Event", "ConnectedEvent").Msg("Raise")
	a.raiseXX(enum.EventConnected)
	acceptor.ShouldRaise.DisconnectedEvent = true
	acceptor.ShouldRaise.ConnectedEvent = false
}

func (a *CAcceptor) RaiseDisconnectedEvent() {
	a.log.Debug().Str("Event", "DisconnectedEvent").Msg("Raise")
	a.raiseXX(enum.EventDisconnected)
	acceptor.ShouldRaise.DisconnectedEvent = false
	acceptor.ShouldRaise.ConnectedEvent = true
}

func (a *CAcceptor) RaiseDownloadRestartEvent() {
	a.log.Debug().Str("Event", "DownloadRestartEvent").Msg("Raise")
	a.raiseXX(enum.EventDownloadRestart)
	acceptor.ShouldRaise.DownloadRestartEvent = false
}

func (a *CAcceptor) RaiseCalibrateProgressEvent() {
	a.log.Debug().Str("Event", "CalibrateProgressEvent").Msg("Raise")
	a.raiseXX(enum.EventCalibrateProgress)
	acceptor.ShouldRaise.CalibrateProgressEvent = false
}
