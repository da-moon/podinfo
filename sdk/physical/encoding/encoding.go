package encoding

import (
	"context"
	"unicode/utf8"

	physical "github.com/da-moon/podinfo/sdk/physical"
	"github.com/palantir/stacktrace"
)

// StorageEncoding is used to add errors into underlying physical requests
type StorageEncoding struct {
	physical.Backend
}

// New returns a wrapped physical backend and verifies the key
// encoding
func New(b physical.Backend) physical.Backend {
	enc := &StorageEncoding{
		Backend: b,
	}
	if bTxn, ok := b.(physical.Transactional); ok {
		return &NewTransactional{
			StorageEncoding: enc,
			Transactional:   bTxn,
		}
	}
	return enc
}

// Put stores the given entry
func (e *StorageEncoding) Put(ctx context.Context, entry *physical.Entry) error {
	if !utf8.ValidString(entry.Key) {
		errMsg := physical.NonUTF8ErrorMessage
		err := stacktrace.NewError(errMsg)
		return err
	}
	if e.containsNonPrintableChars(entry.Key) {
		errMsg := physical.NonPrintableErrorMessage
		err := stacktrace.NewError(errMsg)
		return err
	}
	return e.Backend.Put(ctx, entry)
}

// Delete removes the entry at the given key
func (e *StorageEncoding) Delete(ctx context.Context, key string) error {
	if !utf8.ValidString(key) {
		errMsg := physical.NonUTF8ErrorMessage
		err := stacktrace.NewError(errMsg)
		return err
	}
	if e.containsNonPrintableChars(key) {
		errMsg := physical.NonPrintableErrorMessage
		err := stacktrace.NewError(errMsg)
		return err
	}
	return e.Backend.Delete(ctx, key)
}
