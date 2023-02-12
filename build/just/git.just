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
# Detects the default git pager. If not set, it will fall back to 'less'
DIFF_PAGER := `[[ -n $(git config pager.diff ) ]] && echo "$(git config pager.diff)" || echo 'less'`
alias ga := git-add
# uses fzf to list git changes and help developers stage them
git-add:
    #!/usr/bin/env bash
    set -euo pipefail
    git rev-parse --is-inside-work-tree > /dev/null || return 1;
    [[ $# -ne 0 ]] && git add "$@" && git status -su && return;
    changed=$(git config --get-color color.status.changed red);
    unmerged=$(git config --get-color color.status.unmerged red);
    untracked=$(git config --get-color color.status.untracked red);
    _FZF_DEFAULT_OPTS="--multi --height=40% --reverse --tabstop=4 -0 --prompt=' │ ' --color=prompt:0,hl:178,hl+:178 --bind='ctrl-t:toggle-all,ctrl-g:select-all+accept' --bind='tab:down,shift-tab:up' --bind='?:toggle-preview,ctrl-space:toggle'
    --ansi
    --height='80%'
    --bind='alt-k:preview-up,alt-p:preview-up'
    --bind='alt-j:preview-down,alt-n:preview-down'
    --bind='ctrl-r:toggle-all'
    --bind='ctrl-s:toggle-sort'
    --bind='?:toggle-preview'
    --bind='alt-w:toggle-preview-wrap'
    --preview-window='right:60%'
    +1"
    extract="
        sed 's/^.*]  //' |
        sed 's/.* -> //' |
        sed -e 's/^\\\"//' -e 's/\\\"\$//'";
    preview="
        file=\$(echo {} | $extract)
        if (git status -s -- \$file | grep '^??') &>/dev/null; then  # diff with /dev/null for untracked files
            git diff --color=always --no-index -- /dev/null \$file | {{ DIFF_PAGER }} | sed '2 s/added:/untracked:/'
        else
            git diff --color=always -- \$file | {{ DIFF_PAGER }}
        fi";
    opts="
        $_FZF_DEFAULT_OPTS
        -0 -m --nth 2..,..
    ";
    files=$(git -c color.status=always -c status.relativePaths=true status -su |
        grep -F -e "$changed" -e "$unmerged" -e "$untracked" |
        sed -E 's/^(..[^[:space:]]*)[[:space:]]+(.*)$/[\1]  \2/' |
        FZF_DEFAULT_OPTS="$opts" fzf --preview="$preview" |
        sh -c "$extract");
    [[ -n "$files" ]] && echo "$files" | tr '\n' '\0' | xargs -0 -I% git add % && git status -su && exit ;
    echo 'Nothing to add.'
