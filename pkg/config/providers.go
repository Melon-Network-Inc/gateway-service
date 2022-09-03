package config

import "time"

// CacheConfigProvider provides general Cache access configuration.
type CacheConfigProvider interface {
	GetPassword() string
	GetHost() string
	GetPort() string
}

// RedisConfigProvider provides Redis access configuration.
type RedisConfigProvider interface {
	CacheConfigProvider
	GetDB() int
}

// TokenConfigProvider provides Token secret and validity period configuration.
type TokenConfigProvider interface {
	GetSecretKey() []byte
	GetAuthTokenValidityPeriod() time.Duration
	GetRefreshTokenValidityPeriod() time.Duration
}