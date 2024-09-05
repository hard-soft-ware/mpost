package mpost

////////////////////////////////////

func (a *CAcceptor) raiseXX(e Event) {
	if handler, exists := a.eventHandlers[e]; exists && handler != nil {
		handler(a, 0)
	}
}

////

func (a *CAcceptor) RaiseConnectedEvent() {
	a.log.Debug().Str("Event", "ConnectedEvent").Msg("Raise")
	a.raiseXX(ConnectedEvent)
	a.shouldRaiseDisconnectedEvent = true
	a.shouldRaiseConnectedEvent = false
}

func (a *CAcceptor) RaiseDisconnectedEvent() {
	a.log.Debug().Str("Event", "DisconnectedEvent").Msg("Raise")
	a.raiseXX(DisconnectedEvent)
	a.shouldRaiseDisconnectedEvent = false
	a.shouldRaiseConnectedEvent = true
}

func (a *CAcceptor) RaiseDownloadRestartEvent() {
	a.log.Debug().Str("Event", "DownloadRestartEvent").Msg("Raise")
	a.raiseXX(DownloadRestartEvent)
	a.shouldRaiseDownloadRestartEvent = false
}

func (a *CAcceptor) RaiseCalibrateProgressEvent() {
	a.log.Debug().Str("Event", "CalibrateProgressEvent").Msg("Raise")
	a.raiseXX(CalibrateProgressEvent)
	a.shouldRaiseCalibrateProgressEvent = false
}
