# PR review: the commit and the session end

Part of the `pr-review` skill. Load this file before you commit a review record or a response file, and before the session ends.

## Commit the record

Commit and push the record in the session that writes it.

| After | Commit these files | Who commits |
|---|---|---|
| A review or a repeat review | `docs/reviews/pr-<number>.md` alone | The reviewer |
| An answer to a review | `docs/reviews/pr-<number>-response.md`, each corrected file, and `docs/SESSION-HANDOFF.md` | The author |

The reviewer does not edit the hand-off. The author reads the record, and records the review state in the hand-off with a metadata commit.

The check reads the head of the pull request. So the check sees the record only after the push. A review is complete only when the branch on GitHub holds the record.

Write the commit message in an impersonal voice, for example `The review record of #212`. Name no provider, agent, harness, or model (hard rule 6).

## The end gate

Run these commands after the commit, in this order:

```bash
git push origin <branch>
git fetch origin
git status --short --branch
gh pr view <number> --json headRefOid --jq .headRefOid
```

The status line must show no `[ahead N]`. The hash from `gh pr view` must be the same as `git rev-parse HEAD`. Write the push line in the `## Verification` section of the record.

When the remote refuses the push, the review is not complete. Tell the owner that the record has a commit and no push. A sandbox with no network can refuse the push with no message from git, so read the status line.

At the start of a review, run `git fetch` and `git status --short --branch` too.
