# Response to the review of pull request 258

The author answers the Codex record of `docs/reviews/pr-258.md`. That record read `Changes required` at head `2089b50`, with one finding.

## P2-1: Sign-out skips push cleanup when local storage fails

- The result: full merit.
- The evidence: `releasePushOnSignOut` returned at once when `readFlag()` was false. A browser can refuse `localStorage.setItem`, and `writeFlag` then stores nothing. So a device that the API holds and a live Cloud Messaging registration stayed after the sign-out. That breaks the sign-out rule of D-1005.
- The correction: with no flag, `web/apps/web/src/features/push/push.ts` asks the API when the browser granted the notification permission. The API is the source of truth. A device that the API holds goes through `disablePush`, which also ends the Cloud Messaging registration. A read that fails releases the device all the same. A browser with no permission reads nothing, because the toggle is the only place that asks for the permission.
- The regression check: three new tests in `web/apps/web/src/features/push/push.test.ts`. "releases a device the API holds when the flag is missing" and "releases the device when the read of the API fails and the flag is missing" fail on `2089b50` and pass with the fix. "leaves a device the API does not hold when the flag is missing" holds the case with nothing to release. The push, menu, and auth tests pass, 69 of 69.

## A note on the verification of the record

The record says that `make verify` failed at web test startup with `ERR_REQUIRE_ESM` under Node 20.17.0. The author ran `make verify` with Node 22 on the PATH, and it passed on `be954ee`. The web tests of the repo need Node 22, and pnpm runs from the Node 20 bin.

# Response to round 2

The Codex record read `Changes required` at head `6c49dd6`, with one finding.

## P1-1: The hand-off exceeds the context budget

- The result: full merit.
- The evidence: the review line of `6c49dd6` took `docs/SESSION-HANDOFF.md` to 24,013 bytes, over the limit of 24,000 (D-749). The author ran `make context-budget` before that last edit, and not after it. `verify:shell` failed on `6c49dd6`.
- The correction: the record commit `84dec9f` took the file to 23,735 bytes. The author shortened the resume section again, to 23,621 bytes, so the next review line has room. The hand-off is in the metadata set, so the effective head does not move for this correction (D-752).
- The regression check: `make context-budget` reads 23,621 of 24,000 bytes, and `verify:shell` reads success on `84dec9f` in run 36665252826.

## The audit failure of `verify:web`

`verify:web` failed on `6c49dd6` at `pnpm audit --audit-level high`. The cause is new: advisories GHSA-qhr7-859c-m2p7 and GHSA-6j4f-fj2g-mc7p on `brace-expansion`. `main` holds the same lockfile, so its audit fails too. The owner chose the fix in this pull request (D-1006).

- The correction: `pnpm --dir web update -r brace-expansion` changes `web/pnpm-lock.yaml` alone. The package moves to 1.1.21, 2.1.7, and 5.0.12.
- The regression check: the audit reads 1 moderate advisory and no high one. `pnpm install --frozen-lockfile` passes, and the whole web suite passes, 443 of 443, with lint and typecheck.

# Response to round 3

The Codex record read `Blocked` at head `49dda67`, with no open finding. The one reason: the live gate of the roadmap needs paid evidence on a real device.

## The live gate

- The result: no merit for the merge. The owner decided the order (D-1007).
- The evidence: the live check needs the new API on Cloud Run. Production deploys from `main` alone (D-579, hard rule 9), so no pull request can run the check before its merge. The roadmap entry of PR-26 already named the unit tests and the emulator tests as the gate of this pull request. It named the live check as a step after the deploy. PR-101 and PR-102 took the same order for the blocking function.
- The correction: D-1007 records the order. The roadmap entry now says that the live check comes after the merge, and that it does not stop it. The hand-off names the check as the first action of the next session, and that session asks the owner before the paid build.
- The regression check: none, because no code changed. `make verify` passed on `49dda67` in the record itself.
