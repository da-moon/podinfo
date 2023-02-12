# !/usr/bin/env -S just --justfile
# vim: filetype=just tabstop=2 shiftwidth=2 softtabstop=2 expandtab:
# ────────────────────────────────────────────────────────────────────────────────
# ensures 'git-delta' is installed and set as the default pager for git
_git-delta:
  #!/usr/bin/env bash
  set -euo pipefail
  if  ! command -- delta --version > /dev/null 2>&1 ; then
    just _install-rust-package git-delta
  fi
  if  command -- delta --version > /dev/null 2>&1 ; then
    git config --global core.autocrlf false
    git config --global pager.diff delta
    git config --global pager.log delta
    git config --global pager.reflog delta
    git config --global pager.show delta
    git config --global interactive.difffilter "delta --color-only --features=interactive"
    git config --global delta.features "side-by-side line-numbers decorations"
    git config --global delta.whitespace-error-style "22 reverse"
    git config --global delta.decorations.commit-decoration-style "bold yellow box ul"
    git config --global delta.decorations.file-style "bold yellow ul"
    git config --global delta.decorations.file-decoration-style "none"
    git config --global delta.decorations.commit-style "raw"
    git config --global delta.decorations.hunk-header-decoration-style "blue box"
    git config --global delta.decorations.hunk-header-file-style "red"
    git config --global delta.decorations.hunk-header-line-number-style "#067a00"
    git config --global delta.decorations.hunk-header-style "file line-number syntax"
    git config --global delta.interactive.keep-plus-minus-markers "false"
  fi

alias gf := git-fetch
# fetches latest changes from upstream and removes any local branches that have been deleted in upstream
git-fetch:
    #!/usr/bin/env bash
    set -euo pipefail
    pushd "{{ justfile_directory() }}" > /dev/null 2>&1
    git fetch -p ;
    for branch in $(git branch -vv | grep ': gone]' | grep -v '*' | awk '{print $1}'); do
      git branch -D "$branch";
    done
    popd > /dev/null 2>&1
