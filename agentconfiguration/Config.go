package agentconfiguration

import (
	"github.com/adgs85/gomonmarshalling/monmarshalling/envconfig"
)

var globalCfg *AgentConfig

type AgentConfig struct {
	envconfig.StatsConfig        `mapstructure:",squash"`
	eventDispatcherConfig        `mapstructure:",squash"`
	cpuCollectorConfig           `mapstructure:",squash"`
	statChannelSinkFactoryConfig `mapstructure:",squash"`
	diskCollectorConfig          `mapstructure:",squash"`
}

type eventDispatcherConfig struct {
	ConsoleOnlyMode bool `mapstructure:"console_only_mode"`
}

type cpuCollectorConfig struct {
	CpuStatsPath       string `mapstructure:"cpu_stats_path"`
	UsePortableCpuStat bool   `mapstructure:"cpu_stats_use_portable"`
}

type diskCollectorConfig struct {
	DiskPollingRateMs int    `mapstructure:"disk_polling_rate_ms"`
	DiskFreeSpacePath string `mapstructure:"disk_free_space_path"`
}

type statChannelSinkFactoryConfig struct {
	SenderStatStallMs int `mapstructure:"sender_stall_ms"`
}

func GlobalCfg() *AgentConfig {
	if globalCfg == nil {
		panic("Configuration not initialized")
	}

	return globalCfg
}

func InitConfig(initFunc func(*AgentConfig)) {
	cfg := AgentConfig{}
	initFunc(&cfg)
	globalCfg = &cfg
}
