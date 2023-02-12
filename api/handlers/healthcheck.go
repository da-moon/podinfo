package handlers

import (
	"net/http"
	"time"

	middlewares "github.com/da-moon/northern-labs-interview/api/middlewares"
	registry "github.com/da-moon/northern-labs-interview/api/registry"
	version "github.com/da-moon/northern-labs-interview/build/go/version"
	runtimex "github.com/da-moon/northern-labs-interview/internal/runtimex"
	response "github.com/da-moon/northern-labs-interview/sdk/api/response"
	route "github.com/da-moon/northern-labs-interview/sdk/api/route"
)

func healthcheck() {
	type resp struct {
		Status string                    `json:"status" mapstructure:"status"`
		Build  *version.BuildInformation `json:"build" mapstructure:"build"`
		Stats  *runtimex.Stats           `json:"stats" mapstructure:"stats"`
		Time   int                       `json:"time" mapstructure:"time"`
	}
	name := "healthcheck"
	path := "/health"
	r := route.New()
	r.SetName(name)
	r.SetPath(path)
	r.SetMethod(http.MethodGet)
	r.SetHandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response.WriteSuccessfulJSON(w, r, &resp{
			// TODO: create an enum and have the server respond properly it's current state
			Status: "healthy",
			Build:  version.New(),
			Stats:  runtimex.GetStats(),
			Time:   int(time.Now().Unix()),
		})
		return
	})
	log.Info("Adding log middleware for '%s' handler at '%s'", name, path)
	r.AppendMiddleware(middlewares.Log(log))
	r.AppendMiddleware(middlewares.Metrics(path, log))
	registry.Register("", *r)
}
