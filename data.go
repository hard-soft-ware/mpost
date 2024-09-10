package mpost

////////////////////////////////////

type CDataLinkLayer struct {
	log *LogStruct

	Acceptor                       *CAcceptor
	CurrentCommand, EchoDetect     []byte
	PreviousCommand, PreviousReply []byte
	IdenticalCommandAndReplyCount  int
}

func (a *CAcceptor) NewCDataLinkLayer(lg *LogStruct) *CDataLinkLayer {
	return &CDataLinkLayer{
		Acceptor: a,
		log:      lg.New("Data"),
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
