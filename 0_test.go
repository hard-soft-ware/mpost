package mpost

import (
	"fmt"
	"github.com/hard-soft-ware/mpost/consts"
	"github.com/hard-soft-ware/mpost/enum"
	"go.bug.st/serial"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"testing"
	"time"
)

func TestConnect(t *testing.T) {
	ports, err := serial.GetPortsList()
	if err != nil {
		t.Fatal(err)
		return
	}

	if len(ports) == 0 {
		t.Log("no serial ports found")
		return
	}

	a := New()
	printByte := true
	{
		a.Log.Event = func(eventType enum.EventType, i int) {
			if eventType == enum.EventJamCleared {
				return
			}
			t.Log("Event", eventType.String(), i)
		}
		a.Log.Msg = func(s string) {
			t.Log("Msg", s)
		}
		a.Log.Err = func(s string, err error) {
			t.Error("Err", s, err.Error())
		}

		if printByte { //serial log
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
				fmt.Println("<\t", cmdType.String()+"\t<<<\t", byteToStr(bytes))
			}
			a.Log.SerialSend = func(cmdType consts.CmdType, bytes []byte) {
				fmt.Println(">\t", cmdType.String()+"\t>>>\t", byteToStr(bytes))
			}
		}
	}

	a.Method.Enable.Set.Acceptance(true)

	connCh := make(chan bool)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	{
		a.AddHook(enum.EventConnected, func(i int) {
			connCh <- true
			fmt.Println("HOOK\t\tConnect")
		})
		a.AddHook(enum.EventDisconnected, func(i int) {
			connCh <- false
			fmt.Println("HOOK\t\tDisconnect")
		})
		a.AddHook(enum.EventRejected, func(i int) {
			fmt.Println("HOOK\t\tRejected")
		})
		a.AddHook(enum.EventReturned, func(i int) {
			fmt.Println("HOOK\t\tReturned")
		})
	}

	fmt.Println("Connect to:\t\t" + ports[0])
	a.Open(ports[0], enum.PowerUpE)

	for {
		select {

		case <-sigChan:
			t.Log("Sig Close")
			return

		case <-time.After(time.Second * 100):
			t.Log("close Timeout")
			return

		case status := <-connCh:
			if !status {
				t.Log("Invalid Connect")
				return
			}
			t.Log(a.Method.GetDeviceSerialNumber())
			t.Log(a.Method.Bill.Get.Self())
			t.Log(a.Method.Application.Get.PN())
			t.Log(a.Method.Get.BootPN())
			t.Log(a.Method.GetDeviceType())

			t.Log(a.Method.Get.BNFStatus().String())

		default:
			continue
		}
	}

	a.Close()
}
