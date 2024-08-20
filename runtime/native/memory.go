package native

import (
	"sync"
	"time"

	"github.com/ysugimoto/vintage"
)

func newInMemoryCache() vintage.CacheDriver {
	return &InMemoryCache{
		store: sync.Map{},
	}
}

type inMemoryCacheItem struct {
	Expire time.Time
	Data   []byte
}

type InMemoryCache struct {
	store sync.Map
}

func (c *InMemoryCache) Get(key string) ([]byte, error) {
	v, ok := c.store.Load(key)
	if !ok {
		return nil, nil
	}
	item, ok := v.(inMemoryCacheItem)
	if !ok {
		return nil, nil
	}
	if item.Expire.After(time.Now()) {
		c.store.Delete(key)
		return nil, nil
	}
	return item.Data, nil
}

func (c *InMemoryCache) Set(key string, data []byte, ttl time.Duration) error {
	item := inMemoryCacheItem{
		Expire: time.Now().Add(ttl),
		Data:   data,
	}
	c.store.Store(key, item)
	return nil
}
