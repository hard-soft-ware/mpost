package mpost

import (
	"github.com/hard-soft-ware/mpost/acceptor"
	"github.com/hard-soft-ware/mpost/consts"
)

////////////////////////////////////

func (a *MpostObj) GetCapVariantID() bool {
	a.Log.Method("GetCapVariantID", nil)
	return acceptor.Cap.VariantID
}

func (a *MpostObj) GetCapVariantPN() bool {
	a.Log.Method("GetCapVariantPN", nil)
	return acceptor.Cap.VariantPN
}

func (a *MpostObj) GetVariantNames() []string {
	a.Log.Method("GetVariantNames", nil)

	err := acceptor.Verify(true, "VariantNames")
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

	err := acceptor.Verify(acceptor.Cap.VariantID, "VariantID")
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

	err := acceptor.Verify(acceptor.Cap.VariantPN, "VariantPN")
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
