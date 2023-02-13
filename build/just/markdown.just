# !/usr/bin/env -S just --justfile
# vim: filetype=just tabstop=2 shiftwidth=2 softtabstop=2 expandtab:
# ────────────────────────────────────────────────────────────────────────────────
# installs remark-cli, prettier, and markdown-magic
_format-markdown:
  #!/usr/bin/env bash
  set -euo pipefail
  if [ -z "$(which sponge)" ] > /dev/null 2>&1 ; then
    echo "*** 'sponge' not found. installing ..." ;
    just _install-os-package "moreutils" ;
  fi
  if ! command -- remark -h > /dev/null 2>&1 ; then
    echo "*** 'remark-cli' not found. installing ..." ;
    sudo npm i -g remark remark-stringify remark-cli remark-reference-links remark-frontmatter remark-toc ;
  fi
  if ! command -- prettier -h > /dev/null 2>&1 ; then
    echo "*** 'prettier' not found. installing ..." ;
    sudo npm i -g prettier ;
  fi
  if ! command -- md-magic -h > /dev/null 2>&1 ; then
    echo "*** 'markdown-magic' not found. installing ..." ;
    sudo npm i -g markdown-magic ;
  fi
  if ! command -- cspell-cli -h > /dev/null 2>&1 ; then
    echo "*** 'markdown-magic' not found. installing ..." ;
    sudo npm i -g cspell-cli ;
  fi


# ────────────────────────────────────────────────────────────────────────────────
# install all markdown toolings
bootstrap-markdown: _format-markdown
  @echo bash tools were installed
