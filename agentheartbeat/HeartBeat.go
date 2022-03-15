package agentheartbeat

import (
	"fmt"
	"time"

	"github.com/adgs85/gomonagent/agentconfiguration"
	"github.com/adgs85/gomonagent/agentmessagesdispatcher"
	"github.com/adgs85/gomonmarshalling/monmarshalling"
)

func StartHeartBeat() {
	cfg := agentconfiguration.GlobalCfg()
	heartBeatSleep := time.Duration(agentconfiguration.GlobalCfg().HeartBeatSec) * 1000 * time.Millisecond
	pollRateMs := int(heartBeatSleep.Milliseconds())
	fmt.Printf("%v", heartBeatSleep)
	go func() {
		httpClient := agentmessagesdispatcher.NewHttpClient(cfg.ServerUrl)

		for {
			metaData := *monmarshalling.NewHeartBeatMetaDataWithTs()
			metaData.HostName = cfg.HostName
			metaData.PollRateMs = pollRateMs
			httpClient.PostImmediately(&monmarshalling.Stat{MetaData: metaData})
			time.Sleep(heartBeatSleep)
		}
	}()
}
