package storage

import (
	"context"
	"github.com/pkg/errors"
)

type CacheStoreage struct {
	memStorage   *MemStorage
	redisStorage *RedisStorage
	dbStorage    DBStorage
	codec        Codec
}

func NewCacheStoreage(codec Codec, memCache *MemStorage, redisCache *RedisStorage, database DBStorage) *CacheStoreage {
	return &CacheStoreage{
		codec:        codec,
		memStorage:   memCache,
		redisStorage: redisCache,
		dbStorage:    database,
	}
}

func (s *CacheStoreage) Set(ctx context.Context, key uint64, value any) error {
	newKey := s.codec.FormatKey(key)
	bytes, err := s.codec.Marshal(value)
	if err != nil {
		return err
	}
	if len(bytes) == 0 {
		return errors.New("codec marshal failed: empty result")
	}

	err = s.redisStorage.Set(ctx, newKey, bytes)
	if err != nil {
		return err
	}
	if s.memStorage != nil {
		err = s.memStorage.Set(ctx, newKey, bytes)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *CacheStoreage) Get(ctx context.Context, key uint64) (any, error) {
	newKey := s.codec.FormatKey(key)
	var value []byte
	var err error
	if s.memStorage != nil {
		value, err = s.memStorage.Get(ctx, newKey)
		if err == nil && len(value) > 0 {
			res, err := s.codec.Unmarshal(value)
			return res, err
		}
	}
	value, err = s.redisStorage.Get(ctx, newKey)
	if err == nil && len(value) > 0 {
		if s.memStorage != nil {
			s.memStorage.Set(ctx, newKey, value)
		}
		res, err := s.codec.Unmarshal(value)
		return res, err
	}
	model, err := s.dbStorage.Get(ctx, key)
	if err != nil {
		return nil, errors.Wrap(err, "CacheStoreage.dbStorage.Get failed")
	}
	s.Set(ctx, key, model)
	return model, nil
}

func (s *CacheStoreage) DeleteCache(ctx context.Context, key uint64) error {
	newKey := s.codec.FormatKey(key)
	err := s.redisStorage.Delete(ctx, newKey)
	if err != nil {
		return err
	}

	if s.memStorage != nil {
		err = s.memStorage.Delete(ctx, newKey)
		if err != nil {
			return err
		}
	}
	return nil
}
