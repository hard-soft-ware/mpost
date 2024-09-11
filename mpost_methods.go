package mpost

import (
	"errors"
	"fmt"
	"github.com/hard-soft-ware/mpost/acceptor"
	"github.com/hard-soft-ware/mpost/bill"
	"github.com/hard-soft-ware/mpost/consts"
	"github.com/hard-soft-ware/mpost/enum"
	"github.com/hard-soft-ware/mpost/hook"
	"time"
)

// //////////////////////////////////
func (a *MpostObj) verifyPropertyIsAllowed(capabilityFlag bool, propertyName string) error {
	if !acceptor.Connected {
		return fmt.Errorf("Calling %s not allowed when not connected.", propertyName)
	}

	if !capabilityFlag {
		return fmt.Errorf("Device does not support %s.", propertyName)
	}

	switch acceptor.Device.State {
	case enum.StateDownloadStart, enum.StateDownloading:
		return fmt.Errorf("Calling %s not allowed during flash download.", propertyName)
	case enum.StateCalibrateStart, enum.StateCalibrating:
		return fmt.Errorf("Calling %s not allowed during calibration.", propertyName)
	}

	return nil
}

////

func (a *MpostObj) GetApplicationID() string {
	a.Log.Method("GetApplicationID", nil)

	err := a.verifyPropertyIsAllowed(acceptor.Cap.ApplicationID, "GetApplicationID")
	if err != nil {
		a.Log.Err("GetApplicationID", err)
		return ""
	}

	payload := []byte{consts.CmdAuxiliary.Byte(), 0, 0, consts.CmdAuxAcceptorApplicationID.Byte()}

	reply, err := a.SendSynchronousCommand(payload)
	if err != nil {
		a.Log.Err("GetApplicationID", err)
		return ""
	}

	if len(reply) == 14 {
		s := string(reply[3:12])
		return s
	}

	return ""
}

func (a *MpostObj) GetApplicationPN() string {
	a.Log.Method("GetApplicationPN", nil)

	err := a.verifyPropertyIsAllowed(acceptor.Cap.ApplicationPN, "ApplicationPN")
	if err != nil {
		a.Log.Err("ApplicationPN", err)
		return ""
	}

	payload := []byte{consts.CmdAuxiliary.Byte(), 0, 0, consts.CmdAuxAcceptorApplicationPartNumber.Byte()}

	reply, err := a.SendSynchronousCommand(payload)
	if err != nil {
		a.Log.Err("GetApplicationPN", err)
		return ""
	}

	if len(reply) == 14 {
		s := string(reply[3:12])
		return s
	}

	return ""
}

func (a *MpostObj) GetAuditLifeTimeTotals() []int {
	a.Log.Method("GetAuditLifeTimeTotals", nil)
	values := []int{}

	err := a.verifyPropertyIsAllowed(acceptor.Cap.Audit, "GetAuditLifeTimeTotals")
	if err != nil {
		a.Log.Err("GetAuditLifeTimeTotals", err)
		return values
	}

	payload := []byte{consts.CmdAuxiliary.Byte(), 0, 0, consts.CmdAuxAcceptorAuditLifeTimeTotals.Byte()}

	reply, err := a.SendSynchronousCommand(payload)
	if err != nil {
		a.Log.Err("GetAuditLifeTimeTotals", err)
		return values
	}

	if len(reply) < 13 || ((len(reply)-5)%8 != 0) {
		return values
	}

	fieldCount := (len(reply) - 5) / 8
	for i := 0; i < fieldCount; i++ {
		offset := 8*i + 3
		value := (int((reply)[offset+0]&0x0F) << 28) +
			(int((reply)[offset+1]&0x0F) << 24) +
			(int((reply)[offset+2]&0x0F) << 20) +
			(int((reply)[offset+3]&0x0F) << 16) +
			(int((reply)[offset+4]&0x0F) << 12) +
			(int((reply)[offset+5]&0x0F) << 8) +
			(int((reply)[offset+6]&0x0F) << 4) +
			int((reply)[offset+7]&0x0F)

		values = append(values, value)
	}

	return values
}

