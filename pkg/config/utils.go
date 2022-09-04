package config

import (
	"time"

	"github.com/spf13/viper"
)

type generalConfig struct {
	*viper.Viper
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
