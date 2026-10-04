#!/usr/bin/env bash
# The live evals (D-1132 to D-1139). `./start-live-evals` runs this file.
#
# It prints a summary of every deck that no live eval read. Then it
# polls every five minutes. For each new deck and each revision, it
# starts one headless Claude Code session in its own worktree. That
# session reads the deck against the prompt and the answers of the
# reader, fixes the most important new fault in one pull request, and
# takes the pull request through Gitar and the Codex review. When GitHub
# shows the pull request ready, this script sends the owner a Pushover
# notice. The owner merges. No session merges, deploys, or pushes main.
#
#   ./start-live-evals                    poll until Ctrl-C
#   ./start-live-evals --once             one pass, then stop
#   ./start-live-evals --dry              the summary alone: no mark, no session
#   ./start-live-evals --evaluate-backlog evaluate the decks of the first run
#
# CAUTION: each session spends the Claude plan of the owner, the Codex
# plan, and at most LIVE_EVALS_BUDGET_USD (3.00) of provider money on
# paid targets (D-1134). The live-test lanes stay closed to a session,
# because a deck that a session builds would start another session.
#
# The first run marks every deck that waits as read, after the summary,
# and starts no session for it (D-1133). --evaluate-backlog changes that.
#
# Each new session branches from the newest open live-eval pull request,
# or from main when none is open (D-1138). At most three live-eval pull
# requests stay open. When a parent merges, or a parent changes under its
# child, a restack session rebases the child and takes it through the
# reviews again.
#
# Settings, each with its default:
#   LIVE_EVALS_PROJECT=decktome-prod  LIVE_EVALS_INTERVAL=300 (seconds)
#   LIVE_EVALS_HOME=<repo>/../decktome-live-evals  (worktrees and bundles)
#   LIVE_EVALS_ACCOUNT=<active gcloud account>     (reads Firestore and secrets)
#   LIVE_EVALS_TIMEOUT=14400 (seconds a session may run)
#   LIVE_EVALS_BUDGET_USD=3.00  LIVE_EVALS_MAX_OPEN=3  LIVE_EVALS_MIN_FREE_GB=20
#   LIVE_EVALS_MODEL=<claude default>
set -uo pipefail

ROOT=$(git rev-parse --show-toplevel 2>/dev/null) || { echo "live-evals: run it inside the decktome repository"; exit 1; }
cd "$ROOT" || exit 1

PROJECT=${LIVE_EVALS_PROJECT:-decktome-prod}
INTERVAL=${LIVE_EVALS_INTERVAL:-300}
HOME_DIR=${LIVE_EVALS_HOME:-$(dirname "$ROOT")/decktome-live-evals}
TIMEOUT=${LIVE_EVALS_TIMEOUT:-14400}
BUDGET=${LIVE_EVALS_BUDGET_USD:-3.00}
MAX_OPEN=${LIVE_EVALS_MAX_OPEN:-3}
MIN_FREE_GB=${LIVE_EVALS_MIN_FREE_GB:-20}
MAX_ATTEMPTS=2
MAX_CONTINUES=3
READY_WAIT=7200
LABEL=live-eval
STATE=$ROOT/.local/live-evals
TOOL=$STATE/bin/live-evals
PROMPTS=$ROOT/scripts/live-evals

once=0
dry=0
backlog=0
for arg in "$@"; do
  case "$arg" in
    --once) once=1 ;;
    --dry) dry=1 ;;
    --evaluate-backlog) backlog=1 ;;
    -h|--help) sed -n '2,40p' "$0"; exit 0 ;;
    *) echo "live-evals: unknown flag $arg"; exit 2 ;;
  esac
done

say() { printf '%s %s\n' "$(date '+%H:%M:%S')" "$*"; }
die() { say "STOP: $*"; exit 1; }

# --- preflight ---------------------------------------------------------

mkdir -p "$STATE/decks" "$STATE/bin" "$HOME_DIR" || die "can not make $STATE or $HOME_DIR"

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
# makes worktrees from the same repository.
if pgrep -f "scripts/(autotune|feedback-loop)\.sh" >/dev/null; then
  die "a tuning or feedback loop runs. Stop it first."
fi

for cmd in git go gh jq claude codex gcloud python3; do
  command -v "$cmd" >/dev/null || die "$cmd is not on PATH"
done
gh auth status >/dev/null 2>&1 || die "gh is not signed in"
codex login status 2>&1 | grep -q "ChatGPT" || die "codex is not signed in with ChatGPT (D-833)"