func (a *MpostObj) GetAuditPerformance() []int {
	a.Log.Method("GetAuditPerformance", nil)
	values := []int{}

	err := a.verifyPropertyIsAllowed(acceptor.Cap.Audit, "GetAuditPerformance")
	if err != nil {
		a.Log.Err("GetAuditPerformance", err)
		return values
	}

	payload := []byte{consts.CmdAuxiliary.Byte(), 0, 0, consts.CmdAuxAcceptorAuditPerformanceMeasures.Byte()}

	reply, err := a.SendSynchronousCommand(payload)
	if err != nil {
		a.Log.Err("GetAuditPerformance", err)
		return values
	}

	if len(reply) < 9 || ((len(reply)-5)%4 != 0) {
		return values
	}

	fieldCount := (len(reply) - 5) / 4

	for i := 0; i < fieldCount; i++ {
		offset := 4*i + 3
		value := (int((reply)[offset+0]&0x0F) << 12) +
			(int((reply)[offset+1]&0x0F) << 8) +
			(int((reply)[offset+2]&0x0F) << 4) +
			int((reply)[offset+3]&0x0F)

		values = append(values, value)
	}

	return values
}

func (a *MpostObj) GetAuditQP() []int {
	a.Log.Method("GetAuditQP", nil)
	values := []int{}

	err := a.verifyPropertyIsAllowed(acceptor.Cap.Audit, "GetAuditQP")
	if err != nil {
		a.Log.Err("GetAuditQP", err)
		return values
	}

	payload := []byte{consts.CmdAuxiliary.Byte(), 0, 0, consts.CmdAuxAcceptorAuditQPMeasures.Byte()}

	reply, err := a.SendSynchronousCommand(payload)
	if err != nil {
		a.Log.Err("GetAuditQP", err)
		return values
	}

	if len(reply) < 9 || ((len(reply)-5)%4 != 0) {
		return values
	}

	fieldCount := (len(reply) - 5) / 4

	for i := 0; i < fieldCount; i++ {
		offset := 4*i + 3
		value := (int(reply[offset+0]&0x0F) << 12) +
			(int(reply[offset+1]&0x0F) << 8) +
			(int(reply[offset+2]&0x0F) << 4) +
			int(reply[offset+3]&0x0F)

		values = append(values, value)
	}

	return values
}

func (a *MpostObj) GetAutoStack() bool {
	a.Log.Method("GetAutoStack", nil)
	return acceptor.AutoStack
}

func (a *MpostObj) SetAutoStack(v bool) {
	a.Log.Method("SetAutoStack", nil)
	acceptor.AutoStack = v
}

func (a *MpostObj) GetBarCode() string {
	a.Log.Method("GetBarCode", nil)
	return acceptor.BarCode
}

func (a *MpostObj) GetBill() bill.BillStruct {
	a.Log.Method("GetBill", nil)
	return bill.Bill
}

func (a *MpostObj) GetBillTypes() []bill.BillStruct {
	a.Log.Method("GetBillTypes", nil)
	return bill.Types
}

func (a *MpostObj) GetBillTypeEnables() []bool {
	a.Log.Method("GetBillTypeEnables", nil)
	return bill.TypeEnables
}

