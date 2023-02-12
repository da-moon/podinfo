package server

import (
	"context"
	"fmt"
	"io"
	"strings"

	core "github.com/da-moon/northern-labs-interview/api/core"
	logger "github.com/da-moon/northern-labs-interview/internal/logger"
	version "github.com/da-moon/northern-labs-interview/internal/version"
	cli "github.com/mitchellh/cli"
	pterm "github.com/pterm/pterm"
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
	c.telemetry = &TelemetryFlags{}
	c.telemetry.init()
	// ────────────────────────────────────────────────────────────────────────────────
	c.flags.Merge(c.telemetry.Value())
	c.synopsis = synopsis
	c.help = c.flags.Usage()
	c.flags.Value().Usage = func() {
		c.UI.Info(c.Help())
	}
}

// Run it subcommand entrypoint
func (c *cmd) Run(args []string) int {
	if c.flags == nil {
		c.UI.Error("underlying flag struct was nil")
		return 1
	}
	flags := c.flags.Value()
	err := flags.Parse(args)
	if err != nil {
		return 1
	}
	s, _ := pterm.DefaultBigText.WithLetters(pterm.NewLettersFromString("podinfo")).Srender()
	c.UI.Output(pterm.DefaultCenter.Sprint(s))
	logLevel := c.flags.LogLevel()
	// ────────────────────────────────────────────────────────────────────────────────
	ctx := context.Background()
	l := logger.DefaultWrappedLogger(logLevel)
	// ─── INITIALIZING SERVER CONFIG ─────────────────────────────────────────────────
	conf, err := core.DefaultConfig(l)
	if err != nil {
		c.UI.Error(
			fmt.Sprintf(
				"could not start podinfo server. Err : %v",
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

	c.UI.Output("build info:")
	c.UI.Warn("")
	c.UI.Info(fmt.Sprintf("                   Version Info: '%s'", b.Info()))
	c.UI.Info(fmt.Sprintf("                   Build Context: '%s'", b.BuildContext()))
	c.UI.Warn("")
	c.UI.Output("Node info:")
	c.UI.Warn("")
	c.UI.Info(fmt.Sprintf("                   Development Mode: '%v'", dev))
	c.UI.Info(fmt.Sprintf("                   Node name: '%s'", conf.NodeName))
	c.UI.Info(fmt.Sprintf("                   API addr: '%s'", conf.APIAddr))
	c.UI.Warn("")
	c.UI.Warn("Log data will now stream in as it occurs:\n")
	// ─── REGISTER CONSUL API CORE ───────────────────────────────────────────────────
	srv, err := conf.RestfulServer(ctx)
	if err != nil {
		c.UI.Error(fmt.Sprintf("could not start podinfo. Err : %v", err))
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
