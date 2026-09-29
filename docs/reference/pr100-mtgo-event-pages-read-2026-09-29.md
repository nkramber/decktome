# PR-100: the read of the MTGO event pages on the scheduled run, 2026-09-29

This file holds the reads of PR-100. No paid target ran. No model call ran. Each count comes from the Cloud Run job logs of `decktome-prod`. Each time is UTC.

## The run

The execution `mtg-meta-b56rq` started at 06:00:45 on `worker:bdc5b60`. The service account `mtg-scheduler` created it, so it is the scheduled run of D-987.

- At 18:47 on 2026-09-28 the run was 11 hours away. The owner chose to wait for it in the same session (D-988).
- The run read SUCCESS at 07:05:55, with 1 task complete and 0 failed.
- It stored the quality model `20260929T065701Z`, fitted on 81,431 lists and 1,513 commanders.
- The run `mtg-meta-4wbpp` of PR-99 ran on the same image digest.

## The event pages (F-183)

| Pass | End | Listed | Stored | Fetched | Redirected |
|---|---|---|---|---|---|
| 0 | 06:40:58 | 436 | 432 | 2 | 2 |
| 1 | 06:46:29 | 2 | 0 | 0 | 2 |
| 2 | 06:52:00 | 2 | 0 | 1 | 1 |
| 3 | 06:57:01 | 1 | 0 | 1 | 0 |

- Each pass read 0 fetch errors, and each pass came 5 minutes after the pass before it.
- So 0 event pages stayed redirected after the passes. The log holds no line `mtgo event pages stay redirected`.
- The fetched pages add up to 4, the same as the line `meta source` of MTGO: 4 pages, 0 failures, and 0 fetch errors.
- The time budget and the page cap did not hold.
- Only 2 pages answered 302 on pass 0. The run `mtg-meta-4wbpp` read 73 (PR-99).

The owner closed F-183 on the two runs `mtg-meta-4wbpp` and `mtg-meta-b56rq` (D-989).

## The older months (F-179)

The current month 2026-09 read full on the first read, with 413 links and 142 new event slugs of a known format. Each of the 11 older months read empty on the first read. Each empty page held 29,569 to 29,584 bytes and 0 links.

| Step | Time | Pages | Full | Empty |
|---|---|---|---|---|
| First read | 06:04:40 to 06:08:02 | 11 | 0 | 11 |
| Pass 1 | 06:17:48 to 06:21:21 | 11 | 0 | 11 |
| Pass 2 | 06:26:52 to 06:30:53 | 11 | 1 | 10 |
| Pass 3 | 06:36:23 to 06:40:17 | 10 | 1 | 9 |

- Pass 2 read 2026-01 full, with 463 links and 148 new slugs, 22 minutes after the first read.
- Pass 3 read 2026-06 full, with 461 links and 146 new slugs, 31 minutes after the first read.
- At 06:40:18 the job logged `mtgo month pages stay empty` for 9 months after 3 passes.
- The 9 months are 2026-08, 2026-07, 2026-05, 2026-04, 2026-03, 2026-02, 2025-12, 2025-11, and 2025-10.
- So the 436 listed slugs came from 3 months alone: 142 + 148 + 146 = 436. The run `mtg-meta-4wbpp` listed 1,764 slugs from 12 months.

The runs `mtg-meta-5c425` and `mtg-meta-4wbpp` left 0 older months empty (PR-97, PR-99). This run left 9. The owner kept F-179 closed, and this file records the read (D-989).

UNVERIFIED: the effect of the 9 empty months on the model. The store keeps the event pages of earlier runs, and this run found 432 of 436 listed slugs in it. No read counted the lists of the 9 months in the model `20260929T065701Z`.

## The other sources

The line `meta source` of mtgtop8 read 9 pages and 1 failure. This pull request did not read that failure.
