#!/usr/bin/env bash
# The automated tuning loop (D-133). It runs the question gate, scores
# every question with the eval role, hands the report to a fixer agent,
# and keeps each change the fixer made only when the run got better on
# the rows that change declared (D-181).
#
# CAUTION: every iteration calls real providers and costs money. The loop
# refuses to start without AUTOTUNE_ALLOW_UNATTENDED=1, and it stops at
# the budget.
#
# The loop never runs on a branch it did not make. It reverts a change
# that fails the build, that touches a frozen file, or that the checker
# rejects. Nothing reaches main.
#
# The loop learns. Every decision, kept or dropped, is written to
# docs/reference/autotune-lessons.md with the hypothesis and the questions
# that moved, and the next fixer reads that file before it edits (D-182).
# The evidence of a dropped change is kept under .local/tune/rejected/
# (D-180).
#
# The loop commits to its own branch and does not push. Add --push when
# the owner wants the branch on the remote as it goes.
#
#   AUTOTUNE_ALLOW_UNATTENDED=1 scripts/autotune.sh --budget 3.00 --max 20 \
#     --base pr-7c --baseline .local/tune/run18.json
#
# To stop the loop without losing work, run:
#
#   touch .local/tune/STOP
#
# The iteration that is running finishes and commits, and the loop ends
# before the next one starts. Ctrl-C also works and it abandons the
# iteration in flight.
#
# --base names the branch the owner keeps for loop output. The run cuts a
# working branch off it and pushes nothing. After the review, the owner
# fast-forwards the base, and the next night continues from there.
#
# The loop stops, and rejects nothing, on a tool fault: an unpriced run,
# a partial eval, a checker that exits 2, or a checker that does not
# build. D-171 was a checker fault read as a rejection (T-12).
#
# pipefail is on, so a failed command inside a pipe fails the pipe. The
# merge-failure branch never fired without it (T-11). set -e is not on:
# the loop reads exit codes on purpose in many places.
set -u -o pipefail
cd "$(dirname "$0")/.." || exit 1
ROOT="$(pwd)"

BUDGET="3.00"
BASELINE_JSON=""
# BASELINE_DOC is the eval report that goes with BASELINE_JSON. The fixer
# reads the report, and the checker reads the JSON. Supplying a baseline
# skips the -000 run, so without this the first iteration had no report
# to act on and the fixer failed at once (D-160).
BASELINE_DOC=""
# BASE_REF is the branch the run starts from. Empty means the current
# branch. A named long-lived branch is what makes two nights add up
# instead of diverging from one fixed point (D-142).
BASE_REF=""
MAX_ITERATIONS="20"
TARGET_RATIO="0.05"
EVAL_BUDGET="0.50"
# NOISE is how many bad questions two runs of identical code differ by.
# Empty means the Go constant tune.DefaultNoise, which is 9 (D-230). The
# script carried its own 3 from D-183 for two days after D-230 raised the
# constant, so the two disagreed (T-7, A-3).
NOISE=""
BRANCH_PREFIX="auto-tune"
DRY_RUN="0"
# Nothing leaves the machine unless the owner asks. The loop commits to a
# branch of its own, and the owner pushes after reading it (OQ-25).
PUSH="0"

while [ $# -gt 0 ]; do
  case "$1" in
    --budget) BUDGET="$2"; shift 2 ;;
    --max) MAX_ITERATIONS="$2"; shift 2 ;;
    --target) TARGET_RATIO="$2"; shift 2 ;;
    --eval-budget) EVAL_BUDGET="$2"; shift 2 ;;
    --noise) NOISE="$2"; shift 2 ;;
    --branch) BRANCH_PREFIX="$2"; shift 2 ;;
    --push) PUSH="1"; shift ;;
    --no-push) PUSH="0"; shift ;;
    --base) BASE_REF="$2"; shift 2 ;;
    --baseline) BASELINE_JSON="$2"; shift 2 ;;
    --baseline-doc) BASELINE_DOC="$2"; shift 2 ;;
    --dry-run) DRY_RUN="1"; shift ;;
    *) echo "unknown flag: $1" >&2; exit 2 ;;
  esac
done

