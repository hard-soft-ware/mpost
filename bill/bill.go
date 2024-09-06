package bill

import "fmt"

////////////////////////////////////

type BillStruct struct {
	Country       string
	Value         float64
	Type          rune
	Series        rune
	Compatibility rune
	Version       rune
}

var Bill BillStruct

var Types []BillStruct
var TypeEnables []bool

var Values []BillStruct
var ValueEnables []bool

////

func Reset() {
	Bill = BillStruct{}
}

func (b *BillStruct) ToString() string {
	return fmt.Sprintf("%s %.2f %c %c %c %c", b.Country, b.Value, b.Series, b.Type, b.Compatibility, b.Version)
}

func (b *BillStruct) GetCountry() string {
	return b.Country
}

func (b *BillStruct) GetValue() float64 {
	return b.Value
}

func (b *BillStruct) GetSeries() rune {
	return b.Series
}

func (b *BillStruct) GetType() rune {
	return b.Type
}

func (b *BillStruct) GetCompatibility() rune {
	return b.Compatibility
}

func (b *BillStruct) GetVersion() rune {
	return b.Version
}
