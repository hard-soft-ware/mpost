package consts

/* This file is automatically generated */

type CmdAuxType byte

const (
	CmdAuxSoftwareCRC                   CmdAuxType = 0x00
	CmdAuxCashBoxTotal                  CmdAuxType = 0x01
	CmdAuxDeviceResets                  CmdAuxType = 0x02
	CmdAuxBNFStatus                     CmdAuxType = 0x10
	CmdAuxSetBezel                      CmdAuxType = 0x11
	CmdAuxDeviceCapabilities            CmdAuxType = 0x0d
	CmdAuxCleDeviceResetsarCashBoxTotal CmdAuxType = 0x03

	CmdAuxAcceptorType                     CmdAuxType = 0x04
	CmdAuxAcceptorApplicationID            CmdAuxType = 0x0e
	CmdAuxAcceptorVariantID                CmdAuxType = 0x0f
	CmdAuxAcceptorSerialNumber             CmdAuxType = 0x05
	CmdAuxAcceptorBootPartNumber           CmdAuxType = 0x06
	CmdAuxAcceptorApplicationPartNumber    CmdAuxType = 0x07
	CmdAuxAcceptorVariantName              CmdAuxType = 0x08
	CmdAuxAcceptorVariantPartNumber        CmdAuxType = 0x09
	CmdAuxAcceptorAuditLifeTimeTotals      CmdAuxType = 0x0a
	CmdAuxAcceptorAuditQPMeasures          CmdAuxType = 0x0b
	CmdAuxAcceptorAuditPerformanceMeasures CmdAuxType = 0x0c
)

const (
	CmdAuxTextSoftwareCRC                   = "SoftwareCRC"
	CmdAuxTextCashBoxTotal                  = "CashBoxTotal"
	CmdAuxTextDeviceResets                  = "DeviceResets"
	CmdAuxTextBNFStatus                     = "BNFStatus"
	CmdAuxTextSetBezel                      = "SetBezel"
	CmdAuxTextDeviceCapabilities            = "DeviceCapabilities"
	CmdAuxTextCleDeviceResetsarCashBoxTotal = "CleDeviceResetsarCashBoxTotal"

	CmdAuxTextAcceptorType                     = "AcceptorType"
	CmdAuxTextAcceptorApplicationID            = "AcceptorApplicationID"
	CmdAuxTextAcceptorVariantID                = "AcceptorVariantID"
	CmdAuxTextAcceptorSerialNumber             = "AcceptorSerialNumber"
	CmdAuxTextAcceptorBootPartNumber           = "AcceptorBootPartNumber"
	CmdAuxTextAcceptorApplicationPartNumber    = "AcceptorApplicationPartNumber"
	CmdAuxTextAcceptorVariantName              = "AcceptorVariantName"
	CmdAuxTextAcceptorVariantPartNumber        = "AcceptorVariantPartNumber"
	CmdAuxTextAcceptorAuditLifeTimeTotals      = "AcceptorAuditLifeTimeTotals"
	CmdAuxTextAcceptorAuditQPMeasures          = "AcceptorAuditQPMeasures"
	CmdAuxTextAcceptorAuditPerformanceMeasures = "AcceptorAuditPerformanceMeasures"
)

var CmdAuxMap = map[CmdAuxType]string{
	CmdAuxSoftwareCRC:                      CmdAuxTextSoftwareCRC,
	CmdAuxCashBoxTotal:                     CmdAuxTextCashBoxTotal,
	CmdAuxDeviceResets:                     CmdAuxTextDeviceResets,
	CmdAuxBNFStatus:                        CmdAuxTextBNFStatus,
	CmdAuxSetBezel:                         CmdAuxTextSetBezel,
	CmdAuxDeviceCapabilities:               CmdAuxTextDeviceCapabilities,
	CmdAuxCleDeviceResetsarCashBoxTotal:    CmdAuxTextCleDeviceResetsarCashBoxTotal,
	CmdAuxAcceptorType:                     CmdAuxTextAcceptorType,
	CmdAuxAcceptorApplicationID:            CmdAuxTextAcceptorApplicationID,
	CmdAuxAcceptorVariantID:                CmdAuxTextAcceptorVariantID,
	CmdAuxAcceptorSerialNumber:             CmdAuxTextAcceptorSerialNumber,
	CmdAuxAcceptorBootPartNumber:           CmdAuxTextAcceptorBootPartNumber,
	CmdAuxAcceptorApplicationPartNumber:    CmdAuxTextAcceptorApplicationPartNumber,
	CmdAuxAcceptorVariantName:              CmdAuxTextAcceptorVariantName,
	CmdAuxAcceptorVariantPartNumber:        CmdAuxTextAcceptorVariantPartNumber,
	CmdAuxAcceptorAuditLifeTimeTotals:      CmdAuxTextAcceptorAuditLifeTimeTotals,
	CmdAuxAcceptorAuditQPMeasures:          CmdAuxTextAcceptorAuditQPMeasures,
	CmdAuxAcceptorAuditPerformanceMeasures: CmdAuxTextAcceptorAuditPerformanceMeasures,
}

func (obj CmdAuxType) String() string {
	val, ok := CmdAuxMap[obj]
	if ok {
		return val
	}
	return "Unknown CmdAuxType"
}

func (obj CmdAuxType) Byte() byte {
	return byte(obj)
}