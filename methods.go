package mpost

import (
	"errors"
	"github.com/hard-soft-ware/mpost/acceptor"
	"github.com/hard-soft-ware/mpost/consts"
	"github.com/hard-soft-ware/mpost/enum"
	"github.com/hard-soft-ware/mpost/hook"
	"time"
)

// //////////////////////////////////

type MethodsObj struct {
	a *MpostObj
}

func (a *MpostObj) newMethods() *MethodsObj {
	return &MethodsObj{a: a}
}

////

func (m *MethodsObj) GetBNFStatus() enum.BNFStatusType {
	m.a.Log.Method("Getting BNF status", nil)
	err := acceptor.Verify(acceptor.Cap.BNFStatus, "BNFStatus")

	if err != nil {
		m.a.Log.Err("GetBNFStatus", err)
		return enum.BNFStatusUnknown
	}

	payload := []byte{consts.CmdAuxiliary.Byte(), 0, 0, consts.CmdAuxBNFStatus.Byte()}

	reply, err := m.a.SendSynchronousCommand(payload)
	if err != nil {
		m.a.Log.Err("GetBNFStatus", err)
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

func (m *MethodsObj) GetCapAssetNumber() bool {
	m.a.Log.Method("GetCapAssetNumber", nil)
	return acceptor.Cap.AssetNumber
}

func (m *MethodsObj) GetCapBNFStatus() bool {
	m.a.Log.Method("GetCapBNFStatus", nil)
	return acceptor.Cap.BNFStatus
}

func (m *MethodsObj) GetCapBookmark() bool {
	m.a.Log.Method("GetCapBookmark", nil)
	return acceptor.Cap.Bookmark
}

func (m *MethodsObj) GetCapCalibrate() bool {
	m.a.Log.Method("GetCapCalibrate", nil)
	return acceptor.Cap.Calibrate
}

func (m *MethodsObj) GetCapEscrowTimeout() bool {
	m.a.Log.Method("GetCapEscrowTimeout", nil)
	return acceptor.Cap.EscrowTimeout
}

func (m *MethodsObj) GetCapFlashDownload() bool {
	m.a.Log.Method("GetCapFlashDownload", nil)
	return acceptor.Cap.FlashDownload
}

func (m *MethodsObj) GetCapNoPush() bool {
	m.a.Log.Method("GetCapNoPush", nil)
	return acceptor.Cap.NoPush
}

func (m *MethodsObj) GetCapPupExt() bool {
	m.a.Log.Method("GetCapPupExt", nil)
	return acceptor.Cap.PupExt
}

func (m *MethodsObj) GetCapTestDoc() bool {
	m.a.Log.Method("GetCapTestDoc", nil)
	return acceptor.Cap.TestDoc
}

func (m *MethodsObj) GetConnected() bool {
	m.a.Log.Method("GetConnected", nil)
	return acceptor.Connected
}

func (m *MethodsObj) GetDocType() enum.DocumentType {
	m.a.Log.Method("GetDocType", nil)
	return m.a.DocType
}

func (m *MethodsObj) GetTransactionTimeout() time.Duration {
	m.a.Log.Method("GetTransactionTimeout", nil)
	return acceptor.Timeout.Transaction
}

func (m *MethodsObj) SetTransactionTimeout(v time.Duration) {
	m.a.Log.Method("SetTransactionTimeout", nil)
	acceptor.Timeout.Transaction = v
}

func (m *MethodsObj) GetDownloadTimeout() time.Duration {
	m.a.Log.Method("GetDownloadTimeout", nil)
	return acceptor.Timeout.Download
}

func (m *MethodsObj) SetDownloadTimeout(v time.Duration) {
	m.a.Log.Method("SetDownloadTimeout", nil)
	acceptor.Timeout.Download = v
}

func (m *MethodsObj) GetEnableAcceptance() bool {
	m.a.Log.Method("GetEnableAcceptance", nil)
	return acceptor.Enable.Acceptance
}

func (m *MethodsObj) SetEnableAcceptance(v bool) {
	m.a.Log.Method("SetEnableAcceptance", nil)
	acceptor.Enable.Acceptance = v
}

func (m *MethodsObj) GetEnableBookmarks() bool {
	m.a.Log.Method("GetEnableBookmarks", nil)
	return acceptor.Enable.Bookmarks
}

func (m *MethodsObj) SetEnableBookmarks(v bool) {
	m.a.Log.Method("SetEnableBookmarks", nil)
	acceptor.Enable.Bookmarks = v
}

func (m *MethodsObj) GetEnableNoPush() bool {
	m.a.Log.Method("GetEnableNoPush", nil)
	return acceptor.Enable.NoPush
}

func (m *MethodsObj) SetEnableNoPush(v bool) {
	m.a.Log.Method("SetEnableNoPush", nil)
	acceptor.Enable.NoPush = v
}

func (m *MethodsObj) GetHighSecurity() bool {
	m.a.Log.Method("GetHighSecurity", nil)
	return acceptor.HighSecurity
}

func (m *MethodsObj) SetHighSecurity(v bool) {
	m.a.Log.Method("SetHighSecurity", nil)
	acceptor.HighSecurity = v
}

func (m *MethodsObj) GetVersion() string {
	m.a.Log.Method("GetVersion", nil)
	return acceptor.Version
}

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

func (m *MethodsObj) GetAutoStack() bool {
	m.a.Log.Method("GetAutoStack", nil)
	return acceptor.AutoStack
}

func (m *MethodsObj) SetAutoStack(v bool) {
	m.a.Log.Method("SetAutoStack", nil)
	acceptor.AutoStack = v
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

func (m *MethodsObj) GetCapBootPN() bool {
	m.a.Log.Method("GetCapBootPN", nil)
	return acceptor.Cap.BootPN
}
