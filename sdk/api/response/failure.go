package response

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/palantir/stacktrace"
	"github.com/sirupsen/logrus"
)

// WriteErrorJSON - logs and sends a json response to the client
// showing the error message
// TODO: this method must retry ops on failure
// TODO: turn this into a receiver of Response
func WriteErrorJSON(
	w http.ResponseWriter,
	r *http.Request,
	e HTTPError,
	msg string,
) {
	defer func() {
		LogErrorResponse(r, e, msg)
	}()
	var internalErr error
	defer func() {
		if internalErr != nil {
			LogErrorResponse(r, internalErr, "")
		}
	}()
	var err error
	code := e.StatusCode()

	response, err := e.EncodeJSON()
	if err != nil {
		internalErr = stacktrace.Propagate(err, ErrInternalServerError().Error())
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Content-Length", strconv.Itoa(len(response)))
	w.WriteHeader(int(code))
	if response != nil {
		_, err = w.Write(response)
		if err != nil {
			internalErr = stacktrace.Propagate(err, ErrInternalServerError().Error())
			return
		}
		w.(http.Flusher).Flush()
	}
	return
}

// LogErrorResponse - logs an errornous server server at level debug
// on standard logger
func LogErrorResponse(r *http.Request, err error, msg string) {
	fmt.Println("err", err)
	fmt.Println("msg", msg)
	stacktrace.DefaultFormat = stacktrace.FormatBrief
	logrus.WithFields(logrus.Fields{
		"host":        r.Host,
		"address":     r.RemoteAddr,
		"method":      r.Method,
		"request_uri": r.RequestURI,
		"proto":       r.Proto,
		"useragent":   r.UserAgent(),
	}).WithError(err).Debug(msg)
}
