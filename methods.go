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

func (a *MpostObj) GetBNFStatus() enum.BNFStatusType {
	a.Log.Method("Getting BNF status", nil)
	err := acceptor.Verify(acceptor.Cap.BNFStatus, "BNFStatus")

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

func (a *MpostObj) GetCapAssetNumber() bool {
	a.Log.Method("GetCapAssetNumber", nil)
	return acceptor.Cap.AssetNumber
}

func (a *MpostObj) GetCapBNFStatus() bool {
	a.Log.Method("GetCapBNFStatus", nil)
	return acceptor.Cap.BNFStatus
}

func (a *MpostObj) GetCapBookmark() bool {
	a.Log.Method("GetCapBookmark", nil)
	return acceptor.Cap.Bookmark
}

func (a *MpostObj) GetCapCalibrate() bool {
	a.Log.Method("GetCapCalibrate", nil)
	return acceptor.Cap.Calibrate
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

func (a *MpostObj) GetCapPupExt() bool {
	a.Log.Method("GetCapPupExt", nil)
	return acceptor.Cap.PupExt
}

func (a *MpostObj) GetCapTestDoc() bool {
	a.Log.Method("GetCapTestDoc", nil)
	return acceptor.Cap.TestDoc
}

func (a *MpostObj) GetConnected() bool {
	a.Log.Method("GetConnected", nil)
	return acceptor.Connected
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

func (a *MpostObj) GetEnableBookmarks() bool {
	a.Log.Method("GetEnableBookmarks", nil)
	return acceptor.Enable.Bookmarks
}

func (a *MpostObj) SetEnableBookmarks(v bool) {
	a.Log.Method("SetEnableBookmarks", nil)
	acceptor.Enable.Bookmarks = v
}

func (a *MpostObj) GetEnableNoPush() bool {
	a.Log.Method("GetEnableNoPush", nil)
	return acceptor.Enable.NoPush
}

func (a *MpostObj) SetEnableNoPush(v bool) {
	a.Log.Method("SetEnableNoPush", nil)
	acceptor.Enable.NoPush = v
}

func (a *MpostObj) GetHighSecurity() bool {
	a.Log.Method("GetHighSecurity", nil)
	return acceptor.HighSecurity
}

func (a *MpostObj) SetHighSecurity(v bool) {
	a.Log.Method("SetHighSecurity", nil)
	acceptor.HighSecurity = v
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

func (a *MpostObj) GetAutoStack() bool {
	a.Log.Method("GetAutoStack", nil)
	return acceptor.AutoStack
}

func (a *MpostObj) SetAutoStack(v bool) {
	a.Log.Method("SetAutoStack", nil)
	acceptor.AutoStack = v
}

func (a *MpostObj) GetBootPN() string {
	a.Log.Method("GetBootPN", nil)

	err := acceptor.Verify(acceptor.Cap.BootPN, "GetBootPN")
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

func (a *MpostObj) GetCapBootPN() bool {
	a.Log.Method("GetCapBootPN", nil)
	return acceptor.Cap.BootPN
}
