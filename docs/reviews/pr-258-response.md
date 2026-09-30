# Response to the review of pull request 258

The author answers the Codex record of `docs/reviews/pr-258.md`. That record read `Changes required` at head `2089b50`, with one finding.

## P2-1: Sign-out skips push cleanup when local storage fails

- The result: full merit.
- The evidence: `releasePushOnSignOut` returned at once when `readFlag()` was false. A browser can refuse `localStorage.setItem`, and `writeFlag` then stores nothing. So a device that the API holds and a live Cloud Messaging registration stayed after the sign-out. That breaks the sign-out rule of D-1005.
- The correction: with no flag, `web/apps/web/src/features/push/push.ts` asks the API when the browser granted the notification permission. The API is the source of truth. A device that the API holds goes through `disablePush`, which also ends the Cloud Messaging registration. A read that fails releases the device all the same. A browser with no permission reads nothing, because the toggle is the only place that asks for the permission.
- The regression check: three new tests in `web/apps/web/src/features/push/push.test.ts`. "releases a device the API holds when the flag is missing" and "releases the device when the read of the API fails and the flag is missing" fail on `2089b50` and pass with the fix. "leaves a device the API does not hold when the flag is missing" holds the case with nothing to release. The push, menu, and auth tests pass, 69 of 69.

## A note on the verification of the record

The record says that `make verify` failed at web test startup with `ERR_REQUIRE_ESM` under Node 20.17.0. The author ran `make verify` with Node 22 on the PATH, and it passed on `be954ee`. The web tests of the repo need Node 22, and pnpm runs from the Node 20 bin.
