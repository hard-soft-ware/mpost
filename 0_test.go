package mpost

import (
	"fmt"
	"github.com/hard-soft-ware/mpost/enum"
	"testing"
	"time"
)

func TestConnect(t *testing.T) {
	a := NewCAcceptor(30*time.Second, 30*time.Second)

	a.AddHook(enum.EventConnected, func(acceptor *CAcceptor, i int) {
		fmt.Println("Connect")
	})
	a.AddHook(enum.EventDisconnected, func(acceptor *CAcceptor, i int) {
		fmt.Println("Disconnect")
	})

	a.Open("/dev/ttyUSB0", enum.PowerUpE)

	time.Sleep(2 * time.Second)
	t.Log(a.GetDeviceSerialNumber())
	t.Log(a.GetBNFStatus().String())

	a.Close()
	time.Sleep(1 * time.Second)
}
