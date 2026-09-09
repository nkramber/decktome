#!/usr/bin/env bash
# The review round of the feedback fix cycle (PR-28c, D-645). It waits
# for the review of gitar-bot on one pull request, hands each finding to
# the fixer agent, and replies on the thread with what changed (D-637).
#
#   scripts/feedback-review.sh <pr-number> <rounds> <state-dir>
#
# One round is: wait for the review, ask the fixer, run the free checks,
# commit, push, and reply to every thread. The round ends when the review
# holds no open thread.
#
# It answers 0 when the review holds nothing open, and 1 when a round
# limit or a failure leaves work on the pull request.
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
# out. A round that never sees a review reports it and stops.
wait_for_review() {
  local waited=0
  while [ "$waited" -lt "$WAIT_SECONDS" ]; do
    if reviewed; then return 0; fi
    sleep "$POLL_SECONDS"
    waited=$((waited + POLL_SECONDS))
  done
  return 1
}

round=1
while [ "$round" -le "$ROUNDS" ]; do
  say "round $round of $ROUNDS: waiting for the review of gitar-bot"
  if ! wait_for_review; then
    say "no review inside $WAIT_SECONDS seconds. The pull request keeps whatever is open."
    exit 1
  fi
  threads="$STATE_DIR/threads-$round.jsonl"
  open_threads > "$threads"
  # grep -c prints 0 and exits 1 on an empty file, so an "|| echo 0"
  # here would write a second line and the count would never read 0.
  count="$(awk 'NF' "$threads" 2>/dev/null | wc -l | tr -d ' ')"
  if [ "$count" = "0" ]; then
    say "the review holds nothing open"
    exit 0
  fi
  say "$count open thread(s)"

  # The fixer reads every open finding at once, with the diff of the
  # pull request. One agent run answers the whole review, so a change
  # that serves two findings is written once.
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

  say "handing $count finding(s) to the fixer"
  SUMMARY="$STATE_DIR/summary-$round.md"
  before="$(git rev-parse HEAD)"
  if ! AUTOTUNE_EVAL_DOC="$FINDINGS" \
       AUTOTUNE_LABEL="review-$PR-$round" \
       AUTOTUNE_LAST_GOOD="$before" \
       AUTOTUNE_FIXER_PROMPT="$ROOT/docs/reference/feedback-fixer-prompt.md" \
       "$ROOT/scripts/autotune-fix.sh" > "$SUMMARY" 2>&1; then
    say "the fixer failed on round $round. Read $SUMMARY"
    exit 1
  fi

  after="$(git rev-parse HEAD)"
  if [ "$before" != "$after" ]; then
    if ! ( cd "$ROOT/go" && go build ./... && go vet ./... && go test ./... >/dev/null ) \
       || ! ( cd "$ROOT" && make lint-go >/dev/null 2>&1 ); then
      say "the tree is red after the review fix. The cycle reverts this round."
      git reset -q --hard "$before"
      exit 1
    fi
    git push -q || { say "could not push the review fix"; exit 1; }
    say "pushed $(git rev-list --count "$before".."$after") commit(s)"
  else
    say "the fixer changed nothing this round"
  fi

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

  round=$((round + 1))
done

say "the review still holds open threads after $ROUNDS round(s)"
exit 1
