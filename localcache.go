// Package localcache provides a local cache.
package localcache

import (
	"errors"
)

var (
	// ErrCacheMiss indicates the key is missing
	ErrCacheMiss = errors.New("cache key is missing")
)

// Cache is used to set and get cached data
type Cache interface {
	// Get returns cached data by key
	Get(key string) (interface{}, error)
	// Set sets cached data by key
	Set(key string, data interface{}) error
}