STATE_DIR="$ROOT/.local/tune"
# STOP_FILE lets a person end the loop without losing an iteration.
# Ctrl-C kills the gate or the eval part way, and that iteration is
# abandoned with its spend wasted. Touch this file instead: the iteration
# that is running finishes and commits, and the loop stops before the
# next one. The loop removes the file, so a stale one never blocks a
# later run.
STOP_FILE="$STATE_DIR/STOP"
LOG="$STATE_DIR/autotune.log"
LESSONS="$ROOT/docs/reference/autotune-lessons.md"
REJECTED_DIR="$STATE_DIR/rejected"
# TUNE_CHECK is the checker, built once at the start. A go run each time
# would read the exit code of the Go toolchain when the tree does not
# build, and that code is the reject code (T-12).
TUNE_CHECK="$STATE_DIR/tune-check.bin"
mkdir -p "$STATE_DIR"

# say writes to the log and to stderr, never to stdout. run_gate and
# run_eval hand their path back through stdout, and a log line on the same
# stream would ride along with it.
say() { printf '%s  %s\n' "$(date -u +%H:%M:%S)" "$*" | tee -a "$LOG" >&2; }
die() { say "STOP: $*"; exit 1; }

# --- Frozen files. The fixer may not change how it is measured. --------
# A loop that may edit its own scorer optimizes the scorer. Every path
# here is part of the measurement, not part of the product. The fixer
# prompt and the lessons file are the loop's memory, and a fixer that
# edits its own memory forgets what it was told (D-182). The gate
# command, its metrics, the role table, and the price table are the
# scorecard, and they joined the list on 2026-08-28 (T-10).
FROZEN="
go/cmd/questions-eval
go/cmd/questions-gate/main.go
go/cmd/questions-gate/conversations.json
go/cmd/tune-check
go/internal/tune
go/internal/questions/lint.go
go/internal/questions/lint_test.go
go/internal/questions/metrics.go
go/internal/llm/roles.json
go/internal/llm/prices.json
scripts/autotune.sh
scripts/autotune-fix.sh
docs/owner-questions.md
docs/reference/autotune-fixer-prompt.md
docs/reference/autotune-lessons.md
"

# frozen_touched reads every change since the last good commit: the
# fixer's own commits and whatever it left in the tree. The match is a
# fixed string, so a dot in a path is a dot (T-10).
frozen_touched() {
  local changed path
  changed="$(git diff --name-only "$LAST_GOOD"; git ls-files --others --exclude-standard)"
  for path in $FROZEN; do
    if grep -qxF -- "$path" <<<"$changed" || grep -qF -- "$path/" <<<"$changed"; then
      echo "$path"
      return 0
    fi
  done
  # The M-5 sheets are the owner's record. Nothing may rewrite one.
  if grep -q '^docs/reference/pr7-m5-scoring' <<<"$changed"; then
    echo "docs/reference/pr7-m5-scoring*"
    return 0
  fi
  # docs/decisions.md is append-only: a removed line rewrites history.
  if git diff -U0 "$LAST_GOOD" -- docs/decisions.md | grep -q '^-[^-]'; then
    echo "docs/decisions.md (a line was removed)"
    return 0
  fi
  return 1
}

# --- Budget ------------------------------------------------------------
spent() { [ -f "$LEDGER" ] && awk '{s+=$1} END {printf "%.4f", s+0}' "$LEDGER" || echo "0.0000"; }
# charge adds one priced run to the ledger. It runs at the top level and
# never inside a command substitution, so a cost that is empty or not a
# number stops the loop here. An unpriced run charged $0 before, because
# the die inside $(...) killed a subshell alone (T-1).
charge() {  # $1 = cost in USD, $2 = what it paid for
  case "$1" in
    ''|*[!0-9.]*|.|*.*.*) die "$2 carries no priced cost (got '${1:-}'), so the budget can not count it. Add the model to prices.json." ;;
  esac
  printf '%s %s\n' "$1" "$2" >> "$LEDGER"
}
over_budget() { awk -v s="$(spent)" -v b="$BUDGET" 'BEGIN {exit !(s >= b)}'; }

# cost_of reads the "- Cost: $0.1234." line a run document writes. It
# prints nothing for a document with no priced line, and charge decides.
cost_of() {
  grep -o -- '- Cost: \$[0-9.]*' "$1" 2>/dev/null | tail -1 | tr -d '$' | awk '{print $3+0}'
}

green() {
  ( cd "$ROOT/go" && go build ./... && go vet ./... && go test ./... >/dev/null ) \
    && ( cd "$ROOT" && make lint-go >/dev/null 2>&1 )
}

