package mpost

import (
	"github.com/hard-soft-ware/mpost/acceptor"
	"github.com/hard-soft-ware/mpost/bill"
	"github.com/hard-soft-ware/mpost/enum"
)

////////////////////////////////////

func (a *CAcceptor) processData0(data0 byte) {
	if (data0 & 0x01) != 0 {
		if acceptor.Device.State != enum.StateCalibrating && acceptor.Device.State != enum.StateCalibrateStart {
			acceptor.Device.State = enum.StateIdling
		}
	}

	if (data0 & 0x02) != 0 {
		if acceptor.Device.State != enum.StateCalibrating && acceptor.Device.State != enum.StateCalibrateStart {
			acceptor.Device.State = enum.StateAccepting
		}
	}

	if (data0 & 0x04) != 0 {
		acceptor.Device.State = enum.StateEscrow
		if acceptor.AutoStack {
			acceptor.ShouldRaise.EscrowEvent = false
		}
	} else {
		acceptor.ShouldRaise.EscrowEvent = true
	}

	if (data0 & 0x08) != 0 {
		acceptor.Device.State = enum.StateStacking
	}

	if (data0 & 0x10) != 0 {
		acceptor.Device.State = enum.StateStacked
	} else {
		acceptor.ShouldRaise.StackedEvent = true
	}

	if (data0 & 0x20) != 0 {
		acceptor.Device.State = enum.StateReturning
	}

	if (data0 & 0x40) != 0 {
		acceptor.Device.State = enum.StateReturned
		bill.Reset() // Resetting the bill
	} else {
		acceptor.ShouldRaise.ReturnedEvent = true
	}
}

func (a *CAcceptor) processData1(data1 byte) {
	if (data1 & 0x01) != 0 {
		acceptor.IsCheated = true
	} else {
		acceptor.IsCheated = false
		acceptor.ShouldRaise.CheatedEvent = true
	}

	if (data1 & 0x02) != 0 {
		acceptor.Device.State = enum.StateRejected
	} else {
		acceptor.ShouldRaise.RejectedEvent = true
	}

	if (data1 & 0x04) != 0 {
		acceptor.IsDeviceJammed = true
		acceptor.ShouldRaise.JamDetectedEvent = true
	} else {
		acceptor.IsDeviceJammed = false
		acceptor.ShouldRaise.JamClearedEvent = true
	}

	if (data1 & 0x08) != 0 {
		acceptor.Cash.BoxFull = true
	} else {
		acceptor.Cash.BoxFull = false
		acceptor.ShouldRaise.StackerFullEvent = true
	}

	acceptor.Cash.BoxAttached = (data1 & 0x10) != 0

	if !acceptor.Cash.BoxAttached {
		// Assume a DocumentType exists that handles this
		// _docType = NoValue
	}

	if (data1 & 0x20) != 0 {
		acceptor.Device.Paused = true
		acceptor.ShouldRaise.PauseClearedEvent = true
	} else {
		acceptor.Device.Paused = false
		acceptor.ShouldRaise.PauseDetectedEvent = true
	}

	if (data1 & 0x40) != 0 {
		acceptor.Device.State = enum.StateCalibrating
		if acceptor.ShouldRaise.CalibrateProgressEvent {
			a.RaiseCalibrateProgressEvent()
		}
	} else {
		if acceptor.Device.State == enum.StateCalibrating {
			acceptor.ShouldRaise.CalibrateFinishEvent = true
			acceptor.Device.State = enum.StateIdling
		}
	}
}

