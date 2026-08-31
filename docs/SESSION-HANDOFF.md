# Session hand-off

`CLAUDE.md` is the entry point. It sends you here. Then read `docs/decisions.md` from D-297, `docs/owner-questions.md`, and `docs/open-questions.md`. The session narratives of 2026-08-23 to 2026-08-28 sit in `docs/reference/session-log-2026-08-23-to-26.md` and `docs/reference/session-log-2026-08-27-to-28.md`. The audit of 2026-08-29 is `docs/audit-2026-08-29.md`.

CAUTION: the web tests need Node 22.23.2 (`.nvmrc`). Under Node 20 every test file fails at start with `ERR_REQUIRE_ESM` from jsdom 30. On the owner's machine `~/.nvm/versions/node/v22.23.2/bin` on the PATH fixes it.

## Where things stand (2026-08-30)

- `main` is at `6f871c3`. Merged: PR-0a to PR-8, PR-7B, PR-10 to PR-13, the audits, the Phase 3B roadmap (#46), PR-16 (#47), and PR-16B (#48).
- No pull request is open. The Dependabot bump of the go group merged as #44.
- Branch `pr-17` holds the deck library and the look of the reference design, in nine commits over `main`. The version history and the compare remain.
- The tree is green on `pr-17`: Go build, vet, `-race` tests, golangci-lint, `buf breaking`, `proto-check`, web lint, typecheck, 165 web tests, and the web build. `make lint` reports zero findings.
- The first paint is 361.69 kB raw and 115.72 kB gzipped. The bar of D-323 is 130 kB.
- The revise gate held on run 2 (D-296). The paid gates did not run since 2026-08-28.
- PR-9 is out of the MVP (D-256). Phase 3B comes before Phase 4 (D-316).

CAUTION: branch `pr-17` carries six concerns. They are the contract, the Go side, the reference design, the layout of D-331, the Build menu, and the one deck screen. Guardrail 10 asks for one per pull request. A split before the merge needs the owner's word.

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

CAUTION: no chunk of this app loads behind a Suspense boundary (D-338). React holds a committed fallback for 300 ms, and it holds every later reveal with it. `src/app/deferred.tsx` replaces `React.lazy` everywhere. Do not put `React.lazy` back.


The Go side, in one list:

- `DeckService.UpdateDeck` writes the name and the favorite mark, and `DeleteDeck` removes a deck for good. Both are additive, and `buf breaking` passes.
- `Deck.favorite` and `Deck.card_count` are new. The list view carries no cards, so it sets `card_count`. The deck list on screen reads `cards.length` today and always shows zero.
- `ListDecks` takes `page_size`, `page_token`, `format`, `favorite`, `query`, `power_bracket`, and `power_sixty_step` (D-324).
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

CAUTION: the binder head calls `GetCollection`, and the answer carries every entry. The owner's export holds 4,952 rows, so one page load moves about one megabyte. PR-18 adds paging, and the head reads a page then.

## Reading the app without the owner (2026-08-30)

`@playwright/test` is a dev dependency of `web/apps/web` now, and Chromium sits in the local cache. A session can read its own work.

- Start the Auth emulator and the Vite dev server. Then drive Chromium with a script that answers every `/mtg.v1.*` call from canned JSON, so no backend has to run.
- A Connect Timestamp is an RFC 3339 string in JSON, not a `{seconds}` object. A `{seconds}` stub fails with "cannot decode message google.protobuf.Timestamp".
- The script lives outside the repo, in the session scratchpad. It must run from `web/apps/web`, or the bare import of `@playwright/test` does not resolve.
- PR-23 holds the real smoke flow (D-313). This is a reading tool, not that.

## Next steps, in order

1. The owner reads `pr-17` in the browser and says whether the look and the one deck screen are right.
2. The owner says whether to split `pr-17` before the merge. Six concerns sit on it.
3. Finish PR-17: the version history from the `revised_from_deck_id` chain, and the compare of two decks.
4. The owner runs the question gate and the deck gate to re-baseline (D-302), in parallel. Ask before each run. Write each to a new `GATE_OUT` file (D-65). Record the numbers here and in the roadmap.
5. Then PR-18 to PR-23 in order, one gate each. Before PR-22, ask OQ-45 and OQ-46.
6. After Phase 3B: PR-15, then PR-14.

A ruleset that requires the `verify` check on `main` is not possible. The repo is private on the free plan, and the rulesets API answers 403 (checked 2026-08-28). The owner reads the checks before a merge.

Seven owner rows wait in `docs/owner-questions.md`: OQ-23, OQ-28 to OQ-31, OQ-37, and OQ-39. OQ-20, OQ-44, OQ-45, and OQ-46 wait in `docs/open-questions.md`.

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
