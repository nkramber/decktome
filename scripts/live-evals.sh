#!/usr/bin/env bash
# The live evals (D-1132 to D-1158). `./start-live-evals` runs this file.
#
# It prints a summary of every deck, every thumbs down, and every "Leave
# feedback" note that no live eval read. Then it polls every five
# minutes. The owner gets a Pushover notice for each new deck, each
# revision, and each thumbs down (D-1143, D-1149). A note sent its own
# notice when the reader sent it (D-1156). For each item, the script
# starts one headless Claude Code session in its own clone. That session reads the deck against the
# prompt and the answers of the reader, and the verdict first. It chooses the
# most important new fault, and the owner gets a notice. It replays the
# chat of the reader on the base code and on its fix, and it tries the
# fix at most three times (D-1144). Then it takes one pull request
# through Gitar and the Codex review. When GitHub shows the pull request
# ready, this script sends the owner a notice. The owner merges. No
# session merges, deploys, or pushes main.
#
# A session fixes a product fault alone: the collection, the deck, the
# build, the questions, the chat, and the screens of these (D-1157). It
# answers `out-of-scope` for anything else, such as a new language,
# access, or a change to a user. The owner gets a notice with the
# reason. A pull request that changes a protected path, such as auth,
# the rules, the deploy, or the live evals, waits for the owner as
# blocked (D-1158).
#
# A launchd agent can run one pass every five minutes from a checkout of
# origin/main: `make live-evals-install` (D-1155).
#
#   ./start-live-evals                    poll until Ctrl-C
#   ./start-live-evals --once             one pass, then stop
#   ./start-live-evals --dry              the summary alone: no mark, no session
#   ./start-live-evals --evaluate-backlog evaluate the decks of the first run
#   ./start-live-evals --retry KEY        put one blocked or failed item back in the queue
#
# CAUTION: each session spends the Claude plan of the owner, the Codex
# plan, and at most LIVE_EVALS_BUDGET_USD (3.00) of provider money on
# paid targets, the replays included (D-1134). The live-test lanes stay
# closed to a session, because a deck that a session builds would start
# another session.
#
# Each session runs the pinned Claude Code binary under the Seatbelt
# profile scripts/live-evals/sandbox.sb (D-1136, D-1141, D-1145). It
# reads and writes its run folder, and it can not read the gcloud
# config, the .env of the owner, the state of this script, or any other
# file of the owner. It gets its own Claude token, Codex login, and
# GitHub token, and a .env with the provider keys alone (D-1142).
#
# The first run marks every deck that waits as read, after the summary,
# and starts no session for it (D-1133). The first run with the notes of
# D-1156 marks the notes that wait in the same way. --evaluate-backlog
# changes both.
#
# Each new session branches from the newest open live-eval pull request,
# or from main when none is open (D-1138). At most three live-eval pull
# requests stay open. When a parent merges, or a parent changes under its
# child, a restack session rebases the child and takes it through the
# reviews again.
#
# Settings, each with its default:
#   LIVE_EVALS_PROJECT=decktome-prod  LIVE_EVALS_INTERVAL=300 (seconds)
#   LIVE_EVALS_HOME=<repo>/../decktome-live-evals  (runs, state, caches)
#   LIVE_EVALS_SECRETS=~/.config/decktome-live-evals  (claude-token, gh-token)
#   LIVE_EVALS_ACCOUNT=<active gcloud account>     (reads Firestore and secrets)
#   LIVE_EVALS_CLAUDE_VERSION=2.1.288  (the pinned Claude Code, D-1145)
#   LIVE_EVALS_TIMEOUT=14400 (seconds a session may run)
#   LIVE_EVALS_BUDGET_USD=3.00  LIVE_EVALS_MAX_OPEN=3  LIVE_EVALS_MIN_FREE_GB=20
#   LIVE_EVALS_MODEL=<claude default>
#
# The one-time setup (D-1141, D-1164). Mode 600 for each file, and the
# script refuses a login file that another account can read:
#   umask 077
#   claude setup-token, and the token alone > $LIVE_EVALS_SECRETS/claude-token
#   gh auth token             > $LIVE_EVALS_SECRETS/gh-token
#   cp ~/.codex/auth.json $LIVE_EVALS_HOME/codex-home/auth.json
# The gh token and the Codex login are copies of the logins of the owner
# (D-1164). The gh token reaches each repository of the owner.
#
# Each program that a tick runs has a pinned version in PINS below
# (D-1165). A new version can lose the macOS access to the volume of
# the repository. So the tick stops, and the owner gets one notice.
# Move the pin in a pull request after the check of the new version.
set -uo pipefail

ROOT=$(git rev-parse --show-toplevel 2>/dev/null) || { echo "live-evals: run it inside the decktome repository"; exit 1; }
cd "$ROOT" || exit 1

PROJECT=${LIVE_EVALS_PROJECT:-decktome-prod}
INTERVAL=${LIVE_EVALS_INTERVAL:-300}
HOME_DIR=${LIVE_EVALS_HOME:-$(dirname "$ROOT")/decktome-live-evals}
SECRETS=${LIVE_EVALS_SECRETS:-$HOME/.config/decktome-live-evals}
CLAUDE_VERSION=${LIVE_EVALS_CLAUDE_VERSION:-2.1.288}
TIMEOUT=${LIVE_EVALS_TIMEOUT:-14400}
BUDGET=${LIVE_EVALS_BUDGET_USD:-3.00}
MAX_OPEN=${LIVE_EVALS_MAX_OPEN:-3}
MIN_FREE_GB=${LIVE_EVALS_MIN_FREE_GB:-20}
MAX_ATTEMPTS=2
MAX_CONTINUES=3
READY_WAIT=7200
LABEL=live-eval
PROMPTS=$ROOT/scripts/live-evals
PROFILE=$PROMPTS/sandbox.sb
CARDS=$ROOT/.local/gcs/mtg-local-cards

once=0
dry=0
backlog=0
retry=""
want_retry=0
for arg in "$@"; do
  if [ "$want_retry" = 1 ]; then retry=$arg; want_retry=0; continue; fi
  case "$arg" in
    --once) once=1 ;;
    --dry) dry=1 ;;
    --evaluate-backlog) backlog=1 ;;
    --retry) want_retry=1 ;;
    -h|--help) sed -n '2,80p' "$0"; exit 0 ;;
    *) echo "live-evals: unknown flag $arg"; exit 2 ;;
  esac
