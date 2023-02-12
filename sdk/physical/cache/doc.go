// Package cache provide an LRU cache layer on top of an
// underlying physical backend
package cache

import (
	physical "github.com/da-moon/northern-labs-interview/sdk/physical"
)

// Verify Cache satisfies the correct interfaces
var _ physical.ToggleablePurgemonster = (*Cache)(nil)
var _ physical.ToggleablePurgemonster = (*TransactionalCache)(nil)
var _ physical.Backend = (*Cache)(nil)
var _ physical.Transactional = (*TransactionalCache)(nil)
