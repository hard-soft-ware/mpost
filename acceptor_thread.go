package mpost

import (
	"github.com/hard-soft-ware/mpost/enum"
	"time"
)

////////////////////////////////////

func (a *CAcceptor) OpenThread() {
	lg := a.log.NewLog("Thread")

	replay := a.PollingLoop(lg)

	if a.wasStopped {
		lg.Debug().Msg("thread is stopped")
		return
	}

	a.dataLinkLayer.ProcessReply(replay)
	a.QueryDeviceCapabilities(lg)

	if a.deviceState != enum.StateDownloadRestart {
		a.SetUpBillTable()
		a.connected = true
		if a.shouldRaiseConnectedEvent {
			a.RaiseConnectedEvent()
		}
	} else {
		a.RaiseDownloadRestartEvent()
	}
}

////

func (a *CAcceptor) MessageLoopThread() {
	lg := a.log.NewLog("LoopThread")

	a.dataLinkLayer = a.NewCDataLinkLayer(lg)
	timeoutStart := time.Now()

	for {
		if !a.inSoftResetWaitForReply {
			time.Sleep(10 * time.Millisecond)
		} else {
			time.Sleep(1000 * time.Millisecond)
		}

		if time.Since(timeoutStart) > 30*time.Second {
			if a.deviceState != enum.StateDownloading && a.deviceState != enum.StateDownloadRestart {
				a.connected = false
				if a.shouldRaiseDisconnectedEvent {
					a.RaiseDisconnectedEvent()
				}
				a.wasDisconnected = true
				timeoutStart = time.Now()
			}
		}

		if a.stopWorkerThread {
			a.stopWorkerThread = false
			lg.Debug().Msg("thread is stopped")
			return
		}

		select {
		case message := <-a.messageQueue:
			lg.Debug().Bytes("payload", message.Payload).Msg("MessageLoopThread")

			a.dataLinkLayer.SendPacket(message.Payload)
			reply, err := a.dataLinkLayer.ReceiveReply()
			if err != nil {
				lg.Error().Err(err).Msg("Invalid ReceiveReply")
				continue
			}

			if len(reply) > 0 {
				timeoutStart = time.Now()
				if a.wasDisconnected {
					a.wasDisconnected = false
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
