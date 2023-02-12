# !/usr/bin/env -S just --justfile
# vim: filetype=just tabstop=2 shiftwidth=2 softtabstop=2 expandtab:
# ────────────────────────────────────────────────────────────────────────────────
# ensures 'shfmt' is installed
_format-bash:
  #!/usr/bin/env bash
  set -euo pipefail
  if ! shfmt --version > /dev/null 2>&1 ; then
    echo "*** shfmt not found. installing ..." ;
    go install "mvdan.cc/sh/v3/cmd/shfmt@latest" ;
  fi
