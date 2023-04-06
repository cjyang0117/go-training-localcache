package localcache

import (
	"sync"
	"time"
)

// New initializes a new cache instance
func New(expiry time.Duration) Cache {
	return &cache{
		expiry: expiry,
		data:   make(map[string]interface{}),
	}
}

type cache struct {
	m      sync.Mutex
	expiry time.Duration
	data   map[string]interface{}
}

func (c *cache) Get(key string) (interface{}, error) {
	c.m.Lock()
	defer c.m.Unlock()

	if data, ok := c.data[key]; ok {
		return data, nil
	}

	return nil, ErrCacheMiss
}

func (c *cache) Set(key string, data interface{}) error {
	c.m.Lock()
	defer c.m.Unlock()

	c.data[key] = data
	time.AfterFunc(c.expiry*time.Second, func() {
		c.m.Lock()
		defer c.m.Unlock()

		delete(c.data, key)
	})

	return nil
}
