package latency

import (
	"time"

	"github.com/da-moon/podinfo/internal/urandom"
)

func (l *Injector) addLatency() {
	// Calculate a value between 1 +- jitter%
	percent := 100
	if l.jitterPercent > 0 {
		min := 100 - l.jitterPercent
		max := 100 + l.jitterPercent
		l.randomLock.Lock()
		percent = urandom.Integer(min, max)
		l.randomLock.Unlock()
	}
	latencyDuration := time.Duration(int(l.latency) * percent / 100)
	time.Sleep(latencyDuration)
}
