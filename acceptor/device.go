package acceptor

import "github.com/hard-soft-ware/mpost/enum"

////////////////////////////////////

type DeviceStruct struct {
	Failure      bool
	Model        int
	Paused       bool
	PortName     string
	PowerUp      enum.PowerUpType
	Resets       int
	Revision     int
	SerialNumber string
	Stalled      bool
	State        enum.StateType
	Type         string
}

var Device DeviceStruct
