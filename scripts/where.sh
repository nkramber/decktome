#!/usr/bin/env bash
# Print where the checkout stands, before a commit, a push, or a deploy.
# Every question the session got wrong on 2026-09-07 is a line here.
set -uo pipefail

git fetch --quiet origin 2>/dev/null || true

branch=$(git rev-parse --abbrev-ref HEAD)
dirty=$(git status --porcelain)
printf 'branch          %s\n' "$branch"
printf 'tree            %s\n' "$([ -z "$dirty" ] && echo clean || echo "DIRTY, $(echo "$dirty" | wc -l | tr -d ' ') files")"

if git rev-parse --abbrev-ref '@{u}' >/dev/null 2>&1; then
  read -r behind ahead < <(git rev-list --left-right --count '@{u}...HEAD' | awk '{print $1, $2}')
  printf 'vs upstream     behind %s, ahead %s\n' "$behind" "$ahead"
else
  printf 'vs upstream     none, this branch is not pushed\n'
fi

read -r mahead mbehind < <(git rev-list --left-right --count "main...origin/main" | awk '{print $1, $2}')
printf 'main            %s\n' "$([ "$mbehind" = 0 ] && echo "current with origin" || echo "BEHIND origin by $mbehind")"

if [ "$branch" != "main" ]; then
  printf 'unmerged here   %s commits not in origin/main\n' "$(git rev-list --count origin/main..HEAD)"
fi

if command -v gh >/dev/null 2>&1; then
  pr=$(gh pr list --head "$branch" --state all --json number,state --jq 'if length == 0 then "" else "#\(.[0].number) \(.[0].state)" end' 2>/dev/null || echo "")
  printf 'pull request    %s\n' "${pr:-none}"
  case "$pr" in
    *MERGED*) printf '\nWARNING: this branch is merged. A commit here never reaches main.\n' ;;
  esac
fi

if [ "$branch" = "main" ]; then
  printf '\nWARNING: main takes no commit (D-583). Start a branch first.\n'
fi
