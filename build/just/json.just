# !/usr/bin/env -S just --justfile
# vim: filetype=just tabstop=2 shiftwidth=2 softtabstop=2 expandtab:
# ────────────────────────────────────────────────────────────────────────────────
_format-json:
    #!/usr/bin/env bash
    set -euo pipefail
    if ! command -- jsonfmt -h > /dev/null 2>&1 ; then
      echo "*** 'jsonfmt' not found. Installing ..." ;
      just _install-rust-package jsonfmt
    fi
