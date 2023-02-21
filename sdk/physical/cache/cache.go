package cache

import (
	"context"

	lru "github.com/da-moon/podinfo/internal/golang-lru"
	"github.com/da-moon/podinfo/internal/locksutil"
	"github.com/da-moon/podinfo/internal/pathmanager"
	physical "github.com/da-moon/podinfo/sdk/physical"
)

const (
	// DefaultCacheSize is used if no cache size is specified for NewCache
	// 128 KB
	defaultCacheSize = 128 * 1024
	// refreshCacheCtxKey is a ctx value that denotes the cache should be
	// refreshed during a Get call.
	refreshCacheCtxKey = "refresh_cache"
)

// Cache is used to wrap an underlying physical backend
// and provide an LRU cache layer on top.
type Cache struct {
	backend         physical.Backend
	lru             *lru.TwoQueueCache
	locks           []*locksutil.LockEntry
	enabled         *uint32
	cacheExceptions *pathmanager.PathManager
}

// New returns a physical cache of the given size.
// If no size is provided, the default size is used.
func New(b physical.Backend, pm *pathmanager.PathManager, size int) *Cache {
	if size <= 0 {
		size = defaultCacheSize
	}
	cache, _ := lru.New2Q(size)
	c := &Cache{
		backend:         b,
		lru:             cache,
		locks:           locksutil.CreateLocks(),
		enabled:         new(uint32),
		cacheExceptions: pm,
	}
	return c
}

// Put stores the entry in the cache and the underlying backend.
func (c *Cache) Put(ctx context.Context, entry *physical.Entry) error {
	if entry != nil && !c.ShouldCache(entry.Key) {
		return c.backend.Put(ctx, entry)
	}
	lock := locksutil.LockForKey(c.locks, entry.Key)
	lock.Lock()
	defer lock.Unlock()
	err := c.backend.Put(ctx, entry)
	if err == nil {
		c.lru.Add(entry.Key, entry)
	}
	return err
}

// Get returns the entry from the cache if it is present.
func (c *Cache) Get(ctx context.Context, key string) (*physical.Entry, error) {
	if !c.ShouldCache(key) {
		return c.backend.Get(ctx, key)
	}
	lock := locksutil.LockForKey(c.locks, key)
	lock.RLock()
	defer lock.RUnlock()
	// Check the LRU first
	if !cacheRefreshFromContext(ctx) {
		if raw, ok := c.lru.Get(key); ok {
			if raw == nil {
				return nil, nil
			}
			return raw.(*physical.Entry), nil
		}
	}
	// Read from the underlying backend
	ent, err := c.backend.Get(ctx, key)
	if err != nil {
		return nil, err
	}
	// Cache the result
	c.lru.Add(key, ent)
	return ent, nil
}

// Delete deletes the entry from the cache and the underlying backend.
func (c *Cache) Delete(ctx context.Context, key string) error {
	if !c.ShouldCache(key) {
		return c.backend.Delete(ctx, key)
	}
	lock := locksutil.LockForKey(c.locks, key)
	lock.Lock()
	defer lock.Unlock()
	err := c.backend.Delete(ctx, key)
	if err == nil {
		c.lru.Remove(key)
	}
	return err
}

// List returns the keys from the underlying backend.
func (c *Cache) List(ctx context.Context, prefix string) ([]string, error) {
	// Always pass-through as this would be difficult to cache. For the same
	// reason we don't lock as we can't reasonably know which locks to readlock
	// ahead of time.
	return c.backend.List(ctx, prefix)
}

// Locks returns the locks for the cache.
func (c *TransactionalCache) Locks() []*locksutil.LockEntry {
	return c.locks
}

// LRU returns the LRU cache for the cache.
func (c *TransactionalCache) LRU() *lru.TwoQueueCache {
	return c.lru
}