func (a *MpostObj) SetBillTypeEnables(v []bool) {
	a.Log.Method("SetBillTypeEnables", nil)

	if !acceptor.Connected {
		a.Log.Err("SetBillTypeEnables", errors.New("calling BillTypeEnables not allowed when not connected"))
		return
	}

	if len(bill.TypeEnables) != len(bill.Types) {
		a.Log.Err("SetBillTypeEnables", fmt.Errorf("CBillTypeEnables size must match BillTypes size"))
		return
	}

	bill.TypeEnables = v

	if acceptor.ExpandedNoteReporting {
		payload := make([]byte, 15)
		acceptor.ConstructOmnibusCommand(payload, consts.CmdExpanded, 2, bill.TypeEnables)
		payload[1] = 0x03 // Sub Type

		for i, enable := range bill.TypeEnables {
			enableIndex := i / 7
			bitPosition := i % 7
			bit := 1 << bitPosition

			if enable {
				payload[5+enableIndex] |= byte(bit)
			}
		}

		a.SendAsynchronousCommand(payload)
	}
}

func (a *MpostObj) GetBillValues() []bill.BillStruct {
	a.Log.Method("GetBillValues", nil)
	return bill.Values
}

func (a *MpostObj) GetBillValueEnables() []bool {
	a.Log.Method("GetBillValueEnables", nil)
	return bill.ValueEnables
}

func (a *MpostObj) SetBillValueEnables(v []bool) {
	a.Log.Method("SetBillValueEnables", nil)
	bill.ValueEnables = v

	for _, enabled := range bill.ValueEnables {
		for j, billType := range bill.Types {
			if billType.Value == bill.Values[j].Value && billType.Country == bill.Values[j].Country {
				bill.TypeEnables[j] = enabled
			}
		}
	}

	payload := make([]byte, 15)
	acceptor.ConstructOmnibusCommand(payload, consts.CmdExpanded, 2, bill.TypeEnables)
	payload[1] = 0x03 // Sub Type

	for i, enable := range bill.TypeEnables {
		enableIndex := i / 7
		bitPosition := i % 7
		bit := 1 << bitPosition

		if enable {
			payload[5+enableIndex] |= byte(bit)
		}
	}

	a.SendAsynchronousCommand(payload)
}

func (a *MpostObj) GetBNFStatus() enum.BNFStatusType {
	a.Log.Method("Getting BNF status", nil)
	err := a.verifyPropertyIsAllowed(acceptor.Cap.BNFStatus, "BNFStatus")

	if err != nil {
		a.Log.Err("GetBNFStatus", err)
		return enum.BNFStatusUnknown
	}

	payload := []byte{consts.CmdAuxiliary.Byte(), 0, 0, consts.CmdAuxBNFStatus.Byte()}

	reply, err := a.SendSynchronousCommand(payload)
	if err != nil {
		a.Log.Err("GetBNFStatus", err)
		return enum.BNFStatusUnknown
	}

	if len(reply) == 9 {
		if reply[3] == 0 {
			return enum.BNFStatusNotAttached
		} else {
			if reply[4] == 0 {
				return enum.BNFStatusOK
			} else {
				return enum.BNFStatusError
			}
		}
	}

	return enum.BNFStatusUnknown
}

func (a *MpostObj) GetBootPN() string {
	a.Log.Method("GetBootPN", nil)

	err := a.verifyPropertyIsAllowed(acceptor.Cap.BootPN, "GetBootPN")
	if err != nil {
		a.Log.Err("GetBootPN", err)
		return ""
	}

	payload := []byte{consts.CmdAuxiliary.Byte(), 0, 0, consts.CmdAuxAcceptorBootPartNumber.Byte()}

	reply, err := a.SendSynchronousCommand(payload)
	if err != nil {
		a.Log.Err("GetBootPN", err)
		return ""
	}

	if len(reply) == 14 {
		s := string(reply[3:12]) // Extracting the substring from byte slice
		return s
	}

	return ""
}

func (a *MpostObj) GetCapApplicationID() bool {
	a.Log.Method("GetCapApplicationID", nil)
	return acceptor.Cap.ApplicationID
}

func (a *MpostObj) GetCapApplicationPN() bool {
	a.Log.Method("GetCapApplicationPN", nil)
	return acceptor.Cap.ApplicationPN
}

