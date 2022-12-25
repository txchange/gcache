package redis

import (
	"context"

	"github.com/go-redis/cache/v8"
	"github.com/go-redis/redis/v8"
)

type (
	RedisCache struct {
		client *cache.Cache
	}
)

func New(address string) *RedisCache {
	rdb := redis.NewClient(&redis.Options{
		Addr: address,
	})

	cacheManager := cache.New(&cache.Options{
		Redis: rdb,
	})

	redis := &RedisCache{
		client: cacheManager,
	}

	return redis
}

func (s *RedisCache) Set(ctx context.Context, key string, value interface{}) error {
	return s.client.Set(&cache.Item{
		Ctx:   ctx,
		Key:   key,
		Value: value,
	})
}

func (s *RedisCache) Get(ctx context.Context, key string) (interface{}, error) {
	var values interface{}
	err := s.client.Get(ctx, key, &values)
	return values, err
}

func (s *RedisCache) Del(ctx context.Context, key string) error {
	return s.client.Delete(ctx, key)
}
