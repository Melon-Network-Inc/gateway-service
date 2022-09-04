package config

import (
	"github.com/spf13/viper"
)

// TokenConfiguration keeps all config providers
type TokenConfiguration struct {
	Token      TokenConfigProvider
}

func NewTokenConfig() *TokenConfiguration {
	v := viper.New()

	config := &generalConfig{v}
	return &TokenConfiguration{
		Token:		config.getTokenConfig(),
	}
}
