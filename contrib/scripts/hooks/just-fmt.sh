#!/usr/bin/env bash
# NOTE: simulate this pre-commit hook script:
#
# contrib/scripts/hooks/just-fmt.sh $(git diff --cached --name-only --diff-filter=ACMR | grep -E 'Justfile|\.just$')
to_fmt=()
for file in "$@"; do
  just --unstable --summary --justfile "${file}" > /dev/null 2>&1
  if [ $? -eq 1 ]; then
    echo "There are errors in '${file}', please fix them before committing"
    exit 1
  fi
  just --unstable --fmt --check --justfile "${file}" > /dev/null 2>&1
  if [ $? -eq 1 ]; then
    to_fmt+=(${file})
  fi
done
if (( ${#to_fmt[@]} )); then
  for file in "${to_fmt[@]}"; do
  echo "+++ Formatting '${file}'"
  just --unstable --fmt --justfile "${file}" > /dev/null 2>&1
  done
  echo "Justfile re-formatted, please re-add it to your commit"
  exit 1
fi
