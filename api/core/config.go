package core

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	logger "github.com/da-moon/northern-labs-interview/internal/logger"
	metrics "github.com/da-moon/northern-labs-interview/sdk/api/metrics"
	port "github.com/da-moon/northern-labs-interview/sdk/api/port"
	mapstructure "github.com/mitchellh/mapstructure"
	stacktrace "github.com/palantir/stacktrace"
)

// config is the configuration for the agent.
//
//go:generate gomodifytags -override -file $GOFILE -struct config -add-tags json,yaml,mapstructure -w -transform snakecase
type Config struct {
	mutex           sync.RWMutex          `json:"lock" yaml:"lock" mapstructure:"lock"`
	log             *logger.WrappedLogger `json:"log" yaml:"log" mapstructure:"log"`
	DevelopmentMode bool                  `mapstructure:"development_mode" json:"development_mode,omitempty" yaml:"development_mode"`
	NodeName        string                `mapstructure:"node_name" json:"node_name,omitempty" yaml:"node_name"`
	APIAddr         string                `mapstructure:"api_addr" json:"api_addr,omitempty" yaml:"api_addr"`
	// ─── METRICS ────────────────────────────────────────────────────────────────────
	MetricsPrefix           string        `mapstructure:"metrics_prefix" json:"metrics_prefix" yaml:"metrics_prefix"`
	StatsiteAddr            string        `mapstructure:"statsite_addr" json:"statsite_addr,omitempty" yaml:"statsite_addr"`
	StatsdAddr              string        `mapstructure:"statsd_addr" json:"statsd_addr,omitempty" yaml:"statsd_addr"`
	PrometheusRetentionTime time.Duration `mapstructure:"prometheus_retention_time" json:"prometheus_retention_time" yaml:"prometheus_retention_time"`
	// ─── REDIS ───────────────────────────────────────────────────────────────────
	// host:port address.
	RedisAddr string
	// ClientName will execute the `CLIENT SETNAME ClientName` command for each conn.
	RedisClientName string
	// Use the specified Username to authenticate the current connection
	// with one of the connections defined in the ACL list when connecting
	// to a Redis 6.0 instance, or greater, that is using the Redis ACL system.
	RedisUsername string
	// Optional password. Must match the password specified in the
	// requirepass server configuration option (if connecting to a Redis 5.0 instance, or lower),
	// or the User Password when connecting to a Redis 6.0 instance, or greater,
	// that is using the Redis ACL system.
	RedisPassword string
	// Database to be selected after connecting to the server.
	RedisDB int
	// Maximum number of retries before giving up.
	// -1 disables retries.
	RedisMaxRetries int
	// Minimum backoff between each retry.
	// -1 disables backoff.
	RedisMinRetryBackoff time.Duration
	// Maximum backoff between each retry.
	// -1 disables backoff.
	RedisMaxRetryBackoff time.Duration
	// Dial timeout for establishing new connections.
	RedisDialTimeout time.Duration
	// Timeout for socket reads. If reached, commands will fail
	// with a timeout instead of blocking. Supported values:
	//   - `0` - default timeout (3 seconds).
	//   - `-1` - no timeout (block indefinitely).
	//   - `-2` - disables SetReadDeadline calls completely.
	RedisReadTimeout time.Duration
	// Timeout for socket writes. If reached, commands will fail
	// with a timeout instead of blocking.  Supported values:
	//   - `0` - default timeout (3 seconds).
	//   - `-1` - no timeout (block indefinitely).
	//   - `-2` - disables SetWriteDeadline calls completely.
	RedisWriteTimeout time.Duration
	// ContextTimeoutEnabled controls whether the client respects context timeouts and deadlines.
	// See https://redis.uptrace.dev/guide/go-redis-debugging.html#timeouts
	RedisContextTimeoutEnabled bool
	// Type of connection pool.
	// true for FIFO pool, false for LIFO pool.
	// Note that FIFO has slightly higher overhead compared to LIFO,
	// but it helps closing idle connections faster reducing the pool size.
	RedisPoolFIFO bool
	// Maximum number of socket connections.
	RedisPoolSize int
	// Amount of time client waits for connection if all connections
	// are busy before returning an error.
	RedisPoolTimeout time.Duration
	// Minimum number of idle connections which is useful when establishing
	// new connection is slow.
	RedisMinIdleConns int
	// Maximum number of idle connections.
	RedisMaxIdleConns int
	// ConnMaxIdleTime is the maximum amount of time a connection may be idle.
	// Should be less than server's timeout.
	RedisConnMaxIdleTime time.Duration
	// ConnMaxLifetime is the maximum amount of time a connection may be reused.
	RedisConnMaxLifetime time.Duration
}

