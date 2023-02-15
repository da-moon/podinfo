package server

import (
	"time"

	"github.com/da-moon/northern-labs-interview/api/core"
	flagset "github.com/da-moon/northern-labs-interview/internal/cli/flagset"
	value "github.com/da-moon/northern-labs-interview/internal/cli/value"
	"github.com/palantir/stacktrace"
	redis "github.com/redis/go-redis/v9"
)

// RedisFlags is a struct that
// contains the flags for used with configuring
// the redis related parameters.
type RedisFlags struct {
	*flagset.FlagSet
	addr                  value.String
	clientName            value.String
	username              value.String
	password              value.String
	db                    value.Int
	maxRetries            value.Int
	minRetryBackoff       value.Duration
	maxRetryBackoff       value.Duration
	dialTimeout           value.Duration
	readTimeout           value.Duration
	writeTimeout          value.Duration
	contextTimeoutEnabled value.Bool
	poolFIFO              value.Bool
	poolSize              value.Int
	poolTimeout           value.Duration
	minIdleConns          value.Int
	maxIdleConns          value.Int
	connMaxIdleTime       value.Duration
	connMaxLifetime       value.Duration
}

func (f *RedisFlags) init() {
	redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default db
	})
	f.FlagSet = flagset.New("", "")
	f.Var(&f.addr, "redis-addr",
		[]string{
			"flag used to set  address",
			"it is in host:port address format.",
			"This can also be specified via the 'PODINFO_REDIS_ADDR' env variable",
		})
	f.Var(&f.clientName, "redis-client-name",
		[]string{
			"flag used to set  client name.",
			"ClientName will execute the `CLIENT SETNAME ClientName` command for each conn.",
			"This can also be specified via the 'PODINFO_REDIS_CLIENT_NAME' env variable",
		})
	f.Var(&f.username, "redis-username",
		[]string{
			"flag used to set  username.",
			"Use the specified username to authenticate the current connection" +
				"with one of the connections defined in the ACL list when connecting" +
				"to a  6.0 instance, or greater, that is using the  ACL system.",
			"This can also be specified via the 'PODINFO_REDIS_USERNAME' env variable"},
	)
	f.Var(&f.password, "redis-password",
		[]string{
			"flag used to set  password.",
			"Optional password. Must match the password specified in the" +
				"requirepass server configuration option (if connecting to a  5.0 instance, or lower)," +
				"or the User password when connecting to a  6.0 instance, or greater," +
				"that is using the  ACL system.",
			"This can also be specified via the 'PODINFO_REDIS_PASSWORD' env variable"},
	)
	f.Var(&f.db, "redis-db",
		[]string{
			"flag used to set  database.",
			"Database to be selected after connecting to the server.",
			"This can also be specified via the 'PODINFO_REDIS_db' env variable"},
	)
	f.Var(&f.maxRetries, "redis-max-retries",
		[]string{
			"flag used to set the maximum number of retries before giving up.",
			"Default is 3 retries; -1 disables retries.",
			"This can also be specified via the 'PODINFO_REDIS_MAX_RETRIES' env variable"},
	)
	f.Var(&f.minRetryBackoff, "redis-min-retry-backoff",
		[]string{
			"flag used to set minimum backoff between each retry.",
			"Default is 8 milliseconds; -1 disables backoff.",
			"This can also be specified via the 'PODINFO_REDIS_MIN_RETRY_BACKOFF' env variable"},
	)
	f.Var(&f.maxRetryBackoff, "redis-max-retry-backoff",
		[]string{
			"flag used to set Specifies the maximum backoff between each retry.",
			"Default is 512 milliseconds; -1 disables backoff.",
			"This can also be specified via the 'PODINFO_REDIS_MAX_RETRY_BACKOFF' env variable",
		},
	)
	f.Var(&f.dialTimeout, "redis-dial-timeout",
		[]string{
			"flag used to set timeout for socket reads.",
			"If reached, commands will fail" +
				"with a timeout instead of blocking. Supported values:",
			"  - `0` - default timeout (3 seconds).",
			"  - `-1` - no timeout (block indefinitely).",
			"  - `-2` - disables SetReadDeadline calls completely.",
			"Default is 5 seconds.",
			"This can also be specified via the 'PODINFO_REDIS_DIAL_TIMEOUT' env variable",
		},
	)
	f.Var(&f.readTimeout, "redis-read-timeout",
		[]string{
			"flag used to set  read timeout.",
			"Timeout for socket reads. If reached, commands will fail" +
				"with a timeout instead of blocking. Supported values:",
			"    - '0' - default timeout (3 seconds).",
			"    - '-1' - no timeout (block indefinitely).",
			"    - '-2' - disables SetReadDeadline calls completely.",
			"This can also be specified via the 'PODINFO_REDIS_READ_TIMEOUT' env variable"},
	)
	f.Var(&f.writeTimeout, "redis-write-timeout",
		[]string{
			"flag used to set  write timeout.",
			"Timeout for socket writes. If reached, commands will fail" +
				"with a timeout instead of blocking.  Supported values:",
			"  - `0` - default timeout (3 seconds).",
			"  - `-1` - no timeout (block indefinitely).",
			"  - `-2` - disables SetWriteDeadline calls completely.",
			"This can also be specified via the 'PODINFO_REDIS_WRITE_TIMEOUT' env variable"},
	)
	f.Var(&f.contextTimeoutEnabled, "redis-context-timeout-enabled",
		[]string{
			"flag used to set  context timeout enabled.",
			"contextTimeoutEnabled controls whether the client respects context timeouts and deadlines.",
			"See https://.uptrace.dev/guide/go--debugging.html#timeouts",
			"This can also be specified via the 'PODINFO_REDIS_CONTEXT_TIMEOUT_ENABLED' env variable"},
	)
	f.Var(&f.poolFIFO, "redis-pool-fifo",
		[]string{
			"flag used to set  pool fifo.",
			"Type of connection pool.",
			"true for FIFO pool, false for LIFO pool.",
			"Note that FIFO has slightly higher overhead compared to LIFO," +
				"but it helps closing idle connections faster reducing the pool size.",
			"This can also be specified via the 'PODINFO_REDIS_POOL_FIFO' env variable"},
	)
	f.Var(&f.poolSize, "redis-pool-size",
		[]string{
			"flag used to set  pool size.",
			"Maximum number of socket connections.",
			"Default is 10 connections per every available CPU",
			"This can also be specified via the 'PODINFO_REDIS_POOL_SIZE' env variable",
		})
	f.Var(&f.poolTimeout, "redis-pool-timeout",
		[]string{
			"flag used to set  pool timeout.",
			"Amount of time client waits for connection if all connections" +
				"are busy before returning an error.",
			"Default is readTimeout + 1 second.",
			"This can also be specified via the 'PODINFO_REDIS_POOL_TIMEOUT' env variable"},
	)
	f.Var(&f.minIdleConns, "redis-min-idle-conns",
		[]string{
			"flag used to set the minimum number of idle connections.",
			"it is useful when establishing new connection is slow.",
			"This can also be specified via the 'PODINFO_REDIS_MIN_IDLE_CONNS' env variable"},
	)
	f.Var(&f.maxIdleConns, "redis-max-idle-conns",
		[]string{
			"flag used to set the maximum number of idle connections.",
			"This can also be specified via the 'PODINFO_REDIS_MAX_IDLE_CONNS' env variable"},
	)
	f.Var(&f.connMaxLifetime, "redis-conn-max-idle-time",
		[]string{
			"flag used to set the maximum amount of time a connection may be idle." +
				"Should be less than server's timeout",
			" Expired connections may be closed lazily before reuse" +
				"If d <= 0, connections are not closed due to a connection's idle time.",
			"Default is 30 minutes. -1 disables idle timeout check.",
			"This can also be specified via the 'PODINFO_REDIS_CONN_MAX_IDLE_TIME' env variable"},
	)
	f.Var(&f.connMaxLifetime, "redis-conn-max-lifetime",
		[]string{
			"flag used to set the maximum amount of time a connection may be reused.",
			"Expired connections may be closed lazily before reuse." +
				"If <= 0, connections are not closed due to a connection's age.",
			"Default is to not close idle connections.",
			"This can also be specified via the 'PODINFO_REDIS_CONN_MAX_LIFETIME' env variable"},
	)
}

