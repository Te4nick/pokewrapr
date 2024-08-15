package cache

import (
	"context"
	"sync"
	"time"
)

const CacheDefaultTTL time.Duration = time.Minute * 5

type Cache struct {
	mutex   *sync.RWMutex
	content map[string]interface{}
}

func NewCache() (cache *Cache) {
	return &Cache{
		mutex:   &sync.RWMutex{},
		content: map[string]interface{}{},
	}
}

func (cache *Cache) SetDefault(key string, value interface{}) {
	cache.mutex.Lock()
	cache.content[key] = value
	cache.mutex.Unlock()
}

func (cache *Cache) Set(key string, value interface{}, ttl time.Duration) {
	cache.SetDefault(key, value)
	go cache.purge(key, ttl)
}

func (cache *Cache) Get(key string) interface{} {
	cache.mutex.RLock()
	defer cache.mutex.RUnlock()

	value, ok := cache.content[key]
	if !ok {
		return nil
	}

	return value
}

func (cache *Cache) purge(key string, ttl time.Duration) {
	ctx, cancel := context.WithTimeout(context.Background(), ttl)
	defer cancel()

	<-ctx.Done()

	cache.mutex.Lock()
	delete(cache.content, key)
	cache.mutex.Unlock()
}
