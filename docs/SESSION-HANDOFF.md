# Session hand-off

`CLAUDE.md` is the entry point. It sends you here. Then read `docs/decisions.md` from D-297, `docs/owner-questions.md`, and `docs/open-questions.md`. The session narratives of 2026-08-23 to 2026-08-28 sit in `docs/reference/session-log-2026-08-23-to-26.md` and `docs/reference/session-log-2026-08-27-to-28.md`. The audit of 2026-08-29 is `docs/audit-2026-08-29.md`.

CAUTION: the web tests need Node 22.23.2 (`.nvmrc`). Under Node 20 every test file fails at start with `ERR_REQUIRE_ESM` from jsdom 30. On the owner's machine `~/.nvm/versions/node/v22.23.2/bin` on the PATH fixes it.

## Where things stand (2026-09-02)

- `main` is at `13ca8dd`. Merged: PR-0a to PR-8, PR-7B, PR-10 to PR-13, the audits, the Phase 3B roadmap (#46), PR-16 (#47), PR-16B (#48), PR-17 (#49), PR-17B (#50), PR-18 (#53), and the review fixes of PR-18 (#54).
- Branch `pr-19` holds the chat and build experience (D-432 to D-453), committed and pushed, with its PR open. It also holds the land-swap fix of the revision turn (F-31, D-448) and the split land bucket of the shortlist (F-32, D-450). Revise gate run 6 and deck gate run 11 are due before the merge.
- The tree is green on `nits-and-fixes`: Go build, vet, `-race` tests, golangci-lint, web lint, typecheck, and 219 web tests. The emulator tests of the collection store pass, and `make lint` reports zero findings.
- Question gate run 31 passes every bar. The set deck gate passes 6 of 6.
- Every gate stands and passes: question gate 32, the set deck gate, deck gate 11, and revise gate 7.
- PR-9 is out of the MVP (D-256). Phase 3B comes before Phase 4 (D-316).

CAUTION: branch `pr-17` carries eight concerns. They are the contract, the Go side, the reference design, and the layout of D-331. They are also the Build menu, the one deck screen, the pool picker, and the speed of the app. Guardrail 10 asks for one. The owner chose to ship it whole (D-344).

## The numbers, and why none of them compare with `main` now

The audit changed two prompts (D-302). The generate prompt is at version 10 and the classify prompt at version 14. The shortlist omits the commander, the house format has a sideboard sentence, and the dead `two_plans` and `owned_mode` facts left the classify schema. So the baselines below measure the code before the audit, and the next run re-baselines (D-66).

| Measure | Baseline | Where measured |
|---|---|---|
| Question gate | PASS 25 of 27 counted, 2 invented, 0 premature, 0 lint, $0.16 | run 27, 2026-08-28, ask prompt 13, catalog of D-290 and D-294 |
| Question eval | 19 bad of 382, holdout 8 of 111 (7.2 percent) | eval of run 25, 5 conversations unjudged |
| Deck gate | PASS 18 of 18, 2 repairs, $1.09 | run 8, 2026-08-28, generate prompt 9 |
| Revise gate | PASS 6 of 6, $0.29 | run 2, 2026-08-28 (D-296) |
| Loop | off since 2026-08-26 | seven starts, nothing kept |

CAUTION: `tune-check` paired zero questions between run 24 and run 25, because the catalog and the prompt changed. The paired guard says nothing across that line, and the whole-run margins carry the verdict. The eval leaves a conversation unjudged when the judge returns fewer verdicts than questions. So 382 is the honest count, not a drop from 435.

## The audit of 2026-08-29, what changed

`docs/audit-2026-08-29.md` holds the report. The points a session needs first:

- A turn during a build gets `CodeAborted` with "a build is in progress". The build and its store writes run detached from the client, so a disconnect keeps the deck (D-303). The web app warns before it leaves a page mid-build.
- Every session, deck, and collection id passes `gzstore.ValidID`. An id with a slash is `CodeInvalidArgument`.
- `go/internal/gzstore` holds the gzip and JSON helpers of the three Firestore repos, with one inflate limit.
- The buy cost and the deck cost count the commanders and sum copies per Oracle id. A locked, kept, or commander card above the mana cap stays in the pool (`Revision.Exempt`).
- The house format offers paper cards only (D-306).
- `deck-gate` and `revise-gate` exit 1 on FAIL. A nil cost prints `unpriced` in every gate. `revise-gate` has `-only` and `-dry`. `questions-eval` and `candidates-review` refuse an existing output (D-65).
- The Makefile runs bash with `pipefail`. `candidates-review` guards on a filled score cell.
- `autotune.sh` builds `tune-check` after the branch switch, never switches in a dry run, and restores the start branch on exit.
- The go job checks out with `fetch-depth: 0`, so `make llm-defaults-check` finds a merge base. A weekly `vuln` job runs govulncheck alone on Monday 06:00 UTC (D-305).
- The STE checker flags passive voice, modals and perfect tenses, -ing forms, and the 20-word step limit (D-304). Dated records are exempt: `docs/reference/pr[0-9]*`, the session logs, `docs/audit-*`, and testdata.
- Every code comment states a rule and cites a decision id. No comment carries a date, a session id, a run number, or "the owner".

## PR-13, what it holds (2026-08-29)

- `DeckService.ExportDeck(deck_id, format)` returns the text and a file name. `EXPORT_FORMAT_ARENA_TEXT` is the ManaBox shape, and `EXPORT_FORMAT_BUY_LIST_TEXT` is one "count name" line per card to buy (D-309).
- `go/internal/export`: `ArenaText`, `BuyList`, `BuyListText`, `FileName`, `Render`. The Arena line names the owned printing when the card is owned, else the default paper printing (D-307). The line carries the full card name, and the index resolves it without ambiguity. `TestArenaTextRoundTrip` is the gate.
- The buy list is the shortfall of the commander, the main deck, and the sideboard, summed per Oracle id, and the upgrades apart (D-308). A commander with no entry in the card list counts as one card to buy.
- `web/apps/web/src/features/export`: `export-panel.tsx` (four buttons, the status line, the buy list with a Scryfall link per card) and `buy-list.ts` (the same rules as the Go package, for the screen). The deck view mounts the panel under its header. `export` is a leaf feature, and `deck` imports it.
- The web app reads the text from the API, so the copy and the file match what the API renders.

## PR-16, what it holds (2026-08-29)

CAUTION: three parts of this slice changed on 2026-08-30. The sidebar became a top bar (D-328), the light theme went (D-330), and the color identity of a deck went (D-329). The primitives, the route split, and the test helper below still hold.

- Nine primitives in `web/apps/web/src/components/ui`: Button, Input, Label, Textarea, Checkbox, Card, Skeleton, DropdownMenu, and the toast. Each one is a hand-written shadcn shape on Radix, with relative imports. The shadcn command line installs a `@` alias, and the import boundary of this repo reads relative paths only.
- `src/styles/tokens.css` holds the tokens. The dark theme overrides the neutrals, the surfaces, and the link. The six mana colors do not change. A script in `index.html` paints the class before the first paint.
- The shell is a sidebar on a desktop and a bottom tab bar on a phone. A media query picks one of the two. Two navigations with one name fail the axe landmark-unique rule.
- Sign-out moved into the account menu. A failed sign-out reports as a toast.
- Five things leave the first paint: `firebase/auth`, the Connect client, the five pages, the two shell menus, and the toast host.
- `renderAt` in `src/test-utils.tsx` is async now, and every test awaits it. It warms the page modules and the auth SDK, then flushes one act.
- `src/test-setup.ts` adds four jsdom stubs that the Radix menus need.

## PR-17, what the branch holds (2026-08-30)

The web side, in one list:

- The library grid at `/decks`: the name, the commander, the mana pips, the format, the power, the count, the cost, the date, and a favorite star. Search by name and commander, and filter by format, power, and favorites. Every filter runs on the server.
- `/decks/<id>` is the one screen of a deck (D-335). It holds the deck, the actions, and the conversation that built it.
- `/session/<id>` holds a build with no deck. It hands the reader to the deck's address the moment a turn ends with a deck.
- `src/features/workspace` is the one feature with a path to both chat and deck.
- Build in the header is a menu of the collections (D-332), and a signed-in reader lands there (D-334).
- The look of the reference design: three faces, the navy and gold palette, one top bar, dark alone.
- A "Build from" picker sits under the title of a new chat (D-336). It names the pool, and it changes it without a trip to another page.
- The deck screen carries the version history and the compare (D-343). The decks of one chat are the versions of one deck.

CAUTION: no chunk of this app loads behind a Suspense boundary (D-338). React holds a committed fallback for 300 ms, and it holds every later reveal with it. `src/app/deferred.tsx` replaces `React.lazy` everywhere. Do not put `React.lazy` back.


The Go side, in one list:

- `DeckService.UpdateDeck` writes the name and the favorite mark, and `DeleteDeck` removes a deck for good. Both are additive, and `buf breaking` passes.
- `Deck.favorite` and `Deck.card_count` are new. The list view carries no cards, so it sets `card_count`. The deck list on screen reads `cards.length` today and always shows zero.
- `ListDecks` takes `page_size`, `page_token`, `format`, `favorite`, `query`, `power_bracket`, `power_sixty_step` (D-324), and `session_id` (D-343).
- The filter runs in Go over the rows Firestore returns, not as a query. One read serves every filter, and no composite index has to exist. The scan cap is 500 rows.
- The stored document gains six flat fields. A deck written before PR-17 reads them as zero, and a rename fills them.
- `internal/decks/filter_test.go` is new, and the emulator tests cover Update, Delete, and the filter. `make store-check` passes.

CAUTION: a `t.Cleanup` can not delete from Firestore. Go cancels `t.Context` before a cleanup runs, so the delete fails and the next run reads the leftovers. Each emulator test takes a fresh user id instead.

## The look of 2026-08-30 (D-328 to D-331)

The owner gave a reference design, and it settles the look. Our terms win over its terms: the app is the MtG Deck Builder, and its entries are Build, Decks, and Collection.

- Three faces, three jobs. Cinzel engraves a heading, a label, and a button. Crimson Pro reads a paragraph. JetBrains Mono carries an id, a count, and a date. Each one ships with the build.
- The palette is navy, gold, and purple, with parchment for text. The radius is 4 px, and every panel carries a faint gold hatch.
- The shell is one top bar (D-328). The sidebar, the phone tab bar, and the theme toggle are gone.
- Dark is the only theme (D-330). `src/lib/theme.ts` and the no-flash script left with it.
- A deck has one screen and one address (D-335). `/decks/<id>` holds the deck, its actions, and its conversation. `/session/<id>` holds a build with no deck, and it hands over the moment a deck exists.
- `src/features/workspace` is the one feature with a path to both chat and deck. The lint carries that rule.
- A deck owns the whole page, and the conversation docks at the bottom left (D-331). A History control opens the thread over the composer. Before a deck exists, the conversation is the page.
- The identity wash left (D-329). Color of the game shows in a mana pip and a rarity dot.
- Build in the header is a menu of the collections (D-332). It sets the pool and opens a chat. With no collection it offers the way to the upload screen (D-334).
- A signed-in reader lands on Build, not on the collection (D-334).

CAUTION: a trigger of a Radix menu must pass on every prop it takes. The ref is among them, and Radix measures the trigger through it to place the panel.

A trigger that keeps only the props it names drops the ref. The panel then lands at the top left corner, outside the window, and a click never reaches it. Both menus of the shell carried this defect. jsdom has no layout, so no test there sees it. Playwright found it in one run.
- The binder head sits under the controls, the buy list opens on request, and a page holds 75 rem (D-333).

A session reads its own work with Playwright, and it measures rather than looks. Four checks ran on 2026-08-30.

- The deck page holds 80 percent of a 1440-pixel screen.
- The buy list carries no `open` attribute.
- The Continue-to-chat control does not move when a reader picks an upload.
- The Build menu lists every collection.

## The palette of D-327 (2026-08-29), now amended

D-311 kept every surface neutral and let the card art carry the color. The app read as boring, and the owner said so twice. The five colors of the game are the palette now, and dark leads.

- `src/styles/tokens.css` holds the dark base and a `.light` class. Dark is the class-free base, so a light reader carries the class.
- `src/features/deck/color-identity.ts` reads the color identity of a deck from its commander. `identityVars` sets `--identity-a` and `--identity-b` on one element, and the `identity-wash` and `identity-rule` utilities read them.
- Each card role carries its own hue through `roleToken`.
- `ManaPips` shows an identity as pips, and its screen-reader label names the colors in words.
- The ground of every page carries two soft lights and a fine grain. A panel takes a hairline of its own light along its top edge.
- The collection screen shows the binder: the count, the unique cards, the rarity spread, and the art of the rarest ten cards. It came forward from PR-18.

CAUTION: the binder head calls `GetCollection`, and the answer carries every entry. The owner's export holds 2,657 rows and 4,952 cards, so one page load moves about one megabyte. PR-18 stores the summary on the collection, and the head then reads no entry at all (D-392).

## Reading the app without the owner (2026-08-30)

`@playwright/test` is a dev dependency of `web/apps/web` now, and Chromium sits in the local cache. A session can read its own work.

- Start the Auth emulator and the Vite dev server. Then drive Chromium with a script that answers every `/mtg.v1.*` call from canned JSON, so no backend has to run.
- A Connect Timestamp is an RFC 3339 string in JSON, not a `{seconds}` object. A `{seconds}` stub fails with "cannot decode message google.protobuf.Timestamp".
- The script lives outside the repo, in the session scratchpad. It must run from `web/apps/web`, or the bare import of `@playwright/test` does not resolve.
- PR-23 holds the real smoke flow (D-313). This is a reading tool, not that.

## PR-17B, the set filter (F-29, D-373 to D-383)

Branch `pr-17b` holds it. The app never applied a set as a constraint, and a deck asked for one set held cards of any set.

The eight decisions, in short:

- D-376: a set name resolves to a whole set family, through the Scryfall parent link. "The Hobbit" gives `hob` and `hoc`. A phrase that names two base sets asks.
- D-377: the snapshot carries a fourth file, `sets.json.gz`, from the `/sets` endpoint. No bulk card file holds `parent_set_code`.
- D-378: a set limit never filters basic lands.
- D-379: inside a set limit the theme ranks the shortlist, and neither the theme cut nor the role caps apply.
- D-380: a family under 70 nonbasic cards in the deck colors builds no deck. The floor is 35 for a 60-card format.
- D-381: a card the reader named by name beats the set limit, and the deck marks it.
- D-382: ramp cards and nonbasic lands come from outside the sets only after the reader says yes.
- D-383: `DeckCard.outside_requested_sets` carries the mark, and the deck screen shows it in red.

CAUTION: a card does not have one set. The snapshot of 2026-08-31 holds 988 paper sets over 34,599 playable Oracle cards, and 16,247 of them hold printings in two or more. `Card.set_codes` is a list, 80,193 pairs in all, about two megabytes.

CAUTION: a name prefix can not build a set family. `ltc` is "Tales of Middle-earth Commander" and its base set is "The Lord of the Rings: Tales of Middle-earth". The two names share no prefix, and only `parent_set_code` links them.

CAUTION: the roadmap gate line named two cards that do not exist. The snapshot holds "Thranduil, the Elvenking" and "Smaug, Wicked Worm", both with a comma. Smaug the Impenetrable is in `hoc`, not in `hob`, so only the family rule of D-376 satisfies that line.

What the branch changed, in one list:

- `internal/cards/sets.go` is new: the set table, the family walk, and the resolver. `internal/cards/index.go` fills `Card.set_codes` in the printings walk it already runs, and it shares one string per set code.
- `scryfall.Client.Sets` reads `/sets`. `cards.Refresh` stores the file, and `cards.BackfillSets` fills a snapshot stored before the file existed. The worker calls it every cycle.
- `candidates.Request` takes `SetCodes` and `OutsideRoles`. `CountInSets` and `SetFloor` answer the viability floor, and `CountManaInSets` answers the mana row.
- `generate.Request.SetCodes` marks every deck card the sets do not hold, and one warning counts them.
- `agentsvc` runs the floor twice, before the commander pool and after the colors settle. `ErrThinSet` carries the counts.
- The classify prompt is at version 15 and reports `set_names`. The catalog holds two new rows. The snapshot is at version 3.
- `web/apps/web/src/features/deck/card-tile.tsx` shows the red mark.

CAUTION: the role caps were the second cut, and the first dry run found it. The "other" cap of 10 cut 40 of the 128 cards the Hobbit family offers in black-red. The shortlist reached 61, and a Commander deck needs 99. Both cuts lift under a set limit now.

## The paid runs of PR-17B (2026-08-31)

| Run | File | Result |
|---|---|---|
| Set deck gate 1 | `pr17b-set-gate-run1.md` | PASS. 6 of 6 decks, 0 blocks, 0 notes, 0 repairs. $0.3539 over 306 seconds. |
| Question gate 29 | `pr7-question-gate-run29.md` | FAIL. 29 of 30 catalog-only, over the bar of 25. Two failures, and neither comes from the set filter. $0.16. |
| Question eval 29 | `pr7-question-eval-run29.md` | 11.2 percent bad on the holdout, from 6.8. Read the caution below before you act on that number. $0.09. |
| Question gate 30 | `pr7-question-gate-run30.md` | FAIL. 27 of 27 catalog-only, 0 lint, 0 dead ends. 7 premature, which found two more defects. $0.16. |
| Question gate 31 | `pr7-question-gate-run31.md` | **PASS.** 27 of 27 catalog-only, 0 lint, 0 dead ends, 0 premature. $0.16. |
| Question eval 31 | `pr7-question-eval-run31.md` | 7.9 percent bad on the holdout, from 6.8 on run 28. $0.09. |
| Question gate 32 | `pr7-question-gate-run32.md` | **PASS.** 27 of 27 catalog-only, 0 lint, 0 dead ends, 0 premature. Prompt version 16. $0.16. |
| Question eval 32 | `pr7-question-eval-run32.md` | 22 bad of 393, from 27 of 413. The three changed rows fall from 11 findings to 4. $0.09. |
| Revise gate 3 | `pr12b-revise-gate-run3.md` | FAIL, 7 of 8. It found D-391. $0.57. |
| Revise gate 4 | `pr12b-revise-gate-run4.md` | **PASS, 8 of 8.** $0.54. |
| Revise gate 5 | `pr12b-revise-gate-run5.md` | FAIL, 10 of 11. It found F-32. $0.81, 9 minutes, 12 turns. |
| Revise gate 6 | `pr12b-revise-gate-run6.md` | FAIL, 9 of 10. Revision 9 passed with 14 basics swapped and no repair turn. Revision 1 skipped its question, a variance the prompt now closes. $0.67. |
| Revise gate 7 | `pr12b-revise-gate-run7.md` | **PASS, 11 of 11.** Every land turn asked, then swapped. $0.74, 12.5 minutes. |
| Deck gate 11 | `pr8-deck-gate-run11.md` | **PASS, 24 of 24.** Prompt version 11, the split land bucket. 3 repair turns. $1.46, 22 minutes. It found F-33. |

The revise gate gained a set-limited base on 2026-09-01. A revision reads the same slots as the build, so base 3 is the one run that proves the set filter survives a revision. It does. The revised deck holds Arcane Signet, Delighted Halfling, and Elvish Mystic, and all three are in The Hobbit Eternal. It holds no Sol Ring, which is in neither set.

Run 3 failed on one revision, and the failure had nothing to do with the sets. The model removed six cards and added seven under a mana cap, and the engine blocked the whole deck for one card. D-391 adds the trim, and run 4 passed.

## The question-quality pass of 2026-09-01 (D-387 to D-389)

Eval run 31 found 10 warranted defects over 413 questions, and 6 sat on the budget, the colors, and the format rows. Seven changes answer them, and run 32 measures the result.

| Measure | Run 31 | Run 32 |
|---|---|---|
| Questions scored | 413 | 393 |
| Not warranted | 27 | 22 |
| Detailed findings on budget, colors, and format | 11 | 4 |
| budget | 6 | 1 |
| format | 2 | 0 |
| colors | 3 | 3 |

Gate run 32 passes every bar: 27 of 27 catalog-only, 0 lint, 0 dead ends, 0 premature.

The three colors findings that remain are ones this repo refuses, and each cites a rule that predates them.

- "Whatever is winning" is not a color delegation. The corpus routes "whatever" nowhere, after gate runs 11 to 13 read it as house rules six times (D-111).
- The Atraxa question fires on turn 1 of a conversation whose whole point is a misspelling. The reader wrote "Atraxa, Praetor's Voice", and the card is "Atraxa, Praetors' Voice". The index does not hold the first spelling, so the rule of D-388 claims nothing about it. Guardrail 4 and D-140 both say the engine states nothing it can not prove. The reader corrects the name on turn 2, and the row does not fire again.
- The duplicate color question of conversation 7 follows a format decline. The reader named Pioneer, the decline row retired every open question (D-125), and a retired row asks once more (D-195). That is by design.

CAUTION: the whole-run bad count drifts. Run 28 found 14 of 391, run 29 found 22 of 398, and run 31 found 27 of 413. The step from run 28 to run 31 is 13 questions. D-230 measures the judge noise at up to 9 between two runs of identical code. Each single step sits inside that band, and the sum does not. Watch it on the next run.

The drift belongs to no part of this branch. Not one bad question sits on `set_unresolved`, on `set_outside_mana`, or on `power_sixty`. Every one sits on `budget`, `colors`, `format`, `power_commander`, `theme`, `pool_thin`, `named_card_role`, or `format_unsupported_open`. `budget` leads every eval: 3 of 6 detailed findings on run 28, and 6 of 17 on run 31. The gate is the contract, and the eval is advisory (D-136).

Three runs measure the net of D-351, because run 29 ran without it and the other two ran with it. The table below is the whole case for D-386.

| Run | The net | Conversations that ended before the script did | Questions the reader answered | Turns the net healed |
|---|---|---|---|---|
| 29 | none | 6 | 316 | 0 |
| 30 | closes on the reader's first reply | 45 | 264 | 82 |
| 31 | closes after two turns (D-386) | 11 | 308 | 34 |

The grace period returned 44 of the 52 answers run 30 threw away. It also kept every gain: run 31 reports no dead end.

Run 30 found two more defects beyond the eager net. The gate asked for the `commander` key while `commander_pick` answered it, which called 6 conversations premature. Each of the six holds a commander. `required` reads all four commander keys now.

CAUTION: the eval numbers do not compare cleanly. Run 28 scored 391 questions and found 14 not warranted. Run 29 scored 398 and found 22. The margin is 8 questions, and D-230 measures the judge noise at up to 9 between two runs of identical code. `tune-check` paired zero questions across the line, because the catalog and the prompt both changed. So the run is neither proof of a regression nor proof of none.

The three set conversations reach neither the bad-question list nor the missing-question list of the eval. Every one of the 8 extra findings sits on a `commander_pick`, `colors`, or `budget` row of a conversation that predates this branch.

The set gate proves every deck line. A session read every card of all six decks against the raw snapshot. Only two decks hold a card outside their family, and the app marked every one. Prompt 21 holds 8 lands of the mana fill, and prompt 22 holds Sol Ring. Prompts 19, 20, 23, and 24 hold none.

A free dry run proves the thin-set refusal of D-380. A mono-black Hobbit commander gives 68 cards, the floor is 70, and the turn ends with the count and no deck.

The owner settled every failure of run 29 and run 30 (D-384, D-385, D-386). Run 31 passes every bar.

- The 60-card power row is fixed now, so the ask role sends it word for word. The acronym rule reads the options beside the text, because a reader meets both at once.
- The gate runs the D-351 net, as `agentsvc.Chat` does. A test reads both turn loops and fails when they guard the net on different conditions.

The three defects below are closed. This file still records them, because a session that reads a run 29 or run 30 document needs to know what it holds.

CAUTION: question gate run 29 fails on two things, and PR-17B causes neither. Run 28 and run 29 built the identical conversation for every named case, so a session compared them line for line.

- The linter found 2 questions that say FNM with no expansion. D-374 changed the `power_sixty` row on 2026-08-31 and added that rule, both after run 28. Run 29 is the first run under them. The ask role kept the expansion 4 times and dropped it 2 times, out of 17 firings of the row.
- The dead-end check of D-357 reports 10 conversations, and 8 of them wait on the budget. Run 28 held the same conversations and the same open keys. The check did not exist then, so run 29 is its first full run.

CAUTION: the root cause of those 10 was the gate harness, not the agent. `agentsvc.Chat` closes the open questions on a turn that asks nothing new, and the gate did not. All 10 sat on the last scripted turn, and production heals every one. D-385 puts the net in the gate.

A session marked the three set conversations as probes after this run, so they now sit outside the count. They ran as gate conversations, which raised the counted set from 27 to 30. The bar holds either way, and the next run counts 27 again.

Two defects of the set filter reached the gate, and both are fixed:

- The role caps were a second cut beside the theme cut. The "other" cap of 10 dropped 40 of the 128 cards the Hobbit family offers in black-red. A free dry run found it before any paid call.
- The set row asked about Tarkir on two turns in a row, because nothing recorded the phrase it had already named. It follows the D-210 rule now.

CAUTION: the eight snapshots on the owner's disk held no set file. A session wrote one into `20260831T090157` by hand, so the tests and the gates read a real table. Every other version falls back to a derived table with no family link, and the worker fills the newest one on its next cycle.

## PR-18, and the review fixes on `nits-and-fixes` (2026-09-01)

PR-18 merged as #53. A review the same day found three defects and four gaps, and the owner approved every fix (D-398 to D-404). The fixes sit on branch `nits-and-fixes`, which starts at `main`.

- The binder filter, the sort, and the search run on the server (D-398). `GetCollectionRequest` takes a `BinderFilter` and a `BinderSort`, the page token names both, and the answer carries `matched_rows`. `collections.Filter` and `collections.Apply` hold the rule, and `binder_test.go` proves it.
- The summary lists every set and counts every card type (D-398). `CollectionSummary.sets` replaces `top_sets`, `by_type` is new, and `by_color` left (D-402). The repo reads the card index for the summary of a document stored before the field existed, through `Repo.WithIndex`.
- A derived document id belongs to its hash alone (D-399). `Repo.Put` creates the derived document, and when it exists with another hash the upload takes a fresh id. An emulator test and a service test both walk the sequence: upload X, replace with Y, upload X again.
- The upload dialog offers a choice when a collection is active (D-400): replace it after a diff, or add a new one. The name field shows for a new collection alone.
- The art loads in buckets of 100 rows (D-401). One scroll to the foot of 2,471 rows made 12 page calls and 24 art calls.
- A replacement whose file resolved no row is FailedPrecondition (D-403), and the page names no collection when an answer has no id.
- `ImportCollection` reads its file through `parseUpload`, as `DiffCollections` does. Delete invalidates the `["collection"]` prefix. The head and the binder show an error state with a retry. `statsOf` and `artIds` left the web app.
- The decline control of the budget question reads "No budget" (D-404). A declined budget stores no cap.
- The chat refuses a turn with no card index (D-405). The web app shows the failure as it shows any other.
- A decline applies its rules on both paths (D-406). The "You decide" control on the format question set no format before. No power row or commander row fired, and the build started with no format. `State.DeclineKey` holds the format default and the commander delegation, and the classifier path calls it too.

CAUTION: `make dev` runs the API binary of its start. A Go change needs a restart, or a browser read proves the old server. A session read the merged binary first, and the run reproduced the D-399 overwrite and ignored every filter. The second read built the working tree into the scratchpad and ran it on port 8091 with the environment of `scripts/dev.sh`. The script rerouted the collection and card calls to it with `page.route` and `route.fetch`. Every number below comes from that second read.

A session read the screen in a real browser with Playwright, against the real API and the 2,547-row export in `go/internal/collections/testdata`. It measured rather than looked.

| What | Measured |
|---|---|
| The import | 4,316 cards, 2,547 rows resolved, 1 unresolved, 2,471 rows in the binder |
| The set menu and the type menu | 126 sets, and 8 types with their counts, from the summary |
| Search "sol ring" | 7 of 2,471 rows, one request, the first tile Sol Ring |
| Filters chained: red, then Creature, then 4 or more | 498, then 295, then 15 of 2,471 rows |
| Sort by price | $92.74 first, and the six dearest in order |
| Scroll over 120 frames | median 16.4 ms, p95 19.8 ms, 47 frames over 16.7 ms, with the RPC reroute in the path |
| Art calls over those 120 frames | 0 |
| Scroll to the foot | 12 page calls, 24 art calls, the last virtual row 801 |
| The upload dialog with a collection active | two choices, and the name field under "Add" alone |
| The diff of the synthetic file over the export | 2,602 rows added, 2,438 removed, 21 changed |
| Replace | the request named the collection, and the list still held one, at 6,605 cards |
| The old file uploaded as new after the replace | two collections, and the replaced one kept its 6,605 cards |
| The budget question | "No budget", and "You decide" on the other two |
| Console errors, horizontal overflow | none |

CAUTION: the frame numbers come from headless Chromium against the dev server, with every RPC rerouted through Node. They say the grid is in the right range. They are not the gate line, which is the owner's own measurement on a production build.

CAUTION: `Repo.Get` returns a collection the caller owns. A page is a slice of the entry list, taken in place. A repo that shares one object across calls hands the next reader a collection the last page truncated. The test fake answers a clone.

CAUTION: the active collection is a choice of one visit, and the store keeps only the session id (D-345). A Playwright run can not seed it through localStorage. The script clicks the collection, as a reader does.

## PR-19, the chat and build experience (2026-09-02)

Branch `pr-19` holds it, from `main` at `13ca8dd`. The decisions are D-432 to D-453.

CAUTION: the branch built a start form at `/build` first (D-432, D-434). The owner read it and refused it the same day: the app is chat, and the pool picker is the one control outside it (D-436). The form, its `GetCatalog` rows, `questions/form.go`, and the six form conversations of the gate left. Do not bring a form back.

- The landing is a new chat at `/session/new`, as D-334 said. The unfinished chats sit under its message box, only when there is one, with resume, rename, and delete (D-433, D-438). `ListSessions`, `UpdateSession`, and `DeleteSession` serve it. The store's `List` inflates each session, because a session is a few kilobytes and the summary needs its first message.
- A deck tile of the library carries a delete, behind the same question as the deck screen (D-439).
- Every clickable control shows the pointer cursor, through one base rule in `index.css` (D-440). Tailwind 4 dropped the pointer from buttons.
- The wordmark's shimmer no longer runs forever, and a dialog overlay fades with no backdrop filter (D-441). The page idled at 18 ms a frame under the endless shimmer, and the delete dialogs stuttered on top of it.

CAUTION: measure the idle page before a dialog. The delete dialogs read as the fault, and the fault was a four-second animation on the wordmark that never stopped. `page.addStyleTag` with one style off at a time found it in one run.
- The gap over the message box of a new chat equals the gap under the top bar (D-442). That is 56 pixels on a desktop and 32 on a phone.
- A single card option never zooms, and a half of a commander pair zooms by two to a single card's size (D-443).
- The card art of an option picks it, as the name button does (D-444). The box over the art takes the click, and the image stays readable.
- A commander tile lifts under the pointer with the gold light of a deck tile (D-445).
- A land swap is a counted change (F-31, D-448). Session `vAvg4eteJhmuPEuJwBul` asked for better lands in place of the basics, answered "a mix", and got one Plains moved to one Island. The brief holds `swap_basics` and `land_kinds`. `generate.FitSwapBasics` fits the count to the base deck and the pool. `CheckRevision` blocks a deck that holds fewer new nonbasic lands. The generate prompt is at version 11. The revise gate answers its own questions now (`answer` in `prompts.json`) and holds a clear land-swap row, revision 9. The next run has 12 turns, not 8, and it is due before the merge.
- The land bucket of the shortlist splits, half mana and half theme (F-32, D-450). Revise gate run 5 played the land ask on the Karlov deck, and the shortlist offered 40 lands that gain life and no untapped dual. `capLands` takes the mana half by `Candidate.Fix`, capped at two, then `Pop`, and the theme half by `Themed`. Probe both decks with a throwaway `cmd` before you touch it: Karlov must show Godless Shrine and Isolated Chapel, and Éowyn the three shock lands.
- A turn stores the brief it acted on in `Turn.revision_brief` (D-449). Read it before you guess what the revise role wrote.
- The session spend holds the build (D-447). A built session showed $0.0023 for 9 calls, and the build's own calls never reached the total.
- A stored question closes with its slot, and the turn that builds a deck stales the cached session first (D-446). The docked chat of a fresh deck drew the answered commander question with its art before.
- The commander offer serves an empty theme and reads the set limit (D-437). Session `DrNPxSaisYj2QVlUkTvB` declined the theme under a Hobbit limit and got a bare pick row, then a commander the build chose. `CandidateHints.SetCodes` and `UseSets` carry the sets, and the hint no longer refuses an empty theme.
- `ChatResponse.phase` streams reading, shortlist, building, checking, repairing, and done (D-435). The generator reports its three through `generate.Request.OnPhase`. The stepper in the working row lights them.
- A failed turn offers "Try again" on a retryable failure and "Reload the session" once a session exists. The reload leaves the live panel and reads the stored session again.

A session read the branch in a real browser with Playwright, against the owner's running stack. A canned stream answered the paid Chat call, and every other RPC reached the API. It measured rather than looked.

| What | Measured |
|---|---|
| The landing of a signed-in reader | `/session/new`, and the Build entry of the top bar points there |
| A form on the page | none |
| The unfinished chats under the message box | two real rows, and a rename that stuck |
| A retryable failure | the reason with its code, "Try again", and "Reload the session" |
| Try again | a third request with the same answers as the failed one |
| The pointer cursor | pointer on a nav link, an enabled button, a select, the drop zone, and the dialog buttons, and the arrow on a disabled button |
| The delete dialog, open, before and after D-441 | 8 to 12 frames of 40 over budget, then 6, and the p95 from 27 ms to 17 |
| The idle page, before and after D-441 | 33 frames of 40 over budget, then 10 |
| The gap under the top bar and over the message box | 56 and 56 pixels on a desktop, 32 and 32 on a phone |
| Console errors, horizontal overflow | none |

CAUTION: React mounts a component twice in development, and the cleanup of the first mount aborts a send that already left. A send that starts from an effect must wait one tick and cancel on cleanup, so a double mount sends once. The form's first send hit it before the form left, and no jsdom test finds it: the test renderer mounts once.

CAUTION: a Connect stream request carries a 5-byte envelope before its JSON. A Playwright route that reads the body must skip it.

CAUTION: the owner's Firestore emulator refused every transaction on 2026-09-02, and an untouched sessions test timed out after 60 seconds. A private emulator on port 8282, through `firebase emulators:exec` with a scratch config, ran every store test green. Restart the owner's emulator before `make store-check`.

## The bracket profile, decided (2026-09-02)

The owner asked how a bracket 3 deck can play like a true 3 (D-451 to D-453). PR-14 splits. PR-14A is the bracket profile. It holds the content rules per bracket, a feature vector per built deck with a band per bracket, a goldfish simulation, and a bracket gate. It comes right after PR-19. PR-14B is the learned scorer of D-413, after Phase 3B.

Commander Spellbook's `estimate-bracket` endpoint was verified on 2026-09-02. An anonymous `POST` with a text deck list returned a bracket tag. It also returned a flag per card for Game Changer, mass land denial, and extra turn, and a flag per combo for two-card and speed. The OpenAPI schema is at `backend.commanderspellbook.com/schema/?format=json`. OQ-50 holds the terms check. OQ-51 holds the Moxfield bracket field check for PR-14B.

## The commander offer for a request with no theme (2026-09-01)

Session `t8o1nGGquK6UdTQkfY3V` asked for the best deck, Commander, bracket 5, no colors, no budget, any card. The offer was Toski, Kutzil, and Mondrak. A free test over the snapshot showed why. The theme words were "best", "you", and "can", and each unknown word becomes a text needle. The words "you" and "can" sit in the text of most commanders, so 3,029 of them scored the same. The offer was the three most popular of those.

D-411 drops a needle that more than a tenth of the cards hold, and it adds the words of a request to the stop words. The same request now reaches the unthemed pool, which ranks on EDHREC popularity.

CAUTION: the app holds no power signal for a commander. The bracket drops Game Changers under bracket 3 and nothing else. A bracket 5 request gets the most popular commanders, not the strongest, until PR-14 brings a meta source. OQ-48 records it.

CAUTION: `buf breaking` runs against `main`, and PR-18 merged before the review renamed two summary fields. D-412 reserves the merged numbers. Never rename a field of a merged PR in place, whatever the deployment state.

## PR-14 is the deck quality model now (2026-09-01)

The owner asked for a model of what makes a deck good, bad, and great, over Standard, Modern, and Commander at every quality. Five decisions settle it (D-413 to D-417), and the PR-14 entry of the roadmap holds the plan. `docs/reference/deck-quality-sources-2026-09-01.md` holds the verified facts.

- An MTGO event page embeds every list and the standings in `window.MTGO.decklists.data`. A raw fetch reads it, and the fetch tool of a session does not, because the page renders on the client.
- The owner declined Topdeck.gg, then added it the same day (D-417). cEDH standings and decklists come from its API. The app shows the credit line "Tournament data by TopDeck.gg". The key is a secret in `.env` and Secret Manager.
- The database hosts its lists on Moxfield, and a fetch of the Moxfield terms answered 403. OQ-49 waits on the owner.

## The precon exclusion, decided (2026-09-01)

The owner asked for a deck that uses no card of a precon they own, for any set. Three decisions settle the slice (D-407 to D-409), and it waits in Phase 4 as PR-24. `docs/reference/precon-data-2026-09-01.md` holds the verified facts. MTGJSON lists 3,013 deck products with a Scryfall id per card, the four Marvel Super Heroes Commander decks among them.

CAUTION: the MTGJSON deck endpoint answers 403 to a request with no user agent. A session hit it with Python's default client and passed with curl and a named agent.

## The dev stack of 2026-09-01: no card index, and why

The owner's `make dev` log wrote "snapshot version check failed: bucket doesn't exist" every 15 seconds, and session `YvyZtBJyUiEcGMNnxOwm` failed on it. The API held no card index, so the commander question offered no name, the agent said it chose one, and the build ended with "no card index is loaded".

The cause was one file. fake-gcs-server 1.56.1 keeps the metadata of an object in an extended attribute, `user.metadata`. A session wrote `sets.json.gz` into the snapshot folder by hand on 2026-08-31 (PR-17B), and that file carried no attribute. Every listing of the bucket then answered 404, and a read of one object still answered 200.

The fix was to write the same bytes through the fake GCS API, so the attribute exists. The API loaded the snapshot on its next cycle.

```
curl -X POST "http://127.0.0.1:4443/upload/storage/v1/b/mtg-local-cards/o?uploadType=media&name=scryfall%2F<version>%2Fsets.json.gz" \
  -H "Content-Type: application/gzip" --data-binary @sets.json.gz
```

CAUTION: never write a file into `.local/gcs` by hand. Upload it through the API on port 4443.

The chat ran a turn with no card index before D-405. The commander question then offered no name, and the status line said "I have no more commanders that fit this deck, so I chose one for you", which was not true. `Chat` answers Unavailable before the turn now, as `ImportCollection` does.

## Next steps, in order

1. The owner reads the PR-19 pull request, restarts `make dev`, and builds one deck from the chat to read the stepper. Then the owner merges.
2. Then PR-14A, the bracket profile (D-451 to D-453). It holds the land band of F-33. OQ-50 holds the Commander Spellbook terms check, to do first.
3. Then PR-20 to PR-23 in order, one gate each.
4. After Phase 3B: PR-15, then PR-14B (the deck quality model) and PR-24. OQ-51 holds the Moxfield bracket field check for PR-14B.

Deck gate run 11 ran on 2026-09-02 and passed 24 of 24. The read of every mana base is F-33. The nonbasic count swings from 0 to 33 on the same prompt, run to run, and no rule holds it. The fix is the land band of PR-14A, not a prompt line.

CAUTION: `make revise-gate | tee` hides the exit code. Read the verdict line of the document, never the exit code of a pipe.

CAUTION: the CI step "fake gcs tests" filters on `LiveStore`, and the only live test is `TestLiveFakeGCS`. The step matches no test and passes as a no-op. With the right name it needs a seeded snapshot bucket, which the CI server lacks. A separate change fixes it, on the owner's word.

Deck gate run 10 is done. It ran on 2026-08-31, and `CLAUDE.md` recorded it while this file still asked for it. A session that reads only the prose here spends $1.09 on a run that exists. Read `docs/reference/` before you plan a paid run.

A ruleset that requires the `verify` check on `main` is not possible. The repo is private on the free plan, and the rulesets API answers 403 (checked 2026-08-28). The owner reads the checks before a merge.

No owner row waits in `docs/owner-questions.md`, and no row waits in `docs/open-questions.md`. The owner answered every open question on 2026-09-01 (D-419 to D-431). The judge role runs on Opus 5 now (D-430), and `make llm-defaults-check` reports the change on the next run.

## The dead conversation of 2026-08-31 (D-351 to D-354)

Session `0EqqY19J6A4BxCVgsmAE` stopped for good. The reader answered "No" to two yes-or-no questions, and both keys stayed in the asked state. Readiness needs no key in that state, so no build started. The agent never repeats a question it asked, so no new question came. Every later turn ended with no question, no message, and no deck.

Four changes close the hole:

- `State.DeclineNegative` reads a bare negative and closes the key (D-352). It pairs the answer with its question by id, so it matches no text.
- `Answer.declined` and the "You decide" control on a question card say the same thing outright (D-353).
- The classifier prompt now names a negative answer as a decline (D-354). Its "none closes no key" rule reads as commander_pick alone.
- A turn that asks nothing new and is not ready closes what is out and builds (D-351). That is the net under the three above.

CAUTION: D-352 and D-354 change what the questions engine does with an answer. Read the run 28 gate and its eval against run 27 before you trust a comparison with an older run.

## The paid runs of 2026-08-31

The re-baseline of D-302 is done for the question gate and the deck gate. Total spend: $1.31.

| Run | File | Result |
|---|---|---|
| Question gate 28 | `pr7-question-gate-run28.md` | PASS. 26 of 27 catalog-only, over the bar of 25. Run 27 sat on it at 25. $0.1589. |
| Question eval 28 | `pr7-question-eval-run28.md` | 6.8 percent bad on the holdout, from 7.2. The tune split reads 2.2 percent, from 4.1. $0.0916. |
| Deck gate 9 | `pr8-deck-gate-run9.md` | PASS, and it found a defect the verdict can not see. $1.0450. |

Deck gate run 9 ran with the unbounded owned-first fill. Four decks that cost nothing to buy on run 8 cost $39.81, $80.24, $60.77, and $167.98 on run 9. D-362 bounds the fill at 150 names.

Run 10 proves the bound. Every one of those four decks costs nothing again, and the two other owned-first prompts hold their shape.

| # | Prompt | Run 8 | Run 9 | Run 10 |
|---|---|---|---|---|
| 2 | aristocrats, owned first | $0.00 | $39.81 | $0.00 |
| 5 | blink, owned first | $0.00 | $80.24 | $0.00 |
| 13 | the commander is not owned | $0.00 | $60.77 | $0.00 |
| 15 | delegated commander | $0.00 | $167.98 | $0.00 |
| 16 | a tight budget | $6.17 | $6.83 | $23.03 |
| 17 | upgrade a precon | $73.55 | $77.12 | $78.47 |

Run 10: PASS, 18 of 18, one repair turn, $1.0787.

CAUTION: read the buy cost of each owned-first prompt, not the verdict. All three bars of this gate read legality, never cost. Prompt 16 rose from $6.17 to $23.03 over two runs. Its budget is $25, so no finding fires. Read it again on the next run.

## The dead-end check of the question gate (D-357)

`cmd/questions-gate/stall.go` fails a dead-end conversation. A dead end is a last turn that sent no question and did not report ready. Such a turn also left every slot as it found it, with a question still out. The gate reports a stall that a later turn recovers, and it fails neither.

The check has 11 unit cases and one live smoke run of 8 conversations, which found no dead end. The other 96 conversations of the set are unproven. Read the first full run: a false failure means one line goes, `len(deadEnds) == 0` in the pass expression.

## Speed, measured on 2026-08-30

Every number below comes from Playwright over the built app on `vite preview`, with each RPC stubbed. `scripts` in the scratchpad hold the runs. The same measurement on the dev server gives larger numbers, because Vite serves each module on its own there.

| What | Before | After |
|---|---|---|
| Content of `/session/new`, from navigation start | 347 ms | 44 ms |
| First open of the Build menu | 323 ms | 16 ms |
| Second open of the Build menu | 8 ms | 12 ms |
| Hop to a page it already read | 1 call | 0 calls |
| First-paint chunk, gzipped | 115.93 kB | 116.13 kB |

Three changes give that:

- `src/app/deferred.tsx` replaces `React.lazy` (D-338). A deferred unit starts a download and mounts the component the moment the code is here. Nothing suspends, so React throttles nothing.
- `src/app/chunks.ts` lists every deferred chunk, and the layout warms them all in the idle time after the first paint (D-339).
- The query client keeps server state fresh for 30 seconds (D-340).

CAUTION: the first-paint bar of D-323 is 130 kB gzipped. Read the Vite build report after any change to `src/app/chunks.ts`. A chunk that moves into the entry chunk spends that budget.

## Facts that expire

- Comprehensive Rules: the current file is 2026-08-19 (D-272). The rule citations in the corpus match the 2026-08-07 text.
- Ban-list snapshot: 2026-08-24. Next announcement 2026-10-12, in `internal/cards/announcement_dates.json`. A test fails when that file holds no future date. This is by design.
- Commander brackets: the 2025-10-21 revision. Game Changers: 53 cards, list of 2026-02-09. Lutri is banned as a companion only, per the 2026-02-09 announcement (`companion_bans.json` holds the link).
- Standard: 18 sets, Wilds of Eldraine to The Hobbit. Six sets leave at the first 2027 set. Verified 2026-08-24.
- Card snapshot on disk: `.local/gcs/mtg-local-cards/scryfall/20260824T090152`. Check every card fact against it.
- LLM model ids and prices: `roles.json` and `prices.json`, verified 2026-08-24. The max output per provider in `llm/client.go`, verified 2026-08-29. The OpenAI rows are unverified by anyone but the owner.
- Run cost: the question gate costs $0.152 to $0.165 over about 20 minutes. The eval costs $0.092 to $0.104 over about 13 minutes. The deck gate cost $1.09 on run 8. The revise gate cost $0.29 on run 2. The next runs re-measure all four.
- Toolchain: Go 1.27.0, Node 22.23.2, pnpm 9.2.0, firebase-tools 14.14.0, Java 17.

## How to resume

1. Run `git pull`, then `git status`. Work on `main` or a branch the owner names. The owner commits and pushes.
2. Run `ps aux | grep autotune` before any write. The loop resets the tree when it rejects an iteration.
3. Load the skills. Load `ste-writing` before you write any `.md`. Load `design-doc-style` before you edit the roadmap. Load `mtg-corpus` before you reason about a format, a legality, or a card term.
4. Check the tree: `cd go && go build ./... && go vet ./... && go test -race ./...`, then `make lint` from the root. `make lint` runs the STE check too.
5. Do "Next steps, in order" above. Ask questions as they come up. Record each owner answer in `docs/decisions.md`, and delete the row from `docs/owner-questions.md`.
6. Before you end, update this file.

Six things a fresh session gets wrong without this file.

- Nine targets and the loop script spend money: `make questions-gate`, `make questions-eval`, `make eval-calibrate`, `make deck-gate`, `make revise-gate`, `make chat-probe`, `make generate-probe`, `make summary-judge`, `make test-smoke`, and `scripts/autotune.sh`. Ask the owner before each run. `make autotune` is free.
- A rerun writes to a new file. Every `*_OUT` variable refuses a document that holds a result (D-65).
- A gate run takes about 20 minutes and an eval about 13. A foreground command stops at 10 minutes, so run both in the background.
- `docker compose up` needs provider keys in `.env` now, and fails fast without them (D-267). `make dev` still starts with no keys, and the fake serves the health role only.
- The eval and the agent share a model. Every ratio it reports is a floor, not a measurement (D-136).
- The judge is noisy: two runs of identical code move up to nine bad questions (D-230). One run proves nothing on its own.
