package env

import (
	"net/http"
	"os"
	"sync"

	"github.com/da-moon/northern-labs-interview/api/handlers/readiness"
	logger "github.com/da-moon/northern-labs-interview/internal/logger"
	response "github.com/da-moon/northern-labs-interview/sdk/api/response"
)

// handler struct encapsulates the state this API endpoint
// handler needs
type handler struct {
	// mutex for guard shared state
	mutex sync.RWMutex
	// log is the logger for this handler
	log *logger.WrappedLogger
}

// SetLogger sets the logger for this handler
func (h *handler) SetLogger(l *logger.WrappedLogger) {
	h.mutex.Lock()
	defer h.mutex.Unlock()
	h.log = l
}

// GetLogger returns the logger for this handler
func (h *handler) GetLogger() *logger.WrappedLogger {
	h.mutex.RLock()
	defer h.mutex.RUnlock()
	return h.log
}

// HandlerFn function handles incoming HTTP request
// it satisfies golang's stdlib
// request handler interface (http.HandlerFunc)
var HandlerFn = func(w http.ResponseWriter, r *http.Request) { //nolint:gochecknoglobals //this function is scoped only to this package
	if readiness.Router.GetStatus() == readiness.Unavailable {
		readiness.HandlerFn(w, r)
		return
	}
	var body Response = os.Environ()
	defer func() {
		response.LogSuccessfulResponse(r, body)
		return
	}()
	response.WriteJSON(
		w,
		r,
		http.StatusOK,
		make(map[string]string, 0),
		body,
	)
	return
}