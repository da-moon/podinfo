package chroot

import (
	"context"

	physical "github.com/da-moon/podinfo/sdk/physical"
)

// View represents a prefixed view of a physical backend
type View struct {
	backend physical.Backend
	prefix  string
}

// New takes an underlying physical backend and returns
// a view of it that can only operate with the given prefix.
func New(backend physical.Backend, prefix string) *View {
	return &View{
		backend: backend,
		prefix:  prefix,
	}
}

// List the contents of the prefixed view
func (v *View) List(ctx context.Context, prefix string) ([]string, error) {
	if err := v.sanityCheck(prefix); err != nil {
		return nil, err
	}
	return v.backend.List(ctx, v.expandKey(prefix))
}

// Get the key of the prefixed view
func (v *View) Get(ctx context.Context, key string) (*physical.Entry, error) {
	if err := v.sanityCheck(key); err != nil {
		return nil, err
	}
	entry, err := v.backend.Get(ctx, v.expandKey(key))
	if err != nil {
		return nil, err
	}
	if entry == nil {
		return nil, nil
	}
	if entry != nil {
		entry.Key = v.truncateKey(entry.Key)
	}
	return &physical.Entry{
		Key:   entry.Key,
		Value: entry.Value,
	}, nil
}

// Put the entry into the prefix view
func (v *View) Put(ctx context.Context, entry *physical.Entry) error {
	if err := v.sanityCheck(entry.Key); err != nil {
		return err
	}
	nested := &physical.Entry{
		Key:   v.expandKey(entry.Key),
		Value: entry.Value,
	}
	return v.backend.Put(ctx, nested)
}

// Delete the entry from the prefix view
func (v *View) Delete(ctx context.Context, key string) error {
	if err := v.sanityCheck(key); err != nil {
		return err
	}
	return v.backend.Delete(ctx, v.expandKey(key))
}
