---
name: one-pr-one-session
description: Bind a session to one repository, one branch, one pull request, and one role, and make the pull request the complete unit with its code, tests, decisions, documents, and hand-off. Load before any work for a pull request - a start, a revision, a review, a Gitar answer, or the hand-off. Stops a second pull request in the same session, and stops a pull request that only records an earlier merge.
---

# One pull request, one clean session

The owner decisions are D-746 to D-748. `CLAUDE.md` holds the other rules, and this skill does not repeat them.

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

When you can not do one step, do not start the work. The hook `.claude/hooks/session_bind.py` binds the session to its first branch. When it blocks a command, end the session. Only the owner removes a binding.

## 2. The documentation gate

The pull request body holds the sections of `.github/pull_request_template.md`. The `## Documentation impact` table holds one row for each category. Each entry starts with one of these statuses:

- `Changed: <reason and path>`
- `Reviewed; no change needed: <specific reason and path>`
- `Not applicable: <specific reason>`

Obey these rules:

- Change `docs/SESSION-HANDOFF.md` in every pull request. Record the finished state, the checks, the review state, and "pending owner merge".
- Change every document whose facts or contracts the pull request changes. Name the path in backticks.
- A reason says why the document stays correct. "No documentation impact" is not a reason.
- Never defer a document to "after the merge" or to another pull request.
- The hand-off describes only the work in this pull request and the state of its base.
- Mark the own roadmap item `✅ merged as #N` after the number exists. The text reaches `main` only through the merge (D-747).
- Never write a merge commit, a merge time, or a deploy result that you did not read. Git, GitHub, and Cloud Build hold those facts.

Run `make pr-check` before you ask for the review. For a draft body, set `PR_BODY_FILE` and `PR_TITLE`. CI runs the same check on each body edit and push.

## 3. The completion gate

The pull request is ready for the owner only when all of these are true:

- The code and its regression tests are in the pull request.
- The decisions and questions are in `docs/decisions.md` and the question files.
- The roadmap and the hand-off read the state of this pull request.
- `make pr-check` passes, and every canonical document has its row.
- `make verify` passes. The `gitar-review` skill found no open finding on a current review.
- No work waits for a second pull request.

Then tell the owner the pull request is ready, and write this line with the number:

`This session is bound to PR #N and is complete. End this session. Start a new clean session before beginning another PR.`

Do not offer the next pull request.

## Rules of this repo that win over other skills

- This session answers the Gitar review of its pull request, and a Gitar answer never needs a new session.
- Never call the pull request ready before a current Gitar review lands. Fix each finding on the same pull request. So the merge-first trap of the `gitar-review` skill does not occur.
- A merge or a deploy of an earlier pull request never gets its own pull request. The next item reads the base when its own concern needs it.
- An unattended loop pull request stays red on `pr-contract` until a clean author session completes its rows (D-748).