free_gb=$(df -g "$HOME_DIR" | awk 'NR==2 {print $4}')
[ "${free_gb:-0}" -ge "$MIN_FREE_GB" ] || die "only ${free_gb} GB free under $HOME_DIR, and the floor is $MIN_FREE_GB GB"

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
[ -n "$PO_TOKEN" ] && [ -n "$PO_USER" ] || die "can not read the Pushover secrets of $PROJECT as $ACCOUNT"

say "building the live-evals tool"
(cd "$ROOT/go" && go build -o "$TOOL" ./cmd/live-evals) || die "the tool does not build"

tool() { PROJECT_ID=$PROJECT "$TOOL" "$@"; }

notify() {
  if [ "$dry" = 1 ]; then say "notice (dry, not sent): $1"; return 0; fi
  PUSHOVER_APP_TOKEN=$PO_TOKEN PUSHOVER_USER_KEY=$PO_USER "$TOOL" notify -title "$1" -message "$2" \
    || say "the notice did not send: $1"
}

if [ "$dry" = 0 ] && ! gh label list --limit 200 --json name --jq '.[].name' | grep -qx "$LABEL"; then
  gh label create "$LABEL" --description "A pull request of the live evals (D-1132)" --color 5319e7 >/dev/null \
    || die "can not make the $LABEL label"
fi

# --- deck state ---------------------------------------------------------
# One JSON file per deck in $STATE/decks holds its state. ledger.jsonl
# keeps each change. Both stay in .local, out of git, because they name
# users (D-639).

dstate() { echo "$STATE/decks/$1.json"; }

dget() { # deck key -> value, or empty
  local f
  f=$(dstate "$1")
  [ -f "$f" ] && jq -r --arg k "$2" '.[$k] // empty' "$f"
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
  python3 - "$1" "$2" "$3" "$BUDGET" <<'PY'
import json, sys, pathlib
template, deck, run, budget = sys.argv[1:5]
ctx = json.loads(pathlib.Path(run, "bundle", "context.json").read_text())
text = pathlib.Path(template).read_text()
values = dict(ctx, deck=deck, run=run, bundle=str(pathlib.Path(run, "bundle")), budget=budget)
for key, value in values.items():
    text = text.replace("{{" + key + "}}", str(value))
print(text)
PY
}

run_claude() { # run prompt-file log-file -> exit code
  local run=$1 prompt=$2 log=$3 started rc model_args=()
  mkdir -p "$run/no-gcloud"
  if [ -n "${LIVE_EVALS_MODEL:-}" ]; then model_args=(--model "$LIVE_EVALS_MODEL"); fi
  say "session starts in $run/repo, log $log"
  # No cloud credentials and no Codex key reach the session (D-833,
  # D-1136). It reads the bundle and never the deployed project.
  (cd "$run/repo" && env -u CLOUDSDK_CORE_ACCOUNT -u LIVE_EVALS_OWNER_EMAIL -u LIVE_EVALS_TEST_EMAIL \
    -u CODEX_API_KEY CLOUDSDK_CONFIG="$run/no-gcloud" GOOGLE_APPLICATION_CREDENTIALS="$run/no-gcloud/none.json" \
    LIVE_EVAL_BUNDLE="$run/bundle" LIVE_EVAL_BUDGET_USD="$BUDGET" \
    claude -p "$(cat "$prompt")" --permission-mode bypassPermissions --add-dir "$run/bundle" \
    ${model_args[@]+"${model_args[@]}"} --output-format stream-json --verbose > "$log" 2>&1) &
  child=$!
  started=$(date +%s)
  while kill -0 "$child" 2>/dev/null; do
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
  return "$rc"
}

spent() { # run -> dollars the session wrote to spend.jsonl
  local f=$1/bundle/spend.jsonl
  if [ -f "$f" ]; then jq -s '[.[].usd // 0 | tonumber] | add // 0' "$f"; else echo 0; fi
}

result() { # run key -> value of result.json
  local f=$1/bundle/result.json
  [ -f "$f" ] && jq -r --arg k "$2" '.[$k] // empty | if type == "array" then join("; ") else tostring end' "$f"
}

summary_text() { # run -> the four sections of D-836
  printf 'What: %s\nHow: %s\nCI: %s\nCodex review: %s\nSpent: $%s of $%s' \
    "$(result "$1" what)" "$(result "$1" how)" "$(result "$1" ci)" "$(result "$1" codex)" "$(spent "$1")" "$BUDGET"
}

open_prs() { # -> number<TAB>branch, oldest first
  gh pr list --label "$LABEL" --state open --json number,headRefName,createdAt \
    --jq 'sort_by(.createdAt) | .[] | "\(.number)\t\(.headRefName)"'
}

