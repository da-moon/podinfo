// Package encoding is used to inject certain errors
// into underlying physical requests. It also supports
// transactional physical requests
package encoding

import physical "github.com/da-moon/northern-labs-interview/sdk/physical"

// Verify StorageEncoding satisfies the correct interfaces
var _ physical.Backend = (*StorageEncoding)(nil)
var _ physical.Transactional = (*NewTransactional)(nil)
