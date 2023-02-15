package liveness

import (
	"net/http"

	middlewares "github.com/da-moon/northern-labs-interview/api/middlewares"
	registry "github.com/da-moon/northern-labs-interview/api/registry"
	route "github.com/da-moon/northern-labs-interview/sdk/api/route"
	"github.com/palantir/stacktrace"
)

const (
	// Name stores a human friendly, param-cased
	// unique identifier name for this endpoint
	Name = "kubernetes-liveness-probe"
	// APIGroup stores API group (prefix) for this URI
	// the full URI is /<prefix>/<path>
	APIGroup = ""
	// Path represents the URI path of this endpoint
	Path = "/healthz"
)

var (
	// Router is accessible from outside in case
	// other packages need to access it's state
	Router = new(handler) //nolint:gochecknoglobals // safe as it has mutex guard and access internal state is through getter/setter functions
)

// Register this request handler in the central
// api handlers registry. It also sets the appropriate
// middlewares.
// this function should be called when server is getting ready.
func (h *handler) Register() error {
	l := h.GetLogger()
	if l == nil {
		return stacktrace.NewError("logger is nil")
	}
	r := route.New()
	r.SetName(Name)
	r.SetPath(Path)
	r.SetMethod(http.MethodGet)
	r.SetHandlerFunc(HandlerFn)
	// l.Info("Adding log middleware for '%s' handler", Name)
	// r.AppendMiddleware(middlewares.Log(l))
	r.AppendMiddleware(middlewares.Metrics(Path, l))
	registry.Register(APIGroup, *r)
	return nil
}
