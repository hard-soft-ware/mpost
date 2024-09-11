package mpost

import (
	"fmt"
	"github.com/hard-soft-ware/mpost/consts"
	"github.com/hard-soft-ware/mpost/enum"
	"strings"
	"testing"
	"time"
)

func TestConnect(t *testing.T) {
	//return

	a := DefAcceptor
	{
		a.Log.Event = func(eventType enum.EventType, i int) {
			t.Log("Event", eventType.String(), i)
		}
		a.Log.Msg = func(s string) {
			t.Log("Msg", s)
		}
		a.Log.Err = func(s string, err error) {
			t.Error("Err", s, err.Error())
		}

		byteToStr := func(bytes []byte) string {
			var sb strings.Builder
			for i, byteVal := range bytes {
				if i > 0 {
					sb.WriteString(" ")
				}
				fmt.Fprintf(&sb, "%02X", byteVal)
			}
			return sb.String()
		}

		a.Log.SerialRead = func(cmdType consts.CmdType, bytes []byte) {

			fmt.Println("<<< \t", cmdType.String(), byteToStr(bytes))
		}
		a.Log.SerialSend = func(cmdType consts.CmdType, bytes []byte) {
			fmt.Println(">>> \t", cmdType.String(), byteToStr(bytes))
		}
	}

	a.SetEnableAcceptance(true)

	a.AddHook(enum.EventConnected, func(i int) {
		fmt.Println("Connect")

		a.SetUpBillTable()
	})
	a.AddHook(enum.EventDisconnected, func(i int) {
		fmt.Println("Disconnect")
	})
	a.AddHook(enum.EventRejected, func(i int) {
		fmt.Println("EventRejected")
	})
	a.AddHook(enum.EventReturned, func(i int) {
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
