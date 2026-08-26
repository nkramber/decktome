#!/usr/bin/env bash
# The automated tuning loop (D-133). It runs the question gate, scores
# every question with the eval role, hands the report to a fixer agent,
# and keeps the change only when the run got better.
#
# CAUTION: every iteration calls real providers and costs money. The loop
# refuses to start without AUTOTUNE_ALLOW_UNATTENDED=1, and it stops at
# the budget.
#
# The loop never runs on a branch it did not make. It reverts an
# iteration that fails the build, that touches a frozen file, or that the
# checker rejects. Nothing reaches main.
#
# The loop commits to its own branch and does not push. Add --push when
# the owner wants the branch on the remote as it goes.
#
#   AUTOTUNE_ALLOW_UNATTENDED=1 scripts/autotune.sh --budget 3.00 --max 20
set -u
cd "$(dirname "$0")/.." || exit 1
ROOT="$(pwd)"

BUDGET="3.00"
MAX_ITERATIONS="20"
TARGET_RATIO="0.05"
EVAL_BUDGET="0.50"
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
    --branch) BRANCH_PREFIX="$2"; shift 2 ;;
    --push) PUSH="1"; shift ;;
    --no-push) PUSH="0"; shift ;;
    --dry-run) DRY_RUN="1"; shift ;;
    *) echo "unknown flag: $1" >&2; exit 2 ;;
  esac
done

STATE_DIR="$ROOT/.local/tune"
LEDGER="$STATE_DIR/ledger.txt"
LOG="$STATE_DIR/autotune.log"
mkdir -p "$STATE_DIR"

say() { printf '%s  %s\n' "$(date -u +%H:%M:%S)" "$*" | tee -a "$LOG"; }
die() { say "STOP: $*"; exit 1; }

# --- Frozen files. The fixer may not change how it is measured. --------
# A loop that may edit its own scorer optimizes the scorer. Every path
# here is part of the measurement, not part of the product.
FROZEN="
go/cmd/questions-eval
go/cmd/tune-check
go/internal/tune
go/internal/questions/lint.go
go/internal/questions/lint_test.go
go/cmd/questions-gate/conversations.json
scripts/autotune.sh
scripts/autotune-fix.sh
docs/owner-questions.md
"

frozen_touched() {
  local changed path
  changed="$(git diff --name-only HEAD; git ls-files --others --exclude-standard)"
  for path in $FROZEN; do
    if printf '%s\n' "$changed" | grep -qx -- "$path" || \
       printf '%s\n' "$changed" | grep -q "^$path/"; then
      echo "$path"
      return 0
    fi
  done
  # The M-5 sheets are the owner's record. Nothing may rewrite one.
  if printf '%s\n' "$changed" | grep -q '^docs/reference/pr7-m5-scoring'; then
    echo "docs/reference/pr7-m5-scoring*"
    return 0
  fi
  # docs/decisions.md is append-only: a removed line rewrites history.
  if git diff -U0 -- docs/decisions.md | grep -q '^-[^-]'; then
    echo "docs/decisions.md (a line was removed)"
    return 0
  fi
  return 1
}

# --- Budget ------------------------------------------------------------
spent() { [ -f "$LEDGER" ] && awk '{s+=$1} END {printf "%.4f", s+0}' "$LEDGER" || echo "0.0000"; }
charge() { printf '%s %s\n' "$1" "$2" >> "$LEDGER"; }
over_budget() { awk -v s="$(spent)" -v b="$BUDGET" 'BEGIN {exit !(s >= b)}'; }

# cost_of reads the "- Cost: $0.1234." line a run document writes.
cost_of() { grep -o -- '- Cost: \$[0-9.]*' "$1" 2>/dev/null | tail -1 | tr -d '$' | awk '{print $3+0}'; }

green() {
  ( cd "$ROOT/go" && go build ./... && go vet ./... && go test ./... >/dev/null ) \
    && ( cd "$ROOT" && make lint-go >/dev/null 2>&1 )
}

# --- Preflight ---------------------------------------------------------
[ "${AUTOTUNE_ALLOW_UNATTENDED:-0}" = "1" ] || \
  die "this loop edits code and pushes with no person watching. Set AUTOTUNE_ALLOW_UNATTENDED=1 to allow it."
[ -f "$ROOT/.env" ] || die "no .env, so no API keys"
git diff --quiet && git diff --cached --quiet || die "the working tree is dirty. Commit or stash first."

BASE_BRANCH="$(git rev-parse --abbrev-ref HEAD)"
case "$BASE_BRANCH" in
  main|master) die "refusing to run from $BASE_BRANCH" ;;
esac
STAMP="$(date -u +%Y%m%d-%H%M%S)"
BRANCH="$BRANCH_PREFIX/$STAMP"
say "base branch $BASE_BRANCH, working branch $BRANCH"
[ "$DRY_RUN" = "1" ] || git switch -c "$BRANCH" >/dev/null 2>&1 || die "could not make branch $BRANCH"

# shellcheck disable=SC1091
set -a; . "$ROOT/.env"; set +a
export CARDS_SNAPSHOT_DIR="${CARDS_SNAPSHOT_DIR:-$ROOT/.local/gcs/mtg-local-cards/scryfall}"

