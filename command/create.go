package command

import (
	"github.com/hard-soft-ware/mpost/consts"
)

////////////////////////////////////

var coin bool

func CreateMsg(payload []byte) []byte {
	if !coin {
		coin = true
	} else {
		coin = false
		payload[0] += 1
	}

	commandLength := len(payload) + 4 // STX + Length char + ETX + Checksum

	command := make([]byte, 0, commandLength)
	command = append(command, consts.DataSTX.Byte())
	command = append(command, byte(commandLength))

	command = append(command, payload...)

	command = append(command, consts.DataETX.Byte())
	command = append(command, CRC(command))

	return command
}
