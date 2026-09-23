# Pull request 212 response

Date: 2026-09-23

This file answers the review record `docs/reviews/pr-212.md` of effective head `cc386b4`. The verdict was `Changes required`, with one finding.

## P2-1: The required record title fails the reference check

Result: full merit.

Evidence: the record passed `make ref-check` with its own title. The same record with the template title `# PR-212 review` gave one REF 1 finding. No other file of the skill or the tools writes the GitHub number after `PR-`.

Correction: the skeleton of `.claude/skills/pr-review/references/review-record.md` now starts with `# Pull request <number> review`, the form of the record. The section "The record" also says that the GitHub number never follows `PR-`, because REF 1 reads that form as a roadmap id (D-753). The fix changes the template, and not the checker, so REF 1 keeps its full reach.

Regression check: `test_a_filled_skeleton_passes_the_reference_check` in `docs/tools/test_review_gate.py` fills the skeleton with the number 212. It runs `ref_check.check` against the real registers. It failed on the old title, and it passes on the new title. The tool tests pass, and `make ste-check` and `make ref-check` give no finding.

## New ids

None.

## Final head

The correction commit is the new effective head. `python3 docs/tools/review_gate.py --effective-head 212` prints it.
