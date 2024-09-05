package mpost

////////////////////////////////////

func (dl *CDataLinkLayer) escrowXX(b byte) {
	if !dl.Acceptor.connected {
		dl.log.Debug().Msg("serial not connected")
		return
	}

	payload := make([]byte, 4)
	dl.Acceptor.ConstructOmnibusCommand(payload, CmdOmnibus, 1)

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
	if dl.Acceptor.isPoweredUp && dl.Acceptor.shouldRaisePowerUpEvent {
		dl.log.Debug().Msg("Power Up Event Raised")
		dl.Acceptor.shouldRaisePowerUpEvent = false
	}

	if dl.Acceptor.isVeryFirstPoll {
		dl.Acceptor.isVeryFirstPoll = false
		return
	}

	switch dl.Acceptor.deviceState {
	case Escrow:
		if dl.Acceptor.isPoweredUp && dl.Acceptor.shouldRaisePUPEscrowEvent {
			dl.log.Debug().Msg("PUP Escrow Event Raised")
			dl.Acceptor.shouldRaisePUPEscrowEvent = false
		} else if dl.Acceptor.shouldRaiseEscrowEvent {
			dl.log.Debug().Msg("Escrow Event Raised")
			dl.Acceptor.shouldRaiseEscrowEvent = false
		}
	case Stacked:
		if dl.Acceptor.shouldRaiseStackedEvent {
			dl.log.Debug().Msg("Stacked Event Raised")
			dl.Acceptor.shouldRaiseStackedEvent = false
		}
	case Returned:
		if dl.Acceptor.shouldRaiseReturnedEvent {
			dl.log.Debug().Msg("Returned Event Raised")
			dl.Acceptor.shouldRaiseReturnedEvent = false
		}
	case Rejected:
		if dl.Acceptor.shouldRaiseRejectedEvent {
			dl.log.Debug().Msg("Rejected Event Raised")
			dl.Acceptor.shouldRaiseRejectedEvent = false
		}
	case Stalled:
		if dl.Acceptor.shouldRaiseStallDetectedEvent {
			dl.log.Debug().Msg("Stall Detected Event Raised")
			dl.Acceptor.shouldRaiseStallDetectedEvent = false
		}
	}

	if dl.Acceptor.deviceState != Stalled && dl.Acceptor.shouldRaiseStallClearedEvent {
		dl.log.Debug().Msg("Stall Cleared Event Raised")
		dl.Acceptor.shouldRaiseStallClearedEvent = false
	}

	if dl.Acceptor.cashBoxFull && dl.Acceptor.shouldRaiseStackerFullEvent {
		dl.log.Debug().Msg("Stacker Full Event Raised")
		dl.Acceptor.shouldRaiseStackerFullEvent = false
	}

	if dl.Acceptor.isCheated && dl.Acceptor.shouldRaiseCheatedEvent {
		dl.log.Debug().Msg("Cheated Event Raised")
		dl.Acceptor.shouldRaiseCheatedEvent = false
	}

	if dl.Acceptor.cashBoxAttached && dl.Acceptor.shouldRaiseCashBoxAttachedEvent {
		dl.log.Debug().Msg("Cash Box Attached Event Raised")
		dl.Acceptor.shouldRaiseCashBoxAttachedEvent = false
		dl.Acceptor.shouldRaiseCashBoxRemovedEvent = true
	}

	if !dl.Acceptor.cashBoxAttached && dl.Acceptor.shouldRaiseCashBoxRemovedEvent {
		dl.log.Debug().Msg("Cash Box Removed Event Raised")
		dl.Acceptor.shouldRaiseCashBoxRemovedEvent = false
		dl.Acceptor.shouldRaiseCashBoxAttachedEvent = true
	}

	if dl.Acceptor.devicePaused && dl.Acceptor.shouldRaisePauseDetectedEvent {
		dl.log.Debug().Msg("Pause Detected Event Raised")
		dl.Acceptor.shouldRaisePauseDetectedEvent = false
	}

	if !dl.Acceptor.devicePaused && dl.Acceptor.shouldRaisePauseClearedEvent {
		dl.log.Debug().Msg("Pause Cleared Event Raised")
		dl.Acceptor.shouldRaisePauseClearedEvent = false
	}

	if dl.Acceptor.isDeviceJammed && dl.Acceptor.shouldRaiseJamDetectedEvent {
		dl.log.Debug().Msg("Jam Detected Event Raised")
		dl.Acceptor.shouldRaiseJamDetectedEvent = false
	}

	if !dl.Acceptor.isDeviceJammed && dl.Acceptor.shouldRaiseJamClearedEvent {
		dl.log.Debug().Msg("Jam Cleared Event Raised")
		dl.Acceptor.shouldRaiseJamClearedEvent = false
	}

	if dl.Acceptor.isInvalidCommand && dl.Acceptor.shouldRaiseInvalidCommandEvent {
		dl.log.Debug().Msg("Invalid Command Event Raised")
		dl.Acceptor.shouldRaiseInvalidCommandEvent = false
	}

	if dl.Acceptor.shouldRaiseCalibrateFinishEvent {
		dl.log.Debug().Msg("Calibrate Finish Event Raised")
		dl.Acceptor.shouldRaiseCalibrateFinishEvent = false
	}
}
