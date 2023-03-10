package gcache

import (
	"context"
	"time"
)

type (
	CacheClient interface {
		Set(ctx context.Context, key string, value interface{}, ttl time.Duration) error
		Get(ctx context.Context, key string) (interface{}, error)
		Del(ctx context.Context, key string) error
	}

	Cache struct {
		client CacheClient
	}
)

func New(client CacheClient) *Cache {

	manager := &Cache{
		client: client,
	}

	return manager
}

func (s *Cache) Set(key string, value interface{}) error {
	return s.client.Set(context.TODO(), key, value, time.Duration(1)*time.Hour)
}

func (s *Cache) SetNX(ctx context.Context, key string, value interface{}, ttl time.Duration) error {
	return s.client.Set(ctx, key, value, ttl)
}

func (s *Cache) Get(key string) (interface{}, error) {
	return s.client.Get(context.TODO(), key)
}

func (s *Cache) Del(key string) error {
	return s.client.Del(context.TODO(), key)
}

func (s *Cache) SetX(ctx context.Context, key string, value interface{}) error {
	return s.client.Set(ctx, key, value, time.Duration(1)*time.Hour)
}

func (s *Cache) GetX(ctx context.Context, key string) (interface{}, error) {
	return s.client.Get(ctx, key)
}

func (s *Cache) DelX(ctx context.Context, key string) error {
	return s.client.Del(ctx, key)
}
