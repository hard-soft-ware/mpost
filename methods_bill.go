package mpost

import (
	"errors"
	"fmt"
	"github.com/hard-soft-ware/mpost/acceptor"
	"github.com/hard-soft-ware/mpost/bill"
	"github.com/hard-soft-ware/mpost/consts"
)

////////////////////////////////////

type MethodsBillObj struct {
	a   *MpostObj
	Get MethodsBillGetObj
	Set MethodsBillSetObj
}

type MethodsBillGetObj struct{ a *MpostObj }
type MethodsBillSetObj struct{ a *MpostObj }

func (m *MethodsObj) newBill() *MethodsBillObj {
	obj := MethodsBillObj{}

	obj.a = m.a
	obj.Get.a = m.a
	obj.Set.a = m.a

	return &obj
}

////////////////

func (m *MethodsBillGetObj) Self() bill.BillStruct {
	m.a.Log.Method("GetBill", nil)
	return bill.Bill
}

func (m *MethodsBillGetObj) Types() []bill.BillStruct {
	m.a.Log.Method("GetBillTypes", nil)
	return bill.Types
}

func (m *MethodsBillGetObj) Values() []bill.BillStruct {
	m.a.Log.Method("GetBillValues", nil)
	return bill.Values
}

func (m *MethodsBillGetObj) EnablesType() []bool {
	m.a.Log.Method("GetBillTypeEnables", nil)
	return bill.TypeEnables
}

func (m *MethodsBillGetObj) EnablesValue() []bool {
	m.a.Log.Method("GetBillValueEnables", nil)
	return bill.ValueEnables
}

//

func (m *MethodsBillSetObj) EnablesType(v []bool) {
	m.a.Log.Method("SetBillTypeEnables", v)

	if !acceptor.Connected {
		m.a.Log.Err("SetBillTypeEnables", errors.New("calling BillTypeEnables not allowed when not connected"))
		return
	}

	if len(bill.TypeEnables) != len(bill.Types) {
		m.a.Log.Err("SetBillTypeEnables", fmt.Errorf("CBillTypeEnables size must match BillTypes size"))
		return
	}

	bill.TypeEnables = v

	if acceptor.ExpandedNoteReporting {
		payload := make([]byte, 15)
		acceptor.ConstructOmnibusCommand(payload, consts.CmdExpanded, 2, bill.TypeEnables)
		payload[1] = 0x03 // Sub Type

		for i, enable := range bill.TypeEnables {
			enableIndex := i / 7
			bitPosition := i % 7
			bit := 1 << bitPosition

			if enable {
				payload[5+enableIndex] |= byte(bit)
			}
		}

		m.a.SendAsynchronousCommand(payload)
	}
}

func (m *MethodsBillSetObj) EnablesValue(v []bool) {
	m.a.Log.Method("SetBillValueEnables", v)
	bill.ValueEnables = v

	for _, enabled := range bill.ValueEnables {
		for j, billType := range bill.Types {
			if billType.Value == bill.Values[j].Value && billType.Country == bill.Values[j].Country {
				bill.TypeEnables[j] = enabled
			}
		}
	}

	payload := make([]byte, 15)
	acceptor.ConstructOmnibusCommand(payload, consts.CmdExpanded, 2, bill.TypeEnables)
	payload[1] = 0x03 // Sub Type

	for i, enable := range bill.TypeEnables {
		enableIndex := i / 7
		bitPosition := i % 7
		bit := 1 << bitPosition

		if enable {
			payload[5+enableIndex] |= byte(bit)
		}
	}

	m.a.SendAsynchronousCommand(payload)
}
