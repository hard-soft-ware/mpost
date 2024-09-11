package mpost

import (
	"github.com/hard-soft-ware/mpost/acceptor"
	"github.com/hard-soft-ware/mpost/consts"
	"github.com/hard-soft-ware/mpost/enum"
)

////////////////////////////////////

func (m *MethodsObj) GetCapDevicePaused() bool {
	m.a.Log.Method("GetCapDevicePaused", nil)
	return acceptor.Cap.DevicePaused
}

func (m *MethodsObj) GetCapDeviceSoftReset() bool {
	m.a.Log.Method("GetCapDeviceSoftReset", nil)
	return acceptor.Cap.DeviceSoftReset
}

func (m *MethodsObj) GetCapDeviceType() bool {
	m.a.Log.Method("GetCapDeviceType", nil)
	return acceptor.Cap.DeviceType
}

func (m *MethodsObj) GetCapDeviceResets() bool {
	m.a.Log.Method("GetCapDeviceResets", nil)
	return acceptor.Cap.DeviceResets
}

func (m *MethodsObj) GetCapDeviceSerialNumber() bool {
	m.a.Log.Method("GetCapDeviceSerialNumber", nil)
	return acceptor.Cap.DeviceSerialNumber
}

func (m *MethodsObj) GetDeviceBusy() bool {
	m.a.Log.Method("GetDeviceBusy", nil)
	return acceptor.Device.State != enum.StateIdling
}

func (m *MethodsObj) GetDeviceCRC() int64 {
	m.a.Log.Method("GetDeviceCRC", nil)

	err := acceptor.Verify(true, "DeviceCRC")
	if err != nil {
		m.a.Log.Err("GetDeviceCRC", err)
		return 0
	}

	payload := []byte{consts.CmdAuxiliary.Byte(), 0, 0, consts.CmdAuxSoftwareCRC.Byte()}

	reply, err := m.a.SendSynchronousCommand(payload)
	if err != nil {
		m.a.Log.Err("GetDeviceCRC", err)
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

func (m *MethodsObj) GetDeviceFailure() bool {
	m.a.Log.Method("GetDeviceFailure", nil)
	return acceptor.Device.State == enum.StateFailed
}

func (m *MethodsObj) GetDeviceJammed() bool {
	m.a.Log.Method("GetDeviceJammed", nil)
	return acceptor.Device.Jammed
}

func (m *MethodsObj) GetDeviceModel() int {
	m.a.Log.Method("GetDeviceModel", nil)
	return acceptor.Device.Model
}

func (m *MethodsObj) GetDevicePaused() bool {
	m.a.Log.Method("GetDevicePaused", nil)
	return acceptor.Device.Paused
}

func (m *MethodsObj) GetDevicePortName() string {
	m.a.Log.Method("GetDevicePortName", nil)
	return m.a.port.PortName
}

func (m *MethodsObj) GetDevicePowerUp() enum.PowerUpType {
	m.a.Log.Method("GetDevicePowerUp", nil)
	return acceptor.Device.PowerUp
}

func (m *MethodsObj) GetDeviceResets() int {
	m.a.Log.Method("GetDeviceResets", nil)

	err := acceptor.Verify(acceptor.Cap.DeviceResets, "DeviceResets")
	if err != nil {
		m.a.Log.Err("GetDeviceResets", err)
		return 0
	}

	payload := []byte{consts.CmdAuxiliary.Byte(), 0, 0, consts.CmdAuxDeviceResets.Byte()}

	reply, err := m.a.SendSynchronousCommand(payload)
	if err != nil {
		m.a.Log.Err("GetDeviceResets", err)
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

func (m *MethodsObj) GetDeviceRevision() int {
	m.a.Log.Method("GetDeviceRevision", nil)
	return acceptor.Device.Revision
}

func (m *MethodsObj) GetDeviceSerialNumber() string {
	m.a.Log.Method("GetDeviceSerialNumber", nil)

	err := acceptor.Verify(acceptor.Cap.DeviceSerialNumber, "DeviceSerialNumber")
	if err != nil {
		m.a.Log.Err("GetDeviceSerialNumber", err)
		return ""
	}

	payload := []byte{consts.CmdAuxiliary.Byte(), 0, 0, consts.CmdAuxAcceptorSerialNumber.Byte()}
	reply, err := m.a.SendSynchronousCommand(payload)
	if err != nil {
		m.a.Log.Err("GetDeviceSerialNumber", err)
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

func (m *MethodsObj) GetDeviceStalled() bool {
	m.a.Log.Method("GetDeviceStalled", nil)
	return acceptor.Device.Stalled
}

func (m *MethodsObj) GetDeviceState() enum.StateType {
	m.a.Log.Method("GetDeviceState", nil)
	return acceptor.Device.State
}

func (m *MethodsObj) GetDeviceType() string {
	m.a.Log.Method("GetDeviceType", nil)

	err := acceptor.Verify(acceptor.Cap.DeviceType, "DeviceType")
	if err != nil {
		m.a.Log.Err("GetDeviceType", err)
		return ""
	}

	payload := []byte{consts.CmdAuxiliary.Byte(), 0, 0, consts.CmdAuxAcceptorType.Byte()}

	reply, err := m.a.SendSynchronousCommand(payload)
	if err != nil {
		m.a.Log.Err("GetDeviceType", err)
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