# --- Preflight ---------------------------------------------------------
[ "${AUTOTUNE_ALLOW_UNATTENDED:-0}" = "1" ] || \
  die "this loop edits code and pushes with no person watching. Set AUTOTUNE_ALLOW_UNATTENDED=1 to allow it."
[ -f "$ROOT/.env" ] || die "no .env, so no API keys"
# Not "A && B || C": that runs C when A succeeds and B fails, which is
# the same answer here, but shellcheck can not know it (SC2015).
if ! git diff --quiet || ! git diff --cached --quiet; then
  die "the working tree is dirty. Commit or stash first."
fi
# An untracked file would ride into a loop commit under the owner's name,
# because the loop commits with git add -A (audit H-1).
untracked="$(git ls-files --others --exclude-standard)"
[ -z "$untracked" ] || die "untracked files would ride into a loop commit. Commit, stash, or remove: $(tr '\n' ' ' <<<"$untracked")"

# The checker is built once, from the frozen code, before anything else
# runs. Its exit codes are then its own (T-12).
( cd "$ROOT/go" && go build -o "$TUNE_CHECK" ./cmd/tune-check ) || die "the checker does not build"
if [ -z "$NOISE" ]; then
  NOISE="$("$TUNE_CHECK" -print-noise)" || die "the checker did not print its noise margin"
fi

# The base is a branch the owner keeps. Every night cuts a working branch
# off it, and the owner fast-forwards the base after the review. Two
# nights therefore add up. A loop that always starts from main would give
# the second night none of the first night's accepted work, and the owner
# would merge two branches that changed the same rows (D-142).
if [ -n "$BASE_REF" ]; then
  git rev-parse --verify --quiet "$BASE_REF" >/dev/null || die "no branch $BASE_REF"
  git switch "$BASE_REF" >/dev/null 2>&1 || die "could not switch to $BASE_REF"
fi
BASE_BRANCH="$(git rev-parse --abbrev-ref HEAD)"
case "$BASE_BRANCH" in
  main|master) say "WARNING: the base is $BASE_BRANCH. Two nights from here will diverge, not add up." ;;
esac
STAMP="$(date -u +%Y%m%d-%H%M%S)"
# The ledger is per run. It used to be one file for every run ever, so
# --budget counted spend from earlier nights and each run had less to
# spend than the one before it. The whole record is the set of these
# files (D-179).
LEDGER="$STATE_DIR/ledger-$STAMP.txt"
BRANCH="$BRANCH_PREFIX/$STAMP"
say "base branch $BASE_BRANCH, working branch $BRANCH"
[ "$DRY_RUN" = "1" ] || git switch -c "$BRANCH" >/dev/null 2>&1 || die "could not make branch $BRANCH"

# The directive must sit on the line above the dot, so the three commands
# are split: on one line it attaches to "set -a" and the dot stays
# flagged.
set -a
# The file holds secrets and is gitignored, so shellcheck can not read it.
# shellcheck source=/dev/null
. "$ROOT/.env"
set +a
export CARDS_SNAPSHOT_DIR="${CARDS_SNAPSHOT_DIR:-$ROOT/.local/gcs/mtg-local-cards/scryfall}"

# conv_count reads the size of the conversation set. A hardcoded number
# goes stale: the set grew from 52 to 66 to 100 to 104, and the log still
# said 66 (D-145, D-155).
conv_count() {
  python3 -c "import json,sys;d=json.load(open(sys.argv[1]));print(len(d if isinstance(d,list) else d['conversations']))" \
    "$ROOT/go/cmd/questions-gate/conversations.json" 2>/dev/null || echo "?"
}

# run_gate and run_eval hand a path back on stdout and charge nothing.
# They run inside a command substitution, so a die in them would kill the
# subshell alone. The caller charges at the top level (T-1).
run_gate() {   # $1 = iteration label, $2 = conversation ids to run (optional, comma separated)
  local out="$ROOT/docs/reference/pr7-question-gate-$1.md" only="${2:-}" scope
  # A scored document is never overwritten (D-65, D-177).
  if [ -e "$out" ]; then say "STOP: a gate document already exists at $out"; return 1; fi
  scope="$(conv_count) conversations, about 20 minutes"
  [ -n "$only" ] && scope="conversations $only alone"
  say "gate $1: $scope"
  [ "$DRY_RUN" = "1" ] && { say "dry run: no gate"; return 1; }
  ( cd "$ROOT/go" && QUESTIONS_GATE=1 go run ./cmd/questions-gate \
      -collection internal/collections/testdata/manabox_collection.csv \
      ${only:+-only "$only"} > "$out" ) || true
  # A run that failed leaves an empty file, and an empty file would block
  # the label forever. Remove it and say so.
  [ -s "$out" ] || { rm -f "$out"; return 1; }
  echo "$out"
}

