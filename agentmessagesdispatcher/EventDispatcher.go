package agentmessagesdispatcher

import (
	"github.com/adgs85/gomonmarshalling/monmarshalling"
	"github.com/adgs85/gomonmarshalling/monmarshalling/envconfig"
)

const consoleOnlyModeKey = "console_only_mode"

var sinkArray []StatSinkFuncType = newDispatcherArray()

type StatSinkFuncType = func(stat monmarshalling.Stat)

func newDispatcherArray() []StatSinkFuncType {
	cfg := envconfig.GetViperConfig()

	arr := []StatSinkFuncType{}
	if cfg.GetBool(consoleOnlyModeKey) {
		arr = append(arr, SpewToConsoleSink)
	} else {
		arr = append(arr, StartHttpClientSenderLoopReturnSink())

	}

	return arr
}

func Dispatch(stat monmarshalling.Stat) {
	for _, f := range sinkArray {
		f(stat)
	}
}
