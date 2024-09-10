package mpost

import (
	"errors"
	"fmt"
	"github.com/hard-soft-ware/mpost/acceptor"
	"github.com/hard-soft-ware/mpost/bill"
	"github.com/hard-soft-ware/mpost/consts"
	"github.com/hard-soft-ware/mpost/enum"
)

// //////////////////////////////////
func (a *CAcceptor) verifyPropertyIsAllowed(capabilityFlag bool, propertyName string) error {
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

func (a *CAcceptor) GetDeviceSerialNumber() string {
	a.log.Msg("GetDeviceSerialNumber")

	err := a.verifyPropertyIsAllowed(acceptor.Cap.DeviceSerialNumber, "DeviceSerialNumber")
	if err != nil {
		a.log.Err("GetDeviceSerialNumber", err)
		return ""
	}

	payload := []byte{consts.CmdAuxiliary.Byte(), 0, 0, consts.CmdAuxAcceptorSerialNumber.Byte()}
	reply, err := a.SendSynchronousCommand(payload)
	if err != nil {
		a.log.Err("GetDeviceSerialNumber", err)
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

func (a *CAcceptor) GetApplicationID() string {
	a.log.Msg("GetApplicationID")

	err := a.verifyPropertyIsAllowed(acceptor.Cap.ApplicationID, "GetApplicationID")
	if err != nil {
		a.log.Err("GetApplicationID", err)
		return ""
	}

	payload := []byte{consts.CmdAuxiliary.Byte(), 0, 0, consts.CmdAuxAcceptorApplicationID.Byte()}

	reply, err := a.SendSynchronousCommand(payload)
	if err != nil {
		a.log.Err("GetApplicationID", err)
		return ""
	}

	if len(reply) == 14 {
		s := string(reply[3:12])
		return s
	}

	return ""
}

func (a *CAcceptor) GetApplicationPN() string {
	a.log.Msg("GetApplicationPN")

	err := a.verifyPropertyIsAllowed(acceptor.Cap.ApplicationPN, "ApplicationPN")
	if err != nil {
		a.log.Err("ApplicationPN", err)
		return ""
	}

	payload := []byte{consts.CmdAuxiliary.Byte(), 0, 0, consts.CmdAuxAcceptorApplicationPartNumber.Byte()}

	reply, err := a.SendSynchronousCommand(payload)
	if err != nil {
		a.log.Err("GetApplicationPN", err)
		return ""
	}

	if len(reply) == 14 {
		s := string(reply[3:12])
		return s
	}

	return ""
}

func (a *CAcceptor) GetAuditLifeTimeTotals() []int {
	a.log.Msg("GetAuditLifeTimeTotals")
	values := []int{}

	err := a.verifyPropertyIsAllowed(acceptor.Cap.Audit, "GetAuditLifeTimeTotals")
	if err != nil {
		a.log.Err("GetAuditLifeTimeTotals", err)
		return values
	}

	payload := []byte{consts.CmdAuxiliary.Byte(), 0, 0, consts.CmdAuxAcceptorAuditLifeTimeTotals.Byte()}

	reply, err := a.SendSynchronousCommand(payload)
	if err != nil {
		a.log.Err("GetAuditLifeTimeTotals", err)
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

func (a *CAcceptor) GetAuditPerformance() []int {
	a.log.Msg("GetAuditPerformance")
	values := []int{}

	err := a.verifyPropertyIsAllowed(acceptor.Cap.Audit, "GetAuditPerformance")
	if err != nil {
		a.log.Err("GetAuditPerformance", err)
		return values
	}

	payload := []byte{consts.CmdAuxiliary.Byte(), 0, 0, consts.CmdAuxAcceptorAuditPerformanceMeasures.Byte()}

	reply, err := a.SendSynchronousCommand(payload)
	if err != nil {
		a.log.Err("GetAuditPerformance", err)
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

func (a *CAcceptor) GetAuditQP() []int {
	a.log.Msg("GetAuditQP")
	values := []int{}

	err := a.verifyPropertyIsAllowed(acceptor.Cap.Audit, "GetAuditQP")
	if err != nil {
		a.log.Err("GetAuditQP", err)
		return values
	}

	payload := []byte{consts.CmdAuxiliary.Byte(), 0, 0, consts.CmdAuxAcceptorAuditQPMeasures.Byte()}

	reply, err := a.SendSynchronousCommand(payload)
	if err != nil {
		a.log.Err("GetAuditQP", err)
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

func (a *CAcceptor) GetAutoStack() bool {
	a.log.Msg("GetAutoStack")
	return acceptor.AutoStack
}

func (a *CAcceptor) SetAutoStack(newVal bool) {
	a.log.Msg("SetAutoStack")
	acceptor.AutoStack = newVal
}

func (a *CAcceptor) GetBarCode() string {
	a.log.Msg("GetBarCode")
	return acceptor.BarCode
}

func (a *CAcceptor) GetBill() bill.BillStruct {
	a.log.Msg("GetBill")
	return bill.Bill
}

func (a *CAcceptor) GetBillTypes() []bill.BillStruct {
	a.log.Msg("GetBillTypes")
	return bill.Types
}

func (a *CAcceptor) GetBillTypeEnables() []bool {
	a.log.Msg("GetBillTypeEnables")
	return bill.TypeEnables
}

func (a *CAcceptor) SetBillTypeEnables(newVal []bool) {
	a.log.Msg("SetBillTypeEnables")

	if !acceptor.Connected {
		a.log.Err("SetBillTypeEnables", errors.New("calling BillTypeEnables not allowed when not connected"))
		return
	}

	if len(bill.TypeEnables) != len(bill.Types) {
		a.log.Err("SetBillTypeEnables", fmt.Errorf("CBillTypeEnables size must match BillTypes size"))
		return
	}

	bill.TypeEnables = newVal

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

func (a *CAcceptor) GetBillValues() []bill.BillStruct {
	a.log.Msg("GetBillValues")
	return bill.Values
}

func (a *CAcceptor) GetBillValueEnables() []bool {
	a.log.Msg("GetBillValueEnables")
	return bill.ValueEnables
}

func (a *CAcceptor) SetBillValueEnables(newVal []bool) {
	a.log.Msg("SetBillValueEnables")
	bill.ValueEnables = newVal

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

func (a *CAcceptor) GetBNFStatus() enum.BNFStatusType {
	a.log.Msg("Getting BNF status")
	err := a.verifyPropertyIsAllowed(acceptor.Cap.BNFStatus, "BNFStatus")

	if err != nil {
		a.log.Err("GetBNFStatus", err)
		return enum.BNFStatusUnknown
	}

	payload := []byte{consts.CmdAuxiliary.Byte(), 0, 0, consts.CmdAuxBNFStatus.Byte()}

	reply, err := a.SendSynchronousCommand(payload)
	if err != nil {
		a.log.Err("GetBNFStatus", err)
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
