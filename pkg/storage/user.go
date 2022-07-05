package storage

import (
	"context"

	log "github.com/sirupsen/logrus"

	"github.com/Melon-Network-Inc/gateway-service/pkg/entity"
	"github.com/Melon-Network-Inc/gateway-service/pkg/storage/cache"
)

type usersStorage struct {
	cache cache.Accessor
}

func newUserStorage(cache cache.Accessor) usersDataAccessor {
	return &usersStorage{
		cache: cache,
	}
}

func (s *usersStorage) GetUser(ctx context.Context, token string) (entity.User, error) {
	var user entity.User

	key := cache.Key{"token", token}
	if exists, err := s.cache.GetSingle(ctx, key, &user); !exists {
		log.Info("User not logged in")
		return entity.User{}, err
	}
	return user, nil
}
