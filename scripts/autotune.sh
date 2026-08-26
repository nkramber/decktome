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
set -u
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
# Measured on 2026-08-26: three between two full runs, and two between two
# scorings of one document (D-183).
NOISE="3"
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
# edits its own memory forgets what it was told (D-182).
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
docs/reference/autotune-fixer-prompt.md
docs/reference/autotune-lessons.md
"

# frozen_touched reads every change since the last good commit: the
# fixer's own commits and whatever it left in the tree.
frozen_touched() {
  local changed path
  changed="$(git diff --name-only "$LAST_GOOD"; git ls-files --others --exclude-standard)"
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
  if git diff -U0 "$LAST_GOOD" -- docs/decisions.md | grep -q '^-[^-]'; then
    echo "docs/decisions.md (a line was removed)"
    return 0
  fi
  return 1
}

# --- Budget ------------------------------------------------------------
spent() { [ -f "$LEDGER" ] && awk '{s+=$1} END {printf "%.4f", s+0}' "$LEDGER" || echo "0.0000"; }
charge() { printf '%s %s\n' "$1" "$2" >> "$LEDGER"; }
over_budget() { awk -v s="$(spent)" -v b="$BUDGET" 'BEGIN {exit !(s >= b)}'; }

# cost_of reads the "- Cost: $0.1234." line a run document writes. A
# document with no priced cost is a run on a model with no price row, and
# a loop that reads that as zero has no budget at all (audit H-3).
cost_of() {
  local c
  c="$(grep -o -- '- Cost: \$[0-9.]*' "$1" 2>/dev/null | tail -1 | tr -d '$' | awk '{print $3+0}')"
  [ -n "$c" ] || die "$1 carries no priced cost, so the budget can not count it. Add the model to prices.json."
  echo "$c"
}

green() {
  ( cd "$ROOT/go" && go build ./... && go vet ./... && go test ./... >/dev/null ) \
    && ( cd "$ROOT" && make lint-go >/dev/null 2>&1 )
}

# --- Preflight ---------------------------------------------------------
[ "${AUTOTUNE_ALLOW_UNATTENDED:-0}" = "1" ] || \
  die "this loop edits code and pushes with no person watching. Set AUTOTUNE_ALLOW_UNATTENDED=1 to allow it."
[ -f "$ROOT/.env" ] || die "no .env, so no API keys"
git diff --quiet && git diff --cached --quiet || die "the working tree is dirty. Commit or stash first."
# An untracked file would ride into a loop commit under the owner's name,
# because the loop commits with git add -A (audit H-1).
untracked="$(git ls-files --others --exclude-standard)"
[ -z "$untracked" ] || die "untracked files would ride into a loop commit. Commit, stash, or remove: $(printf '%s ' $untracked)"

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

# shellcheck disable=SC1091
set -a; . "$ROOT/.env"; set +a
export CARDS_SNAPSHOT_DIR="${CARDS_SNAPSHOT_DIR:-$ROOT/.local/gcs/mtg-local-cards/scryfall}"

# conv_count reads the size of the conversation set. A hardcoded number
# goes stale: the set grew from 52 to 66 to 100 to 104, and the log still
# said 66 (D-145, D-155).
conv_count() {
  python3 -c "import json,sys;d=json.load(open(sys.argv[1]));print(len(d if isinstance(d,list) else d['conversations']))" \
    "$ROOT/go/cmd/questions-gate/conversations.json" 2>/dev/null || echo "?"
}

