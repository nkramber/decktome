#!/usr/bin/env bash
# Print the state of the Gitar review: "on" or "paused" (D-802).
#
# The file .github/gitar-review holds one word. "paused" means that no
# session and no script waits for gitar-bot. "on" means that the author
# answers each Gitar finding before the cross-provider review (D-803).
# An absent file or another word is a fault, and the exit code is 2.
set -uo pipefail
cd "$(dirname "$0")/.." || exit 2

file=".github/gitar-review"
if [ ! -f "$file" ]; then
  echo "gitar-state: $file does not exist" >&2
  exit 2
fi
state="$(tr -d '[:space:]' < "$file")"
case "$state" in
  on|paused) echo "$state" ;;
  *) echo "gitar-state: $file reads '$state', and it must read on or paused" >&2; exit 2 ;;
esac
