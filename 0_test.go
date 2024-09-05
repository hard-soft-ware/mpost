package mpost

import (
	"testing"
	"time"
)

func TestName(t *testing.T) {
	a := NewCAcceptor(30*time.Second, 30*time.Second)
	a.Open("/dev/ttyUSB0", B)

	time.Sleep(5 * time.Second)
	t.Log(a.GetDeviceSerialNumber())

	a.Close()
}
