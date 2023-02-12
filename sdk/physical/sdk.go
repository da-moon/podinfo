package physical

import (
	"context"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"strings"
	"time"

	"github.com/palantir/stacktrace"
)

// constants
// TODO Make this dynamic ?
const (
	// DefaultParallelOperations - Default number of parallel operations
	DefaultParallelOperations = 128
	// DefaultJitterPercent is used if no cache size is specified for NewCache
	DefaultJitterPercent = 20
	// DefaultErrorPercent is used to determin how often we error
	DefaultErrorPercent = 20
)

// Entry is used to represent data stored by the physical backend
type Entry struct {
	Key   string
	Value []byte
}

// MD5CurrentHexString returns the MD5 hash of the current value
// wrapped in Entry struct.
func (e *Entry) MD5CurrentHexString() string {
	hash := md5.New()   //nolint:gosec // this is not a sensitive hash so using md5 is fine
	hash.Write(e.Value) //nolint:errcheck // it shouldn't cause any issues to ignore this error
	md5sumCurr := hash.Sum(nil)
	var appendHyphen bool
	if len(md5sumCurr) == 0 {
		md5sumCurr = make([]byte, 16)
		rand.Read(md5sumCurr) //nolint:errcheck // it shouldn't cause any issues to ignore this error
		appendHyphen = true
	}
	if appendHyphen {
		return hex.EncodeToString(md5sumCurr)[:32] + "-1"
	}
	return hex.EncodeToString(md5sumCurr)
}

// Backend is the interface required for a physical
// backend. A physical backend is used to durably store
// data outside of Vault. As such, it is completely untrusted,
// and is only accessed via a security barrier. The backends
// must represent keys in a hierarchical manner. All methods
// are expected to be thread safe.
type Backend interface {
	// Put is used to insert or update an entry
	Put(ctx context.Context, entry *Entry) error
	// Get is used to fetch an entry
	Get(ctx context.Context, key string) (*Entry, error)
	// Delete is used to permanently delete an entry
	Delete(ctx context.Context, key string) error
	// List is used to list all the keys under a given
	// prefix, up to the next prefix.
	List(ctx context.Context, prefix string) ([]string, error)
}

// ToggleablePurgemonster is an interface for backends that can toggle on or
// off special functionality and/or support purging. This is only used for the
// cache, don't use it for other things.
type ToggleablePurgemonster interface {
	Purge(ctx context.Context)
	SetEnabled(bool)
}

// Transactional is an optional interface for backends that
// support doing transactional updates of multiple keys. This is
// required for some features such as replication.
type Transactional interface {
	// The function to run a transaction
	Transaction(context.Context, []*TxnEntry) []ErrTransactionalOperations
}

// TransactionalBackend is an optional interface for backends that
// support doing transactional updates of multiple keys.
type TransactionalBackend interface {
	Backend
	Transactional
}

// PseudoTransactional is an optional interface for backends that
// support doing transactional updates of multiple keys.
// This interface is used to ensure that required functions for
// GenericTransactionHandler are implemented.
type PseudoTransactional interface {
	// An internal function should do no locking or permit pool acquisition.
	// Depending on the backend and if it natively supports transactions, these
	// may simply chain to the normal backend functions.
	GetInternal(context.Context, string) (*Entry, error)
	PutInternal(context.Context, *Entry) error
	DeleteInternal(context.Context, string) error
}

// Backoff used with retirable operations
type Backoff interface {
	NextInterval(retry int) time.Duration
}

// Operation enum represents a CRUD operation
type Operation string

// operation enum value
const (
	DeleteOperation Operation = "delete"
	GetOperation              = "get"
	ListOperation             = "list"
	PutOperation              = "put"
)

// TxnEntry is an operation that takes atomically as part of
// a transactional update. Only supported by Transactional backends.
type TxnEntry struct {
	Operation Operation
	Entry     Entry
}

