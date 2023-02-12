package retry

import (
	"context"
	"sync"

	constant "github.com/da-moon/northern-labs-interview/internal/backoff/constant"
	exponential "github.com/da-moon/northern-labs-interview/internal/backoff/exponential"
	logger "github.com/da-moon/northern-labs-interview/internal/logger"
	physical "github.com/da-moon/northern-labs-interview/sdk/physical"
)

type fallbackFn func(error) error

type Retry struct {
	sync.RWMutex
	log        *logger.WrappedLogger
	backend    physical.Backend
	backoff    physical.Backoff
	retryCount int
	name       string
	// TODO: figure it out
	// fallbackFunc fallbackFn
}

// NewConstantBackoff returns a new instance of the retirable physical
// backend that uses constant backoff .
func NewConstantBackoff(log *logger.WrappedLogger, backend physical.Backend, opts ...Option) *Retry {
	if log == nil {
		panic("no logger was provided")
	}
	if backend == nil {
		panic("given physical backend is nil")
	}
	result := &Retry{
		name:       "physical retrier with constant backoff",
		log:        log,
		backend:    backend,
		retryCount: 5,
	}
	b := constant.Default()
	for _, opt := range opts {
		if opt.retryOption != nil {
			opt.retryOption(result)
		}
		if opt.constantBackoffOption != nil {
			opt.constantBackoffOption(b)
		}
	}
	result.backoff = b
	return result
}

// NewExponentialBackoff returns a new instance of the retirable physical
// backend that uses exponential backoff .
func NewExponentialBackoff(log *logger.WrappedLogger, backend physical.Backend, opts ...Option) *Retry {
	if log == nil {
		panic("no logger was provided")
	}
	if backend == nil {
		panic("given physical backend is nil")
	}
	result := &Retry{
		name:       "physical retrier with exponential backoff",
		log:        log,
		backend:    backend,
		retryCount: 5,
	}
	b := exponential.Default()
	for _, opt := range opts {
		if opt.retryOption != nil {
			opt.retryOption(result)
		}
		if opt.exponentialBackoffOption != nil {
			opt.exponentialBackoffOption(b)
		}
	}
	result.backoff = b
	return result
}

// Put is used to insert or update an entry.
func (c *Retry) Put(ctx context.Context, entry *physical.Entry) error {
	return c.PutInternal(ctx, entry)
}

// Get is used to fetch an entry.
func (c *Retry) Get(ctx context.Context, key string) (*physical.Entry, error) {
	return c.GetInternal(ctx, key)
}

// Delete deletes the given key.
func (c *Retry) Delete(ctx context.Context, key string) error {
	return c.DeleteInternal(ctx, key)
}

// List lists all the keys under the given prefix.
func (c *Retry) List(ctx context.Context, prefix string) ([]string, error) {
	return c.ListInternal(ctx, prefix)
}
