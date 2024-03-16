package storage

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisStorage struct {
	cache          *redis.Client
	expireDuration time.Duration
}

func NewRedisCache(redisClient *redis.Client, expireDuration time.Duration) *RedisStorage {
	return &RedisStorage{
		cache:          redisClient,
		expireDuration: expireDuration,
	}
}

func (c *RedisStorage) Set(ctx context.Context, key string, value []byte) error {
	err := c.cache.Set(ctx, key, value, c.expireDuration).Err()
	return err
}

func (c *RedisStorage) Get(ctx context.Context, key string) ([]byte, error) {
	result, err := c.cache.Get(ctx, key).Bytes()
	return result, err
}

func (c *RedisStorage) Delete(ctx context.Context, key string) error {
	err := c.cache.Del(ctx, key).Err()
	if err == redis.Nil {
		return nil
	}
	return err
}
