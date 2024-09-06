package mpost

import (
	"github.com/hard-soft-ware/mpost/acceptor"
	"github.com/hard-soft-ware/mpost/enum"
	"time"
)

////////////////////////////////////

func (a *CAcceptor) OpenThread() {
	lg := a.log.New("Thread")

	replay := a.PollingLoop(lg)

	if acceptor.WasStopped {
		lg.Msg("thread is stopped")
		return
	}

	a.dataLinkLayer.ProcessReply(replay)
	a.QueryDeviceCapabilities()

	if acceptor.Device.State != enum.StateDownloadRestart {
		a.SetUpBillTable()
		acceptor.Connected = true
		if acceptor.ShouldRaise.ConnectedEvent {
			a.RaiseConnectedEvent()
		}
	} else {
		a.RaiseDownloadRestartEvent()
	}
}

////

func (a *CAcceptor) MessageLoopThread() {
	lg := a.log.New("LoopThread")

	a.dataLinkLayer = a.NewCDataLinkLayer(lg)
	timeoutStart := time.Now()

	for {
		if !acceptor.InSoftResetWaitForReply {
			time.Sleep(10 * time.Millisecond)
		} else {
			time.Sleep(1000 * time.Millisecond)
		}

		if time.Since(timeoutStart) > 30*time.Second {
			if acceptor.Device.State != enum.StateDownloading && acceptor.Device.State != enum.StateDownloadRestart {
				acceptor.Connected = false
				if acceptor.ShouldRaise.DisconnectedEvent {
					a.RaiseDisconnectedEvent()
				}
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
		case message := <-a.messageQueue:

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
		}
	}
}
