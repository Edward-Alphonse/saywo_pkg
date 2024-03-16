package utils

import (
	"context"
	"encoding/json"
	"time"

	"github.com/Edward-Alphonse/saywo_pkg/logs"
	"github.com/redis/go-redis/v9"
)

type RedisObjCache[T any, U interface{ *T }] struct {
	ctx       context.Context
	redis     redis.UniversalClient
	redisKey  string
	cacheTime time.Duration      // 正常缓存时间
	queryFunc func() (*T, error) // 查询不到缓存，执行的闭包查询
}

func NewRedisObjCache[T any, U interface{ *T }](ctx context.Context, redis redis.UniversalClient,
	redisKey string, cacheTime time.Duration, queryFunc func() (*T, error)) *RedisObjCache[T, U] {
	return &RedisObjCache[T, U]{ctx: ctx, redis: redis, redisKey: redisKey, cacheTime: cacheTime, queryFunc: queryFunc}
}

func (ro RedisObjCache[T, U]) Get() (result *T, err error) {
	vCache, err := ro.redis.Get(ro.ctx, ro.redisKey).Bytes()
	if err != nil && err != redis.Nil {
		// redis 查询失败，则直接查询方法
		logs.InfoByArgs("RedisObjCache redis get failed :%v", err.Error())
		result, err = ro.queryFunc()
		return result, err
	}

	// 缓存为空
	if err == redis.Nil {
		// 执行查询
		result, err = ro.queryFunc()
		if err != nil {
			return result, err
		}

		cacheByte, err := json.Marshal(result)
		if err != nil {
			return result, err
		}

		// 缓存对象
		rsErr := ro.redis.SetEx(ro.ctx, ro.redisKey, cacheByte, ro.cacheTime).Err()
		if rsErr != nil {
			return result, rsErr
		}

		return result, err
	}

	t := new(T)
	err = json.Unmarshal(vCache, t)
	if err != nil {
		logs.InfoByArgs("RedisObjCache json unmarshal failed :%v", err.Error())
	}
	return t, nil
}
