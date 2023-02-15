package disable

import (
	"net/http"
	"sync"

	readiness "github.com/da-moon/northern-labs-interview/api/handlers/readiness"
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

var HandlerFn = func(w http.ResponseWriter, r *http.Request) { //nolint:gochecknoglobals //this function is scoped only to this package
	defer func() {
		response.LogSuccessfulResponse(r, nil)
		return
	}()
	if readiness.Router.GetStatus() == readiness.OK {
		readiness.Router.SetStatus(readiness.Unavailable)
	}
	response.WriteJSON(
		w,
		r,
		http.StatusAccepted,
		make(map[string]string, 0),
		nil,
	)
	return
}
