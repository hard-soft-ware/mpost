package acceptor

import "github.com/hard-soft-ware/mpost/enum"

////////////////////////////////////

type DeviceStruct struct {
	Failure  bool
	Model    int
	Paused   bool
	PowerUp  enum.PowerUpType
	Resets   int
	Revision int
	Stalled  bool
	State    enum.StateType
	Type     string

	Jammed bool
}

var Device DeviceStruct
