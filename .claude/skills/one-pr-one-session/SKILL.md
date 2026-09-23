---
name: one-pr-one-session
description: Bind a session to one repository, one branch, one pull request, and one role, and make the pull request the complete unit with its code, tests, decisions, documents, and hand-off. Load before any work for a pull request - a start, a revision, a review, a Gitar answer, a merge message, or the hand-off. Stops a second pull request in the same session, and stops a pull request that only records an earlier merge.
---

# One pull request, one clean session

The owner decisions are D-746 to D-748, and D-823 to D-833 for the review loop and the auto-merge. `CLAUDE.md` holds the other rules, and this skill does not repeat them.

## The rule

A session works on one pull request. A session can make many turns and commits for it. More than one clean session can work on the same pull request, for example a review or a correction. The pull request carries all its work: code, tests, decisions, documents, the review answers, and the hand-off.

## 1. The start gate

Do these steps before the first edit.

1. Read this conversation, not the checkout. Look for work on another pull request, another repository, or a finished pull request.
2. When you find such work, stop. Answer only `Blocked: start a new clean session for this PR.`
3. A fork, a subagent, a compaction, or a summary of such a session is not clean. Stop for these too.
4. Run `make where`. Record the branch, the pull request number or the intent, and the base commit.
5. Keep changes in the checkout that are not yours. Make a worktree for your branch. Such changes do not block a clean session.
6. Name your role: author, reviewer, or correction author.
7. Name the one concern of the pull request. When you see two concerns, ask the owner.
8. Read `docs/SESSION-HANDOFF.md` and the decisions that the change touches.
9. Write the draft matrix of section 2 before the code.

A merge message for the pull request of this session is the one exception to step 2. Section 5 gives the answer for it.

When you can not do one step, do not start the work. The hook `.claude/hooks/session_bind.py` binds the session to its first branch. When it blocks a command, end the session. Only the owner removes a binding.

## 2. The documentation gate

The pull request body holds the sections of `.github/pull_request_template.md`. The `## Documentation impact` table holds one row for each category. Each entry starts with one of these statuses:

- `Changed: <reason and path>`
- `Reviewed; no change needed: <specific reason and path>`
- `Not applicable: <specific reason>`

Obey these rules:

- Change `docs/SESSION-HANDOFF.md` in every pull request. Record the finished state, the checks, the review state, and "pending the auto-merge" (D-828).
- Write the line `Author provider: Claude Code` or `Author provider: Codex` in the hand-off record. The Codex review reads it (D-811).
- Change every document whose facts or contracts the pull request changes. Name the path in backticks.
- A reason says why the document stays correct. "No documentation impact" is not a reason.
- Never defer a document to "after the merge" or to another pull request.
- The hand-off describes only the work in this pull request and the state of its base.
- Mark the own roadmap item `✅ merged as #N` right after the pull request opens, before the Gitar pass. A roadmap commit moves the effective head, so a later mark needs a new review (D-822). The text reaches `main` only through the merge (D-747).
- Never write a merge commit, a merge time, or a deploy result that you did not read. Git, GitHub, and Cloud Build hold those facts.

Run `make pr-check` before you ask for the review. For a draft body, set `PR_BODY_FILE` and `PR_TITLE`. CI runs the same check on each body edit and push.

## 3. The author loop and the completion gate

Do these steps for each round of changes (D-823, D-828):

1. Commit the round, and push it one time.
2. Do the Gitar pass with the `gitar-review` skill. Resolve each thread after its answer.
3. Run `make codex-review PR=<number>` in the background, and wait for the notice of its end.
4. Run `git pull --ff-only`, and read the outcome line.
5. For `changes`, answer each finding with `answer-review.md` of the `pr-review` skill, then go to step 1.
6. For `three-strike stop`, stop the loop, and ask the owner (D-826).
7. For `approve`, go to the completion gate below.

The pull request is complete only when all of these are true:

- The code and its regression tests are in the pull request.
- The decisions and questions are in `docs/decisions.md` and the question files.
- The roadmap and the hand-off read the state of this pull request.
- `make pr-check` passes, and every canonical document has its row.
- `make verify` passes. The `gitar-review` skill found no open finding on a current review.
- The `review-gate` check passes. A Codex record approves the effective head, or the `review-override` label applies (D-811, D-812).
- No work waits for a second pull request.

### The auto-merge

Turn on the auto-merge only when each of these conditions is true (D-828):

- The last metadata commit is on origin. It holds the record and the hand-off.
- The Gitar pass is complete, and each top-level Gitar comment has its answer.
- The record says `Ready for owner merge` for the effective head.

Then run these commands, in this order:

```bash
gh pr merge <number> --auto --squash
gh pr checks <number> --watch
gh pr view <number> --json state,mergedAt,mergeCommit
```

The ruleset of `main` merges the pull request when each required check passes (D-828). When the state is `MERGED`, write the transitional prompt of section 5 at once. When the merge does not come, name the check that blocks it, and ask the owner.

After the prompt, write this line with the number:

`This session is bound to PR #N and is complete. End this session. Start a new clean session before beginning another PR.`

