package hook

import "github.com/hard-soft-ware/mpost/enum"

////////////////////////////////////

var eventHandlers = make(map[enum.EventType]func(int), enum.Event_End)

func Add(ev enum.EventType, h func(int)) {
	eventHandlers[ev] = h
}
