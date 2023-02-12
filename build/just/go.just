# !/usr/bin/env -S just --justfile
# vim: filetype=just tabstop=2 shiftwidth=2 softtabstop=2 expandtab:
# ────────────────────────────────────────────────────────────────────────────────
_lint-go:
  #!/usr/bin/env bash
  set -euo pipefail
  if ! golangci-lint --version > /dev/null 2>&1 ; then
    echo "*** golangci-lint not found. installing ..." ;
    wget \
      -O- \
      -nv \
    "https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh" \
    | sh \
      -s -- \
      -b "$(go env GOPATH)/bin/" \
      -d "latest"
  fi
