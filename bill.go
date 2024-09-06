package mpost

import (
	"fmt"
	"github.com/hard-soft-ware/mpost/acceptor"
	"github.com/hard-soft-ware/mpost/enum"
	"strconv"
	"strings"
)

////////////////////////////////////

// представление купюры
type CBill struct {
	Country       string
	Value         float64
	Type          rune
	Series        rune
	Compatibility rune
	Version       rune
}

func (a *CAcceptor) ParseBillData(reply []byte, extDataIndex int) CBill {
	var bill CBill

	if len(reply) < extDataIndex+15 {
		return bill
	}

	country := string(reply[extDataIndex+1 : extDataIndex+4])
	bill.Country = strings.TrimSpace(country)

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

	bill.Value = billValue
	a.docType = enum.DocumentBill
	a.wasDocTypeSetOnEscrow = acceptor.Device.State == enum.StateEscrow

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

	bill.Type = rune(reply[extDataIndex+11])
	bill.Series = rune(reply[extDataIndex+12])
	bill.Compatibility = rune(reply[extDataIndex+13])
	bill.Version = rune(reply[extDataIndex+14])

	return bill
}

////

func (b *CBill) ToString() string {
	return fmt.Sprintf("%s %.2f %c %c %c %c", b.Country, b.Value, b.Series, b.Type, b.Compatibility, b.Version)
}

func (b *CBill) GetCountry() string {
	return b.Country
}

func (b *CBill) GetValue() float64 {
	return b.Value
}

func (b *CBill) GetSeries() rune {
	return b.Series
}

func (b *CBill) GetType() rune {
	return b.Type
}

func (b *CBill) GetCompatibility() rune {
	return b.Compatibility
}

func (b *CBill) GetVersion() rune {
	return b.Version
}
