#!/usr/bin/env bash
# The review step of the feedback fix cycle (PR-28c, D-645, D-878). It
# reads the review of gitar-bot on one pull request, then runs the Codex
# review, and hands each finding to the fixer agent.
#
#   scripts/feedback-review.sh <pr-number> <rounds> <state-dir>
#
# One round is: wait for the checks of the tip, and read Gitar. An open
# Gitar thread goes to the fixer, who changes the code or says why not,
# and the round replies on the thread (D-637). With no open thread, the
# round runs the Codex review (D-811). A record that asks for changes
# goes to the fixer, who writes the response file (D-878). Each fix runs
# the free checks, and the round pushes it.
#
# During the Gitar pause the round reads Gitar one time and never waits
# for it (D-838). An open thread or an issue on the dashboard stops it.
#
# It answers 0 when Codex approves the effective head, and 1 when a round
# limit or a failure leaves work on the pull request. It answers 3 when
# the Gitar pause stops it on a finding (D-838), and 4 when a Codex
# finding is open at its third head (D-826).
#
# The script never merges, and it never turns on the auto-merge. The
# owner decides the merge after an approval (D-878).
#
# CAUTION: this writes to a public pull request with no person watching.
# The cycle calls it, and the cycle refuses to run without
# FEEDBACK_LOOP_ALLOW=1.
set -uo pipefail
cd "$(dirname "$0")/.." || exit 1
ROOT="$(pwd)"

PR="${1:-}"
ROUNDS="${2:-3}"
STATE_DIR="${3:-$ROOT/.local/tune/review-$(date -u +%Y%m%d-%H%M%S)}"
[ -n "$PR" ] || { echo "feedback-review: give the pull request number" >&2; exit 2; }
mkdir -p "$STATE_DIR"

REPO="$(gh repo view --json nameWithOwner --jq .nameWithOwner 2>/dev/null)"
[ -n "$REPO" ] || { echo "feedback-review: gh does not read this repository" >&2; exit 2; }

say() { printf '%s  review: %s\n' "$(date -u +%H:%M:%S)" "$*" >&2; }

# WAIT_SECONDS is how long one round waits for the review to land. Gitar
# took 6 minutes 48 seconds on pull request #121, so the default gives it
# room and still ends a stuck round.
WAIT_SECONDS="${FEEDBACK_REVIEW_WAIT:-900}"
POLL_SECONDS="${FEEDBACK_REVIEW_POLL:-30}"

# CI_WAIT_SECONDS is how long one round waits for the checks of the tip.
# The default gives the whole verify workflow room.
CI_WAIT_SECONDS="${FEEDBACK_CI_WAIT:-3600}"

# CODEX_REVIEW_CMD runs the Codex review of D-823. A test names a fake.
CODEX_REVIEW_CMD="${CODEX_REVIEW_CMD:-python3 $ROOT/docs/tools/codex_review.py}"

# PAUSE_FILE is the switch of the Gitar pause (D-838). While it exists, a
# round reads Gitar one time and goes on to the Codex review. An open
# thread or an issue on the Gitar dashboard stops the cycle before the
# fixer, and a comment tells the owner, who reads it first.
# GITAR_PAUSE_FILE moves the switch for a test.
PAUSE_FILE="${GITAR_PAUSE_FILE:-$ROOT/docs/reference/gitar-pause.md}"
paused() { [ -f "$PAUSE_FILE" ]; }

# dashboard_clean answers 0 when the newest Gitar dashboard reports no
# issue. The rule is dashboard_issue of docs/tools/codex_review.py, so the
# cycle and make codex-review read a dashboard the same way.
dashboard_clean() {
  gh api --paginate "repos/$REPO/issues/$PR/comments" \
    --jq '.[] | select(.user.login == "gitar-bot[bot]") | .body | tojson' 2>/dev/null \
    | python3 -c 'import json, sys
sys.path.insert(0, "docs/tools")
import codex_review
bodies = [json.loads(line) for line in sys.stdin if line.strip()]
comments = [{"user": {"login": codex_review.GITAR}, "created_at": f"{i:08d}", "body": b} for i, b in enumerate(bodies)]
sys.exit(1 if codex_review.dashboard_issue(comments) else 0)'
}

