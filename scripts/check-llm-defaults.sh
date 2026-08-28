#!/usr/bin/env bash
# Warn when a change touches the LLM role defaults or the price table.
# Advisory only (guardrail 3): a model swap belongs in a bake-off (PR-15).
# In GitHub Actions the warning becomes an annotation. It never fails.
set -u
# An array, so the paths reach git as two arguments and not one word
# that the shell must split (SC2086).
FILES=(go/internal/llm/roles.json go/internal/llm/prices.json)
base="${LLM_DEFAULTS_BASE:-}"
if [ -z "$base" ]; then
  if [ -n "${GITHUB_BASE_REF:-}" ]; then
    git fetch -q origin "$GITHUB_BASE_REF" 2>/dev/null || true
    base="origin/$GITHUB_BASE_REF"
  else
    base="origin/main"
  fi
fi
mb=$(git merge-base "$base" HEAD 2>/dev/null) || { echo "check-llm-defaults: no merge base with $base, skip"; exit 0; }
changed=$(git diff --name-only "$mb" HEAD -- "${FILES[@]}")
if [ -z "$changed" ]; then
  echo "check-llm-defaults: no LLM default change"
  exit 0
fi
for f in $changed; do
  msg="LLM default changed in $f. A model or price swap needs a bake-off before merge (PR-15) and a dated verified_at."
  if [ -n "${GITHUB_ACTIONS:-}" ]; then
    echo "::warning file=$f::$msg"
  else
    echo "WARNING: $msg"
  fi
  git diff "$mb" HEAD -- "$f" | head -80
done
exit 0