func (a *MpostObj) GetCapAssetNumber() bool {
	a.Log.Method("GetCapAssetNumber", nil)
	return acceptor.Cap.AssetNumber
}

func (a *MpostObj) GetCapAudit() bool {
	a.Log.Method("GetCapAudit", nil)
	return acceptor.Cap.Audit
}

func (a *MpostObj) GetCapBarCodes() bool {
	a.Log.Method("GetCapBarCodes", nil)
	return acceptor.Cap.BarCodes
}

func (a *MpostObj) GetCapBarCodesExt() bool {
	a.Log.Method("GetCapBarCodesExt", nil)
	return acceptor.Cap.BarCodesExt
}

func (a *MpostObj) GetCapBNFStatus() bool {
	a.Log.Method("GetCapBNFStatus", nil)
	return acceptor.Cap.BNFStatus
}

func (a *MpostObj) GetCapBookmark() bool {
	a.Log.Method("GetCapBookmark", nil)
	return acceptor.Cap.Bookmark
}

func (a *MpostObj) GetCapBootPN() bool {
	a.Log.Method("GetCapBootPN", nil)
	return acceptor.Cap.BootPN
}

func (a *MpostObj) GetCapCalibrate() bool {
	a.Log.Method("GetCapCalibrate", nil)
	return acceptor.Cap.Calibrate
}

func (a *MpostObj) GetCapCashBoxTotal() bool {
	a.Log.Method("GetCapCashBoxTotal", nil)
	return acceptor.Cap.CashBoxTotal
}

func (a *MpostObj) GetCapCouponExt() bool {
	a.Log.Method("GetCapCouponExt", nil)
	return acceptor.Cap.CouponExt
}

func (a *MpostObj) GetCapDevicePaused() bool {
	a.Log.Method("GetCapDevicePaused", nil)
	return acceptor.Cap.DevicePaused
}

func (a *MpostObj) GetCapDeviceSoftReset() bool {
	a.Log.Method("GetCapDeviceSoftReset", nil)
	return acceptor.Cap.DeviceSoftReset
}

func (a *MpostObj) GetCapDeviceType() bool {
	a.Log.Method("GetCapDeviceType", nil)
	return acceptor.Cap.DeviceType
}

func (a *MpostObj) GetCapDeviceResets() bool {
	a.Log.Method("GetCapDeviceResets", nil)
	return acceptor.Cap.DeviceResets
}

func (a *MpostObj) GetCapDeviceSerialNumber() bool {
	a.Log.Method("GetCapDeviceSerialNumber", nil)
	return acceptor.Cap.DeviceSerialNumber
}

func (a *MpostObj) GetCapEscrowTimeout() bool {
	a.Log.Method("GetCapEscrowTimeout", nil)
	return acceptor.Cap.EscrowTimeout
}

func (a *MpostObj) GetCapFlashDownload() bool {
	a.Log.Method("GetCapFlashDownload", nil)
	return acceptor.Cap.FlashDownload
}

func (a *MpostObj) GetCapNoPush() bool {
	a.Log.Method("GetCapNoPush", nil)
	return acceptor.Cap.NoPush
}

func (a *MpostObj) GetCapOrientationExt() bool {
	a.Log.Method("GetCapOrientationExt", nil)
	return acceptor.Cap.OrientationExt
}

func (a *MpostObj) GetCapPupExt() bool {
	a.Log.Method("GetCapPupExt", nil)
	return acceptor.Cap.PupExt
}

func (a *MpostObj) GetCapTestDoc() bool {
	a.Log.Method("GetCapTestDoc", nil)
	return acceptor.Cap.TestDoc
}

func (a *MpostObj) GetCapVariantID() bool {
	a.Log.Method("GetCapVariantID", nil)
	return acceptor.Cap.VariantID
}

func (a *MpostObj) GetCapVariantPN() bool {
	a.Log.Method("GetCapVariantPN", nil)
	return acceptor.Cap.VariantPN
}

