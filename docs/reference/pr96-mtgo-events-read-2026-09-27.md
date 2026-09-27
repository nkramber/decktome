# PR-96: the read of the MTGO event fetch, 2026-09-27

This file holds the reads of PR-96. No paid target ran. No model call ran. Each count comes from the Cloud Run job logs of `decktome-prod`. Each time is UTC.

## The deploy of #247

- #247 merged as `e8b2236` at 19:27:28 on 2026-09-27.
- The Cloud Build `b0249c2f` of `deploy-api` read SUCCESS at 19:33:50.
- The jobs `mtg-meta` and `mtg-snapshot` run `worker:e8b2236`.
- The API serves revision `mtg-api-00094-xsc` on `api:e8b2236`.
- The schedule `mtg-meta-schedule` reads `0 6 * * *`, ENABLED.

The first scheduled run on the new image comes at 06:00 on 2026-09-28, about 10 hours after the deploy. The owner chose a manual run in place of that wait (D-981).

## The manual run

The execution `mtg-meta-4nps4` started at 19:46:59 on `worker:e8b2236`. Its task started at 19:49:34. The task timeout of 9,000 seconds stopped it 2.5 hours later, at 22:19:34. The execution read failed at 22:19:36, with the message "The configured timeout was reached". The run logged no `meta job done` line, so it stored no model.

## The first read of the month pages

| Month | Time | Bytes | Links | Covered |
|---|---|---|---|---|
| 2026-09 | 19:49:48 | 319,535 | 387 | 134 |
| 2026-08 | 19:49:55 | 366,146 | 451 | 155 |
| 2026-07 | 19:50:25 | 29,572 | 0 | 0 |
| 2026-06 | 19:50:55 | 29,572 | 0 | 0 |
| 2026-05 | 19:51:26 | 29,569 | 0 | 0 |
| 2026-04 | 19:51:56 | 29,575 | 0 | 0 |
| 2026-03 | 19:52:26 | 29,575 | 0 | 0 |
| 2026-02 | 19:52:56 | 29,584 | 0 | 0 |
| 2026-01 | 19:53:05 | 376,320 | 463 | 148 |
| 2025-12 | 19:53:35 | 29,584 | 0 | 0 |
| 2025-11 | 19:54:05 | 29,584 | 0 | 0 |
| 2025-10 | 19:54:08 | 29,581 | 0 | 0 |

Nine older month pages read empty. Eight empty reads came 30 or 31 seconds after the line before them. The empty read of 2025-10 came 3 seconds after the line before it.

## The month retry passes

The lines `mtgo month page again` by pass:

| Pass | Time | Months read | Full | Empty |
|---|---|---|---|---|
| 1 | 20:04:25 to 20:07:28 | 9 | 2025-11 (139 covered), 2025-10 (149 covered) | 7 |
| 2 | 20:12:48 to 20:15:50 | 7 | none | 7 |
| 3 | 20:21:16 to 20:21:26 | 7 | none | 7 |

- The line `mtgo month pages wait` read pass=2 and wait=4m50s at 20:07:28, and pass=3 and wait=5m0s at 20:15:50.
- In passes 1 and 2, each empty read came 30 or 31 seconds after the line before it. In pass 3, the 7 empty reads came 0 to 4 seconds apart.
- The line `mtgo month pages stay empty` at 20:21:27 named 2026-07, 2026-06, 2026-05, 2026-04, 2026-03, 2026-02, and 2025-12, with passes=3.

So 5 months gave slugs: 134 + 155 + 148 + 139 + 149 = 725 covered slugs.

## The event pages

| Line | Pass | Time | Listed | Stored | Fetched | Redirected |
|---|---|---|---|---|---|---|
| `mtgo event slugs` | 0 | 20:21:27 to 21:18:34 | 725 | 352 | 102 | 98 |
| `mtgo event pages again` | 1 | 21:23:34 to 21:49:59 | 98 | 0 | 49 | 49 |
| `mtgo event pages again` | 2 | 21:54:59 to 22:05:41 | 49 | 0 | 16 | 33 |
| `mtgo event pages again` | 3 | 22:10:41 to 22:19:20 | 33 | 0 | 27 | 6 |

- Each line read fetch_errors=0.
- Pass 0 read 200 pages in 57 minutes, so the page cap held. That is about 17 seconds a request.
- The line `mtgo event pages wait` read wait=5m0s before each of passes 1 to 3.
- The line `mtgo event pages stay redirected` at 22:19:20 read slugs=6 and passes=3.
- The lines `mtgo event page answered 302` counted 98, 49, 33, and 6 for passes 0 to 3.
- In pass 3, four 302 answers in a row came 30 to 35 seconds apart.

The retry of D-980 read 92 of the 98 redirected pages full. So most 302 answers were a fault of the site at that time, as D-980 expected.

## The end of the run

The line `meta source` for `mtgo` at 22:19:20 read pages=194, lists=5046, fetch_errors=6, and the page cap as the skip reason. The job stores each event page and merges its lists at its fetch. So the 194 pages and their lists stay in the store.

The quality fit started at 22:19:20. The line `worker failed` at 22:19:34 read "quality fit: quality: read sixty lists: meta: read meta/lists/sixty/2011-02.jsonl.gz: context canceled".

The time of the run:

| Step | Time | Minutes |
|---|---|---|
| The card index and the month pages | 19:49:34 to 19:54:08 | 5 |
| The other sources | 19:54:08 to 20:04:25 | 10 |
| The month retry | 20:04:25 to 20:21:27 | 17 |
| The event pages, pass 0 | 20:21:27 to 21:18:34 | 57 |
| The 302 retry, 3 passes | 21:18:34 to 22:19:20 | 61 |
| The quality fit, until the timeout | 22:19:20 to 22:19:34 | 0 |

The 302 retry added 61 minutes, and 15 minutes of it were waits. The run `mtg-meta-87jlg` fit its model in about 8 minutes, from 18:21:35 to 18:29:56. So the run needed about 10 more minutes.

## The stopgap and the fix

The session raised the task timeout of `mtg-meta` from 150 to 240 minutes at 22:21 UTC (D-982). The audit log reads `ReplaceJob` at 22:21:41. The job still runs `worker:e8b2236`. The step `update-jobs` of `cloudbuild/api.yaml` changes the image alone, so a later deploy keeps the timeout.

The code stops the MTGO lane at `DefaultMTGOBudget`, 2 hours from its start (D-982). The budget stops each fetch and each wait that would start after it. On this run, the budget would have stopped pass 1 of the 302 retry at about 21:49:45. The fit would then have started at about 21:50, 29 minutes before the old timeout.