done
if [ "$want_retry" = 1 ] || { [ -n "$retry" ] && [ "$dry" = 1 ]; }; then
  echo "live-evals: --retry needs the key of one item, and it does not run with --dry"
  exit 2
fi

say() { printf '%s %s\n' "$(date '+%H:%M:%S')" "$*"; }
die() { say "STOP: $*"; exit 1; }

# --- preflight ---------------------------------------------------------

# The profile reads real paths, so each folder is resolved once here.
mkdir -p "$HOME_DIR" || die "can not make $HOME_DIR"
HOME_DIR=$(cd "$HOME_DIR" && pwd -P)
STATE=$HOME_DIR/state
BIN=$HOME_DIR/bin
CACHE=$HOME_DIR/cache
CODEX_DIR=$HOME_DIR/codex-home
RUNS=$STATE/runs
TOOL=$BIN/live-evals
CLAUDE_BIN=$BIN/claude-$CLAUDE_VERSION
mkdir -p "$STATE/decks" "$RUNS" "$BIN" "$CACHE" "$CODEX_DIR" || die "can not make the folders under $HOME_DIR"

# A session writes its run folder, and it can put a link at any path in
# it. So the script never reads or writes there by path. The helper opens
# each part with no link (D-1141). Prompts, logs, and context stay in
# $RUNS, where no session can write.
runfs() { python3 "$ROOT/docs/tools/live_evals_runfs.py" "$@"; }
rget() { runfs get "$HOME_DIR" "$1" "$2" "${3:-1048576}" 2>/dev/null; } # deck rel [max]

# One copy at a time. mkdir is atomic, and a dead holder frees the lock.
LOCK=$STATE/lock
if ! mkdir "$LOCK" 2>/dev/null; then
  holder=$(cat "$LOCK/pid" 2>/dev/null || true)
  if [ -n "$holder" ] && kill -0 "$holder" 2>/dev/null; then
    die "another start-live-evals runs as PID $holder"
  fi
  rm -rf "$LOCK"
  mkdir "$LOCK" || die "can not take the lock $LOCK"
fi
echo $$ > "$LOCK/pid"
child=""
cleanup() {
  if [ -n "$child" ]; then kill "$child" 2>/dev/null; fi
  rm -rf "$LOCK"
}
trap cleanup EXIT
trap 'exit 130' INT TERM

# The tuning loop resets the checkout of the owner, and this script
# clones with that checkout as its reference.
if pgrep -f "scripts/(autotune|feedback-loop)\.sh" >/dev/null; then
  die "a tuning or feedback loop runs. Stop it first."
fi

for cmd in git go gh jq codex gcloud python3 sandbox-exec shasum pnpm curl; do
  command -v "$cmd" >/dev/null || die "$cmd is not on PATH"
done

# The pinned version of each program that a tick runs (D-1165). Each
# line holds the name, the version command, and the first line of its
# output. The macOS build pins the programs of /bin and /usr/bin, and
# the Command Line Tools pin git and python3 apart. Claude Code has its
# own pin (D-1145), and .nvmrc pins Node 22.
PINS=(
  "macOS|sw_vers -buildVersion|25F84"
  "jq|jq --version|jq-1.7.1-apple"
  "git|git --version|git version 2.50.1 (Apple Git-155)"
  "python3|python3 --version|Python 3.9.6"
  "go|go version|go version go1.27.1 darwin/arm64"
  "gh|gh --version|gh version 2.102.0 (2026-09-30)"
  "codex|codex --version|codex-cli 0.39.0"
  "gcloud|gcloud version|Google Cloud SDK 533.0.0"
  "node|node --version|v20.17.0"
  "pnpm|pnpm --version|9.2.0"
)
pin_misses=""
for pin in "${PINS[@]}"; do
  IFS='|' read -r name cmd want <<<"$pin"
  got=$($cmd 2>/dev/null | head -1)
  [ "$got" = "$want" ] || pin_misses="${pin_misses}$name reads \"$got\", and the pin is \"$want\". "
done

gh auth status >/dev/null 2>&1 || die "gh is not signed in"

free_gb=$(df -g "$HOME_DIR" | awk 'NR==2 {print $4}')
if [ "${free_gb:-0}" -lt "$MIN_FREE_GB" ]; then
  die "only ${free_gb} GB free under $HOME_DIR, and the floor is $MIN_FREE_GB GB"
fi

ACCOUNT=${LIVE_EVALS_ACCOUNT:-$(gcloud config get-value account 2>/dev/null)}
[ -n "$ACCOUNT" ] || die "no gcloud account. Set LIVE_EVALS_ACCOUNT."
export CLOUDSDK_CORE_ACCOUNT=$ACCOUNT
# The tool labels the decks of the owner and of the test account. Neither
# email leaves this process (D-1135).
export LIVE_EVALS_OWNER_EMAIL=${LIVE_EVALS_OWNER_EMAIL:-$ACCOUNT}
if [ -z "${LIVE_EVALS_TEST_EMAIL:-}" ] && [ -f "$ROOT/.env" ]; then
  LIVE_EVALS_TEST_EMAIL=$(sed -n 's/^API_BUILD_EMAIL=//p' "$ROOT/.env" | tr -d "\"'" | head -1)
fi
export LIVE_EVALS_TEST_EMAIL=${LIVE_EVALS_TEST_EMAIL:-}

# The Pushover pair lives in Secret Manager (D-893). It stays in this
# shell, and only the notify call sees it.
PO_TOKEN=$(gcloud secrets versions access latest --secret=pushover-app-token --project "$PROJECT" 2>/dev/null) || PO_TOKEN=""
PO_USER=$(gcloud secrets versions access latest --secret=pushover-user-key --project "$PROJECT" 2>/dev/null) || PO_USER=""
if [ -z "$PO_TOKEN" ] || [ -z "$PO_USER" ]; then
  die "can not read the Pushover secrets of $PROJECT as $ACCOUNT"
fi

# A missed pin stops the tick before the build. The tool can not send
# the notice, because go can be the missed pin. So curl sends it, and
# it reads the secrets from stdin. One notice goes for each new miss.
PIN_NOTICE=$STATE/pin-notice
if [ -n "$pin_misses" ]; then
  if [ "$dry" = 0 ] && [ "$(cat "$PIN_NOTICE" 2>/dev/null)" != "$pin_misses" ]; then
    msg=$(printf '%s' "$pin_misses" | sed 's/\\/\\\\/g; s/"/\\"/g')
    if printf 'form-string = "token=%s"\nform-string = "user=%s"\nform-string = "title=%s"\nform-string = "message=%s"\n' \
      "$PO_TOKEN" "$PO_USER" "decktome live evals stopped: a version pin" "$msg" \
      | curl -sS -o /dev/null --fail -K - https://api.pushover.net/1/messages.json; then
      printf '%s' "$pin_misses" > "$PIN_NOTICE"
    else
      say "the notice of the pins did not send"
    fi
  fi
  die "a version pin missed (D-1165): $pin_misses"
