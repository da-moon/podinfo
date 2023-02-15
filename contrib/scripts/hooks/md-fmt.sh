#!/usr/bin/env bash
# NOTE: simulate this pre-commit hook script:
#
# contrib/scripts/hooks/pre-commit.sh $(git diff --cached --name-only --diff-filter=ACM | grep -E '\.(md|markdown)$')
if [ -z "$(which sponge)" ] >/dev/null  2>&1; then
  echo "'sponge' not found, please install 'moreutils' before committing markdown files"
  exit 1
fi
if [ -z "$(which prettier)" ] >/dev/null  2>&1; then
  echo "'prettier' not found, please install it before committing markdown files"
  echo ""
  echo "sudo npm i -g prettier"
  exit 1
fi
if ! command -- remark -h >/dev/null  2>&1; then
  echo "'remark-cli' not found, please install it before committing markdown files"
  echo ""
  echo "sudo npm i -g remark remark-stringify remark-cli remark-reference-links remark-frontmatter remark-toc"
  exit 1
fi
to_fmt=()
for file in "$@"; do
  diff -au "${file}" <(remark \
    --use remark-toc='"heading": "Table of contents"' \
    --use remark-frontmatter \
    --use remark-reference-links \
    --use remark-stringify='bullet: "-",listItemIndent: "one"' \
    "${file}" 2>/dev/null |
      prettier --print-width=79 --prose-wrap=always --parser markdown) >/dev/null
    #
  if [ $? -eq 1 ]; then
    to_fmt+=(${file})
  fi
done
if ((${#to_fmt[@]})); then
  for file in "${to_fmt[@]}"; do
    remark \
      --use remark-toc='"heading": "Table of contents"' \
      --use remark-frontmatter \
      --use remark-reference-links \
      --use remark-stringify='bullet: "-",listItemIndent: "one"' \
      "${file}" |
      prettier --print-width=79 --prose-wrap=always --parser markdown |
      sponge "${file}"
  done
  exit 1
else
  exit 0
fi
