// Package errorinjector helps with error injection into
// underlying physical requests . It also supports transactional
// physical requests
package errorinjector

import "github.com/da-moon/podinfo/sdk/physical"

// Verify ErrorInjector satisfies the correct interfaces
var _ physical.Backend = (*ErrorInjector)(nil)
var _ physical.Transactional = (*TransactionalErrorInjector)(nil)
