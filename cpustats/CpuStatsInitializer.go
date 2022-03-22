package cpustats

import (
	"github.com/adgs85/gomonagent/agentconfiguration"
	"github.com/adgs85/gomonagent/agentlogger"
	"github.com/adgs85/gomonagent/agentmessagesdispatcher"
)

var logger = agentlogger.Logger()

func StartCpuUsageInfoLoop(sink agentmessagesdispatcher.StatSinkFuncType) {
	print("Cpu stats implementation: ")
	if agentconfiguration.GlobalCfg().UsePortableCpuStat {
		println("portable")
		startPortableCpuUsageInfoLoop(sink)
	} else {
		println("low level")
		startCpuUsageInfoLoop(sink)
	}
}
