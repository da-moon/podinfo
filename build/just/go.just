# !/usr/bin/env -S just --justfile
# vim: filetype=just tabstop=2 shiftwidth=2 softtabstop=2 expandtab:
# ────────────────────────────────────────────────────────────────────────────────
# ensure go toolchain is installed
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
# install mage and upx
_build: _go
  #!/usr/bin/env bash
  set -euo pipefail
  if ! upx --version > /dev/null 2>&1 ; then
    just _install-os-package "upx"
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
_lint-go: _go
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
alias gofmt := format-go
alias go-fmt := format-go
# format all go files
format-go: _go
    #!/usr/bin/env bash
    set -euo pipefail
    gofmt -l -w {{ justfile_directory() }}
# ────────────────────────────────────────────────────────────────────────────────
alias golangci-lint := lint-go
lint-go: _lint-go
    #!/usr/bin/env bash
    set -euo pipefail
    if ! golangci-lint --version > /dev/null 2>&1 ; then
      echo "automatic install of 'golangci-lint' failed. please install it manually."
      exit 0 ;
    fi
    golangci-lint run \
    --print-issued-lines=false \
    --exclude-use-default=false \
    --config "{{ justfile_directory() }}/.golangci.yml"
# ────────────────────────────────────────────────────────────────────────────────
# install all go toolings
bootstrap-go: _go _build _lint-go
  #!/usr/bin/env bash
  set -euo pipefail
  go env -w "GO111MODULE=on"
  go env -w "CGO_ENABLED=0"
  go env -w "CGO_LDFLAGS=-s -w -extldflags '-static'"
  if [ -r "{{ justfile_directory() }}/go.mod" ];then
      go clean -modcache
      go mod tidy
  fi
  if [ -r "{{ justfile_directory() }}/tools.go" ];then
    go generate -tags tools tools.go
  fi
