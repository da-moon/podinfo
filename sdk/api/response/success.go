package response

import (
	"net/http"
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
	e := LogEntry(r)
	if e != nil {
		e.Debug(data)
	}
}