run_eval() {   # $1 = iteration label, $2 = gate document
  local doc="$ROOT/docs/reference/pr7-question-eval-$1.md"
  local sum="$STATE_DIR/$1.json"
  if [ -e "$doc" ]; then say "STOP: an eval report already exists at $doc"; return 1; fi
  if [ -e "$sum" ]; then say "STOP: an eval summary already exists at $sum"; return 1; fi
  say "eval $1: scoring every question"
  ( cd "$ROOT/go" && QUESTIONS_EVAL=1 go run ./cmd/questions-eval \
      -in "$2" -out "$doc" -json "$sum" -budget "$EVAL_BUDGET" ) || true
  [ -s "$sum" ] || { rm -f "$sum" "$doc"; return 1; }
  # A budget-cut eval scores fewer questions, and a smaller count read as
  # an improvement once. It is no baseline and no candidate (T-2).
  if is_partial "$sum"; then
    say "STOP: the eval of $1 is partial: $(stopped_reason "$sum")"
    return 2
  fi
  echo "$sum"
}

# eval_doc_of names the eval report that goes with an iteration label.
eval_doc_of() { echo "$ROOT/docs/reference/pr7-question-eval-$1.md"; }

# fixer_json_of writes the summary the fixer's checkout may hold: the
# same file with every holdout verdict removed (T-8).
fixer_json_of() {  # $1 = eval summary
  local out="${1%.json}.fixer.json"
  [ -s "$out" ] || "$TUNE_CHECK" -fixer -next "$1" -out "$out" || die "could not write the fixer summary for $1"
  echo "$out"
}

# commit_all commits everything in the tree. The result is checked: a
# commit that failed left accepted work uncommitted, and the next revert
# to LAST_GOOD destroyed it with no log line (audit M-11).
commit_all() {  # $1 = version, $2 = subject, $3 = eval summary (optional)
  if [ "$DRY_RUN" = "1" ]; then say "dry run: no commit"; return 0; fi
  case "$(git rev-parse --abbrev-ref HEAD)" in
    main|master) die "refusing to commit to $(git rev-parse --abbrev-ref HEAD)" ;;
  esac
  git add -A
  git diff --cached --quiet && { say "nothing to commit"; return 0; }
  # Declared apart from the assignment, so a failure inside the command
  # substitution is not masked by local's own exit status (SC2155).
  local body
  body="Automated tuning loop, budget spent \$$(spent) of \$$BUDGET."
  if [ -n "${3:-}" ]; then
    body="Bad-question ratio $(ratio_of "$3")%. Questions asked $(metric_of "$3" questions).
$body"
  fi
  git commit -q -F - <<EOF || die "git commit failed for $1, and the tree holds the work uncommitted"
$1: $2

$body
EOF
  if [ "$PUSH" = "1" ]; then
    git push -q -u origin "$BRANCH" || say "WARNING: the push of $BRANCH failed. The commit is local."
  fi
  say "committed $1"
}

ratio_of()  { python3 -c "import json,sys;print(round(json.load(open(sys.argv[1]))['bad_ratio']*100,1))" "$1" 2>/dev/null || echo "?"; }
metric_of() { python3 -c "import json,sys;print(json.load(open(sys.argv[1]))['metrics'][sys.argv[2]])" "$1" "$2" 2>/dev/null || echo "?"; }
is_partial() { python3 -c "import json,sys;sys.exit(0 if json.load(open(sys.argv[1])).get('partial') else 1)" "$1" 2>/dev/null; }
stopped_reason() { python3 -c "import json,sys;print(json.load(open(sys.argv[1])).get('stopped_reason',''))" "$1" 2>/dev/null || echo "?"; }
# scored_bad_of reads the count the checker scores: the holdout's bad
# questions when there is a holdout, and every bad question otherwise.
scored_bad_of() {
  python3 -c "import json,sys;d=json.load(open(sys.argv[1]));print(d['holdout_bad'] if d.get('holdout_judged',0)>0 else d['bad'])" "$1" 2>/dev/null || echo "?"
}

