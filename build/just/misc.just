# !/usr/bin/env -S just --justfile
# vim: filetype=just tabstop=2 shiftwidth=2 softtabstop=2 expandtab:
# ────────────────────────────────────────────────────────────────────────────────
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
# ────────────────────────────────────────────────────────────────────────────────
alias vt := vscode-tasks
vscode-tasks:
    #!/usr/bin/env bash
    set -euo pipefail
    if command -- jq -h > /dev/null 2>&1 ; then
      IFS=' ' read -a TASKS <<< "$(just --summary --color never -f "{{ justfile() }}" 2>/dev/null)"
      if [ ${#TASKS[@]} -ne 0  ];then
        mkdir -p "{{ justfile_directory() }}/.vscode"
        json=$(jq -n --arg version "2.0.0" '{"version":$version,"tasks":[]}')
        for task in "${TASKS[@]}";do
          taskjson=$(jq -n --arg task "${task}" --arg command "just ${task}" '[{"type": "shell","label": $task,  "command": $command }]')
          json=$(echo "${json}" | jq ".tasks += ${taskjson}")
        done
        echo "${json}" | jq -r '.' > "{{ justfile_directory() }}/.vscode/tasks.json"
      fi
    fi
    just format-just
