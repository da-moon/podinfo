# !/usr/bin/env -S just --justfile
# vim: filetype=just tabstop=2 shiftwidth=2 softtabstop=2 expandtab:
# ────────────────────────────────────────────────────────────────────────────────
# ensures 'shfmt' is installed
_format-bash:
  #!/usr/bin/env bash
  set -euo pipefail
  if ! shfmt --version > /dev/null 2>&1 ; then
    echo "*** shfmt not found. installing ..." ;
    go install "mvdan.cc/sh/v3/cmd/shfmt@latest" ;
  fi
# ────────────────────────────────────────────────────────────────────────────────
# ensures bash linter is installed
_lint-bash:
  #!/usr/bin/env bash
  set -euo pipefail
  if ! shellcheck --version > /dev/null 2>&1 ; then
    just _install-os-package "shellcheck" ;
  fi
# ────────────────────────────────────────────────────────────────────────────────
alias shfmt := format-bash
alias bash-fmt := format-bash
# detect and format all bash scripts
format-bash: _format-bash
    #!/usr/bin/env bash
    if ! command -- shfmt --version > /dev/null 2>&1 ; then
      echo "automatic install of 'shfmt' failed. please install it manually."
      exit 0 ;
    fi
    targets=($(find . \
        -type f \
        -not -path '*/\.git/*' \
        -exec grep -Il '.' {} \; \
        | xargs -r -P 0 -I {} \
        gawk 'FNR>4 {nextfile} /#!.*sh/ { print FILENAME ; nextfile }' {})) ;
    if [ ${#targets[@]} -ne 0  ];then
        for target in "${targets[@]}";do
            chmod +x "${target}" ;
            shfmt -kp -i 2 -ci -w "${target}" ;
        done
    fi
# ────────────────────────────────────────────────────────────────────────────────
alias shellcheck := lint-bash
lint-bash: _lint-bash
    #!/usr/bin/env bash
    set -euo pipefail
    if ! shellcheck --version > /dev/null 2>&1 ; then
      echo "automatic install of 'shellcheck' failed. please install it manually."
      exit 0 ;
    fi
    targets=($(find . \
        -type f \
        -not -path '*/\.git/*' \
        -exec grep -Il '.' {} \; \
        | xargs -r -P 0 -I {} \
        gawk 'FNR>4 {nextfile} /#!.*sh/ { print FILENAME ; nextfile }' {})) ;
    if [ ${#targets[@]} -ne 0  ];then
        for target in "${targets[@]}";do
            shellcheck "${target}" || true ;
        done
    fi

# ────────────────────────────────────────────────────────────────────────────────
# install all bash toolings
bootstrap-bash: _format-bash _lint-bash
    @echo bash tools were installed
