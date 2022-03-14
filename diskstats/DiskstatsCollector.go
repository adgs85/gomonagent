package diskstats

import (
	"fmt"

	"github.com/adgs85/gomonmarshalling/monmarshalling"
	"github.com/adgs85/gomonmarshalling/monmarshalling/envconfig"
	"github.com/davecgh/go-spew/spew"
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

func New() {
	path := cfg.DiskFreeSpacePath
	CollectDiskInfo(path, sink)
}

func sink(stats monmarshalling.Stat) {
	fmt.Println(spew.Sdump(stats))
}