fi
rm -f "$PIN_NOTICE"

say "building the live-evals tool"
(cd "$ROOT/go" && go build -o "$TOOL" ./cmd/live-evals) || die "the tool does not build"

# The script checks each binary before each call. The profile already
# keeps a session out of $BIN, and the sum catches any other change.
sum() { shasum -a 256 < "$1" | cut -d' ' -f1; }
TOOL_SUM=$(sum "$TOOL")
check_sum() { # file sum
  [ "$(sum "$1")" = "$2" ] || die "$1 changed after the start. Stop, and read who wrote it."
}

tool() { check_sum "$TOOL" "$TOOL_SUM"; PROJECT_ID=$PROJECT "$TOOL" "$@"; }

notify() {
  if [ "$dry" = 1 ]; then say "notice (dry, not sent): $1"; return 0; fi
  check_sum "$TOOL" "$TOOL_SUM"
  PUSHOVER_APP_TOKEN=$PO_TOKEN PUSHOVER_USER_KEY=$PO_USER "$TOOL" notify -title "$1" -message "$2" \
    || say "the notice did not send: $1"
}

# A login file that another account can read stops the script (D-1164).
# A mode bit of the group or of others can give the read, and so can an
# entry of an access control list. `ls -lde` prints one line for each
# entry after the line of the file.
private_file() { # file
  local mode
  mode=$(stat -f %Lp "$1") || die "can not read the mode of $1"
  [ $(( 8#$mode & 8#077 )) = 0 ] || die "$1 has mode $mode, and another account can read it. Run chmod 600 on it (D-1164)."
  # shellcheck disable=SC2012 # find prints no access control list, and the path is fixed.
  [ "$(ls -lde "$1" | wc -l)" -eq 1 ] || die "$1 has an access control list. Run chmod -N on it (D-1164)."
}

# The flags of sandbox-exec for one run folder. The session and the probe
# of the sandbox use the same flags (D-1168). LOGS is the logs folder of
# the deck, and LOGS_RX is the same path as a regular expression. So the
# profile allows a list of the folder and a read of its session-<n>.log
# files alone (D-1173).
sandbox_flags() { # run ledger logs
  SANDBOX_FLAGS=(-f "$PROFILE" -D HOME="$HOME" -D RUN="$1" -D LEDGER="$2" -D LOGS="$3" -D LOGS_RX="$(rx_quote "$3")"
    -D CACHE="$CACHE" -D CODEX="$CODEX_DIR" -D BIN="$BIN" -D CARDS="$CARDS" -D NODE22="$NODE22" -D NODE20="$NODE20")
}

rx_quote() { # text -> the text with each regular-expression character escaped
  printf '%s' "$1" | sed 's/[][\\.*^$+?(){}|]/\\&/g'
}

# The ledger of a deck holds the measured spend of each paid run (D-1171).
# It sits outside the run folder, and its append-only flag stays. The
# profile allows a session to append to it, and nothing more. So a
# session can not delete, truncate, or rewrite the spend that each start
# of a paid command reads.
make_ledger() { # file -> 0 when the file exists with the append-only flag
  mkdir -p "$(dirname "$1")" || return 1
  [ -f "$1" ] || : > "$1" || return 1
  chflags uappnd "$1" || return 1
  case ",$(stat -f %Sf "$1")," in *,uappnd,*) return 0 ;; esac
  return 1
}

drop_ledger() { # file -> remove a probe ledger
  chflags nouappnd "$1" 2>/dev/null
  rm -rf "$(dirname "$1")"
}

if [ "$dry" = 0 ]; then
  # The pinned Claude Code (D-1145). A copy outside the folder of the
  # updater keeps the version, and the profile keeps a session out of it.
  if [ ! -x "$CLAUDE_BIN" ]; then
    src=$HOME/.local/share/claude/versions/$CLAUDE_VERSION
    [ -f "$src" ] || die "Claude Code $CLAUDE_VERSION is not installed at $src. Set LIVE_EVALS_CLAUDE_VERSION."
    if ! cp "$src" "$CLAUDE_BIN" || ! chmod 755 "$CLAUDE_BIN"; then die "can not copy $src to $CLAUDE_BIN"; fi
  fi
  "$CLAUDE_BIN" --version 2>/dev/null | grep -q "^$CLAUDE_VERSION " || die "$CLAUDE_BIN is not Claude Code $CLAUDE_VERSION"
  CLAUDE_SUM=$(sum "$CLAUDE_BIN")

  # The session gets its own Claude token. The gh token and the Codex
  # login are copies of the logins of the owner (D-1141, D-1164).
  [ -s "$SECRETS/claude-token" ] || die "no $SECRETS/claude-token. Run claude setup-token, and save the token there."
  [ -s "$SECRETS/gh-token" ] || die "no $SECRETS/gh-token. Save the output of gh auth token there (D-1164)."
  private_file "$SECRETS/claude-token"
  private_file "$SECRETS/gh-token"
  CLAUDE_TOKEN=$(cat "$SECRETS/claude-token")
  GH_SESSION_TOKEN=$(cat "$SECRETS/gh-token")
  # A session can write the Codex home, so this script never runs Codex
  # with it. It reads that the login file exists, and nothing more.
  if [ ! -f "$CODEX_DIR/auth.json" ] || [ -L "$CODEX_DIR/auth.json" ]; then
    die "the session Codex has no login. Copy ~/.codex/auth.json to $CODEX_DIR (D-833, D-1164)"
  fi
  private_file "$CODEX_DIR/auth.json"
  [ -d "$CARDS/scryfall" ] || die "no card store at $CARDS. A replay needs it."
  # The checkout of the launchd agent links .local to the checkout of the
  # owner (D-1155), and the profile reads real paths.
  CARDS=$(cd "$CARDS" && pwd -P)
  REPO_URL=$(gh repo view --json url --jq .url).git

  # The tools of a session: Node 22 for the web tests, and the pnpm of
  # the owner (CLAUDE.md). The profile lets the session read both.
  NODE22=$HOME/.nvm/versions/node/v$(tr -d 'v \n' < "$ROOT/.nvmrc")
  [ -x "$NODE22/bin/node" ] || die "no Node $(cat "$ROOT/.nvmrc") at $NODE22"
  pnpm_bin=$(command -v pnpm) || die "pnpm is not on PATH"
  NODE20=$(cd "$(dirname "$pnpm_bin")/.." && pwd -P)
  NODE22=$(cd "$NODE22" && pwd -P)
  # PATH keeps no folder of the home or of a volume, because the profile
  # hides them. The session prefix and Node come first.
  SESSION_PATH=$CACHE/npm-global/bin:$NODE22/bin:$NODE20/bin
  IFS=: read -r -a path_parts <<<"$PATH"
  for p in "${path_parts[@]}"; do
    case "$p" in "$HOME"/*|/Volumes/*|"") ;; *) SESSION_PATH=$SESSION_PATH:$p ;; esac
  done

  # Claude Code makes the folder of its Bash tool under
  # $CLAUDE_CODE_TMPDIR/claude-<uid>, or under /tmp, and it ignores
  # TMPDIR. The profile refuses /tmp, so a session gets $run/tmp there
  # (D-1168). Each tick makes that folder under the profile before any
  # session starts. A refusal stops the tick, so no session starts and
  # no item gets its mark. One notice goes for each new refusal.
  probe=$HOME_DIR/.probe
  SANDBOX_NOTICE=$STATE/sandbox-notice
  if ! rm -rf "$probe" || ! mkdir -p "$probe/tmp"; then die "can not make the probe folder $probe"; fi
  sandbox_flags "$probe" "$probe/no-ledger" "$probe/no-logs"
  if ! env -i PATH=/usr/bin:/bin sandbox-exec "${SANDBOX_FLAGS[@]}" \
    /bin/mkdir -p "$probe/tmp/claude-$(id -u)/probe" 2>/dev/null; then
    if [ ! -f "$SANDBOX_NOTICE" ]; then
      notify "decktome live evals stopped: the sandbox" "The profile refuses $probe/tmp/claude-$(id -u). No session starts until the probe passes (D-1168)."
      touch "$SANDBOX_NOTICE"
    fi
    die "the profile refuses the temporary folder of a session, $probe/tmp/claude-$(id -u) (D-1168)"
  fi
  # The ledger takes an append under the profile, and no truncate and no
  # change of its flag (D-1171). A refusal stops the tick the same way.
  probe_ledger=$HOME_DIR/.probe-ledger/spend.jsonl
  drop_ledger "$probe_ledger"
  make_ledger "$probe_ledger" || die "can not make the probe ledger $probe_ledger"
  sandbox_flags "$probe" "$probe_ledger" "$probe/no-logs"
  ledger_ok=0
  # shellcheck disable=SC2016 # $1 expands in the shell under the profile
  if env -i PATH=/usr/bin:/bin sandbox-exec "${SANDBOX_FLAGS[@]}" /bin/sh -c 'echo probe >> "$1"' sh "$probe_ledger" 2>/dev/null \
    && ! env -i PATH=/usr/bin:/bin sandbox-exec "${SANDBOX_FLAGS[@]}" /bin/sh -c ': > "$1"' sh "$probe_ledger" 2>/dev/null \
    && ! env -i PATH=/usr/bin:/bin sandbox-exec "${SANDBOX_FLAGS[@]}" /usr/bin/chflags nouappnd "$probe_ledger" 2>/dev/null \
    && [ "$(cat "$probe_ledger")" = probe ]; then
    ledger_ok=1
  fi
  drop_ledger "$probe_ledger"
  if [ "$ledger_ok" != 1 ]; then
    if [ ! -f "$SANDBOX_NOTICE" ]; then
      notify "decktome live evals stopped: the ledger" "The profile does not keep the spend ledger append-only. No session starts until the probe passes (D-1171)."
      touch "$SANDBOX_NOTICE"
    fi
    die "the profile does not keep the spend ledger append-only (D-1171)"
  fi
  # A paid command reads the session logs of its deck, and no other file
  # of the logs folder (D-1173). A refusal stops the tick the same way.
  probe_logs=$HOME_DIR/.probe-logs
  if ! rm -rf "$probe_logs" || ! mkdir -p "$probe_logs"; then die "can not make the probe logs $probe_logs"; fi
  echo probe > "$probe_logs/session-1.log"
  echo probe > "$probe_logs/context.json"
  sandbox_flags "$probe" "$probe_ledger" "$probe_logs"
  if env -i PATH=/usr/bin:/bin sandbox-exec "${SANDBOX_FLAGS[@]}" /bin/ls "$probe_logs" >/dev/null 2>&1 \
    && env -i PATH=/usr/bin:/bin sandbox-exec "${SANDBOX_FLAGS[@]}" /bin/cat "$probe_logs/session-1.log" >/dev/null 2>&1 \
    && ! env -i PATH=/usr/bin:/bin sandbox-exec "${SANDBOX_FLAGS[@]}" /bin/cat "$probe_logs/context.json" >/dev/null 2>&1; then
    logs_ok=1
  else
    logs_ok=0
  fi
  rm -rf "$probe_logs"
  if [ "$logs_ok" != 1 ]; then
    if [ ! -f "$SANDBOX_NOTICE" ]; then
      notify "decktome live evals stopped: the logs" "The profile does not limit a session to the read of its own logs. No session starts until the probe passes (D-1173)."
      touch "$SANDBOX_NOTICE"
    fi
    die "the profile does not limit a session to the read of its own logs (D-1173)"
  fi
  rm -rf "$probe" "$SANDBOX_NOTICE"
fi

if [ "$dry" = 0 ] && ! gh label list --limit 200 --json name --jq '.[].name' | grep -qx "$LABEL"; then
  gh label create "$LABEL" --description "A pull request of the live evals (D-1132)" --color 5319e7 >/dev/null \
    || die "can not make the $LABEL label"
fi

# --- deck state ---------------------------------------------------------
# One JSON file per deck in $STATE/decks holds its state. ledger.jsonl
# keeps each change. Both stay outside the repository and outside the
# reach of a session, because they name users (D-639).

dstate() { echo "$STATE/decks/$1.json"; }

dget() { # deck key -> value, or empty
  local f
  f=$(dstate "$1")
  if [ -f "$f" ]; then jq -r --arg k "$2" '.[$k] // empty' "$f"; fi
}

dnum() { # deck key -> the number, or 0 when the key is absent
  local v
  v=$(dget "$1" "$2")
  echo "${v:-0}"
}

dset() { # deck key value [key value ...]
  local deck=$1 f tmp
  shift
  f=$(dstate "$deck")
  [ -f "$f" ] || echo '{}' > "$f"
  tmp=$f.tmp
  cp "$f" "$tmp"
  while [ $# -ge 2 ]; do
    jq --arg k "$1" --arg v "$2" '.[$k] = $v' "$tmp" > "$tmp.2" && mv "$tmp.2" "$tmp"
    shift 2
  done
  mv "$tmp" "$f"
  jq -c --arg deck "$deck" --arg at "$(date -u +%Y-%m-%dT%H:%M:%SZ)" '. + {deck: $deck, at: $at}' "$f" >> "$STATE/ledger.jsonl"
}

mark() { # uid deck
  if [ "$dry" = 1 ]; then return 0; fi
  tool mark -uid "$1" -deck "$2" || say "the mark of $2 did not write"
}

# --- the session --------------------------------------------------------

render() { # template deck run -> prompt on stdout
  python3 - "$1" "$2" "$3" "$BUDGET" "$RUNS/$2/context.json" <<'PY'
import json, sys, pathlib
template, deck, run, budget, context = sys.argv[1:6]
ctx = json.loads(pathlib.Path(context).read_text())
text = pathlib.Path(template).read_text()
values = dict(ctx, deck=deck, run=run, bundle=str(pathlib.Path(run, "bundle")),
              replay=str(pathlib.Path(run, "replay")), budget=budget)
for key, value in values.items():
    text = text.replace("{{" + key + "}}", str(value))
print(text)
PY
}

run_claude() { # deck run prompt-file log-file -> exit code
  local deck=$1 run=$2 prompt=$3 log=$4 started rc model_args=() ledger=$RUNS/$1/ledger/spend.jsonl
  runfs mkdir "$HOME_DIR" "$deck" tmp claude-config gh-config xdg-config replay || return 1
  if [ -n "${LIVE_EVALS_MODEL:-}" ]; then model_args=(--model "$LIVE_EVALS_MODEL"); fi
  check_sum "$CLAUDE_BIN" "$CLAUDE_SUM"
  make_ledger "$ledger" || { say "can not make the ledger $ledger"; return 1; }
  sandbox_flags "$run" "$ledger" "$RUNS/$deck"
  say "session starts in $run/repo, log $log"
  # env -i passes only the names below. No cloud credential, no login of
  # the owner, and no provider key reach the process. A provider key in
  # the environment would also replace the Claude plan (D-1142).
  (cd "$run/repo" && env -i \
    HOME="$HOME" USER="$(id -un)" LOGNAME="$(id -un)" LANG="${LANG:-en_US.UTF-8}" TERM=dumb SHELL=/bin/bash \
    PATH="$SESSION_PATH" TMPDIR="$run/tmp" CLAUDE_CODE_TMPDIR="$run/tmp" \
    CLAUDE_CONFIG_DIR="$run/claude-config" CLAUDE_CODE_OAUTH_TOKEN="$CLAUDE_TOKEN" DISABLE_AUTOUPDATER=1 \
    GH_TOKEN="$GH_SESSION_TOKEN" GH_CONFIG_DIR="$run/gh-config" XDG_CONFIG_HOME="$run/xdg-config" \
    GOCACHE="$CACHE/go-build" GOMODCACHE="$CACHE/go-mod" GOPATH="$CACHE/gopath" \
    npm_config_prefix="$CACHE/npm-global" npm_config_cache="$CACHE/npm" \
    npm_config_store_dir="$CACHE/pnpm-store" PNPM_HOME="$CACHE/pnpm-home" CODEX_HOME="$CODEX_DIR" \
    LIVE_EVAL_BUNDLE="$run/bundle" LIVE_EVAL_BUDGET_USD="$BUDGET" LIVE_EVAL_LEDGER="$ledger" LIVE_EVAL_LOGS="$RUNS/$deck" \
    sandbox-exec "${SANDBOX_FLAGS[@]}" \
    "$CLAUDE_BIN" -p "$(cat "$prompt")" --permission-mode bypassPermissions \
      --add-dir "$run/bundle" --add-dir "$run/replay" \
      ${model_args[@]+"${model_args[@]}"} --output-format stream-json --verbose > "$log" 2>&1) &
  child=$!
  started=$(date +%s)
  while kill -0 "$child" 2>/dev/null; do
    fix_notice "$deck"
    if [ $(( $(date +%s) - started )) -ge "$TIMEOUT" ]; then
      say "the session passed $TIMEOUT seconds, and it stops"
      pkill -TERM -P "$child" 2>/dev/null
      kill "$child" 2>/dev/null
      break
    fi
    sleep 10
  done
  wait "$child"
  rc=$?
  child=""
  fix_notice "$deck"
  return "$rc"
}

fix_notice() { # deck -> notify once when the session chose a fix (D-1143)
  local fix
  [ -z "$(dget "$1" fix_notice)" ] || return 0
  fix=$(rget "$1" bundle/fix.json 65536) || return 0
  dset "$1" fix_notice sent
  notify "decktome live eval: deck $1 needs a fix" \
    "$(jq -r '"Finding: \(.finding // "?")\nBar: \(.bar // "?")"' <<<"$fix" 2>/dev/null || echo "fix.json does not read")"
}

spent() { # deck -> the measured spend of the item, for a notice (D-1169, D-1170)
  # The marked lines of the session logs and the append-only ledger are
  # the sources (D-1171). The session can not edit a log or lower the
  # ledger, and spend.jsonl of the bundle is a cross-check. The tool
  # refuses a line that no paid command wrote, and an unmeasured run
  # counts as the full budget.
  local r ledger=$RUNS/$1/ledger/spend.jsonl
  [ -f "$ledger" ] || ledger=""
  r=$({ rget "$1" bundle/spend.jsonl || true; } | tool spend -logs "$RUNS/$1" -ledger "$ledger" -budget "$BUDGET" 2>/dev/null) || r=""
  if [ -z "$r" ]; then
    echo "unread, counted as \$$BUDGET of \$$BUDGET"
    return 0
  fi
  jq -r --arg b "$BUDGET" '"$\(.charged * 10000 | round / 10000) of $\($b): \(.text)"' <<<"$r" 2>/dev/null \
    || echo "unread, counted as \$$BUDGET of \$$BUDGET"
}

result() { # deck key -> value of result.json
  rget "$1" bundle/result.json 262144 \
    | jq -r --arg k "$2" '.[$k] // empty | if type == "array" then join("; ") else tostring end' 2>/dev/null
}

summary_text() { # deck -> the four sections of D-836
  printf 'What: %s\nHow: %s\nReplay: %s\nCI: %s\nCodex review: %s\nSpent: %s' \
    "$(result "$1" what)" "$(result "$1" how)" "$(result "$1" replay)" "$(result "$1" ci)" \
    "$(result "$1" codex)" "$(spent "$1")"
}

open_prs() { # -> number<TAB>branch, oldest first
  gh pr list --label "$LABEL" --state open --json number,headRefName,createdAt \
    --jq 'sort_by(.createdAt) | .[] | "\(.number)\t\(.headRefName)"'
}

clone() { # run branch base -> a clone of the base on a new branch
  local run=$1 branch=$2 base=$3 repo=$1/repo
  # --dissociate copies the objects, so the clone never writes or needs
  # the checkout of the owner after this step.
  git clone --quiet --reference "$ROOT" --dissociate --branch "$base" "$REPO_URL" "$repo" || return 1
  git -C "$repo" switch --quiet -c "$branch" || return 1
  git -C "$repo" config credential.helper '!gh auth git-credential'
  git -C "$repo" config user.name "$(git config user.name)"
  git -C "$repo" config user.email "$(git config user.email)"
  # The session gets the provider keys alone (D-1142). The test account
  # of make api-build stays out, because the live-test lanes stay closed.
  if [ -f "$ROOT/.env" ]; then
    (umask 077 && grep -E '^(OPENAI_API_KEY|ANTHROPIC_API_KEY|LLM_[A-Z_]+)=' "$ROOT/.env" > "$repo/.env")
  fi
  # The replay reads the card store through this link, and the profile
  # keeps the store read-only.
  mkdir -p "$repo/.local/gcs" && ln -s "$CARDS" "$repo/.local/gcs/mtg-local-cards"
}

prepare() { # uid deck kind who -> 0 when the run folder is ready
  local uid=$1 deck=$2 kind=$3 who=$4 run=$HOME_DIR/$2 logs=$RUNS/$2 build branch base parent newest nonce f
  branch=live-eval/$(echo "$deck" | tr '[:upper:]' '[:lower:]' | cut -c1-12)
  git fetch --quiet origin || return 1
  mkdir -p "$logs"
  if [ -z "$(dget "$deck" branch)" ]; then
    newest=$(open_prs | tail -1)
    base=main
    parent=""
    if [ -n "$newest" ]; then
      parent=${newest%%$'\t'*}
      base=${newest#*$'\t'}
    fi
    # The deck state names no branch, so no session ran in the folder.
    # A folder there is what an interrupted clone left, and it goes.
    runfs clear "$HOME_DIR" "$deck" || return 1
    # A new run folder holds no file of a session yet, so the clone may
    # use paths in it. mkdir fails when the folder exists.
    mkdir "$run" || return 1
    if ! clone "$run" "$branch" "$base"; then runfs clear "$HOME_DIR" "$deck"; return 1; fi
    dset "$deck" uid "$uid" kind "$kind" who "$who" branch "$branch" base "$base" parent_pr "$parent" status running
  fi
  # The bundle is built in $RUNS, where no session can write. A
  # continuation keeps what the earlier session wrote: its findings, its
  # spend, and its fix (D-1134). Then the helper copies the build in.
  build=$logs/bundle
  rm -rf "${build:?}"
  # The nonce marks where reader text ends. The prompt names it, and a
  # reader can not guess it (D-1139).
  nonce=$(python3 -c 'import secrets; print(secrets.token_hex(8))')
  tool bundle -uid "$uid" -deck "$deck" -out "$build" -nonce "$nonce" || return 1
  for f in findings.md spend.jsonl fix.json; do
    rget "$deck" "bundle/$f" > "$build/$f" || rm -f "${build:?}/${f:?}"
  done
  open_prs | jq -R -s -c 'split("\n") | map(select(length > 0) | split("\t") | {number: .[0], branch: .[1]})' > "$build/open-prs.json"
  cp "$STATE/findings-index.md" "$build/earlier-findings.md" 2>/dev/null || echo "No earlier finding." > "$build/earlier-findings.md"
  runfs put "$HOME_DIR" "$deck" bundle "$build" || return 1
  jq -n --arg branch "$(dget "$deck" branch)" --arg base "$(dget "$deck" base)" --arg parent_pr "$(dget "$deck" parent_pr)" \
    --arg kind "$kind" --arg who "$who" --arg label "$LABEL" --arg nonce "$nonce" \
    '{branch: $branch, base: $base, parent_pr: (if $parent_pr == "" then "none" else $parent_pr end), kind: $kind, who: $who, label: $label, nonce: $nonce}' \
    > "$logs/context.json"
}

finish() { # deck run -> act on result.json
  local deck=$1 run=$2 status pr uid n
  uid=$(dget "$deck" uid)
  status=$(result "$deck" status)
  pr=$(result "$deck" pr)
  if [ -n "$(result "$deck" findings)" ]; then
    printf -- '- deck %s: %s\n' "$deck" "$(result "$deck" findings)" >> "$STATE/findings-index.md"
  fi
  case "$status" in
    ready)
      dset "$deck" status waiting pr "$pr" waiting_since "$(date +%s)"
      check_ready "$deck"
      ;;
    out-of-scope)
      # The item asks for no product fix (D-1157). The owner decides on it.
      say "deck $deck: out of scope. $(result "$deck" reason)"
      dset "$deck" status out-of-scope
      mark "$uid" "$deck"
      runfs rm "$HOME_DIR" "$deck" repo
      notify "decktome live eval: item $deck is out of scope" "$(result "$deck" reason)
No fix and no pull request. The owner decides."
      ;;
    no-new-issues)
      say "deck $deck: no new fault. $(result "$deck" findings)"
      dset "$deck" status "done"
      mark "$uid" "$deck"
      runfs rm "$HOME_DIR" "$deck" repo
      ;;
    fix-failed)
      # Three tries of the fix did not beat the replay of the base code
      # (D-1144). The run folder stays for the owner.
      dset "$deck" status fix-failed
      mark "$uid" "$deck"
      notify "decktome live eval: the fix failed on deck $deck" "$(result "$deck" reason)
Replay: $(result "$deck" replay)
Spent: $(spent "$deck"). Read $run and $RUNS/$deck"
      ;;
    blocked)
      dset "$deck" status blocked pr "$pr"
      mark "$uid" "$deck"
      notify "decktome live eval: blocked on deck $deck" "$(result "$deck" reason)
Questions: $(result "$deck" questions)
$( [ -n "$pr" ] && echo "PR #$pr")"
      ;;
    checkpoint)
      n=$(( $(dnum "$deck" continues) + 1 ))
      dset "$deck" status checkpoint continues "$n"
      ;;
    *)
      n=$(( $(dnum "$deck" attempts) + 1 ))
      dset "$deck" status failed attempts "$n"
      if [ "$n" -ge "$MAX_ATTEMPTS" ]; then
        mark "$uid" "$deck"
        notify "decktome live eval: failed on deck $deck" "The session ended $n times with no result. Logs: $RUNS/$deck"
      fi
      ;;
  esac
}

check_ready() { # deck -> notify once when GitHub shows the pull request ready
  local deck=$1 pr out since
  pr=$(dget "$deck" pr)
  [ -n "$pr" ] || return 0
  check_sum "$TOOL" "$TOOL_SUM"
  # GitHub names the changed files, and never the session (D-1158).
  out=$("$TOOL" guard -pr "$pr" 2>&1)
  case $? in
    0) ;;
    4)
      dset "$deck" status blocked guard yes
      mark "$(dget "$deck" uid)" "$deck"
      notify "decktome live eval: PR #$pr changes a protected path" "${out#live-evals: }
The pull request waits for the owner, and no ready notice follows."
      say "PR #$pr changes a protected path, so it waits for the owner"
      return 0
      ;;
    *) say "PR #$pr: the guard did not read: $out"; return 0 ;;
  esac
  if out=$("$TOOL" ready -pr "$pr" 2>&1); then
    dset "$deck" status ready
    mark "$(dget "$deck" uid)" "$deck"
    notify "decktome: PR #$pr is ready to merge" "$(summary_text "$deck")
${out#ready }"
    say "PR #$pr is ready, and the owner has the notice"
    return 0
  fi
  say "PR #$pr: $out"
  since=$(dget "$deck" waiting_since)
  if [ -n "$since" ] && [ $(( $(date +%s) - since )) -ge "$READY_WAIT" ] && [ -z "$(dget "$deck" late_notice)" ]; then
    dset "$deck" late_notice sent
    notify "decktome: PR #$pr is not ready after 2 hours" "$out"
  fi
}

evaluate() { # uid deck kind who
  local uid=$1 deck=$2 kind=$3 who=$4 run=$HOME_DIR/$2 logs=$RUNS/$2 template n status
  if ! prepare "$uid" "$deck" "$kind" "$who"; then
    n=$(( $(dnum "$deck" attempts) + 1 ))
    dset "$deck" status failed attempts "$n"
    say "deck $deck: the clone or the bundle failed"
    return
  fi
  template=$PROMPTS/eval-prompt.md
  if [ "$(dget "$deck" status)" = checkpoint ]; then template=$PROMPTS/continue-prompt.md; fi
  render "$template" "$deck" "$run" > "$logs/prompt.md"
  runfs rm "$HOME_DIR" "$deck" bundle/result.json
  n=$(( $(dnum "$deck" sessions) + 1 ))
  dset "$deck" sessions "$n" status running
  run_claude "$deck" "$run" "$logs/prompt.md" "$logs/session-$n.log"
  status=$?
  say "deck $deck: the session ended with exit $status and result $(result "$deck" status)"
  finish "$deck" "$run"
}

restack() { # deck target -> rebase the pull request of deck onto target
  local deck=$1 target=$2 run=$HOME_DIR/$1 logs=$RUNS/$1 n
  dset "$deck" restack_target "$target"
  jq --arg t "$target" --arg pr "$(dget "$deck" pr)" '.restack_target = $t | .pr = $pr' "$logs/context.json" > "$logs/context.tmp" \
    && mv "$logs/context.tmp" "$logs/context.json"
  render "$PROMPTS/restack-prompt.md" "$deck" "$run" > "$logs/restack-prompt.md"
  runfs rm "$HOME_DIR" "$deck" bundle/result.json
  n=$(( $(dnum "$deck" sessions) + 1 ))
  dset "$deck" sessions "$n" status restacking
  run_claude "$deck" "$run" "$logs/restack-prompt.md" "$logs/session-$n.log"
  if [ "$(result "$deck" status)" = ready ]; then
    local parent=""
    if [ "$target" != main ]; then parent=$(dget "$deck" parent_pr); fi
    dset "$deck" status waiting waiting_since "$(date +%s)" late_notice "" base "$target" parent_pr "$parent"
    check_ready "$deck"
  else
    dset "$deck" status blocked
    notify "decktome live eval: the restack of PR #$(dget "$deck" pr) stopped" "$(result "$deck" reason)"
  fi
}

# --- the passes ---------------------------------------------------------

restack_pass() {
  local f deck pr parent state base phead
  git fetch --quiet origin || return
  for f in "$STATE"/decks/*.json; do
    [ -f "$f" ] || continue
    deck=$(basename "$f" .json)
    pr=$(dget "$deck" pr)
    parent=$(dget "$deck" parent_pr)
    case "$(dget "$deck" status)" in waiting|ready) ;; *) continue ;; esac
    if [ -z "$pr" ] || [ -z "$parent" ]; then continue; fi
    state=$(gh pr view "$pr" --json state --jq .state 2>/dev/null)
    [ "$state" = OPEN ] || continue
    read -r state phead < <(gh pr view "$parent" --json state,headRefName --jq '"\(.state) \(.headRefName)"' 2>/dev/null)
    case "$state" in
      MERGED) say "PR #$parent merged, so PR #$pr moves onto main"; restack "$deck" main ;;
      CLOSED)
        dset "$deck" status chain-broken
        notify "decktome live eval: PR #$parent closed unmerged" "PR #$pr branched from it and carries its commits. Decide by hand: rebase PR #$pr onto main, or close it."
        ;;
      OPEN)
        base=$(dget "$deck" branch)
        if ! git merge-base --is-ancestor "origin/$phead" "origin/$base" 2>/dev/null; then
          say "PR #$parent changed under PR #$pr, so PR #$pr moves onto it"
          restack "$deck" "$phead"
        fi
        ;;
    esac
  done
}

waiting_pass() {
  local f deck
  for f in "$STATE"/decks/*.json; do
    [ -f "$f" ] || continue
    deck=$(basename "$f" .json)
    if [ "$(dget "$deck" status)" = waiting ]; then check_ready "$deck"; fi
  done
}

cleanup_pass() { # remove the clone of a merged or closed pull request, once
  local f deck pr state
  for f in "$STATE"/decks/*.json; do
    [ -f "$f" ] || continue
    deck=$(basename "$f" .json)
    pr=$(dget "$deck" pr)
    if [ -z "$pr" ] || [ -n "$(dget "$deck" cleaned)" ]; then continue; fi
    state=$(gh pr view "$pr" --json state --jq .state 2>/dev/null)
    case "$state" in
      MERGED|CLOSED)
        runfs rm "$HOME_DIR" "$deck" repo && say "removed the clone of PR #$pr ($state)"
        dset "$deck" status "$(echo "$state" | tr '[:upper:]' '[:lower:]')" cleaned yes
        ;;
    esac
  done
}

announce_pass() { # one notice for each new deck, revision, and thumbs down (D-1143, D-1149, D-1156)
  local line deck kind who
  while IFS= read -r line; do
    [ -n "$line" ] || continue
    deck=$(jq -r .deck <<<"$line")
    [ -z "$(dget "$deck" announced)" ] || continue
    kind=$(jq -r .kind <<<"$line")
    who=$(jq -r .who <<<"$line")
    dset "$deck" announced "$(date +%s)"
    # The API sent the notice of a note when the reader sent it.
    [ "$kind" != note ] || continue
    if [ "$kind" = thumbs-down ]; then
      notify "decktome: a thumbs down of a $who account" "On a $(jq -r .target <<<"$line"). Item $deck. A live eval reads it next."
      continue
    fi
    notify "decktome: a new $kind of a $who account" "Deck $deck. A live eval reads it next."
  done < "$STATE/pending.jsonl"
}

eval_pass() {
  local line uid deck kind who status open
  while IFS= read -r line; do
    [ -n "$line" ] || continue
    uid=$(jq -r .uid <<<"$line")
    deck=$(jq -r .deck <<<"$line")
    kind=$(jq -r .kind <<<"$line")
    who=$(jq -r .who <<<"$line")
    status=$(dget "$deck" status)
    case "$status" in waiting|ready|blocked|done|restacking|fix-failed|out-of-scope) continue ;; esac
    if [ "$status" = failed ] && [ "$(dnum "$deck" attempts)" -ge "$MAX_ATTEMPTS" ]; then continue; fi
    if [ "$status" = checkpoint ] && [ "$(dnum "$deck" continues)" -gt "$MAX_CONTINUES" ]; then
      dset "$deck" status failed attempts "$MAX_ATTEMPTS"
      mark "$uid" "$deck"
      notify "decktome live eval: deck $deck passed $MAX_CONTINUES checkpoints" "Read $HOME_DIR/$deck"
      continue
    fi
    open=$(open_prs | grep -c .)
    if [ "$open" -ge "$MAX_OPEN" ] && [ "$status" != checkpoint ]; then
      say "$open live-eval pull requests are open, so deck $deck waits for a merge"
      return
    fi
    say "deck $deck: a $kind of a $who account"
    evaluate "$uid" "$deck" "$kind" "$who"
  done < "$STATE/pending.jsonl"
}

# --- the run --------------------------------------------------------------

# --retry puts one item back in the queue after a fault of the harness
# (D-1168). It clears the mark, and it keeps the old state and logs under
# a new name. The next tick starts a new session for the item.
if [ -n "$retry" ]; then
  [ -f "$(dstate "$retry")" ] || die "--retry: no state for $retry"
  rstatus=$(dget "$retry" status)
  case "$rstatus" in
    blocked|failed|fix-failed) ;;
    *) die "--retry: $retry has the status \"$rstatus\". Only blocked, failed, or fix-failed takes a retry." ;;
  esac
  [ -z "$(dget "$retry" pr)" ] || die "--retry: $retry has PR #$(dget "$retry" pr). Close it, then retry."
  rbranch=$(dget "$retry" branch)
  if [ -n "$rbranch" ] && git ls-remote --exit-code --heads origin "$rbranch" >/dev/null 2>&1; then
    die "--retry: the branch $rbranch is on origin. Delete it, then retry."
  fi
  tool unmark -uid "$(dget "$retry" uid)" -deck "$retry" || die "--retry: the unmark of $retry did not write"
  stamp=$(date +%s)
  mv "$(dstate "$retry")" "$(dstate "$retry").retried-$stamp"
  if [ -d "$RUNS/$retry" ]; then mv "$RUNS/$retry" "$RUNS/$retry.retried-$stamp"; fi
  say "--retry: $retry waits for a live eval again. The old state and logs end with .retried-$stamp."
  exit 0
fi

say "summary of every deck that no live eval read, from $PROJECT:"
tool summary || die "the summary did not read"
if [ "$dry" = 1 ]; then
  say "--dry: no mark and no session"
  exit 0
fi

# A first run marks each item that waits read, and starts no session for
# it. A flag is set only after the read and each mark pass. Else the next
# tick of the launchd agent starts a paid session for each old item
# (D-1155).
first_run() { # jq filter -> 0 when the read and each mark passed
  local list=$STATE/first-run.jsonl line deck ok=0
  tool pending > "$list" || return 1
  while IFS= read -r line; do
    deck=$(jq -r .deck <<<"$line")
    if tool mark -uid "$(jq -r .uid <<<"$line")" -deck "$deck"; then
      dset "$deck" status backlog announced backlog
    else
      ok=1
    fi
  done < <(jq -c "$1" "$list")
  return "$ok"
}

if [ ! -f "$STATE/initialized" ] && [ "$backlog" = 0 ]; then
  say "first run: every item above is marked read, and no session starts for it (D-1133)"
  first_run . || die "the first run did not read or mark each item. The next run tries again."
fi
touch "$STATE/initialized"

# The notes join the queue with D-1156. The notes that wait at that time
# get the same first run as the decks of D-1133.
if [ ! -f "$STATE/initialized-notes" ] && [ "$backlog" = 0 ]; then
  say "first run with the notes: every note that waits is marked read, and no session starts for it (D-1156)"
  first_run 'select(.kind == "note")' || die "the first run of the notes did not read or mark each note. The next run tries again."
fi
touch "$STATE/initialized-notes"

while true; do
  cleanup_pass
  restack_pass
  waiting_pass
  if tool pending > "$STATE/pending.jsonl"; then
    announce_pass
    eval_pass
  else
    say "the pending list did not read"
  fi
  [ "$once" = 1 ] && break
  say "next poll in $INTERVAL seconds"
  sleep "$INTERVAL"
done
