# PR-86: the MTGO month pages of the meta job (2026-09-26)

This record holds the reads of PR-86 (F-179, D-962). The session read the deploy of #237 and the logs and the store of `mtg-meta` on `decktome-prod`. Every read is of 2026-09-26, between 22:40 and 23:30 UTC.

## The deploy of #237

| Item | Read |
|---|---|
| Merge commit | `b739386`, from #237 |
| Web build | `deploy-web`, SUCCESS at 22:42:31 UTC |
| API build | none, because #237 changes no file of `go/` |
| `/version.json` | commit `b739386b18f4ef621193acb28db2d9a11c68712c` |
| Job image | `worker:1e6f73c`, from #236 |

## The scheduled runs

| Date | Execution | End | MTGO pages | MTGO lists | MTGO fetch errors |
|---|---|---|---|---|---|
| 2026-09-20 | `mtg-meta-44fks` | "meta job done" | 0 | 0 | 0 |
| 2026-09-21 | `mtg-meta-cznwb` | "meta job done" | 4 | 64 | 23 |
| 2026-09-22 | `mtg-meta-qpr5c` | "meta job done" | 9 | 64 | 19 |
| 2026-09-23 | `mtg-meta-8hg9q` | "meta job done" | 8 | 69 | 17 |
| 2026-09-24 | `mtg-meta-qlfmq` | "meta job done" | 0 | 0 | 0 |
| 2026-09-25 | `mtg-meta-dgdnq` | "meta job done" | 0 | 0 | 0 |
| 2026-09-26 | `mtg-meta-4spgm` | "worker failed", exit 1 | 0 | 0 | 0 |
| 2026-09-26, manual | `mtg-meta-8qg4p` | "meta job done" | 200 | 5,391 | 89 |

The run of 2026-09-26 at 06:00 UTC failed the precon bar: "commander: precon over own copy 0.94 of 827, the bar is 0.95". This is F-178, and #236 fixed it.

No run logged "meta source failed" or a retry of an MTGO page. So each month page of a 0 run answered 200 with no event link, or 404. The job logged no line for one month page, so the logs can not tell the two apart.

CAUTION: the job writes the level of a log line into `jsonPayload.level`, and Cloud Logging reads no severity from it. A filter on `severity>=ERROR` misses an error line of the job. Filter on `jsonPayload.level="ERROR"`.

## The stored event pages

The store holds 830 MTGO event pages under `meta/raw/mtgo/`. This table counts them by the hour of storage and by the month of the event.

| Stored (UTC) | Event months before September | September events |
|---|---|---|
| 2026-09-07, 17:00 to 18:59 | 321: June 37, July 129, August 155 | 33 |
| 2026-09-11, 06:00 | 100: June | 9 |
| 2026-09-14, 06:00 | 111: March | 6 |
| 2026-09-26, 20:00 to 21:59 | 169: February 33, March 33, April 94, June 9 | 31 |
| Each other 06:00 run, 2026-09-08 to 2026-09-23 | 0 | 0 to 10 |

These facts follow from the table and the code of `runMTGO`:

- The runs of 2026-09-24 to 2026-09-26 stored no September page. The manual run of the same day stored 31 new September pages. So those runs got no event link from any month page. The cause is not "every page already stored".
- The run takes the links newest first, up to 200 pages. The runs of 2026-09-19 and 2026-09-21 to 2026-09-23 read September links alone, with no cap. Each refused page was a September page. April, May, and October 2025 to January 2026 were not stored, so the older month pages gave those runs no link.
- The 06:00 runs of 2026-09-11 and 2026-09-14 stored 100 and 111 older pages. So the loss is not a fault of the hour.
- No run stored a page of May 2026, or of a month before February 2026.

## The month pages from this Mac

At 22:48 UTC, the session fetched four month pages with the user agent and the accept header of the job.

| Month | Status | Bytes | Modern and Standard event links |
|---|---|---|---|
| 2026/09 | 200 | 307,533 | 129 |
| 2026/08 | 200 | 366,146 | 155 |
| 2026/03 | 200 | 412,285 | 158 |
| 2025/10 | 200 | 342,848 | 149 |

## The cause

UNVERIFIED. The month pages give the job event links at some times and none at others. The pattern follows no hour and no month. The site served each page in full to this Mac.

## The change

`runMTGO` logs one line for each month page (D-962):

- `mtgo month page`, with `month`, `found`, `bytes`, `links`, and `covered`. The `covered` count holds the new Modern and Standard links.
- `mtgo event slugs`, with `listed`, `stored`, `fetched`, and `fetch_errors`, at the end of the MTGO read.

A month page with no event link stays in the store under `meta/raw/mtgo-month-empty/`, with the month and the time of the run in its name. A reader can then see what the site served.

## Open

- The cause of a month page with no event link. The first scheduled run after the deploy of this pull request logs each month page.
