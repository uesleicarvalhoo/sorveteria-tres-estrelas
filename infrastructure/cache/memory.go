package cache

import (
	"context"
	"fmt"
	"time"

	"github.com/dgraph-io/ristretto"
	"github.com/pkg/errors"
)

const (
	numCounters = 1e7     // number of keys to track frequency of (10M).
	maxCost     = 1 << 30 // maximum cost of cache (1GB).
	bufferItems = 64      // number of keys per Get buffer.
)

type MemoryCache struct {
	Client *ristretto.Cache
}

func NewMemoryCacheClient() (*MemoryCache, error) {
	client, err := ristretto.NewCache(&ristretto.Config{
		NumCounters: numCounters, // number of keys to track frequency of (10M).
		MaxCost:     maxCost,     // maximum cost of cache (1GB).
		BufferItems: bufferItems, // number of keys per Get buffer.
	})

	cache := &MemoryCache{
		Client: client,
	}

	return cache, err
}

func (cache *MemoryCache) Set(ctx context.Context, key, value string, expiration time.Duration) error {
	ok := cache.Client.SetWithTTL(key, value, 0, expiration)
	cache.Client.Wait()

	if !ok {
		return errors.Wrap(ErrCache, "Error on set")
	}

	return nil
}

func (cache *MemoryCache) HSet(ctx context.Context, key string, values ...any) error {
	return nil
}

func (cache *MemoryCache) HSetExp(ctx context.Context, key string, expiration time.Duration, values ...any) error {
	return nil
}

func (cache *MemoryCache) Get(ctx context.Context, key string) (string, error) {
	data, ok := cache.Client.Get(key)
	if !ok {
		return "", errors.Wrap(ErrCache, "Error on get")
	}

	return fmt.Sprintf("%s", data), nil
}

func (cache *MemoryCache) Del(ctx context.Context, key string) error {
	cache.Client.Del(key)

	return nil
}
