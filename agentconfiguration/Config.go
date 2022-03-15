package agentconfiguration

import (
	"github.com/adgs85/gomonmarshalling/monmarshalling/envconfig"
)

var globalCfg = initCfg()

func GlobalCfg() *envconfig.StatsConfig {
	return globalCfg
}

func initCfg() *envconfig.StatsConfig {
	cfg := new(envconfig.StatsConfig)
	envconfig.GetViperConfig().Unmarshal(cfg)
	return cfg
}
