package diskstats

import (
//	"fmt"

	"github.com/adgs85/gomonmarshalling/monmarshalling/envconfig"
	"github.com/davecgh/go-spew/spew"
)

type diskCollectorConfig struct {
	envconfig.StatsConfig
	DiskPollingRateMs int `mapstructure:"disk_polling_rate_ms"`
	DiskFreeSpaceParg string `mapstructure:"disk_free_space_path"`
}

func initConfig() diskCollectorConfig {
	cfg := new(diskCollectorConfig)
	envconfig.GetViperConfig().Unmarshal(cfg)
	envconfig.GetViperConfig().Unmarshal(&cfg.StatsConfig)
	return *cfg
}

func New() {

	println("Disk stats config: \n" + spew.Sdump(initConfig()))
	
}
