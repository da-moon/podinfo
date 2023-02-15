package readiness

import (
	"bytes"
	"encoding/json"
	"net/http"
	"sync"
	"sync/atomic"

	logger "github.com/da-moon/northern-labs-interview/internal/logger"
	response "github.com/da-moon/northern-labs-interview/sdk/api/response"
	stacktrace "github.com/palantir/stacktrace"
)

// handler struct encapsulates the state this API endpoint
// handler needs
type handler struct {
	// mutex for guard shared state
	mutex sync.RWMutex
	// log is the logger for this handler
	log *logger.WrappedLogger
	// status represents server readiness status
	status atomic.Bool
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

// SetStatus sets the readiness status for this handler
// mutex is not used here as atomic.Bool is more performant
func (h *handler) SetStatus(status Status) {
	if status == OK {
		h.status.Store(true)
		return
	}
	h.status.Store(false)
}

// GetStatus returns the readiness status for this handler
func (h *handler) GetStatus() Status {
	if h.status.Load() {
		return OK
	}
	return Unavailable
}

// HandlerFn function handles incoming HTTP request
// it satisfies golang's stdlib
// request handler interface (http.HandlerFunc)
var HandlerFn = func(w http.ResponseWriter, r *http.Request) { //nolint:gochecknoglobals //this function is scoped only to this package
	s := Router.GetStatus()
	code := http.StatusOK
	body := &Response{
		Status: s.String(),
	}
	loggerfn := func(rr *http.Request, data interface{}) func() {
		return func() {
			// NOTE: this is just to stay safe from nil pointer dereference
			if data != nil {
				response.LogSuccessfulResponse(rr, data)
			}
			return
		}
	}
	if s != OK {
		code = http.StatusServiceUnavailable
		loggerfn = func(rr *http.Request, data interface{}) func() {
			return func() {
				// NOTE: this is just to stay safe from nil pointer dereference
				if data != nil {
					compact := &bytes.Buffer{}
					// NOTE: this is just to stay safe from nil pointer dereference
					resp, _ := json.Marshal(data) //nolint:gosec // failure here does not matter
					if resp != nil {
						json.Compact(compact, resp) //nolint:gosec // failure here does not matter
						// msg := fmt.Sprintf("%s : %s", Name, Path)
						err := stacktrace.NewErrorWithCode(stacktrace.ErrorCode(code), compact.String())
						response.LogErrorResponse(rr, err, "")
					}
				}
				return
			}
		}
	}
	defer loggerfn(r, body)()
	response.WriteJSON(
		w,
		r,
		code,
		make(map[string]string, 0),
		body,
	)
}
