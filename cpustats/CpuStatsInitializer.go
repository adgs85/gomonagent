package cpustats

import (
	"github.com/adgs85/gomonagent/agentmessagesdispatcher"
)

func StartCpuUsageInfoLoop(sink agentmessagesdispatcher.StatSinkFuncType) {
	print("Cpu stats implementation: ")
	if cpuStatCfg.UsePortableCpuStat {
		println("portable")
		startPortableCpuUsageInfoLoop(sink)
	} else {
		println("low level")
		startCpuUsageInfoLoop(sink)
	}
}
