package mpost

////////////////////////////////////

type CDataLinkLayer struct {
	log *LogGlobalStruct

	Acceptor                       *CAcceptor
	AckToggleBit                   byte
	NakCount                       uint8
	CurrentCommand, EchoDetect     []byte
	PreviousCommand, PreviousReply []byte
	IdenticalCommandAndReplyCount  int
}

func (a *CAcceptor) NewCDataLinkLayer(lg *LogGlobalStruct) *CDataLinkLayer {
	return &CDataLinkLayer{
		Acceptor: a,
		log:      lg.NewLog("Data"),
	}
}

////

func (dl *CDataLinkLayer) ComputeCheckSum(command []byte) byte {
	var result byte

	end := int(command[1]) - 2
	for i := 1; i < end; i++ {
		result ^= command[i]
	}

	return result
}

//
