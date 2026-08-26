#!/usr/bin/env bash
# One fix step of the tuning loop (D-133). It builds the prompt for a
# fixer agent and runs whatever command the owner named.
#
# The loop calls this script. It reads two variables:
#   AUTOTUNE_EVAL_DOC  the eval report to act on
#   AUTOTUNE_LABEL     the iteration label, for logs
#
# AUTOTUNE_FIXER_CMD names the agent. This script ships no default on
# purpose. An unattended agent that edits a repository is the owner's
# call, and it must be written down where the owner can read it, not
# buried in a script. See docs/reference/autotune-design.md.
#
# The command reads the prompt on standard input and edits the working
# tree in place. autotune.sh checks the result: a frozen path, a red
# build, or a worse run all revert the change.
set -u
cd "$(dirname "$0")/.." || exit 1
ROOT="$(pwd)"

DOC="${AUTOTUNE_EVAL_DOC:-}"
LABEL="${AUTOTUNE_LABEL:-auto}"
[ -f "$DOC" ] || { echo "autotune-fix: no eval report at $DOC" >&2; exit 1; }

if [ -z "${AUTOTUNE_FIXER_CMD:-}" ]; then
  echo "autotune-fix: set AUTOTUNE_FIXER_CMD to the agent that applies fixes." >&2
  echo "autotune-fix: docs/reference/autotune-design.md names the trade-off." >&2
  exit 1
fi

PROMPT="$(mktemp)"
trap 'rm -f "$PROMPT"' EXIT

{
  cat "$ROOT/docs/reference/autotune-fixer-prompt.md"
  echo
  echo "## The owner's open questions"
  echo
  echo "Decide none of these. When a fix needs one of them, skip that fix and say so in your summary."
  echo
  sed -n '/^| OQ-/p' "$ROOT/docs/owner-questions.md" 2>/dev/null || true
  echo
  echo "## The eval report for $LABEL"
  echo
  cat "$DOC"
} > "$PROMPT"

# The provider keys do not reach the fixer, for two reasons.
#
# The loop exports every name in .env, and Claude Code reads
# ANTHROPIC_API_KEY in preference to a claude.ai login. The fixer would
# then bill per token to that key instead of the owner's monthly plan,
# with no cap, unattended, all night (D-161).
#
# OPENAI_API_KEY goes for a second reason. The loop owns every paid
# measurement and counts it against --budget. A fixer that ran its own
# gate to check its work would spend outside that accounting.
#
# The offline tests need neither key. They run against the fake provider.
# shellcheck disable=SC2086
exec env -u ANTHROPIC_API_KEY -u OPENAI_API_KEY $AUTOTUNE_FIXER_CMD < "$PROMPT"
