package main

import (
	"sync"

	"github.com/adgs85/gomonagent/agentheartbeat"
	"github.com/adgs85/gomonagent/agentmessagesdispatcher"
	"github.com/adgs85/gomonagent/cpustats"
	"github.com/adgs85/gomonagent/diskstats"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	startStatCollectors()

	wg.Wait()
}

func startStatCollectors() {

	agentheartbeat.StartHeartBeat()

	diskstats.StartDiskInfoLoopGoRoutine(agentmessagesdispatcher.Dispatch)

	cpustats.StartCpuUsageInfoLoop(agentmessagesdispatcher.Dispatch)

}
