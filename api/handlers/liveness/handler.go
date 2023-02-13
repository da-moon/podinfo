package liveness

import (
	"net/http"

	response "github.com/da-moon/northern-labs-interview/sdk/api/response"
)

const (
	// Name stores a human friendly, param-cased
	// unique identifier name for this endpoint
	Name = "kubernetes-liveness-probe"
	// Path represents the URI path of this endpoint
	Path = "/healthz"
)

// handler function handles incoming HTTP request
// it satisfies golang's stdlib
// request handler interface (http.HandlerFunc)
var handler = func(w http.ResponseWriter, r *http.Request) { //nolint:gochecknoglobals //this function is scoped only to this package
	response.WriteSuccessfulJSON(w, r, &Response{
		Status: "OK",
	})
	return
}
