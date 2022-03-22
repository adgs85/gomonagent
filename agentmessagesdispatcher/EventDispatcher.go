package agentmessagesdispatcher

import (
	"github.com/adgs85/gomonagent/agentconfiguration"
	"github.com/adgs85/gomonmarshalling/monmarshalling"
)

type StatSinkFuncType = func(stat monmarshalling.Stat)

type EventDispatcher interface {
	Dispatch(stat monmarshalling.Stat)
}

type dispatcher struct {
	sinkArray []StatSinkFuncType
}

func (d dispatcher) Dispatch(stat monmarshalling.Stat) {
	dispatch(stat, d.sinkArray)
}

func NewDispatcher() EventDispatcher {
	return dispatcher{newDispatcherArray()}
}

func newDispatcherArray() []StatSinkFuncType {
	arr := []StatSinkFuncType{}
	if agentconfiguration.GlobalCfg().ConsoleOnlyMode {
		arr = append(arr, SpewToConsoleSink)
	} else {
		arr = append(arr, StartHttpClientSenderLoopReturnSink())

	}

	return arr
}

func dispatch(stat monmarshalling.Stat, sinkArray []StatSinkFuncType) {
	for _, f := range sinkArray {
		f(stat)
	}
}