func (a *MpostObj) GetCashBoxAttached() bool {
	a.Log.Method("GetCashBoxAttached", nil)
	return acceptor.Cash.BoxAttached
}

func (a *MpostObj) GetCashBoxFull() bool {
	a.Log.Method("GetCashBoxFull", nil)
	return acceptor.Cash.BoxFull
}

func (a *MpostObj) GetCashBoxTotal() int {
	a.Log.Method("GetCashBoxTotal", nil)

	err := a.verifyPropertyIsAllowed(acceptor.Cap.CashBoxTotal, "GetCashBoxTotal")
	if err != nil {
		a.Log.Err("GetCashBoxTotal", err)
		return 0
	}

	payload := []byte{consts.CmdOmnibus.Byte(), 0x7F, 0x3C, 0x02}

	reply, err := a.SendSynchronousCommand(payload)
	if err != nil {
		a.Log.Err("GetCashBoxTotal", err)
		return 0
	}

	if len(reply) < 9 {
		return 0
	}

	total := int(reply[3]&0x0F)<<20 |
		int(reply[4]&0x0F)<<16 |
		int(reply[5]&0x0F)<<12 |
		int(reply[6]&0x0F)<<8 |
		int(reply[7]&0x0F)<<4 |
		int(reply[8]&0x0F)

	return total
}

func (a *MpostObj) GetConnected() bool {
	a.Log.Method("GetConnected", nil)
	return acceptor.Connected
}

func (a *MpostObj) GetCoupon() *CouponObj {
	a.Log.Method("GetCoupon", nil)
	return a.coupon
}

func (a *MpostObj) GetDeviceBusy() bool {
	a.Log.Method("GetDeviceBusy", nil)
	return acceptor.Device.State != enum.StateIdling
}

func (a *MpostObj) GetDeviceCRC() int64 {
	a.Log.Method("GetDeviceCRC", nil)

	err := a.verifyPropertyIsAllowed(true, "DeviceCRC")
	if err != nil {
		a.Log.Err("GetDeviceCRC", err)
		return 0
	}

	payload := []byte{consts.CmdAuxiliary.Byte(), 0, 0, consts.CmdAuxSoftwareCRC.Byte()}

	reply, err := a.SendSynchronousCommand(payload)
	if err != nil {
		a.Log.Err("GetDeviceCRC", err)
		return 0
	}

	if len(reply) < 7 {
		return 0
	}

	crc := int64(reply[3]&0x0F)<<12 |
		int64(reply[4]&0x0F)<<8 |
		int64(reply[5]&0x0F)<<4 |
		int64(reply[6]&0x0F)

	return crc
}

func (a *MpostObj) GetDeviceFailure() bool {
	a.Log.Method("GetDeviceFailure", nil)
	return acceptor.Device.State == enum.StateFailed
}

func (a *MpostObj) GetDeviceJammed() bool {
	a.Log.Method("GetDeviceJammed", nil)
	return acceptor.Device.Jammed
}

func (a *MpostObj) GetDeviceModel() int {
	a.Log.Method("GetDeviceModel", nil)
	return acceptor.Device.Model
}

func (a *MpostObj) GetDevicePaused() bool {
	a.Log.Method("GetDevicePaused", nil)
	return acceptor.Device.Paused
}

func (a *MpostObj) GetDevicePortName() string {
	a.Log.Method("GetDevicePortName", nil)
	return a.port.PortName
}

func (a *MpostObj) GetDevicePowerUp() enum.PowerUpType {
	a.Log.Method("GetDevicePowerUp", nil)
	return acceptor.Device.PowerUp
}