// Mergeconfig takes in two config objects
// and merges them into one, giving precedence to the second config.
// nolint:gocognit,gocyclo // this function is well tested and won't be used often
func Mergeconfig(a, b *Config) *Config { // revive:disable:unexported-return
	result := *a
	result.DevelopmentMode = b.DevelopmentMode
	if b.NodeName != "" {
		result.NodeName = b.NodeName
	}
	if b.APIAddr != "" {
		result.APIAddr = b.APIAddr
	}
	// ─── TELEMETRY ───────────────────────────────────────────────────────────────
	if b.MetricsPrefix != "" {
		result.MetricsPrefix = b.MetricsPrefix
	}
	if b.StatsiteAddr != "" {
		result.StatsiteAddr = b.StatsiteAddr
	}
	if b.StatsdAddr != "" {
		result.StatsdAddr = b.StatsdAddr
	}
	if b.PrometheusRetentionTime.Nanoseconds() >= 1 {
		result.PrometheusRetentionTime = b.PrometheusRetentionTime
	}
	// ─── REDIS
	// ───────────────────────────────────────────────────────────────────
	if b.RedisAddr != "" {
		result.RedisAddr = b.RedisAddr
	}
	if b.RedisClientName != "" {
		result.RedisClientName = b.RedisClientName
	}
	if b.RedisUsername != "" {
		result.RedisUsername = b.RedisUsername
	}
	if b.RedisPassword != "" {
		result.RedisPassword = b.RedisPassword
	}
	result.RedisDB = b.RedisDB
	if b.RedisMaxRetries >= 0 {
		result.RedisMaxRetries = b.RedisMaxRetries
	}
	if b.RedisMinRetryBackoff.Nanoseconds() >= -1 {
		result.RedisMinRetryBackoff = b.RedisMinRetryBackoff
	}
	if b.RedisMaxRetryBackoff.Nanoseconds() >= -1 {
		result.RedisMaxRetryBackoff = b.RedisMaxRetryBackoff
	}
	if b.RedisDialTimeout.Nanoseconds() >= 0 {
		result.RedisDialTimeout = b.RedisDialTimeout
	}
	if b.RedisReadTimeout.Nanoseconds() >= -2 {
		result.RedisReadTimeout = b.RedisReadTimeout
	}
	if b.RedisWriteTimeout.Nanoseconds() >= -2 {
		result.RedisWriteTimeout = b.RedisWriteTimeout
	}
	result.RedisContextTimeoutEnabled = b.RedisContextTimeoutEnabled
	result.RedisPoolFIFO = b.RedisPoolFIFO
	if b.RedisPoolSize >= 0 {
		result.RedisPoolSize = b.RedisPoolSize
	}
	if b.RedisPoolTimeout >= 0 {
		result.RedisPoolTimeout = b.RedisPoolTimeout
	}
	if b.RedisMinIdleConns != 0 {
		result.RedisMinIdleConns = b.RedisMinIdleConns
	}
	if b.RedisMaxIdleConns >= 0 {
		result.RedisMaxIdleConns = b.RedisMaxIdleConns
	}
	if b.RedisConnMaxIdleTime.Nanoseconds() >= 0 {
		result.RedisConnMaxIdleTime = b.RedisConnMaxIdleTime
	}
	result.RedisConnMaxLifetime = b.RedisConnMaxLifetime
	return &result
	// revive:enable:unexported-return
}

//
// ──────────────────────────────────────────────────────────────────── I ──────────
//   :::::: D E F A U L T   V A L U E S : :  :   :    :     :        :          :
// ──────────────────────────────────────────────────────────────────────────────
//

