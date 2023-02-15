package core

import (
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"

	stacktrace "github.com/palantir/stacktrace"
)

//
// ──────────────────────────────────────────────────────────────────── I ──────────
//   :::::: D E F A U L T   V A L U E S : :  :   :    :     :        :          :
// ──────────────────────────────────────────────────────────────────────────────
//

// DefaultRedisAddr returns the default redis address
func DefaultRedisAddr() (string, error) {
	result := strings.TrimSpace(os.Getenv("PODINFO_REDIS_ADDR"))
	if result == "" {
		result = "0.0.0.0:6379"
	}
	return result, nil
}

// DefaultRedisClientName returns the default redis client name
func DefaultRedisClientName() (string, error) {
	result := os.Getenv("PODINFO_REDIS_CLIENT_NAME")
	if len(result) == 0 {
		result = "podinfo"
	}
	return result, nil
}

// DefaultRedisUsername returns the default redis username
func DefaultRedisUsername() (string, error) {
	result := strings.TrimSpace(os.Getenv("PODINFO_REDIS_USERNAME"))
	var err error
	if result == "" {
		err = stacktrace.NewError("PODINFO_REDIS_USERNAME is not set")
	}
	return result, err
}

// DefaultRedisPassword returns the default redis password
func DefaultRedisPassword() (string, error) {
	result := strings.TrimSpace(os.Getenv("PODINFO_REDIS_PASSWORD"))
	var err error
	if result == "" {
		err = stacktrace.NewError("PODINFO_REDIS_PASSWORD is not set")
	}
	return result, err
}

// DefaultRedisDB returns the default redis db
func DefaultRedisDB() (int, error) {
	var (
		result = 0
		err    error
	)
	resultStr := strings.TrimSpace(os.Getenv("PODINFO_REDIS_DB"))
	if resultStr != "" {
		result, err = strconv.Atoi(resultStr)
		if err != nil {
			err = stacktrace.Propagate(err, "cannot parse PODINFO_REDIS_DB")
		}

	}
	return result, err
}

// DefaultRedisMaxRetries returns the default redis max retries
func DefaultRedisMaxRetries() (int, error) {
	var (
		result = 3
		err    error
	)
	resultStr := strings.TrimSpace(os.Getenv("PODINFO_REDIS_MAX_RETRIES"))
	if resultStr != "" {
		result, err = strconv.Atoi(resultStr)
		if err != nil {
			err = stacktrace.Propagate(err, "cannot parse PODINFO_REDIS_MAX_RETRIES")
		}

	}
	return result, err
}

// DefaultRedisMinRetryBackoff returns the default redis min retry backoff
func DefaultRedisMinRetryBackoff() (time.Duration, error) {
	var (
		result = 8 * time.Millisecond
		err    error
	)
	resultStr := strings.TrimSpace(os.Getenv("PODINFO_REDIS_MIN_RETRY_BACKOFF"))
	if resultStr != "" {
		result, err = time.ParseDuration(resultStr)
		if err != nil {
			err = stacktrace.Propagate(err, "cannot parse PODINFO_REDIS_MIN_RETRY_BACKOFF")
		}
	}
	return result, err
}

// DefaultRedisMaxRetryBackoff returns the default redis max retry backoff
func DefaultRedisMaxRetryBackoff() (time.Duration, error) {
	var (
		result = 512 * time.Millisecond
		err    error
	)
	resultStr := strings.TrimSpace(os.Getenv("PODINFO_REDIS_MAX_RETRY_BACKOFF"))
	if resultStr != "" {
		result, err = time.ParseDuration(resultStr)
		if err != nil {
			err = stacktrace.Propagate(err, "cannot parse PODINFO_REDIS_MAX_RETRY_BACKOFF")
		}
	}
	return result, err
}

// DefaultRedisDialTimeout returns the default redis dial timeout
func DefaultRedisDialTimeout() (time.Duration, error) {
	var (
		result = 5 * time.Second
		err    error
	)
	resultStr := strings.TrimSpace(os.Getenv("PODINFO_REDIS_DIAL_TIMEOUT"))
	if resultStr != "" {
		result, err = time.ParseDuration(resultStr)
		if err != nil {
			err = stacktrace.Propagate(err, "cannot parse PODINFO_REDIS_DIAL_TIMEOUT")
		}

	}
	return result, err
}

