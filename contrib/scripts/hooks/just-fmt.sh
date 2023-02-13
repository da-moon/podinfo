#!/usr/bin/env bash
# NOTE: simulate this pre-commit hook script:
#
# contrib/scripts/hooks/just-fmt.sh
#
# create justfile
WD="$(cd "$(dirname "${BASH_SOURCE[0]}")/../../../" && pwd)"
pushd "$WD" >/dev/null 2>&1
rm -f Justfile.commit
cat <<'EOF' >Justfile.commit
# !/usr/bin/env -S just --justfile
# vim: filetype=just tabstop=2 shiftwidth=2 softtabstop=2 expandtab:
# ────────────────────────────────────────────────────────────────────────────────
EOF
while read file; do
  sed \
    -e '/env -S just --justfile/d' \
    -e '/# vim: /d' \
    -e '/# ────/d' \
    "${file}" |
    tee -a Justfile.commit >/dev/null
done < <(find . -name '*.just' -type f)
just --unstable --summary --justfile Justfile.commit >/dev/null 2>&1
if [ $? -eq 1 ]; then
  echo "There are errors in concatenated Justfile, please fix them before committing."
  echo "You can run the following to see error messages:"
  echo ""
  echo "just --unstable --summary --justfile Justfile.commit"
  exit 1
fi
just --unstable --fmt --justfile Justfile.commit
if [ ! -r Justfile ]; then
  mv Justfile.commit Justfile
  echo "Main Justfile was formed. Stage the file and commit again"
  exit 1
fi
git checkout HEAD -- Justfile
diff -au Justfile Justfile.commit >/dev/null
if [ $? -eq 1 ]; then
  echo "There is a difference between the Original Justfile and updated one."
  echo "Stage Justfile and commit again"
  mv Justfile.commit Justfile
  exit 1
fi
rm -f Justfile.commit
exit 0
popd >/dev/null 2>&1
