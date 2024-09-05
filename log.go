package mpost

import (
	"github.com/rs/zerolog"
)

/////////////////////////////////////////////////////////

const TextLogMsgInit = "INIT"

type LogGlobalStruct struct {
	log   zerolog.Logger
	index string
}

////

func NewLog(log zerolog.Logger, root string) *LogGlobalStruct {
	obj := LogGlobalStruct{index: root}

	newLogger := log.With().Str("index", obj.index).Logger()
	obj.log = newLogger
	newLogger.Debug().Msg(TextLogMsgInit)

	return &obj
}

func (obj *LogGlobalStruct) NewLog(point string) *LogGlobalStruct {
	newObj := LogGlobalStruct{index: obj.index + "/" + point}

	newLogger := obj.log.With().Str("index", newObj.index).Logger()
	newObj.log = newLogger
	newLogger.Debug().Msg(TextLogMsgInit)

	return &newObj
}

////

func (obj *LogGlobalStruct) Debug() *zerolog.Event {
	return obj.log.Debug()
}
func (obj *LogGlobalStruct) Info() *zerolog.Event {
	return obj.log.Info()
}
func (obj *LogGlobalStruct) Log() *zerolog.Event {
	return obj.log.Log()
}
func (obj *LogGlobalStruct) Error() *zerolog.Event {
	return obj.log.Error()
}
func (obj *LogGlobalStruct) Panic() *zerolog.Event {
	return obj.log.Panic()
}
func (obj *LogGlobalStruct) Fatal() *zerolog.Event {
	return obj.log.Fatal()
}
func (obj *LogGlobalStruct) Warn() *zerolog.Event {
	return obj.log.Warn()
}
func (obj *LogGlobalStruct) Trace() *zerolog.Event {
	return obj.log.Trace()
}
