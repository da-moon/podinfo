package cache

import (
	"sync"
	"sync/atomic"

	"github.com/da-moon/podinfo/api/handlers/cache/delete"
	"github.com/da-moon/podinfo/api/handlers/cache/get"
	"github.com/da-moon/podinfo/api/handlers/cache/post"
	"github.com/da-moon/podinfo/api/handlers/cache/put"
	shared "github.com/da-moon/podinfo/api/handlers/cache/shared"
	logger "github.com/da-moon/podinfo/internal/logger"
	stacktrace "github.com/palantir/stacktrace"
	redis "github.com/redis/go-redis/v9"
)

var (
	initialized atomic.Bool
	once        sync.Once
)

// Init initializes the shared state of this api group
// This function will be called
// from the server subcommand library
func Init(l *logger.WrappedLogger, dev bool, opt *redis.Options) error {
	if !dev {
		if l == nil {
			return stacktrace.NewError("logger is nil")
		}
		if opt == nil {
			return stacktrace.NewError("Redis Client options is nil")
		}
	}
	once.Do(func() {
		shared.Group.SetLogger(l)
		shared.Group.SetRedisOptions(opt)
		initialized.Store(true)
		return
	})
	return nil
}

// Register adds all endpoints of this api group to the
// router registry
func Register() error {
	if !initialized.Load() {
		err := stacktrace.NewError("%s api group (%s) has not been initialized", shared.Name, shared.GroupPrefix)
		return err
	}
	l := shared.Group.GetLogger()
	if l == nil {
		return stacktrace.NewError("logger is nil")
	}
	// POST /cache/{key}
	post.Router.SetLogger(l)
	err := post.Router.Register()
	if err != nil {
		err = stacktrace.Propagate(err, "failed to initialize [ POST ] HTTP request handlers for '%s' (%s%s)", post.Name, shared.GroupPrefix, post.Path)
		return err
	}
	// PUT /cache/{key}
	put.Router.SetLogger(l)
	err = put.Router.Register()
	if err != nil {
		err = stacktrace.Propagate(err, "failed to initialize [ PUT ] HTTP request handlers for '%s' (%s%s)", put.Name, shared.GroupPrefix, put.Path)
		return err
	}
	// GET /cache/{key}
	get.Router.SetLogger(l)
	err = get.Router.Register()
	if err != nil {
		err = stacktrace.Propagate(err, "failed to initialize [ GET ] HTTP request handlers for '%s' (%s%s)", get.Name, shared.GroupPrefix, get.Path)
		return err
	}
	// DELETE /cache/{key}
	delete.Router.SetLogger(l)
	err = delete.Router.Register()
	if err != nil {
		err = stacktrace.Propagate(err, "failed to initialize [ DELETE ] HTTP request handlers for '%s' (%s%s)", delete.Name, shared.GroupPrefix, delete.Path)
		return err
	}
	return nil
}
