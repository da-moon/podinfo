# !/usr/bin/env -S just --justfile
# vim: filetype=just tabstop=2 shiftwidth=2 softtabstop=2 expandtab:
# ────────────────────────────────────────────────────────────────────────────────
# name of built binary
BINARY_NAME := 'podinfo'
# ────────────────────────────────────────────────────────────────────────────────
kill:
  #!/usr/bin/env bash
  set -euo pipefail
  pkill -9 "{{BINARY_NAME}}" || true
