# !/usr/bin/env -S just --justfile
# vim: filetype=just tabstop=2 shiftwidth=2 softtabstop=2 expandtab:
# ────────────────────────────────────────────────────────────────────────────────
MASTER_BRANCH_NAME := 'master'
MAJOR_VERSION := `[[ -n $(git tag -l | head -n 1 ) ]] && convco version --major 2>/dev/null || echo '0.0.1'`
MINOR_VERSION := `[[ -n $(git tag -l | head -n 1 ) ]] && convco version --minor 2>/dev/null || echo '0.0.1'`
PATCH_VERSION := `[[ -n $(git tag -l | head -n 1 ) ]] && convco version --patch 2>/dev/null || echo '0.0.1'`
