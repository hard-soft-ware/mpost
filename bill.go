package mpost

import (
	"github.com/hard-soft-ware/mpost/acceptor"
	"github.com/hard-soft-ware/mpost/bill"
	"github.com/hard-soft-ware/mpost/consts"
	"github.com/hard-soft-ware/mpost/enum"
	"strconv"
	"strings"
	"time"
)

////////////////////////////////////

func (a *MpostObj) parseBillData(reply []byte, extDataIndex int) bill.BillStruct {
	var bil bill.BillStruct

	if len(reply) < extDataIndex+15 {
		return bil
	}

	country := string(reply[extDataIndex+1 : extDataIndex+4])
	bil.Country = strings.TrimSpace(country)

	valueString := string(reply[extDataIndex+4 : extDataIndex+7])
	billValue, err := strconv.ParseFloat(valueString, 64)
	if err != nil {
		billValue = 0
	}

	exponentSign := reply[extDataIndex+7]
	exponentString := string(reply[extDataIndex+8 : extDataIndex+10])
	exponent, err := strconv.Atoi(exponentString)
	if err != nil {
		exponent = 0
	}

	if exponentSign == '+' {
		for i := 1; i <= exponent; i++ {
			billValue *= 10.0
		}
	} else {
		for i := 1; i <= exponent; i++ {
			billValue /= 10.0
		}
	}

	bil.Value = billValue
	a.DocType = enum.DocumentBill
	acceptor.WasDocTypeSetOnEscrow = acceptor.Device.State == enum.StateEscrow

	orientation := reply[extDataIndex+10]
	switch orientation {
	case 0x00:
		acceptor.EscrowOrientation = enum.OrientationRightUp
	case 0x01:
		acceptor.EscrowOrientation = enum.OrientationRightDown
	case 0x02:
		acceptor.EscrowOrientation = enum.OrientationLeftUp
	case 0x03:
		acceptor.EscrowOrientation = enum.OrientationLeftDown
	}

	bil.Type = rune(reply[extDataIndex+11])
	bil.Series = rune(reply[extDataIndex+12])
	bil.Compatibility = rune(reply[extDataIndex+13])
	bil.Version = rune(reply[extDataIndex+14])

	return bil
}

////////

func (a *MpostObj) retrieveBillTable() {
	var index byte = 1
	for {
		payload := make([]byte, 6)
		acceptor.ConstructOmnibusCommand(payload, consts.CmdExpanded, 2, bill.TypeEnables)
		payload[1] = 0x02
		payload[5] = index

		var reply []byte
		var err error
		{
			for {
				reply, err = a.SendSynchronousCommand(payload)
				if err != nil {
					a.Log.Err("Error sending command", err)
					break
				}
				if len(reply) == 30 {
					break
				}
				time.Sleep(100 * time.Millisecond)
			}
		}

		if err != nil || len(reply) != 30 {
			break
		}

		ctl := reply[2]
		if (ctl&0x70) != 0x70 || reply[3] != 0x02 {
			break
		}

		if reply[10] == 0 {
			break
		}

		billFromTable := a.parseBillData(reply, 10)
		bill.Types = append(bill.Types, billFromTable)
		index++
	}

	for range bill.Types {
		bill.TypeEnables = append(bill.TypeEnables, true)
	}

	a.Log.Msg("Bill table retrieved")
}

func (a *MpostObj) setUpBillTable() {
	bill.SetUpTable(acceptor.ExpandedNoteReporting, func() {
		a.retrieveBillTable()
	})
}
