package constant_test

import (
	"testing"
	"time"

	"github.com/da-moon/podinfo/internal/backoff/constant"
	"github.com/stretchr/testify/assert"
)

func TestBackoffNextTime(t *testing.T) {
	tests := []struct {
		name     string
		interval int64
		jitter   int64
		counter  int
	}{
		{name: "negative_retry", interval: 100, jitter: 0, counter: -1},
		{name: "simple", interval: 100, jitter: 0, counter: 4},
		{name: "zero_jitter", interval: 100, jitter: 0, counter: 1000},
		{name: "ten_jitter", interval: 100, jitter: 10, counter: 5},
		{name: "fifty_jitter", interval: 100, jitter: 50, counter: 1000},
		{name: "hundred_and_fifty_jitter", interval: 100, jitter: 150, counter: 1000},
		{name: "negative_jitter", interval: 100, jitter: -1, counter: 4},
	}
	for _, tt := range tests {
		name := tt.name
		interval := tt.interval * int64(time.Second/time.Millisecond)
		jitter := tt.jitter * int64(time.Second/time.Millisecond)
		counter := tt.counter
		t.Run(name, func(t *testing.T) {
			backoff := &constant.Backoff{
				Interval:          interval,
				MaxJitterInterval: jitter,
			}
			if counter <= 0 {
				actual := backoff.NextInterval(counter)
				expected := time.Duration(interval) * time.Millisecond
				assert.Equal(t, expected, actual)
				return
			}
			for i := 0; i < counter; i++ {
				actual := backoff.NextInterval(counter)
				expected := time.Duration(interval) * time.Millisecond
				if jitter <= 0 {
					assert.Equal(t, expected, actual)
				} else {
					assert.LessOrEqual(t, expected, actual)
					assert.LessOrEqual(t, backoff.NextInterval(1), time.Duration(interval+jitter)*time.Millisecond)
				}
			}
		})
	}
}