// DefaultConfig returns a new config struct
// nolint:gocognit,gocyclo // this function is well tested and won't be used often
func DefaultConfig(log *logger.WrappedLogger) (*Config, error) { // revive:disable:unexported-return
	if log == nil {
		err := stacktrace.NewError("no logger was provided")
		return nil, err
	}
	// ─── DEVELOPMENT MODE DEFAULT ───────────────────────────────────────────────────
	developmentMode := DefaultDevelopmentMode()
	// ─── NODE NAME DEFAULT ──────────────────────────────────────────────────────────
	nodeName, err := DefaultNodeName()
	if err != nil {
		err = stacktrace.Propagate(err, "cannot prepare default api config struct")
		return nil, err
	}
	// ─── API ADDRESS DEFAULT ─────────────────────────────────────────────────────────
	apiAddr, err := DefaultAPIAddr()
	if err != nil {
		err = stacktrace.Propagate(err, "cannot prepare default api config struct")
		return nil, err
	}
	// // ─── REDIS DEFAULTS ──────────────────────────────────────────────────────────
	// redisAddr, err := DefaultRedisAddr()
	// if err != nil {
	// 	err = stacktrace.Propagate(err, "cannot prepare default api config struct")
	// 	return nil, err
	// }
	// redisClientName, err := DefaultRedisClientName()
	// if err != nil {
	// 	err = stacktrace.Propagate(err, "cannot prepare default api config struct")
	// 	return nil, err
	// }
	// redisUsername, err := DefaultRedisUsername()
	// if err != nil {
	// 	err = stacktrace.Propagate(err, "cannot prepare default api config struct")
	// 	return nil, err
	// }
	// redisPassword, err := DefaultRedisPassword()
	// if err != nil {
	// 	err = stacktrace.Propagate(err, "cannot prepare default api config struct")
	// 	return nil, err
	// }
	// redisDB, err := DefaultRedisDB()
	// if err != nil {
	// 	err = stacktrace.Propagate(err, "cannot prepare default api config struct")
	// 	return nil, err
	// }
	// redisMaxRetries, err := DefaultRedisMaxRetries()
	// if err != nil {
	// 	err = stacktrace.Propagate(err, "cannot prepare default api config struct")
	// 	return nil, err
	// }
	// redisMinRetryBackoff, err := DefaultRedisMinRetryBackoff()
	// if err != nil {
	// 	err = stacktrace.Propagate(err, "cannot prepare default api config struct")
	// 	return nil, err
	// }
	// redisMaxRetryBackoff, err := DefaultRedisMaxRetryBackoff()
	// if err != nil {
	// 	err = stacktrace.Propagate(err, "cannot prepare default api config struct")
	// 	return nil, err
	// }
	// redisDialTimeout, err := DefaultRedisDialTimeout()
	// if err != nil {
	// 	err = stacktrace.Propagate(err, "cannot prepare default api config struct")
	// 	return nil, err
	// }
	// redisReadTimeout, err := DefaultRedisReadTimeout()
	// if err != nil {
	// 	err = stacktrace.Propagate(err, "cannot prepare default api config struct")
	// 	return nil, err
	// }
	// redisWriteTimeout, err := DefaultRedisWriteTimeout()
	// if err != nil {
	// 	err = stacktrace.Propagate(err, "cannot prepare default api config struct")
	// 	return nil, err
	// }
	// redisContextTimeoutEnabled, err := DefaultRedisContextTimeoutEnabled()
	// if err != nil {
	// 	err = stacktrace.Propagate(err, "cannot prepare default api config struct")
	// 	return nil, err
	// }
	// redisPoolFIFO, err := DefaultRedisPoolFIFO()
	// if err != nil {
	// 	err = stacktrace.Propagate(err, "cannot prepare default api config struct")
	// 	return nil, err
	// }
	// redisPoolSize, err := DefaultRedisPoolSize()
	// if err != nil {
	// 	err = stacktrace.Propagate(err, "cannot prepare default api config struct")
	// 	return nil, err
	// }
	// redisPoolTimeout, err := DefaultRedisPoolTimeout()
	// if err != nil {
	// 	err = stacktrace.Propagate(err, "cannot prepare default api config struct")
	// 	return nil, err
	// }
	// redisMinIdleConns, err := DefaultRedisMinIdleConns()
	// if err != nil {
	// 	err = stacktrace.Propagate(err, "cannot prepare default api config struct")
	// 	return nil, err
	// }
	// redisMaxIdleConns, err := DefaultRedisMaxIdleConns()
	// if err != nil {
	// 	err = stacktrace.Propagate(err, "cannot prepare default api config struct")
	// 	return nil, err
	// }
	// redisConnMaxIdleTime, err := DefaultRedisConnMaxIdleTime()
	// if err != nil {
	// 	err = stacktrace.Propagate(err, "cannot prepare default api config struct")
	// 	return nil, err
	// }
	// redisConnMaxLifetime, err := DefaultRedisConnMaxLifetime()
	// if err != nil {
	// 	err = stacktrace.Propagate(err, "cannot prepare default api config struct")
	// 	return nil, err
	// }
	// ─── RESULT ─────────────────────────────────────────────────────────────────────
	// ────────────────────────────────────────────────────────────────────────────────
	// ────────────────────────────────────────────────────────────────────────────────
	result := &Config{
		log:             log,
		DevelopmentMode: developmentMode,
		NodeName:        nodeName,
		APIAddr:         apiAddr,
		// ─── REDIS ────────────────────────────────────────────────────
		// RedisAddr:                  redisAddr,
		// RedisClientName:            redisClientName,
		// RedisUsername:              redisUsername,
		// RedisPassword:              redisPassword,
		// RedisDB:                    redisDB,
		// RedisMaxRetries:            redisMaxRetries,
		// RedisMinRetryBackoff:       redisMinRetryBackoff,
		// RedisMaxRetryBackoff:       redisMaxRetryBackoff,
		// RedisDialTimeout:           redisDialTimeout,
		// RedisReadTimeout:           redisReadTimeout,
		// RedisWriteTimeout:          redisWriteTimeout,
		// RedisContextTimeoutEnabled: redisContextTimeoutEnabled,
		// RedisPoolFIFO:              redisPoolFIFO,
		// RedisPoolSize:              redisPoolSize,
		// RedisPoolTimeout:           redisPoolTimeout,
		// RedisMinIdleConns:          redisMinIdleConns,
		// RedisMaxIdleConns:          redisMaxIdleConns,
		// RedisConnMaxIdleTime:       redisConnMaxIdleTime,
		// RedisConnMaxLifetime:       redisConnMaxLifetime,
		// ─── METRICS ─────────────────────────────────────────────────────
		MetricsPrefix:           metrics.DefaultMetricsPrefix(),
		StatsiteAddr:            metrics.DefaultStatsiteAddr(),
		StatsdAddr:              metrics.DefaultStatsdAddr(),
		PrometheusRetentionTime: metrics.DefaultPrometheusRetentionTime(),
	}
	return result, nil
	// revive:enable:unexported-return
}

