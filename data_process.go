package mpost

import (
	"github.com/hard-soft-ware/mpost/acceptor"
	"github.com/hard-soft-ware/mpost/enum"
)

////////////////////////////////////

func (dl *dataObj) ProcessStandardOmnibusReply(reply []byte) {
	if len(reply) < 9 {
		return
	}

	dl.Acceptor.processData4(reply[7])
	dl.Acceptor.processData0(reply[3])
	dl.Acceptor.processData1(reply[4])
	dl.Acceptor.processData2(reply[5])
	dl.Acceptor.processData3(reply[6])
	dl.Acceptor.processData5(reply[8])

	dl.RaiseEvents()
}

func (dl *dataObj) ProcessExtendedOmnibusBarCodeReply(reply []byte) {
	if len(reply) < 38 {
		return
	}

	dl.Acceptor.processData4(reply[8])
	dl.Acceptor.processData0(reply[4])
	dl.Acceptor.processData1(reply[5])
	dl.Acceptor.processData2(reply[6])
	dl.Acceptor.processData3(reply[7])
	dl.Acceptor.processData5(reply[9])

	if acceptor.Device.State == enum.StateEscrow {
		acceptor.BarCode = ""
		for i := 10; i < 38; i++ {
			if reply[i] != '(' {
				acceptor.BarCode += string(reply[i])
			} else {
				break
			}
		}
		dl.Acceptor.DocType = enum.DocumentBarcode
	}
}

func (dl *dataObj) ProcessExtendedOmnibusExpandedNoteReply(reply []byte) {
	if len(reply) < 10 {
		return
	}

	dl.Acceptor.processData4(reply[8])
	dl.Acceptor.processData0(reply[4])
	dl.Acceptor.processData1(reply[5])
	dl.Acceptor.processData2(reply[6])
	dl.Acceptor.processData3(reply[7])
	dl.Acceptor.processData5(reply[9])
}

func (dl *dataObj) ProcessExtendedOmnibusExpandedCouponReply(reply []byte) {
	if len(reply) < 15 {
		return
	}

	dl.Acceptor.processData4(reply[8])
	dl.Acceptor.processData0(reply[4])
	dl.Acceptor.processData1(reply[5])
	dl.Acceptor.processData2(reply[6])
	dl.Acceptor.processData3(reply[7])
	dl.Acceptor.processData5(reply[9])

	if acceptor.Device.State == enum.StateEscrow || (acceptor.Device.State == enum.StateStacked && !acceptor.WasDocTypeSetOnEscrow) {
		couponData := ((int(reply[10]) & 0x0F) << 12) +
			((int(reply[11]) & 0x0F) << 8) +
			((int(reply[12]) & 0x0F) << 4) +
			(int(reply[13]) & 0x0F)

		value := float64(couponData & 0x07)
		if value == 3.0 {
			value = 5.0
		}

		ownerID := (couponData & 0xFFF8) >> 3

		dl.Acceptor.coupon = newCoupon(ownerID, value)

		dl.Acceptor.DocType = enum.DocumentCoupon
		acceptor.WasDocTypeSetOnEscrow = acceptor.Device.State == enum.StateEscrow
	}
}
