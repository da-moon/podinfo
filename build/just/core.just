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
