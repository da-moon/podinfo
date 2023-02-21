package errorinjector

import (
	"github.com/da-moon/podinfo/internal/urandom"
	"github.com/palantir/stacktrace"
)

func (e *ErrorInjector) addError() error {
	e.randomLock.Lock()
	roll := urandom.Int(100)
	e.randomLock.Unlock()
	if roll < e.errorPercent {
		return stacktrace.NewError("random error")
	}
	return nil
}

// SetErrorPercentage -
func (e *ErrorInjector) SetErrorPercentage(p int) {
	e.errorPercent = p
}