prepare() { # uid deck kind who -> 0 when the run directory is ready
  local uid=$1 deck=$2 kind=$3 who=$4 run=$HOME_DIR/$2 branch base parent newest nonce
  branch=live-eval/$(echo "$deck" | tr '[:upper:]' '[:lower:]' | cut -c1-12)
  git fetch --quiet origin || return 1
  if [ ! -d "$run/repo" ]; then
    newest=$(open_prs | tail -1)
    base=main
    parent=""
    if [ -n "$newest" ]; then
      parent=${newest%%$'\t'*}
      base=${newest#*$'\t'}
    fi
    mkdir -p "$run"
    git worktree add --quiet -b "$branch" "$run/repo" "origin/$base" || return 1
    ln -s "$ROOT/.env" "$run/repo/.env" 2>/dev/null
    ln -s "$ROOT/.local" "$run/repo/.local" 2>/dev/null
    dset "$deck" uid "$uid" kind "$kind" who "$who" branch "$branch" base "$base" parent_pr "$parent" status running
  fi
  rm -rf "$run/bundle"
  # The nonce marks where reader text ends. The prompt names it, and a
  # reader can not guess it (D-1139).
  nonce=$(python3 -c 'import secrets; print(secrets.token_hex(8))')
  tool bundle -uid "$uid" -deck "$deck" -out "$run/bundle" -nonce "$nonce" || return 1
  open_prs | jq -R -s -c 'split("\n") | map(select(length > 0) | split("\t") | {number: .[0], branch: .[1]})' > "$run/bundle/open-prs.json"
  cp "$STATE/findings-index.md" "$run/bundle/earlier-findings.md" 2>/dev/null || echo "No earlier finding." > "$run/bundle/earlier-findings.md"
  jq -n --arg branch "$(dget "$deck" branch)" --arg base "$(dget "$deck" base)" --arg parent_pr "$(dget "$deck" parent_pr)" \
    --arg kind "$kind" --arg who "$who" --arg label "$LABEL" --arg nonce "$nonce" \
    '{branch: $branch, base: $base, parent_pr: (if $parent_pr == "" then "none" else $parent_pr end), kind: $kind, who: $who, label: $label, nonce: $nonce}' \
    > "$run/bundle/context.json"
}

finish() { # deck run -> act on result.json
  local deck=$1 run=$2 status pr uid
  uid=$(dget "$deck" uid)
  status=$(result "$run" status)
  pr=$(result "$run" pr)
  [ -n "$(result "$run" findings)" ] && printf -- '- deck %s: %s\n' "$deck" "$(result "$run" findings)" >> "$STATE/findings-index.md"
  case "$status" in
    ready)
      dset "$deck" status waiting pr "$pr" waiting_since "$(date +%s)"
      check_ready "$deck"
      ;;
    no-new-issues)
      say "deck $deck: no new fault. $(result "$run" findings)"
      dset "$deck" status "done"
      mark "$uid" "$deck"
      git worktree remove --force "$run/repo" 2>/dev/null
      ;;
    blocked)
      dset "$deck" status blocked pr "$pr"
      mark "$uid" "$deck"
      notify "decktome live eval: blocked on deck $deck" "$(result "$run" reason)
Questions: $(result "$run" questions)
$( [ -n "$pr" ] && echo "PR #$pr")"
      ;;
    checkpoint)
      local n
      n=$(( $(dget "$deck" continues || echo 0) + 1 ))
      dset "$deck" status checkpoint continues "$n"
      ;;
    *)
      local n
      n=$(( $(dget "$deck" attempts || echo 0) + 1 ))
      dset "$deck" status failed attempts "$n"
      if [ "$n" -ge "$MAX_ATTEMPTS" ]; then
        mark "$uid" "$deck"
        notify "decktome live eval: failed on deck $deck" "The session ended $n times with no result. Log: $run"
      fi
      ;;
  esac
}

check_ready() { # deck -> notify once when GitHub shows the pull request ready
  local deck=$1 pr run out since
  pr=$(dget "$deck" pr)
  run=$HOME_DIR/$deck
  [ -n "$pr" ] || return 0
  if out=$("$TOOL" ready -pr "$pr" 2>&1); then
    dset "$deck" status ready
    mark "$(dget "$deck" uid)" "$deck"
    notify "decktome: PR #$pr is ready to merge" "$(summary_text "$run")
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
  local uid=$1 deck=$2 kind=$3 who=$4 run=$HOME_DIR/$2 template n status
  if ! prepare "$uid" "$deck" "$kind" "$who"; then
    n=$(( $(dget "$deck" attempts || echo 0) + 1 ))
    dset "$deck" status failed attempts "$n"
    say "deck $deck: the worktree or the bundle failed"
    return
  fi
  template=$PROMPTS/eval-prompt.md
  [ "$(dget "$deck" status)" = checkpoint ] && template=$PROMPTS/continue-prompt.md
  render "$template" "$deck" "$run" > "$run/prompt.md"
  rm -f "$run/bundle/result.json"
  n=$(( $(dget "$deck" sessions || echo 0) + 1 ))
  dset "$deck" sessions "$n" status running
  run_claude "$run" "$run/prompt.md" "$run/session-$n.log"
  status=$?
  say "deck $deck: the session ended with exit $status and result $(result "$run" status)"
  finish "$deck" "$run"
}

