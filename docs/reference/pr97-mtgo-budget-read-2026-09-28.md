# PR-97: the read of the MTGO time budget, 2026-09-28

This file holds the reads of PR-97. No paid target ran. No model call ran. Each count comes from the Cloud Run job logs of `decktome-prod`. Each time is UTC.

## The deploy of #248

- #248 merged as `220818b` at 23:21:13 on 2026-09-27.
- The Cloud Build `68d50e1e` of `deploy-api` read SUCCESS at 23:25:59.
- The jobs `mtg-meta` and `mtg-snapshot` run `worker:220818b`.
- The task timeout of `mtg-meta` reads 14,400 seconds (D-982).
- The API serves revision `mtg-api-00095-d4v` on `api:220818b`.
- The schedule `mtg-meta-schedule` reads `0 6 * * *`, ENABLED.

The owner chose the scheduled run of 06:00 on 2026-09-28 over a manual run (D-983).

## The scheduled run

The execution `mtg-meta-5c425` started at 06:00:35 on `worker:220818b`. Its task started at 06:00:52, and the execution ended SUCCESS at 08:09:24. So the task ran 7,712 seconds, 54% of the task timeout of 14,400 seconds. The run also ends inside the old timeout of 9,000 seconds, with 1,288 seconds to spare.

## The month pages

| Step | Time | Pages | Full | Empty |
|---|---|---|---|---|
| First read | 06:03:37 to 06:05:29 | 12 | 1 (2026-07) | 11 |
| Pass 1 | 06:50:38 to 06:53:10 | 10 | 7 | 3 (2025-10, 2025-11, 2026-08) |
| Pass 2 | 06:58:17 to 06:58:33 | 3 | 3 | 0 |

- Each older month read full after pass 1 or pass 2. The job ran no pass 3.
- The current month 2026-09 read empty, with 29,587 bytes and 0 links. The code reads no current month again (D-974), so it read once.
- The two runs of 2026-09-27 read 2026-09 full: 386 links in `mtg-meta-87jlg` and 387 in `mtg-meta-4nps4`.
- So the run fetched no September event of the last day. The store holds the September events of the runs before (F-185).

## The event pages

| Pass | End | Listed | In the store | Fetched | Redirected |
|---|---|---|---|---|---|
| 0 | 07:34:46 | 1,626 | 1,290 | 145 | 55 |
| 1 | 08:00:18 | 55 | 0 | 26 | 29 |

- The first 302 came at 07:00:15, and the last came at 07:59:45. The log holds 84 lines "mtgo event page answered 302".
- Pass 1 came after a wait of 5 minutes, and it read 26 of the 55 redirected pages full.

## The time budget

- The line `mtgo time budget held` came at 08:00:18, with `stage="events again"` and `budget=2h0m0s`.
- The check held before pass 2 of the 302 retry. The wait of pass 2 ends after the deadline, so the pass did not start.
- The line `mtgo event pages stay redirected` names 29 slugs after 1 pass. The next run reads them again.
- The report of the MTGO source reads 171 pages, 4,756 lists, and 29 fetch errors. It names the time budget as the skip reason.

## The end of the run

| Line | Time | Fields |
|---|---|---|
| `quality model fitted` | 08:09:19 | 75,929 lists, 1,510 commanders |
| `quality model stored` | 08:09:20 | `20260928T080018Z`, 20,494,012 bytes |
| `meta job done` | 08:09:20 | `20260928T080018Z` |

- The fit started after the budget held, and it took 9 minutes.
- The fit accuracy reads 0.6742 for Modern, 0.5595 for Commander, and 0.6346 for Standard.

## The other sources

| Source | Pages | Lists | Note |
|---|---|---|---|
| TopDeck | 1 | 348 | One 429, then a retry after 10 seconds |
| EDHREC | 2,089 | 6 | |
| MTGTop8 | 9 | 0 | The cEDH format page holds no event, 1 failure |
| MTGGoldfish | 200 | 46 | The page cap of 200 held |
| MTGJSON | 1 | 0 | The deck list of version 5.3.0+20260927 names the products of the stored table |
| cEDH DB | 1 | 0 | |

## The fix

The owner chose a fix in this pull request over a finding alone and over a note alone (D-984). The current month page reads again when it lists no event, the same as an older month. `TestMTGORetriesAnEmptyCurrentMonth` fails on the code of #248.
