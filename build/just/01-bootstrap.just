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
