package constant

import (
	"sync"
	"time"

	urandom "github.com/da-moon/northern-labs-interview/internal/urandom"
	physical "github.com/da-moon/northern-labs-interview/sdk/physical"
)

var _ physical.Backoff = (*Backoff)(nil)

// TODO: make the fields private
type Backoff struct {
	sync.RWMutex
	Interval          int64
	MaxJitterInterval int64
}

func Default() *Backoff {
	return &Backoff{
		Interval:          int64(2 * time.Millisecond),
		MaxJitterInterval: int64(1 * time.Millisecond),
	}
}

func (b *Backoff) NextInterval(_ int) time.Duration {
	if b.MaxJitterInterval <= 0 {
		b.MaxJitterInterval = 0
	}
	backOffInterval := b.Interval
	randJitter := urandom.Int64(b.MaxJitterInterval + 1)
	return (time.Duration(backOffInterval) * time.Millisecond) + (time.Duration(randJitter) * time.Millisecond)
}