func (a *CAcceptor) processData2(data2 byte) {
	if !acceptor.ExpandedNoteReporting {
		billTypeIndex := (data2 & 0x38) >> 3
		if billTypeIndex > 0 {
			if acceptor.Device.State == enum.StateEscrow || (acceptor.Device.State == enum.StateStacked && !acceptor.WasDocTypeSetOnEscrow) {
				bill.Bill = bill.Types[billTypeIndex-1]
				a.docType = enum.DocumentBill
				acceptor.WasDocTypeSetOnEscrow = acceptor.Device.State == enum.StateEscrow
			}
		} else {
			if acceptor.Device.State == enum.StateStacked || acceptor.Device.State == enum.StateEscrow {
				bill.Reset()
				a.docType = enum.DocumentNoValue
				acceptor.WasDocTypeSetOnEscrow = false
			}
		}
	} else {
		if acceptor.Device.State == enum.StateStacked {
			if a.docType == enum.DocumentBill && bill.Bill.Value == 0.0 {
				a.docType = enum.DocumentNoValue
			}
		} else if acceptor.Device.State == enum.StateEscrow {
			bill.Reset()
			a.docType = enum.DocumentNoValue
		}
	}

	if (data2 & 0x01) != 0 {
		acceptor.IsPoweredUp = true
		a.docType = enum.DocumentNoValue
	} else {
		acceptor.ShouldRaise.PowerUpEvent = true
		if !acceptor.IsVeryFirstPoll {
			acceptor.IsPoweredUp = false
		}
	}

	if (data2 & 0x02) != 0 {
		acceptor.IsInvalidCommand = true
	} else {
		acceptor.IsInvalidCommand = false
		acceptor.ShouldRaise.InvalidCommandEvent = true
	}

	if (data2 & 0x04) != 0 {
		acceptor.Device.State = enum.StateFailed
	}
}

func (a *CAcceptor) processData3(data3 byte) {
	if (data3 & 0x01) != 0 {
		acceptor.Device.State = enum.StateStalled
		acceptor.ShouldRaise.StallClearedEvent = true
	} else {
		acceptor.ShouldRaise.StallDetectedEvent = true
	}

	if (data3 & 0x02) != 0 {
		acceptor.Device.State = enum.StateDownloadRestart
	}

	if (data3 & 0x08) != 0 {
		acceptor.Cap.BarCodesExt = true
	}

	if (data3 & 0x10) != 0 {
		acceptor.IsQueryDeviceCapabilitiesSupported = true
	}
}

func (a *CAcceptor) processData4(data4 byte) {
	acceptor.Device.Model = int(data4 & 0x7F) //todo проверить валидность перевода
	m := acceptor.Device.Model
	d := m

	acceptor.Cap.ApplicationPN = m == 'T' || m == 'U'
	acceptor.Cap.AssetNumber = m == 'T' || m == 'U'
	acceptor.Cap.Audit = m == 'T' || m == 'U'
	acceptor.Cap.BarCodes = m == 'T' || m == 'U' || d == 15 || d == 23
	acceptor.Cap.Bookmark = true
	acceptor.Cap.BootPN = m == 'T' || m == 'U'
	acceptor.Cap.Calibrate = true
	acceptor.Cap.CashBoxTotal = m == 'A' || m == 'B' || m == 'C' || m == 'D' || m == 'G' || m == 'M' || m == 'P' || m == 'W' || m == 'X'
	acceptor.Cap.CouponExt = m == 'P' || m == 'X'
	acceptor.Cap.DevicePaused = m == 'P' || m == 'X' || d == 31
	acceptor.Cap.DeviceSoftReset = m == 'A' || m == 'B' || m == 'C' || m == 'D' || m == 'G' || m == 'M' || m == 'P' || m == 'T' || m == 'U' || m == 'W' || m == 'X' || d == 31
	acceptor.Cap.DeviceType = m == 'T' || m == 'U'
	acceptor.Cap.DeviceResets = m == 'A' || m == 'B' || m == 'C' || m == 'D' || m == 'G' || m == 'M' || m == 'P' || m == 'T' || m == 'U' || m == 'W' || m == 'X'
	acceptor.Cap.DeviceSerialNumber = m == 'T' || m == 'U'
	acceptor.Cap.FlashDownload = true
	acceptor.Cap.EscrowTimeout = m == 'T' || m == 'U'
	acceptor.Cap.NoPush = m == 'P' || m == 'X' || d == 31 || d == 23
	acceptor.Cap.VariantPN = m == 'T' || m == 'U'
	acceptor.ExpandedNoteReporting = m == 'T' || m == 'U' // This setting might be toggled in debug or production builds
}

func (a *CAcceptor) processData5(data5 byte) {
	switch {
	case acceptor.Device.Model < 23, // S1K
		acceptor.Device.Model == 30 || acceptor.Device.Model == 31, // S3K
		acceptor.Device.Model == 74,                                // CFMC
		acceptor.Device.Model == 84 || acceptor.Device.Model == 85: // CFSC
		acceptor.Device.Revision = int(data5 & 0x7F)

	default: // S2K
		acceptor.Device.Revision = int(data5&0x0F) + (int(data5&0x70)>>4)*10
	}
}
