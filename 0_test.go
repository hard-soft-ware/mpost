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
	a := DefAcceptor
	a.AddLog(
		log.Output(zerolog.ConsoleWriter{
			Out:        os.Stdout,
			NoColor:    false,
			TimeFormat: "15:04:05",
		}),
		"TEST",
	)

	a.AddHook(enum.EventConnected, func(acceptor *CAcceptor, i int) {
		fmt.Println("Connect")

		acceptor.SetEnableAcceptance(true)
		acceptor.SetEnableBarCodes(true)
		acceptor.SetEnableBookmarks(true)
	})
	a.AddHook(enum.EventDisconnected, func(acceptor *CAcceptor, i int) {
		fmt.Println("Disconnect")
	})

	a.Open("/dev/ttyUSB0", enum.PowerUpE)

	time.Sleep(2 * time.Second)
	t.Log(a.GetDeviceSerialNumber())
	t.Log(a.GetBill())
	t.Log(a.GetApplicationPN())
	t.Log(a.GetBootPN())
	t.Log(a.GetDeviceState().String())

	a.Close()
	time.Sleep(1 * time.Second)
}