# pause_stop ends the cycle with 3 when Gitar left a finding during the
# pause: an open thread, or an issue on the dashboard (D-838).
pause_stop() {
  local what=""
  [ "$1" = "0" ] || what="$1 open thread(s)"
  dashboard_clean || what="${what:+$what and }an issue on the dashboard"
  [ -n "$what" ] || return 0
  alert="@${REPO%%/*} Gitar left $what during the Gitar pause (D-838). The cycle stopped before the fixer. The owner reads each finding first."
  gh pr comment "$PR" --repo "$REPO" --body "$alert" >/dev/null 2>&1 </dev/null \
    || say "the comment to the owner failed"
  say "the Gitar pause stops the cycle on $what. Tell the owner (D-838)."
  exit 3
}

# open_threads writes every unresolved thread of the pull request as one
# JSON object a line: the thread id, the path, and the body of the first
# comment. A thread the cycle already answered is resolved by gitar on
# its next pass, so an answered finding leaves this list on its own.
open_threads() {
  # The query names its own variables with a dollar sign, and the shell
  # must leave every one of them alone.
  # shellcheck disable=SC2016
  gh api graphql -f query='
    query($owner:String!,$name:String!,$pr:Int!){
      repository(owner:$owner,name:$name){
        pullRequest(number:$pr){
          reviewThreads(first:50){
            nodes{
              id isResolved isOutdated path line
              comments(first:1){nodes{author{login} body}}
            }
          }
        }
      }
    }' -F owner="${REPO%%/*}" -F name="${REPO##*/}" -F pr="$PR" \
    --jq '.data.repository.pullRequest.reviewThreads.nodes[]
          | select(.isResolved == false)
          | {id: .id, path: .path, line: .line,
             author: .comments.nodes[0].author.login,
             body: .comments.nodes[0].body}' 2>/dev/null
}

# reviewed answers whether gitar has written anything on this commit yet.
# A pull request with no review is not a pull request with no finding.
reviewed() {
  gh pr checks "$PR" --json name,bucket 2>/dev/null \
    | python3 -c 'import json,sys
try:
    rows = json.load(sys.stdin)
except Exception:
    sys.exit(1)
for r in rows:
    if r.get("name") == "Gitar":
        sys.exit(0 if r.get("bucket") != "pending" else 1)
sys.exit(1)'
}

# wait_for_review holds until the Gitar check settles or the wait runs
# out. The round calls it only when the Gitar review is not paused. During
# the pause the round reads Gitar one time and never waits (D-838).
wait_for_review() {
  local waited=0
  while [ "$waited" -lt "$WAIT_SECONDS" ]; do
    if reviewed; then return 0; fi
    sleep "$POLL_SECONDS"
    waited=$((waited + POLL_SECONDS))
  done
  return 1
}

# ci_state prints "done" when each check of the pushed head completed, and
# "fail <names>" when one failed. It prints "pending" while GitHub reads an
# older head or a check still runs. Gitar is not CI, and review-gate waits
# for the Codex record, so neither one counts.
ci_state() {
  local head
  head="$(gh pr view "$PR" --json headRefOid --jq .headRefOid 2>/dev/null)"
  if [ -z "$head" ] || [ "$head" != "$(git rev-parse HEAD 2>/dev/null)" ]; then
    echo "pending"
    return
  fi
  gh pr checks "$PR" --json name,bucket 2>/dev/null \
    | python3 -c 'import json, sys
try:
    rows = json.load(sys.stdin)
except Exception:
    print("pending")
    sys.exit(0)
rows = [r for r in rows if r.get("name") not in ("Gitar", "review-gate")]
if not rows or any(r.get("bucket") == "pending" for r in rows):
    print("pending")
    sys.exit(0)
bad = [r["name"] for r in rows if r.get("bucket") in ("fail", "cancel")]
print("fail " + ", ".join(bad) if bad else "done")'
}

# wait_for_ci holds until each check of the tip completes. The Codex
# review starts after CI, as the pr-review skill says (D-878).
wait_for_ci() {
  local waited=0 state
  while :; do
    state="$(ci_state)"
    case "$state" in
      done) return 0 ;;
      fail*) say "a check of the tip failed: ${state#fail }. A person reads it."; return 1 ;;
    esac
    if [ "$waited" -ge "$CI_WAIT_SECONDS" ]; then
      say "the checks of the tip did not complete inside $CI_WAIT_SECONDS seconds"
      return 1
    fi
    sleep "$POLL_SECONDS"
    waited=$((waited + POLL_SECONDS))
  done
}

# frozen_paths prints the FROZEN block of the cycle. The fixer of a review
# round keeps the same rule as the fixer of the cycle (D-645), and it never
# edits the review record, which belongs to the reviewer (D-878).
frozen_paths() {
  awk '/^FROZEN="$/ {f = 1; next} f && /^"$/ {exit} f && NF {print $1}' "$ROOT/scripts/feedback-loop.sh"
  echo "docs/reviews/pr-$PR.md"
}

