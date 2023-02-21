package encoding

import (
	"context"

	physical "github.com/da-moon/podinfo/sdk/physical"
)

// Purge is used to purge the underlying backend.
func (e *StorageEncoding) Purge(ctx context.Context) {
	if purgeable, ok := e.Backend.(physical.ToggleablePurgemonster); ok {
		purgeable.Purge(ctx)
	}
}

// SetEnabled enables or disables the purgeable backend.
func (e *StorageEncoding) SetEnabled(enabled bool) {
	if purgeable, ok := e.Backend.(physical.ToggleablePurgemonster); ok {
		purgeable.SetEnabled(enabled)
	}
}
