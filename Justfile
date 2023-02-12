# !/usr/bin/env -S just --justfile
# vim: filetype=just tabstop=2 shiftwidth=2 softtabstop=2 expandtab:
# ────────────────────────────────────────────────────────────────────────────────
# this is needed for properly passing user input arguments to just targets

set positional-arguments := true

# loads environment variables from .env

set dotenv-load := true

# sets shell to bash, and enables pipefail

set shell := ["/bin/bash", "-o", "pipefail", "-c"]

# sets project name to the name of the current directory

project_name := `basename $PWD`

# `default` target, i.e target execued when just is called without any arguments
default:
    @just --choose
