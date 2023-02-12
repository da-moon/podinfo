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

# this target updates all os packages
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
