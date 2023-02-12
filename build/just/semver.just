# !/usr/bin/env -S just --justfile
# vim: filetype=just tabstop=2 shiftwidth=2 softtabstop=2 expandtab:
# ────────────────────────────────────────────────────────────────────────────────
bootstrap-semver:
  #!/usr/bin/env bash
  set -euo pipefail
  if ! command -- convco -h > /dev/null 2>&1 ; then
      if command -- cargo -h > /dev/null 2>&1 ; then
      cargo install -j `nproc` --locked --all-features --git "https://github.com/convco/convco.git"
    fi
  fi

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
major-release: git-fetch bootstrap-semver
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
minor-release: git-fetch bootstrap-semver
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
# ────────────────────────────────────────────────────────────────────────────────
alias pr := patch-release
# generate changelog and create and push a new patch release tag
patch-release: git-fetch bootstrap-semver
  #!/usr/bin/env bash
  set -euo pipefail
  IFS=':' read -a paths <<< "$(printenv PATH)" ;
  [[ ! " ${paths[@]} " =~ " ${HOME}/bin " ]] && export PATH="${PATH}:${HOME}/bin" || true;
  pushd "{{ justfile_directory() }}" > /dev/null 2>&1
  git checkout "{{ MASTER_BRANCH_NAME }}"
  git pull
  git tag -a "v{{ PATCH_VERSION }}" -m 'patch release {{ PATCH_VERSION }}'
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
  fi
  just git-fetch
  popd > /dev/null 2>&1
# ────────────────────────────────────────────────────────────────────────────────
alias gc := generate-changelog
# generate markdown and pdf changelog files
generate-changelog: bootstrap-semver
  #!/usr/bin/env bash
  set -euo pipefail
  rm -rf "{{ justfile_directory() }}/tmp"
  mkdir -p "{{ justfile_directory() }}/tmp"
  convco changelog > "{{ justfile_directory() }}/tmp/$(basename {{ justfile_directory() }})-changelog-$(date -u +%Y-%m-%d).md"
  if command -- pandoc -h >/dev/null 2>&1; then
  pandoc \
    --from markdown \
    --pdf-engine=xelatex \
    -o  "{{ justfile_directory() }}/tmp/$(basename {{ justfile_directory() }})-changelog-$(date -u +%Y-%m-%d).pdf" \
    "{{ justfile_directory() }}/tmp/$(basename {{ justfile_directory() }})-changelog-$(date -u +%Y-%m-%d).md"
  fi
  if [ -d /workspace ]; then
    cp -f "{{ justfile_directory() }}/tmp/$(basename {{ justfile_directory() }})-changelog-$(date -u +%Y-%m-%d).pdf" /workspace/
    cp -f "{{ justfile_directory() }}/tmp/$(basename {{ justfile_directory() }})-changelog-$(date -u +%Y-%m-%d).md" /workspace/
  fi
