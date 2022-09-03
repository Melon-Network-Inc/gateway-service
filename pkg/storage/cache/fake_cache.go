package cache

import "context"

type fakeCache struct{}

// NewFakeCache creates new instance of fake cache which imitates caching values.
// Implements all Accessor methods but each of them is doing nothing.
func NewFakeCache() Accessor {
	return &fakeCache{}
}

func (cache *fakeCache) Set(ctx context.Context, entries ...Entry) error {
	return nil
}

func (cache *fakeCache) GetSingle(ctx context.Context, key Key, value Value) (bool, error) {
	return false, nil
}

func (cache *fakeCache) Get(ctx context.Context, entries ...Entry) ([]bool, error) {
	results := make([]bool, len(entries))
	for i := range results {
		results[i] = false
	}

	return results, nil
}

func (cache *fakeCache) Delete(ctx context.Context, keys ...string) error {
	return nil
}

func (cache *fakeCache) Incr(ctx context.Context, keys ...string) error {
	return nil
}

func (cache *fakeCache) Decr(ctx context.Context, keys ...string) error {
	return nil
}

func (cache *fakeCache) SAdd(ctx context.Context, key string, values Values) error {
	return nil
}

func (cache *fakeCache) SMembers(ctx context.Context, key string, values Values) (bool, error) {
	return false, nil
}

func (cache *fakeCache) SRemove(ctx context.Context, key string, values Values) error {
	return nil
}

func (cache *fakeCache) Flush(ctx context.Context) error {
	return nil
}
