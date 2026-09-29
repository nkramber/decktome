# The live check of the blocking function, 2026-09-29 (PR-102, D-991, D-993)

This record holds each read of the registration of the `beforeCreate` trigger on `decktome-prod`. The route came with #253 (D-990). All times are UTC. The addresses of the check are plus-addresses of the check account of D-779, and this record holds none of them (D-639).

## The deploy

| Read | Result |
|---|---|
| The merge of #253 on `main` | `0bfef4e` |
| The Cloud Build run of `deploy-api` | `d29eef54`, SUCCESS, 19:00:50 to 19:05:26 |
| The live revision of `mtg-api` | `mtg-api-00097-8hd`, 100 percent of the traffic |
| The env var `DEPLOY_COMMIT` of that revision | `0bfef4e238cb1b9bd025a208c2f485896a240519` |
| `POST /auth/before-create` with an empty body, 19:05:41 | 400, `{"error":{"message":"no token","status":"INVALID_ARGUMENT"}}` |

Before the deploy, the live revision was `mtg-api-00096-ldb` with `DEPLOY_COMMIT` `bdc5b60`. The build of `0bfef4e` ran when the session started, and the session waited for it.

## The registration

| Read | Result |
|---|---|
| The auth configuration before the change | `subtype: IDENTITY_PLATFORM`, `blockingFunctions: {}` |
| The owner confirmation of the PATCH | D-993 |
| The PATCH of `blockingFunctions.triggers` | HTTP 200 |
| The trigger after the PATCH | `beforeCreate`, `functionUri` `https://mtg-api-qk2ackpb3q-uc.a.run.app/auth/before-create`, `updateTime` 19:05:59.486 |

Identity Platform accepted the URL of a Cloud Run service as `functionUri`. Google documents the field for a Cloud Function, so D-990 marked this item UNVERIFIED. The PATCH and the two paths below verify it.

## The invited path

| Step | Result |
|---|---|
| `make allow` of plus-address 1 | `invited ... in config/allowlist of project decktome-prod` |
| A wait for the cache of the list | 70 seconds |
| `accounts:signUp` of plus-address 1 | HTTP 200, with an `idToken` |
| The call of the route in the request log | 19:08:06.294, status 200, latency 0.110 s, user agent `Google-Firebase` |
| `accounts:delete` with that `idToken` | HTTP 200 |
| `make disallow` of plus-address 1 | `removed ... from config/allowlist in project decktome-prod` |

## The refused path

| Step | Result |
|---|---|
| `accounts:signUp` of plus-address 2, which is not on the list | HTTP 400, with no `idToken` |
| The error message | `BLOCKING_FUNCTION_ERROR_RESPONSE : HTTP Cloud Function returned an error: {"error":{"message":"not-invited","status":"PERMISSION_DENIED"}}` |
| The call of the route in the request log | 19:08:40.347, status 403, latency 0.002 s, user agent `Google-Firebase` |

The refusal made no account, so the check deleted nothing. The error text of Google names a "Cloud Function" for a Cloud Run route too. The web form of #253 reads `not-invited` in that text.

## The end state

- The trigger stays registered. Section 8.5 of `docs/deploy-and-rollback.md` holds its removal.
- The list holds neither plus-address. The project holds no account of either address.
- No paid target ran.
