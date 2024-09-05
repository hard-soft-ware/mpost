package mpost

import "fmt"

// //////////////////////////////////
func (a *CAcceptor) verifyPropertyIsAllowed(capabilityFlag bool, propertyName string) error {
	if !a.connected {
		return fmt.Errorf("Calling %s not allowed when not connected.", propertyName)
	}

	if !capabilityFlag {
		return fmt.Errorf("Device does not support %s.", propertyName)
	}

	switch a.deviceState {
	case DownloadStart, Downloading:
		return fmt.Errorf("Calling %s not allowed during flash download.", propertyName)
	case CalibrateStart, Calibrating:
		return fmt.Errorf("Calling %s not allowed during calibration.", propertyName)
	}

	return nil
}

func (a *CAcceptor) GetDeviceSerialNumber() string {
	err := a.verifyPropertyIsAllowed(a.capDeviceSerialNumber, "DeviceSerialNumber")
	if err != nil {
		a.log.Debug().Err(err).Msg("GetDeviceSerialNumber")
		return ""
	}

	payload := []byte{CmdAuxiliary, 0, 0, CmdAuxQueryAcceptorSerialNumber}
	reply, err := a.SendSynchronousCommand(payload)
	if err != nil {
		a.log.Debug().Err(err).Msg("GetDeviceSerialNumber")
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
