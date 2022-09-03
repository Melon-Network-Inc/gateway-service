package storage

import (
	"context"

	"github.com/Melon-Network-Inc/gateway-service/pkg/storage/cache"

	"github.com/Melon-Network-Inc/gateway-service/pkg/config"
)

type storage struct{
	usersDataAccessor
}

// New constructs Accessor
func New(ctx context.Context, redisConfig config.RedisConfigProvider) Accessor {
	redisCache := cache.NewRedisCache(ctx, redisConfig)
	if redisCache == nil {
		panic("failed to connect to Redis instance")
	}
	usersStorage := newUserStorage(redisCache)
	return &storage{
		usersDataAccessor:  usersStorage,
	}
}
