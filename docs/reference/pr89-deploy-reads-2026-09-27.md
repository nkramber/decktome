# PR-89: the deploy reads of #230 and #227, 2026-09-27

This record holds two reads on the deployed app. The first read is the feedback list of all verdicts, after the index of #230. The second read is the eleven calls of `CheckInvite` of D-907, after the deploy of #227. The second read failed, and it opened F-181 (D-967).

## The order of the work (D-966)

The session started at about 01:45 UTC on 2026-09-27, before the scheduled run of `mtg-meta` at 06:00 UTC. The owner put these two reads before the read of the meta job. The meta job moves to the third pull request after #240.

## The deployed API

- The API serves revision `mtg-api-00089-z54`. The Cloud Run Admin API read its create time as 01:26:32 UTC on 2026-09-27, the deploy of #239.
- Its image is `mtg/api@sha256:e3b7238ae661fd8c2376858a12f0c08889a3d9579872755b72c857ba9b43b72a`.
- #227 (`72290db`) and #230 (`3b26cde`) come before `2258cb0` of #239 on `main`. So the revision holds the key rule of D-907.

## The feedback list of #230

`make feedback-list VERDICT=` ran at about 01:50 UTC. The empty value reads every verdict, with no filter on "up" or "down".

- The command ended with exit 0 and read "6 verdict(s) in decktome-prod".
- The list held one user, newest first, from 2026-09-24 20:45 to 2026-09-08 02:44.
- It held five "down" verdicts and one "up" verdict. Three card verdicts and one import verdict read "down". Of the two question verdicts, one reads "up".
- The newest verdict is an import "down" of 2026-09-24 20:45, with the reason `parse_fault`. It came after the harvest of 15:28 UTC, so the harvest read five.

The query orders a collection group by `created_at` descending, with no filter. Before #230, the store held "verdict ascending, created_at descending" alone. #230 declared the collection-group override of `created_at` in `firestore.indexes.json` (D-938). So the read proves that the deployed store holds that index.

UNVERIFIED: `gcloud` asked for a new login, so no command described the index itself. The read passed on the application default credentials.

## The eleven calls of `CheckInvite`

The owner approved the calls. Each call ran at 01:53:24 to 01:53:27 UTC on 2026-09-27 with this shape:

- `POST https://mtg-api-qk2ackpb3q-uc.a.run.app/mtg.v1.InviteService/CheckInvite`
- The body `{"email":"probe@example.invalid"}`. The domain is a reserved name, and no person owns it.
- The header `X-Forwarded-For: 203.0.113.N`, with N from 1 to 11. The range is a documentation range.

| Call | Status | Body | Seconds |
|---|---|---|---|
| 1 | 200 | `{}` | 1.08 |
| 2 to 10 | 200 | `{}` | 0.12 to 0.24 |
| 11 | 200 | `{}` | 0.17 |

D-907 expects a 429 on the eleventh call. The eleventh call passed, so the check failed.

## The cause

The Cloud Run request logs of 01:53 to 01:54 UTC held 11 entries for `CheckInvite`. Each entry named the instance `...aab2f88abd9c` of revision `mtg-api-00089-z54`. So one limiter counted all eleven calls, and the key changed with each call. The remote address of the logs is a public IPv6 address, and `proxyHop` never skips one.

`Limiter.Allow` refuses the eleventh call of one key in a window. So the key followed the spoofed address. The interceptor read the header with `Header.Get`, which answers the first header line alone. A proxy that appends its entry as a second line leaves the line of the caller first.

UNVERIFIED: no Google page states how the front end of a `run.app` URL writes this header. The page of the external Application Load Balancer states an append of `<client-ip>,<load-balancer-ip>` to the value of the caller, read 2026-09-27 at `https://cloud.google.com/load-balancing/docs/https`. No log line holds the header that the API read.

## The fix (D-967)

`forwardedFor` joins every `X-Forwarded-For` line in order, and both interceptors read it. `TestASpoofedFirstLineGetsNoNewBucket` sends 11 calls with a new first line and one second line. It read 11 of 11 allowed on the old code, and 10 of 11 on the fix.

The second pull request after #240 makes the eleven calls again, on the deploy of this pull request (D-968). When the eleventh call still passes, the front end writes no client address that the API can read, and the fix needs a new design.
