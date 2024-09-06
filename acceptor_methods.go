package mpost

import (
	"fmt"
	"github.com/hard-soft-ware/mpost/acceptor"
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

func (a *CAcceptor) GetBNFStatus() enum.BNFStatusType {
	a.log.Msg("Getting BNF status")
	err := a.verifyPropertyIsAllowed(acceptor.Cap.BNFStatus, "BNFStatus")

	if err != nil {
		a.log.Err("GetBNFStatus", err)
		return enum.BNFStatusUnknown
	}

	payload := []byte{0x02, 0x00, 0x00, 0x03} // Assuming CmdAuxiliary is 0x02 and CmdAuxQueryBNFStatus is 0x03

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
