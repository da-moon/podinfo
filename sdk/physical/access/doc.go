// Package access allows a service to expose its physical
// storage operations through PhysicalAccess() while
// restricting the ability to modify Service.physical itself.
// It should be used when returning a `physical` field
// of a parent to wrap the field for security purposes
package access

import "github.com/da-moon/podinfo/sdk/physical"

// Verify PhysicalAccess satisfies the correct interfaces
var _ physical.Backend = (*PhysicalAccess)(nil)
