# PR-98: the read of the current MTGO month, 2026-09-28

This file holds the reads of PR-98. No paid target ran. No model call ran. Each count comes from the Cloud Run job logs of `decktome-prod`. Each time is UTC.

## The deploy of #249

- #249 merged as `bdc5b60`.
- The Cloud Build `0dfd1851` of the deploy trigger started at 14:16:44, and it read SUCCESS at 14:20:59.
- The jobs `mtg-meta` and `mtg-snapshot` run `worker:bdc5b60`.
- The task timeout of `mtg-meta` reads 14,400 seconds.
- The API serves revision `mtg-api-00096-ldb` on `api:bdc5b60`.

The owner chose a manual run over the scheduled run of 06:00 on 2026-09-29 (D-985).

## The manual run

The execution `mtg-meta-4wbpp` started at 16:27:21 on `worker:bdc5b60`. This file reads its month pages alone. The run was still in progress at the time of this read, at 16:46.

## The current month

The first read of 2026-09 came at 16:29:29. It read 330,023 bytes, 401 links, and 138 covered events.

- The month read full on the first read. So the job did not start the retry of an empty current month, and the log holds no line `mtgo month page again` for 2026-09.
- The runs before read 2026-09 as follows: 386 links in `mtg-meta-87jlg`, 387 in `mtg-meta-4nps4`, and 0 in `mtg-meta-5c425` (F-185).
- The retry of an empty current month has the proof of `TestMTGORetriesAnEmptyCurrentMonth` alone. No run on the image of #249 read an empty current month yet.

The owner chose to close F-185 on a full first read (D-985).

## The older months

| Step | Time | Pages | Full | Empty |
|---|---|---|---|---|
| First read | 16:29:29 to 16:33:31 | 12 | 3 (2026-09, 2026-07, 2026-03) | 9 |
| Pass 1 | 16:44:17 to 16:46:05 | 9 | 6 | 3 (2026-08, 2026-04, 2026-02) |

- Each empty page holds 29,569 to 29,584 bytes and 0 links. Each full page holds 330,023 to 412,285 bytes.
- Pass 1 read each empty month again after 12 to 15 minutes.
- The line `mtgo month pages wait` came at 16:46:05, with 3 months, pass 2, and a wait of 4m9s.
- The older months follow the path of F-179 and D-974. PR-98 changes no part of that path.
