// Package latency is used to introduce latency
// into underlying physical requests. it also supports transactional physical requests
package latency

import "github.com/da-moon/northern-labs-interview/sdk/physical"

// Verify Injector satisfies the correct interfaces
var _ physical.Backend = (*Injector)(nil)
var _ physical.Transactional = (*TransactionalLatencyInjector)(nil)
