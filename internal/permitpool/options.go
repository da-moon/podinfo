package permitpool

import (
	"sync"
)

type option func(*permitPool)

type permitPool struct {
	mutex   sync.Mutex
	sem     chan int
	permits int
}

// WithPermits sets number of permits in the permit pool
func WithPermits(arg int) option {
	return func(e *permitPool) {
		e.mutex.Lock()
		defer e.mutex.Unlock()
		e.permits = arg
	}
}
