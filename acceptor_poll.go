package mpost

import "github.com/hard-soft-ware/mpost/acceptor"

////////////////////////////////////

type CSuppressStandardPoll struct {
	acceptor *CAcceptor
}

func NewCSuppressStandardPoll(a *CAcceptor) *CSuppressStandardPoll {
	s := &CSuppressStandardPoll{
		acceptor: a,
	}

	acceptor.SuppressStandardPoll = true
	return s
}

func (s *CSuppressStandardPoll) Release() {
	acceptor.SuppressStandardPoll = false
}
