package main

import (
	"sync"

	"github.com/adgs85/gomonagent/agentconfiguration"
	"github.com/adgs85/gomonagent/agentheartbeat"
	"github.com/adgs85/gomonagent/agentmessagesdispatcher"
	"github.com/adgs85/gomonagent/cpustats"
	"github.com/adgs85/gomonagent/diskstats"
	"github.com/adgs85/gomonmarshalling/monmarshalling/envconfig"
)

func main() {

	agentconfiguration.InitConfig(func(cfg *agentconfiguration.AgentConfig) {
		envconfig.GetViperConfig().Unmarshal(&cfg)
	})

	var wg sync.WaitGroup
	wg.Add(1)

	startStatCollectors()

	wg.Wait()
}

func startStatCollectors() {

	agentheartbeat.StartHeartBeat()

	dispatcher := agentmessagesdispatcher.NewDispatcher()

	diskstats.StartDiskInfoLoopGoRoutine(dispatcher.Dispatch)

	cpustats.StartCpuUsageInfoLoop(dispatcher.Dispatch)

}
