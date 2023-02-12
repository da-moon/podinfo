package core

import (
	"context"
	"net"
	"net/http"
	"sync"
	"time"

	"github.com/da-moon/northern-labs-interview/api"
	logger "github.com/da-moon/northern-labs-interview/internal/logger"
	metrics "github.com/da-moon/northern-labs-interview/sdk/api/metrics"
	mux "github.com/gorilla/mux"
	stacktrace "github.com/palantir/stacktrace"
)

// RestfulServer enables clients to start a graceful
// http server
type RestfulServer struct {
	mutex sync.RWMutex
	sync.Once
	config   *config
	listener net.Listener
	api      *http.Server
	router   *mux.Router
	metrics  *metrics.Metrics
	log      *logger.WrappedLogger
	stop     bool
	stopCh   chan struct{}
}

// RestfulServer spawns a goroutine
// that starts Podinfo Server and
// returns the struct
func (c *config) RestfulServer(
	ctx context.Context,
) (*RestfulServer, error) { // revive:disable:unexported-return
	var (
		err    error
		result = &RestfulServer{
			config: c,
			stop:   false,
			log:    c.Log(),
			stopCh: make(chan struct{}),
		}
	)

	err = result.initListener(ctx)
	if err != nil {
		err = stacktrace.Propagate(err, "restful-server failed at very early stage")
		return nil, err
	}
	result.log.Info("restful-server successfully bound to host port")
	r, err := result.config.Router()
	if err != nil {
		err = stacktrace.Propagate(err, "restful-server failed at very early stage")
		return nil, err
	}
	if r == nil {
		err = stacktrace.NewError("restful-server router initialization failed")
		return nil, err
	}

	result.log.Info("restful-server initializing NotFound route handler")
	r.NotFoundHandler = http.HandlerFunc(api.NotFoundHandler())
	result.log.Info("restful-server routers are ready to serve client requests")
	result.router = r
	result.Do(func() {
		result.log.Info("asynchronous API endpoint initialization started")
		go result.start(ctx)
	})
	return result, nil
	// revive:enable:unexported-return
}

// start sets up router and starts the
// graceful Rest server.
// NOTE(damoon) for now, this function does not
// need any arguments
func (s *RestfulServer) start(_ context.Context) {
	s.api = &http.Server{
		Addr:        s.listener.Addr().String(),
		Handler:     s.router,
		IdleTimeout: 90 * time.Second,
	}
	err := s.api.Serve(s.listener)
	if err != nil && err != http.ErrServerClosed {
		s.log.Error("restful-server start failed: %v", err)
		s.Shutdown()
		return
	}
	s.log.Info("restful-server has started accepting connection")
	return
}

// Shutdown gracefully terminates the http server
func (s *RestfulServer) Shutdown() {
	s.log.Trace("restful-server start Shutdown called")
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if s.stop {
		return
	}
	var err error
	s.stop = true
	close(s.stopCh)
	s.log.Info("restful-server gracefully tearing down api server")
	err = s.api.Shutdown(context.Background())
	if err != nil {
		s.log.Error(
			"restful-server graceful tear down of the api server failed: %v",
			err,
		)
	}
}

// ShutdownCh returns the servers shutdown channel
func (s *RestfulServer) ShutdownCh() <-chan struct{} {
	return s.stopCh
}
