# !/usr/bin/env -S just --justfile
# vim: filetype=just tabstop=2 shiftwidth=2 softtabstop=2 expandtab:
# ────────────────────────────────────────────────────────────────────────────────
# ensures 'jsonfmt' is installed
_format-json:
    #!/usr/bin/env bash
    set -euo pipefail
    if ! command -- jsonfmt -h > /dev/null 2>&1 ; then
      echo "*** 'jsonfmt' not found. Installing ..." ;
      just _install-rust-package jsonfmt
    fi
# ────────────────────────────────────────────────────────────────────────────────
alias json-fmt := format-json
# detect and format all json files
format-json: _format-json
  #!/usr/bin/env bash
  set -euo pipefail
  if ! command -- jsonfmt -h > /dev/null 2>&1 ; then
    echo "automatic install of 'jsonfmt' failed. please install it manually."
    exit 0 ;
  fi
  while read file;do
    echo "*** formatting $file"
    jsonfmt "$file" || true
    echo '' >> "$file"
  done < <(find -type f -not -path '*/\.git/*' -name '*.json')
# ────────────────────────────────────────────────────────────────────────────────
bootstrap-json: _format-json
    @echo json tools were installed
