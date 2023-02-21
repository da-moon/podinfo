package core

import (
	"fmt"

	handlers "github.com/da-moon/podinfo/api/handlers"
	registry "github.com/da-moon/podinfo/api/registry"
	mux "github.com/gorilla/mux"
	stacktrace "github.com/palantir/stacktrace"
)

// Routes returns the router for the api server core
func (c *Config) Router() (*mux.Router, error) {
	err := handlers.Initialize(c.log)
	if err != nil {
		err = stacktrace.Propagate(err, "failed to initialize HTTP router")
		return nil, err
	}
	routes, err := registry.Dispense()
	if err != nil {
		err = stacktrace.Propagate(err, "router registry failed to dispense routes")
		return nil, err
	}
	err = c.Telemetry(routes)
	if err != nil {
		err = stacktrace.Propagate(err, "router registry failed to dispense routes")
		return nil, err
	}
	c.log.Info("metrics exporter route was successfully initialized.")
	l, err := c.Listener()
	if err != nil {
		err = stacktrace.Propagate(err, "could not initialize request router")
		return nil, err
	}
	routes.Synopsis(c.log.Writer(), "127.0.0.1:"+l.Port)
	result := routes.Router()
	// PrintRoutes(result) //nolint:gocritic // this is a debug function
	return result, nil
}
func PrintRoutes(r *mux.Router) {
	r.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		tpl, _ := route.GetPathTemplate()
		met, _ := route.GetMethods()
		pathRegex, _ := route.GetPathRegexp()
		queriesRegexp, _ := route.GetQueriesRegexp()
		fmt.Printf("\napi: route path <%s> | path method <%v> | Path Regex <%v> | queries regex <%v>", tpl, met, pathRegex, queriesRegexp)
		return nil
	})
}
