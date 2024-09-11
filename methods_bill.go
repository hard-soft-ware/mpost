package mpost

import (
	"errors"
	"fmt"
	"github.com/hard-soft-ware/mpost/acceptor"
	"github.com/hard-soft-ware/mpost/bill"
	"github.com/hard-soft-ware/mpost/consts"
)

////////////////////////////////////

func (a *MpostObj) GetBill() bill.BillStruct {
	a.Log.Method("GetBill", nil)
	return bill.Bill
}

func (a *MpostObj) GetBillTypes() []bill.BillStruct {
	a.Log.Method("GetBillTypes", nil)
	return bill.Types
}

func (a *MpostObj) GetBillTypeEnables() []bool {
	a.Log.Method("GetBillTypeEnables", nil)
	return bill.TypeEnables
}

func (a *MpostObj) SetBillTypeEnables(v []bool) {
	a.Log.Method("SetBillTypeEnables", nil)

	if !acceptor.Connected {
		a.Log.Err("SetBillTypeEnables", errors.New("calling BillTypeEnables not allowed when not connected"))
		return
	}

	if len(bill.TypeEnables) != len(bill.Types) {
		a.Log.Err("SetBillTypeEnables", fmt.Errorf("CBillTypeEnables size must match BillTypes size"))
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

		a.SendAsynchronousCommand(payload)
	}
}

func (a *MpostObj) GetBillValues() []bill.BillStruct {
	a.Log.Method("GetBillValues", nil)
	return bill.Values
}

func (a *MpostObj) GetBillValueEnables() []bool {
	a.Log.Method("GetBillValueEnables", nil)
	return bill.ValueEnables
}

func (a *MpostObj) SetBillValueEnables(v []bool) {
	a.Log.Method("SetBillValueEnables", nil)
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

	a.SendAsynchronousCommand(payload)
}
