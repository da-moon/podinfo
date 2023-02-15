package delay

import (
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/da-moon/northern-labs-interview/api/handlers/readiness"
	logger "github.com/da-moon/northern-labs-interview/internal/logger"
	response "github.com/da-moon/northern-labs-interview/sdk/api/response"
	mux "github.com/gorilla/mux"
	"github.com/palantir/stacktrace"
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
	code := http.StatusOK
	vars := mux.Vars(r)
	body := new(Response)
	defer func() {
		response.WriteJSON(
			w,
			r,
			code,
			make(map[string]string, 0),
			body,
		)
	}()
	delayString, ok := vars["seconds"]
	if !ok {
		code = http.StatusBadRequest
		msg := Name
		err := stacktrace.NewErrorWithCode(stacktrace.ErrorCode(code), "{second} variable was not included in the URI")
		response.LogErrorResponse(r, err, msg)
		body = nil
		return
	}
	delay, err := strconv.ParseUint(delayString, 10, 64)
	if err != nil {
		code = http.StatusBadRequest
		msg := Name
		err = stacktrace.PropagateWithCode(err, stacktrace.ErrorCode(code), "{second} section of the URI could not be parsed as an integer : %s", delayString)
		response.LogErrorResponse(r, err, msg)
		body = nil
		return
	}
	defer func() {
		response.LogSuccessfulResponse(r, body)
		return
	}()
	// Actual work starts here
	time.Sleep(time.Duration(delay) * time.Second)
	body.Delay = delay
	return
}
