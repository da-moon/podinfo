package readiness

import (
	"sync"
	"sync/atomic"

	logger "github.com/da-moon/northern-labs-interview/internal/logger"
)

type handler struct {
	mutex  sync.RWMutex
	log    *logger.WrappedLogger
	status atomic.Bool
}
