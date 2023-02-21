package retry

import (
	"time"

	constant "github.com/da-moon/podinfo/internal/backoff/constant"
	exponential "github.com/da-moon/podinfo/internal/backoff/exponential"
)

// ────────────────────────────────────────────────────────────────────────────────

// Option struct contains configuration options for the retirable operation.
type Option struct {
	retryOption              func(*Retry)
	constantBackoffOption    func(*constant.Backoff)
	exponentialBackoffOption func(*exponential.Backoff)
}

// WithRetryCount sets the retry count.
func WithRetryCount(retryCount int) Option {
	return Option{
		constantBackoffOption:    nil,
		exponentialBackoffOption: nil,
		retryOption: func(b *Retry) {
			b.Lock()
			defer b.Unlock()
			b.retryCount = retryCount
		},
	}
}

// WithInterval sets the interval for the constant backoff.
func WithInterval(arg time.Duration) Option {
	return Option{
		retryOption:              nil,
		exponentialBackoffOption: nil,
		constantBackoffOption: func(b *constant.Backoff) {
			b.Lock()
			defer b.Unlock()
			b.Interval = int64(arg / time.Millisecond)
		},
	}
}

// WithMaxJitterInterval sets the maximum jitter interval for the exponential backoff.
func WithMaxJitterInterval(arg time.Duration) Option {
	return Option{
		retryOption: nil,
		exponentialBackoffOption: func(b *exponential.Backoff) {
			b.Lock()
			defer b.Unlock()
			b.MaxJitterInterval = int64(arg / time.Millisecond)
		},
		constantBackoffOption: func(b *constant.Backoff) {
			b.Lock()
			defer b.Unlock()
			b.MaxJitterInterval = int64(arg / time.Millisecond)
		},
	}
}

// WithInitialTimeout sets the initial timeout for the exponential backoff.
func WithInitialTimeout(arg time.Duration) Option {
	return Option{
		constantBackoffOption: nil,
		retryOption:           nil,
		exponentialBackoffOption: func(b *exponential.Backoff) {
			b.Lock()
			defer b.Unlock()
			b.InitialTimeout = float64(arg / time.Millisecond)
		},
	}
}

// WithMaxTimeout sets the maximum timeout for the exponential backoff.
func WithMaxTimeout(arg time.Duration) Option {
	return Option{
		constantBackoffOption: nil,
		retryOption:           nil,
		exponentialBackoffOption: func(b *exponential.Backoff) {
			b.Lock()
			defer b.Unlock()
			b.MaxTimeout = float64(arg / time.Millisecond)
		},
	}
}

// WithExponent sets the exponent for the exponential backoff.
func WithExponent(arg float64) Option {
	return Option{
		constantBackoffOption: nil,
		retryOption:           nil,
		exponentialBackoffOption: func(b *exponential.Backoff) {
			b.Lock()
			defer b.Unlock()
			b.Exponent = arg
		},
	}
}
