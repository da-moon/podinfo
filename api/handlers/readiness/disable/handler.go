package disable

import (
	"sync"

	logger "github.com/da-moon/northern-labs-interview/internal/logger"
)

// handler struct encapsulates the state this API endpoint
// handler needs
type handler struct {
	// mutex for guard shared state
	mutex sync.RWMutex
	// log is the logger for this handler
	log *logger.WrappedLogger
}

func (h *handler) SetLogger(l *logger.WrappedLogger) {
	h.mutex.Lock()
	defer h.mutex.Unlock()
	h.log = l
}