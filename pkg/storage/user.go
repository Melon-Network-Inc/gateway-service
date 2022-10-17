package storage

import (
	"context"
	"errors"

	log "github.com/sirupsen/logrus"

	"github.com/Melon-Network-Inc/common/pkg/entity"
	"github.com/Melon-Network-Inc/gateway-service/pkg/storage/cache"
	"github.com/Melon-Network-Inc/gateway-service/pkg/token"
)

type usersStorage struct {
	cache    cache.Accessor
	tokenMgr token.HashedTokenManager
}

func newUserStorage(cache cache.Accessor, tokenMgr token.HashedTokenManager) usersDataAccessor {
	return &usersStorage{
		cache:    cache,
		tokenMgr: tokenMgr,
	}
}

func (s *usersStorage) GetCachedUserByToken(ctx context.Context, token string) (entity.CachedUser, error) {
	var user entity.CachedUser

	userID, _, err := s.tokenMgr.ValidateAuthToken(token)
	if err != nil {
		log.Error("unable to validate token")
		return entity.CachedUser{}, err
	}

	key := cache.Key{"user", userID}
	exists, err := s.cache.GetSingle(ctx, key, &user)
	if err != nil {
		return entity.CachedUser{}, err
	}
	if !exists {
		log.Info("user information is not present in cache")
		return entity.CachedUser{}, errors.New("user information is not present in cache")
	}
	return user, nil
}
