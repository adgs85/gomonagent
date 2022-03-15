package diskstats

import (
	"github.com/adgs85/gomonmarshalling/monmarshalling/envconfig"
)

type DiskCollectorConfig struct {
	HostName          string `mapstructure:"host_name"`
	DiskPollingRateMs int    `mapstructure:"disk_polling_rate_ms"`
	DiskFreeSpacePath string `mapstructure:"disk_free_space_path"`
}

func initConfig() DiskCollectorConfig {
	cfg := new(DiskCollectorConfig)
	envconfig.GetViperConfig().Unmarshal(cfg)
	return *cfg
}

var cfg DiskCollectorConfig = initConfig()
