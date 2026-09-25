# Response to the review of pull request 231

The author answers the Codex record of `docs/reviews/pr-231.md`. That record read `Changes required` at head `64a4891`, with one finding.

## P1-1: A delayed build can deploy after a newer build

- The result: partial merit.
- The evidence: the trigger holds. Build A reads an older live commit in its guard step, and a later step deploys. Build B can pass its guard and deploy in that gap, and A then deploys over B. The roadmap text of PR-81 said that an old build can no longer replace a newer one, and that text claimed too much. The window is the one minute of a deploy. The owner accepted it with the guards, in the words of D-943: "Two builds can still overlap during the one minute of one deploy". A lock or a serial pipeline is a change of that decision, so the author asked the owner.
- The owner decision: keep the window of D-943, correct the text, and make each inversion fail a build (D-945). The owner declined a deploy lock, because it needs new grants and a stale-lock rule.
- The correction: each API revision names its commit in `DEPLOY_COMMIT`, and each Hosting release names its commit in its message. After its deploy, each build reads the recent deploys. `api-order` and `web-order` of `docs/tools/deploy_order.py` fail the build when a deploy made before its own names a later commit. The log names the build to run again. The API list covers the jobs too. One build updates the API and the jobs, so a run again of the newer build restores both. The roadmap entry and `docs/deploy-and-rollback.md` now state the window and the failure.
- The regression check: `test_two_builds_that_overlap` runs the trigger of the finding. Both guards read the older live commit and pass. B deploys, then A deploys. The readiness check of A passes, and its order checks of the API and the web fail and name B. `test_order_checks_pass_on_a_right_order` covers a later deploy after this one, a deploy again of the same commit, and deploys with no commit. `BuildFilesTest` requires the stamps and the order checks in both build files. `python3 -m unittest test_deploy_order` passes.
