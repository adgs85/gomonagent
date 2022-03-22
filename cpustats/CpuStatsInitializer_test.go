package cpustats

import (
	"os"
	. "testing"
	"time"

	"github.com/adgs85/gomonagent/agentconfiguration"
	"github.com/adgs85/gomonmarshalling/monmarshalling"
	"github.com/stretchr/testify/assert"
)

func TestPortableCpuStatsMetaDataPayload(t *T) {

	const hostName = "Test"

	initTestConfig(func(cfg *agentconfiguration.AgentConfig) {
		cfg.UsePortableCpuStat = true
		cfg.HostName = hostName
	})

	c := make(chan struct{})

	StartCpuUsageInfoLoop(func(stat monmarshalling.Stat) {
		assert.Equal(t, monmarshalling.StatsBeatMessageTypeName, stat.MetaData.MessageType)
		assert.Equal(t, CpuUsagesStatType, stat.MetaData.StatType)
		assert.Equal(t, hostName, stat.MetaData.HostName)
		assert.Equal(t, int(time.Second.Milliseconds()), stat.MetaData.PollRateMs)
		hostName, _ := os.Hostname()
		assert.Equal(t, hostName, stat.MetaData.InstanceName)
		defer close(c)
	})

	select {
	case <-c:
	case <-time.After(2 * time.Second):
		t.Error("Timeout")
	}
}

func initTestConfig(testCfgAcceptor agentconfiguration.AgentConfigAcceptor) {
	agentconfiguration.InitConfig(func(cfg *agentconfiguration.AgentConfig) {
		testCfgAcceptor(cfg)
	})

}
