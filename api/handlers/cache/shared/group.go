package shared

import (
	"sync"

	logger "github.com/da-moon/podinfo/internal/logger"
	redis "github.com/redis/go-redis/v9"
)

const (
	// Name stores a human friendly, param-cased
	// unique identifier name for this api group
	Name = "redis-proxy"
	// GroupPrefix stores API group (prefix) for this URI
	// the full URI is /<prefix>/<path>
	GroupPrefix = "/cache"
)

var (
	// Group is used from outside this package to register
	// all the endpoints in this API group
	Group = &group{} //nolint:gochecknoglobals // safe as it has mutex guard and access internal state is through getter/setter functions
)

// group struct encapsulates the state of this API group
type group struct {
	// mutex for guard shared state
	mutex sync.RWMutex
	// los is the logger for this handler
	log *logger.WrappedLogger
	// a redis client options, which is used
	// to generate a new redis client
	// for each request handler.
	// A request handler will use this to get it's self
	// a dedicated client
	redisOptions *redis.Options
}

// SetLogger sets the logger for handlers under this API group
func (g *group) SetLogger(l *logger.WrappedLogger) {
	g.mutex.Lock()
	defer g.mutex.Unlock()
	g.log = l
}

// GetLogger returns the logger for handlers under this API group
func (g *group) GetLogger() *logger.WrappedLogger {
	g.mutex.RLock()
	defer g.mutex.RUnlock()
	return g.log
}

// SetRedisOptions sets the redis client options for handlers under this API group
func (g *group) SetRedisOptions(opts *redis.Options) {
	g.mutex.Lock()
	defer g.mutex.Unlock()
	g.redisOptions = opts
}

// GetRedisOptions returns the redis client options for handlers under this API group
func (g *group) GetRedisOptions() *redis.Options {
	g.mutex.RLock()
	defer g.mutex.RUnlock()
	return g.redisOptions
}
