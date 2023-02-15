package server

import (
	"bytes"
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
	redis "github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
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
	logLevel = c.flags.LogLevel()
	// ────────────────────────────────────────────────────────────────────────────────
	ctx := context.Background()
	l := logger.DefaultWrappedLogger(logLevel)
	logrusLevel, err := logrus.ParseLevel(strings.ToLower(logLevel))
	if err == nil {
		logrus.SetLevel(logrusLevel)
	}
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
	disableCache := c.flags.DisableCache()
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
	redisInfo := ""
	if !disableCache {
		redisInfo, err = c.initCache(l, conf)
		if err != nil {
			c.UI.Error(fmt.Sprintf("issues detected while initializing server. err: %v", err))
			return 1
		}
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
	if disableCache {
		c.UI.Info("                   /cache/ API Group State : 'disabled'")
	}
	c.UI.Warn("")
	if !disableCache {
		c.UI.Output("Redis Info:")
		c.UI.Warn("")
		c.UI.Info(redisInfo)
		c.UI.Warn("")
	}
	c.UI.Output("Telemetry Info:")
	c.UI.Warn("")
	c.UI.Info(fmt.Sprintf("                   MetricsPrefix: '%v'", metricsPrefix))
	if statsiteAddr != "" {
		c.UI.Info(fmt.Sprintf("                   StatsiteAddr: '%v'", statsiteAddr))
	}
	if statsdAddr != "" {
		c.UI.Info(fmt.Sprintf("                   StatsdAddr: '%v'", statsdAddr))
	}
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

// initCache function initializes /cache/ api group
// It returns a string which is redis information (to be shown to the user) and
// an error
func (c *cmd) initCache(l *logger.WrappedLogger, conf *core.Config) (string, error) {
	dev := c.flags.Dev()
	var result bytes.Buffer

	// ─── REDIS FLAGS ─────────────────────────────────────────────────────────────
	redisAddr, err := c.redis.Addr()
	if err != nil {
		return "", err
	}
	if redisAddr == "" {
		if !dev {
			return "", err
		}
	}
	if redisAddr != "" {
		conf.SetRedisAddr(redisAddr)
	} else {
		if err != nil {
			return "", err
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
			return "", err
		}
	}
	redisMaxRetries, err := c.redis.MaxRetries()
	if redisMaxRetries >= -1 {
		conf.SetRedisMaxRetries(redisMaxRetries)
	} else {
		if err != nil {
			return "", err
		}
	}
	redisMinRetryBackoff, err := c.redis.MinRetryBackoff()
	if redisMinRetryBackoff >= -1 {
		conf.SetRedisMinRetryBackoff(redisMinRetryBackoff)
	} else {
		if err != nil {
			return "", err
		}
	}
	redisMaximumRetryBackoff, err := c.redis.MaximumRetryBackoff()
	if redisMaximumRetryBackoff >= -1 {
		conf.SetRedisMaxRetryBackoff(redisMaximumRetryBackoff)
	} else {
		if err != nil {
			return "", err
		}
	}
	redisDialTimeout, err := c.redis.DialTimeout()
	if redisDialTimeout >= -2 {
		conf.SetRedisDialTimeout(redisDialTimeout)
	} else {
		if err != nil {
			return "", err
		}
	}
	redisReadTimeout, err := c.redis.ReadTimeout()
	if redisReadTimeout >= -2 {
		conf.SetRedisReadTimeout(redisReadTimeout)
	} else {
		if err != nil {
			return "", err
		}
	}
	redisWriteTimeout, err := c.redis.WriteTimeout()
	if redisWriteTimeout >= -2 {
		conf.SetRedisWriteTimeout(redisWriteTimeout)
	} else {
		if err != nil {
			return "", err
		}
	}
	redisPoolFIFO, _ := c.redis.PoolFIFO()
	conf.SetRedisPoolFIFO(redisPoolFIFO)
	redisPoolSize, err := c.redis.PoolSize()
	if redisPoolSize > 0 {
		conf.SetRedisPoolSize(redisPoolSize)
	} else {
		if err != nil {
			return "", err
		}
	}
	redisPoolTimeout, _ := c.redis.PoolTimeout()
	conf.SetRedisPoolTimeout(redisPoolTimeout)
	redisMinIdleConns, _ := c.redis.MinIdleConns()
	conf.SetRedisMinIdleConns(redisMinIdleConns)
	redisMaxIdleConns, _ := c.redis.MaxIdleConns()
	conf.SetRedisMaxIdleConns(redisMaxIdleConns)
	redisConnMaxIdleTime, _ := c.redis.ConnMaxIdleTime()
	conf.SetRedisConnMaxIdleTime(redisConnMaxIdleTime)
	redisConnMaxLifetime, _ := c.redis.ConnMaxLifetime()
	conf.SetRedisConnMaxLifetime(redisConnMaxLifetime)
	opts := &redis.Options{
		Addr:                  conf.RedisAddr,
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
	}
	err = cache.Init(l, dev, opts)
	if err != nil {
		return "", err
	}
	// TODO move them to previous if statements
	fmt.Fprintf(&result, "                   Address: '%v'\n", redisAddr)
	if redisClientName != "" {
		fmt.Fprintf(&result, "                   Client Name: '%v'\n", redisClientName)
	}
	if redisUsername != "" {
		fmt.Fprintf(&result, "                   Username: '%v'\n", redisUsername)
	}
	// fmt.Fprintf(&result, "                   Password: '%v'\n", redisPassword)
	fmt.Fprintf(&result, "                   DB: '%v'\n", redisDB)
	fmt.Fprintf(&result, "                   MaxRetries: '%v'\n", redisMaxRetries)
	fmt.Fprintf(&result, "                   MinRetryBackoff: '%v'\n", redisMinRetryBackoff)
	fmt.Fprintf(&result, "                   MaximumRetryBackoff: '%v'\n", redisMaximumRetryBackoff)
	fmt.Fprintf(&result, "                   DialTimeout: '%v'\n", redisDialTimeout)
	fmt.Fprintf(&result, "                   ReadTimeout: '%v'\n", redisReadTimeout)
	fmt.Fprintf(&result, "                   WriteTimeout: '%v'\n", redisWriteTimeout)
	fmt.Fprintf(&result, "                   PoolFIFO: '%v'\n", redisPoolFIFO)
	fmt.Fprintf(&result, "                   PoolSize: '%v'\n", redisPoolSize)
	fmt.Fprintf(&result, "                   PoolTimeout: '%v'\n", redisPoolTimeout)
	fmt.Fprintf(&result, "                   MinIdleConns: '%v'\n", redisMinIdleConns)
	fmt.Fprintf(&result, "                   MaxIdleConns: '%v'\n", redisMaxIdleConns)
	fmt.Fprintf(&result, "                   ConnMaxIdleTime: '%v'\n", redisConnMaxIdleTime)
	fmt.Fprintf(&result, "                   ConnMaxLifetime: '%v'\n", redisConnMaxLifetime)
	return result.String(), nil
}
