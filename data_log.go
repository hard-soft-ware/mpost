package mpost

////////////////////////////////////

func (dl *CDataLinkLayer) LogCommandAndReply(command, reply []byte, wasEchoDiscarded bool) {
	dl.log.Debug().Bytes("command", command).Bytes("reply", reply).Bool("wasEchoDiscarded", wasEchoDiscarded).Bool("wasEchoDiscarded", wasEchoDiscarded).Send()
}

func (dl *CDataLinkLayer) FlushIdenticalTransactionsToLog() {
	if dl.IdenticalCommandAndReplyCount > 0 {
		dl.log.Debug().Int("IdenticalCommandAndReplyCount", dl.IdenticalCommandAndReplyCount).Send()
	}
}