Do not offer the next pull request.

## 4. While the pull request waits

The session stays bound to the pull request while it waits for Gitar, for the Codex review, or for the owner. It answers each finding on the same pull request (D-746).

After the Gitar pass, start the Codex review with `make codex-review` (D-823). A pull request of documents alone takes the label in place of that review.

- Tell the owner that the session is ready for a context compaction while the pull request waits (D-754).
- Say the same when the context of the session passes 300K tokens (D-750).
- Read the resume section of `docs/SESSION-HANDOFF.md` again after a context compaction.
- A context compaction of this session keeps its binding. It starts no new pull request.

## 5. The transitional prompt

### The trigger

Two events start this section. The session reads the state `MERGED` after the auto-merge of section 3 (D-828). Or the owner says that the pull request merged (D-764). The owner asks for no prompt, and the session waits for no other word. Each of these messages is a trigger, and any other variant that names the merge of this pull request:

- `Merged`
- `Merged PR #N`
- `PR #N is merged`
- `#N merged`
- `merged it`

The trigger is the one exception to step 2 of section 1. A merge message for another pull request is not an exception, and it gets the blocked answer of step 2.

Two cases stop the prompt. Ask the owner, and write no prompt until the answer arrives:

- The message names no pull request, and this session holds no binding. Ask which pull request it names.
- The owner merged the pull request before section 3 called it ready. Name each part that did not land, such as an open Gitar finding, a document, or a check. Then ask the owner for the next step.

### The procedure

1. Read the merge commit: `git fetch origin && git log --oneline -1 origin/main`.
2. Confirm that the commit names this pull request.
3. Read the next step of `docs/SESSION-HANDOFF.md`, and name the next item.
4. Read `docs/owner-questions.md`, and name each open question of that item.
5. Name each check that needs `main` or the deploy of this merge.
6. Write the block below in the last message, and stop.

The pick of step 3 is provisional, and the owner can name a different item.

The prompt is one fenced block, and the owner pastes it into the next clean session:

```
Start <item>: <the one concern>

PR #<x> merged to `main` as <sha>. Read `docs/SESSION-HANDOFF.md` first.
Branch: `<prefix>/<slug>`. Base: `<sha>`. Role: author.
Load the `one-pr-one-session` skill and the skills of the task before any change.
<Each check that needs `main` or the deploy of this merge. Run it before the item work.>
Open questions for this item: <each OQ-# with its subject, or `none`>.
First action: <the first concrete action>.
```

### The rules of the prompt

- The prompt carries one item. A second item needs a second session, and a second prompt.
- The prompt never asks the next session to record this merge. Git holds the merge (D-747).
- A check that needs `main` comes first. The branch of the next item gives no such result.
- Remove that line of the block when this merge needs no such check.
- The next step of the hand-off holds the same first action. The two agree, or the hand-off wins.

The session ends with this prompt. It makes no branch and no change for the next pull request.

## Enforcement

| Rule | Enforced by |
|---|---|
| The body, the table, the hand-off change, and a deferred document | `make pr-check` and the `pr-contract` workflow (D-748) |
| A commit on `main` | The pre-commit hook of `make hooks` (D-585) |
| The roadmap mark `✅ merged as #N` in a diff that changes the roadmap | `make pr-check` and the `pr-contract` workflow (D-822) |
| A merge with no approved Codex record, label, or Dependabot exemption | The `review-gate` workflow and the ruleset of `main` (D-815) |
| A merge with a red check, or an open review thread | The ruleset of `main`, and `make ruleset-check` for its content (D-828) |
| A Codex review before the Gitar pass, or on a dirty tree | `make codex-review` (D-832) |
| The third open round of one finding | `make codex-review`, exit 4 (D-826) |
| An API key in a Codex process | `make codex-review` (D-833) |
| A second branch in one session | `.claude/hooks/session_bind.py` (D-748) |
| The skill frontmatter and the wiring | `make lifecycle-check` (D-748) |
| The byte budget of the start read | `make context-budget` (D-749) |
| Each cited id and each repository path | `make ref-check` (D-753) |
| One pull request in each session, and a clean session for each one | The agent. No check reads the conversation |
| The trigger of the transitional prompt, and its two stop cases | The agent. No check reads the conversation (D-764) |
| The truth of each reason, and the one concern | The agent, then the owner |
| The merge | The auto-merge under the ruleset (D-828) |
| The deploy | Cloud Build, from `main` alone (D-579) |

## Rules of this repo that win over other skills

- This session answers the Gitar review of its pull request, and a Gitar answer never needs a new session.
- Never call the pull request ready before a current Gitar review lands. Fix each finding on the same pull request. So the merge-first trap of the `gitar-review` skill does not occur.
- Never turn on the auto-merge before the `review-gate` check passes on its head. The ruleset of `main` refuses the merge until then (D-815).
- A merge or a deploy of an earlier pull request never gets its own pull request. The next item reads the base when its own concern needs it.
- An unattended loop pull request stays red on `pr-contract` until a clean author session completes its rows (D-748).
