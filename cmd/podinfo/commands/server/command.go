package server

import (
	"context"
	"fmt"
	"io"
	"strings"

	core "github.com/da-moon/northern-labs-interview/api/core"
	"github.com/da-moon/northern-labs-interview/api/handlers/cache"
	version "github.com/da-moon/northern-labs-interview/build/go/version"
	logger "github.com/da-moon/northern-labs-interview/internal/logger"
	cli "github.com/mitchellh/cli"
	pterm "github.com/pterm/pterm"
	putils "github.com/pterm/pterm/putils"
	"github.com/redis/go-redis/v9"
)

const (
	entrypoint = "server"
	synopsis   = "podinfo rest server."
	help       = `
Usage: podinfo server [options]

Starts podinfo server.
`
)

// New is the subcommand's constructor function
func New(ui cli.Ui) *cmd { // revive:disable:unexported-return
	c := &cmd{UI: ui}
	c.init()
	return c
	// revive:enable:unexported-return
}

type cmd struct {
	UI         cli.Ui
	telemetry  *TelemetryFlags
	redis      *RedisFlags
	flags      *ServerFlags
	shutdownCh <-chan struct{}
	help       string
	synopsis   string
	// testStdin is the input for testing.
	testStdin io.Reader
}

func (c *cmd) init() {
	c.UI = &cli.PrefixedUi{
		OutputPrefix: "",
		InfoPrefix:   "",
		ErrorPrefix:  "",
		Ui:           c.UI,
	}

	c.flags = &ServerFlags{}
	c.flags.init()
	c.redis = &RedisFlags{}
	c.redis.init()
	c.telemetry = &TelemetryFlags{}
	c.telemetry.init()
	// ────────────────────────────────────────────────────────────────────────────────
	c.flags.Merge(c.redis.Value())
	c.flags.Merge(c.telemetry.Value())
	c.synopsis = synopsis
	c.help = c.flags.Usage()
	c.flags.Value().Usage = func() {
		c.UI.Info(c.Help())
	}
}

