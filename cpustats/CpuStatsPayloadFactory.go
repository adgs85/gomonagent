package cpustats

import (
	"encoding/json"
	"time"

	"github.com/adgs85/gomonagent/agentlogger"
	"github.com/adgs85/gomonmarshalling/monmarshalling"
)

const CpuUsagesStatType = "cpu"

type cpuStatsPayload struct {
	CpuName         string  `mapstructure:"cpuName"`
	CpuUsagePercent float64 `mapstructure:"cpuUsagePercent"`
}

var hardCodedCpuPollRateMs = int(time.Second.Milliseconds())

func newCpuUsageStat(cpuName string, cpuUsagePercent float64) *monmarshalling.Stat {
	stat := newStat()
	payload := []cpuStatsPayload{{CpuName: cpuName, CpuUsagePercent: cpuUsagePercent}}

	payloadStr, err := json.Marshal(payload)
	if err != nil {
		agentlogger.Logger().Fatalln(err)
	}
	stat.Payload = string(payloadStr)
	return stat
}

func newStat() *monmarshalling.Stat {
	statsMetadata := monmarshalling.NewStatsMetaDataWithTs()
	statsMetadata.StatType = CpuUsagesStatType
	statsMetadata.PollRateMs = hardCodedCpuPollRateMs
	statsMetadata.HostName = cpuStatCfg.StatsConfig.HostName
	stat := monmarshalling.Stat{MetaData: *statsMetadata}
	return &stat
}
