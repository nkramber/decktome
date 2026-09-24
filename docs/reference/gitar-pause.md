# The Gitar pause (D-838)

The owner paused the Gitar requirement on 2026-09-23. The pause stays until the owner ends it in a later pull request. While this file exists, its rules win over each Gitar rule of `CLAUDE.md`, `AGENTS.md`, the skills, and the reference documents.

This file is also the switch of the unattended feedback cycle. `scripts/feedback-review.sh` reads it (D-838).

## What the pause changes

- No step waits for a Gitar review. Do not do the push wait, and do not comment `Gitar review`.
- The Codex review needs no Gitar pass. Run `make codex-review PR=<n> -- --skip-gitar-review` for each Codex review.
- A pull request of documents alone takes the `review-override` label when each other check is green. It does not wait for Gitar (amends D-679).
- The Codex reviewer needs no current Gitar review. It still verifies each Gitar claim that exists.
- The feedback cycle needs no Gitar review. Its review step reads Gitar one time, never waits for it, and then runs the Codex review (D-878). A Gitar finding stops it before the fixer, and a comment on its pull request tells the owner.

## What stays

- Gitar can still review a push. Each Gitar finding still gets its answer on the same pull request.
- The ruleset of `main` still requires each review thread resolved. So `--skip-gitar-review` still refuses an open review thread or a Gitar finding, and it spends nothing.
- The flag `--skip-gitar-review` stays after the pause ends (D-838).

## A Gitar finding stops the work

A Gitar finding is a review thread of `gitar-bot`, or an issue in the `Code Review` block of its dashboard comment. A dashboard that reports no issue is not a finding. A pause note of Gitar is not a finding. A plan notice, and each other Gitar comment that names no specific item, is not a finding. Do not reply to it (D-842).

`dashboard_issue` of `docs/tools/codex_review.py` reads the dashboard for the flag and for the feedback cycle. A clean dashboard has the verdict `✅ Approved` or `✅ No issues found`, and each of its findings is resolved or closed. Each other form counts as a finding.

1. After each push, wait for each check on the tip to complete.
2. Read the Gitar threads and the newest dashboard comment one time, with command B of the `gitar-review` skill.
3. Read them again before the Codex review, the merge question, and the auto-merge.
4. When Gitar wrote a finding, stop the work on the pull request at once.
5. Tell the owner the finding, its link, and its claim. The owner then decides about the Gitar requirement.
6. Answer the finding only after the owner replies.

## The end of the pause

The owner ends the pause in a later pull request. Do these steps in that pull request:

1. Run `git grep -n -F "**Gitar pause (D-838).**"` to find each pause note.
2. Delete each paragraph or bullet that the command finds.
3. Delete this file. The feedback cycle then waits for Gitar again, and it answers each finding before the Codex review.
4. Keep the flag `--skip-gitar-review` and the pause branch of `scripts/feedback-review.sh`.
5. Run `make lint`. The test `docs/tools/test_gitar_pause.py` fails when a pause note stays without this file.
6. Record the end as a new decision.
