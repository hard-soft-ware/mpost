package mpost

////////////////////////////////////

type CSuppressStandardPoll struct {
	acceptor *CAcceptor
}

func NewCSuppressStandardPoll(acceptor *CAcceptor) *CSuppressStandardPoll {
	s := &CSuppressStandardPoll{
		acceptor: acceptor,
	}
	s.acceptor.suppressStandardPoll = true
	return s
}

func (s *CSuppressStandardPoll) Release() {
	s.acceptor.suppressStandardPoll = false
}
