package util

import (
	"github.com/bixlabs/authentication/tools"
	"github.com/caarlos0/env"
)

var config GlobalConfig

func IsDevEnvironment() bool {
	return newConfig().Env == "dev"
}

type GlobalConfig struct {
   Env  string `env:"AUTH_SERVER_APP_ENV" envDefault:"dev"`
}

func newConfig() GlobalConfig {
	if config.Env != "" {
		return config
	}
	parseConfiguration()
	return config
}

func parseConfiguration() GlobalConfig {
	config = GlobalConfig{}
	err := env.Parse(&config)
	if err != nil {
		tools.Log().WithError(err).Panic("parsing the env variables for global configuration failed")
	}
	return config
}



