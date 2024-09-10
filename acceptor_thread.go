package mpost

import (
	"github.com/hard-soft-ware/mpost/acceptor"
	"github.com/hard-soft-ware/mpost/bill"
	"github.com/hard-soft-ware/mpost/consts"
	"github.com/hard-soft-ware/mpost/enum"
	"time"
)

////////////////////////////////////

func (a *CAcceptor) OpenThread() {
	lg := a.log.New("Thread")

	replay := a.PollingLoop()

	if acceptor.WasStopped {
		lg.Msg("thread is stopped")
		return
	}

	a.dataLinkLayer.ProcessReply(replay)
	a.QueryDeviceCapabilities()

	if acceptor.Device.State != enum.StateDownloadRestart {
		a.SetUpBillTable()
		acceptor.Connected = true
		a.RaiseConnectedEvent()
	} else {
		a.RaiseDownloadRestartEvent()
	}
}

////

func (a *CAcceptor) MessageLoopThread() {
	lg := a.log.New("LoopThread")

	a.dataLinkLayer = a.NewCDataLinkLayer(lg)
	timeoutStart := time.Now()
	loopCycleCounter := 0

	for {
		if !acceptor.InSoftResetWaitForReply {
			time.Sleep(10 * time.Millisecond)
		} else {
			time.Sleep(1000 * time.Millisecond)
		}

		if time.Since(timeoutStart) > 30*time.Second {
			if acceptor.Device.State != enum.StateDownloading && acceptor.Device.State != enum.StateDownloadRestart {
				acceptor.Connected = false
				a.Close()
				acceptor.WasDisconnected = true
				timeoutStart = time.Now()
			}
		}

		if acceptor.StopWorkerThread {
			acceptor.StopWorkerThread = false
			lg.Msg("thread is stopped")
			return
		}

		select {
		case <-a.Ctx.Done():
			a.Close()
			return

		case message := <-a.messageQueue:
			loopCycleCounter = 0

			a.dataLinkLayer.SendPacket(message.Payload)
			reply, err := a.dataLinkLayer.ReceiveReply()
			if err != nil {
				a.log.Err("Invalid ReceiveReply", err)
				continue
			}

			if len(reply) > 0 {
				timeoutStart = time.Now()
				if acceptor.WasDisconnected {
					acceptor.WasDisconnected = false
					a.RaiseConnectedEvent()
				}
				if message.IsSynchronous {
					a.replyQueue <- reply
				} else {
					a.dataLinkLayer.ProcessReply(reply)
				}
			}

		default:
			loopCycleCounter++

			if loopCycleCounter > 9 {
				loopCycleCounter = 0

				payload := make([]byte, 4)
				acceptor.ConstructOmnibusCommand(payload, consts.CmdOmnibus, 1, bill.TypeEnables)

				a.dataLinkLayer.SendPacket(payload)

				reply, err := a.dataLinkLayer.ReceiveReply()
				if err != nil {
					a.Close()
					a.log.Err("Invalid loopCycleCounter", err)
					return
				}

				if len(reply) > 0 {
					timeoutStart = time.Now()

					if acceptor.WasDisconnected {
						acceptor.WasDisconnected = false

						if reply[2]&0x70 != 0x50 {
							acceptor.Connected = true
							a.RaiseConnectedEvent()
						} else {
							a.RaiseDownloadRestartEvent()
						}
					}

					if acceptor.InSoftResetWaitForReply {
						acceptor.InSoftResetWaitForReply = false
					}

					a.dataLinkLayer.ProcessReply(reply)
				}
			}
		}
	}
}
