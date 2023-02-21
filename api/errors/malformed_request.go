package errors

import (
	"fmt"
	"net/http"
	"strings"

	response "github.com/da-moon/podinfo/sdk/api/response"
	stacktrace "github.com/palantir/stacktrace"
)

// ─── MALFORMED REQUEST ──────────────────────────────────────────────────────────

var _ response.HTTPError = errMalformedRequest{}

// errMalformedRequest is the error shown when the
// incoming request is malformed and we cannot unmarshal it.
// TODO: maybe add this feature to response.errMalformedRequest
// TODO: add information about the missing field that led to malformed request
type errMalformedRequest struct {
	cause string
	err   string
}

// ErrMalformedRequest is the constructor of errMalformedRequest
func ErrMalformedRequest(cause string) errMalformedRequest {
	return errMalformedRequest{
		cause: cause,
		err:   "Malformed request",
	}
}

// Error satisfies the error interface
// it is the root error cause (no stacktrace)
func (e errMalformedRequest) Error() string {
	msg := e.err
	if e.cause != "" {
		msg += fmt.Sprintf(". cause : %s", e.cause)
	}
	return stacktrace.NewMessageWithCode(e.StatusCode(), msg).Error()
}

// StatusCode returns 400 status code
func (e errMalformedRequest) StatusCode() stacktrace.ErrorCode {
	return stacktrace.ErrorCode(http.StatusBadRequest)
}

// EncodeJSON encodes the error as JSON
func (e errMalformedRequest) EncodeJSON() ([]byte, error) {
	resp := &response.Response{
		Success: false,
		Body: struct {
			Msg string `json:"msg"`
		}{
			Msg: e.Error(),
		},
	}
	return resp.EncodeJSON()
}

// IsMalformedRequest checks if a error is of
// the same type as errMalformedRequest
func IsMalformedRequest(err error) bool {
	return strings.Contains(err.Error(), "Malformed request")
}
