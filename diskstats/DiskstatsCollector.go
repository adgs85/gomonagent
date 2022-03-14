package diskstats

import (
	"github.com/adgs85/gomonagent/agentmessagesdispatcher"
	"github.com/adgs85/gomonmarshalling/monmarshalling/envconfig"
)

type DiskCollectorConfig struct {
	envconfig.StatsConfig
	DiskPollingRateMs int    `mapstructure:"disk_polling_rate_ms"`
	DiskFreeSpacePath string `mapstructure:"disk_free_space_path"`
}

func initConfig() DiskCollectorConfig {
	cfg := new(DiskCollectorConfig)
	envconfig.GetViperConfig().Unmarshal(cfg)
	envconfig.GetViperConfig().Unmarshal(&cfg.StatsConfig)
	return *cfg
}

var cfg DiskCollectorConfig = initConfig()

func New(sink agentmessagesdispatcher.StatSinkFuncType) {
	path := cfg.DiskFreeSpacePath
	CollectDiskInfo(path, sink)
}
