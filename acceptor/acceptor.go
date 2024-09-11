package acceptor

import (
	"github.com/hard-soft-ware/mpost/consts"
	"github.com/hard-soft-ware/mpost/enum"
)

////////////////////////////////////

var Connected bool
var AutoStack bool

var EscrowOrientation enum.OrientationType
var HighSecurity bool
var OrientationCtl enum.OrientationControlType
var OrientationCtlExt enum.OrientationControlType

var Version string
var BarCode string
var BootPN string

var InSoftResetOneSecondIgnore bool
var InSoftResetWaitForReply bool
var ExpandedNoteReporting bool
var IsQueryDeviceCapabilitiesSupported bool
var IsCheated bool
var IsPoweredUp bool
var IsInvalidCommand bool
var WasDocTypeSetOnEscrow bool
var WasDisconnected bool
var IsVeryFirstPoll bool

var StopWorkerThread bool
var StopOpenThread bool
var StopFlashDownloadThread bool
var SuppressStandardPoll bool
var WasStopped bool

var OpenThread bool
var FlashDownloadThread bool

////

func ConstructOmnibusCommand(payload []byte, controlCode consts.CmdType, data0Index int, billTypeEnables []bool) {
	payload[0] = controlCode.Byte()

	if Enable.Bookmarks && Enable.Acceptance && Device.State != enum.StateCalibrating {
		payload[0] |= 0x20
	}

	data0 := byte(0)

	if Enable.Acceptance && Device.State != enum.StateCalibrating {
		if ExpandedNoteReporting {
			data0 |= 0x7F
		} else {
			if len(billTypeEnables) == 0 {
				data0 |= 0x7F
			} else {
				for i, enable := range billTypeEnables {
					if enable {
						data0 |= 1 << i
					}
				}
			}
		}
	}

	data1 := byte(0) // Ignore bit 0 since we are not supporting special interrupt mode.

	if HighSecurity {
		data1 |= 0x02
	}

	switch OrientationCtl {
	case enum.OrientationControlTwoWay:
		data1 |= 0x04
	case enum.OrientationControlFourWay:
		data1 |= 0x0C
	}

	data1 |= 0x10 // Always enable escrow mode.

	data2 := byte(0)

	if Enable.NoPush {
		data2 |= 0x01
	}

	if Enable.BarCodes && Enable.Acceptance && Device.State != enum.StateCalibrating {
		data2 |= 0x02
	}

	switch Device.PowerUp {
	case enum.PowerUpB:
		data2 |= 0x04
	case enum.PowerUpC:
		data2 |= 0x0C
	}

	if ExpandedNoteReporting {
		data2 |= 0x10
	}

	if Enable.CouponExt && Cap.CouponExt {
		data2 |= 0x20
	}

	payload[data0Index] = data0
	payload[data0Index+1] = data1
	payload[data0Index+2] = data2
}
