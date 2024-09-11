package mpost

////////////////////////////////////

type CDataLinkLayer struct {
	Acceptor                      *MpostObj
	IdenticalCommandAndReplyCount int
}

func (a *MpostObj) newCDataLinkLayer() *CDataLinkLayer {
	return &CDataLinkLayer{
		Acceptor: a,
	}
}
