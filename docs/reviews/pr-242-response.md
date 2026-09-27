# Response to the review of pull request 242

The author answers the Codex record of `docs/reviews/pr-242.md`. That record read `Blocked` at head `7e4ec70`, with no finding. It named two open items.

## The Gitar CI notice has no author reply

- The result: full merit.
- The evidence: the dashboard comment of Gitar names one specific item, the failure of `review-gate` for want of `docs/reviews/pr-242.md`. D-842 exempts a Gitar comment that names no specific item, so this notice needs an answer.
- The correction: comment 5852326252 on the pull request answers the notice. The failure is by design, because the record comes from the second review after the Gitar pass (D-815, D-823). Commit `e4a2c1a` adds the record. No file of the code changed.
- The regression check: `gh api repos/nkramber/decktome/issues/242/comments` lists the answer after the dashboard comment.

## The result of `review-gate` after publication

- The result: no merit as a defect of the pull request.
- The evidence: `review-gate` reads the verdict of the record. It fails while the record reads `Blocked`, and it passes on `Ready for owner merge` (D-815). The repeat review reads the check on its own record.
- The correction: none. The repeat review runs after this response.
- The regression check: the repeat review reads `gh pr checks 242`.
