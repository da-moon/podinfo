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

# this target updates all os packages . supports (debian, arch, alpine)
_update-os-pkgs:
    #!/usr/bin/env bash
    set -euo pipefail
    if  command -- apt -h > /dev/null 2>&1 ; then
      echo >&2 "*** Debian based distribution detected."
      export DEBIAN_FRONTEND=noninteractive
      sudo apt-get update -qq
      sudo apt-get -f install -y
      sudo apt-get upgrade -yq
      sudo apt-get autoremove --purge -y
    elif command -- pacman --version > /dev/null 2>&1 ; then
      echo >&2 "*** Arch Linux based distribution detected."
      echo >&2 "*** updating official Arch packages with pacman."
      sudo pacman -Syyu --noconfirm || true ;
    elif sudo apk --version > /dev/null 2>&1 ; then
      echo >&2 "*** Alpine Linux based distribution detected."
      echo >&2 "*** updating Alpine packages with apk."
      sudo apk update && sudo apk update
    else
      echo >&2 "*** Your Operating system is not supported."
    fi

# this target installs an os package. supports (debian, arch, alpine)
_install-os-package pkg:
    #!/usr/bin/env bash
    set -euo pipefail
    if  command -- apt -h > /dev/null 2>&1 ; then
      PKG_OK=$((dpkg-query -W --showformat='${Status}\n' {{ pkg }} || true )|(grep "install ok installed" || true))
      if [ "" = "$PKG_OK" ]; then
        sudo apt-get -yq install {{ pkg }}
      else
        echo >&2 "*** '{{ pkg }}' has already been installed.skipping "
      fi
    elif command -- pacman --version > /dev/null 2>&1 ; then
      if ! pacman -Qi "{{ pkg }}" > /dev/null 2>&1 ; then
        sudo pacman -Sy --needed --noconfirm {{ pkg }} || true ;
      else
        echo >&2 "*** '{{ pkg }}' has already been installed.skipping "
      fi
    elif sudo apk --version > /dev/null 2>&1 ; then
      if ! sudo apk info -L "{{ pkg }}" > /dev/null 2>&1 ; then
        sudo apk update && sudo apk add --no-cache {{ pkg }}
      else
        echo >&2 "*** '{{ pkg }}' has already been installed.skipping "
      fi
    else
      echo >&2 "*** Your Operating system is not supported."
      exit 1
    fi

# this target validates and installs nodejs and additional dependencies
_bootstrap-nodejs:
    #!/usr/bin/env bash
    set -euo pipefail
    if ! command -- $(which node) --version > /dev/null 2>&1 ; then
      echo >&2 "*** nodejs is required."
      exit 1
    else
      echo >&2 "*** Node.JS installation has been validated."
    fi
    if ! command -- $(which npm) --version > /dev/null 2>&1 ; then
      echo >&2 "*** npm is required."
      exit 1
    else
      echo >&2 "*** npm installation has been validated."
    fi
    if ! command -- $(which yarn) --version > /dev/null 2>&1 ; then
      echo >&2 "*** yarn not found. installing"
      sudo npm install -g yarn
      exit 1
    else
      echo >&2 "*** yarn installation has been validated."
    fi

# this target uses npm to install a nodejs package globally
_install-nodejs-package pkg:
    #!/usr/bin/env bash
    set -euo pipefail
    if ! command -- $(which npm) --version > /dev/null 2>&1 ; then
      just _bootstrap-nodejs
    else
      true
    fi
    sudo $(which npm) install -g {{ pkg }}

# This target is used to validate that the rust toolchain is installed
_validate-rust:
    #!/usr/bin/env bash
    set -euo pipefail
    if ! command -- rustup -h > /dev/null 2>&1 ; then
      echo >&2 "*** rustup is required."
      exit 1
    else
      true
    fi
    if ! command -- cargo -h > /dev/null 2>&1 ; then
      echo >&2 "*** cargo is required."
      exit 1
    else
      true
    fi

# This target ensures rust toolchain and all the tools compiled and installed with it are up to date
_update-rust: _validate-rust
    #!/usr/bin/env bash
    set -euo pipefail
    echo >&2 "*** ensuring rustup has been updated."
    rustup update >/dev/null 2>&1
    echo >&2 "*** ensuring rust nightly and stable toolchains are installed."
    rustup toolchain install nightly stable >/dev/null 2>&1
    rustup default stable
    if ! command -- cargo-install-update -h >/dev/null 2>&1; then
      just _install-rust-package cargo-update
    else
      true
    fi
    echo >&2 "*** ensuring all installed rust-based command line utilities, compiled with stable toolchain, have been updated to latest versions"
    cargo-install-update install-update --all || true
    rustup default nightly
    echo >&2 "*** ensuring all installed rust-based command line utilities, compiled with nightly toolchain, have been updated to latest versions"
    cargo-install-update install-update --all || true
    rustup default stable

# this target builds and installs a rust package from source
_install-rust-package name:
    #!/usr/bin/env bash
    set -euo pipefail
    if  ! command -- cargo --version > /dev/null 2>&1 ; then
        echo >&2 "*** cannot install '{{ name }}' as rust toolchain has not been installed"
        exit 1
    else
        true
    fi
    installed_packages=($(cargo install --list | awk '/:/{print $1}'))
    mkdir -p {{ justfile_directory() }}/tmp
    rm -rf {{ justfile_directory() }}/tmp/rust-fail.txt
    if [[ ! " ${installed_packages[@]} " =~ " {{ name }} " ]]; then
        echo >&2 "***  building and installing '{{ name }}' from source ..."
        cargo install -j `nproc` --locked --all-features '{{ name }}' || (echo '{{ name }}' >> {{ justfile_directory() }}/tmp/rust-fail.txt ; true)
    else
        echo >&2 "***  '{{ name }}' installation detected. Skipping build ..."
    fi

