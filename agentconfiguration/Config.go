package agentconfiguration

import (
	"github.com/adgs85/gomonmarshalling/monmarshalling/envconfig"
	"github.com/davecgh/go-spew/spew"
)

var globalCfg *AgentConfig

type AgentConfig struct {
	envconfig.StatsConfig `mapstructure:",squash"`
	eventDispatcherConfig `mapstructure:",squash"`
	cpuCollectorConfig    `mapstructure:",squash"`
}

type eventDispatcherConfig struct {
	ConsoleOnlyMode bool `mapstructure:"console_only_mode"`
}

type cpuCollectorConfig struct {
	CpuStatsPath       string `mapstructure:"cpu_stats_path"`
	UsePortableCpuStat bool   `mapstructure:"cpu_stats_use_portable"`
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
	println(spew.Sdump(cfg))
	globalCfg = &cfg
}
