# !/usr/bin/env -S just --justfile
# vim: filetype=just tabstop=2 shiftwidth=2 softtabstop=2 expandtab:
# ────────────────────────────────────────────────────────────────────────────────
# this target installs a collection of core os packages. supports (debian, arch, alpine)
bootstrap-os-pkgs: _update-os-pkgs
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
    core_dependencies+=("moreutils")
    if command -- apt -h > /dev/null 2>&1 ; then
      core_dependencies+=("python3-distutils")
      core_dependencies+=("pdftk")
      core_dependencies+=("libgconf-2-4")
      core_dependencies+=("libssl-dev")
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
      core_dependencies+=("cargo")
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
      core_dependencies+=("poppler")
      core_dependencies+=("librsvg")
      core_dependencies+=("xorg-xfontsel")
      core_dependencies+=("texlive-most")
      core_dependencies+=("git-delta")
      core_dependencies+=("python-pre-commit")
      core_dependencies+=("rustup")
    fi
    if sudo apk --version > /dev/null 2>&1 ; then
      core_dependencies+=("glow")
      core_dependencies+=("yarn")
      core_dependencies+=("npm")
      core_dependencies+=("nodejs")
      core_dependencies+=("delta")
      core_dependencies+=("pre-commit")
      core_dependencies+=("rust")
    fi
    if [ ${#core_dependencies[@]} -ne 0  ]; then
      for dep in "${core_dependencies[@]}"; do
        just _install-os-package "${dep}"
      done
    else
      true
    fi
# ────────────────────────────────────────────────────────────────────────────────
alias b := bootstrap
# installs dependencies and prepares development environment
bootstrap: bootstrap-os-pkgs bootstrap-git bootstrap-semver bootstrap-pre-commit bootstrap-go bootstrap-bash bootstrap-json bootstrap-markdown
    @echo all developer tools were installed
