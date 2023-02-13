package liveness

import (
	"net/http"

	middlewares "github.com/da-moon/northern-labs-interview/api/middlewares"
	registry "github.com/da-moon/northern-labs-interview/api/registry"
	logger "github.com/da-moon/northern-labs-interview/internal/logger"
	route "github.com/da-moon/northern-labs-interview/sdk/api/route"
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

// Initialize register this request handler in the central
// api handlers registry. It also sets the appropriate
// middlewares.
// this function should be called when server is getting ready.
func (h *Handler) Initialize() {
	r := route.New()
	r.SetName(Name)
	r.SetPath(Path)
	r.SetMethod(http.MethodGet)
	r.SetHandlerFunc(handler)
	h.log.Info("Adding log middleware for '%s' handler at '%s'", Name, Path)
	r.AppendMiddleware(middlewares.Log(h.log))
	r.AppendMiddleware(middlewares.Metrics(Path, h.log))
	registry.Register("", *r)
}
