#!/usr/bin/env bash
# The feedback fix cycle (PR-28c, D-557 to D-559, D-645). It reads a
# harvest, turns every thumbs down into a test case, proves each case
# fails, hands the failures to a fixer agent, and proves they pass.
#
# The result is one pull request that holds the case AND the fix, so the
# gate on main is never red (D-645).
#
# CAUTION: a live cycle calls real providers and costs money. It refuses
# to start without FEEDBACK_LOOP_ALLOW=1, and it stops at --cap, which is
# $2.00 by default (D-559).
#
# The cycle never runs on a branch it did not make. It reverts the fixer
# when the tree goes red, when the fixer touches a frozen path, or when a
# case still fails.
#
#   FEEDBACK_LOOP_ALLOW=1 scripts/feedback-loop.sh --cap 2.00
#   scripts/feedback-loop.sh --dry
#
# AUTOTUNE_FIXER_CMD names the agent, the way the tuning loop does. This
# script ships no default: an unattended agent that edits a repository is
# the owner's call (docs/reference/autotune-design.md).
#
# pipefail is on, so a failed command inside a pipe fails the pipe. set -e
# is not on: the cycle reads exit codes on purpose in many places.
set -uo pipefail
cd "$(dirname "$0")/.." || exit 1
ROOT="$(pwd)"

CAP="2.00"
HARVEST=""
BASE_REF=""
DRY_RUN="0"
ROUNDS="3"
OPEN_PR="1"

# FROZEN holds what the fixer may not touch. The cases are the
# measurement: a fixer that edits one makes the gate agree with the code
# instead of with the reader. The rest is the scorer, the accept rules,
# and the loop itself (T-10).
FROZEN="
go/cmd/questions-gate/conversations.json
go/cmd/deck-gate/prompts.json
go/cmd/bracket-gate/prompts.json
go/cmd/case-check
go/cmd/feedback-triage
go/internal/triage
go/cmd/questions-eval
go/cmd/tune-check
go/internal/tune
go/internal/questions/lint.go
go/internal/questions/lint_test.go
go/internal/llm/roles.json
go/internal/llm/prices.json
scripts/feedback-loop.sh
scripts/autotune.sh
scripts/autotune-fix.sh
docs/owner-questions.md
docs/reference/feedback-fixer-prompt.md
docs/reference/autotune-fixer-prompt.md
docs/reference/autotune-lessons.md
docs/reference/eval/baselines.json
"

usage() {
  sed -n '2,26p' "$0" | sed 's/^# \{0,1\}//'
  exit "${1:-0}"
}

while [ $# -gt 0 ]; do
  case "$1" in
    --cap) CAP="${2:-}"; shift 2 ;;
    --in) HARVEST="${2:-}"; shift 2 ;;
    --base) BASE_REF="${2:-}"; shift 2 ;;
    --rounds) ROUNDS="${2:-}"; shift 2 ;;
    --dry) DRY_RUN="1"; shift ;;
    --no-pr) OPEN_PR="0"; shift ;;
    -h|--help) usage 0 ;;
    *) echo "feedback-loop: unknown flag $1" >&2; usage 1 ;;
  esac
done

say() { printf '%s  %s\n' "$(date -u +%H:%M:%S)" "$*" | tee -a "$LOG" >&2; }
die() { printf 'feedback-loop: %s\n' "$*" >&2; exit 1; }

STAMP="$(date -u +%Y%m%d-%H%M%S)"
STATE_DIR="$ROOT/.local/tune/feedback-$STAMP"
mkdir -p "$STATE_DIR"
LOG="$STATE_DIR/loop.log"
: > "$LOG"
MANIFEST="$STATE_DIR/cases.json"
TRIAGE_DOC="$STATE_DIR/triage.md"
REPORT="$STATE_DIR/report.md"
LEDGER="$STATE_DIR/ledger"
: > "$LEDGER"

# --- Preflight ---------------------------------------------------------

