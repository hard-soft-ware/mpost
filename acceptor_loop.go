package mpost

import (
	"github.com/hard-soft-ware/mpost/acceptor"
	"github.com/hard-soft-ware/mpost/consts"
	"github.com/hard-soft-ware/mpost/enum"
	"time"
)

////////////////////////////////////

func (a *CAcceptor) PollingLoop(lg *LogStruct) []byte {
	startTickCount := time.Now()

	for {
		payload := []byte{consts.CmdOmnibus.Byte(), 0x00, 0x10, 0x00}

		reply, err := a.SendSynchronousCommand(payload)
		if err != nil {
			a.log.Err("PollingLoop", err)
		}

		if time.Since(startTickCount) > PollingDisconnectTimeout {
			a.RaiseDisconnectedEvent()
			startTickCount = time.Now()
		}

		if a.flashDownloadThread != nil {
			if acceptor.StopFlashDownloadThread {
				acceptor.StopFlashDownloadThread = true
				<-a.flashDownloadThread
				acceptor.Device.State = enum.StateIdling
				acceptor.WasStopped = true
				return nil
			}
		} else if a.openThread != nil {
			if acceptor.StopOpenThread {
				acceptor.StopOpenThread = false
				acceptor.StopWorkerThread = true

				<-a.openThread

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
