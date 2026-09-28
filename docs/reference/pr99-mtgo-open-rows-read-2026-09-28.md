# PR-99: the read of the open MTGO rows, 2026-09-28

This file holds the reads of PR-99. No paid target ran. No model call ran. Each count comes from the Cloud Run job logs of `decktome-prod`. Each time is UTC.

## The run

The execution `mtg-meta-4wbpp` started at 16:27:21 on `worker:bdc5b60`. `docs/reference/pr98-mtgo-current-month-read-2026-09-28.md` holds its month pages up to 16:46. This file reads the rest of the run.

- At 17:08 the run was still in progress. The scheduled run of 06:00 on 2026-09-29 was 13 hours away.
- The owner chose to read the end of this run before the pick of a row, and not to wait for the scheduled run (D-986).
- The run read SUCCESS at 18:22:08, with 1 task complete and 0 failed.

## The older months (F-179)

| Step | Time | Pages | Full | Empty |
|---|---|---|---|---|
| First read | 16:29:29 to 16:33:31 | 12 | 3 | 9 |
| Pass 1 | 16:44:17 to 16:46:05 | 9 | 6 | 3 |
| Pass 2 | 16:50:44 to 16:51:30 | 3 | 0 | 3 |
| Pass 3 | 16:56:41 to 16:56:54 | 3 | 3 | 0 |

- The 3 months of pass 2 and pass 3 were 2026-08, 2026-04, and 2026-02.
- Pass 2 read each one with 29,575 to 29,584 bytes and 0 links.
- Pass 3 read each one with 419 to 451 links, 25 to 27 minutes after the first read.
- So 0 older months stayed empty after the passes.
- The scheduled run `mtg-meta-5c425` read 10 older months empty, and 2 passes read all of them full (PR-97).

The owner closed F-179 on these two runs (D-987).

## The event pages (F-183)

| Pass | End | Listed | Fetched | Redirected |
|---|---|---|---|---|
| 0 | 17:32:44 | 1,764 | 96 | 73 |
| 1 | 17:59:31 | 73 | 58 | 15 |
| 2 | 18:07:53 | 15 | 14 | 1 |
| 3 | 18:13:12 | 1 | 1 | 0 |

- Pass 0 found 1,595 of the 1,764 listed slugs in the store. So 1,595 + 96 + 73 = 1,764, and the page cap of 200 did not hold.
- The fetched pages add up to 169, the same as the line `meta source` of MTGO: 169 pages, 5,199 lists, 0 failures, and 0 fetch errors.
- The last page to answer 302 on pass 2 was `modern-league-2025-10-059742`. Pass 3 read it full.
- The 302 slugs of pass 0 came from 2025-10 to 2026-05.
- So 0 event pages stayed redirected after the passes.
- The run `mtg-meta-4nps4` left 6 pages redirected after pass 3 (PR-96). In the run `mtg-meta-5c425`, the time budget held before pass 2 (PR-97).

This run is the first run with 0 redirected pages at the end. The owner kept F-183 open for a second run (D-987).

## The end of the run

- The MTGO lane ended at 18:13:12, 1 hour 46 minutes after the start of the run. The time budget of 2 hours did not hold (D-982).
- The quality fit stored the model `20260928T181312Z` at 18:22:01.
- The line `meta job done` came at 18:22:01.