run_gate() {   # $1 = iteration label, $2 = conversation ids to run (optional, comma separated)
  local out="$ROOT/docs/reference/pr7-question-gate-$1.md" only="${2:-}" scope
  # A scored document is never overwritten (D-65, D-177). This runs
  # inside a command substitution, so it returns and never exits: die
  # would kill the subshell alone and hide the reason.
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
  charge "$(cost_of "$out")" "gate-$1"
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
  charge "$(cost_of "$doc")" "eval-$1"
  echo "$sum"
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
  local body="Automated tuning loop, budget spent \$$(spent) of \$$BUDGET."
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

# strip_attribution removes AI-attribution lines from the fixer's commit
# messages. The house rule forbids them in any commit (CLAUDE.md rule 6),
# and a fixer tool may add one by default.
strip_attribution() {
  git log --format=%B "$LAST_GOOD..HEAD" | grep -qiE 'co-authored-by|generated with|claude code' || return 0
  say "a fixer commit carried attribution text, which the house rule forbids: removed"
  FILTER_BRANCH_SQUELCH_WARNING=1 git filter-branch -f \
    --msg-filter "grep -viE 'co-authored-by|generated with|claude code' || true" \
    "$LAST_GOOD..HEAD" >/dev/null 2>&1 || die "could not rewrite the commit messages"
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
  for f in "$ROOT"/docs/reference/pr7-question-gate-"$1"*.md "$ROOT"/docs/reference/pr7-question-eval-"$1"*.md "$STATE_DIR/$1"*.json; do
    [ -f "$f" ] && cp "$f" "$dir/"
  done
  printf '%s\n' "$2" > "$dir/why.txt"
  say "kept the evidence of $1 under $dir"
}

revert_to_last_good() { git reset -q --hard "$LAST_GOOD"; git clean -qfd; }

# record_lesson appends what the checker learned to the lessons file and
# commits it, so a rejected iteration still leaves its lesson on the
# branch. The checker wrote the lesson to a scratch file, because the
# revert wipes the tree (D-182).
record_lesson() {  # $1 = label, $2 = outcome word
  local scratch="$STATE_DIR/$1-lesson.md"
  [ -s "$scratch" ] || return 0
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

if [ -n "$BASELINE_JSON" ]; then
  [ -s "$BASELINE_JSON" ] || die "no readable baseline at $BASELINE_JSON"
  PREV="$BASELINE_JSON"
  # The fixer needs the report, not the summary. Take the name the owner
  # gave, or the one that matches the summary: .local/tune/run18.json goes
  # with docs/reference/pr7-question-eval-run18.md.
  if [ -z "$BASELINE_DOC" ]; then
    BASELINE_DOC="$ROOT/docs/reference/pr7-question-eval-$(basename "$BASELINE_JSON" .json).md"
  fi
  [ -f "$BASELINE_DOC" ] || die "no eval report at $BASELINE_DOC. Name it with --baseline-doc."
  PREV_DOC="$BASELINE_DOC"
  # A summary from before D-181 carries no per-conversation counts, and a
  # partial re-measure needs them. The gate document supplies them.
  PREV_GATE="$ROOT/docs/reference/pr7-question-gate-$(basename "$BASELINE_JSON" .json).md"
  say "baseline read from $BASELINE_JSON, ratio $(ratio_of "$PREV")%"
  say "baseline report $PREV_DOC"
else
  GATE="$(run_gate "$STAMP-000")" || die "the baseline gate produced nothing"
  PREV="$(run_eval "$STAMP-000" "$GATE")" || die "the baseline eval produced nothing"
  PREV_DOC="$ROOT/docs/reference/pr7-question-eval-$STAMP-000.md"
  PREV_GATE="$GATE"
  say "baseline ratio $(ratio_of "$PREV")%, spent \$$(spent)"
  commit_all "v0.0" "baseline run for the tuning loop" "$PREV"
fi
LAST_GOOD="$(git rev-parse HEAD)"
[ -f "$LESSONS" ] || { printf '# What the tuning loop learned\n\nThe loop appends one block per iteration (D-182).\n' > "$LESSONS"; commit_all "lessons" "start the lessons file"; LAST_GOOD="$(git rev-parse HEAD)"; }

# --- Loop --------------------------------------------------------------
i=0
rejects=0
while [ "$i" -lt "$MAX_ITERATIONS" ]; do
  i=$((i+1))
  # The label carries the run stamp, so no two runs ever write the same
  # document (D-177).
  LABEL="$(printf '%s-%03d' "$STAMP" "$i")"
  if [ -f "$STOP_FILE" ]; then
    say "stop file found, so the loop ends after $((i-1)) iterations"
    rm -f "$STOP_FILE"
    break
  fi
  if over_budget; then say "the budget of \$$BUDGET is spent"; break; fi
  if [ "$rejects" -ge 3 ]; then say "three iterations in a row changed nothing that held"; break; fi

  say "--- iteration $i, spent \$$(spent) of \$$BUDGET"

  if ! AUTOTUNE_EVAL_DOC="$PREV_DOC" AUTOTUNE_LABEL="$LABEL" AUTOTUNE_LESSONS="$LESSONS" \
       AUTOTUNE_LAST_GOOD="$LAST_GOOD" "$ROOT/scripts/autotune-fix.sh"; then
    say "the fixer failed"; revert_to_last_good; rejects=$((rejects+1)); continue
  fi
  # Edits the fixer left uncommitted become one change with no rows. The
  # checker can charge nothing to it, so the whole-run rules decide it.
  if ! git diff --quiet HEAD || [ -n "$(git ls-files --others --exclude-standard)" ]; then
    commit_all "change" "unsplit edits of iteration $i"
  fi
  if [ "$(git rev-parse HEAD)" = "$LAST_GOOD" ]; then
    say "the fixer changed nothing"; rejects=$((rejects+1)); continue
  fi
  if path="$(frozen_touched)"; then
    say "REVERT: the fixer touched a frozen path: $path"
    preserve "$LABEL" "touched a frozen path: $path"
    revert_to_last_good; rejects=$((rejects+1)); continue
  fi
  strip_attribution
  CHANGES="$STATE_DIR/$LABEL-changes.json"
  say "the fixer made $(changes_json "$CHANGES") changes"
  if ! green; then
    say "REVERT: the tree is not green"
    preserve "$LABEL" "the tree was not green"
    revert_to_last_good; rejects=$((rejects+1)); continue
  fi

  GATE="$(run_gate "$LABEL")" || { say "REVERT: the gate produced nothing"; preserve "$LABEL" "the gate produced nothing"; revert_to_last_good; rejects=$((rejects+1)); continue; }
  NEXT="$(run_eval "$LABEL" "$GATE")" || { say "REVERT: the eval produced nothing"; preserve "$LABEL" "the eval produced nothing"; revert_to_last_good; rejects=$((rejects+1)); continue; }

  VERDICT="$STATE_DIR/$LABEL-verdict.txt"
  rm -f "$STATE_DIR/$LABEL-lesson.md"
  ( cd "$ROOT/go" && go run ./cmd/tune-check -next "$NEXT" -prev "$PREV" -target "$TARGET_RATIO" \
      -changes "$CHANGES" -noise "$NOISE" -lessons "$STATE_DIR/$LABEL-lesson.md" -label "$LABEL" ) | tee -a "$LOG" > "$VERDICT"
  code="${PIPESTATUS[0]}"
  case "$code" in
    0|3)
      commit_all "v0.$i" "tuning iteration $i" "$NEXT"
      LAST_GOOD="$(git rev-parse HEAD)"
      record_lesson "$LABEL" "accepted"
      PREV="$NEXT"; PREV_DOC="$ROOT/docs/reference/pr7-question-eval-$LABEL.md"; PREV_GATE="$GATE"; rejects=0
      [ "$code" = "3" ] && { say "target reached"; break; }
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
        say "REVERT: the kept changes do not apply without the dropped ones"
        revert_to_last_good; record_lesson "$LABEL" "rejected"; rejects=$((rejects+1)); continue
      fi
      if ! green; then
        say "REVERT: the kept changes alone are not green"
        revert_to_last_good; record_lesson "$LABEL" "rejected"; rejects=$((rejects+1)); continue
      fi
      RLABEL="$LABEL-kept"
      RGATE="$(run_gate "$RLABEL" "$only")" || { say "REVERT: the re-measure gate produced nothing"; revert_to_last_good; record_lesson "$LABEL" "rejected"; rejects=$((rejects+1)); continue; }
      RNEXT="$(run_eval "$RLABEL" "$RGATE")" || { say "REVERT: the re-measure eval produced nothing"; revert_to_last_good; record_lesson "$LABEL" "rejected"; rejects=$((rejects+1)); continue; }
      MERGED="$STATE_DIR/$RLABEL-merged.json"
      MERGED_DOC="$ROOT/docs/reference/pr7-question-eval-$RLABEL-merged.md"
      ( cd "$ROOT/go" && go run ./cmd/tune-check -merge -prev "$PREV" -prev-doc "$PREV_GATE" \
          -next "$RNEXT" -out "$MERGED" -out-doc "$MERGED_DOC" ) | tee -a "$LOG" >&2 \
        || { say "REVERT: the merge failed"; revert_to_last_good; record_lesson "$LABEL" "rejected"; rejects=$((rejects+1)); continue; }
      ( cd "$ROOT/go" && go run ./cmd/tune-check -next "$MERGED" -prev "$PREV" -target "$TARGET_RATIO" -noise "$NOISE" ) | tee -a "$LOG" >&2
      rcode="${PIPESTATUS[0]}"
      case "$rcode" in
        0|3)
          commit_all "v0.$i" "tuning iteration $i, kept changes alone" "$MERGED"
          LAST_GOOD="$(git rev-parse HEAD)"
          record_lesson "$LABEL" "partly kept"
          PREV="$MERGED"; PREV_DOC="$MERGED_DOC"; PREV_GATE=""; rejects=0
          [ "$rcode" = "3" ] && { say "target reached"; break; }
          ;;
        *)
          say "REVERT: the kept changes alone did not hold"
          preserve "$RLABEL" "the kept changes alone did not hold"
          revert_to_last_good; record_lesson "$LABEL" "rejected"; rejects=$((rejects+1))
          ;;
      esac
      ;;
    *)
      say "REVERT: the checker rejected iteration $i"
      preserve "$LABEL" "$(grep '^reject: ' "$VERDICT" | head -3)"
      revert_to_last_good
      record_lesson "$LABEL" "rejected"
      rejects=$((rejects+1))
      ;;
  esac
done

say "done. Spent \$$(spent) of \$$BUDGET over $i iterations. Branch $BRANCH."
say "read $LOG, $LESSONS, and the eval documents in docs/reference/."
