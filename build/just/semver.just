# !/usr/bin/env -S just --justfile
# vim: filetype=just tabstop=2 shiftwidth=2 softtabstop=2 expandtab:
# ────────────────────────────────────────────────────────────────────────────────
# stores upstream master branch name
MASTER_BRANCH_NAME := 'master'
# this variable stores the next major release tag
MAJOR_VERSION := `[[ -n $(git tag -l | head -n 1 ) ]] && convco version --major 2>/dev/null || echo '0.0.1'`
# this variable stores the next minor release tag
MINOR_VERSION := `[[ -n $(git tag -l | head -n 1 ) ]] && convco version --minor 2>/dev/null || echo '0.0.1'`
# this variable stores the next patch release tag
PATCH_VERSION := `[[ -n $(git tag -l | head -n 1 ) ]] && convco version --patch 2>/dev/null || echo '0.0.1'`
