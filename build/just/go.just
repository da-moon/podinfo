# !/usr/bin/env -S just --justfile
# vim: filetype=just tabstop=2 shiftwidth=2 softtabstop=2 expandtab:
# ────────────────────────────────────────────────────────────────────────────────
# ensures go toolchain is installed
_go:
  #!/usr/bin/env bash
  set -euo pipefail
  if ! go version > /dev/null 2>&1 ; then
    dep="go"
    if command -- apt -h > /dev/null 2>&1 ; then
      dep="golang"
    fi
    just _install-os-package "${dep}"
  fi
# ────────────────────────────────────────────────────────────────────────────────
_build:
  #!/usr/bin/env bash
  set -euo pipefail
  if ! upx --version > /dev/null 2>&1 ; then
    just _install-os-package "${dep}"
  fi
  if ! mage --version > /dev/null 2>&1 ; then
    echo "*** mage not found. installing ..." ;
    sudo rm -rf "/tmp/mage"
    pushd "/tmp/mage" > /dev/null 2>&1
      git clone "https://github.com/magefile/mage" "/tmp/mage"
      go run bootstrap.go
      sudo rm -rf "/tmp/mage"
    popd > /dev/null 2>&1
  fi
# ────────────────────────────────────────────────────────────────────────────────
# ensure golangci-lint is installed
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
# ────────────────────────────────────────────────────────────────────────────────