if [ "$DRY_RUN" != "1" ]; then
  [ "${FEEDBACK_LOOP_ALLOW:-0}" = "1" ] || \
    die "this cycle edits code, commits, and opens a pull request with no person watching. Set FEEDBACK_LOOP_ALLOW=1 to allow it."
  [ -n "${AUTOTUNE_FIXER_CMD:-}" ] || \
    die "set AUTOTUNE_FIXER_CMD to the agent that applies fixes. docs/reference/autotune-design.md names the trade-off."
  [ -f "$ROOT/.env" ] || die "no .env, so no provider keys"
fi
# A dry run makes no branch and no commit, so it reads a dirty tree the
# way a person does: while they work. Every other run commits with
# git add -A, and a stray file would ride into a commit of the cycle.
if [ "$DRY_RUN" != "1" ]; then
  if ! git diff --quiet || ! git diff --cached --quiet; then
    die "the working tree is dirty. Commit or stash first."
  fi
  untracked="$(git ls-files --others --exclude-standard)"
  [ -z "$untracked" ] || die "untracked files would ride into a cycle commit. Commit, stash, or remove: $(tr '\n' ' ' <<<"$untracked")"
fi

START_BRANCH="$(git rev-parse --abbrev-ref HEAD)"
if [ "$DRY_RUN" != "1" ]; then
  case "$START_BRANCH" in
    main|master) die "never commit on $START_BRANCH (D-583). Start a branch first." ;;
  esac
fi
if [ -n "$BASE_REF" ]; then
  git rev-parse --verify --quiet "$BASE_REF" >/dev/null || die "no branch $BASE_REF"
  [ "$DRY_RUN" = "1" ] || git switch "$BASE_REF" >/dev/null 2>&1 || die "could not switch to $BASE_REF"
fi

BRANCH="feedback-fix/$STAMP"
if [ "$DRY_RUN" = "1" ]; then
  say "dry run: no branch, no model call, no commit. The checkout stays on $START_BRANCH."
else
  git switch -c "$BRANCH" >/dev/null 2>&1 || die "could not cut $BRANCH"
  say "branch $BRANCH"
fi
START_COMMIT="$(git rev-parse HEAD)"

# --- The ledger --------------------------------------------------------

# charge adds the cost of one run file to the ledger. A run with no
# priced cost is a fault: an uncounted call would let the cycle pass its
# cap without knowing (T-10).
charge() {
  local file="$1" cost
  cost="$(head -1 "$file" 2>/dev/null | python3 -c 'import json,sys
try:
    h = json.load(sys.stdin)
except Exception:
    sys.exit(1)
c = h.get("cost_usd")
print("" if c is None else "%.6f" % c)' 2>/dev/null)"
  [ -n "$cost" ] || return 1
  echo "$cost" >> "$LEDGER"
  return 0
}

spent() { awk '{s+=$1} END {printf "%.4f", s+0}' "$LEDGER"; }

# under_cap answers whether the ledger still has room. The cycle checks
# it before every paid run, so it never starts one it can not afford.
under_cap() {
  awk -v s="$(spent)" -v c="$CAP" 'BEGIN {exit !(s < c)}'
}

# --- The gates ---------------------------------------------------------

# run_gate runs one gate over the case ids of one suite and writes its
# run file. It calls the Makefile target, so every guard of that target
# holds: the env variable, the overwrite check, and the key loading.
run_gate() {
  local gate="$1" ids="$2" label="$3" doc run
  doc="$STATE_DIR/$gate-$label.md"
  run="$STATE_DIR/$gate-$label.jsonl"
  case "$gate" in
    questions) make questions-gate GATE_OUT="$doc" GATE_RUN="$run" GATE_ARGS="-only $ids" >>"$LOG" 2>&1 ;;
    decks)     make deck-gate DECK_GATE_OUT="$doc" DECK_GATE_RUN="$run" DECK_GATE_ARGS="-only $ids" >>"$LOG" 2>&1 ;;
    brackets)  make bracket-gate BRACKET_GATE_OUT="$doc" BRACKET_GATE_RUN="$run" BRACKET_GATE_ARGS="-only $ids" >>"$LOG" 2>&1 ;;
    *) say "no gate named $gate"; return 2 ;;
  esac
  # A gate exits 1 on a FAIL verdict and writes its document first, which
  # is the expected answer on the confirm run. A missing run file is the
  # real fault.
  [ -f "$run" ] || { say "the $gate gate wrote no run file. Read $LOG"; return 2; }
  charge "$run" || { say "the $gate run is unpriced, so the cap can not count it"; return 2; }
  echo "$run"
  return 0
}

