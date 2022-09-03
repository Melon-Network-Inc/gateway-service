package storage

import (
	"context"
	"errors"
	"github.com/Melon-Network-Inc/common/pkg/log"

	"github.com/Melon-Network-Inc/gateway-service/pkg/config"
	"github.com/Melon-Network-Inc/gateway-service/pkg/storage/cache"
	"github.com/Melon-Network-Inc/gateway-service/pkg/token"
)

type storage struct {
	usersDataAccessor
}

// New constructs Accessor
func New(ctx context.Context, redisConfig config.RedisConfigProvider, tokenConfig config.TokenConfigProvider, logger log.Logger) (Accessor, error) {
	redisCache := cache.NewRedisCache(ctx, redisConfig)
	if redisCache == nil {
		logger.Error("failed to connect to Redis instance")
		return &storage{}, errors.New("failed to connect to Redis instance")
	}

	tokenMgr := token.NewHashedTokenManager(tokenConfig, logger)
	if tokenConfig == nil {
		logger.Error("failed to create hashed token manager")
		return &storage{}, errors.New("failed to set up hashed token manager")
	}

	usersStorage := newUserStorage(redisCache, tokenMgr)
	return &storage{
		usersDataAccessor: usersStorage,
	}, nil
}
