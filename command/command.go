package command

import (
	"github.com/hard-soft-ware/mpost/consts"
)

////////////////////////////////////

func CRC(command []byte) byte {
	var result byte

	end := int(command[1]) - 2
	for i := 1; i < end; i++ {
		result ^= command[i]
	}

	return result
}

func Create(payload []byte) []byte {
	commandLength := len(payload) + 4 // STX + Length char + ETX + Checksum

	command := make([]byte, 0, commandLength)
	command = append(command, consts.DataSTX.Byte())
	command = append(command, byte(commandLength))

	command = append(command, payload...)

	command = append(command, consts.DataETX.Byte())
	command = append(command, CRC(command))

	return command
}