# check_cases reads the run file against the manifest. want is pass or
# fail. It answers 0 when every case reads that way, 1 when one does not,
# and 2 on a fault.
check_cases() {
  local gate="$1" run="$2" want="$3"
  ( cd "$ROOT/go" && go run ./cmd/case-check -manifest "$MANIFEST" -run "$run" -gate "$gate" -want "$want" ) \
    | tee -a "$REPORT" | tee -a "$LOG"
  return "${PIPESTATUS[0]}"
}

# --- Step 1: the cases -------------------------------------------------

say "cap \$$CAP. State under $STATE_DIR."
{
  echo "# The feedback fix cycle, $STAMP"
  echo
  echo "Cap \$$CAP. Branch \`$BRANCH\`. Start commit \`$START_COMMIT\`."
  echo
} > "$REPORT"

triage_args=(-root "$ROOT" -out "$TRIAGE_DOC" -manifest "$MANIFEST")
# The triage runs with the go directory as its working directory, so a
# harvest named on the command line is resolved here.
if [ -n "$HARVEST" ]; then
  case "$HARVEST" in
    /*) ;;
    *) HARVEST="$ROOT/$HARVEST" ;;
  esac
  [ -f "$HARVEST" ] || die "no harvest at $HARVEST"
  triage_args+=(-in "$HARVEST")
fi
if [ "$DRY_RUN" = "1" ]; then
  # A dry cycle writes no case into a gate file. It names what a live one
  # would write, and it leaves the checkout alone.
  triage_args+=(-dry)
  say "step 1: the triage, dry. The reason keys place a case, no judge runs, and no gate file changes."
else
  triage_args+=(-apply)
  say "step 1: the triage. The judge reads the verdicts the reason keys can not place."
fi
if [ "$DRY_RUN" = "1" ]; then
  ( cd "$ROOT/go" && go run ./cmd/feedback-triage "${triage_args[@]}" ) >>"$LOG" 2>&1
else
  ( cd "$ROOT/go" && FEEDBACK_TRIAGE=1 go run ./cmd/feedback-triage "${triage_args[@]}" ) >>"$LOG" 2>&1
fi
triage_code=$?
[ -f "$MANIFEST" ] || die "the triage wrote no manifest. Read $LOG"
if [ "$triage_code" -ne 0 ]; then
  say "WARNING: the triage exited $triage_code. Read $TRIAGE_DOC"
fi

CASE_COUNT="$(python3 -c 'import json,sys; print(len(json.load(open(sys.argv[1]))["cases"]))' "$MANIFEST")"
say "the triage wrote $CASE_COUNT case(s) a gate measures"
if [ "$CASE_COUNT" = "0" ]; then
  say "no case, so there is nothing to fix. The cycle stops."
  exit 0
fi
GATES="$(python3 -c 'import json,sys
m = json.load(open(sys.argv[1]))
seen = []
for c in m["cases"]:
    if c["gate"] not in seen:
        seen.append(c["gate"])
print(" ".join(seen))' "$MANIFEST")"

if [ "$DRY_RUN" != "1" ]; then
  git add -A
  if ! git commit -q -m "The cases of the feedback harvest, $STAMP

The triage of PR-28b read the harvest and wrote one case per thumbs
down. Each one is the fault a reader met, in the shape of the gate that
owns it. The fix follows in this branch (D-645)."; then
    die "could not commit the cases"
  fi
fi
CASES_COMMIT="$(git rev-parse HEAD)"

# --- Step 2: confirm ---------------------------------------------------

say "step 2: confirm. Every case must fail before the fixer touches anything."
{
  echo "## The confirm run"
  echo
  echo "A case that already passes never measured the reader's fault. The cycle keeps it as a case a change must not flip, and the fixer never sees it."
  echo
} >> "$REPORT"

TO_FIX=""
if [ "$DRY_RUN" = "1" ]; then
  say "dry run: no gate ran. A live cycle would run: $GATES"
  for gate in $GATES; do
    ids="$(python3 -c 'import json,sys
m = json.load(open(sys.argv[1]))
print(",".join(str(c["id"]) for c in m["cases"] if c["gate"] == sys.argv[2]))' "$MANIFEST" "$gate")"
    say "  $gate gate, -only $ids"
    echo "- The $gate gate would run over \`-only $ids\`." >> "$REPORT"
  done
  say "dry run: the plan is written and no fixer ran. Read $REPORT"
  exit 0
fi

# A gate the cap stopped leaves cases nobody measured. Those cases are in
# the gate files already, so the cycle must not go on to a pull request
# that carries a case no run ever read.
capped=0
for gate in $GATES; do
  if ! under_cap; then
    say "the cap of \$$CAP is spent before the $gate gate"
    capped=1
    break
  fi
  ids="$(python3 -c 'import json,sys
m = json.load(open(sys.argv[1]))
print(",".join(str(c["id"]) for c in m["cases"] if c["gate"] == sys.argv[2]))' "$MANIFEST" "$gate")"
  say "  the $gate gate over -only $ids"
  echo "### The $gate gate, before the fix" >> "$REPORT"
  run="$(run_gate "$gate" "$ids" confirm)" || die "the $gate gate faulted. Read $LOG"
  check_cases "$gate" "$run" fail
  case $? in
    0) TO_FIX="$TO_FIX $gate" ;;
    1) say "  some cases of the $gate gate already pass. The fixer sees the failures alone." ; TO_FIX="$TO_FIX $gate" ;;
    *) die "case-check faulted on the $gate gate. Read $LOG" ;;
  esac
done
say "spent \$$(spent) of \$$CAP"
if [ "$capped" != "0" ]; then
  say "some cases reached no gate, so the cycle stops before the fixer."
  say "the branch $BRANCH holds the cases alone. Raise --cap and run again, or drop the branch."
  exit 1
fi

# --- Step 3: the fixer -------------------------------------------------

say "step 3: the fixer"
if ! AUTOTUNE_EVAL_DOC="$REPORT" \
     AUTOTUNE_LABEL="feedback-$STAMP" \
     AUTOTUNE_LAST_GOOD="$CASES_COMMIT" \
     AUTOTUNE_FIXER_PROMPT="$ROOT/docs/reference/feedback-fixer-prompt.md" \
     "$ROOT/scripts/autotune-fix.sh" >>"$LOG" 2>&1; then
  say "the fixer failed. Read $LOG"
  git reset -q --hard "$CASES_COMMIT"
  exit 1
fi

# --- Step 4: the free guards -------------------------------------------

say "step 4: the frozen paths and the build"
changed="$(git diff --name-only "$CASES_COMMIT"; git ls-files --others --exclude-standard)"
for path in $FROZEN; do
  if grep -qxF -- "$path" <<<"$changed" || grep -qF -- "$path/" <<<"$changed"; then
    say "the fixer touched a frozen path: $path. The cycle reverts it."
    git reset -q --hard "$CASES_COMMIT"
    exit 1
  fi
done
if git diff -U0 "$CASES_COMMIT" -- docs/decisions.md | grep -q '^-[^-]'; then
  say "the fixer removed a line of docs/decisions.md, which is append-only. The cycle reverts it."
  git reset -q --hard "$CASES_COMMIT"
  exit 1
fi
if ! ( cd "$ROOT/go" && go build ./... && go vet ./... && go test ./... >/dev/null ) \
   || ! ( cd "$ROOT" && make lint-go >/dev/null 2>&1 ); then
  say "the tree is red after the fixer. The cycle reverts it."
  git reset -q --hard "$CASES_COMMIT"
  exit 1
fi
say "  the tree is green"

# --- Step 5: measure ---------------------------------------------------

say "step 5: measure. Every case must pass now."
{
  echo
  echo "## The measure run"
  echo
} >> "$REPORT"
failed=0
for gate in $TO_FIX; do
  under_cap || { say "the cap of \$$CAP is spent before the $gate gate can be measured."; failed=1; break; }
  ids="$(python3 -c 'import json,sys
m = json.load(open(sys.argv[1]))
print(",".join(str(c["id"]) for c in m["cases"] if c["gate"] == sys.argv[2]))' "$MANIFEST" "$gate")"
  echo "### The $gate gate, after the fix" >> "$REPORT"
  run="$(run_gate "$gate" "$ids" measure)" || die "the $gate gate faulted. Read $LOG"
  check_cases "$gate" "$run" pass
  case $? in
    0) say "  every case of the $gate gate passes" ;;
    1) say "  a case of the $gate gate still fails"; failed=1 ;;
    *) die "case-check faulted on the $gate gate. Read $LOG" ;;
  esac
done
say "spent \$$(spent) of \$$CAP"

if [ "$failed" != "0" ]; then
  say "the cycle did not answer every case. The fixer's work is reverted, and the cases stay."
  git reset -q --hard "$CASES_COMMIT"
  say "read $REPORT and $LOG. The branch $BRANCH holds the cases alone."
  exit 1
fi

# --- Step 6: the baselines ---------------------------------------------

say "step 6: eval-check, so no baseline flipped"
if ! make eval-check >>"$LOG" 2>&1; then
  say "a baseline flipped. The fixer traded one reader's complaint for a regression. The cycle reverts it."
  git reset -q --hard "$CASES_COMMIT"
  exit 1
fi

# --- Step 7: the evidence and the pull request -------------------------

mkdir -p "$ROOT/docs/reference/feedback"
EVIDENCE="docs/reference/feedback/fix-cycle-$STAMP.md"
cp "$REPORT" "$ROOT/$EVIDENCE"
git add -A
git commit -q -m "The evidence of the feedback fix cycle, $STAMP

Every case failed before the fix and passes after it. make eval-check
shows no flip on the baselines (D-645)." || die "could not commit the evidence"

say "spent \$$(spent) of \$$CAP over $(wc -l < "$LEDGER" | tr -d ' ') gate run(s)"
if [ "$OPEN_PR" != "1" ]; then
  say "the cycle wrote $BRANCH and pushed nothing. To open it:"
  say "  git push -u origin $BRANCH && gh pr create --fill"
  exit 0
fi

say "step 7: push and open the pull request"
git push -q -u origin "$BRANCH" || die "could not push $BRANCH"
PR_BODY="$STATE_DIR/pr-body.md"
{
  echo "The feedback fix cycle of $STAMP (PR-28c, D-645)."
  echo
  echo "It reads the harvest, writes one case per thumbs down, proves each case fails, fixes the causes, and proves each case passes. The case and its fix are in one pull request, so the gate on \`main\` is never red."
  echo
  echo "Cap \$$CAP. Spent \$$(spent)."
  echo
  sed -n '/^## The measure run/,$p' "$REPORT"
  echo
  echo "The whole evidence is \`$EVIDENCE\`."
} > "$PR_BODY"
gh pr create --title "The feedback fix cycle of $STAMP (PR-28c)" --body-file "$PR_BODY" >>"$LOG" 2>&1 \
  || die "could not open the pull request. Read $LOG"
PR_NUM="$(gh pr view --json number --jq .number 2>/dev/null)"
[ -n "$PR_NUM" ] || die "the pull request opened and its number did not read back"
say "pull request #$PR_NUM"

# --- Step 8: the review ------------------------------------------------

if [ "$ROUNDS" -lt 1 ]; then
  say "no review round was asked for. The cycle ends at #$PR_NUM."
  exit 0
fi
say "step 8: the review of gitar-bot, up to $ROUNDS round(s)"
"$ROOT/scripts/feedback-review.sh" "$PR_NUM" "$ROUNDS" "$STATE_DIR" 2>&1 | tee -a "$LOG"
review_code="${PIPESTATUS[0]}"
if [ "$review_code" -ne 0 ]; then
  say "the review rounds ended with work open on #$PR_NUM. Read the pull request."
  exit "$review_code"
fi
say "the cycle is done. #$PR_NUM holds the cases, the fix, and the answered review."