# fix_and_push hands one findings document to the fixer. It reverts the
# round when the fixer touches a frozen path, removes a decision line, or
# turns the tree red. It pushes what is left. It answers 1 on a failure.
fix_and_push() {
  local doc="$1" label="$2" summary="$3" before after changed path
  before="$(git rev-parse HEAD)"
  if ! AUTOTUNE_EVAL_DOC="$doc" \
       AUTOTUNE_LABEL="$label" \
       AUTOTUNE_LAST_GOOD="$before" \
       AUTOTUNE_FIXER_PROMPT="$ROOT/docs/reference/feedback-fixer-prompt.md" \
       "$ROOT/scripts/autotune-fix.sh" > "$summary" 2>&1; then
    say "the fixer failed on $label. Read $summary"
    git reset -q --hard "$before"
    return 1
  fi
  after="$(git rev-parse HEAD)"
  if [ "$before" = "$after" ]; then
    say "the fixer changed nothing on $label"
    return 0
  fi
  changed="$(git diff --name-only "$before"; git ls-files --others --exclude-standard)"
  for path in $(frozen_paths); do
    if grep -qxF -- "$path" <<<"$changed" || grep -qF -- "$path/" <<<"$changed"; then
      say "the fixer touched a frozen path: $path. The cycle reverts $label."
      git reset -q --hard "$before"
      return 1
    fi
  done
  if git diff -U0 "$before" -- docs/decisions.md | grep -q '^-[^-]'; then
    say "the fixer removed a line of docs/decisions.md, which is append-only. The cycle reverts $label."
    git reset -q --hard "$before"
    return 1
  fi
  if ! ( cd "$ROOT/go" && go build ./... && go vet ./... && go test ./... >/dev/null ) \
     || ! ( cd "$ROOT" && make lint-go >/dev/null 2>&1 ); then
    say "the tree is red after the fix of $label. The cycle reverts it."
    git reset -q --hard "$before"
    return 1
  fi
  git push -q || { say "could not push the fix of $label"; return 1; }
  say "pushed $(git rev-list --count "$before".."$after") commit(s)"
}

# gitar_round hands the open Gitar threads to the fixer, and replies on
# each thread with what the fixer said (D-637). Only a round with no
# Gitar pause reaches it.
gitar_round() {
  local threads="$1" count="$2"
  FINDINGS="$STATE_DIR/findings-$round.md"
  {
    echo "# The review of pull request #$PR, round $round"
    echo
    echo "\`gitar-bot\` reviewed this pull request. Every finding below is open. Read each one on its merit and never on its tone."
    echo
    echo "For each finding, answer in your summary with the thread id and one of two words."
    echo
    echo "- \`change\`: you made a change, a test, and a commit. Say what you changed."
    echo "- \`no-merit\`: you made no change. Say why in one paragraph."
    echo
    echo "## The findings"
    echo
    # The script is python and not shell, so no name in it expands here.
    # shellcheck disable=SC2016
    python3 -c '
import json, sys
for line in open(sys.argv[1]):
    line = line.strip()
    if not line:
        continue
    t = json.loads(line)
    print("### Thread `%s`" % t["id"])
    print()
    print("File `%s`, line %s, by %s." % (t.get("path") or "the pull request", t.get("line") or "none", t.get("author") or "unknown"))
    print()
    print(t.get("body") or "")
    print()
' "$threads"
    echo "## The change under review"
    echo
    echo '```diff'
    git diff "$(git merge-base HEAD origin/main)"...HEAD | head -c 200000
    echo '```'
  } > "$FINDINGS"

  say "handing $count Gitar finding(s) to the fixer"
  SUMMARY="$STATE_DIR/summary-$round.md"
  fix_and_push "$FINDINGS" "review-$PR-$round" "$SUMMARY" || exit 1

  # One reply per thread, with what the fixer said about it. A thread the
  # summary does not name gets the general answer, so no finding is left
  # without a word (D-637).
  while read -r line; do
    [ -n "$line" ] || continue
    tid="$(python3 -c 'import json,sys; print(json.loads(sys.argv[1])["id"])' "$line")"
    reply="$(python3 -c '
import re, sys
tid, path = sys.argv[1], sys.argv[2]
text = open(path, encoding="utf-8", errors="replace").read()
# The fixer answers with the thread id and then its words. Take from the
# id to the next id or the end.
m = re.search(re.escape(tid) + r"(.{0,4000}?)(?=\n#|\Z)", text, re.S)
print((m.group(1).strip() if m else "")[:3500])
' "$tid" "$SUMMARY")"
    if [ -z "$reply" ]; then
      reply="The fixer read this finding and its summary named no answer for it. A person reads it next."
    fi
    # shellcheck disable=SC2016
    gh api graphql -f query='
      mutation($tid:ID!,$body:String!){
        addPullRequestReviewThreadReply(input:{pullRequestReviewThreadId:$tid, body:$body}){
          comment{ id }
        }
      }' -F tid="$tid" -F body="$reply" >/dev/null 2>&1 </dev/null \
      || say "could not reply on thread $tid"
  done < "$threads"
  say "replied to $count thread(s)"
}

