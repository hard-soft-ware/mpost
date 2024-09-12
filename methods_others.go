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

type MethodsOtherObj struct {
	a *MpostObj
}

func (m *MethodsObj) newOther() *MethodsOtherObj {
	obj := MethodsOtherObj{}
	obj.a = m.a
	return &obj
}

////////////////

func (m *MethodsOtherObj) Calibrate() {
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

func (m *MethodsOtherObj) SoftReset() {
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

func (m *MethodsOtherObj) RawTransaction(command []byte) ([]byte, error) {
	m.a.Log.Method("RawTransaction", command)

	reply, err := m.a.SendSynchronousCommand(command)
	if err != nil {
		m.a.Log.Err("RawTransaction", err)
		return reply, err
	}

	return reply, nil
}

////

func (m *MethodsOtherObj) GetConnected() bool {
	m.a.Log.Method("GetConnected", nil)
	return acceptor.Connected
}

func (m *MethodsOtherObj) GetDocType() enum.DocumentType {
	m.a.Log.Method("GetDocType", nil)
	return m.a.DocType
}

func (m *MethodsOtherObj) GetVersion() string {
	m.a.Log.Method("GetVersion", nil)
	return acceptor.Version
}

func (m *MethodsOtherObj) GetAutoStack() bool {
	m.a.Log.Method("GetAutoStack", nil)
	return acceptor.AutoStack
}

func (m *MethodsOtherObj) GetHighSecurity() bool {
	m.a.Log.Method("GetHighSecurity", nil)
	return acceptor.HighSecurity
}

func (m *MethodsOtherObj) GetBootPN() string {
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

func (m *MethodsOtherObj) GetCapAssetNumber() bool {
	m.a.Log.Method("GetCapAssetNumber", nil)
	return acceptor.Cap.AssetNumber
}

func (m *MethodsOtherObj) GetCapEscrowTimeout() bool {
	m.a.Log.Method("GetCapEscrowTimeout", nil)
	return acceptor.Cap.EscrowTimeout
}

func (m *MethodsOtherObj) GetCapFlashDownload() bool {
	m.a.Log.Method("GetCapFlashDownload", nil)
	return acceptor.Cap.FlashDownload
}

func (m *MethodsOtherObj) GetCapPupExt() bool {
	m.a.Log.Method("GetCapPupExt", nil)
	return acceptor.Cap.PupExt
}

func (m *MethodsOtherObj) GetCapTestDoc() bool {
	m.a.Log.Method("GetCapTestDoc", nil)
	return acceptor.Cap.TestDoc
}

func (m *MethodsOtherObj) GetCapCalibrate() bool {
	m.a.Log.Method("GetCapCalibrate", nil)
	return acceptor.Cap.Calibrate
}

func (m *MethodsOtherObj) GetCapBookmark() bool {
	m.a.Log.Method("GetCapBookmark", nil)
	return acceptor.Cap.Bookmark
}

func (m *MethodsOtherObj) GetCapNoPush() bool {
	m.a.Log.Method("GetCapNoPush", nil)
	return acceptor.Cap.NoPush
}

func (m *MethodsOtherObj) GetCapBootPN() bool {
	m.a.Log.Method("GetCapBootPN", nil)
	return acceptor.Cap.BootPN
}

////

func (m *MethodsOtherObj) SetAutoStack(v bool) {
	m.a.Log.Method("SetAutoStack", v)
	acceptor.AutoStack = v
}

func (m *MethodsOtherObj) SetHighSecurity(v bool) {
	m.a.Log.Method("SetHighSecurity", v)
	acceptor.HighSecurity = v
}

func (m *MethodsOtherObj) SetBezel(bezel enum.BezelType) {
	m.a.Log.Method("SetBezel", bezel)

	if !acceptor.Connected {
		m.a.Log.Err("SetBezel", errors.New("SetBezel called when not connected"))
		return
	}

	payload := []byte{consts.CmdAuxiliary.Byte(), byte(bezel), 0x00, consts.CmdAuxSetBezel.Byte()}
	m.a.SendAsynchronousCommand(payload)
}

func (m *MethodsOtherObj) SetAssetNumber(asset string) {
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
