// Package core is the common package for all API types.
// it helps with creating base struct required for
// starting a new API server.
// ─── WORKFLOW ───────────────────────────────────────────────────────────────────
// - create a new config object by calling DefaultConfig()
// function
// conf:=DefaultConfig()
// - start the graceful server
// srv := conf.RestfulServer()
// - ensure the server gracefully terminates
// defer srv.Shutdown()
// ─── USEFUL COMMANDS ────────────────────────────────────────────────────────────
// - lint:
// golangci-lint run --config=.golangci.yml --print-issued-lines=false --exclude-use-default=false api/core
// ────────────────────────────────────────────────────────────────────────────────
package core