// Run it subcommand entrypoint
func (c *cmd) Run(args []string) int { // nolint:gocyclo // this function is well tested
	if c.flags == nil {
		c.UI.Error("underlying flag struct was nil")
		return 1
	}
	flags := c.flags.Value()
	err := flags.Parse(args)
	if err != nil {
		return 1
	}

	s, _ := pterm.DefaultBigText.WithLetters(putils.LettersFromString("podinfo")).Srender()
	c.UI.Output(pterm.DefaultCenter.Sprint(s))
	var logLevel string
	// logLevel = c.flags.LogLevel()
	// ────────────────────────────────────────────────────────────────────────────────
	ctx := context.Background()
	// FIXME: do not hardcode
	logLevel = "TRACE"
	l := logger.DefaultWrappedLogger(logLevel)
	// ─── INITIALIZING SERVER CONFIG ─────────────────────────────────────────────────
	conf, err := core.DefaultConfig(l)
	if err != nil {
		c.UI.Error(
			fmt.Sprintf(
				"issues detected while initializing server. Err : %v",
				err,
			),
		)
		return 1
	}
	if conf == nil {
		c.UI.Error("nil server config")
		return 1
	}
	// ─── SERVER FLAGS ───────────────────────────────────────────────────────────────
	dev := c.flags.Dev()
	if dev {
		conf.SetDevelopmentMode()
	}
	nodeName := c.flags.NodeName()
	if nodeName != "" {
		conf.SetNodeName(nodeName)
	}
	apiAddr := c.flags.APIAddr()
	if apiAddr != "" {
		conf.SetAPIAddr(apiAddr)
	}
	// ─── REDIS FLAGS ─────────────────────────────────────────────────────────────
	redisAddr, err := c.redis.Addr()
	if err != nil {
		c.UI.Error(fmt.Sprintf("issues detected while initializing server. err: %v", err))
	}
	if redisAddr == "" {
		if !dev {
			return 1
		}
	}
	if redisAddr != "" {
		conf.SetRedisAddr(redisAddr)
	} else {
		if err != nil {
			c.UI.Error(fmt.Sprintf("issues detected while initializing server. err: %v", err))
			return 1
		}
	}
	redisClientName, _ := c.redis.ClientName()
	if redisClientName != "" {
		conf.SetRedisClientName(redisClientName)
	}
	redisUsername, _ := c.redis.Username()
	if redisUsername != "" {
		conf.SetRedisUsername(redisUsername)
	}
	redisPassword, _ := c.redis.Password()
	if redisPassword != "" {
		conf.SetRedisPassword(redisPassword)
	}
	redisDB, err := c.redis.DB()
	if redisDB >= 0 {
		conf.SetRedisDB(redisDB)
	} else {
		if err != nil {
			c.UI.Error(fmt.Sprintf("issues detected while initializing server. err: %v", err))
			return 1
		}
	}
	redisMaxRetries, err := c.redis.MaxRetries()
	if redisMaxRetries >= -1 {
		conf.SetRedisMaxRetries(redisMaxRetries)
	} else {
		if err != nil {
			c.UI.Error(fmt.Sprintf("issues detected while initializing server. err: %v", err))
			return 1
		}
	}
	redisMinRetryBackoff, err := c.redis.MinRetryBackoff()
	if redisMinRetryBackoff >= -1 {
		conf.SetRedisMinRetryBackoff(redisMinRetryBackoff)
	} else {
		if err != nil {
			c.UI.Error(fmt.Sprintf("issues detected while initializing server. err: %v", err))
			return 1
		}
	}
	redisMaximumRetryBackoff, err := c.redis.MaximumRetryBackoff()
	if redisMaximumRetryBackoff >= -1 {
		conf.SetRedisMaxRetryBackoff(redisMaximumRetryBackoff)
	} else {
		if err != nil {
			c.UI.Error(fmt.Sprintf("issues detected while initializing server. err: %v", err))
			return 1
		}
	}
	redisDialTimeout, err := c.redis.DialTimeout()
	if redisDialTimeout >= -2 {
		conf.SetRedisDialTimeout(redisDialTimeout)
	} else {
		if err != nil {
			c.UI.Error(fmt.Sprintf("issues detected while initializing server. err: %v", err))
			return 1
		}
	}
	redisReadTimeout, err := c.redis.ReadTimeout()
	if redisReadTimeout >= -2 {
		conf.SetRedisReadTimeout(redisReadTimeout)
	} else {
		if err != nil {
			c.UI.Error(fmt.Sprintf("issues detected while initializing server. err: %v", err))
			return 1
		}
	}
	redisWriteTimeout, err := c.redis.WriteTimeout()
	if redisWriteTimeout >= -2 {
		conf.SetRedisWriteTimeout(redisWriteTimeout)
	} else {
		if err != nil {
			c.UI.Error(fmt.Sprintf("issues detected while initializing server. err: %v", err))
			return 1
		}
	}
	redisPoolFIFO, _ := c.redis.PoolFIFO()
	conf.SetRedisPoolFIFO(redisPoolFIFO)
	redisPoolSize, err := c.redis.PoolSize()
	if redisPoolSize > 0 {
		conf.SetRedisPoolSize(redisPoolSize)
	} else {
		if err != nil {
			c.UI.Error(fmt.Sprintf("issues detected while initializing server. err: %v", err))
			return 1
		}
	}
	redisPoolTimeout, _ := c.redis.PoolTimeout()
	conf.SetRedisPoolTimeout(redisPoolTimeout)
	redisMinIdleConns, _ := c.redis.MinIdleConns()
	conf.SetRedisMinIdleConns(redisMinIdleConns)
	redisMaxIdleConns, _ := c.redis.MaxIdleConns()
	redisConnMaxIdleTime, _ := c.redis.ConnMaxIdleTime()
	conf.SetRedisConnMaxIdleTime(redisConnMaxIdleTime)
	redisConnMaxLifetime, _ := c.redis.ConnMaxLifetime()
	conf.SetRedisConnMaxLifetime(redisConnMaxLifetime)
	err = cache.Init(l, dev, &redis.Options{
		Addr:                  conf.APIAddr,
		ClientName:            conf.RedisClientName,
		Username:              conf.RedisUsername,
		Password:              conf.RedisPassword,
		DB:                    conf.RedisDB,
		MaxRetries:            conf.RedisMaxRetries,
		MinRetryBackoff:       conf.RedisMinRetryBackoff,
		MaxRetryBackoff:       conf.RedisMaxRetryBackoff,
		DialTimeout:           conf.RedisDialTimeout,
		ReadTimeout:           conf.RedisReadTimeout,
		WriteTimeout:          conf.RedisWriteTimeout,
		ContextTimeoutEnabled: conf.RedisContextTimeoutEnabled,
		PoolFIFO:              conf.RedisPoolFIFO,
		PoolSize:              conf.RedisPoolSize,
		PoolTimeout:           conf.RedisPoolTimeout,
		MinIdleConns:          conf.RedisMinIdleConns,
		MaxIdleConns:          conf.RedisMaxIdleConns,
		ConnMaxIdleTime:       conf.RedisConnMaxIdleTime,
		ConnMaxLifetime:       conf.RedisConnMaxLifetime,
	})
	if err != nil {
		c.UI.Error(fmt.Sprintf("issues detected while initializing server. err: %v", err))
		return 1
	}
	// ─── TELEMETRY FLAGS ────────────────────────────────────────────────────────────
	metricsPrefix := c.telemetry.MetricsPrefix()
	if metricsPrefix != "" {
		conf.SetMetricsPrefix(metricsPrefix)
	}
	statsdAddr := c.telemetry.StatsdAddr()
	if statsdAddr != "" {
		conf.SetStatsdAddr(statsdAddr)
	}
	statsiteAddr := c.telemetry.StatsiteAddr()
	if statsiteAddr != "" {
		conf.SetStatsiteAddr(statsiteAddr)
	}
	prometheusRetentionTime := c.telemetry.PrometheusRetentionTime()
	if prometheusRetentionTime != 0 {
		conf.SetPrometheusRetentionTime(prometheusRetentionTime)
	}
	// ────────────────────────────────────────────────────────────────────────────────
	b := version.New()
	c.UI.Warn("")
	c.UI.Output(pterm.DefaultCenter.Sprint(pterm.Info.Sprint("podinfo running!")))
	c.UI.Warn("")
	c.UI.Warn("")
	c.UI.Output("build info:")
	c.UI.Warn("")
	c.UI.Info(fmt.Sprintf("                   Version Info: '%s'", b.Info()))
	c.UI.Info(fmt.Sprintf("                   Build Context: '%s'", b.BuildContext()))
	c.UI.Warn("")
	c.UI.Output("Node info:")
	c.UI.Warn("")
	c.UI.Info(fmt.Sprintf("                   Log Level: '%v'", logLevel))
	c.UI.Info(fmt.Sprintf("                   Development Mode: '%v'", dev))
	c.UI.Info(fmt.Sprintf("                   Node name: '%s'", conf.NodeName))
	c.UI.Info(fmt.Sprintf("                   API addr: '%s'", conf.APIAddr))
	c.UI.Warn("")
	c.UI.Output("Redis Info:")
	c.UI.Warn("")
	c.UI.Info(fmt.Sprintf("                   Address: '%v'", redisAddr))
	c.UI.Info(fmt.Sprintf("                   Client Name: '%v'", redisClientName))
	c.UI.Info(fmt.Sprintf("                   Username: '%v'", redisUsername))
	c.UI.Info(fmt.Sprintf("                   DB: '%v'", redisDB))
	c.UI.Info(fmt.Sprintf("                   MaxRetries: '%v'", redisMaxRetries))
	c.UI.Info(fmt.Sprintf("                   MinRetryBackoff: '%v'", redisMinRetryBackoff))
	c.UI.Info(fmt.Sprintf("                   MaximumRetryBackoff: '%v'", redisMaximumRetryBackoff))
	c.UI.Info(fmt.Sprintf("                   DialTimeout: '%v'", redisDialTimeout))
	c.UI.Info(fmt.Sprintf("                   ReadTimeout: '%v'", redisReadTimeout))
	c.UI.Info(fmt.Sprintf("                   WriteTimeout: '%v'", redisWriteTimeout))
	c.UI.Info(fmt.Sprintf("                   PoolFIFO: '%v'", redisPoolFIFO))
	c.UI.Info(fmt.Sprintf("                   PoolSize: '%v'", redisPoolSize))
	c.UI.Info(fmt.Sprintf("                   PoolTimeout: '%v'", redisPoolTimeout))
	c.UI.Info(fmt.Sprintf("                   MinIdleConns: '%v'", redisMinIdleConns))
	c.UI.Info(fmt.Sprintf("                   MaxIdleConns: '%v'", redisMaxIdleConns))
	c.UI.Info(fmt.Sprintf("                   ConnMaxIdleTime: '%v'", redisConnMaxIdleTime))
	c.UI.Info(fmt.Sprintf("                   ConnMaxLifetime: '%v'", redisConnMaxLifetime))
	c.UI.Warn("")
	c.UI.Output("Telemetry Info:")
	c.UI.Warn("")
	c.UI.Info(fmt.Sprintf("                   MetricsPrefix: '%v'", metricsPrefix))
	c.UI.Info(fmt.Sprintf("                   StatsiteAddr: '%v'", statsiteAddr))
	c.UI.Info(fmt.Sprintf("                   StatsdAddr: '%v'", statsdAddr))
	c.UI.Info(fmt.Sprintf("                   PrometheusRetentionTime: '%v'", prometheusRetentionTime))
	c.UI.Warn("")
	c.UI.Warn("Log data will now stream in as it occurs:\n")
	srv, err := conf.RestfulServer(ctx)
	if err != nil {
		c.UI.Error(fmt.Sprintf("issues detected while initializing. Err : %v", err))
		return 1
	}
	if srv == nil {
		c.UI.Error("instantiated podinfo was a nil pointer")
		return 1
	}
	defer srv.Shutdown()
	return c.handleSignals(srv)
}

// Synopsis shows the short description of
// 'podinfo server' command.
// it is the string on the right hand side
// of 'server' command when a user runs 'podinfo --help'
func (c *cmd) Synopsis() string {
	return strings.TrimSpace(c.synopsis)
}

// Help represents the long form Subcommand help.
// it is what is shown when a user runs 'podinfo server --help'
func (c *cmd) Help() string {
	return strings.TrimSpace(c.help)
}
