# Pull request 212 response

Date: 2026-09-23

This file answers the review record `docs/reviews/pr-212.md` of effective head `cc386b4`. The verdict was `Changes required`, with one finding.

## P2-1: The required record title fails the reference check

Result: full merit.

Evidence: the record passed `make ref-check` with its own title. The same record with the template title, with the number 212 after the roadmap prefix, gave one REF 1 finding. No other file of the skill or the tools writes the GitHub number after `PR-`.

Correction: the skeleton of `.claude/skills/pr-review/references/review-record.md` now starts with `# Pull request <number> review`, the form of the record. The section "The record" also says that the GitHub number never follows `PR-`, because REF 1 reads that form as a roadmap id (D-753). The fix changes the template, and not the checker, so REF 1 keeps its full reach.

Regression check: `test_a_filled_skeleton_passes_the_reference_check` in `docs/tools/test_review_gate.py` fills the skeleton with the number 212. It runs `ref_check.check` against the real registers. It failed on the old title, and it passes on the new title. The tool tests pass, and `make ste-check` and `make ref-check` give no finding.

## P2-2: The roadmap mark accepts a longer pull request number

Result: full merit.

Evidence: `check_pr` with the number 212 and the roadmap text `✅ merged as #2120` gave no mark error on `f500798`.

Correction: `docs/tools/pr_check.py` now needs a non-digit, or the end of the text, after the number of the mark (D-822).

Regression check: `test_the_mark_of_a_longer_number_fails` tries `#2120` with and without a period, and `#21`. The two `#2120` cases failed on `f500798`, and each case passes now. `test_the_mark_at_the_end_of_the_text_passes` covers the mark at the end of the text.

## New ids

None.

## Final head

The correction commit is the new effective head. `python3 docs/tools/review_gate.py --effective-head 212` prints it.
