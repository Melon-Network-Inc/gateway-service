package config

import (
	"github.com/spf13/viper"
)

// TokenConfiguration keeps all config providers
type TokenConfiguration struct {
	Token      TokenConfigProvider
}

func NewTokenConfig(path string) *TokenConfiguration {
	v := viper.New()

	v.SetConfigFile(path)
	v.SetConfigType("yml")

	config := &generalConfig{v}
	return &TokenConfiguration{
		Token:		config.getTokenConfig(),
	}
}
