package response

import (
	"net/http"

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
	defer func() {
		LogSuccessfulResponse(r, body)
		return
	}()
	WriteJSON(
		w,
		r,
		http.StatusOK,
		make(map[string]string, 0),
		&Response{
			Success: true,
			Body:    body,
		},
	)
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
