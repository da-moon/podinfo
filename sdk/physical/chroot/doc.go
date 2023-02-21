// Package chroot helps with representig
// a prefixed view ( chroot ) of a physical backend
// this ensures program or users cannot access data they are
// not supposed to access
package chroot

import "github.com/da-moon/podinfo/sdk/physical"

// Verify View satisfies the correct interfaces
var _ physical.Backend = (*View)(nil)
