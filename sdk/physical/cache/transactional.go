package cache

import (
	"context"
	"sync/atomic"

	"github.com/da-moon/podinfo/internal/locksutil"
	"github.com/da-moon/podinfo/internal/pathmanager"
	physical "github.com/da-moon/podinfo/sdk/physical"
)

// TransactionalCache struct wraps the physical backend and adds transactional
// cache capabilities to it.
type TransactionalCache struct {
	*Cache
	Transactional physical.Transactional
}

// NewTransactional returns a new instance of the transactional cache.
func NewTransactional(b physical.Backend, pm *pathmanager.PathManager, size int) *TransactionalCache {
	c := &TransactionalCache{
		Cache:         New(b, pm, size),
		Transactional: b.(physical.Transactional),
	}
	return c
}

// Transaction runs a transactional (atomic) operation
// on the underlying physical backend and
func (c *TransactionalCache) Transaction(ctx context.Context, txns []*physical.TxnEntry) []physical.ErrTransactionalOperations {
	// Bypass the locking below
	if atomic.LoadUint32(c.enabled) == 0 {
		return c.Transactional.Transaction(ctx, txns)
	}
	// Collect keys that  need to be locked
	// var keys []string
	keys := make([]string, 0)
	for _, curr := range txns {
		keys = append(keys, curr.Entry.Key)
	}
	// Lock the keys
	for _, l := range locksutil.LocksForKeys(c.locks, keys) {
		l.Lock()
		defer l.Unlock()
	}
	if err := c.Transactional.Transaction(ctx, txns); err != nil {
		return err
	}
	for _, txn := range txns {
		if !c.ShouldCache(txn.Entry.Key) {
			continue
		}
		switch txn.Operation {
		case physical.PutOperation:
			c.lru.Add(txn.Entry.Key, txn.Entry)
		case physical.DeleteOperation:
			c.lru.Remove(txn.Entry.Key)
		}
	}
	return nil
}
