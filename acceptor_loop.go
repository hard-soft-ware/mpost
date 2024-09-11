package mpost

import (
	"github.com/hard-soft-ware/mpost/acceptor"
	"github.com/hard-soft-ware/mpost/consts"
	"github.com/hard-soft-ware/mpost/enum"
	"time"
)

////////////////////////////////////

func (a *CAcceptor) PollingLoop() []byte {
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

		if !a.flashDownloadThread {
			if acceptor.StopFlashDownloadThread {
				acceptor.StopFlashDownloadThread = true
				a.flashDownloadThread = true
				acceptor.Device.State = enum.StateIdling
				acceptor.WasStopped = true
				return nil
			}
		}
		if !a.openThread {
			if acceptor.StopOpenThread {
				acceptor.StopOpenThread = false
				acceptor.StopWorkerThread = true
				a.openThread = true
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
