package config

import (
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/Melon-Network-Inc/gateway-service/pkg/utils"
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
	if cp := os.Getenv(`MELON_CONFIG_PATH`); cp != "" {
		v.AddConfigPath(cp)
	}
	v.AddConfigPath(`$GOPATH/src/github.com/Melon-Network-Inc/gateway-service`)

	cn := utils.GetenvOrDefault(`MELON_CONFIG_NAME`, "config")
	v.SetConfigName(cn)

	err := v.ReadInConfig()
	if err != nil {
		log.WithError(err).Error("Error reading config file.")
		return nil
	}

	ct := utils.GetenvOrDefault(`MELON_CONFIG_TYPE`, "development")
	subConfig := v.Sub(ct)
	if subConfig == nil {
		log.Errorf("Failed to read '%s' config.", ct)
		return nil
	}

	config := &generalConfig{subConfig}
	return &Configuration{
		Redis: config.getRedisConfig(),
	}
}
