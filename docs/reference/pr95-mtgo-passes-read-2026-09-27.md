# PR-95: the read of the 3 MTGO month passes, 2026-09-27

This file holds the reads of PR-95. No paid target ran. No model call ran. Each count comes from the Cloud Run job logs of `decktome-prod`, and from reads on the Mac of the owner. Each time is UTC.

## The deploy of #246

- #246 merged as `0414f27` at 16:26:06 on 2026-09-27.
- The Cloud Build `245c2d7b` read SUCCESS at 16:31:13.
- The jobs `mtg-meta` and `mtg-snapshot` run `worker:0414f27`.
- The API serves revision `mtg-api-00093-rx2` on `api:0414f27`.
- The schedule `mtg-meta-schedule` reads `0 6 * * *`, ENABLED.

The first scheduled run on the new image comes at 06:00 on 2026-09-28, about 13.5 hours after the deploy. The owner chose a manual run in place of that wait (D-978).

## The end of the manual run of PR-94

The execution `mtg-meta-hlzf5` ran `worker:9781286` from 07:42:34 to 09:29:44, and it succeeded.

- The line `mtgo event slugs` at 09:11:00 read listed=1758, stored=830, fetched=200, and fetch_errors=134.
- The line `meta job done` at 09:29:41 named the model `20260927T092053Z`.
- The run logged no `mtgo month page again` line and no `mtgo month pages stay empty` line.

## The manual run on the new image

The execution `mtg-meta-87jlg` ran `worker:0414f27` from 16:36:35 to 18:30:01, and it succeeded. The line `meta job done` at 18:29:56 named the model `20260927T182135Z`.

The first read of the 12 month pages came from 16:40:49 to 16:43:50.

| Month | Time | Bytes | Links | Covered |
|---|---|---|---|---|
| 2026-09 | 16:40:49 | 318,780 | 386 | 133 |
| 2026-08 | 16:41:19 | 29,578 | 0 | 0 |
| 2026-07 | 16:41:34 | 377,346 | 467 | 158 |
| 2026-06 | 16:42:04 | 29,572 | 0 | 0 |
| 2026-05 | 16:42:34 | 29,569 | 0 | 0 |
| 2026-04 | 16:42:41 | 362,752 | 447 | 138 |
| 2026-03 | 16:43:00 | 412,285 | 513 | 161 |
| 2026-02 | 16:43:30 | 29,584 | 0 | 0 |
| 2026-01 | 16:43:35 | 376,320 | 463 | 148 |
| 2025-12 | 16:43:38 | 362,923 | 445 | 149 |
| 2025-11 | 16:43:46 | 332,700 | 405 | 139 |
| 2025-10 | 16:43:50 | 342,848 | 419 | 149 |

Four older month pages read empty. Each empty read came 30 seconds after the line before it. The full reads came 3 to 19 seconds after the line before them. PR-94 found the same time on three of the four empty reads of its Mac probe.

## The event pages of the first pass

The line `mtgo event slugs` at 18:05:44 read listed=1175, stored=549, fetched=200, and fetch_errors=141.

- The 8 full months gave 1175 covered slugs.
- The job fetched 200 event pages, so the page cap of 200 held.
- Each of the 141 fetch errors was a 302 answer.
- The fetch took 82 minutes, from 16:43:50 to 18:05:44, for 341 requests. That is about 14 seconds a request.

## The retry passes

| Time | Pass | Month | After | Bytes | Links | Covered |
|---|---|---|---|---|---|---|
| 18:15:37 | 1 | 2026-08 | 1h34m18s | 366,146 | 451 | 155 |
| 18:15:47 | 1 | 2026-06 | 1h33m43s | 372,737 | 461 | 146 |
| 18:15:55 | 1 | 2026-05 | 1h33m21s | 369,533 | 457 | 150 |
| 18:16:25 | 1 | 2026-02 | 1h32m55s | 29,584 | 0 | 0 |
| 18:21:35 | 2 | 2026-02 | 1h38m5s | 343,813 | 419 | 133 |

- Pass 1 read 3 of the 4 months full. The empty read of 2026-02 came 30 seconds after the line before it.
- The line `mtgo month pages wait` at 18:16:25 read months=1, pass=2, and wait=5m0s.
- Pass 2 read 2026-02 full, so pass 3 did not run.
- The run logged no `mtgo month pages stay empty` line.

So the 3 passes of #246 got each month page. The one retry of #245 would have left 2026-02 empty in this run.

## The event pages of the retry

- After pass 1, the line `mtgo event slugs` read listed=451, stored=0, fetched=0, and fetch_errors=0.
- After pass 2, the same line read listed=133, stored=0, fetched=0, and fetch_errors=0.

The page cap held before the retry. So the loop stopped before the first slug, and the 584 new slugs fetched no page. Each slug of the retry came after each slug of the first pass, 2025-10 included. So the fetch did not go newest month first, as the comment of `MaxPages` said (F-183).

## The 302 answers

- `mtg-meta-87jlg` logged 141 slugs with a 302, and `mtg-meta-hlzf5` logged 134.
- Only 8 slugs occur in both runs.
- After the run, the Mac read the first 3 of these slugs, with the user agent of a browser. Each answered 200, with 334,177, 34,189, and 349,528 bytes.

So a 302 is most often a fault of the site at that time, as the empty month page is. The comment of `NewFetcher` said the site answers 302 for an event page that it no longer serves. That is true of few of these slugs.

## The time of a 302 answer

The 141 lines of the 302 answers leave 140 gaps. 52 gaps are less than 5 seconds, and 41 gaps are about 30 seconds. So a 302 answer took 0 to 30 seconds. Each empty month page took 30 seconds too.

- The first fetch of the new code stops at 200 requests, because each 302 page holds a place under the cap. At this run, that is about 47 minutes.
- Pass 1 then reads about 80 pages again, in at most about 40 minutes.
- So the new code needs about the time of this run, 1 hour 54 minutes. The task timeout is 9,000 seconds.
- The old code read past each 302 answer until it fetched 200 pages. So only the count of slugs, 1175 in this run, limited a spell of 302 answers. The new code reads at most 4 times each of at most 200 slugs.

A spell of 302 answers to each page for hours can still pass the task timeout. That risk stays open.

## The Mac reads of the month pages

After the run, the Mac read 2026-08, 2026-06, 2026-05, and 2026-02 twice each, with the user agent of the job. Each read was full, with the links of the retry passes. Each took 1.0 to 16.4 seconds.

## The change of this pull request

The owner chose to fetch the event pages after the month retry, newest month first, under one page cap (D-979). The owner also chose a retry of each 302 event page (D-980).

- The job reads the 12 month pages first, then the other sources, then the retry passes of the month pages.
- The job then fetches the event pages of every month, newest month first, up to `MaxPages`.
- An event page that answers 302 holds its place under the cap. It reads again up to `MTGORetryPasses` times, each pass at least `MTGORetryWait` after the fetch before it.
- The line `mtgo event page answered 302` names the slug and the pass. The line `mtgo event slugs` names the pass 0 and the count of `redirected` pages.
- The line `mtgo event pages again` names each retry pass. The warning `mtgo event pages stay redirected` counts the pages after the last pass.
- A page that stays redirected counts as one fetch error. The next run reads it again.

`TestMTGOFetchesNewestMonthFirst` fails when the job fetches the event pages before the retry. `TestMTGORetriesARedirectedEventPage` fails when the job reads no 302 page again.
