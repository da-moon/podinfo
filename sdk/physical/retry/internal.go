package retry

import (
	"context"
	"time"

	multierror "github.com/da-moon/northern-labs-interview/internal/multierror"
	physical "github.com/da-moon/northern-labs-interview/sdk/physical"
	"github.com/palantir/stacktrace"
)

// GetInternal gets the value at the given key.
func (c *Retry) GetInternal(ctx context.Context, key string) (*physical.Entry, error) {
	multiErr := &multierror.MultiError{}
	for i := 0; i <= c.retryCount; i++ {
		entry, err := c.backend.Get(ctx, key)
		if err != nil {
			if physical.IsErrNotExist(err) {
				return nil, err
			}
			err = stacktrace.Propagate(err, "try [ % d ]", i)
			multiErr.Add(err.Error())
			backoffTime := c.backoff.NextInterval(i)
			time.Sleep(backoffTime)
			continue
		}
		return entry, nil
	}
	return nil, multiErr
}

// PutInternal puts the given entry.
func (c *Retry) PutInternal(ctx context.Context, entry *physical.Entry) error {
	multiErr := &multierror.MultiError{}
	for i := 0; i <= c.retryCount; i++ {
		err := c.backend.Put(ctx, entry)
		if err != nil {
			err = stacktrace.Propagate(err, "try [ % d ]", i)
			multiErr.Add(err.Error())
			backoffTime := c.backoff.NextInterval(i)
			time.Sleep(backoffTime)
			continue
		}
		return nil
	}
	return multiErr
}

// DeleteInternal deletes the given key.
func (c *Retry) DeleteInternal(ctx context.Context, key string) error {
	multiErr := &multierror.MultiError{}
	for i := 0; i <= c.retryCount; i++ {
		err := c.backend.Delete(ctx, key)
		if err != nil {
			if physical.IsErrNotExist(err) {
				return err
			}
			err = stacktrace.Propagate(err, "try [ % d ]", i)
			multiErr.Add(err.Error())
			backoffTime := c.backoff.NextInterval(i)
			time.Sleep(backoffTime)
			continue
		}
		return nil
	}
	return multiErr
}

// ListInternal lists all the keys under the given prefix.
func (c *Retry) ListInternal(ctx context.Context, prefix string) ([]string, error) {
	multiErr := &multierror.MultiError{}
	for i := 0; i <= c.retryCount; i++ {
		result, err := c.backend.List(ctx, prefix)
		if err != nil {
			if physical.IsErrNotExist(err) {
				return nil, err
			}
			err = stacktrace.Propagate(err, "try [ % d ]", i)
			multiErr.Add(err.Error())
			backoffTime := c.backoff.NextInterval(i)
			time.Sleep(backoffTime)
			continue
		}
		return result, nil
	}
	return nil, multiErr
}
