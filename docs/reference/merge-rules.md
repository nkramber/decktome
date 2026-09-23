# The merge rules of `main`

This file describes the ruleset and the merge settings that let a pull request merge itself (D-828). `.github/rulesets/` holds both. `make ruleset-check` compares them with GitHub, and it exits 1 on a difference.

## The green light

A pull request merges itself when three conditions hold (D-828):

1. Each required check is green on the head.
2. The Codex record approves the effective head. The `review-gate` check reads the record.
3. Each Gitar comment has its answer.

The owner then confirms the merge after a summary in four sections: What, How, CI, and Codex review (D-834, D-836). The session turns on the auto-merge after that confirmation alone.

The ruleset enforces conditions 1 and 2, and the thread part of condition 3. A top-level Gitar comment is not a review thread, so the ruleset can not read its answer. So the author session reads each one with the `gitar-review` skill before it turns on the auto-merge.

## The ruleset

`.github/rulesets/review-gate.json` is the body of ruleset 23858584 on the default branch. It holds two rules and no bypass actor (D-815, D-830).

| Rule | Content |
|---|---|
| `required_status_checks` | `review-gate`, `pr-contract`, and nine jobs of `verify`, each from the GitHub Actions app, id 15368. |
| `pull_request` | No approval, each review thread resolved, and the squash merge alone. |

The nine jobs of `verify` are `skip`, `go`, `vuln`, `emulator`, `web`, `proto`, `shell`, `eval`, and `docker`. The ruleset omits `verify:changes`, because that job runs on a dispatch alone (D-832). A test of `docs/tools/test_ruleset_check.py` compares the list with the job names of `.github/workflows/verify.yml`. So a new job fails `make lint` until the ruleset file names it.

The check policy is not strict. A pull request does not need the newest `main` before it merges. The second ruleset, `naub`, refuses a deletion and a force push of `main`, and this file does not hold it.

## The merge settings

`.github/rulesets/merge-settings.json` holds four settings of the repository. The auto-merge is on, and the squash merge is the one merge method. The rebase merge and the merge commit are off, to agree with `AGENTS.md`.

## A skipped job

A job that its `if` condition skips reports the conclusion `skipped` under its own name. A required check accepts `skipped` as a pass. The skip job of D-818 skips the six heavy jobs on a change of documents alone. So the names stay the same on each kind of head, and no aggregate job is necessary (D-832).

A check that never reports blocks the merge. A job with a new name, or a workflow with a path filter, does that. Change the ruleset file and the live ruleset in the same pull request as such a change.

## Change the rules

Get the approval of the owner before each change. The settings are outward-facing. Then run these commands from the root of the repo:

```bash
gh api -X PUT repos/nkramber/decktome/rulesets/23858584 --input .github/rulesets/review-gate.json
gh api -X PATCH repos/nkramber/decktome --input .github/rulesets/merge-settings.json
make ruleset-check
```

The last command must print that the live rules match the files.
