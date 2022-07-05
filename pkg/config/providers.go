package config

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