// Addr returns parsed 'redis-addr' flag.
func (f *RedisFlags) Addr() (string, error) {
	var err error
	result := f.addr.Get()
	if result == "" {
		result, err = core.DefaultRedisAddr()
		err = stacktrace.Propagate(err, "failed to parse 'redis-addr' flag")
	}
	return result, err
}

// ClientName returns parsed 'redis-client-name' flag.
func (f *RedisFlags) ClientName() (string, error) {
	var err error
	result := f.clientName.Get()
	if result == "" {
		result, err = core.DefaultRedisClientName()
		err = stacktrace.Propagate(err, "failed to parse 'redis-client-name' flag")
	}
	return result, err
}

// Username returns parsed 'redis-username' flag.
func (f *RedisFlags) Username() (string, error) {
	var err error
	result := f.username.Get()
	if result == "" {
		result, err = core.DefaultRedisPassword()
		err = stacktrace.Propagate(err, "failed to parse 'redis-username' flag")
	}
	return result, err
}

// Password returns parsed 'redis-password' flag.
func (f *RedisFlags) Password() (string, error) {
	var err error
	result := f.password.Get()
	if result == "" {
		result, err = core.DefaultRedisPassword()
		err = stacktrace.Propagate(err, "failed to parse 'redis-password' flag")
	}
	return result, err
}

