package server

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	core "github.com/da-moon/podinfo/api/core"
)

const (
	gracefulTimeout = 3 * time.Second
)

func (c *cmd) handleSignals(srv *core.RestfulServer) int {
	signalCh := make(chan os.Signal, 4)
	signal.Notify(signalCh, os.Interrupt, syscall.SIGTERM, syscall.SIGHUP)
	var sig os.Signal
	select {
	case s := <-signalCh:
		sig = s
	case <-c.shutdownCh:
		sig = os.Interrupt
	case <-srv.ShutdownCh():
		return 0
	}
	c.UI.Output(fmt.Sprintf("Caught signal: %v", sig))
	graceful := false
	if sig == os.Interrupt {
		graceful = true
	} else if sig == syscall.SIGTERM {
		graceful = true
	}
	if !graceful {
		return 1
	}
	gracefulCh := make(chan struct{})
	c.UI.Output("Gracefully shutting down agent...")
	go func() {
		srv.Shutdown()
		close(gracefulCh)
	}()
	select {
	case <-signalCh:
		return 1
	case <-time.After(gracefulTimeout):
		return 1
	case <-gracefulCh:
		return 0
	}
}
