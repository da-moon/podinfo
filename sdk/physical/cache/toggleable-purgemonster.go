package cache

import (
	"context"
	"sync/atomic"
)

// SetEnabled is used to toggle whether the cache is on or off. It must be
// called with true to actually activate the cache after creation.
func (c *Cache) SetEnabled(enabled bool) {
	if enabled {
		atomic.StoreUint32(c.enabled, 1)
		return
	}
	atomic.StoreUint32(c.enabled, 0)
}

// Purge is used to clear the cache
func (c *Cache) Purge(ctx context.Context) {
	// Lock the world
	for _, lock := range c.locks {
		lock.Lock()
		defer lock.Unlock()
	}
	c.lru.Purge()
}
