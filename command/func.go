package command

////////////////////////////////////

func CRC(command []byte) byte {
	var result byte

	end := int(command[1]) - 2
	for i := 1; i < end; i++ {
		result ^= command[i]
	}

	return result
}
