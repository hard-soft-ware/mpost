package acceptor

import "time"

////////////////////////////////////

type TimeoutStruct struct {
	Transaction time.Duration
	Download    time.Duration
}

var Timeout TimeoutStruct
