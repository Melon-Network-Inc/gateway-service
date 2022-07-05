package config

import (
	"github.com/spf13/viper"
)

// Configuration keeps all config providers
type Configuration struct {
	Redis RedisConfigProvider
}

// New reads and creates configuration from path provided in env `$MELON_CONFIG_PATH`
// or from default path `$GOPATH/src/github.com/Melon-Network-Inc/gateway-service`.
//
// By setting `$MELON_CONFIG_NAME` variable you can specify the name
// of config file. Default name is `config`.
//
// By setting `$MELON_CONFIG_TYPE` variable you can specify which config type will
// be chosen: `development`, `production` or `test`. Default is `development`.
func New() *Configuration {
	v := viper.New()

	config := &generalConfig{v}
	return &Configuration{
		Redis: config.getRedisConfig(),
	}
}
