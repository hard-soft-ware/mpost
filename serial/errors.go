package serial

import "errors"

////////////////////////////////////

var (
	ErrNotConnect      = errors.New("not connect")
	ErrFailSendCommand = errors.New("fail send command")

	ErrReceivedInsufficientBytes = errors.New("received insufficient bytes")
	ErrInvalidSTX                = errors.New("invalid STX")
)
