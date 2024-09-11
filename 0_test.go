package mpost

import (
	"fmt"
	"github.com/hard-soft-ware/mpost/enum"
	"testing"
	"time"
)

func TestConnect(t *testing.T) {
	//return

	a := DefAcceptor
	a.Log.Event = func(eventType enum.EventType, i int) {
		t.Log("Event", eventType.String(), i)
	}
	a.Log.Msg = func(s string) {
		t.Log("Msg", s)
	}
	a.Log.Err = func(s string, err error) {
		t.Error("Err", s, err.Error())
	}

	a.SetEnableAcceptance(true)
	a.SetEnableBarCodes(true)
	a.SetEnableBookmarks(true)

	a.AddHook(enum.EventConnected, func(acceptor *CAcceptor, i int) {
		fmt.Println("Connect")

		acceptor.SetUpBillTable()
	})
	a.AddHook(enum.EventDisconnected, func(acceptor *CAcceptor, i int) {
		fmt.Println("Disconnect")
	})
	a.AddHook(enum.EventRejected, func(acceptor *CAcceptor, i int) {
		fmt.Println("EventRejected")
	})
	a.AddHook(enum.EventReturned, func(acceptor *CAcceptor, i int) {
		fmt.Println("EventReturned")
	})

	a.Open("/dev/ttyUSB0", enum.PowerUpE)

	time.Sleep(2 * time.Second)
	t.Log(a.GetDeviceSerialNumber())
	t.Log(a.GetBill())
	t.Log(a.GetApplicationPN())
	t.Log(a.GetBootPN())
	t.Log(a.GetDeviceType())

	t.Log(a.GetBNFStatus().String())

	time.Sleep(100 * time.Second)
	a.Close()
}