// DefaultRedisReadTimeout returns the default redis read timeout
func DefaultRedisReadTimeout() (time.Duration, error) {
	var (
		result = 3 * time.Second
		err    error
	)
	resultStr := strings.TrimSpace(os.Getenv("PODINFO_REDIS_READ_TIMEOUT"))
	if resultStr != "" {
		result, err = time.ParseDuration(resultStr)
		if err != nil {
			err = stacktrace.Propagate(err, "cannot parse PODINFO_REDIS_READ_TIMEOUT")
		}
	}

	return result, err
}

// DefaultRedisWriteTimeout returns the default redis write timeout
func DefaultRedisWriteTimeout() (time.Duration, error) {
	var (
		result = 3 * time.Second
		err    error
	)
	resultStr := strings.TrimSpace(os.Getenv("PODINFO_REDIS_WRITE_TIMEOUT"))
	if resultStr != "" {
		result, err = time.ParseDuration(resultStr)
		if err != nil {
			err = stacktrace.Propagate(err, "cannot parse PODINFO_REDIS_WRITE_TIMEOUT")
		}
	}
	return result, err
}

// DefaultRedisContextTimeoutEnabled returns the state of the context timeout
func DefaultRedisContextTimeoutEnabled() (bool, error) {
	var (
		result bool
		err    error
	)
	resultStr := strings.TrimSpace(os.Getenv("PODINFO_REDIS_CONTEXT_TIMEOUT_ENABLED"))
	if resultStr != "" {
		result, err = strconv.ParseBool(resultStr)
		if err != nil {
			err = stacktrace.Propagate(err, "cannot parse PODINFO_REDIS_CONTEXT_TIMEOUT_ENABLED")
		}
	}
	return result, err
}

// DefaultRedisPoolFIFO returns the default connection pool type
func DefaultRedisPoolFIFO() (bool, error) {
	var (
		result bool
		err    error
	)
	resultStr := strings.TrimSpace(os.Getenv("PODINFO_REDIS_POOL_FIFO"))
	if resultStr != "" {
		result, err = strconv.ParseBool(resultStr)
		if err != nil {
			err = stacktrace.Propagate(err, "cannot parse PODINFO_REDIS_POOL_FIFO")
			result = false
		}
	}
	return result, err
}

// DefaultRedisPoolSize returns the default connection pool size
func DefaultRedisPoolSize() (int, error) {
	var (
		result = 10 * runtime.GOMAXPROCS(0)
		err    error
	)
	resultStr := strings.TrimSpace(os.Getenv("PODINFO_REDIS_POOL_SIZE"))
	if resultStr != "" {
		result, err = strconv.Atoi(resultStr)
		if err != nil {
			err = stacktrace.Propagate(err, "cannot parse PODINFO_REDIS_POOL_SIZE")
		}
	}
	return result, err
}

// DefaultRedisPoolTimeout returns the default connection pool timeout
func DefaultRedisPoolTimeout() (time.Duration, error) {
	var (
		result = 4 * time.Second
		err    error
	)
	resultStr := strings.TrimSpace(os.Getenv("PODINFO_REDIS_POOL_TIMEOUT"))
	if resultStr != "" {
		result, err = time.ParseDuration(resultStr)
		if err != nil {
			err = stacktrace.Propagate(err, "cannot parse PODINFO_REDIS_POOL_TIMEOUT")
		}
	}
	return result, err
}

// DefaultRedisMinIdleConns returns the default redis min idle conns
func DefaultRedisMinIdleConns() (int, error) {
	var (
		result int
		err    error
	)
	resultStr := strings.TrimSpace(os.Getenv("PODINFO_REDIS_MIN_IDLE_CONNS"))
	if resultStr != "" {
		result, err = strconv.Atoi(resultStr)
		if err != nil {
			err = stacktrace.Propagate(err, "cannot parse PODINFO_REDIS_MIN_IDLE_CONNS")
			result = -1
		}
	}
	return result, err
}

// DefaultRedisMaxIdleConns returns the default redis max idle conns
func DefaultRedisMaxIdleConns() (int, error) {
	var (
		result int
		err    error
	)
	resultStr := strings.TrimSpace(os.Getenv("PODINFO_REDIS_MAX_IDLE_CONNS"))
	if resultStr != "" {
		result, err = strconv.Atoi(resultStr)
		if err != nil {
			err = stacktrace.Propagate(err, "cannot parse PODINFO_REDIS_MAX_IDLE_CONNS")
			result = -1
		}
	}
	return result, err
}

