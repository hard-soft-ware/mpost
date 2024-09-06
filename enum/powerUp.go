package enum

/* This file is automatically generated */

type PowerUpType byte

const (
	PowerUpA PowerUpType = 0
	PowerUpB PowerUpType = 1
	PowerUpC PowerUpType = 2
	PowerUpE PowerUpType = 3
)

const (
	PowerUpTextA = "A"
	PowerUpTextB = "B"
	PowerUpTextC = "C"
	PowerUpTextE = "E"
)

var PowerUpMap = map[PowerUpType]string{
	PowerUpA: PowerUpTextA,
	PowerUpB: PowerUpTextB,
	PowerUpC: PowerUpTextC,
	PowerUpE: PowerUpTextE,
}

func (obj PowerUpType) String() string {
	val, ok := PowerUpMap[obj]
	if ok {
		return val
	}
	return "Unknown PowerUpType"
}
