package storage

import "github.com/Melon-Network-Inc/gateway-service/pkg/storage/cache"

// FakeStorage exports all fields which can be necessary for tests.
type FakeStorage struct {
	Cache   cache.Accessor
	Storage Accessor
}

// NewFakeStorage creates new fake storage necessary for tests.
func NewFakeStorage() *FakeStorage {
	cache := cache.NewFakeCache()
	return &FakeStorage{
		Cache:   cache,
		Storage: &storage{},
	}
}
