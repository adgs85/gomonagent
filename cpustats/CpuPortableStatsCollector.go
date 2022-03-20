package cpustats

import (
	"time"

	"github.com/adgs85/gomonagent/agentmessagesdispatcher"
	"github.com/shirou/gopsutil/v3/cpu"
)

func startPortableCpuUsageInfoLoop(sink agentmessagesdispatcher.StatSinkFuncType) {
	go func() {
		collectPortableTotalCpuUsage(sink)
	}()
}

func collectPortableTotalCpuUsage(sink agentmessagesdispatcher.StatSinkFuncType) {
	for {
		usage, err := cpu.Percent(time.Second, false) //this waits 1 sec
		if err != nil || len(usage) == 0 {
			logger.Panic("ERROR fetching cpu stats")
		}

		sink(*newCpuUsageStat("cpu", usage[0]))
	}
}