# codex_round runs the Codex review of the pushed head (D-811, D-878). An
# approval ends the review step with 0, and the owner then decides the
# merge. A record that asks for changes goes to the fixer, who answers each
# finding in the response file. The cycle never merges (D-878).
codex_round() {
  local log="$STATE_DIR/codex-$round.log" code args
  args=(--pr "$PR")
  if paused; then args+=(--skip-gitar-review); fi
  say "round $round: the Codex review (D-878)"
  # CODEX_REVIEW_CMD holds a command and its arguments, so it splits here.
  # shellcheck disable=SC2086
  $CODEX_REVIEW_CMD "${args[@]}" > "$log" 2>&1 </dev/null
  code=$?
  tail -n 3 "$log" >&2
  case "$code" in
    0|3|4) git pull -q --ff-only || { say "could not pull the review record"; exit 1; } ;;
  esac
  case "$code" in
    0)
      say "Codex approved the effective head. The owner decides the merge, and the cycle never merges (D-878)."
      exit 0 ;;
    3) ;;
    4)
      say "a Codex finding is open at its third head (D-826). The cycle stops, and the owner decides."
      exit 4 ;;
    *)
      say "the Codex review ended with exit $code. Read $log"
      exit 1 ;;
  esac

  FINDINGS="$STATE_DIR/codex-findings-$round.md"
  {
    echo "# The Codex review of pull request #$PR, round $round"
    echo
    echo "Codex reviewed this pull request, and its record asks for changes. A finding is a claim, not a fact."
    echo "Answer each open finding with the steps of \`.claude/skills/pr-review/references/answer-review.md\`, under \"Answer the findings of a review\"."
    echo "The cycle does the push and the next review, so skip steps 8 and 9 of that list."
    echo
    echo "- Decide full merit, partial merit, or no merit for each open finding."
    echo "- Correct each part with merit, with a test, in a commit."
    echo "- Write \`docs/reviews/pr-$PR-response.md\` as the reference file says, and commit it."
    echo "- Never edit \`docs/reviews/pr-$PR.md\`. The record belongs to the reviewer, and the cycle reverts a round that edits it."
    echo
    echo "## The record"
    echo
    cat "$ROOT/docs/reviews/pr-$PR.md" 2>/dev/null || echo "The record did not read. Answer nothing, and say so in your summary."
    echo
    echo "## The change under review"
    echo
    echo '```diff'
    git diff "$(git merge-base HEAD origin/main)"...HEAD | head -c 200000
    echo '```'
  } > "$FINDINGS"
  say "handing the Codex findings to the fixer"
  fix_and_push "$FINDINGS" "codex-$PR-$round" "$STATE_DIR/codex-summary-$round.md" || exit 1
}

round=1
while [ "$round" -le "$ROUNDS" ]; do
  say "round $round of $ROUNDS: the checks of the tip"
  wait_for_ci || exit 1
  if paused; then
    say "round $round: one read of Gitar, because the Gitar review is paused (D-838)"
  else
    say "round $round: waiting for the review of gitar-bot"
    if ! wait_for_review; then
      say "no review inside $WAIT_SECONDS seconds. The pull request keeps whatever is open."
      exit 1
    fi
  fi
  threads="$STATE_DIR/threads-$round.jsonl"
  open_threads > "$threads"
  # grep -c prints 0 and exits 1 on an empty file, so an "|| echo 0"
  # here would write a second line and the count would never read 0.
  count="$(awk 'NF' "$threads" 2>/dev/null | wc -l | tr -d ' ')"
  if paused; then
    pause_stop "$count"
  fi
  if [ "$count" != "0" ]; then
    say "$count open thread(s)"
    gitar_round "$threads" "$count"
  else
    say "Gitar holds nothing open"
    codex_round
  fi
  round=$((round + 1))
done

say "the review still holds open work after $ROUNDS round(s)"
exit 1
