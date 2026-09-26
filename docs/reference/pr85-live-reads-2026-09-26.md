# PR-85: the reads after #236, 2026-09-26

This record holds the reads after the merge of #236 (`1e6f73c`): the deploy, a manual meta job, and two "anime" sessions on the deployed app. It also holds the first runs of the `live-test` skill and of the live sweep (D-958 to D-961).

## The deploy of `1e6f73c`

- The merge changed `go/` and no file of `web/`. So the API build ran alone, and no web build started.
- The `deploy-api` build of `1e6f73c` had a create time of 20:05:00 UTC. It started at 20:05:50 and ended SUCCESS at 20:10:05.
- `/readyz` then read `ok` with version `1e6f73c90392d850d6e99557bdfac2fe7f5bbb71`, and card snapshot 2026-09-26T09:01:56Z.
- The build pointed the job `mtg-meta` at image `worker:1e6f73c`.

## The meta job, run by hand (D-958)

The last scheduled run of `mtg-meta` started at 06:00:10 UTC on 2026-09-26, and it failed the precon bar (F-178). The next scheduled run is 06:00 UTC on 2026-09-27. The owner chose a manual run on the new image (D-958).

| Item | Value |
|---|---|
| Execution | `mtg-meta-8qg4p`, no argument and no env override, as the schedule runs it |
| Start and end | 20:12:15 and 21:32:21 UTC, exit 0 |
| End log | "meta job done", model `20260926T212435Z` |
| Stored model | `meta/model/20260926T212435Z/`, `quality.json.gz` of 20,280,781 bytes and `complete` |
| Card snapshot | `20260926T090156` |
| Lists fitted | 53,431 lists, 1,525 commanders |

### The sources

| Source | Pages | Lists | Fetch errors | Skipped |
|---|---|---|---|---|
| mtgo | 200 | 5,391 | 89 | the page cap of 200 held |
| mtggoldfish | 200 | 82 | 0 | the page cap of 200 held |
| mtgtop8 | 99 | 76 | 0 | none |
| topdeck | 1 | 104 | 0 | none |
| cedhdb | 0 | 0 | 0 | read today |
| edhrec | 0 | 0 | 0 | read on 2026-09-21 |
| mtgjson | 1 | 0 | 0 | the stored table `5.3.0+20260921` names the products |

The MTGO pages that did not fetch answered 302 to `/decklists`. A fetch from this Mac got the same answer, so the site, not an address block, refuses them. The scheduled runs of 2026-09-21 to 2026-09-23 had 17 to 23 such errors each.

UNVERIFIED: the scheduled runs of 2026-09-24 to 2026-09-26 read 0 MTGO pages and 0 errors. The manual run read 200 pages, the newest from March 2026. This record found no cause for the difference.

### The Commander own shares

The job logs no own share. The stored model holds the pairs of fold 0 alone: colors, curve, and lands 132 of 132, and synergy 52 of 68. The bar reads the sum over every fold (`fr.Holdout` in `go/internal/quality/fit.go`).

So a free refit read the same data. It ran `cmd/quality-gate` at `1e6f73c` with `-audit-out`, over a copy of the lists, the commanders, the precons, and snapshot `20260926T090156`. The refit read the counts of the job: 25,987 Commander lists, 9,244 used, 212 immaterial, and a holdout of 15,252. It took 147 seconds, and it called no provider.

| Format | Other own copies | Synergy own copies | Verdict |
|---|---|---|---|
| Commander | 1.00, 572 of 572 (bar 0.95) | 0.80, 210 of 264 (bar 0.75) | PASS |
| Modern | 1.00, 832 of 835 (bar 0.95) | 0.94, 807 of 860 (bar 0.75) | PASS |
| Standard | 1.00, 23 of 23 (bar 0.95) | 1.00, 17 of 17 (bar 0.75) | PASS |

The Commander other copies are colors 180 of 180, curve 196 of 196, and lands 196 of 196. The failed job of the same morning read 205 of 255 synergy copies (F-178).

## The "anime" sessions (D-955)

The message of each session was "Build me an anime-themed commander deck from my collection." Each run used the check account of D-779.

| Lane | Session | Time (UTC) | Turn 1 |
|---|---|---|---|
| `make api-build`, `-max-turns 1` | `q3u6SJljKgGOHcFYEN6x` | 21:43:18 | one question, slot `theme_unmatched` |
| `make live-web` | `SXD19vlkrXKtKjXbtJqR` | 21:48:40 | one question, slot `theme_unmatched` |

Turn 1 of each session asked the theme row alone: "No card I know matches the theme anime-themed. What should the deck do: a creature type, a mechanic, or a play style?" No power row and no color row came with it. D-955 holds on the deployed app.

`scripts/read-session.sh` read both sessions, with the same collection and the same state: status ASKING, one turn, and the slot states format, pool rule, and theme FILLED, and `theme_unmatched` ASKED.

The screen lane showed the question with the thumbs, the "Your answer" field, and "You decide". The screenshot read "Session spend: 1 calls, 2548 in, 128 out, $0.0007".

The first run of `make live-web` stopped before the send. A new browser starts on "Any card", and the lane then read no active collection. The lane now picks the first collection of the account. That run sent no message, and it cost nothing.

## The live sweep (D-961)

`make live-sweep` walks every screen of the deployed web app as the check account, on a desktop of 1280 by 800 and a phone of 390 by 844. Each screen reads console errors, uncaught errors, failed requests, sideways overflow (D-624), and broken images.

| Run | Build | Screens | Faults | Turns | Time | Cost |
|---|---|---|---|---|---|---|
| free 1 | no | 33 | 2 | 0 | 41 s | $0 |
| free 2 | no | 36 | 0 | 0 | 53 s | $0 |
| paid 1 | yes | 41 | 0 | 3 | 131 s | $0.0545, 8 calls |

- Free run 1 found two faults of the sweep itself: the browser logs each expected 404 of an error state as a console error. The sweep now passes that line on an error-state screen alone.
- Free run 1 also showed no binder, because a new browser holds no active collection. The sweep now picks the collection first. It also unfolds the scroll frame of D-364 for each screenshot.
- Paid run 1 sent "Build me a Commander deck from my collection." in session `kFDKDWGakAJDR1od5VUW`. It picked bracket 1 and the commander Syr Konrad, the Grim, and the turn after that ended on deck `ttH2nGfapNo5eE6U0kfz`.

| Turn | Questions | Picks of the sweep |
|---|---|---|
| 1 | the theme, the power bracket, and the colors | You decide, "1 exhibition", You decide |
| 2 | the spending limit, and the commander of three | No budget, Syr Konrad, the Grim |
| 3 | the deck | none |

Turn 1 asked the theme row beside the power and color rows. D-955 holds the row of D-725 alone, the row of a theme that matches no card. The message named no theme, so this turn follows the rules.

## What stays open

- The scheduled run of 2026-09-27 at 06:00 UTC. It must log "meta job done" and store a new model.
- The cause of 0 MTGO pages in the scheduled runs of 2026-09-24 to 2026-09-26.
