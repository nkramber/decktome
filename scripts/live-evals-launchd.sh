#!/usr/bin/env bash
# The launchd agent of the live evals (D-1155). It copies the night fixer
# of what-you-carry: a launchd agent runs one pass every five minutes,
# and no terminal stays open.
#
#   scripts/live-evals-launchd.sh install     write the agent, and load it
#   scripts/live-evals-launchd.sh uninstall   stop the agent, and remove it
#   scripts/live-evals-launchd.sh status      the state of the agent, and the log
#   scripts/live-evals-launchd.sh print       the agent file on stdout, and no load
#   scripts/live-evals-launchd.sh retry KEY   put one blocked or failed item back in the queue
#
# Each tick moves a separate clone, $LIVE_EVALS_HOME/main, to
# origin/main, and runs `start-live-evals --once` there. So the agent
# runs merged code alone, and it never touches the branch of the owner or
# of a session. That clone links .env and .local to the checkout of the
# owner.
#
# launchd starts no tick while the last one runs, and the lock of
# scripts/live-evals.sh refuses a second copy, such as a run by hand.
# One tick can run a session for up to LIVE_EVALS_TIMEOUT seconds.
#
# CAUTION: the agent starts a paid session for each new item, with no
# one at the terminal (D-1134). Install it only after the one-time setup
# in the header of scripts/live-evals.sh. uninstall stops a tick that
# runs, and its session with it.
#
# Settings, each with its default:
#   LIVE_EVALS_HOME=<checkout of the owner>/../decktome-live-evals
#   LIVE_EVALS_SECRETS=~/.config/decktome-live-evals
#   LIVE_EVALS_ACCOUNT=<active gcloud account>, written into the agent
#   LIVE_EVALS_INTERVAL=300 (seconds between two ticks)
set -euo pipefail

LABEL=com.decktome.live-evals
PLIST=$HOME/Library/LaunchAgents/$LABEL.plist
LOG=$HOME/Library/Logs/decktome-live-evals.log
DOMAIN=gui/$(id -u)

say() { printf 'live-evals-launchd: %s\n' "$*"; }
die() { say "STOP: $*"; exit 1; }

[ "$(uname -s)" = Darwin ] || die "launchd is on macOS alone"

# The checkout of the owner is the main working tree. A worktree of a
# session can go away, and the main one stays.
common=$(git rev-parse --path-format=absolute --git-common-dir 2>/dev/null) || die "run it inside the decktome repository"
SOURCE=$(dirname "$common")
HOME_DIR=${LIVE_EVALS_HOME:-$(dirname "$SOURCE")/decktome-live-evals}
SECRETS=${LIVE_EVALS_SECRETS:-$HOME/.config/decktome-live-evals}
INTERVAL=${LIVE_EVALS_INTERVAL:-300}

preflight() {
  # A tick that stops at setup would only fill the log, so the setup of
  # the owner comes first.
  [ -f "$SOURCE/.env" ] || die "no .env in $SOURCE"
  [ -d "$SOURCE/.local/gcs/mtg-local-cards/scryfall" ] || die "no card store under $SOURCE/.local"
  [ -s "$SECRETS/claude-token" ] || die "no $SECRETS/claude-token. Do the setup in the header of scripts/live-evals.sh"
  [ -s "$SECRETS/gh-token" ] || die "no $SECRETS/gh-token. Do the setup in the header of scripts/live-evals.sh"
  [ -f "$HOME_DIR/codex-home/auth.json" ] || die "the session Codex has no login. Do the setup in the header of scripts/live-evals.sh"
  # A login file that another account can read stops the install (D-1164).
  local f mode
  for f in "$SECRETS/claude-token" "$SECRETS/gh-token" "$HOME_DIR/codex-home/auth.json"; do
    mode=$(stat -f %Lp "$f") || die "can not read the mode of $f"
    [ $(( 8#$mode & 8#077 )) = 0 ] || die "$f has mode $mode, and another account can read it. Run chmod 600 on it."
    # shellcheck disable=SC2012 # find prints no access control list, and the path is fixed.
    [ "$(ls -lde "$f" | wc -l)" -eq 1 ] || die "$f has an access control list. Run chmod -N on it."
  done
}

# launchd gives a small PATH. The agent gets the folder of each tool that
# the script needs, as this shell finds it now. A retry uses the same
# PATH, so the version pins read the same programs (D-1165).
agent_path() {
  local path="" dir cmd
  for cmd in git go gh jq codex gcloud python3 sandbox-exec shasum pnpm; do
    dir=$(command -v "$cmd") || die "$cmd is not on PATH"
    dir=$(dirname "$dir")
    case ":$path:" in *":$dir:"*) ;; *) path=${path:+$path:}$dir ;; esac
  done
  echo "$path:/usr/bin:/bin:/usr/sbin:/sbin"
}

write_plist() { # file -> the agent file
  mkdir -p "$HOME_DIR"
  HOME_DIR=$(cd "$HOME_DIR" && pwd -P)
  case "$INTERVAL" in '' | *[!0-9]*) die "LIVE_EVALS_INTERVAL must be a number of seconds" ;; esac
  local account
  account=${LIVE_EVALS_ACCOUNT:-$(gcloud config get-value account 2>/dev/null)}
  [ -n "$account" ] || die "no gcloud account. Set LIVE_EVALS_ACCOUNT."

  local path
  path=$(agent_path) || exit 1

  # The tick runs from launchd. It reads no file of a branch before it
  # moves its own clone to origin/main. The clone is a full clone, and no
  # worktree, because the script clones each session with it as the
  # reference, and git refuses a linked worktree there. /bin/bash reads
  # each script of the volume, because macOS refuses the volume to the
  # children of a script that the kernel runs itself (D-1166).
  local tick url
  url=$(git -C "$SOURCE" remote get-url origin) || die "the checkout has no origin"
  tick=$(cat <<EOF
set -uo pipefail
src=$(printf %q "$SOURCE")
wt=$(printf %q "$HOME_DIR/main")
url=$(printf %q "$url")
echo "\$(date '+%Y-%m-%d %H:%M:%S') tick"
if [ ! -e "\$wt/.git" ]; then
  git clone --quiet --reference "\$src" --dissociate "\$url" "\$wt" || exit 1
  printf '.local\n' >> "\$wt/.git/info/exclude"
fi
git -C "\$wt" fetch --quiet origin main || { echo "the fetch of origin/main failed"; exit 1; }
git -C "\$wt" checkout --quiet --detach --force origin/main || exit 1
git -C "\$wt" clean --quiet -fd || exit 1
ln -sfn "\$src/.env" "\$wt/.env"
ln -sfn "\$src/.local" "\$wt/.local"
cd "\$wt" || exit 1
exec /bin/bash ./start-live-evals --once
EOF
)
  python3 - "$1" "$LABEL" "$tick" "$path" "$HOME_DIR" "$SECRETS" "$account" "$LOG" "$INTERVAL" <<'PY'
import plistlib, sys
plist, label, tick, path, home, secrets, account, log, interval = sys.argv[1:10]
with open(plist, "wb") as f:
    plistlib.dump({
        "Label": label,
        "ProgramArguments": ["/bin/bash", "-c", tick],
        "EnvironmentVariables": {
            "PATH": path,
            "LANG": "en_US.UTF-8",
            "LIVE_EVALS_HOME": home,
            "LIVE_EVALS_SECRETS": secrets,
            "LIVE_EVALS_ACCOUNT": account,
        },
        "StartInterval": int(interval),
        "RunAtLoad": True,
        "StandardOutPath": log,
        "StandardErrorPath": log,
    }, f)
PY
  ACCOUNT=$account
}

