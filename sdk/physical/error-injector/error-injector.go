package errorinjector

import (
	"context"
	"sync"

	"github.com/da-moon/northern-labs-interview/internal/logger"
	physical "github.com/da-moon/northern-labs-interview/sdk/physical"
)

// ErrorInjector is used to add errors into underlying physical requests
type ErrorInjector struct {
	backend      physical.Backend
	errorPercent int
	randomLock   *sync.Mutex
	log          *logger.WrappedLogger
}

// New returns a wrapped physical backend to inject error
func New(b physical.Backend, errorPercent int, log *logger.WrappedLogger) *ErrorInjector {
	if errorPercent < 0 || errorPercent > 100 {
		errorPercent = physical.DefaultErrorPercent
	}
	// logger.Logger().Println("creating error injector")
	return &ErrorInjector{
		backend:      b,
		log:          log,
		errorPercent: errorPercent,
		randomLock:   new(sync.Mutex),
	}
}

// Put stores the given entry
// and based on backend configuration, injects errors
func (e *ErrorInjector) Put(ctx context.Context, entry *physical.Entry) error {
	if err := e.addError(); err != nil {
		return err
	}
	return e.backend.Put(ctx, entry)
}

// Get retrieves the entry at the given key
// and based on backend configuration, injects errors.
func (e *ErrorInjector) Get(ctx context.Context, key string) (*physical.Entry, error) {
	if err := e.addError(); err != nil {
		return nil, err
	}
	return e.backend.Get(ctx, key)
}

// Delete removes the entry at the given key
// and based on backend configuration, injects errors
func (e *ErrorInjector) Delete(ctx context.Context, key string) error {
	if err := e.addError(); err != nil {
		return err
	}
	return e.backend.Delete(ctx, key)
}

// List returns a list of keys under the given prefix
// and based on backend configuration, injects errors
func (e *ErrorInjector) List(ctx context.Context, prefix string) ([]string, error) {
	if err := e.addError(); err != nil {
		return nil, err
	}
	return e.backend.List(ctx, prefix)
}
