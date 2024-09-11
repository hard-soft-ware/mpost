package mpost

import (
	"errors"
	"github.com/hard-soft-ware/mpost/acceptor"
	"github.com/hard-soft-ware/mpost/consts"
	"github.com/hard-soft-ware/mpost/enum"
	"github.com/hard-soft-ware/mpost/hook"
	"time"
)

////////////////////////////////////

type MethodsObj struct {
	a *MpostObj

	Get MethodsGetObj
	Set MethodsSetObj

	Enable      *MethodsEnableObj
	Application *MethodsApplicationObj
	Audit       *MethodsAuditObj
	BarCode     *MethodsBarCodeObj
}

type MethodsGetObj struct{ a *MpostObj }

type MethodsSetObj struct{ a *MpostObj }

////

func (a *MpostObj) newMethods() *MethodsObj {
	obj := MethodsObj{}

	obj.a = a
	obj.Get.a = a
	obj.Set.a = a

	obj.Enable = obj.newEnable()
	obj.Application = obj.newApplication()
	obj.Audit = obj.newAudit()
	obj.BarCode = obj.newBarCode()

	return &obj
}

////////////////

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

////

func (m *MethodsGetObj) Connected() bool {
	m.a.Log.Method("GetConnected", nil)
	return acceptor.Connected
}

func (m *MethodsGetObj) DocType() enum.DocumentType {
	m.a.Log.Method("GetDocType", nil)
	return m.a.DocType
}

func (m *MethodsGetObj) TransactionTimeout() time.Duration {
	m.a.Log.Method("GetTransactionTimeout", nil)
	return acceptor.Timeout.Transaction
}

func (m *MethodsGetObj) DownloadTimeout() time.Duration {
	m.a.Log.Method("GetDownloadTimeout", nil)
	return acceptor.Timeout.Download
}

func (m *MethodsGetObj) Version() string {
	m.a.Log.Method("GetVersion", nil)
	return acceptor.Version
}

func (m *MethodsGetObj) AutoStack() bool {
	m.a.Log.Method("GetAutoStack", nil)
	return acceptor.AutoStack
}

func (m *MethodsGetObj) HighSecurity() bool {
	m.a.Log.Method("GetHighSecurity", nil)
	return acceptor.HighSecurity
}

func (m *MethodsGetObj) BNFStatus() enum.BNFStatusType {
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

func (m *MethodsGetObj) BootPN() string {
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

func (m *MethodsGetObj) CapAssetNumber() bool {
	m.a.Log.Method("GetCapAssetNumber", nil)
	return acceptor.Cap.AssetNumber
}

func (m *MethodsGetObj) CapEscrowTimeout() bool {
	m.a.Log.Method("GetCapEscrowTimeout", nil)
	return acceptor.Cap.EscrowTimeout
}

func (m *MethodsGetObj) CapFlashDownload() bool {
	m.a.Log.Method("GetCapFlashDownload", nil)
	return acceptor.Cap.FlashDownload
}

func (m *MethodsGetObj) CapPupExt() bool {
	m.a.Log.Method("GetCapPupExt", nil)
	return acceptor.Cap.PupExt
}

func (m *MethodsGetObj) CapTestDoc() bool {
	m.a.Log.Method("GetCapTestDoc", nil)
	return acceptor.Cap.TestDoc
}

func (m *MethodsGetObj) CapBNFStatus() bool {
	m.a.Log.Method("GetCapBNFStatus", nil)
	return acceptor.Cap.BNFStatus
}

func (m *MethodsGetObj) CapCalibrate() bool {
	m.a.Log.Method("GetCapCalibrate", nil)
	return acceptor.Cap.Calibrate
}

func (m *MethodsGetObj) CapBookmark() bool {
	m.a.Log.Method("GetCapBookmark", nil)
	return acceptor.Cap.Bookmark
}

func (m *MethodsGetObj) CapNoPush() bool {
	m.a.Log.Method("GetCapNoPush", nil)
	return acceptor.Cap.NoPush
}

func (m *MethodsGetObj) CapBootPN() bool {
	m.a.Log.Method("GetCapBootPN", nil)
	return acceptor.Cap.BootPN
}

////

func (m *MethodsSetObj) TransactionTimeout(v time.Duration) {
	m.a.Log.Method("SetTransactionTimeout", nil)
	acceptor.Timeout.Transaction = v
}

func (m *MethodsSetObj) DownloadTimeout(v time.Duration) {
	m.a.Log.Method("SetDownloadTimeout", nil)
	acceptor.Timeout.Download = v
}

func (m *MethodsSetObj) AutoStack(v bool) {
	m.a.Log.Method("SetAutoStack", nil)
	acceptor.AutoStack = v
}

func (m *MethodsSetObj) HighSecurity(v bool) {
	m.a.Log.Method("SetHighSecurity", nil)
	acceptor.HighSecurity = v
}
