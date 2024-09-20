package mpost

import (
	"github.com/hard-soft-ware/mpost/acceptor"
	"github.com/hard-soft-ware/mpost/consts"
)

////////////////////////////////////

func (m *MethodsObj) GetAuditLifeTimeTotals() []int {
	m.a.Log.Method("GetAuditLifeTimeTotals", nil)
	values := []int{}

	err := acceptor.Verify(acceptor.Cap.Audit, "GetAuditLifeTimeTotals")
	if err != nil {
		m.a.Log.Err("GetAuditLifeTimeTotals", err)
		return values
	}

	payload := []byte{consts.CmdAuxiliary.Byte(), 0, 0, consts.CmdAuxAcceptorAuditLifeTimeTotals.Byte()}

	reply, err := m.a.SendSynchronousCommand(payload)
	if err != nil {
		m.a.Log.Err("GetAuditLifeTimeTotals", err)
		return values
	}

	if len(reply) < 13 || ((len(reply)-5)%8 != 0) {
		return values
	}

	fieldCount := (len(reply) - 5) / 8
	for i := 0; i < fieldCount; i++ {
		offset := 8*i + 3
		value := (int((reply)[offset+0]&0x0F) << 28) +
			(int((reply)[offset+1]&0x0F) << 24) +
			(int((reply)[offset+2]&0x0F) << 20) +
			(int((reply)[offset+3]&0x0F) << 16) +
			(int((reply)[offset+4]&0x0F) << 12) +
			(int((reply)[offset+5]&0x0F) << 8) +
			(int((reply)[offset+6]&0x0F) << 4) +
			int((reply)[offset+7]&0x0F)

		values = append(values, value)
	}

	return values
}

func (m *MethodsObj) GetAuditPerformance() []int {
	m.a.Log.Method("GetAuditPerformance", nil)
	values := []int{}

	err := acceptor.Verify(acceptor.Cap.Audit, "GetAuditPerformance")
	if err != nil {
		m.a.Log.Err("GetAuditPerformance", err)
		return values
	}

	payload := []byte{consts.CmdAuxiliary.Byte(), 0, 0, consts.CmdAuxAcceptorAuditPerformanceMeasures.Byte()}

	reply, err := m.a.SendSynchronousCommand(payload)
	if err != nil {
		m.a.Log.Err("GetAuditPerformance", err)
		return values
	}

	if len(reply) < 9 || ((len(reply)-5)%4 != 0) {
		return values
	}

	fieldCount := (len(reply) - 5) / 4

	for i := 0; i < fieldCount; i++ {
		offset := 4*i + 3
		value := (int((reply)[offset+0]&0x0F) << 12) +
			(int((reply)[offset+1]&0x0F) << 8) +
			(int((reply)[offset+2]&0x0F) << 4) +
			int((reply)[offset+3]&0x0F)

		values = append(values, value)
	}

	return values
}

func (m *MethodsObj) GetAuditQP() []int {
	m.a.Log.Method("GetAuditQP", nil)
	values := []int{}

	err := acceptor.Verify(acceptor.Cap.Audit, "GetAuditQP")
	if err != nil {
		m.a.Log.Err("GetAuditQP", err)
		return values
	}

	payload := []byte{consts.CmdAuxiliary.Byte(), 0, 0, consts.CmdAuxAcceptorAuditQPMeasures.Byte()}

	reply, err := m.a.SendSynchronousCommand(payload)
	if err != nil {
		m.a.Log.Err("GetAuditQP", err)
		return values
	}

	if len(reply) < 9 || ((len(reply)-5)%4 != 0) {
		return values
	}

	fieldCount := (len(reply) - 5) / 4

	for i := 0; i < fieldCount; i++ {
		offset := 4*i + 3
		value := (int(reply[offset+0]&0x0F) << 12) +
			(int(reply[offset+1]&0x0F) << 8) +
			(int(reply[offset+2]&0x0F) << 4) +
			int(reply[offset+3]&0x0F)

		values = append(values, value)
	}

	return values
}

//

func (m *MethodsObj) GetCapAudit() bool {
	m.a.Log.Method("GetCapAudit", nil)
	return acceptor.Cap.Audit
}
