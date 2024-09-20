package mpost

import (
	"github.com/hard-soft-ware/mpost/acceptor"
	"github.com/hard-soft-ware/mpost/consts"
	"github.com/hard-soft-ware/mpost/enum"
)

////////////////////////////////////

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

//

func (m *MethodsObj) GetCapBNFStatus() bool {
	m.a.Log.Method("GetCapBNFStatus", nil)
	return acceptor.Cap.BNFStatus
}