// DefaultDevelopmentMode returns the default value for the development mode
func DefaultDevelopmentMode() bool {
	var result bool
	podinfoDevelString := os.Getenv("PODINFO_DEVEL")
	if podinfoDevelString != "" {
		var err error
		result, err = strconv.ParseBool(podinfoDevelString)
		if err != nil {
			result = false
		}
	}
	return result
}

// DefaultNodeName returns the default value for the node name
func DefaultNodeName() (string, error) {
	var err error
	result := os.Getenv("PODINFO_NODE_NAME")
	if result == "" {
		result, err = os.Hostname()
		if err != nil {
			err = stacktrace.Propagate(err, "cannot get default node name")
			return "", err
		}
	}
	return result, nil
}

// DefaultAPIAddr returns the default value for the api address
// TODO: add validation to ensure address is valid when the user provided the address
func DefaultAPIAddr() (string, error) {
	apiAddr := os.Getenv("PODINFO_API_ADDR")
	if apiAddr == "" {
		tcpAddr, err := port.New().TCP()
		if err != nil {
			err = stacktrace.Propagate(err, "cannot find a random port to bind to")
			return "", err
		}
		apiAddr = fmt.Sprintf("0.0.0.0:%s", tcpAddr.Port)
	}
	return apiAddr, nil
}

//
// ──────────────────────────────────────────────────── I ──────────
//   :::::: G E T T E R : :  :   :    :     :        :          :
// ──────────────────────────────────────────────────────────────
//

// Decodeconfig takes an io.reader and decodes
// underlying byte stream into a config struct
// nolint:gocognit // This is a well tested function
func Decodeconfig(r io.Reader) (*Config, error) { // revive:disable:unexported-return
	var raw interface{}
	dec := json.NewDecoder(r)
	if err := dec.Decode(&raw); err != nil {
		return nil, err
	}
	var md mapstructure.Metadata
	var result Config
	msdec, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		Metadata:    &md,
		Result:      &result,
		ErrorUnused: true,
	})
	if err != nil {
		return nil, err
	}
	if err := msdec.Decode(raw); err != nil {
		return nil, err
	}
	return &result, nil
	// revive:enable:unexported-return
}

