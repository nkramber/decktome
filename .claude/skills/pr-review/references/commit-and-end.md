# PR review: commit and session end

Part of the `pr-review` skill. Load this file before you commit a review record or a response file, and before the session ends.

## Commit the record

Commit the review record and the hand-off entry, then push them to the branch of the pull request. Do it in the session that writes them.

| After | Commit these files | Who commits |
|---|---|---|
| A review or a repeat review | `docs/reviews/pr-<number>.md` and `docs/SESSION-HANDOFF.md` | The reviewer |
| Work that answers a review | `docs/reviews/pr-<number>-response.md`, each corrected file, and `docs/SESSION-HANDOFF.md` | The author |

Make one commit that holds the record and its hand-off entry. Never leave either file without a commit or a push. The `review-gate` check reads the head of the pull request, so a record that is not on the remote does not exist for it.

A record with no commit has three effects:

- The next commit of the other provider takes it in, and the history no longer shows who wrote what.
- An author can commit an approval that the author never read.
- The `review-gate` check cannot read the record.

Write the commit message in an impersonal voice. Name no provider, agent, harness, or model (hard rule 6). Add no trailer.

## The hand-off entry

Fetch the remote, and read the hand-off again before you write the entry.

- Add a new entry at the top of "The three most recent sessions" in `docs/SESSION-HANDOFF.md`.
- Start the entry with one line: `Author: Claude Code` or `Author: Codex` (D-806).
- Write the pull request, the role, the verdict, and the effective head.
- Move the oldest entry to `docs/reference/session-handoff-archive.md` with no change to its text.
- Keep the hand-off under the byte limits of `make context-budget`.

The reviewer changes the resume section only to state the verdict and the next step. Never change the words of an entry of another session. Both hand-off files are in the metadata set, so a change to them never moves the effective head.

## Session end gate

Run these four commands after the commit, in this sequence. The evidence comes from the remote, not from the local checkout.

```
git push origin <branch>
git fetch origin
git status --short --branch
gh pr view <number> --json headRefOid --jq .headRefOid
```

The status line must show no `[ahead N]`. The hash from `gh pr view` must equal `git rev-parse HEAD`. Write the push line in the Verification section of the review record.

If the remote refuses the push, the review is not complete. Do not end the session. Ask the owner for the push, and say in the hand-off that the record has a commit and no push.

At the start of a review or a repeat review, run `git fetch` and `git status --short --branch` too.
