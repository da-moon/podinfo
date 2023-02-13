package liveness

import (
	logger "github.com/da-moon/northern-labs-interview/internal/logger"
)

// Handler struct encapsulates the state this API endpoint
// handler needs
type Handler struct {
	// log is the logger for this handler
	log *logger.WrappedLogger
}

// New function returns a new instance of request
// Handler
func New(l *logger.WrappedLogger) *Handler {
	return &Handler{
		log: l,
	}
}
