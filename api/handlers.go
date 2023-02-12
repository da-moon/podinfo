package api

import (
	"net/http"

	response "github.com/da-moon/northern-labs-interview/sdk/api/response"
)

// NotFoundHandler returns a simple error on non-existing routes
func NotFoundHandler() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		response.WriteErrorJSON(w, r, response.ErrMethodNotAllowed(), "")
		return
	}
}
