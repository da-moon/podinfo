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
# ────────────────────────────────────────────────────────────────────────────────
alias pc := bootstrap-pre-commit
# ensures tools for making sane commits are installed and initializes pre-commit
bootstrap-pre-commit: _pre-commit
    #!/usr/bin/env bash
    set -euo pipefail
    IFS=':' read -a paths <<< "$(printenv PATH)" ;
    [[ ! " ${paths[@]} " =~ " ${HOME}/bin " ]] && export PATH="${PATH}:${HOME}/bin" || true ;
    pushd "{{ justfile_directory() }}" > /dev/null 2>&1
    if [ -r .pre-commit-config.yaml ]; then
      pre-commit autoupdate
      git add ".pre-commit-config.yaml"
      pre-commit install > /dev/null 2>&1
      pre-commit install --install-hooks
      pre-commit
    fi
    popd > /dev/null 2>&1
# ────────────────────────────────────────────────────────────────────────────────
alias c := commit
# help guide the developers make conventional commits. it is recommended to use this target to make new commits
commit: git-fetch bootstrap-pre-commit
    #!/usr/bin/env bash
    set -euo pipefail
    pushd "{{ justfile_directory() }}" > /dev/null 2>&1
    # echo 'hello world' | commitlint -x @commitlint/config-conventional
    if command -- convco -h > /dev/null 2>&1 ; then
      convco commit
    else
      git commit
    fi
    popd > /dev/null 2>&1
