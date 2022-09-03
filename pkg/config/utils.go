package config

import (
	"time"

	"github.com/spf13/viper"
)

type generalConfig struct {
	*viper.Viper
}

func (config *generalConfig) getPort() int {
	return 8080
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

func (config *generalConfig) getTokenConfig() *tokenConfig {
	secretKey := []byte("secret")
	authTokenValidityPeriod := time.Hour * 24
	refreshTokenValidityPeriod := time.Hour * 24

	return &tokenConfig{
		secretKey:                  secretKey,
		authTokenValidityPeriod:    authTokenValidityPeriod,
		refreshTokenValidityPeriod: refreshTokenValidityPeriod,
	}
}

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

type tokenConfig struct {
	secretKey                  []byte
	authTokenValidityPeriod    time.Duration
	refreshTokenValidityPeriod time.Duration
}

func (config *tokenConfig) GetSecretKey() []byte {
	return config.secretKey
}

func (config *tokenConfig) GetAuthTokenValidityPeriod() time.Duration {
	return config.authTokenValidityPeriod
}

func (config *tokenConfig) GetRefreshTokenValidityPeriod() time.Duration {
	return config.refreshTokenValidityPeriod
}
