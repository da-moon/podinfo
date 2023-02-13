package response

import (
	"fmt"
	"net/http"
	"strconv"

	fastjson "github.com/da-moon/northern-labs-interview/sdk/api/fastjson"
	stacktrace "github.com/palantir/stacktrace"
	logrus "github.com/sirupsen/logrus"
)

// WriteSuccessfulJSON logs and sends a new json response to the client
// TODO: this method must retry ops on failure
// TODO: turn this into a receiver of Response
func WriteSuccessfulJSON(
	w http.ResponseWriter,
	r *http.Request,
	body interface{},
) {
	var internalErr error
	defer func() {
		if internalErr != nil {
			LogErrorResponse(r, internalErr, "")
		}
	}()

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	response := &Response{
		Success: true,
		Body:    body,
	}
	result, err := response.EncodeJSON()
	if err != nil {
		WriteErrorJSON(w, r, ErrInternalServerError(), fmt.Sprintf("root cause: %s", err.Error()))
		return
	}
	w.Header().Set("Content-Length", strconv.Itoa(len(result)))
	_, err = w.Write(result)
	if err != nil {
		internalErr = stacktrace.Propagate(err, ErrInternalServerError().Error())
		return
	}
	w.(http.Flusher).Flush()
	LogSuccessfulResponse(r, body)
	return
}

func WriteSuccessfulJSONRaw(
	w http.ResponseWriter,
	r *http.Request,
	body interface{},
) {
	var internalErr error
	defer func() {
		if internalErr != nil {
			LogErrorResponse(r, internalErr, "")
		}
	}()

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	result, err := fastjson.EncodeJSON(body)
	if err != nil {
		WriteErrorJSON(w, r, ErrInternalServerError(), fmt.Sprintf("root cause: %s", err.Error()))
		return
	}
	w.Header().Set("Content-Length", strconv.Itoa(len(result)))
	_, err = w.Write(result)
	if err != nil {
		internalErr = stacktrace.Propagate(err, ErrInternalServerError().Error())
		return
	}
	w.(http.Flusher).Flush()
	LogSuccessfulResponse(r, body)
	return
}

// LogSuccessfulResponse logs a successful server response at level debug
// on standard logger
func LogSuccessfulResponse(r *http.Request, data interface{}) {
	logrus.WithFields(logrus.Fields{
		"host":        r.Host,
		"address":     r.RemoteAddr,
		"method":      r.Method,
		"request_uri": r.RequestURI,
		"proto":       r.Proto,
		"useragent":   r.UserAgent(),
	}).Debug(data)
}
