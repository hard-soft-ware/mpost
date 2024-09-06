package serial

import (
	"go.bug.st/serial"
	"time"
)

////////////////////////////////////

func (obj *SerialStruct) Port() serial.Port {
	return obj.port
}

func (obj *SerialStruct) Write(p []byte) (n int, err error) {
	if !*obj.connect {
		return 0, ErrNotConnect
	}
	return obj.port.Write(p)
}

func (obj *SerialStruct) SetTimeout(t time.Duration) error {
	if !*obj.connect {
		return ErrNotConnect
	}
	return obj.port.SetReadTimeout(t)
}
