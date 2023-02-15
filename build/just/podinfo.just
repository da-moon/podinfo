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
  bin/podinfo server -log-level="TRACE" > "{{ justfile_directory() }}/tmp/server/log" 2>&1 &
# ────────────────────────────────────────────────────────────────────────────────
# send a GET API request to /healthz endpoint
liveness-probe:
  #!/usr/bin/env bash
  URI="healthz"
  VERB="GET"
  URL="http://localhost:${PODINFO_SERVER_PORT}/${URI}"
  resp="$(curl -o - -sSl --request "${VERB}" "${URL}" )";
  echo "${resp}" | jq -r || true
  status_code="$(curl -s -o /dev/null -w "%{http_code}" "${URL}" || true)"
  echo "Status Code: ${status_code}"
# ────────────────────────────────────────────────────────────────────────────────
readiness-probe:
  #!/usr/bin/env bash
  echo "─── SUCCESS ──────────────────────────────────────────────────────────────────"
  URI="readyz"
  VERB="GET"
  echo "❯ Sending ${VERB} request to /${URI}"
  URL="http://localhost:${PODINFO_SERVER_PORT}/${URI}"
  resp="$(curl -o - -sSl --request "${VERB}" "${URL}" )";
  echo "${resp}" | jq -r || true
  status_code="$(curl -s -o /dev/null -w "%{http_code}" "${URL}" || true)"
  echo "Status Code: ${status_code}"
