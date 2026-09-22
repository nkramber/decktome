# PR review: repeat review

Part of the `pr-review` skill. Load this file when the session reviews a pull request again after a correction. The scope rules of `SKILL.md` apply to each new finding.

## Repeat review procedure

Do these steps in sequence after the author changes the pull request.

1. Read the response file when one exists.
2. Check the provider gate again. A fix of the reviewer changes the eligibility.
3. Read the new head, the new base, the whole `git diff --stat`, and the diff since the reviewed head.
4. Verify each claimed fix against its first trigger and its regression check.
5. Set the `Status` line of each earlier finding. Keep each id and each piece of evidence.
6. Inspect the new diff for new defects and affected consumers.
7. Give each new finding the next index of its severity.
8. Update the Identity list to the new effective head.
9. Update the Verification section with the commands that ran on the new head.
10. Write the verdict for the new head. Keep one verdict name in the Verdict section.
11. Commit the review record and the hand-off entry together, then run the session end gate.

Edit the existing `docs/reviews/pr-<number>.md`. Do not make a second file for the same pull request. Move the earlier verdict to a section above the Verdict section, under the heading `## Earlier verdicts`. The gate reads the section under the exact heading `## Verdict`, and it fails a section that holds two bold names.

Close a finding only when the evidence proves the fix, or an owner decision settles it. Record each required check that still waits for a result.

## When a finding closes

A finding closes when its stated trigger passes and its regression check passes. Then set the status to `fixed in <sha>`.

A new trigger for the same class of defect is a new finding with a new id. Assess it against the scope rules of `SKILL.md`. It is not a reason to hold the old id open.

Stop at the third assessment of one id. Write the pattern in the review record, and ask the owner which pull request carries the whole surface. A fourth correction of one finding is a scope question, and not a defect.
