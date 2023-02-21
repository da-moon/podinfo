package handlers

import (
	"net/http"

	middlewares "github.com/da-moon/podinfo/api/middlewares"
	registry "github.com/da-moon/podinfo/api/registry"
	route "github.com/da-moon/podinfo/sdk/api/route"
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
	// registry.Register(Prefix, *r)
	registry.Register("", *r)
}
