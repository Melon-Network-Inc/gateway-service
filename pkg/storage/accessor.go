package storage

import (
	"context"

	"github.com/Melon-Network-Inc/common/pkg/entity"
)

type usersDataAccessor interface {
	GetCachedUserByToken(ctx context.Context, userID string) (entity.CachedUser, error)
}

// Accessor is interface which defines all functions used on database/cache/fts
// in the system. Any other packages should use this Accessor instead of using
// eg. database directly.
type Accessor interface {
	usersDataAccessor
}
