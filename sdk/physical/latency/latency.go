package latency

import (
	"context"
	"sync"
	"time"

	"github.com/da-moon/podinfo/internal/logger"
	physical "github.com/da-moon/podinfo/sdk/physical"
)

// Injector struct is used to add
// latency into underlying physical requests
type Injector struct {
	log           *logger.WrappedLogger
	backend       physical.Backend
	latency       time.Duration
	jitterPercent int
	randomLock    *sync.Mutex
}

// New returns a wrapped physical backend to simulate latency
func New(b physical.Backend, latency time.Duration, jitter int, log *logger.WrappedLogger) *Injector {
	if jitter < 0 || jitter > 100 {
		jitter = physical.DefaultJitterPercent
	}
	// logger.Logger().Println("creating latency injector")
	return &Injector{
		log:           log,
		backend:       b,
		latency:       latency,
		jitterPercent: jitter,
		randomLock:    new(sync.Mutex),
	}
}

// SetLatency -
func (l *Injector) SetLatency(latency time.Duration) {
	l.log.Debug("Changing backend latency to %#v", latency)
	l.latency = latency
}

// Put is a latent put request
func (l *Injector) Put(ctx context.Context, entry *physical.Entry) error {
	l.addLatency()
	return l.backend.Put(ctx, entry)
}

// Get is a latent get request
func (l *Injector) Get(ctx context.Context, key string) (*physical.Entry, error) {
	l.addLatency()
	return l.backend.Get(ctx, key)
}

// Delete is a latent delete request
func (l *Injector) Delete(ctx context.Context, key string) error {
	l.addLatency()
	return l.backend.Delete(ctx, key)
}

// List is a latent list request
func (l *Injector) List(ctx context.Context, prefix string) ([]string, error) {
	l.addLatency()
	return l.backend.List(ctx, prefix)
}