func (a *MpostObj) GetDeviceResets() int {
	a.Log.Method("GetDeviceResets", nil)

	err := a.verifyPropertyIsAllowed(acceptor.Cap.DeviceResets, "DeviceResets")
	if err != nil {
		a.Log.Err("GetDeviceResets", err)
		return 0
	}

	payload := []byte{consts.CmdAuxiliary.Byte(), 0, 0, consts.CmdAuxDeviceResets.Byte()}

	reply, err := a.SendSynchronousCommand(payload)
	if err != nil {
		a.Log.Err("GetDeviceResets", err)
		return 0
	}

	if len(reply) < 9 {
		return 0
	}

	resets := int(reply[3]&0x0F)<<20 |
		int(reply[4]&0x0F)<<16 |
		int(reply[5]&0x0F)<<12 |
		int(reply[6]&0x0F)<<8 |
		int(reply[7]&0x0F)<<4 |
		int(reply[8]&0x0F)

	return resets
}

func (a *MpostObj) GetDeviceRevision() int {
	a.Log.Method("GetDeviceRevision", nil)
	return acceptor.Device.Revision
}

func (a *MpostObj) GetDeviceSerialNumber() string {
	a.Log.Method("GetDeviceSerialNumber", nil)

	err := a.verifyPropertyIsAllowed(acceptor.Cap.DeviceSerialNumber, "DeviceSerialNumber")
	if err != nil {
		a.Log.Err("GetDeviceSerialNumber", err)
		return ""
	}

	payload := []byte{consts.CmdAuxiliary.Byte(), 0, 0, consts.CmdAuxAcceptorSerialNumber.Byte()}
	reply, err := a.SendSynchronousCommand(payload)
	if err != nil {
		a.Log.Err("GetDeviceSerialNumber", err)
		return ""
	}

	validCharIndex := 3
	for validCharIndex < len(reply) && reply[validCharIndex] > 0x20 && reply[validCharIndex] < 0x7F && validCharIndex <= 22 {
		validCharIndex++
	}
	returnedStringLength := validCharIndex - 3

	s := string(reply[3 : 3+returnedStringLength])
	return s
}

func (a *MpostObj) GetDeviceStalled() bool {
	a.Log.Method("GetDeviceStalled", nil)
	return acceptor.Device.Stalled
}

func (a *MpostObj) GetDeviceState() enum.StateType {
	a.Log.Method("GetDeviceState", nil)
	return acceptor.Device.State
}

func (a *MpostObj) GetDeviceType() string {
	a.Log.Method("GetDeviceType", nil)

	err := a.verifyPropertyIsAllowed(acceptor.Cap.DeviceType, "DeviceType")
	if err != nil {
		a.Log.Err("GetDeviceType", err)
		return ""
	}

	payload := []byte{consts.CmdAuxiliary.Byte(), 0, 0, consts.CmdAuxAcceptorType.Byte()}

	reply, err := a.SendSynchronousCommand(payload)
	if err != nil {
		a.Log.Err("GetDeviceType", err)
		return ""
	}

	// Specified to check for non-printable characters from 0x00 to 0x1F or 0x7F as termination.
	validCharIndex := 3
	for ; validCharIndex < len(reply) && reply[validCharIndex] > 0x20 && reply[validCharIndex] < 0x7F && validCharIndex <= 22; validCharIndex++ {
	}

	returnedStringLength := validCharIndex - 3
	if returnedStringLength > 0 {
		return string(reply[3 : 3+returnedStringLength])
	}

	return ""
}

func (a *MpostObj) GetDocType() enum.DocumentType {
	a.Log.Method("GetDocType", nil)
	return a.DocType
}

func (a *MpostObj) GetTransactionTimeout() time.Duration {
	a.Log.Method("GetTransactionTimeout", nil)
	return acceptor.Timeout.Transaction
}

func (a *MpostObj) SetTransactionTimeout(v time.Duration) {
	a.Log.Method("SetTransactionTimeout", nil)
	acceptor.Timeout.Transaction = v
}

