package storage

import (
	"context"

	"github.com/Melon-Network-Inc/gateway-service/pkg/storage/cache"

	"github.com/Melon-Network-Inc/gateway-service/pkg/config"
)

type storage struct{}

// New constructs Accessor
func New(ctx context.Context, redisConfig config.RedisConfigProvider) {
	cache := cache.NewRedisCache(ctx, redisConfig)
	if cache == nil {
		panic("failed to connect to Redis instance")
	}
}
