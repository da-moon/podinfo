package handlers

import (
	"net/http"

	middlewares "github.com/da-moon/northern-labs-interview/api/middlewares"
	registry "github.com/da-moon/northern-labs-interview/api/registry"
	route "github.com/da-moon/northern-labs-interview/sdk/api/route"
)

// Handler preflight is a handler used to reply to
// OPTIONS request
func preflight() {
	r := route.New()
	r.SetName("preflight")
	r.SetPath("")
	r.SetMethod(http.MethodOptions)
	r.SetHandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		return
	})
	r.AppendMiddleware(middlewares.Cors)
	registry.Register(Prefix, *r)
}
