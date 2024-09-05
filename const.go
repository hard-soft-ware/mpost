package mpost

////////////////////////////////////

const (
	CmdOmnibus       = 0x10
	CmdCalibrate     = 0x40
	CmdFlashDownload = 0x50
	CmdAuxiliary     = 0x60
	CmdExpanded      = 0x70

	CmdAuxQuerySoftwareCRC                      = 0x00
	CmdAuxQueryCashBoxTotal                     = 0x01
	CmdAuxQueryDeviceResets                     = 0x02
	CmdAuxClearCashBoxTotal                     = 0x03
	CmdAuxQueryAcceptorType                     = 0x04
	CmdAuxQueryAcceptorSerialNumber             = 0x05
	CmdAuxQueryAcceptorBootPartNumber           = 0x06
	CmdAuxQueryAcceptorApplicationPartNumber    = 0x07
	CmdAuxQueryAcceptorVariantName              = 0x08
	CmdAuxQueryAcceptorVariantPartNumber        = 0x09
	CmdAuxQueryAcceptorAuditLifeTimeTotals      = 0x0A
	CmdAuxQueryAcceptorAuditQPMeasures          = 0x0B
	CmdAuxQueryAcceptorAuditPerformanceMeasures = 0x0C
	CmdAuxQueryDeviceCapabilities               = 0x0D
	CmdAuxQueryAcceptorApplicationID            = 0x0E
	CmdAuxQueryAcceptorVariantID                = 0x0F
	CmdAuxQueryBNFStatus                        = 0x10
	CmdAuxSetBezel                              = 0x11
)

const (
	CommunicationDisconnectTimeout = 3000
	PollingDisconnectTimeout       = 3000
	CalibrateTimeout               = 3000
)
