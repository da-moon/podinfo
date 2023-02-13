package response

import (
	"fmt"
	"net/http"

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
	WriteJSON(
		w,
		r,
		int(e.StatusCode()),
		make(map[string]string, 0),
		&Response{
			Success: false,
			Body: struct {
				Msg string `json:"msg"`
			}{
				Msg: e.Error(),
			},
		},
	)
	return
}

// LogErrorResponse logs an erroneous server server at level debug
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
