package storage

import (
	"context"
	"gateway-service/pkg/config"
	"gateway-service/pkg/storage/cache"
)

type storage struct{}

// New constructs Accessor
func New(ctx context.Context, redisConfig config.RedisConfigProvider) {
	cache := cache.NewRedisCache(ctx, redisConfig)
	if cache == nil {
		panic("failed to connect to Redis instance")
	}
}
