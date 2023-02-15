package disable

import (
	"sync"

	logger "github.com/da-moon/northern-labs-interview/internal/logger"
)

type handler struct {
	// mutex for guard shared state
	mutex sync.RWMutex
	// log is the logger for this handler
	log *logger.WrappedLogger
}