// DB returns parsed 'redis-db' flag.
func (f *RedisFlags) DB() (int, error) {
	var (
		result = f.db.Get()
		err    error
	)
	if result == 0 {
		result, err = core.DefaultRedisDB()
		err = stacktrace.Propagate(err, "failed to parse 'redis-db' flag")
		return result, err
	}
	if result < 0 {
		result, err = core.DefaultRedisDB()
		err = stacktrace.Propagate(err, "failed to parse 'redis-db' flag")
		return result, err
	}
	return result, err
}

// MaxRetries returns parsed 'redis-max-retries' flag.
func (f *RedisFlags) MaxRetries() (int, error) {
	var (
		result = f.maxRetries.Get()
		err    error
	)
	if result == 0 {
		result, err = core.DefaultRedisMaxRetries()
		err = stacktrace.Propagate(err, "failed to parse 'redis-max-retries' flag")
		return result, err
	}
	if result < -1 {
		result, err = core.DefaultRedisMaxRetries()
		err = stacktrace.Propagate(err, "failed to parse 'redis-max-retries' flag")
		return result, err
	}
	return result, err
}

// MinRetryBackoff returns parsed 'redis-min-retry-backoff' flag.
func (f *RedisFlags) MinRetryBackoff() (time.Duration, error) {
	var (
		result = f.minRetryBackoff.Get()
		err    error
	)
	if result == 0 {
		result, err = core.DefaultRedisMinRetryBackoff()
		if err != nil {
			err = stacktrace.Propagate(err, "failed to parse 'redis-min-retry-backoff' flag")
			return result, err
		}
	}
	result = f.minRetryBackoff.Get()
	if result < -1 {
		result, err = core.DefaultRedisMinRetryBackoff()
		if err != nil {
			err = stacktrace.Propagate(err, "failed to parse 'redis-min-retry-backoff' flag")
			return result, err
		}
	}
	return result, err
}

// MaxRetryBackoff returns parsed '-redis-max-retry-backoff' flag.
func (f *RedisFlags) MaximumRetryBackoff() (time.Duration, error) {
	var (
		result = f.maxRetryBackoff.Get()
		err    error
	)
	if result == 0 {
		result, err = core.DefaultRedisMaxRetryBackoff()
		if err != nil {
			err = stacktrace.Propagate(err, "failed to parse 'redis-max-retry-backoff' flag")
			return result, err
		}
	}
	result = f.maxRetryBackoff.Get()
	if result < -1 {
		result, err = core.DefaultRedisMaxRetryBackoff()
		if err != nil {
			err = stacktrace.Propagate(err, "failed to parse 'redis-max-retry-backoff' flag")
			return result, err
		}
	}
	return result, err
}

// DialTimeout returns parsed 'redis-dial-timeout' flag.
func (f *RedisFlags) DialTimeout() (time.Duration, error) {
	var (
		result = f.dialTimeout.Get()
		err    error
	)
	if result == 0 {
		result, err = core.DefaultRedisDialTimeout()
		if err != nil {
			err = stacktrace.Propagate(err, "failed to parse'redis-dial-timeout' flag")
			return result, err
		}
	}
	result = f.dialTimeout.Get()
	if result < 0 {
		result, err = core.DefaultRedisDialTimeout()
		if err != nil {
			err = stacktrace.Propagate(err, "failed to parse'redis-dial-timeout' flag")
			return result, err
		}
	}
	return result, err
}

// ReadTimeout returns parsed '-redis-read-timeout' flag.
func (f *RedisFlags) ReadTimeout() (time.Duration, error) {
	var (
		result = f.readTimeout.Get()
		err    error
	)
	if result == 0 {
		result, err = core.DefaultRedisReadTimeout()
		if err != nil {
			err = stacktrace.Propagate(err, "failed to parse 'redis-read-timeout' flag")
			return result, err
		}
	}
	result = f.readTimeout.Get()
	if result < -2 {
		result, err = core.DefaultRedisReadTimeout()
		if err != nil {
			err = stacktrace.Propagate(err, "failed to parse 'redis-read-timeout' flag")
			return result, err
		}
	}
	return result, err
}

