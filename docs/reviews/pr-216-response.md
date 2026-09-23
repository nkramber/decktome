# Pull request 216: the answer to the review

The author answers each finding of `docs/reviews/pr-216.md` here, with the procedure of `.claude/skills/pr-review/references/answer-review.md`.

## P2-1: The pause misses dashboard-only findings

- Result: full merit.
- Evidence: `docs/reference/gitar-pause.md` defines a Gitar finding as a review thread or an issue in the `Code Review` block of the dashboard. At `9474ed1`, `scripts/feedback-review.sh` read the threads alone. So a dashboard issue with no thread ended the round with 0. The flag `--skip-gitar-review` of `docs/tools/codex_review.py` had the same gap, so the correction covers both.
- Correction: `dashboard_issue` of `docs/tools/codex_review.py` reads the summary line of the newest Gitar dashboard (D-838). A dashboard is clean when its verdict is `✅ Approved` or `✅ No issues found`, and each finding is resolved or closed. Each other form counts as an issue. The flag refuses such a dashboard. During the pause, `pause_stop` of `scripts/feedback-review.sh` calls the same function, and it stops the round with 3 and a comment to the owner.
- The forms: the 90 Gitar dashboards of #120 to #216, read 2026-09-23, hold 8 forms of the summary line, and each one reads clean. No stored dashboard shows an open issue, because Gitar edits its dashboard in place. So the rule counts each unknown form as an issue.
- Regression check: `TestTheGitarPauseStopsTheReviewBeforeTheFixer` runs the round with a fake `gh` in 6 cases. The 2 dashboard cases fail on the script of `9474ed1`, and all 6 pass now. `test_codex_review.py` adds 7 tests of `dashboard_issue` and 1 test of the flag with a dashboard issue and no thread. `make verify` passes.
