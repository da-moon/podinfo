package errors

import (
	"net/http"

	response "github.com/da-moon/podinfo/sdk/api/response"
	physical "github.com/da-moon/podinfo/sdk/physical"
	stacktrace "github.com/palantir/stacktrace"
)

var _ response.HTTPError = ErrTransactionalOperations{}

// ErrTransactionalOperations aggregates multiple ErrTransactionalOperations
type ErrTransactionalOperations []physical.ErrTransactionalOperations

// Error satisfies the error interface
// it is the root error cause (no stacktrace)
func (t ErrTransactionalOperations) Error() string {
	return "some/all transactional operations were not successful"
}

// StatusCode returns the HTTP status code
func (t ErrTransactionalOperations) StatusCode() stacktrace.ErrorCode {
	return stacktrace.ErrorCode(http.StatusInternalServerError)
}

// EncodeJSON encodes the error as JSON
func (t ErrTransactionalOperations) EncodeJSON() ([]byte, error) {
	resp := &response.Response{
		Success: false,
		Body: struct {
			Msg         string                     `json:"msg"`
			Diagnostics ErrTransactionalOperations `json:"diagnostics"`
		}{
			Msg:         t.Error(),
			Diagnostics: t,
		},
	}
	return resp.EncodeJSON()
}
