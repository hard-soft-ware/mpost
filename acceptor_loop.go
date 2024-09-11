package mpost

import (
	"github.com/hard-soft-ware/mpost/acceptor"
	"github.com/hard-soft-ware/mpost/bill"
	"github.com/hard-soft-ware/mpost/consts"
	"github.com/hard-soft-ware/mpost/enum"
	"github.com/hard-soft-ware/mpost/hook"
	"time"
)

////////////////////////////////////

func (a *MpostObj) pollingLoop() []byte {
	startTickCount := time.Now()

	for {
		payload := []byte{consts.CmdOmnibus.Byte(), 0x00, 0x10, 0x00}

		reply, err := a.SendSynchronousCommand(payload)
		if err != nil {
			a.Log.Err("PollingLoop", err)
		}

		if time.Since(startTickCount) > PollingDisconnectTimeout {
			a.Close()
			startTickCount = time.Now()
		}

		if !acceptor.FlashDownloadThread {
			if acceptor.StopFlashDownloadThread {
				acceptor.StopFlashDownloadThread = true
				acceptor.FlashDownloadThread = true
				acceptor.Device.State = enum.StateIdling
				acceptor.WasStopped = true
				return nil
			}
		}
		if !acceptor.OpenThread {
			if acceptor.StopOpenThread {
				acceptor.StopOpenThread = false
				acceptor.StopWorkerThread = true
				acceptor.OpenThread = true
				acceptor.WasStopped = true
				a.Close()
				return nil
			}
		}

		if len(reply) > 0 {
			return reply
		}
	}
}

////////

func (a *MpostObj) openThread() {
	replay := a.pollingLoop()

	if acceptor.WasStopped {
		a.Log.Msg("thread is stopped")
		return
	}

	a.dataLinkLayer.ProcessReply(replay)
	a.queryDeviceCapabilities()

	if acceptor.Device.State != enum.StateDownloadRestart {
		a.SetUpBillTable()
		acceptor.Connected = true
		hook.Raise.Connected()
	} else {
		hook.Raise.Download.Restart()
	}
}

func (a *MpostObj) messageLoopThread() {
	a.dataLinkLayer = a.newCDataLinkLayer()
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
			a.Log.Msg("thread is stopped")
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
				a.Log.Err("Invalid ReceiveReply", err)
				continue
			}

			if len(reply) > 0 {
				timeoutStart = time.Now()
				if acceptor.WasDisconnected {
					acceptor.WasDisconnected = false
					hook.Raise.Connected()
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
					a.Log.Err("Invalid loopCycleCounter", err)
					return
				}

				if len(reply) > 0 {
					timeoutStart = time.Now()

					if acceptor.WasDisconnected {
						acceptor.WasDisconnected = false

						if reply[2]&0x70 != 0x50 {
							acceptor.Connected = true
							hook.Raise.Connected()
						} else {
							hook.Raise.Download.Restart()
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
