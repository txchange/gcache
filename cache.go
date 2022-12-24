package trem

import (
	"context"
)

type (
	CacheClient interface {
		SetCache(ctx context.Context, key string, value interface{}) error
		GetCache(ctx context.Context, key string) (interface{}, error)
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

func (s Cache) SetCache(key string, value interface{}) error {
	return s.client.SetCache(context.TODO(), key, value)
}

func (s Cache) GetCache(key string) (interface{}, error) {
	return s.client.GetCache(context.TODO(), key)
}