// WriteTimeout returns parsed '-redis-write-timeout' flag.
func (f *RedisFlags) WriteTimeout() (time.Duration, error) {
	var (
		result = f.writeTimeout.Get()
		err    error
	)
	if result == 0 {
		result, err = core.DefaultRedisWriteTimeout()
		if err != nil {
			err = stacktrace.Propagate(err, "failed to parse 'redis-write-timeout' flag")
			return result, err
		}
	}
	result = f.writeTimeout.Get()
	if result < -2 {
		result, err = core.DefaultRedisWriteTimeout()
		if err != nil {
			err = stacktrace.Propagate(err, "failed to parse 'redis-write-timeout' flag")
			return result, err
		}
	}
	return result, err
}

// PoolFIFO returns parsed '-redis-redis-pool-fifo' flag.
func (f *RedisFlags) PoolFIFO() (bool, error) {
	return f.poolFIFO.Get(), nil
}

// PoolSize returns parsed '-redis-redis-pool-size' flag.
func (f *RedisFlags) PoolSize() (int, error) {
	var (
		result = f.poolSize.Get()
		err    error
	)
	if result == 0 {
		result, err = core.DefaultRedisPoolSize()
		err = stacktrace.Propagate(err, "failed to parse 'redis-pool-size' flag")
		return result, err
	}
	if result <= 0 {
		result, err = core.DefaultRedisPoolSize()
		err = stacktrace.Propagate(err, "failed to parse 'redis-pool-size' flag")
		return result, err
	}
	return result, err
}

// PoolTimeout returns parsed '-redis-redis-pool-timeout' flag.
func (f *RedisFlags) PoolTimeout() (time.Duration, error) {
	var (
		result = f.poolTimeout.Get()
		err    error
	)
	if result == 0 {
		result, err = core.DefaultRedisPoolTimeout()
		if err != nil {
			err = stacktrace.Propagate(err, "failed to parse 'redis-pool-timeout' flag")
			return result, err
		}
	}
	result = f.poolTimeout.Get()
	if result < -2 {
		result, err = core.DefaultRedisPoolTimeout()
		if err != nil {
			err = stacktrace.Propagate(err, "failed to parse 'redis-pool-timeout' flag")
			return result, err
		}
	}
	return result, err
}

// MinIdleConns returns parsed '-redis-redis-min-idle-conns' flag.
func (f *RedisFlags) MinIdleConns() (int, error) {
	var (
		result = f.minIdleConns.Get()
		err    error
	)
	if result == 0 {
		result, err = core.DefaultRedisMinIdleConns()
		err = stacktrace.Propagate(err, "failed to parse 'redis-min-idle-conns' flag")
		return result, err
	}
	if result < 0 {
		result, err = core.DefaultRedisMinIdleConns()
		err = stacktrace.Propagate(err, "failed to parse 'redis-min-idle-conns' flag")
		return result, err
	}
	return result, err
}

// MaxIdleConns returns parsed '-redis-redis-min-idle-conns' flag.
func (f *RedisFlags) MaxIdleConns() (int, error) {
	var (
		result = f.maxIdleConns.Get()
		err    error
	)
	if result == 0 {
		result, err = core.DefaultRedisMaxIdleConns()
		err = stacktrace.Propagate(err, "failed to parse 'redis-min-idle-conns' flag")
		return result, err
	}
	if result < -1 {
		result, err = core.DefaultRedisMaxIdleConns()
		err = stacktrace.Propagate(err, "failed to parse 'redis-min-idle-conns' flag")
		return result, err
	}
	return result, err
}

// ConnMaxIdleTime returns parsed '-redis-redis-conn-max-idle-time' flag.
func (f *RedisFlags) ConnMaxIdleTime() (time.Duration, error) {
	var (
		result = f.connMaxLifetime.Get()
		err    error
	)
	if result == 0 {
		result, err = core.DefaultRedisConnMaxIdleTime()
		if err != nil {
			err = stacktrace.Propagate(err, "failed to parse 'redis-conn-max-idle-time' flag")
			return result, err
		}
	}
	if result < -1 {
		result, err = core.DefaultRedisConnMaxIdleTime()
		if err != nil {
			err = stacktrace.Propagate(err, "failed to parse 'redis-conn-max-idle-time' flag")
			return result, err
		}
	}
	return result, err
}

// ConnMaxLifetime returns parsed '-redis-redis-conn-max-lifetime' flag.
func (f *RedisFlags) ConnMaxLifetime() (time.Duration, error) {
	var (
		result = f.connMaxLifetime.Get()
		err    error
	)
	if result == 0 {
		result, err = core.DefaultRedisConnMaxLifetime()
		if err != nil {
			err = stacktrace.Propagate(err, "failed to parse 'redis-conn-max-lifetime' flag")
			return result, err
		}
	}
	if result < -2 {
		result, err = core.DefaultRedisConnMaxLifetime()
		if err != nil {
			err = stacktrace.Propagate(err, "failed to parse 'redis-conn-max-lifetime' flag")
			return result, err
		}
	}
	return result, err
}
