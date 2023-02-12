package main

import (
	"os"

	server "github.com/da-moon/northern-labs-interview/cmd/podinfo/commands/server"
	version "github.com/da-moon/northern-labs-interview/cmd/podinfo/commands/version"
	cli "github.com/mitchellh/cli"
)

var Commands map[string]cli.CommandFactory

func init() {
	ui := &cli.BasicUi{
		Reader:      os.Stdin,
		Writer:      os.Stdout,
		ErrorWriter: os.Stderr,
	}
	Commands = map[string]cli.CommandFactory{
		"server": func() (cli.Command, error) {
			return server.New(ui), nil
		},
		"version": func() (cli.Command, error) {
			return version.New(ui), nil
		},
	}
}
