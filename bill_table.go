package mpost

import (
	"github.com/hard-soft-ware/mpost/consts"
	"time"
)

////////////////////////////////////

func (a *CAcceptor) SetUpBillTable() {
	a.ClearBillTable()

	if a.expandedNoteReporting {
		a.RetrieveBillTable()
	} else {
		a.BuildHardCodedBillTable()
	}

	a.BuildBillValues()
}

////

func (a *CAcceptor) ClearBillTable() {
	a.billTypes = []CBill{}
	a.billTypeEnables = []bool{}
	a.billValues = []CBill{}
	a.billValueEnables = []bool{}

	a.log.Debug().Msg("Bill table cleared")
}

func (a *CAcceptor) RetrieveBillTable() {
	index := 1
	for {
		payload := make([]byte, 6)
		a.ConstructOmnibusCommand(payload, consts.CmdExpanded.Byte(), 2)
		payload[1] = 0x02
		payload[5] = byte(index)

		var reply []byte
		var err error
		for {
			reply, err = a.SendSynchronousCommand(payload)
			if err != nil || len(reply) == 30 {
				break
			}
			time.Sleep(100 * time.Millisecond)
		}

		if err != nil || len(reply) != 30 {
			a.log.Debug().Err(err).Msg("Error sending command")
			break
		}

		ctl := reply[2]
		if (ctl&0x70) != 0x70 || reply[3] != 0x02 {
			break
		}

		if reply[10] == 0 {
			break
		}

		billFromTable := a.ParseBillData(reply, 10)
		a.billTypes = append(a.billTypes, billFromTable)
		index++
	}

	for range a.billTypes {
		a.billTypeEnables = append(a.billTypeEnables, true)
	}

	a.log.Debug().Msg("Bill table retrieved")
}

