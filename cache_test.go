package gcache

import (
	"context"
	"log"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type CacheMockup struct {
	data map[string]interface{}
	mu   *sync.Mutex
}

func NewCacheMockup() *CacheMockup {
	return &CacheMockup{
		data: make(map[string]interface{}),
		mu:   &sync.Mutex{},
	}
}

func (s *CacheMockup) Set(ctx context.Context, key string, value interface{}, ttl time.Duration) error {
	s.mu.Lock()
	s.data[key] = value
	s.mu.Unlock()

	return nil
}

func (s *CacheMockup) Get(ctx context.Context, key string) (interface{}, error) {
	s.mu.Lock()
	val := s.data[key]
	s.mu.Unlock()
	return val, nil
}

func (s *CacheMockup) Del(ctx context.Context, key string) error {
	delete(s.data, key)
	return nil
}

func TestCache(t *testing.T) {

	var (
		cache     *Cache
		KeyTest   = "keytest"
		ValueTest = "valuetest"
	)

	type JsonTest struct {
		Value string `json:"value"`
	}

	cache = New(NewCacheMockup())
	t.Run("set", func(t *testing.T) {
		err := cache.Set(KeyTest, ValueTest)
		assert.NoError(t, err)
	})

	t.Run("get", func(t *testing.T) {

		value, err := cache.Get(KeyTest)

		assert.NoError(t, err)

		assert.Equal(t, ValueTest, value)
	})

	t.Run("del", func(t *testing.T) {
		err := cache.Del(KeyTest)

		assert.NoError(t, err)
	})

	t.Run("get empty value", func(t *testing.T) {
		value, _ := cache.Get(KeyTest)
		assert.Empty(t, value)
	})

	t.Run("set json", func(t *testing.T) {
		err := cache.Set(KeyTest, JsonTest{
			Value: ValueTest,
		})
		assert.NoError(t, err)
	})

	t.Run("get json", func(t *testing.T) {
		val, err := cache.Get(KeyTest)

		assert.NoError(t, err)

		v, ok := val.(JsonTest)

		if !ok {
			log.Fatalln("casting failed")
		}

		assert.Equal(t, ValueTest, v.Value)
	})
}