run_gate() {   # $1 = iteration label
  local out="$ROOT/docs/reference/pr7-question-gate-$1.md"
  say "gate $1: 66 conversations, about 14 minutes"
  [ "$DRY_RUN" = "1" ] && { say "dry run: no gate"; return 1; }
  ( cd "$ROOT/go" && QUESTIONS_GATE=1 go run ./cmd/questions-gate \
      -collection internal/collections/testdata/manabox_collection.csv > "$out" ) || true
  [ -s "$out" ] || return 1
  charge "$(cost_of "$out")" "gate-$1"
  echo "$out"
}

run_eval() {   # $1 = iteration label, $2 = gate document
  local doc="$ROOT/docs/reference/pr7-question-eval-$1.md"
  local sum="$STATE_DIR/$1.json"
  say "eval $1: scoring every question"
  ( cd "$ROOT/go" && QUESTIONS_EVAL=1 go run ./cmd/questions-eval \
      -in "$2" -out "$doc" -json "$sum" -budget "$EVAL_BUDGET" ) || true
  [ -s "$sum" ] || return 1
  charge "$(cost_of "$doc")" "eval-$1"
  echo "$sum"
}

commit_push() {  # $1 = version, $2 = subject
  git add -A
  git diff --cached --quiet && { say "nothing to commit"; return 0; }
  git commit -q -F - <<EOF
$1: $2

Bad-question ratio $(ratio_of "$3")%. Questions asked $(metric_of "$3" questions).
Automated tuning loop, budget spent \$$(spent) of \$$BUDGET.
EOF
  [ "$PUSH" = "1" ] && git push -q -u origin "$BRANCH" 2>/dev/null
  say "committed $1"
}

ratio_of()  { python3 -c "import json,sys;print(round(json.load(open(sys.argv[1]))['bad_ratio']*100,1))" "$1" 2>/dev/null || echo "?"; }
metric_of() { python3 -c "import json,sys;print(json.load(open(sys.argv[1]))['metrics'][sys.argv[2]])" "$1" "$2" 2>/dev/null || echo "?"; }

# --- Baseline ----------------------------------------------------------
say "budget \$$BUDGET, target ratio $TARGET_RATIO, at most $MAX_ITERATIONS iterations"
GATE="$(run_gate "auto-000")" || die "the baseline gate produced nothing"
PREV="$(run_eval "auto-000" "$GATE")" || die "the baseline eval produced nothing"
say "baseline ratio $(ratio_of "$PREV")%, spent \$$(spent)"
commit_push "v0.0" "baseline run for the tuning loop" "$PREV"
LAST_GOOD="$(git rev-parse HEAD)"

# --- Loop --------------------------------------------------------------
i=0
rejects=0
while [ "$i" -lt "$MAX_ITERATIONS" ]; do
  i=$((i+1))
  LABEL="$(printf 'auto-%03d' "$i")"
  if over_budget; then say "the budget of \$$BUDGET is spent"; break; fi
  if [ "$rejects" -ge 3 ]; then say "three iterations in a row changed nothing that held"; break; fi

  say "--- iteration $i, spent \$$(spent) of \$$BUDGET"
  EVAL_DOC="$ROOT/docs/reference/pr7-question-eval-$(printf 'auto-%03d' $((i-1))).md"
  [ -f "$EVAL_DOC" ] || EVAL_DOC="$ROOT/docs/reference/pr7-question-eval-auto-000.md"

  if ! AUTOTUNE_EVAL_DOC="$EVAL_DOC" AUTOTUNE_LABEL="$LABEL" "$ROOT/scripts/autotune-fix.sh"; then
    say "the fixer failed"; git checkout -- . ; rejects=$((rejects+1)); continue
  fi
  if git diff --quiet HEAD && [ -z "$(git ls-files --others --exclude-standard)" ]; then
    say "the fixer changed nothing"; rejects=$((rejects+1)); continue
  fi
  if path="$(frozen_touched)"; then
    say "REVERT: the fixer touched a frozen path: $path"
    git reset -q --hard "$LAST_GOOD"; git clean -qfd; rejects=$((rejects+1)); continue
  fi
  if ! green; then
    say "REVERT: the tree is not green"
    git reset -q --hard "$LAST_GOOD"; git clean -qfd; rejects=$((rejects+1)); continue
  fi

  GATE="$(run_gate "$LABEL")" || { say "REVERT: the gate produced nothing"; git reset -q --hard "$LAST_GOOD"; git clean -qfd; rejects=$((rejects+1)); continue; }
  NEXT="$(run_eval "$LABEL" "$GATE")" || { say "REVERT: the eval produced nothing"; git reset -q --hard "$LAST_GOOD"; git clean -qfd; rejects=$((rejects+1)); continue; }

  ( cd "$ROOT/go" && go run ./cmd/tune-check -next "$NEXT" -prev "$PREV" -target "$TARGET_RATIO" ) | tee -a "$LOG"
  code="${PIPESTATUS[0]}"
  case "$code" in
    0) commit_push "v0.$i" "tuning iteration $i" "$NEXT"; LAST_GOOD="$(git rev-parse HEAD)"; PREV="$NEXT"; rejects=0 ;;
    3) commit_push "v0.$i" "tuning iteration $i, target reached" "$NEXT"; say "target reached"; break ;;
    *) say "REVERT: the checker rejected iteration $i"
       git reset -q --hard "$LAST_GOOD"; git clean -qfd; rejects=$((rejects+1)) ;;
  esac
done

say "done. Spent \$$(spent) of \$$BUDGET over $i iterations. Branch $BRANCH."
say "read $LOG and the eval documents in docs/reference/."