func (a *CAcceptor) BuildHardCodedBillTable() {
	a.billTypes = []CBill{}

	switch a.deviceModel {
	case 1, 12, 23, 30, 31, 'J', 'X', 'T':
		a.billTypes = append(a.billTypes, CBill{"USD", 1, '*', '*', '*', '*'})
		a.billTypes = append(a.billTypes, CBill{"USD", 2, '*', '*', '*', '*'})
		a.billTypes = append(a.billTypes, CBill{"USD", 5, '*', '*', '*', '*'})
		a.billTypes = append(a.billTypes, CBill{"USD", 10, '*', '*', '*', '*'})
		a.billTypes = append(a.billTypes, CBill{"USD", 20, '*', '*', '*', '*'})
		a.billTypes = append(a.billTypes, CBill{"USD", 50, '*', '*', '*', '*'})
		a.billTypes = append(a.billTypes, CBill{"USD", 100, '*', '*', '*', '*'})

	case 'P':
		a.billTypes = append(a.billTypes, CBill{"USD", 1, '*', '*', '*', '*'})
		a.billTypes = append(a.billTypes, CBill{"USD", 2, '*', '*', '*', '*'})
		a.billTypes = append(a.billTypes, CBill{"USD", 5, '*', '*', '*', '*'})
		a.billTypes = append(a.billTypes, CBill{"USD", 10, '*', '*', '*', '*'})
		a.billTypes = append(a.billTypes, CBill{"USD", 20, '*', '*', '*', '*'})

	case 'G':
		a.billTypes = append(a.billTypes, CBill{})
		a.billTypes = append(a.billTypes, CBill{"ARS", 2, '*', '*', '*', '*'})
		a.billTypes = append(a.billTypes, CBill{"ARS", 5, '*', '*', '*', '*'})
		a.billTypes = append(a.billTypes, CBill{"ARS", 10, '*', '*', '*', '*'})
		a.billTypes = append(a.billTypes, CBill{"ARS", 20, '*', '*', '*', '*'})
		a.billTypes = append(a.billTypes, CBill{"ARS", 50, '*', '*', '*', '*'})
		a.billTypes = append(a.billTypes, CBill{"ARS", 100, '*', '*', '*', '*'})

	case 'A':
		a.billTypes = append(a.billTypes, CBill{"AUD", 5, '*', '*', '*', '*'})
		a.billTypes = append(a.billTypes, CBill{"AUD", 10, '*', '*', '*', '*'})
		a.billTypes = append(a.billTypes, CBill{"AUD", 100, '*', '*', '*', '*'})

	case 15:
		a.billTypes = append(a.billTypes, CBill{})
		a.billTypes = append(a.billTypes, CBill{})
		a.billTypes = append(a.billTypes, CBill{"AUD", 5, '*', '*', '*', '*'})
		a.billTypes = append(a.billTypes, CBill{"AUD", 10, '*', '*', '*', '*'})
		a.billTypes = append(a.billTypes, CBill{"AUD", 20, '*', '*', '*', '*'})
		a.billTypes = append(a.billTypes, CBill{"AUD", 50, '*', '*', '*', '*'})
		a.billTypes = append(a.billTypes, CBill{"AUD", 100, '*', '*', '*', '*'})

	case 'W':
		a.billTypes = append(a.billTypes, CBill{"BRL", 1, '*', '*', '*', '*'})
		a.billTypes = append(a.billTypes, CBill{"BRL", 2, '*', '*', '*', '*'})
		a.billTypes = append(a.billTypes, CBill{"BRL", 5, '*', '*', '*', '*'})
		a.billTypes = append(a.billTypes, CBill{"BRL", 10, '*', '*', '*', '*'})
		a.billTypes = append(a.billTypes, CBill{"BRL", 20, '*', '*', '*', '*'})
		a.billTypes = append(a.billTypes, CBill{"BRL", 50, '*', '*', '*', '*'})
		a.billTypes = append(a.billTypes, CBill{"BRL", 100, '*', '*', '*', '*'})

	case 'C':
		a.billTypes = append(a.billTypes, CBill{})
		a.billTypes = append(a.billTypes, CBill{})
		a.billTypes = append(a.billTypes, CBill{"CAD", 5, '*', '*', '*', '*'})
		a.billTypes = append(a.billTypes, CBill{"CAD", 10, '*', '*', '*', '*'})
		a.billTypes = append(a.billTypes, CBill{"CAD", 20, '*', '*', '*', '*'})
		a.billTypes = append(a.billTypes, CBill{"CAD", 50, '*', '*', '*', '*'})
		a.billTypes = append(a.billTypes, CBill{"CAD", 100, '*', '*', '*', '*'})

	case 'D':
		a.billTypes = append(a.billTypes, CBill{"EUR", 5, '*', '*', '*', '*'})
		a.billTypes = append(a.billTypes, CBill{"EUR", 10, '*', '*', '*', '*'})

	case 'M':
		a.billTypes = append(a.billTypes, CBill{"MXP", 20, '*', '*', '*', '*'})
		a.billTypes = append(a.billTypes, CBill{"MXP", 50, '*', '*', '*', '*'})

	case 'B':
		a.billTypes = append(a.billTypes, CBill{"RUR", 10, '*', '*', '*', '*'})
		a.billTypes = append(a.billTypes, CBill{"RUR", 50, '*', '*', '*', '*'})
		a.billTypes = append(a.billTypes, CBill{"RUR", 100, '*', '*', '*', '*'})
		a.billTypes = append(a.billTypes, CBill{"RUR", 500, '*', '*', '*', '*'})

	default:
		a.billTypes = append(a.billTypes, CBill{"USD", 1, '*', '*', '*', '*'})
		a.billTypes = append(a.billTypes, CBill{"USD", 2, '*', '*', '*', '*'})
		a.billTypes = append(a.billTypes, CBill{"USD", 5, '*', '*', '*', '*'})
		a.billTypes = append(a.billTypes, CBill{"USD", 10, '*', '*', '*', '*'})
		a.billTypes = append(a.billTypes, CBill{"USD", 20, '*', '*', '*', '*'})
		a.billTypes = append(a.billTypes, CBill{"USD", 50, '*', '*', '*', '*'})
		a.billTypes = append(a.billTypes, CBill{"USD", 100, '*', '*', '*', '*'})
	}

	a.billTypeEnables = make([]bool, len(a.billTypes))
	for i, bill := range a.billTypes {
		if bill.Value > 0 {
			a.billTypeEnables[i] = true
		} else {
			a.billTypeEnables[i] = false
		}
	}

	a.log.Debug().Msg("Hardcoded bill table built")
}

func (a *CAcceptor) BuildBillValues() {
	a.billValues = []CBill{}
	a.billValueEnables = []bool{}

	for i := range a.billTypes {
		valueExists := false

		for j := range a.billValues {
			if a.billTypes[i].Value == a.billValues[j].Value && a.billTypes[i].Country == a.billValues[j].Country {
				valueExists = true
				break
			}
		}

		if !valueExists {
			a.billValues = append(a.billValues, CBill{
				Country:       a.billTypes[i].Country,
				Value:         a.billTypes[i].Value,
				Type:          '*',
				Series:        '*',
				Compatibility: '*',
				Version:       '*',
			})
			a.billValueEnables = append(a.billValueEnables, a.billTypes[i].Value > 0)
		}
	}

	a.log.Debug().Msg("Bill values built")
}
