package response

import (
	"net/http"

	"github.com/palantir/stacktrace"
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
	// fmt.Printf("*** err %s\n", err)
	// fmt.Printf("*** msg %s\n", msg)
	e := LogEntry(r)
	if e != nil {
		stacktrace.DefaultFormat = stacktrace.FormatBrief
		e.WithError(err).Error(msg)
	}
}