// DefaultRedisConnMaxIdleTime returns the default redis conn max idle time
func DefaultRedisConnMaxIdleTime() (time.Duration, error) {
	var (
		result = 30 * time.Minute
		err    error
	)
	resultStr := strings.TrimSpace(os.Getenv("PODINFO_REDIS_CONN_MAX_IDLE_TIME"))
	if resultStr != "" {
		result, err = time.ParseDuration(resultStr)
		if err != nil {
			err = stacktrace.Propagate(err, "cannot parse PODINFO_REDIS_CONN_MAX_IDLE_TIME")
		}
	}
	return result, err
}

// DefaultRedisConnMaxLifetime returns the default redis conn max lifetime
func DefaultRedisConnMaxLifetime() (time.Duration, error) {
	var (
		result = -1 * time.Nanosecond
		err    error
	)
	resultStr := strings.TrimSpace(os.Getenv("PODINFO_REDIS_CONN_MAX_LIFETIME"))
	if resultStr != "" {
		result, err = time.ParseDuration(resultStr)
		if err != nil {
			err = stacktrace.Propagate(err, "cannot parse PODINFO_REDIS_CONN_MAX_LIFETIME")
		}
	}
	return result, err
}

// ──────────────────────────────────────────────────── I ──────────
//
//	:::::: S E T T E R : :  :   :    :     :        :          :
//

// ──────────────────────────────────────────────────────────────
// SetRedisAddr sets the redis address
func (c *config) SetRedisAddr(value string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.RedisAddr = value
}

// SetRedisClientName sets the redis client name
func (c *config) SetRedisClientName(value string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.RedisClientName = value
}

// SetRedisUsername sets the redis username
func (c *config) SetRedisUsername(value string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.RedisUsername = value
}

// SetRedisPassword sets the redis password
func (c *config) SetRedisPassword(value string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.RedisPassword = value
}

// SetRedisDB sets the redis db
func (c *config) SetRedisDB(value int) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.RedisDB = value
}

// SetRedisMaxRetries sets the redis max retries
func (c *config) SetRedisMaxRetries(value int) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.RedisMaxRetries = value
}

// SetRedisMinRetryBackoff sets the redis min retry backoff
func (c *config) SetRedisMinRetryBackoff(value time.Duration) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.RedisMinRetryBackoff = value
}

// SetRedisMaxRetryBackoff sets the redis max retry backoff
func (c *config) SetRedisMaxRetryBackoff(value time.Duration) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.RedisMaxRetryBackoff = value
}

// SetRedisDialTimeout sets the redis dial timeout
func (c *config) SetRedisDialTimeout(value time.Duration) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.RedisDialTimeout = value
}

// SetRedisReadTimeout sets the redis read timeout
func (c *config) SetRedisReadTimeout(value time.Duration) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.RedisReadTimeout = value
}

// SetRedisWriteTimeout sets the redis write timeout
func (c *config) SetRedisWriteTimeout(value time.Duration) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.RedisWriteTimeout = value
}

// SetRedisContextTimeoutEnabled sets the redis context timeout enabled
func (c *config) SetRedisContextTimeoutEnabled(value bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.RedisContextTimeoutEnabled = value
}

// SetRedisPoolFIFO sets the redis connection pool type
func (c *config) SetRedisPoolFIFO(value bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.RedisPoolFIFO = value
}

// SetRedisPoolSize sets the redis pool size
func (c *config) SetRedisPoolSize(value int) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.RedisPoolSize = value
}

// SetRedisPoolTimeout sets the redis pool timeout
func (c *config) SetRedisPoolTimeout(value time.Duration) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.RedisPoolTimeout = value
}

// SetRedisMinIdleConns sets the redis min idle conns
func (c *config) SetRedisMinIdleConns(value int) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.RedisMinIdleConns = value
}

// SetRedisMaxIdleConns sets the redis max idle conns
func (c *config) SetRedisMaxIdleConns(value int) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.RedisMaxIdleConns = value
}

// SetRedisConnMaxIdleTime
func (c *config) SetRedisConnMaxIdleTime(value time.Duration) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.RedisConnMaxIdleTime = value
}

// SetRedisConnMaxLifetime sets the redis conn max lifetime
func (c *config) SetRedisConnMaxLifetime(value time.Duration) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.RedisConnMaxLifetime = value
}