func (a *MpostObj) GetDownloadTimeout() time.Duration {
	a.Log.Method("GetDownloadTimeout", nil)
	return acceptor.Timeout.Download
}

func (a *MpostObj) SetDownloadTimeout(v time.Duration) {
	a.Log.Method("SetDownloadTimeout", nil)
	acceptor.Timeout.Download = v
}

func (a *MpostObj) GetEnableAcceptance() bool {
	a.Log.Method("GetEnableAcceptance", nil)
	return acceptor.Enable.Acceptance
}

func (a *MpostObj) SetEnableAcceptance(v bool) {
	a.Log.Method("SetEnableAcceptance", nil)
	acceptor.Enable.Acceptance = v
}

func (a *MpostObj) GetEnableBarCodes() bool {
	a.Log.Method("GetEnableBarCodes", nil)
	return acceptor.Enable.BarCodes
}

func (a *MpostObj) SetEnableBarCodes(v bool) {
	a.Log.Method("SetEnableBarCodes", nil)
	acceptor.Enable.BarCodes = v
}

func (a *MpostObj) GetEnableBookmarks() bool {
	a.Log.Method("GetEnableBookmarks", nil)
	return acceptor.Enable.Bookmarks
}

func (a *MpostObj) SetEnableBookmarks(v bool) {
	a.Log.Method("SetEnableBookmarks", nil)
	acceptor.Enable.Bookmarks = v
}

func (a *MpostObj) GetEnableCouponExt() bool {
	a.Log.Method("GetEnableCouponExt", nil)
	return acceptor.Enable.CouponExt
}

func (a *MpostObj) SetEnableCouponExt(v bool) {
	a.Log.Method("SetEnableCouponExt", nil)
	acceptor.Enable.CouponExt = v
}

func (a *MpostObj) GetEnableNoPush() bool {
	a.Log.Method("GetEnableNoPush", nil)
	return acceptor.Enable.NoPush
}

func (a *MpostObj) SetEnableNoPush(v bool) {
	a.Log.Method("SetEnableNoPush", nil)
	acceptor.Enable.NoPush = v
}

func (a *MpostObj) GetEscrowOrientation() enum.OrientationType {
	a.Log.Method("GetEscrowOrientation", nil)
	if acceptor.Cap.OrientationExt {
		return acceptor.EscrowOrientation
	}
	return enum.OrientationUnknownOrientation
}

func (a *MpostObj) GetHighSecurity() bool {
	a.Log.Method("GetHighSecurity", nil)
	return acceptor.HighSecurity
}

func (a *MpostObj) SetHighSecurity(v bool) {
	a.Log.Method("SetHighSecurity", nil)
	acceptor.HighSecurity = v
}

func (a *MpostObj) GetOrientationControl() enum.OrientationControlType {
	a.Log.Method("GetOrientationControl", nil)
	return acceptor.OrientationCtl
}

func (a *MpostObj) SetOrientationControl(v enum.OrientationControlType) {
	a.Log.Method("SetOrientationControl", nil)
	acceptor.OrientationCtl = v
}

func (a *MpostObj) GetOrientationCtlExt() enum.OrientationControlType {
	a.Log.Method("GetOrientationCtlExt", nil)
	return acceptor.OrientationCtlExt
}

func (a *MpostObj) SetOrientationCtlExt(v enum.OrientationControlType) {
	a.Log.Method("SetOrientationCtlExt", nil)
	acceptor.OrientationCtlExt = v
}

