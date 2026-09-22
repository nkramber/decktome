# Review records

This folder holds one review record for each pull request, `pr-<number>.md`, and the answer of the author, `pr-<number>-response.md` (D-803).

The provider that wrote a pull request does not review it. The other provider writes the record, and the `review-gate` check reads its Head field and its verdict (D-804). Load `.claude/skills/pr-review/SKILL.md` before you write or answer a record.

A record is history. Never delete a finding, and never change the words of an earlier verdict. A repeat review moves that verdict under `## Earlier verdicts`.
