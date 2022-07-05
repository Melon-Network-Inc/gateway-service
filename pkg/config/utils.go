package config

import (
	"github.com/spf13/viper"
)

type redisConfig struct {
	host     string
	port     string
	password string
	db       int
}

func (config *redisConfig) GetPassword() string {
	return config.password
}

func (config *redisConfig) GetHost() string {
	return config.host
}

func (config *redisConfig) GetPort() string {
	return config.port
}

func (config *redisConfig) GetDB() int {
	return config.db
}

type generalConfig struct {
	*viper.Viper
}

func (config *generalConfig) getRedisConfig() *redisConfig {
	host := "localhost"
	port := "6379"
	password := ""
	db := 0

	return &redisConfig{
		host:     host,
		port:     port,
		password: password,
		db:       db,
	}
}
