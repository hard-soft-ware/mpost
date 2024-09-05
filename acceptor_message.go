package mpost

////////////////////////////////////

type CMessage struct {
	Payload       []byte
	PayloadLength int
	IsSynchronous bool
}

func NewCMessage(payload []byte, isSynchronous bool) *CMessage {
	return &CMessage{
		Payload:       payload,
		PayloadLength: len(payload),
		IsSynchronous: isSynchronous,
	}
}