# --- The fixer's changes -----------------------------------------------
# The fixer commits each independent change on its own, with two trailers:
# "Rows:" names the catalog rows the change means to move, and
# "Hypothesis:" says why. The checker charges every moved question to the
# change that owns its row (D-181).
changes_json() {  # $1 = output path
  git log --reverse --format='%H%x1f%s%x1f%b%x1e' "$LAST_GOOD..HEAD" | python3 -c '
import sys, json
out = []
for rec in sys.stdin.read().split("\x1e"):
    if not rec.strip():
        continue
    parts = (rec.strip("\n").split("\x1f") + ["", ""])[:3]
    sha, subject, body = parts
    rows, hyp = [], ""
    for line in body.splitlines():
        l = line.strip()
        if l.lower().startswith("rows:"):
            rows = [r.strip().strip("`") for r in l[5:].split(",") if r.strip()]
        elif l.lower().startswith("hypothesis:"):
            hyp = l[11:].strip()
    out.append({"commit": sha, "subject": subject, "rows": rows, "hypothesis": hyp})
json.dump(out, open(sys.argv[1], "w"), indent=2)
print(len(out))' "$1"
}

# strip_attribution removes AI-attribution lines from the fixer's tip
# commit. The house rule forbids them in any commit (CLAUDE.md rule 6),
# and a fixer tool may add one by default. Only the tip is amended: a
# filter-branch over the range rewrote every commit, and an older commit
# that carries one is a rejection instead.
strip_attribution() {
  local pattern='co-authored-by|generated with|claude code' msgs
  msgs="$(git log --format=%B "$LAST_GOOD..HEAD")"
  grep -qiE "$pattern" <<<"$msgs" || return 0
  if git log -1 --format=%B | grep -qiE "$pattern"; then
    say "the tip commit carried attribution text, which the house rule forbids: removed"
    git log -1 --format=%B | grep -viE "$pattern" | git commit -q --amend -F - || die "could not amend the tip commit"
  fi
  msgs="$(git log --format=%B "$LAST_GOOD..HEAD")"
  if grep -qiE "$pattern" <<<"$msgs"; then
    say "a commit below the tip carries attribution text, and the loop amends the tip alone"
    return 1
  fi
}

# preserve keeps the evidence of an iteration before the tree is reset:
# the diff, the commits, the gate and eval documents, and the verdict.
# A rejected iteration is a paid measurement, and the fixer's diff is the
# one artifact that says what was tried (D-180).
preserve() {  # $1 = label, $2 = why
  local dir="$REJECTED_DIR/$1" f
  mkdir -p "$dir"
  git diff "$LAST_GOOD" > "$dir/changes.patch"
  git log --format='%H %s%n%b%n' "$LAST_GOOD..HEAD" > "$dir/commits.txt"
  for f in "$ROOT"/docs/reference/pr7-question-gate-"$1"*.md "$ROOT"/docs/reference/pr7-question-eval-"$1"*.md "$STATE_DIR/$1"*.json "$STATE_DIR/$1"*.txt; do
    [ -f "$f" ] && cp "$f" "$dir/"
  done
  printf '%s\n' "$2" > "$dir/why.txt"
  say "kept the evidence of $1 under $dir"
}

# revert_to_last_good resets the working branch. It never runs in a dry
# run, where the branch is the owner's base and LAST_GOOD is its tip: a
# reset there threw away the owner's uncommitted state once (T-13).
revert_to_last_good() {
  if [ "$DRY_RUN" = "1" ]; then say "dry run: no revert"; return 0; fi
  git reset -q --hard "$LAST_GOOD"
  git clean -qfd
}

# reject preserves the evidence, reverts, records the lesson, and counts
# the rejection. Every reject path runs through here, so none can skip
# the preserve step (audit 2026-08-28).
reject() {  # $1 = label, $2 = why
  say "REVERT: $2"
  preserve "$1" "$2"
  revert_to_last_good
  record_lesson "$1" "rejected"
  rejects=$((rejects+1))
}

