# !/usr/bin/env -S just --justfile
# vim: filetype=just tabstop=2 shiftwidth=2 softtabstop=2 expandtab:
# ────────────────────────────────────────────────────────────────────────────────
# ensures dependencies for creating sane commit messages are installed
_pre-commit:
  #!/usr/bin/env bash
  set -euo pipefail
  IFS=':' read -a paths <<< "$(printenv PATH)" ;
  [[ ! " ${paths[@]} " =~ " ${HOME}/bin " ]] && export PATH="${PATH}:${HOME}/bin" || true ;
  if ! command -- commitlint -h > /dev/null 2>&1 ; then
    if command -- sudo -E npm -h > /dev/null 2>&1 ; then
      echo >&2 "*** npm (running with 'sudo') not found. Please install npm and try again."
      exit 1
    fi
    sudo npm i -g @commitlint/config-conventional @commitlint/cli
  fi
  if ! command -- pre-commit -h > /dev/null 2>&1 ; then
    curl "https://pre-commit.com/install-local.py" | "$(command -v python3)" -
  fi
  if ! command -- convco -h > /dev/null 2>&1 ; then
      if command -- cargo -h > /dev/null 2>&1 ; then
      cargo install -j `nproc` --locked --all-features --git "https://github.com/convco/convco.git"
    fi
  fi
