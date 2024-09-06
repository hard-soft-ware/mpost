package mpost

import (
	"github.com/hard-soft-ware/mpost/enum"
	"testing"
	"time"
)

func TestName(t *testing.T) {
	a := NewCAcceptor(30*time.Second, 30*time.Second)
	a.Open("/dev/ttyUSB0", enum.PowerUpB)

	time.Sleep(30 * time.Second)
	t.Log(a.GetBNFStatus().String())
	t.Log(a.GetDeviceSerialNumber())

	a.Close()
}
