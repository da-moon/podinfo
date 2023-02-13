package handlers

import (
	liveness "github.com/da-moon/northern-labs-interview/api/handlers/liveness"
	logger "github.com/da-moon/northern-labs-interview/internal/logger"
	"github.com/palantir/stacktrace"
)

var (
	log *logger.WrappedLogger //nolint:gochecknoglobals //this is exposed so it can be set from caller
)

// Initialize function sets up routes.
// this function should be called when server is getting ready.
func Initialize(l *logger.WrappedLogger) error {
	log = l
	// log.Info("initializing API '%s' request handler set", Prefix)
	// preflight()
	debug()
	liveness.Router.SetLogger(l)
	err := liveness.Router.Register()
	if err != nil {
		err = stacktrace.Propagate(err, "failed to initialize HTTP request handlers for '%s' (%s)", liveness.Name, liveness.Path)
		return err
	}
	return nil
}
