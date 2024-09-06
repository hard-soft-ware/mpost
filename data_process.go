package mpost

import "github.com/hard-soft-ware/mpost/enum"

////////////////////////////////////

func (dl *CDataLinkLayer) ProcessStandardOmnibusReply(reply []byte) {
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

func (dl *CDataLinkLayer) ProcessExtendedOmnibusBarCodeReply(reply []byte) {
	if len(reply) < 38 {
		return
	}

	dl.Acceptor.processData4(reply[8])
	dl.Acceptor.processData0(reply[4])
	dl.Acceptor.processData1(reply[5])
	dl.Acceptor.processData2(reply[6])
	dl.Acceptor.processData3(reply[7])
	dl.Acceptor.processData5(reply[9])

	if dl.Acceptor.deviceState == enum.StateEscrow {
		dl.Acceptor.barCode = ""
		for i := 10; i < 38; i++ {
			if reply[i] != '(' {
				dl.Acceptor.barCode += string(reply[i])
			} else {
				break
			}
		}
		dl.Acceptor.docType = enum.DocumentBarcode
	}
}

func (dl *CDataLinkLayer) ProcessExtendedOmnibusExpandedNoteReply(reply []byte) {
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

func (dl *CDataLinkLayer) ProcessExtendedOmnibusExpandedCouponReply(reply []byte) {
	if len(reply) < 15 {
		return
	}

	dl.Acceptor.processData4(reply[8])
	dl.Acceptor.processData0(reply[4])
	dl.Acceptor.processData1(reply[5])
	dl.Acceptor.processData2(reply[6])
	dl.Acceptor.processData3(reply[7])
	dl.Acceptor.processData5(reply[9])

	if dl.Acceptor.deviceState == enum.StateEscrow || (dl.Acceptor.deviceState == enum.StateStacked && !dl.Acceptor.wasDocTypeSetOnEscrow) {
		couponData := ((int(reply[10]) & 0x0F) << 12) +
			((int(reply[11]) & 0x0F) << 8) +
			((int(reply[12]) & 0x0F) << 4) +
			(int(reply[13]) & 0x0F)

		value := float64(couponData & 0x07)
		if value == 3.0 {
			value = 5.0
		}

		ownerID := (couponData & 0xFFF8) >> 3

		dl.Acceptor.coupon = NewCCoupon(ownerID, value)

		dl.Acceptor.docType = enum.DocumentCoupon
		dl.Acceptor.wasDocTypeSetOnEscrow = dl.Acceptor.deviceState == enum.StateEscrow
	}
}
