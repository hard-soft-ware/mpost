package mpost

import "time"

////////////////////////////////////

func (a *CAcceptor) PollingLoop(lg *LogGlobalStruct) []byte {
	startTickCount := time.Now()

	for {
		payload := []byte{CmdOmnibus, 0x00, 0x10, 0x00}

		reply, err := a.SendSynchronousCommand(payload)
		if err != nil {
			lg.Debug().Err(err).Msg("PollingLoop")
		}

		if time.Since(startTickCount) > PollingDisconnectTimeout {
			if a.shouldRaiseDisconnectedEvent {
				a.RaiseDisconnectedEvent()
			}
			startTickCount = time.Now()
		}

		if a.flashDownloadThread != nil {
			if a.stopFlashDownloadThread {
				a.stopFlashDownloadThread = true
				<-a.flashDownloadThread
				a.deviceState = Idling
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