func (a *MpostObj) GetVariantNames() []string {
	a.Log.Method("GetVariantNames", nil)

	err := a.verifyPropertyIsAllowed(true, "VariantNames")
	if err != nil {
		a.Log.Err("GetVariantNames", err)
		return nil
	}

	payload := []byte{consts.CmdAuxiliary.Byte(), 0, 0, consts.CmdAuxAcceptorVariantName.Byte()}

	reply, err := a.SendSynchronousCommand(payload)
	if err != nil {
		a.Log.Err("GetVariantNames", err)
		return nil
	}

	var names []string
	validCharIndex := 3

	for validCharIndex < len(reply) && reply[validCharIndex] > 0x20 && reply[validCharIndex] < 0x7F && validCharIndex <= 34 {
		if validCharIndex+2 < len(reply) {
			names = append(names, string(reply[validCharIndex:validCharIndex+3]))
			validCharIndex += 4
		} else {
			break
		}
	}

	return names
}

func (a *MpostObj) GetVariantID() string {
	a.Log.Method("GetVariantID", nil)

	err := a.verifyPropertyIsAllowed(acceptor.Cap.VariantID, "VariantID")
	if err != nil {
		a.Log.Err("GetVariantID", err)
		return ""
	}

	payload := []byte{consts.CmdAuxiliary.Byte(), 0, 0, consts.CmdAuxAcceptorVariantID.Byte()}

	reply, err := a.SendSynchronousCommand(payload)
	if err != nil {
		a.Log.Err("GetVariantID", err)
		return ""
	}

	if len(reply) == 14 {
		return string(reply[3:12]) // Extracting a 9-byte string starting from index 3
	}

	return ""
}

func (a *MpostObj) GetVariantPN() string {
	a.Log.Method("GetVariantPN", nil)

	err := a.verifyPropertyIsAllowed(acceptor.Cap.VariantPN, "VariantPN")
	if err != nil {
		a.Log.Err("GetVariantPN", err)
		return ""
	}

	payload := []byte{consts.CmdAuxiliary.Byte(), 0, 0, consts.CmdAuxAcceptorVariantPartNumber.Byte()}

	reply, err := a.SendSynchronousCommand(payload)
	if err != nil {
		a.Log.Err("GetVariantPN", err)
		return ""
	}

	if len(reply) == 14 {
		return string(reply[3:12])
	}

	return ""
}

func (a *MpostObj) GetVersion() string {
	a.Log.Method("GetVersion", nil)
	return acceptor.Version
}

func (a *MpostObj) Calibrate() {
	a.Log.Method("Calibrate", nil)
	if !acceptor.Connected {
		a.Log.Err("Calibrate", errors.New("Calibrate called when not connected"))
		return
	}

	if acceptor.Device.State != enum.StateIdling {
		a.Log.Err("Calibrate", errors.New("Calibrate allowed only when DeviceState == Idling"))
		return
	}

	payload := []byte{consts.CmdCalibrate.Byte(), 0x00, 0x00, 0x00}

	acceptor.SuppressStandardPoll = true
	acceptor.Device.State = enum.StateCalibrateStart

	hook.Raise.Calibrate.Start()

	hook.CalibrateProgress = true

	startTickCount := time.Now()

	for {
		reply, err := a.SendSynchronousCommand(payload)
		if err != nil {
			a.Log.Err("Calibrate", errors.New("Failed to send synchronous command during calibration"))
			return
		}

		if len(reply) == 11 && (reply[2]&0x70) == 0x40 {
			break
		}

		if time.Since(startTickCount) > CalibrateTimeout {
			hook.Raise.Calibrate.Finish()
			return
		}
	}
}

func (a *MpostObj) ClearCashBoxTotal() (err error) {
	a.Log.Method("ClearCashBoxTotal", nil)

	if !acceptor.Connected {
		err = errors.New("ClearCashBoxTotal called when not connected")
		a.Log.Err("ClearCashBoxTotal", err)
		return
	}

	payload := []byte{consts.CmdCalibrate.Byte(), 0x00, 0x00, consts.CmdAuxCashBoxTotal.Byte()}

	reply, err := a.SendSynchronousCommand(payload)
	if err != nil {
		a.Log.Err("ClearCashBoxTotal", err)
		return
	}

	a.dataLinkLayer.ProcessReply(reply)
	return
}