restack() { # deck target -> rebase the pull request of deck onto target
  local deck=$1 target=$2 run=$HOME_DIR/$1 n
  dset "$deck" restack_target "$target"
  jq --arg t "$target" --arg pr "$(dget "$deck" pr)" '.restack_target = $t | .pr = $pr' "$run/bundle/context.json" > "$run/bundle/context.tmp" \
    && mv "$run/bundle/context.tmp" "$run/bundle/context.json"
  render "$PROMPTS/restack-prompt.md" "$deck" "$run" > "$run/restack-prompt.md"
  rm -f "$run/bundle/result.json"
  n=$(( $(dget "$deck" sessions || echo 0) + 1 ))
  dset "$deck" sessions "$n" status restacking
  run_claude "$run" "$run/restack-prompt.md" "$run/session-$n.log"
  if [ "$(result "$run" status)" = ready ]; then
    local parent=""
    if [ "$target" != main ]; then parent=$(dget "$deck" parent_pr); fi
    dset "$deck" status waiting waiting_since "$(date +%s)" late_notice "" base "$target" parent_pr "$parent"
    check_ready "$deck"
  else
    dset "$deck" status blocked
    notify "decktome live eval: the restack of PR #$(dget "$deck" pr) stopped" "$(result "$run" reason)"
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
    [ -n "$pr" ] && [ -n "$parent" ] || continue
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
    [ "$(dget "$deck" status)" = waiting ] && check_ready "$deck"
  done
}

cleanup_pass() { # remove the worktree of a merged or closed pull request
  local f deck pr state
  for f in "$STATE"/decks/*.json; do
    [ -f "$f" ] || continue
    deck=$(basename "$f" .json)
    pr=$(dget "$deck" pr)
    [ -n "$pr" ] && [ -d "$HOME_DIR/$deck/repo" ] || continue
    state=$(gh pr view "$pr" --json state --jq .state 2>/dev/null)
    case "$state" in
      MERGED|CLOSED)
        git worktree remove --force "$HOME_DIR/$deck/repo" 2>/dev/null && say "removed the worktree of PR #$pr ($state)"
        dset "$deck" status "$(echo "$state" | tr '[:upper:]' '[:lower:]')"
        ;;
    esac
  done
}

eval_pass() {
  local line uid deck kind who status open
  tool pending > "$STATE/pending.jsonl" || { say "the pending list did not read"; return; }
  while IFS= read -r line; do
    [ -n "$line" ] || continue
    uid=$(jq -r .uid <<<"$line")
    deck=$(jq -r .deck <<<"$line")
    kind=$(jq -r .kind <<<"$line")
    who=$(jq -r .who <<<"$line")
    status=$(dget "$deck" status)
    case "$status" in waiting|ready|blocked|done|restacking) continue ;; esac
    if [ "$status" = failed ] && [ "$(dget "$deck" attempts || echo 0)" -ge "$MAX_ATTEMPTS" ]; then continue; fi
    if [ "$status" = checkpoint ] && [ "$(dget "$deck" continues || echo 0)" -gt "$MAX_CONTINUES" ]; then
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

say "summary of every deck that no live eval read, from $PROJECT:"
tool summary || die "the summary did not read"
if [ "$dry" = 1 ]; then
  say "--dry: no mark and no session"
  exit 0
fi

if [ ! -f "$STATE/initialized" ] && [ "$backlog" = 0 ]; then
  say "first run: every deck above is marked read, and no session starts for it (D-1133)"
  tool pending | while IFS= read -r line; do
    deck=$(jq -r .deck <<<"$line")
    tool mark -uid "$(jq -r .uid <<<"$line")" -deck "$deck" && dset "$deck" status backlog
  done
fi
touch "$STATE/initialized"

while true; do
  cleanup_pass
  restack_pass
  waiting_pass
  eval_pass
  [ "$once" = 1 ] && break
  say "next poll in $INTERVAL seconds"
  sleep "$INTERVAL"
done
