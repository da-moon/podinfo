package access

import (
	"context"

	physical "github.com/da-moon/northern-labs-interview/sdk/physical"
)

// PhysicalAccess is a wrapper struct
// to satisfy around physical.Backend
// physical.ToggleablePurgemonster interfaces
type PhysicalAccess struct {
	backend physical.Backend
}

// New returns a new instance of the physical access
func New(physical physical.Backend) *PhysicalAccess {
	return &PhysicalAccess{backend: physical}
}

// Put stores the given entry.
func (p *PhysicalAccess) Put(ctx context.Context, entry *physical.Entry) error {
	return p.backend.Put(ctx, entry)
}

// Get retrieves the entry at the given key.
func (p *PhysicalAccess) Get(ctx context.Context, key string) (*physical.Entry, error) {
	return p.backend.Get(ctx, key)
}

// Delete removes the entry at the given key.
func (p *PhysicalAccess) Delete(ctx context.Context, key string) error {
	return p.backend.Delete(ctx, key)
}

// List returns a list of keys under the given
func (p *PhysicalAccess) List(ctx context.Context, prefix string) ([]string, error) {
	return p.backend.List(ctx, prefix)
}

// Purge is a wrapper around physical.ToggleablePurgemonster.Purge
func (p *PhysicalAccess) Purge(ctx context.Context) {
	if purgeable, ok := p.backend.(physical.ToggleablePurgemonster); ok {
		purgeable.Purge(ctx)
	}
}
