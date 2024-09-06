package mpost

import (
	"fmt"
	"github.com/hard-soft-ware/mpost/enum"
	"github.com/rs/zerolog"
	"strings"
)

/////////////////////////////////////////////////////////

type LogStruct struct {
	isEnable bool
	log      zerolog.Logger
	index    string
}

func NewLog(log zerolog.Logger, root string) LogStruct {
	obj := LogStruct{
		isEnable: true,
		index:    root,
		log:      log,
	}
	return obj
}

func (obj *LogStruct) New(point string) *LogStruct {
	newObj := *obj
	newObj.index = obj.index + "/" + point
	return &newObj
}

////

func (obj *LogStruct) Msg(msg string) {
	if !obj.isEnable {
		return
	}
	obj.log.Debug().Msg(msg)
}

func (obj *LogStruct) Err(msg string, err error) {
	if !obj.isEnable {
		return
	}
	obj.log.Debug().Err(err).Msg(msg)
}

func (obj *LogStruct) Bytes(msg string, data []byte) {
	if !obj.isEnable {
		return
	}

	var sb strings.Builder
	for i, byteVal := range data {
		if i > 0 {
			sb.WriteString(" ")
		}
		fmt.Fprintf(&sb, "%02X", byteVal)
	}

	obj.log.Debug().Str("data", sb.String()).Int("len", len(data)).Msg(msg)
}

//

func (obj *LogStruct) Event(event enum.EventType) {
	obj.log.Debug().Str("event", event.String()).Msg("Event")
}