# record_lesson appends what the checker learned to the lessons file and
# commits it, so a rejected iteration still leaves its lesson on the
# branch. The checker wrote the lesson to a scratch file, because the
# revert wipes the tree (D-182).
record_lesson() {  # $1 = label, $2 = outcome word
  local scratch="$STATE_DIR/$1-lesson.md"
  [ -s "$scratch" ] || return 0
  [ "$DRY_RUN" = "1" ] && return 0
  cat "$scratch" >> "$LESSONS"
  git add "$LESSONS"
  git commit -q -m "lesson: $1 $2" -m "The tuning loop recorded what iteration $1 taught (D-182)." \
    || die "git commit of the lesson failed"
  LAST_GOOD="$(git rev-parse HEAD)"
}

# --- Baseline ----------------------------------------------------------
say "budget \$$BUDGET, target ratio $TARGET_RATIO, noise margin $NOISE, at most $MAX_ITERATIONS iterations"
# A scored run already on disk is the baseline. The loop pays for its own
# only when the owner gives it none. A baseline the owner named and the
# loop can not read stops the run: a typo must not buy a new baseline
# (audit H-4).
# tune-check runs from go/, so a relative baseline path resolves against
# the wrong directory. Make both absolute before anything reads them.
case "$BASELINE_JSON" in ""|/*) ;; *) BASELINE_JSON="$ROOT/$BASELINE_JSON" ;; esac
case "$BASELINE_DOC" in ""|/*) ;; *) BASELINE_DOC="$ROOT/$BASELINE_DOC" ;; esac

# BEST is the best accepted summary of the night, and the checker
# compares every candidate against it. An accepted run that is a margin
# worse than BEST is committed and does not replace it, so the baseline
# can not drift down one margin per iteration (T-6). PREV_DOC is the
# latest accepted report, which the fixer acts on.
if [ -n "$BASELINE_JSON" ]; then
  [ -s "$BASELINE_JSON" ] || die "no readable baseline at $BASELINE_JSON"
  is_partial "$BASELINE_JSON" && die "the baseline $BASELINE_JSON is a partial eval: $(stopped_reason "$BASELINE_JSON")"
  BEST="$BASELINE_JSON"
  BEST_LABEL="$(basename "$BASELINE_JSON" .json)"
  # The fixer needs the report, not the summary. Take the name the owner
  # gave, or the one that matches the summary: .local/tune/run18.json goes
  # with docs/reference/pr7-question-eval-run18.md.
  if [ -z "$BASELINE_DOC" ]; then
    BASELINE_DOC="$(eval_doc_of "$BEST_LABEL")"
  fi
  [ -f "$BASELINE_DOC" ] || die "no eval report at $BASELINE_DOC. Name it with --baseline-doc."
  PREV_DOC="$BASELINE_DOC"
  PREV_JSON="$BEST"
  # A summary from before D-181 carries no per-conversation counts, and a
  # partial re-measure needs them. The gate document supplies them.
  PREV_GATE="$ROOT/docs/reference/pr7-question-gate-$BEST_LABEL.md"
  say "baseline read from $BEST, ratio $(ratio_of "$BEST")%"
  say "baseline report $PREV_DOC"
else
  GATE="$(run_gate "$STAMP-000")" || die "the baseline gate produced nothing"
  charge "$(cost_of "$GATE")" "gate-$STAMP-000"
  BEST="$(run_eval "$STAMP-000" "$GATE")" || die "the baseline eval produced nothing"
  charge "$(cost_of "$(eval_doc_of "$STAMP-000")")" "eval-$STAMP-000"
  BEST_LABEL="$STAMP-000"
  PREV_DOC="$(eval_doc_of "$STAMP-000")"
  PREV_JSON="$BEST"
  PREV_GATE="$GATE"
  say "baseline ratio $(ratio_of "$BEST")%, spent \$$(spent)"
  commit_all "v0.0" "baseline run for the tuning loop" "$BEST"
fi
LAST_GOOD="$(git rev-parse HEAD)"
[ -f "$LESSONS" ] || { printf '# What the tuning loop learned\n\nThe loop appends one block per iteration (D-182).\n' > "$LESSONS"; commit_all "lessons" "start the lessons file"; LAST_GOOD="$(git rev-parse HEAD)"; }

# accept_run commits an accepted iteration and moves the baseline. BEST
# moves only when the new run scored at least as well (T-6).
accept_run() {  # $1 = version, $2 = subject, $3 = eval summary, $4 = eval report, $5 = gate document, $6 = label, $7 = outcome word
  commit_all "$1" "$2" "$3"
  LAST_GOOD="$(git rev-parse HEAD)"
  record_lesson "$6" "$7"
  PREV_DOC="$4"; PREV_JSON="$3"; PREV_GATE="$5"; rejects=0
  local got was
  got="$(scored_bad_of "$3")"; was="$(scored_bad_of "$BEST")"
  if [ "$got" != "?" ] && [ "$was" != "?" ] && [ "$got" -le "$was" ]; then
    BEST="$3"; BEST_LABEL="$6"
    say "best so far: $BEST_LABEL with $got scored bad questions"
  else
    say "accepted inside the margin, and the baseline stays $BEST_LABEL ($was scored bad questions against $got)"
  fi
}

# --- Loop --------------------------------------------------------------
i=0
rejects=0
while [ "$i" -lt "$MAX_ITERATIONS" ]; do
  i=$((i+1))
  # The label carries the run stamp, so no two runs ever write the same
  # document (D-177).
  LABEL="$(printf '%s-%03d' "$STAMP" "$i")"
  # An exit here happens before the iteration ran, so the count steps
  # back and the closing line reports the iterations that did run.
  if [ -f "$STOP_FILE" ]; then
    i=$((i-1))
    say "stop file found, so the loop ends after $i iterations"
    rm -f "$STOP_FILE"
    break
  fi
  if over_budget; then i=$((i-1)); say "the budget of \$$BUDGET is spent"; break; fi
  if [ "$rejects" -ge 3 ]; then i=$((i-1)); say "three iterations in a row changed nothing that held"; break; fi

  say "--- iteration $i, spent \$$(spent) of \$$BUDGET, baseline $BEST_LABEL"

  FIXER_JSON="$(fixer_json_of "$PREV_JSON")"
  if ! AUTOTUNE_EVAL_DOC="$PREV_DOC" AUTOTUNE_LABEL="$LABEL" AUTOTUNE_LESSONS="$LESSONS" \
       AUTOTUNE_LAST_GOOD="$LAST_GOOD" AUTOTUNE_FIXER_JSON="$FIXER_JSON" AUTOTUNE_DRY_RUN="$DRY_RUN" \
       "$ROOT/scripts/autotune-fix.sh"; then
    say "the fixer failed"; revert_to_last_good; rejects=$((rejects+1)); continue
  fi
  # A dry run stops here. The fixer wrote its prompt and ran no agent, so
  # there is nothing to commit, to measure, or to revert (T-13).
  if [ "$DRY_RUN" = "1" ]; then
    say "dry run: the fixer prompt is written, and no gate, eval, commit, or revert runs"
    break
  fi
  # Edits the fixer left uncommitted become one change with no rows. The
  # checker drops a change with no rows (T-4).
  if ! git diff --quiet HEAD || [ -n "$(git ls-files --others --exclude-standard)" ]; then
    commit_all "change" "unsplit edits of iteration $i"
  fi
  if [ "$(git rev-parse HEAD)" = "$LAST_GOOD" ]; then
    say "the fixer changed nothing"; rejects=$((rejects+1)); continue
  fi
  if path="$(frozen_touched)"; then
    reject "$LABEL" "the fixer touched a frozen path: $path"; continue
  fi
  if ! strip_attribution; then
    reject "$LABEL" "a commit below the tip carries attribution text"; continue
  fi
  CHANGES="$STATE_DIR/$LABEL-changes.json"
  say "the fixer made $(changes_json "$CHANGES") changes"
  if ! green; then
    reject "$LABEL" "the tree is not green"; continue
  fi

  GATE="$(run_gate "$LABEL")" || { reject "$LABEL" "the gate produced nothing"; continue; }
  charge "$(cost_of "$GATE")" "gate-$LABEL"
  NEXT="$(run_eval "$LABEL" "$GATE")"
  rc=$?
  [ "$rc" = "2" ] && { preserve "$LABEL" "the eval was partial"; die "the eval of $LABEL is partial, and a partial run decides nothing"; }
  [ "$rc" = "0" ] || { reject "$LABEL" "the eval produced nothing"; continue; }
  charge "$(cost_of "$(eval_doc_of "$LABEL")")" "eval-$LABEL"

  VERDICT="$STATE_DIR/$LABEL-verdict.txt"
  rm -f "$STATE_DIR/$LABEL-lesson.md"
  ( cd "$ROOT/go" && "$TUNE_CHECK" -next "$NEXT" -prev "$BEST" -target "$TARGET_RATIO" \
      -changes "$CHANGES" -noise "$NOISE" -lessons "$STATE_DIR/$LABEL-lesson.md" -label "$LABEL" ) | tee -a "$LOG" > "$VERDICT"
  code="${PIPESTATUS[0]}"
  case "$code" in
    0|3)
      accept_run "v0.$i" "tuning iteration $i" "$NEXT" "$(eval_doc_of "$LABEL")" "$GATE" "$LABEL" "accepted"
      [ "$code" = "3" ] && { say "target reached"; break; }
      ;;
    1)
      reject "$LABEL" "the checker rejected iteration $i: $(grep '^reject: ' "$VERDICT" | head -3 | tr '\n' ' ')"
      ;;
    4)
      # Some changes are kept and some are dropped. The kept subset was
      # never measured alone, so the loop runs the conversations its rows
      # touch again, and folds that measurement into the baseline (D-181).
      keeps="$(awk '/^keep: /{print $2}' "$VERDICT")"
      only="$(awk '/^remeasure: /{print $2}' "$VERDICT")"
      preserve "$LABEL" "partial: $(awk '/^drop: /{printf "%s ", $2}' "$VERDICT")dropped"
      revert_to_last_good
      ok=1
      for sha in $keeps; do
        git cherry-pick -q "$sha" >/dev/null 2>&1 || { git cherry-pick --abort >/dev/null 2>&1; ok=0; break; }
      done
      if [ "$ok" != "1" ]; then
        reject "$LABEL" "the kept changes do not apply without the dropped ones"; continue
      fi
      if ! green; then
        reject "$LABEL" "the kept changes alone are not green"; continue
      fi
      RLABEL="$LABEL-kept"
      RGATE="$(run_gate "$RLABEL" "$only")" || { reject "$RLABEL" "the re-measure gate produced nothing"; continue; }
      charge "$(cost_of "$RGATE")" "gate-$RLABEL"
      RNEXT="$(run_eval "$RLABEL" "$RGATE")"
      rc=$?
      [ "$rc" = "2" ] && { preserve "$RLABEL" "the re-measure eval was partial"; die "the eval of $RLABEL is partial, and a partial run decides nothing"; }
      [ "$rc" = "0" ] || { reject "$RLABEL" "the re-measure eval produced nothing"; continue; }
      charge "$(cost_of "$(eval_doc_of "$RLABEL")")" "eval-$RLABEL"
      MERGED="$STATE_DIR/$RLABEL-merged.json"
      MERGED_DOC="$(eval_doc_of "$RLABEL-merged")"
      if ! ( cd "$ROOT/go" && "$TUNE_CHECK" -merge -prev "$PREV_JSON" -prev-doc "$PREV_GATE" \
          -next "$RNEXT" -out "$MERGED" -out-doc "$MERGED_DOC" ) | tee -a "$LOG" >&2; then
        preserve "$RLABEL" "the merge failed"
        die "the merge of $RLABEL failed, which is a tool fault"
      fi
      ( cd "$ROOT/go" && "$TUNE_CHECK" -next "$MERGED" -prev "$BEST" -target "$TARGET_RATIO" -noise "$NOISE" ) | tee -a "$LOG" >&2
      rcode="${PIPESTATUS[0]}"
      case "$rcode" in
        0|3)
          accept_run "v0.$i" "tuning iteration $i, kept changes alone" "$MERGED" "$MERGED_DOC" "" "$LABEL" "partly kept"
          [ "$rcode" = "3" ] && { say "target reached"; break; }
          ;;
        1)
          reject "$RLABEL" "the kept changes alone did not hold"
          ;;
        *)
          preserve "$RLABEL" "the checker exited $rcode"
          die "the checker exited $rcode on $RLABEL, which is a tool fault and not a verdict (T-12)"
          ;;
      esac
      ;;
    *)
      # Exit 2 is a fault: a partial summary, two owners of one row, or a
      # file the checker could not read. It is nobody's rejection (T-12).
      preserve "$LABEL" "the checker exited $code"
      die "the checker exited $code on $LABEL, which is a tool fault and not a verdict (T-12). Read $VERDICT."
      ;;
  esac
done

say "done. Spent \$$(spent) of \$$BUDGET over $i iterations. Branch $BRANCH. Best run $BEST_LABEL."
say "read $LOG, $LESSONS, and the eval documents in docs/reference/."
