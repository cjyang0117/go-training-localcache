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
	timer  *time.Timer
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

	if c.timer != nil && c.data[key] != nil {
		c.timer.Stop()
	}

	c.data[key] = data
	c.timer = time.AfterFunc(c.expiry, func() {
		c.m.Lock()
		defer c.m.Unlock()

		delete(c.data, key)
	})

	return nil
}
