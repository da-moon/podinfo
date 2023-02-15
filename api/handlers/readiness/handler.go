package readiness

import (
	"sync"
	"sync/atomic"

	logger "github.com/da-moon/northern-labs-interview/internal/logger"
)

// handler struct encapsulates the state this API endpoint
// handler needs
type handler struct {
	// mutex for guard shared state
	mutex sync.RWMutex
	// log is the logger for this handler
	log *logger.WrappedLogger
	// status represents server readiness status
	status atomic.Bool
}

// SetLogger sets the logger for this handler
func (h *handler) SetLogger(l *logger.WrappedLogger) {
	h.mutex.Lock()
	defer h.mutex.Unlock()
	h.log = l
}

// GetLogger returns the logger for this handler
func (h *handler) GetLogger() *logger.WrappedLogger {
	h.mutex.RLock()
	defer h.mutex.RUnlock()
	return h.log
}

// SetStatus sets the readiness status for this handler
// mutex is not used here as atomic.Bool is more performant
func (h *handler) SetStatus(status Status) {
	if status == OK {
		h.status.Store(true)
		return
	}
	h.status.Store(false)
}
