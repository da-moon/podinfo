package cache

import (
	"context"
	"sync/atomic"
)

func (c *Cache) ShouldCache(key string) bool {
	if atomic.LoadUint32(c.enabled) == 0 {
		return false
	}
	return !c.cacheExceptions.HasPath(key)
}

// cacheRefreshFromContext is a helper to look up if the provided context is
// requesting a cache refresh.
func cacheRefreshFromContext(ctx context.Context) bool {
	r, ok := ctx.Value(refreshCacheCtxKey).(bool)
	if !ok {
		return false
	}
	return r
}

// RefreshContext returns a context with an added value denoting if the
// cache should attempt a refresh.
func RefreshContext(ctx context.Context, r bool) context.Context {
	//nolint:staticcheck // SA1029 allow using a string as a context key just because I am too lazy to create a new type
	return context.WithValue(ctx, refreshCacheCtxKey, r) // revive:disable:context-keys-type
	//lint:staticcheck
}