install() {
  mkdir -p "$HOME_DIR" "$(dirname "$PLIST")" "$(dirname "$LOG")"
  HOME_DIR=$(cd "$HOME_DIR" && pwd -P)
  preflight
  write_plist "$PLIST"
  launchctl bootout "$DOMAIN/$LABEL" 2>/dev/null || true
  launchctl bootstrap "$DOMAIN" "$PLIST" || die "launchctl bootstrap failed"
  say "loaded $LABEL: one pass every $INTERVAL seconds from $HOME_DIR/main"
  say "the gcloud account is $ACCOUNT. Check that it reads decktome-prod."
  say "log: $LOG"
}

uninstall() {
  if launchctl bootout "$DOMAIN/$LABEL" 2>/dev/null; then say "stopped $LABEL"; else say "$LABEL was not loaded"; fi
  rm -f "$PLIST"
  say "removed $PLIST. The state under $HOME_DIR stays."
}

status() {
  if launchctl print "$DOMAIN/$LABEL" >/dev/null 2>&1; then
    launchctl print "$DOMAIN/$LABEL" | grep -E '^\s*(state|last exit code|run interval) =' || true
  else
    say "$LABEL is not loaded"
  fi
  [ -f "$LOG" ] && tail -n 5 "$LOG"
  return 0
}

# retry runs `start-live-evals --retry` in the clone of the agent at
# origin/main, with the environment of a tick (D-1168). It starts no
# session. The next tick starts one for the item.
retry() { # key
  [ -n "${1:-}" ] || die "retry needs the key of one item, for example v-<verdict id>"
  local wt path account
  HOME_DIR=$(cd "$HOME_DIR" && pwd -P) || die "no $HOME_DIR"
  wt=$HOME_DIR/main
  [ -e "$wt/.git" ] || die "no clone at $wt. Install the agent first."
  if launchctl print "$DOMAIN/$LABEL" 2>/dev/null | grep -qE '^\s*state = running'; then
    die "a tick runs now. Retry after it ends."
  fi
  path=$(agent_path) || exit 1
  account=${LIVE_EVALS_ACCOUNT:-$(gcloud config get-value account 2>/dev/null)}
  [ -n "$account" ] || die "no gcloud account. Set LIVE_EVALS_ACCOUNT."
  git -C "$wt" fetch --quiet origin main || die "the fetch of origin/main failed"
  git -C "$wt" checkout --quiet --detach --force origin/main || die "the checkout of origin/main failed"
  (cd "$wt" && env -i HOME="$HOME" USER="$(id -un)" LANG=en_US.UTF-8 PATH="$path" \
    LIVE_EVALS_HOME="$HOME_DIR" LIVE_EVALS_SECRETS="$SECRETS" LIVE_EVALS_ACCOUNT="$account" \
    /bin/bash ./start-live-evals --retry "$1")
}

case "${1:-}" in
  install) install ;;
  uninstall) uninstall ;;
  status) status ;;
  print) write_plist /dev/stdout ;;
  retry) retry "${2:-}" ;;
  *) echo "usage: scripts/live-evals-launchd.sh install|uninstall|status|print|retry KEY"; exit 2 ;;
esac
