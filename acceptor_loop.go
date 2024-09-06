package mpost

import (
	"github.com/hard-soft-ware/mpost/acceptor"
	"github.com/hard-soft-ware/mpost/consts"
	"github.com/hard-soft-ware/mpost/enum"
	"time"
)

////////////////////////////////////

func (a *CAcceptor) PollingLoop(lg *LogGlobalStruct) []byte {
	startTickCount := time.Now()

	for {
		payload := []byte{consts.CmdOmnibus.Byte(), 0x00, 0x10, 0x00}

		reply, err := a.SendSynchronousCommand(payload)
		if err != nil {
			lg.Debug().Err(err).Msg("PollingLoop")
		}

		if time.Since(startTickCount) > PollingDisconnectTimeout {
			if acceptor.ShouldRaise.DisconnectedEvent {
				a.RaiseDisconnectedEvent()
			}
			startTickCount = time.Now()
		}

		if a.flashDownloadThread != nil {
			if a.stopFlashDownloadThread {
				a.stopFlashDownloadThread = true
				<-a.flashDownloadThread
				acceptor.Device.State = enum.StateIdling
				a.wasStopped = true
				return nil
			}
		} else if a.openThread != nil {
			if a.stopOpenThread {
				a.stopOpenThread = false
				a.stopWorkerThread = true

				<-a.openThread

				a.wasStopped = true
				a.Close()
				return nil
			}
		}

		if len(reply) > 0 {
			return reply
		}
	}
}
