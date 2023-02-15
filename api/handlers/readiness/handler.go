package readiness

import (
	"sync"
	"sync/atomic"

	logger "github.com/da-moon/northern-labs-interview/internal/logger"
)

// handler needs
type handler struct {
	// mutex for guard shared state
	mutex sync.RWMutex
	// log is the logger for this handler
	log *logger.WrappedLogger
	// status represents server readiness status
	status atomic.Bool
}
