# PR-91: the IPv6 check on the deploy of #242, and the move of the checkout, 2026-09-27

This file holds the reads of PR-91. No paid target ran. No model call ran. The file names no address of the owner. Each count comes from Cloud Build and from the Cloud Run request logs of `decktome-prod`.

## The deployed API

- The `deploy-api` build `f10c2f8f` of `f178f05` read SUCCESS. It started at 03:54:24 UTC and ended at 03:59:44 UTC on 2026-09-27.
- The API serves revision `mtg-api-00091-xd9` with all the traffic. The Cloud Run Admin API read its create time as 03:59:19 UTC, inside that build.
- Its image is `mtg/api@sha256:14eb127d52922c9e171bdfee47eeac24f5823d2b0039a32d362db849999362b2`.
- `/readyz` read `ok` with the version `f178f05bc5c3f23eb69100274597687feef5d6a6`.

## The shared reads over IPv6 (F-182, D-970)

The calls ran at 04:00:17 to 04:00:29 UTC, with the shape of PR-90:

- `POST https://mtg-api-qk2ackpb3q-uc.a.run.app/mtg.v1.DeckService/GetSharedDeck`
- The body `{"token":"probe-no-such-token"}`. No share holds that token, so the calls read no deck of a user.
- The header `X-Forwarded-For: 2001:db8::N`, with N from 1 to 61.
- `curl -6 --interface`, with the source address in turn from two addresses of one /64. The script refused two addresses from two prefixes.

| Call | Status | Body |
|---|---|---|
| 1 to 60 | 404 | `not_found`, "deck not found" |
| 61 | 429 | `resource_exhausted`, "too many calls from this address, try again in a minute" |

The request logs held 61 entries, all on revision `mtg-api-00091-xd9` and on one instance. They held two IPv6 remote addresses of one /64: 31 calls and 30 calls. The first read of the logs found 55 entries, and a second read one minute later found all 61. So the delay came from the log intake.

PR-90 read 61 answers of 404 from the same shape of source, before its fix. So the /64 key of D-970 holds on the deploy, and F-182 is fixed.

## The move of the checkout (D-971, D-972)

The session started in `/Users/nate/Repos/decktome`, on the branch of #236, which merged. Its tree held two untracked files alone.

1. The owner chose to delete the two untracked review prompts (D-972). The old tree then held no change.
2. The session cloned `git@github.com:nkramber/decktome.git` to `/Volumes/SSD-1TB/decktome`.
3. It copied `.env`, `.local` (1.8 GB, no symbolic link), and `.claude/settings.local.json` with their attributes. `diff -rq` found no difference.
4. It copied the auto-memory directory of 52 files to the key `-Volumes-SSD-1TB-decktome`. `diff -rq` found no difference.
5. It ran `make hooks` and `pnpm install --frozen-lockfile` in `web/`. Both exited 0.
6. It ran `make verify` in the new checkout. The pull request body holds the result.

No script and no `Makefile` target of this repository names the old path. The user settings of Claude Code and the launch agents of the Mac also name no path of the old checkout.

The old checkout stays until the owner approves its deletion (D-971). It holds `.bin/buf`, which `make` builds again in the new checkout.
