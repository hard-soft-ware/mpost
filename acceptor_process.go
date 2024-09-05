package mpost

////////////////////////////////////

func (a *CAcceptor) processData0(data0 byte) {
	if (data0 & 0x01) != 0 {
		if a.deviceState != Calibrating && a.deviceState != CalibrateStart {
			a.deviceState = Idling
		}
	}

	if (data0 & 0x02) != 0 {
		if a.deviceState != Calibrating && a.deviceState != CalibrateStart {
			a.deviceState = Accepting
		}
	}

	if (data0 & 0x04) != 0 {
		a.deviceState = Escrow
		if a.autoStack {
			a.shouldRaiseEscrowEvent = false
		}
	} else {
		a.shouldRaiseEscrowEvent = true
	}

	if (data0 & 0x08) != 0 {
		a.deviceState = Stacking
	}

	if (data0 & 0x10) != 0 {
		a.deviceState = Stacked
	} else {
		a.shouldRaiseStackedEvent = true
	}

	if (data0 & 0x20) != 0 {
		a.deviceState = Returning
	}

	if (data0 & 0x40) != 0 {
		a.deviceState = Returned
		a.bill = CBill{} // Resetting the bill
	} else {
		a.shouldRaiseReturnedEvent = true
	}
}

func (a *CAcceptor) processData1(data1 byte) {
	if (data1 & 0x01) != 0 {
		a.isCheated = true
	} else {
		a.isCheated = false
		a.shouldRaiseCheatedEvent = true
	}

	if (data1 & 0x02) != 0 {
		a.deviceState = Rejected
	} else {
		a.shouldRaiseRejectedEvent = true
	}

	if (data1 & 0x04) != 0 {
		a.isDeviceJammed = true
		a.shouldRaiseJamDetectedEvent = true
	} else {
		a.isDeviceJammed = false
		a.shouldRaiseJamClearedEvent = true
	}

	if (data1 & 0x08) != 0 {
		a.cashBoxFull = true
	} else {
		a.cashBoxFull = false
		a.shouldRaiseStackerFullEvent = true
	}

	a.cashBoxAttached = (data1 & 0x10) != 0

	if !a.cashBoxAttached {
		// Assume a DocumentType exists that handles this
		// _docType = NoValue
	}

	if (data1 & 0x20) != 0 {
		a.devicePaused = true
		a.shouldRaisePauseClearedEvent = true
	} else {
		a.devicePaused = false
		a.shouldRaisePauseDetectedEvent = true
	}

	if (data1 & 0x40) != 0 {
		a.deviceState = Calibrating
		if a.shouldRaiseCalibrateProgressEvent {
			a.RaiseCalibrateProgressEvent()
		}
	} else {
		if a.deviceState == Calibrating {
			a.shouldRaiseCalibrateFinishEvent = true
			a.deviceState = Idling
		}
	}
}

func (a *CAcceptor) processData2(data2 byte) {
	if !a.expandedNoteReporting {
		billTypeIndex := (data2 & 0x38) >> 3
		if billTypeIndex > 0 {
			if a.deviceState == Escrow || (a.deviceState == Stacked && !a.wasDocTypeSetOnEscrow) {
				a.bill = a.billTypes[billTypeIndex-1]
				a.docType = Bill
				a.wasDocTypeSetOnEscrow = a.deviceState == Escrow
			}
		} else {
			if a.deviceState == Stacked || a.deviceState == Escrow {
				a.bill = CBill{}
				a.docType = NoValue
				a.wasDocTypeSetOnEscrow = false
			}
		}
	} else {
		if a.deviceState == Stacked {
			if a.docType == Bill && a.bill.Value == 0.0 {
				a.docType = NoValue
			}
		} else if a.deviceState == Escrow {
			a.bill = CBill{}
			a.docType = NoValue
		}
	}

	if (data2 & 0x01) != 0 {
		a.isPoweredUp = true
		a.docType = NoValue
	} else {
		a.shouldRaisePowerUpEvent = true
		if !a.isVeryFirstPoll {
			a.isPoweredUp = false
		}
	}

	if (data2 & 0x02) != 0 {
		a.isInvalidCommand = true
	} else {
		a.isInvalidCommand = false
		a.shouldRaiseInvalidCommandEvent = true
	}

	if (data2 & 0x04) != 0 {
		a.deviceState = Failed
	}
}

func (a *CAcceptor) processData3(data3 byte) {
	if (data3 & 0x01) != 0 {
		a.deviceState = Stalled
		a.shouldRaiseStallClearedEvent = true
	} else {
		a.shouldRaiseStallDetectedEvent = true
	}

	if (data3 & 0x02) != 0 {
		a.deviceState = DownloadRestart
	}

	if (data3 & 0x08) != 0 {
		a.capBarCodesExt = true
	}

	if (data3 & 0x10) != 0 {
		a.isQueryDeviceCapabilitiesSupported = true
	}
}

func (a *CAcceptor) processData4(data4 byte) {
	a.deviceModel = int(data4 & 0x7F) //todo проверить валидность перевода
	m := a.deviceModel
	d := m

	a.capApplicationPN = m == 'T' || m == 'U'
	a.capAssetNumber = m == 'T' || m == 'U'
	a.capAudit = m == 'T' || m == 'U'
	a.capBarCodes = m == 'T' || m == 'U' || d == 15 || d == 23
	a.capBookmark = true
	a.capBootPN = m == 'T' || m == 'U'
	a.capCalibrate = true
	a.capCashBoxTotal = m == 'A' || m == 'B' || m == 'C' || m == 'D' || m == 'G' || m == 'M' || m == 'P' || m == 'W' || m == 'X'
	a.capCouponExt = m == 'P' || m == 'X'
	a.capDevicePaused = m == 'P' || m == 'X' || d == 31
	a.capDeviceSoftReset = m == 'A' || m == 'B' || m == 'C' || m == 'D' || m == 'G' || m == 'M' || m == 'P' || m == 'T' || m == 'U' || m == 'W' || m == 'X' || d == 31
	a.capDeviceType = m == 'T' || m == 'U'
	a.capDeviceResets = m == 'A' || m == 'B' || m == 'C' || m == 'D' || m == 'G' || m == 'M' || m == 'P' || m == 'T' || m == 'U' || m == 'W' || m == 'X'
	a.capDeviceSerialNumber = m == 'T' || m == 'U'
	a.capFlashDownload = true
	a.capEscrowTimeout = m == 'T' || m == 'U'
	a.capNoPush = m == 'P' || m == 'X' || d == 31 || d == 23
	a.capVariantPN = m == 'T' || m == 'U'
	a.expandedNoteReporting = m == 'T' || m == 'U' // This setting might be toggled in debug or production builds
}

func (a *CAcceptor) processData5(data5 byte) {
	switch {
	case a.deviceModel < 23, // S1K
		a.deviceModel == 30 || a.deviceModel == 31, // S3K
		a.deviceModel == 74,                        // CFMC
		a.deviceModel == 84 || a.deviceModel == 85: // CFSC
		a.deviceRevision = int(data5 & 0x7F)

	default: // S2K
		a.deviceRevision = int(data5&0x0F) + (int(data5&0x70)>>4)*10
	}
}
