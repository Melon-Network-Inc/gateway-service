package config

import "time"

// TokenConfigProvider provides Token secret and validity period configuration.
type TokenConfigProvider interface {
	GetSecretKey() []byte
	GetAuthTokenValidityPeriod() time.Duration
	GetRefreshTokenValidityPeriod() time.Duration
}