# PR-90: the rate limit checks on the deploy of #241, 2026-09-27

This file holds the reads of the second pull request after #240 (D-968). No paid target ran. No model call ran. The file names no address of the owner. Each count comes from the Cloud Run request logs of `decktome-prod`.

## The deployed API

- The `deploy-api` build `28dfd4b1` of `ea587fc` read SUCCESS. It started at 02:41:52 UTC and ended at 02:46:23 UTC on 2026-09-27.
- The API serves revision `mtg-api-00090-zzs`. The Cloud Run Admin API read its create time as 02:45:48 UTC, inside that build.
- Its image is `mtg/api@sha256:33cba794516377a250aa6b93aa691883d84f459e32938e81ed92cb6dbd64f779`.
- `/readyz` read `ok` with the version `ea587fcac4c8e18b569594888f7146f6732925ac`.

## The eleven calls of `CheckInvite` (F-181, D-968)

The owner approved the calls. They ran at 02:46:48 to 02:46:50 UTC, with the shape of #241:

- `POST https://mtg-api-qk2ackpb3q-uc.a.run.app/mtg.v1.InviteService/CheckInvite`
- The body `{"email":"probe@example.invalid"}`.
- The header `X-Forwarded-For: 203.0.113.N`, with N from 1 to 11.

| Call | Status | Body |
|---|---|---|
| 1 to 10 | 200 | `{}` |
| 11 | 429 | `resource_exhausted`, "too many calls from this address, try again in a minute" |

The request logs held 11 entries. Each one named the instance `...b340796944c7` of `mtg-api-00090-zzs`, and one IPv6 remote address. So the fix of D-967 holds on the deploy. The key no longer follows a spoofed first line.

## The shared reads over the default route (F-182)

The owner named the shared-read limit as the second concern of this pull request (D-969). The calls ran at 02:51:58 to 02:52:09 UTC:

- `POST https://mtg-api-qk2ackpb3q-uc.a.run.app/mtg.v1.DeckService/GetSharedDeck`
- The body `{"token":"probe-no-such-token"}`. No share holds that token, so the calls read no deck of a user.
- The header `X-Forwarded-For: 198.51.100.N`, with N from 1 to 61.

All 61 calls answered 404 `not_found`. D-315 expects a 429 on the 61st call.

The request logs held 61 entries on the one instance `...b340796944c7`. They held two IPv6 remote addresses in one /64. The Mac sent 56 calls from one address and 5 from the other, in five runs. So no key reached 61 calls.

A scratch test built the handler of `go/cmd/api/main.go` with the same interceptors and the constant of 60. It sent 61 calls through a real HTTP server, each with a new first line and one second line. It read 60 answers of 501 and one of 429. So the wiring holds, and the key caused the miss.

## The shared reads over IPv4

The owner chose the fix of F-182 in this pull request, with a probe over IPv4 first (D-970). The calls ran at 02:56:21 to 02:56:32 UTC with `curl -4`, the same body, and `X-Forwarded-For: 192.0.2.N`, with N from 1 to 61.

| Call | Status | Body |
|---|---|---|
| 1 to 60 | 404 | `not_found`, "deck not found" |
| 61 | 429 | `resource_exhausted`, "too many calls from this address, try again in a minute" |

The request logs held 61 entries on the instance `...b340796944c7`, from one IPv4 remote address. So the shared-read limit of D-315 holds on the deploy for a key that stays the same.

## The cause (F-182)

`ClientAddress` answers one full IPv6 address, and the limiter used it as the key. A network gives each IPv6 host a /64. The host can pick any address in it, and macOS changes its temporary address on its own. So one host can get a new bucket for each call. This defeats the shared-read limit of D-315 and the invite limit of D-592. The invite limit exists to stop a scan of the invite list.

## The fix (D-970)

`bucket` turns the client address into the key. An IPv4 address stays its own key. An IPv6 address keys on its /64. Both interceptors read the key through `clientKey`.

`TestANewAddressInOneSlash64GetsNoNewBucket` sends 11 calls of `CheckInvite`, each with a new first line and a new address in one /64. It read 11 of 11 allowed on the old code, and 10 of 11 on the fix. `TestBucketKeysIPv6OnItsSlash64` reads the key of each address form.

The next deploy proves the fix. 61 shared reads over IPv6 from one /64 must end with a 429, and so must 11 calls of `CheckInvite`.
