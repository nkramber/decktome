# PR-92: the read of the scheduled meta job, 2026-09-27

This file holds the reads of PR-92. No paid target ran. No model call ran. Each count comes from the Cloud Run job logs and the store of `decktome-prod`, and from fetches on the Mac of the owner.

## The execution

- The schedule `mtg-meta-schedule` reads `0 6 * * *`, ENABLED, time zone `Etc/UTC`.
- The job `mtg-meta` runs `worker:f178f05`, the image of #242.
- The execution `mtg-meta-xpt2x` started at 06:00:51 UTC and ended at 06:26:23 UTC on 2026-09-27. It read one success.
- The job wrote 34 log lines: 29 INFO and 5 WARNING. Each line held `jsonPayload.message` and a severity, and no line held `msg` or `level` (F-180).

## The four checks of next step 3

| Check | Result |
|---|---|
| "meta job done" | Pass. It logged at 06:26:16 UTC with the model `20260927T061832Z`. |
| A new `meta/model/` version | Pass. The store holds `meta/model/20260927T061832Z/`. The newest version before the run was `20260926T212435Z`. |
| 12 `mtgo month page` lines | Pass. The job logged one line for each month, 2025-10 to 2026-09. |
| Each `links=0` page | 11 months read `links=0`. The store holds the 11 pages under `meta/raw/mtgo-month-empty/`. |

The model fit read 53,845 lists and 1,511 commanders. The model file holds 20,284,445 bytes. The three fits read these accuracies: Standard 0.6473, Modern 0.6787, and Commander 0.5601.

## The month pages of the job

| Month | Time (UTC) | Bytes | Links | Covered |
|---|---|---|---|---|
| 2026-09 | 06:03:48 | 316,513 | 383 | 132 |
| 2026-08 | 06:03:49 | 29,578 | 0 | 0 |
| 2026-07 | 06:03:50 | 29,572 | 0 | 0 |
| 2026-06 | 06:04:21 | 29,572 | 0 | 0 |
| 2026-05 | 06:04:51 | 29,569 | 0 | 0 |
| 2026-04 | 06:05:21 | 29,575 | 0 | 0 |
| 2026-03 | 06:05:51 | 29,575 | 0 | 0 |
| 2026-02 | 06:06:22 | 29,584 | 0 | 0 |
| 2026-01 | 06:06:52 | 29,581 | 0 | 0 |
| 2025-12 | 06:07:22 | 29,584 | 0 | 0 |
| 2025-11 | 06:07:52 | 29,584 | 0 | 0 |
| 2025-10 | 06:08:22 | 29,581 | 0 | 0 |

The page of 2026-06 and each older page came about 30 seconds after the page before it. The job logged no retry of an MTGO page. The cause of the 30 seconds is not known.

The September links gave 132 Modern and Standard slugs. The store held 128 of them. The job fetched one new event page, and three event pages answered 302. The MTGO lane read 32 lists.

## The stored empty pages

The session read the pages of 2026-08 and 2025-10 from `meta/raw/mtgo-month-empty/`.

- Each page is the real MTGO page. Its title names the month, for example "August 2026 Decklists | Magic: The Gathering Online".
- Each page holds the header, the menu, and the footer of the site.
- The deck list of each page is an empty `ul` element. The page holds no link to `/decklist/`.

So the site served the page, and it served the deck list with no item.

## The month pages from the Mac

From 06:28:58 to about 06:31 UTC on 2026-09-27, the session fetched the month pages from the Mac of the owner.

| Client | Month | Status | Bytes | Event links |
|---|---|---|---|---|
| curl, default headers | 2026-08 | 200 | 366,146 | 451 |
| curl, default headers | 2026-09 | 200 | 316,513 | 383 |
| curl, the headers of the job | 2026-08 | 200 | 366,146 | 451 |
| curl, the headers of the job, gzip | 2026-08 | 200 | 366,146 | 451 |
| Go `net/http`, the headers of the job | 2026-08 | 200 | 366,146 | 451 |
| Go `net/http`, the headers of the job | 2025-10 | 200 | 342,848 | 419 |

The headers of the job are the user agent `mtg-deck-builder/0.1 (github.com/nkramber/decktome)` and the accept header of `go/internal/meta/meta.go`. The site answered each month page with `Cache-Control: no-cache, no-store`, `Server: Apache`, and `Vary: Accept-Encoding,User-Agent`. The site names no cache in front of the page.

## What the reads prove

- The user agent, the accept header, gzip, and the Go client do not cause the empty list. Each one got the full list from the Mac.
- The site builds each month page for each request. So the job got a page that the site built with an empty list.
- The September page of the same run held its list. So the fault touches the older months alone.

Two causes stay open:

- The site answers an older month with an empty list to the address of the job.
- The site answers an older month with an empty list at times, to any client.

Facts weaken both causes. The 06:00 runs of 2026-09-11 and 2026-09-14 stored 100 and 111 older event pages. So the hour alone does not empty the list, and the job got full lists before. The manual run `mtg-meta-8qg4p` from Cloud Run stored 169 older event pages. `docs/reference/pr86-mtgo-month-pages-2026-09-26.md` holds those counts.

## The owner decision

The owner chose to record the reads in this pull request (D-973). A later pull request makes the job fetch an empty older month page again, once, after a delay. That change can fix both causes, and its log lines can tell them apart.
