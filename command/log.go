package command

import "github.com/hard-soft-ware/mpost/consts"

////////////////////////////////////

func commandByte(b byte) consts.CmdType {
	if b%2 != 0 {
		b--
	}
	return consts.CmdType(b)
}

func Parse(msg []byte) (consts.CmdType, []byte) {
	if len(msg) < 4 {
		return 250, msg
	}
	return commandByte(msg[2]), msg[3 : len(msg)-2]
}
