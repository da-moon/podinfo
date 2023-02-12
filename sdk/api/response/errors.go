package response

import (
	"net/http"
	"strings"

	stacktrace "github.com/palantir/stacktrace"
)

// HTTPError is a propagatable error
// that can also be marshaled into a human radable json response
type HTTPError interface {
	// Error satisfies the error interface
	// it is the root error cause (no stacktrace)
	Error() string
	// StatusCode returns the HTTP status code
	StatusCode() stacktrace.ErrorCode
	// EncodeJSON encodes the error as JSON
	EncodeJSON() ([]byte, error)
}

// ─── INTERNAL SERVER ERROR ──────────────────────────────────────────────────────
type errInternalServerError string

// ErrInternalServerError is the constructor of errInternalServerError
func ErrInternalServerError() HTTPError {
	return errInternalServerError("Internal Server Error")
}

// Error satisfies the error interface
// it is the root error cause (no stacktrace)
func (e errInternalServerError) Error() string {
	return stacktrace.NewMessageWithCode(e.StatusCode(), string(e)).Error()
}

// StatusCode returns the HTTP status code
func (e errInternalServerError) StatusCode() stacktrace.ErrorCode {
	return stacktrace.ErrorCode(http.StatusInternalServerError)
}

// EncodeJSON encodes the error as JSON
func (e errInternalServerError) EncodeJSON() ([]byte, error) {
	response := &Response{
		Success: false,
		Body: struct {
			Msg string `json:"msg"`
		}{
			Msg: string(e),
		},
	}

	return response.EncodeJSON()
}

// IsInternalServerError checks if a error is of
// the same type as errInternalServerError
func IsInternalServerError(err error) bool {
	return strings.EqualFold(ErrInternalServerError().Error(), err.Error())
}

// ─── MALFORMED REQUEST ──────────────────────────────────────────────────────────

// errMalformedRequest is the error shown when the
// incoming request is malformed and we cannot unmarshal it.
type errMalformedRequest string

// ErrMalformedRequest is the constructor of errMalformedRequest
func ErrMalformedRequest() HTTPError {
	return errMalformedRequest("Malformed request")
}

// Error satisfies the error interface
// it is the root error cause (no stacktrace)
func (e errMalformedRequest) Error() string {
	return stacktrace.NewMessageWithCode(e.StatusCode(), string(e)).Error()
}

// StatusCode returns 400 status code
func (e errMalformedRequest) StatusCode() stacktrace.ErrorCode {
	return stacktrace.ErrorCode(http.StatusBadRequest)
}

// EncodeJSON encodes the error as JSON
func (e errMalformedRequest) EncodeJSON() ([]byte, error) {
	response := &Response{
		Success: false,
		Body: struct {
			Msg string `json:"msg"`
		}{
			Msg: e.Error(),
		},
	}
	return response.EncodeJSON()
}

// IsMalformedRequest checks if a error is of
// the same type as errMalformedRequest
func IsMalformedRequest(err error) bool {
	return strings.EqualFold(ErrMalformedRequest().Error(), err.Error())
}

type errMethodNotAllowed string

// ErrMethodNotAllowed is the constructor of errMethodNotAllowed
func ErrMethodNotAllowed() HTTPError {
	return errMethodNotAllowed("The specified method is not allowed against this resource")
}

// Error satisfies the error interface
func (e errMethodNotAllowed) Error() string {
	return stacktrace.NewMessageWithCode(e.StatusCode(), string(e)).Error()
}

// StatusCode returns the HTTP status code
func (e errMethodNotAllowed) StatusCode() stacktrace.ErrorCode {
	return stacktrace.ErrorCode(http.StatusMethodNotAllowed)
}

// EncodeJSON encodes the error as JSON
func (e errMethodNotAllowed) EncodeJSON() ([]byte, error) {
	response := &Response{
		Success: false,
		Body: struct {
			Msg string `json:"msg"`
		}{
			Msg: string(e),
		},
	}

	return response.EncodeJSON()
}

// IsMethodNotAllowed checks if a error is of
// the same type as errMethodNotAllowed
func IsMethodNotAllowed(err error) bool {
	return strings.EqualFold(ErrMethodNotAllowed().Error(), err.Error())
}
