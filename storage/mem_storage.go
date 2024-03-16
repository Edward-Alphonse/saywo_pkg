package storage

import (
	"context"
	"time"

	"github.com/allegro/bigcache/v3"
)

type MemStorage struct {
	cache          *bigcache.BigCache
	expireDuration time.Duration
}

func NewMemStorage(ctx context.Context, expireDuration time.Duration) (*MemStorage, error) {
	config := bigcache.DefaultConfig(expireDuration)
	config.MaxEntriesInWindow = 6000
	cache, err := bigcache.New(ctx, config)
	if err != nil {
		return nil, err
	}
	return &MemStorage{
		cache:          cache,
		expireDuration: expireDuration,
	}, nil
}

func (c *MemStorage) Set(ctx context.Context, key string, value []byte) error {
	err := c.cache.Set(key, value)
	return err
}

func (c *MemStorage) Get(ctx context.Context, key string) ([]byte, error) {
	value, err := c.cache.Get(key)
	return value, err
}

func (c *MemStorage) Delete(ctx context.Context, key string) error {
	err := c.cache.Delete(key)
	if err == bigcache.ErrEntryNotFound {
		return nil
	}
	return err
}
