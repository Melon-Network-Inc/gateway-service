package config

import "time"

// CacheConfigProvider provides general Cache access configuration.
type CacheConfigProvider interface {
	GetPassword() string
	GetHost() string
	GetPort() string
	GetExpirationTime() time.Duration
}

// RedisConfigProvider provides Redis access configuration.
type RedisConfigProvider interface {
	CacheConfigProvider
	GetDB() int
}
