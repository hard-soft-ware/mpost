package mpost

import (
	"errors"
	"github.com/hard-soft-ware/mpost/acceptor"
	"github.com/hard-soft-ware/mpost/bill"
	"github.com/hard-soft-ware/mpost/consts"
	"github.com/hard-soft-ware/mpost/enum"
	"github.com/hard-soft-ware/mpost/hook"
	"time"
)

////////////////////////////////////

func (m *MethodsObj) Calibrate() {
	m.a.Log.Method("Calibrate", nil)
	if !acceptor.Connected {
		m.a.Log.Err("Calibrate", errors.New("Calibrate called when not connected"))
		return
	}

	if acceptor.Device.State != enum.StateIdling {
		m.a.Log.Err("Calibrate", errors.New("Calibrate allowed only when DeviceState == Idling"))
		return
	}

	payload := []byte{consts.CmdCalibrate.Byte(), 0x00, 0x00, 0x00}

	acceptor.SuppressStandardPoll = true
	acceptor.Device.State = enum.StateCalibrateStart

	hook.Raise.Calibrate.Start()

	hook.CalibrateProgress = true

	startTickCount := time.Now()

	for {
		reply, err := m.a.SendSynchronousCommand(payload)
		if err != nil {
			m.a.Log.Err("Calibrate", errors.New("Failed to send synchronous command during calibration"))
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

func (m *MethodsObj) SoftReset() {
	m.a.Log.Method("SoftReset", nil)

	err := acceptor.Verify(acceptor.Cap.DeviceSoftReset, "SoftReset")
	if err != nil {
		m.a.Log.Err("SoftReset", err)
		return
	}

	m.a.DocType = enum.DocumentNoValue

	payload := []byte{consts.CmdAuxiliary.Byte(), 0x7F, 0x7F, 0x7F}
	m.a.SendAsynchronousCommand(payload)
	acceptor.InSoftResetWaitForReply = true
	acceptor.InSoftResetOneSecondIgnore = true
}

func (m *MethodsObj) RawTransaction(command []byte) ([]byte, error) {
	m.a.Log.Method("RawTransaction", command)

	reply, err := m.a.SendSynchronousCommand(command)
	if err != nil {
		m.a.Log.Err("RawTransaction", err)
		return reply, err
	}

	return reply, nil
}

////

func (m *MethodsObj) GetConnected() bool {
	m.a.Log.Method("GetConnected", nil)
	return acceptor.Connected
}

func (m *MethodsObj) GetDocType() enum.DocumentType {
	m.a.Log.Method("GetDocType", nil)
	return m.a.DocType
}

func (m *MethodsObj) GetVersion() string {
	m.a.Log.Method("GetVersion", nil)
	return acceptor.Version
}

func (m *MethodsObj) GetAutoStack() bool {
	m.a.Log.Method("GetAutoStack", nil)
	return acceptor.AutoStack
}

func (m *MethodsObj) GetHighSecurity() bool {
	m.a.Log.Method("GetHighSecurity", nil)
	return acceptor.HighSecurity
}

func (m *MethodsObj) GetBootPN() string {
	m.a.Log.Method("GetBootPN", nil)

	err := acceptor.Verify(acceptor.Cap.BootPN, "GetBootPN")
	if err != nil {
		m.a.Log.Err("GetBootPN", err)
		return ""
	}

	payload := []byte{consts.CmdAuxiliary.Byte(), 0, 0, consts.CmdAuxAcceptorBootPartNumber.Byte()}

	reply, err := m.a.SendSynchronousCommand(payload)
	if err != nil {
		m.a.Log.Err("GetBootPN", err)
		return ""
	}

	if len(reply) == 14 {
		s := string(reply[3:12]) // Extracting the substring from byte slice
		return s
	}

	return ""
}

//

func (m *MethodsObj) GetCapAssetNumber() bool {
	m.a.Log.Method("GetCapAssetNumber", nil)
	return acceptor.Cap.AssetNumber
}

func (m *MethodsObj) GetCapEscrowTimeout() bool {
	m.a.Log.Method("GetCapEscrowTimeout", nil)
	return acceptor.Cap.EscrowTimeout
}

func (m *MethodsObj) GetCapFlashDownload() bool {
	m.a.Log.Method("GetCapFlashDownload", nil)
	return acceptor.Cap.FlashDownload
}

func (m *MethodsObj) GetCapPupExt() bool {
	m.a.Log.Method("GetCapPupExt", nil)
	return acceptor.Cap.PupExt
}

func (m *MethodsObj) GetCapTestDoc() bool {
	m.a.Log.Method("GetCapTestDoc", nil)
	return acceptor.Cap.TestDoc
}

func (m *MethodsObj) GetCapCalibrate() bool {
	m.a.Log.Method("GetCapCalibrate", nil)
	return acceptor.Cap.Calibrate
}

func (m *MethodsObj) GetCapBookmark() bool {
	m.a.Log.Method("GetCapBookmark", nil)
	return acceptor.Cap.Bookmark
}

func (m *MethodsObj) GetCapNoPush() bool {
	m.a.Log.Method("GetCapNoPush", nil)
	return acceptor.Cap.NoPush
}

func (m *MethodsObj) GetCapBootPN() bool {
	m.a.Log.Method("GetCapBootPN", nil)
	return acceptor.Cap.BootPN
}

////

func (m *MethodsObj) SetAutoStack(v bool) {
	m.a.Log.Method("SetAutoStack", v)
	acceptor.AutoStack = v
}

func (m *MethodsObj) SetHighSecurity(v bool) {
	m.a.Log.Method("SetHighSecurity", v)
	acceptor.HighSecurity = v
}

func (m *MethodsObj) SetBezel(bezel enum.BezelType) {
	m.a.Log.Method("SetBezel", bezel)

	if !acceptor.Connected {
		m.a.Log.Err("SetBezel", errors.New("SetBezel called when not connected"))
		return
	}

	payload := []byte{consts.CmdAuxiliary.Byte(), byte(bezel), 0x00, consts.CmdAuxSetBezel.Byte()}
	m.a.SendAsynchronousCommand(payload)
}

func (m *MethodsObj) SetAssetNumber(asset string) {
	m.a.Log.Method("SetAssetNumber", asset)

	if !acceptor.Connected {
		m.a.Log.Err("SetAssetNumber", errors.New("SetAssetNumber called when not connected"))
		return
	}

	payload := make([]byte, 21)
	acceptor.ConstructOmnibusCommand(payload, consts.CmdExpanded, 2, bill.TypeEnables)
	payload[1] = 0x05 // Setting the sub command or option byte

	maxLen := len(asset)
	if maxLen > 16 {
		maxLen = 16
	}
	copy(payload[5:], asset[:maxLen])

	m.a.SendAsynchronousCommand(payload)
}
