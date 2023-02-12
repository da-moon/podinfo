# !/usr/bin/env -S just --justfile
# vim: filetype=just tabstop=2 shiftwidth=2 softtabstop=2 expandtab:
# ────────────────────────────────────────────────────────────────────────────────

set positional-arguments := true
set dotenv-load := true
set shell := ["/bin/bash", "-o", "pipefail", "-c"]

project_name := `basename $PWD`

default:
    @just --choose
