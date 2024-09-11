package mpost

import (
	"fmt"
	"github.com/hard-soft-ware/mpost/enum"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"testing"
	"time"
)

func TestConnect(t *testing.T) {
	return

	a := DefAcceptor
	a.AddLog(
		log.Output(zerolog.ConsoleWriter{
			Out:        os.Stdout,
			NoColor:    false,
			TimeFormat: "15:04:05",
		}),
		"TEST",
		true,
	)

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
