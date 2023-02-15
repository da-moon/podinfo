package put

import (
	"io"
	"net/http"
	"sync"

	"github.com/da-moon/northern-labs-interview/api/handlers/cache/shared"
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
	var (
		code = http.StatusAccepted
		body = make([]byte, 0)
		vars = mux.Vars(r)
		ctx  = r.Context()
		err  error
	)
	defer func() {
		if err != nil {
			msg := Name
			response.LogErrorResponse(r, err, msg)
			return
		}
		response.LogSuccessfulResponse(r, body)
		return
	}()
	// grab the client
	client := shared.RedisClient(ctx, w, r)
	// RedisClient has sent the initial preflight check and
	// responded in case redis was not accessible
	if client == nil {
		err = stacktrace.NewError("redis client is not ready")
		code = http.StatusInternalServerError
		return
	}
	defer func() {
		response.Write(
			w,
			r,
			code,
			make(map[string]string, 0),
			nil,
		)
	}()
	// Actual work starts here
	key, ok := vars["key"]
	if !ok {
		err = stacktrace.NewErrorWithCode(stacktrace.ErrorCode(code), "{key} variable was not included in the URI")
		code = http.StatusBadRequest
		return
	}
	body, err = io.ReadAll(r.Body)
	if err != nil {
		code = http.StatusBadRequest
		err = stacktrace.PropagateWithCode(err, stacktrace.ErrorCode(code), "failed to read request body")
		return
	}
	err = client.Set(ctx, key, body, 0).Err()
	if err != nil {
		code = http.StatusInternalServerError
		err = stacktrace.PropagateWithCode(err, stacktrace.ErrorCode(code), "failed to set key in redis")
		return
	}
	return
}
