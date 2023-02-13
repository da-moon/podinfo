# !/usr/bin/env -S just --justfile
# vim: filetype=just tabstop=2 shiftwidth=2 softtabstop=2 expandtab:
# ────────────────────────────────────────────────────────────────────────────────
# name of built binary
BINARY_NAME := 'podinfo'
# ────────────────────────────────────────────────────────────────────────────────
# send SIGTERM to running binary to stop it
kill:
  #!/usr/bin/env bash
  set -euo pipefail
  pkill -9 "{{BINARY_NAME}}" || true
  just clean-go
# ────────────────────────────────────────────────────────────────────────────────
# build and start the server and forward logs to ./tmp/server/log
run: build-go
  #!/usr/bin/env bash
  set -euo pipefail
  rm -rf "{{ justfile_directory() }}/tmp/server"
  mkdir -p "{{ justfile_directory() }}/tmp/server"
  bin/podinfo server > "{{ justfile_directory() }}/tmp/server/log" 2>&1 &
# ────────────────────────────────────────────────────────────────────────────────
liveness-probe:
  #!/usr/bin/env bash
  set -euo pipefail
  URI="healthz"
  VERB="GET"
  curl -fsSl \
    --request "${VERB}" \
  "http://localhost:${PODINFO_SERVER_PORT}/${URI}" \
  | jq -r
