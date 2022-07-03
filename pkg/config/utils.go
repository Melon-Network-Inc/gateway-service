package config

import (
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type redisConfig struct {
	password       string
	host           string
	port           string
	db             int
	expirationTime time.Duration
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

func (config *redisConfig) GetExpirationTime() time.Duration {
	return config.expirationTime
}

type generalConfig struct {
	*viper.Viper
}

func (config *generalConfig) getRedisConfig() *redisConfig {
	password := config.GetString("redis.password")
	host := config.GetString("redis.host")
	port := config.GetString("redis.port")
	db := config.GetInt("redis.db")
	expirationTime := config.GetDuration("redis.expiration_time")

	if host == "" || port == "" || db < 0 || expirationTime < 0 {
		log.WithFields(log.Fields{
			"password":        password,
			"host":            host,
			"port":            port,
			"db":              db,
			"expiration time": expirationTime,
		}).Fatal("Config file doesn't contain valid redis access data.")
	}

	return &redisConfig{
		password:       password,
		host:           host,
		port:           port,
		db:             db,
		expirationTime: expirationTime,
	}
}
