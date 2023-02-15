package handlers

import (
	env "github.com/da-moon/northern-labs-interview/api/handlers/env"
	headers "github.com/da-moon/northern-labs-interview/api/handlers/headers"
	liveness "github.com/da-moon/northern-labs-interview/api/handlers/liveness"
	readiness "github.com/da-moon/northern-labs-interview/api/handlers/readiness"
	readinessDisable "github.com/da-moon/northern-labs-interview/api/handlers/readiness/disable"
	readinessEnable "github.com/da-moon/northern-labs-interview/api/handlers/readiness/enable"
	logger "github.com/da-moon/northern-labs-interview/internal/logger"
	stacktrace "github.com/palantir/stacktrace"
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
	// GET /healthz endpoint
	liveness.Router.SetLogger(l)
	err := liveness.Router.Register()
	if err != nil {
		err = stacktrace.Propagate(err, "failed to initialize HTTP request handlers for '%s' (%s)", liveness.Name, liveness.Path)
		return err
	}
	// GET /readyz endpoint
	readiness.Router.SetLogger(l)
	err = readiness.Router.Register()
	if err != nil {
		err = stacktrace.Propagate(err, "failed to initialize HTTP request handlers for '%s' (%s)", readiness.Name, readiness.Path)
		return err
	}
	// GET /readyz/enable endpoint
	readinessEnable.Router.SetLogger(l)
	err = readinessEnable.Router.Register()
	if err != nil {
		err = stacktrace.Propagate(err, "failed to initialize HTTP request handlers for '%s' (%s)", readinessEnable.Name, readinessEnable.APIGroup+readinessEnable.Path)
		return err
	}
	// GET /readyz/disable endpoint
	readinessDisable.Router.SetLogger(l)
	err = readinessDisable.Router.Register()
	if err != nil {
		err = stacktrace.Propagate(err, "failed to initialize HTTP request handlers for '%s' (%s)", readinessDisable.Name, readinessDisable.APIGroup+readinessDisable.Path)
		return err
	}
	// GET /env endpoint
	env.Router.SetLogger(l)
	err = env.Router.Register()
	if err != nil {
		err = stacktrace.Propagate(err, "failed to initialize HTTP request handlers for '%s' (%s)", env.Name, env.Path)
		return err
	}
	// GET /headers endpoint
	headers.Router.SetLogger(l)
	err = headers.Router.Register()
	if err != nil {
		err = stacktrace.Propagate(err, "failed to initialize HTTP request handlers for '%s' (%s)", headers.Name, headers.Path)
		return err
	}
	return nil
}
