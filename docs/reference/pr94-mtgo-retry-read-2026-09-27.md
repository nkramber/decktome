# PR-94: the read of the MTGO month retry, 2026-09-27

This file holds the reads of PR-94. No paid target ran. No model call ran. Each count comes from the Cloud Run job logs of `decktome-prod`, and from reads on the Mac of the owner. Each time is UTC.

## The deploy of #245

- #245 merged as `9781286` at 07:24:26 on 2026-09-27.
- The Cloud Build `e4988aa0` read SUCCESS at 07:28:39.
- The jobs `mtg-meta` and `mtg-snapshot` run `worker:9781286`.
- The API serves revision `mtg-api-00092-f9g` on `api:9781286`.
- The schedule `mtg-meta-schedule` reads `0 6 * * *`, ENABLED.

The run of 06:00 on 2026-09-27 came before the merge. So the first scheduled run on the new image is at 06:00 on 2026-09-28. The owner chose a manual run in place of that wait (D-975).

## The manual run on Cloud Run

The execution `mtg-meta-hlzf5` started at 07:42:25. It read 12 month pages from 07:44:20 to 07:45:33.

| Month | Bytes | Links | Covered |
|---|---|---|---|
| 2026-09 | 316,513 | 383 | 132 |
| 2026-08 | 366,146 | 451 | 155 |
| 2026-07 | 377,346 | 467 | 158 |
| 2026-06 | 372,737 | 461 | 146 |
| 2026-05 | 369,533 | 457 | 150 |
| 2026-04 | 362,752 | 447 | 138 |
| 2026-03 | 412,285 | 513 | 161 |
| 2026-02 | 343,813 | 419 | 133 |
| 2026-01 | 376,320 | 463 | 148 |
| 2025-12 | 362,923 | 445 | 149 |
| 2025-11 | 332,700 | 405 | 139 |
| 2025-10 | 342,848 | 419 | 149 |

No month page was empty. So the run started no retry, and it logged no `mtgo month page again` line. The pages came 3 to 12 seconds apart. On the run of 06:00 on 2026-09-27, the empty pages came about 30 seconds apart.

The same Cloud Run job got the full lists at 07:44. So the address of the job does not cause the empty list by itself.

## The forced retry on the Mac

The owner chose a forced run on the Mac over a flag in the production job (D-976). A scratch test in the package `meta` runs `runMTGO`, then `retryMTGOMonths`, against the real site.

- A wrapper of the HTTP transport removes each event link from the first read of each older month.
- Each later read of that month comes from the real site, with no change.
- The test reads two months, with a page cap of 3 and `MTGORetryWait` of 1 minute.
- The test marks each event of the current month as stored, so the page cap stays for the retried month.
- The test writes to a local folder under `.local`, and not to the store of `decktome-prod`.

The first run used the code of #245.

| Time | Log line | Month | Bytes | Links |
|---|---|---|---|---|
| 07:52:48 | `mtgo month page` | 2026-09 | 316,513 | 383 |
| 07:52:57 | `mtgo month page`, forced | 2026-08 | 367,048 | 0 |
| 07:53:57 | `mtgo month page again`, after=1m0s | 2026-08 | 29,578 | 0 |

The real retry read an empty page. Its size is the size of the empty 2026-08 page of the run of 06:00 on 2026-09-27. So the site served the fault to the Mac, at 07:53, one minute after a full read.

## The probes from the Mac

The session then read three month pages from the Mac with curl and the user agent of the job.

Probe 1 ran from 07:54:09 to 07:57:39, with six rounds.

| Month | Empty reads | Full reads |
|---|---|---|
| 2026-08 | 3 | 3 |
| 2026-07 | 1 | 5 |
| 2026-09 | 0 | 6 |

- Each empty read held 29,570 to 29,576 bytes and no event link.
- Full and empty reads of 2026-08 came 19 to 48 seconds apart.
- Probe 1 did not time each read. The gaps of its log show that three of the four empty reads took about 30 seconds, and one took about 1 second.

Probe 2 ran from 07:58:02 to 07:58:44. It read 2026-06, 2026-07, and 2026-08 ten times. Each read was full, and each took 0.5 to 7.3 seconds.

## What the reads prove

- The site serves an empty older month page at times, to any client. The Mac and Cloud Run both got it.
- The fault comes in spells. Full and empty reads of one page alternate inside a spell.
- The current month read full on each read of this session.
- One retry can get the same empty page. The retry of 07:53:57 did.

So the one retry of #245 lowers the fault, but it does not end it. The owner chose a stronger retry in this pull request (D-977).

## The change of this pull request

The job reads each empty older month page again up to 3 times (D-977).

- Pass 1 comes after the other sources, at least `MTGORetryWait` after the first read, as in #245.
- Each later pass comes at least `MTGORetryWait` after the reads of the pass before it. The default is 5 minutes.
- Each pass reads only the months that are still empty.
- A month page that did not fetch reads again in the next pass.
- The job fetches the event pages of each pass after the pass.
- The log line `mtgo month page again` names the pass, and `after` is the time since the first read.
- A page that is empty again stays under `meta/raw/mtgo-month-empty/`, with `-again1`, `-again2`, or `-again3` at the end of its name.
- After the last pass, the warning `mtgo month pages stay empty` names each month that stayed empty.

A second forced run on the Mac used the new code. The first read of 2026-08 at 08:18:24 lost its links. Pass 1 read the real page at 08:19:53, after=1m29s, with 451 links and 155 covered. So the real site answered pass 1 full, and passes 2 and 3 did not run. `TestMTGORetriesAnEmptyOlderMonth` covers them with a fake site.
