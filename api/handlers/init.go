package handlers

import (
	logger "github.com/da-moon/northern-labs-interview/internal/logger"
)

const (
	// Prefix represent the version of the api request handler functions in this package
	// map to
	Prefix = "/v1"
)

var (
	log *logger.WrappedLogger
)

// Initialize function sets up routes.
// this function should be called when server is getting ready.
func Initialize(l *logger.WrappedLogger) {
	log = l
	// log.Info("initializing API '%s' request handler set", Prefix)
	// preflight()
	healthcheck()
	debug()
}
