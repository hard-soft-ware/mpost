package bill

import (
	"github.com/hard-soft-ware/mpost/acceptor"
)

////////////////////////////////////

func ClearTable() {
	Types = []BillStruct{}
	TypeEnables = []bool{}
	Values = []BillStruct{}
	ValueEnables = []bool{}
}

func BuildHardCodedTable() {
	Types = []BillStruct{}

	switch acceptor.Device.Model {
	case 1, 12, 23, 30, 31, 'J', 'X', 'T':
		Types = append(Types, BillStruct{"USD", 1, '*', '*', '*', '*'})
		Types = append(Types, BillStruct{"USD", 2, '*', '*', '*', '*'})
		Types = append(Types, BillStruct{"USD", 5, '*', '*', '*', '*'})
		Types = append(Types, BillStruct{"USD", 10, '*', '*', '*', '*'})
		Types = append(Types, BillStruct{"USD", 20, '*', '*', '*', '*'})
		Types = append(Types, BillStruct{"USD", 50, '*', '*', '*', '*'})
		Types = append(Types, BillStruct{"USD", 100, '*', '*', '*', '*'})

	case 'P':
		Types = append(Types, BillStruct{"USD", 1, '*', '*', '*', '*'})
		Types = append(Types, BillStruct{"USD", 2, '*', '*', '*', '*'})
		Types = append(Types, BillStruct{"USD", 5, '*', '*', '*', '*'})
		Types = append(Types, BillStruct{"USD", 10, '*', '*', '*', '*'})
		Types = append(Types, BillStruct{"USD", 20, '*', '*', '*', '*'})

	case 'G':
		Types = append(Types, BillStruct{})
		Types = append(Types, BillStruct{"ARS", 2, '*', '*', '*', '*'})
		Types = append(Types, BillStruct{"ARS", 5, '*', '*', '*', '*'})
		Types = append(Types, BillStruct{"ARS", 10, '*', '*', '*', '*'})
		Types = append(Types, BillStruct{"ARS", 20, '*', '*', '*', '*'})
		Types = append(Types, BillStruct{"ARS", 50, '*', '*', '*', '*'})
		Types = append(Types, BillStruct{"ARS", 100, '*', '*', '*', '*'})

	case 'A':
		Types = append(Types, BillStruct{"AUD", 5, '*', '*', '*', '*'})
		Types = append(Types, BillStruct{"AUD", 10, '*', '*', '*', '*'})
		Types = append(Types, BillStruct{"AUD", 100, '*', '*', '*', '*'})

	case 15:
		Types = append(Types, BillStruct{})
		Types = append(Types, BillStruct{})
		Types = append(Types, BillStruct{"AUD", 5, '*', '*', '*', '*'})
		Types = append(Types, BillStruct{"AUD", 10, '*', '*', '*', '*'})
		Types = append(Types, BillStruct{"AUD", 20, '*', '*', '*', '*'})
		Types = append(Types, BillStruct{"AUD", 50, '*', '*', '*', '*'})
		Types = append(Types, BillStruct{"AUD", 100, '*', '*', '*', '*'})

	case 'W':
		Types = append(Types, BillStruct{"BRL", 1, '*', '*', '*', '*'})
		Types = append(Types, BillStruct{"BRL", 2, '*', '*', '*', '*'})
		Types = append(Types, BillStruct{"BRL", 5, '*', '*', '*', '*'})
		Types = append(Types, BillStruct{"BRL", 10, '*', '*', '*', '*'})
		Types = append(Types, BillStruct{"BRL", 20, '*', '*', '*', '*'})
		Types = append(Types, BillStruct{"BRL", 50, '*', '*', '*', '*'})
		Types = append(Types, BillStruct{"BRL", 100, '*', '*', '*', '*'})

	case 'C':
		Types = append(Types, BillStruct{})
		Types = append(Types, BillStruct{})
		Types = append(Types, BillStruct{"CAD", 5, '*', '*', '*', '*'})
		Types = append(Types, BillStruct{"CAD", 10, '*', '*', '*', '*'})
		Types = append(Types, BillStruct{"CAD", 20, '*', '*', '*', '*'})
		Types = append(Types, BillStruct{"CAD", 50, '*', '*', '*', '*'})
		Types = append(Types, BillStruct{"CAD", 100, '*', '*', '*', '*'})

	case 'D':
		Types = append(Types, BillStruct{"EUR", 5, '*', '*', '*', '*'})
		Types = append(Types, BillStruct{"EUR", 10, '*', '*', '*', '*'})

	case 'M':
		Types = append(Types, BillStruct{"MXP", 20, '*', '*', '*', '*'})
		Types = append(Types, BillStruct{"MXP", 50, '*', '*', '*', '*'})

	case 'B':
		Types = append(Types, BillStruct{"RUR", 10, '*', '*', '*', '*'})
		Types = append(Types, BillStruct{"RUR", 50, '*', '*', '*', '*'})
		Types = append(Types, BillStruct{"RUR", 100, '*', '*', '*', '*'})
		Types = append(Types, BillStruct{"RUR", 500, '*', '*', '*', '*'})

	default:
		Types = append(Types, BillStruct{"USD", 1, '*', '*', '*', '*'})
		Types = append(Types, BillStruct{"USD", 2, '*', '*', '*', '*'})
		Types = append(Types, BillStruct{"USD", 5, '*', '*', '*', '*'})
		Types = append(Types, BillStruct{"USD", 10, '*', '*', '*', '*'})
		Types = append(Types, BillStruct{"USD", 20, '*', '*', '*', '*'})
		Types = append(Types, BillStruct{"USD", 50, '*', '*', '*', '*'})
		Types = append(Types, BillStruct{"USD", 100, '*', '*', '*', '*'})
	}

	TypeEnables = make([]bool, len(Types))
	for i, bill := range Types {
		if bill.Value > 0 {
			TypeEnables[i] = true
		} else {
			TypeEnables[i] = false
		}
	}
}

func BuildValues() {
	Values = []BillStruct{}
	ValueEnables = []bool{}

	for i := range Types {
		valueExists := false

		for j := range Values {
			if Types[i].Value == Values[j].Value && Types[i].Country == Values[j].Country {
				valueExists = true
				break
			}
		}

		if !valueExists {
			Values = append(Values, BillStruct{
				Country:       Types[i].Country,
				Value:         Types[i].Value,
				Type:          '*',
				Series:        '*',
				Compatibility: '*',
				Version:       '*',
			})
			ValueEnables = append(ValueEnables, Types[i].Value > 0)
		}
	}
}

////

func SetUpTable(expandedNoteReporting bool, RetrieveBillTable func()) {
	ClearTable()

	if expandedNoteReporting {
		RetrieveBillTable()
	} else {
		BuildHardCodedTable()
	}

	BuildValues()
}