// GenericTransactionHandler Implements the transaction interface
// func (t PseudoTransactional)GenericTransactionHandler(ctx context.Context txns []*TxnEntry) (retErr error) {
func GenericTransactionHandler(ctx context.Context, t PseudoTransactional, txns []*TxnEntry) []ErrTransactionalOperations {
	retErr := make([]ErrTransactionalOperations, 0)
	rollbackStack := make([]*TxnEntry, 0, len(txns))
	var dirty bool
	// We walk the transactions in order; each successful operation goes into a
	// LIFO for rollback if we hit an error along the way
TxnWalk:
	for _, txn := range txns {
		switch txn.Operation {
		case DeleteOperation:
			entry, err := t.GetInternal(ctx, txn.Entry.Key)
			if err != nil && !IsErrNotExist(err) {
				retErr = append(retErr, ErrTransactionalOperation(*txn, err.Error()))
				dirty = true
				break TxnWalk
			}
			if entry == nil {
				// Nothing to delete or roll back
				continue
			}
			rollbackEntry := &TxnEntry{
				Operation: PutOperation,
				Entry: Entry{
					Key:   entry.Key,
					Value: entry.Value,
				},
			}
			err = t.DeleteInternal(ctx, txn.Entry.Key)
			if err != nil {
				retErr = append(retErr, ErrTransactionalOperation(*txn, err.Error()))
				dirty = true
				break TxnWalk
			}
			rollbackStack = append([]*TxnEntry{rollbackEntry}, rollbackStack...)
		case PutOperation:
			entry, err := t.GetInternal(ctx, txn.Entry.Key)
			if err != nil && !IsErrNotExist(err) {
				retErr = append(retErr, ErrTransactionalOperation(*txn, err.Error()))
				dirty = true
				break TxnWalk
			}
			// Nothing existed so in fact rolling back requires a delete
			var rollbackEntry *TxnEntry
			if entry == nil {
				rollbackEntry = &TxnEntry{
					Operation: DeleteOperation,
					Entry: Entry{
						Key: txn.Entry.Key,
					},
				}
			} else {
				rollbackEntry = &TxnEntry{
					Operation: PutOperation,
					Entry: Entry{
						Key:   entry.Key,
						Value: entry.Value,
					},
				}
			}
			err = t.PutInternal(ctx, &txn.Entry)
			if err != nil {
				retErr = append(retErr, ErrTransactionalOperation(*txn, err.Error()))
				dirty = true
				break TxnWalk
			}
			rollbackStack = append([]*TxnEntry{rollbackEntry}, rollbackStack...)
		}
	}
	// Need to roll back because we hit an error along the way
	if dirty {
		// While traversing this, if we get an error, we continue anyways in
		// best-effort fashion
		for _, txn := range rollbackStack {
			switch txn.Operation {
			case DeleteOperation:
				err := t.DeleteInternal(ctx, txn.Entry.Key)
				if err != nil {
					err = stacktrace.Propagate(err, "rollback error")
					retErr = append(retErr, ErrTransactionalOperation(*txn, err.Error()))
				}
			case PutOperation:
				err := t.PutInternal(ctx, &txn.Entry)
				if err != nil {
					err = stacktrace.Propagate(err, "rollback error")
					retErr = append(retErr, ErrTransactionalOperation(*txn, err.Error()))
				}
			}
		}
	}
	return retErr
}

// Prefixes is a shared helper function returns all parent 'folders' for a
// given vault key.
// e.g. for 'foo/bar/baz', it returns ['foo', 'foo/bar']
func Prefixes(s string) []string {
	components := strings.Split(s, "/")
	result := []string{}
	for i := 1; i < len(components); i++ {
		result = append(result, strings.Join(components[:i], "/"))
	}
	return result
}

// SanitizeEntryPath is a shared helper function that takes a path and
// returns a sanitized version of it. This is used to ensure that paths
// are always in a consistent format.
func SanitizeEntryPath(key string) string {
	splitted := make([]string, 0)
	for _, j := range strings.Split(key, "/") {
		trimmed := strings.Trim(j, "/")
		trimmed = strings.TrimSpace(trimmed)
		if trimmed != "" {
			splitted = append(splitted, trimmed)
		}
	}
	return strings.Join(splitted, "/")
}
