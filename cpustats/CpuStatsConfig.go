package cpustats

import (
	"github.com/adgs85/gomonagent/agentconfiguration"
	"github.com/adgs85/gomonmarshalling/monmarshalling/envconfig"
)

type cpuCollectorConfig struct {
	CpuStatsPath string `mapstructure:"cpu_stats_path"`
	*envconfig.StatsConfig
}

func initConfig() cpuCollectorConfig {
	cfg := new(cpuCollectorConfig)
	envconfig.GetViperConfig().Unmarshal(cfg)
	cfg.StatsConfig = agentconfiguration.GlobalCfg()
	return *cfg
}

var cpuStatCfg = initConfig()
