package mpost

import (
	"github.com/hard-soft-ware/mpost/acceptor"
	"github.com/hard-soft-ware/mpost/bill"
	"github.com/hard-soft-ware/mpost/consts"
)

////////////////////////////////////

type dataObj struct {
	Acceptor *MpostObj
}

////

func (dl *dataObj) escrowXX(b byte) {
	if !acceptor.Connected {
		dl.Acceptor.Log.Msg("serial not connected")
		return
	}

	payload := make([]byte, 4)
	acceptor.ConstructOmnibusCommand(payload, consts.CmdOmnibus, 1, bill.TypeEnables)

	payload[2] |= b

	dl.Acceptor.SendAsynchronousCommand(payload)
}
