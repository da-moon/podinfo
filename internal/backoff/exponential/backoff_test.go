package exponential_test

import (
	"math"
	"testing"
	"time"

	"github.com/da-moon/northern-labs-interview/internal/backoff/exponential"
	"github.com/stretchr/testify/assert"
)

func TestBackoffNextTime(t *testing.T) {
	tests := []struct {
		name           string
		jitter         int64
		initialTimeout float64
		maxTimeout     float64
		exponent       float64
		counter        int
		single         bool
	}{
		{name: "simple", initialTimeout: 100, maxTimeout: 1000, exponent: 2, jitter: 0, counter: 4},
		{name: "negative_retry", initialTimeout: 100, maxTimeout: 1000, exponent: 2, jitter: 0, counter: -1},
		{name: "negative_jitter", initialTimeout: 100, maxTimeout: 1000, exponent: 2, jitter: -1, counter: 4},
		{name: "larger_than_max_timeout", initialTimeout: 100, maxTimeout: 1000, exponent: 2, jitter: 0, counter: 5, single: true},
		{name: "equal_max_timeout", initialTimeout: 100, maxTimeout: 800, exponent: 2, jitter: 0, counter: 3, single: true},
		{name: "no_jitter", initialTimeout: 100, maxTimeout: 1000, exponent: 2, jitter: 0, counter: 10000},
		{name: "some_jitter", initialTimeout: 100, maxTimeout: 1000, exponent: 2, jitter: 1, counter: 10000},
		{name: "higher_jitter", initialTimeout: 100, maxTimeout: 1000, exponent: 2, jitter: 20, counter: 10000},
	}
	for _, tt := range tests {
		name := tt.name
		initialTimeout := tt.initialTimeout * float64(time.Millisecond)
		jitter := tt.jitter * int64(time.Second/time.Millisecond)
		maxTimeout := tt.maxTimeout * float64(time.Second/time.Millisecond)
		exponent := tt.exponent
		counter := tt.counter
		single := tt.single
		t.Run(name, func(t *testing.T) {
			backoff := &exponential.Backoff{
				MaxJitterInterval: jitter,
				InitialTimeout:    initialTimeout,
				MaxTimeout:        maxTimeout,
				Exponent:          exponent,
			}
			if counter <= 0 || single {
				actual := backoff.NextInterval(counter)
				timeout := initialTimeout
				if maxTimeout < timeout {
					timeout = maxTimeout
				}
				expected := time.Duration(timeout) * time.Millisecond
				assert.Equal(t, expected, actual)
				return
			}
			for i := 0; i < counter; i++ {
				actual := backoff.NextInterval(counter)
				timeout := initialTimeout
				timeout *= math.Pow(exponent, float64(i))
				if maxTimeout < timeout {
					timeout = maxTimeout
				}
				expected := time.Duration(timeout) * time.Millisecond
				if jitter <= 0 {
					assert.Equal(t, expected, actual)
				} else {
					assert.LessOrEqual(t, expected, actual)
					assert.LessOrEqual(t, backoff.NextInterval(1), time.Duration(int64(timeout)+jitter)*time.Millisecond)
				}
			}
		})
	}
}
