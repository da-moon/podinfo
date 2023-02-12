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
# ────────────────────────────────────────────────────────────────────────────────
alias mar := major-release
# generate changelog and create and push a new major release tag
major-release: git-fetch
  #!/usr/bin/env bash
  set -euo pipefail
  IFS=':' read -a paths <<< "$(printenv PATH)" ;
  [[ ! " ${paths[@]} " =~ " ${HOME}/bin " ]] && export PATH="${PATH}:${HOME}/bin" || true;
  pushd "{{ justfile_directory() }}" > /dev/null 2>&1
  git checkout "{{ MASTER_BRANCH_NAME }}"
  git pull
  git tag -a "v{{ MAJOR_VERSION }}" -m 'major release {{ MAJOR_VERSION }}'
  git push origin --tags
  if command -- convco -h > /dev/null 2>&1 ; then
    convco changelog > CHANGELOG.md
    git add CHANGELOG.md
    if command -- pre-commit -h > /dev/null 2>&1 ; then
      pre-commit || true
      git add CHANGELOG.md
    fi
    git commit -m 'docs(changelog): updated changelog for v{{ MAJOR_VERSION }}'
    git push
  fi
  just git-fetch
  popd > /dev/null 2>&1
# ────────────────────────────────────────────────────────────────────────────────
# generate changelog and create and push a new minor release tag
minor-release: git-fetch
  #!/usr/bin/env bash
  set -euo pipefail
  IFS=':' read -a paths <<< "$(printenv PATH)" ;
  [[ ! " ${paths[@]} " =~ " ${HOME}/bin " ]] && export PATH="${PATH}:${HOME}/bin" || true;
  pushd "{{ justfile_directory() }}" > /dev/null 2>&1
  git checkout "{{ MASTER_BRANCH_NAME }}"
  git pull
  git tag -a "v{{ MINOR_VERSION }}" -m 'minor release {{ MINOR_VERSION }}'
  git push origin --tags
  if command -- convco -h > /dev/null 2>&1 ; then
    convco changelog > CHANGELOG.md
    git add CHANGELOG.md
    if command -- pre-commit -h > /dev/null 2>&1 ; then
      pre-commit || true
      git add CHANGELOG.md
    fi
    git commit -m 'docs(changelog): updated changelog for v{{ MINOR_VERSION }}'
    git push
    just git-fetch
  fi
  popd > /dev/null 2>&1
