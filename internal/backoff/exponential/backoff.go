package exponential

import (
	"math"
	"sync"
	"time"

	urandom "github.com/da-moon/podinfo/internal/urandom"
	physical "github.com/da-moon/podinfo/sdk/physical"
)

var _ physical.Backoff = (*Backoff)(nil)

// TODO: maybe remove mutex?
type Backoff struct {
	sync.RWMutex
	MaxJitterInterval int64
	InitialTimeout    float64
	MaxTimeout        float64
	Exponent          float64
}

func Default() *Backoff {
	return &Backoff{
		InitialTimeout:    float64(2 * time.Millisecond),
		MaxTimeout:        float64(10 * time.Millisecond),
		MaxJitterInterval: int64(1 * time.Millisecond),
		Exponent:          2.0,
	}
}

func (b *Backoff) NextInterval(retry int) time.Duration {
	if retry < 0 {
		retry = 0
	}
	if b.MaxJitterInterval < 0 {
		b.MaxJitterInterval = 0
	}
	// Min ( initial timeout * ( (e ^ retry) , maxTimeout)
	base := math.Min(
		b.InitialTimeout*
			math.Pow(b.Exponent, float64(retry)),
		b.MaxTimeout,
	)
	jitter := float64(urandom.Int64(b.MaxJitterInterval + 1))
	return time.Duration(base+jitter) * time.Millisecond
}
