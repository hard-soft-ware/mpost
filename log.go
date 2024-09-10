package mpost

import (
	"fmt"
	"github.com/hard-soft-ware/mpost/enum"
	"github.com/rs/zerolog"
	"strings"
)

/////////////////////////////////////////////////////////

type LogStruct struct {
	isEnable   bool
	log        zerolog.Logger
	index      string
	printBytes bool
}

func NewLog(log zerolog.Logger, root string, printBytes bool) LogStruct {
	obj := LogStruct{
		isEnable:   true,
		index:      root,
		log:        log,
		printBytes: printBytes,
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
	obj.log.Debug().Str("index", obj.index).Msg(msg)
}

func (obj *LogStruct) Err(msg string, err error) {
	if !obj.isEnable {
		return
	}
	obj.log.Debug().Str("index", obj.index).Err(err).Msg(msg)
}

func (obj *LogStruct) Bytes(msg string, data []byte) {
	if !obj.isEnable || !obj.printBytes {
		return
	}

	var sb strings.Builder
	for i, byteVal := range data {
		if i > 0 {
			sb.WriteString(" ")
		}
		fmt.Fprintf(&sb, "%02X", byteVal)
	}

	obj.log.Debug().Str("index", obj.index).Str("data", sb.String()).Int("len", len(data)).Msg(msg)
}

//

func (obj *LogStruct) Event(event enum.EventType) {
	obj.log.Debug().Str("index", obj.index).Str("event", event.String()).Msg("Event")
}
