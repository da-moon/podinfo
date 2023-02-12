#!/usr/bin/env bash
# NOTE: simulate this pre-commit hook script:
#
# contrib/scripts/hooks/just-fmt.sh $(git diff --cached --name-only --diff-filter=ACMR | grep -E '\.just$')
#
# create justfile
cat <<'EOF' > Justfile.commit
# !/usr/bin/env -S just --justfile
# vim: filetype=just tabstop=2 shiftwidth=2 softtabstop=2 expandtab:
# ────────────────────────────────────────────────────────────────────────────────
EOF
for file in "$@"; do
sed \
  -e '/env -S just --justfile/d' \
  -e '/# vim: /d' "${file}" \
  -e '/# ────/d' | tee -a Justfile.commit > /dev/null
done
just --unstable --summary --justfile Justfile.commit > /dev/null 2>&1
if [ $? -eq 1 ]; then
  echo "There are errors in concatenated Justfile, please fix them before committing."
  echo "You can run the following to see error messages:"
  echo ""
  echo "just --unstable --summary --justfile Justfile.commit Justfile.commit"
  exit 1
fi
just --unstable --fmt --justfile Justfile.commit
if [ ! -r Justfile ] ;then
  mv Justfile.commit Justfile
  echo "Main Justfile was formed. Stage the file and commit again"
  exit 1
fi
diff -au Justfile Justfile.commit > /dev/null
if [ $? -eq 1 ]; then
  echo "There is a difference between the Original Justfile and updated one."
  echo "Stage Justfile and commit again"
  mv Justfile.commit Justfile
  exit 1
fi
exit 0