// Log returns the underlying log writer
func (c *Config) Log() *logger.WrappedLogger {
	c.mutex.RLock()
	return c.log
}

// ──────────────────────────────────────────────────── I ──────────
//
//	:::::: S E T T E R : :  :   :    :     :        :          :
//
// ──────────────────────────────────────────────────────────────
func (c *Config) SetDevelopmentMode() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.DevelopmentMode = true
}

func (c *Config) SetNodeName(value string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if value != "" {
		c.NodeName = value
	}
}
func (c *Config) SetAPIAddr(value string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if value != "" {
		c.APIAddr = value
	}
}
func (c *Config) SetMetricsPrefix(value string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if value != "" {
		c.MetricsPrefix = value
	}
}
func (c *Config) SetStatsiteAddr(value string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if value != "" {
		c.StatsiteAddr = value
	}
}
func (c *Config) SetStatsdAddr(value string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if value != "" {
		c.StatsdAddr = value
	}
}
func (c *Config) SetPrometheusRetentionTime(value time.Duration) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if value > 0 {
		c.PrometheusRetentionTime = value
	}
}

//
// ────────────────────────────────────────────────────────────── I ──────────
//   :::::: F I L E   L O A D E R : :  :   :    :     :        :          :
// ────────────────────────────────────────────────────────────────────────
//

// ReadconfigPaths takes in an array of file (json)
// paths and returns a config struct.
// TODO: read configs in parallel with help of a waitgroup as calling defer in loop might lead to bugs
// nolint:gocognit,gocyclo // I can't think of any ways to lower complexity
func ReadconfigPaths(paths []string) (*Config, error) { // revive:disable:unexported-return,defer
	result := &Config{
		log: logger.DefaultWrappedLogger(string(logger.ErrorLevel)),
	}
	for _, p := range paths {
		path := filepath.Clean(p)
		f, err := os.Open(path)
		if err != nil {
			return nil, stacktrace.Propagate(err, "error reading '%s'", path)
		}
		// note(damoon) this may cause the application to open many
		// file descriptors which might cause the kernel to panic
		if f != nil {
			defer func() {
				err = f.Close()
				if err != nil {
					result.log.Error("cannot close file:%v", err)
				}
			}()
		}
		fi, err := f.Stat()
		if err != nil {
			return nil, stacktrace.Propagate(err, "error reading '%s'", path)
		}
		if !fi.IsDir() {
			dec := new(Config) // nolint:staticcheck // SA4006 dec is used right away
			dec, err = Decodeconfig(f)
			if err != nil {
				return nil, stacktrace.Propagate(err, "error decoding '%s'", path)
			}
			result = Mergeconfig(result, dec)
			continue
		}
		contents, err := f.Readdir(-1)
		if err != nil {
			return nil, stacktrace.Propagate(err, "error reading '%s'", path)
		}
		sort.Sort(dirEnts(contents))
		for _, fi := range contents {
			if fi.IsDir() {
				continue
			}
			if !strings.HasSuffix(fi.Name(), ".json") {
				continue
			}
			subpath := filepath.Clean(
				filepath.Join(path, fi.Name()),
			)
			ff, err := os.Open(subpath)
			if err != nil {
				return nil, stacktrace.Propagate(err, "error reading '%s'", subpath)
			}
			if ff != nil {
				defer func() {
					err = ff.Close()
					if err != nil {
						result.log.Error("cannot close file:%v", err)
					}
				}()
			}
			config, err := Decodeconfig(ff)
			if err != nil {
				return nil, stacktrace.Propagate(err, "error decoding '%s'", subpath)
			}
			result = Mergeconfig(result, config)
		}
	}
	return result, nil
	// revive:enable:unexported-return,defer
}

// ─── SORT INTERFACE ─────────────────────────────────────────────────────────────
type dirEnts []os.FileInfo

// Len implements sort.Interface.
func (d dirEnts) Len() int { return len(d) }

// Less implements sort.Interface
func (d dirEnts) Less(i, j int) bool { return d[i].Name() < d[j].Name() }

// Swap implement sort.Interface
func (d dirEnts) Swap(i, j int) { d[i], d[j] = d[j], d[i] }

// ────────────────────────────────────────────────────────────────────────────────