alias kc := kary-comments

# adds support for extra languages to Kary Comments VSCode extension
kary-comments:
    #!/usr/bin/env bash
    set -euo pipefail
    path_pattern="*/karyfoundation.comment*/dictionary.js";
    while read path; do
      if test -n "${path}"; then
        sed "/shellscript/r"<( \
        leading_whitespaces="$(grep -Po "[[:space:]]+(?=case 'shellscript':)" "${path}")";
          language='terraform'; ! grep -q "case '${language}':" "${path}" && (echo -n "${leading_whitespaces}";echo "case '${language}':" );
          language='dockerfile'; ! grep -q "case '${language}':" "${path}" && (echo -n "${leading_whitespaces}";echo "case '${language}':" );
          language='just'; ! grep -q "case '${language}':" "${path}" && (echo -n "${leading_whitespaces}";echo "case '${language}':" );
          language='hcl'; ! grep -q "case '${language}':" "${path}" && (echo -n "${leading_whitespaces}";echo "case '${language}':" );
          language='packer'; ! grep -q "case '${language}':" "${path}" && (echo -n "${leading_whitespaces}";echo "case '${language}':" );
        ) -i -- "${path}" ;
      fi ;
    done <<< "$(find "${HOME}" -type f -path "${path_pattern}" 2>/dev/null || true )" ;

# this target installs a collection of core os packages. supports (debian, arch, alpine)
_core-pkgs: _update-os-pkgs
    #!/usr/bin/env bash
    set -euo pipefail
    core_dependencies=()
    core_dependencies+=("jq")
    core_dependencies+=("parallel")
    core_dependencies+=("cmake")
    core_dependencies+=("make")
    core_dependencies+=("git")
    core_dependencies+=("fzf")
    core_dependencies+=("sshpass")
    core_dependencies+=("bash-completion")
    core_dependencies+=("pandoc")
    core_dependencies+=("texmaker")
    core_dependencies+=("ripgrep")
    core_dependencies+=("exa")
    core_dependencies+=("shellcheck")
    core_dependencies+=("moreutils")
    if command -- apt -h > /dev/null 2>&1 ; then
      core_dependencies+=("python3-distutils")
      core_dependencies+=("pdftk")
      core_dependencies+=("libgconf-2-4")
      core_dependencies+=("libssl-dev")
      core_dependencies+=("golang")
      core_dependencies+=("build-essential")
      core_dependencies+=("software-properties-common")
      core_dependencies+=("poppler-utils")
      core_dependencies+=("librsvg2-bin")
      core_dependencies+=("lmodern")
      core_dependencies+=("fonts-symbola")
      core_dependencies+=("xfonts-utils ")
      core_dependencies+=("texlive-xetex")
      core_dependencies+=("texlive-fonts-recommended")
      core_dependencies+=("texlive-fonts-extra")
      core_dependencies+=("texlive-latex-extra")
    fi
    if command -- pacman --version > /dev/null 2>&1 ; then
      core_dependencies+=("glow")
      core_dependencies+=("pdftk")
      core_dependencies+=("yarn")
      core_dependencies+=("npm")
      core_dependencies+=("nodejs")
      core_dependencies+=("pacman-contrib")
      core_dependencies+=("expac")
      core_dependencies+=("base-devel")
      core_dependencies+=("go")
      core_dependencies+=("poppler")
      core_dependencies+=("librsvg")
      core_dependencies+=("xorg-xfontsel")
      core_dependencies+=("texlive-most")
      core_dependencies+=("git-delta")
      core_dependencies+=("python-pre-commit")
    fi
    if sudo apk --version > /dev/null 2>&1 ; then
      core_dependencies+=("glow")
      core_dependencies+=("yarn")
      core_dependencies+=("npm")
      core_dependencies+=("nodejs")
      core_dependencies+=("go")
      core_dependencies+=("delta")
      core_dependencies+=("shellcheck-doc")
      core_dependencies+=("pre-commit")
    fi
    if [ ${#core_dependencies[@]} -ne 0  ]; then
      for dep in "${core_dependencies[@]}"; do
        just _install-os-package "${dep}"
      done
    else
      true
    fi

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
      pre-commit --install-hooks
      pre-commit
    fi
    popd > /dev/null 2>&1

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

# stores upstream master branch name

MASTER_BRANCH_NAME := 'master'

# this variable stores the next major release tag

MAJOR_VERSION := `[[ -n $(git tag -l | head -n 1 ) ]] && convco version --major 2>/dev/null || echo '0.0.1'`

# this variable stores the next minor release tag

MINOR_VERSION := `[[ -n $(git tag -l | head -n 1 ) ]] && convco version --minor 2>/dev/null || echo '0.0.1'`

# this variable stores the next patch release tag

PATCH_VERSION := `[[ -n $(git tag -l | head -n 1 ) ]] && convco version --patch 2>/dev/null || echo '0.0.1'`
