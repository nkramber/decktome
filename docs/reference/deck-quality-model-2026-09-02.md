# The deck quality model (2026-09-02)

This note records every fact and every call the deck quality model of PR-14B rests on (D-413 to D-417, D-451, D-460). A session verified each source fact on 2026-09-02 against the live site. `docs/reference/deck-quality-sources-2026-09-01.md` holds the first read of the sources. Re-verify before a paid run.

## The sources, re-verified

| Source | What it gives | Verified 2026-09-02 |
|---|---|---|
| MTGO decklists | Challenge lists with standings, league lists with a record | A month page answers 200, and an event page embeds `window.MTGO.decklists.data`. |
| MTGJSON deck products | Every deck product with a Scryfall id per card | `DeckList.json` held 3,029 products at version `5.3.0+20260902`. |
| EDHREC | Deck count, rank, bracket counts, and the average deck per commander | `json.edhrec.com/pages/commanders/<slug>.json` and `average-decks/<slug>.json` answer 200. The top pages `commanders/year.json`, `month.json`, and `week.json` hold 100 commanders each. `commanders.json` answers 403. |
| cEDH Decklist Database | The competitive tier and the commanders of each entry | One page, 1.1 MB, 137 entries: 56 competitive, 9 brews, 72 outdated. |
| Topdeck.gg | cEDH tournaments with standings and decklists | The docs at `topdeck.gg/docs/tournaments-v2` answer 200. No key was in `.env`, so no session has read a live answer. |
| Moxfield | The lists the database links | A plain client with a named agent gets 403 from Cloudflare on the deck page, the v2 API, and the v3 API. The `moxfield-api` library, 2.1.0 on Node 22, got the same 403 page on 2026-09-03 (D-493). |
| MTGTop8 | Events with placements | The format page answers 200 and holds no event link a plain reader can find. |

CAUTION: the Moxfield lists are out of reach for a page reader. The database gives the tier and the commander alone, as the first read foresaw (D-419, D-470). No session can check OQ-51 on the live endpoint.

CAUTION: the MTGO event object of a challenge and of a league differ. A challenge carries `standings`, `final_rank`, and `winloss`, and a league carries `decklists` alone, each with a `wins` block. The `format` key reads `CMODERN` or `CSTANDARD` on a challenge and is absent on a league, so the reader takes the format from the slug.

## The live read of 2026-09-02

`make meta-refresh` ran with `-meta-months 1 -meta-pages 5` on the local stack. MTGO read 5 event pages into 133 lists with no parse failure. EDHREC read 13 pages into 5 lists, and the database read its 137 entries. MTGJSON read 557 deck files, and then one connection reset ended the source. The fetcher retries a transport error or a 5xx once since, after two seconds. A second run read the last 146 files and stored the table, 701 products.

One MTGO page timed out on both tries, and the job counts that as a fetch error and reads on since. The fit then ran over 194 Commander, 468 Standard, and 539 Modern lists.

## The full read of 2026-09-02

`make meta-refresh` ran whole after the owner set the windows (D-479, D-480). MTGO read 200 pages into 5,288 lists, and 29 pages failed. The site answered a 302 to the month listing for an event page it no longer serves. The client followed it, and the reader stored the listing as the event. The client follows no redirect since. A page that does not parse goes under `raw/mtgo-failed/`, and the event fetches again.

Topdeck.gg answered 21,146 lists over 90 days in 13 weekly calls, and the bulk endpoint answered 429 twice with waits of 14 and 35 seconds. A 429 retries up to four times on the header since, and the bulk calls sit ten seconds apart. EDHREC read 403 pages into 195 lists before the page cap held, and the cap no longer applies to that read. The fit then failed on a month of Commander lists past the 16 MB inflate limit of the shared gzip helper. The store inflates its own objects under 1 GB since, and a fit keeps at most 4,000 lists per tier (D-483).

## The MTGO read

The slug carries the format, the event kind, the date, and the id: `modern-challenge-32-2026-09-0212853228`. The reader keeps the Modern and the Standard events and skips the rest. Duel Commander is not Commander (D-112).

A top-8 finish in a challenge is great, and the rest of the challenge is good. A league page shows the 5-0 lists alone, so every list there is a finish, and good (D-414).

The job reads the month pages back 12 months, this month included (D-471). Then it reads every event page of a covered format the store lacks, newest first, up to 200 pages a run. The September 2026 month page held 21 events over two days. So a year holds about 3,500 events, and about a third of them are Modern or Standard. The first full read takes about six runs.

The raw page stays in the store under `meta/raw/mtgo/<slug>.gz`. A parser fix runs `worker -meta -meta-reparse`, which re-reads every stored page and fetches nothing (M-6).

## The precon table

The table holds the Commander decks and the 60-card constructed decks (D-407). The types are these:

- Commander Deck and MTGO Commander Deck.
- Brawl Deck.
- Starter Deck, Intro Pack, Planeswalker Deck, and Theme Deck.
- Challenger Deck, Pioneer Challenger Deck, Event Deck, and Modern Event Deck.

The counts per type on 2026-09-02 give about 720 products.

A Commander product is a baseline list of the Commander format. A 60-card product carries the format word `sixty`, and the Standard fit and the Modern fit both read it as a baseline (D-472). A product older than the format's pool serves no baseline: Modern reads from Eighth Edition on, and Standard from the last three years (D-478). A Brawl deck sits in the table for the ownership check of PR-24 and serves no fit.

The table lives at `meta/precons/<MTGJSON version>/precons.jsonl.gz`. A row holds the product name, the set code, the type, the release date, and the cards as printings with counts and both Scryfall ids. The job reads the deck list on every run and the deck files once per version.

CAUTION: the table follows the MTGJSON version and not the Scryfall one, so it lives under `meta/` and not beside `sets.json.gz`. PR-24 reads it there (amends the shape of D-407).

## The commander reads

One file per day, `meta/commanders/<YYYY-MM-DD>.jsonl.gz`, holds one row per commander slug. Three sources fold into a row:

- EDHREC gives the deck count, the rank, and the bracket counts. The read runs every seven days over the three top pages and the 1,000 most built legends of the card index, about 1,100 commanders (D-489).
- The database marks the competitive commanders. A pair marks the pair slug and each partner.
- Topdeck.gg counts the entries and the top-cut finishes over the last 90 days.

The cEDH signal of a commander is the top-cut share of its entries. A competitive commander with no entry reads 0.5, and one with entries reads at least 0.5. A commander with neither reads 0 (D-476).

The EDHREC slug rule: lower case, letters and digits kept, the rest cut, words joined by hyphens. "Y'shtola, Night's Blessed" is `yshtola-nights-blessed`, checked against the top page. A pair joins the two slugs with a hyphen. No session checked that rule against the site.

## The features

The scorer reads a deck through the profile of PR-14A and the corpus of its format. `go/internal/quality/model.go` names the keys.

| Group | Features |
|---|---|
| The cards | `card_rate`, the mean inclusion rate of the nonland cards in the great and good lists. `unseen_share`, the share no top list held. |
| The pairs | `synergy`, the mean log lift of the nonland pairs over chance. |
| The mana base | `land`, `color_sources`, `tapped_share`, `mana_turn_four`, `hands_two_to_four_lands`. |
| The curve | `avg_mana_value`, `curve_low` (share at mana value 2 or less), `curve_high` (share at 5 or more). |
| The jobs | `ramp`, `draw`, `removal`, `wipe`, `interaction`, `empty_roles`. |
| The power signals | `fast_mana`, `game_changer`, `playset_share` and `singleton_share` (60-card only), `source_spread`. |
| The commander | `cedh_signal`, `high_bracket_share`, `commander_decks` (Commander only). |

The inclusion rate of a card has a prior. It is the placement-weighted count of the lists that hold the card, plus five times the uniform prior, over the total weight plus five. A great list weighs two and a good list one. The lift of a pair is the pair's weight times the total over the two card weights. A pair stays when three lists hold it, the lift is 1.5 or more, and each card sits in five lists. The store keeps at most 200,000 pairs.

## The fit

One list in five goes to the holdout, by the hash of its key. The corpus rates and pairs come from the training split alone, so the holdout measures lists the scorer never saw (D-473).

Each real list of a format makes one synthetic bad list. The axis cycles over lands, curve, colors, copies, and synergy (D-414, D-473):

- lands: a third of the lands become random legal spells.
- curve: the fifteen cheapest spells become random six-drops and up.
- colors: every nonbasic land becomes a basic of one color.
- copies: every playset becomes one copy, and random singletons fill in. A singleton deck breaks on synergy instead.
- synergy: half the spells become random legal cards in the identity.

The scorer is a proportional odds model over the standardized features. Gradient descent fits it: 3,000 iterations at a rate of 0.05 with an L2 weight of 0.01. The cut points stay ordered through a first cut and positive gaps. A feature with no spread in a format leaves that format's model.

The grade is the most probable rung, and the score is the expected rung over the ladder. The reasons are the three largest products of weight and standardized value, in words.

## The defect detector

One linear ladder can not rank a top list over a precon and a precon over its broken copy at once. Gate run 1 read the second in 0.16 to 0.46 of the pairs. A logistic detector over the same features reads the second question (D-485). Every real list is a negative, in three groups of equal weight: the broken copies, the precons with the average decks, and the rest. The detector picks its cut on the training rows.

A flagged deck scores under 0.15 of the range, and every other deck scores above it. So the order between a precon and its copy follows the detector alone.

The broken copies come from the precons and the average decks alone, on every axis (D-484). A break draws its replacement cards from the ones the real lists of the format play (D-488). So it breaks the pairing or the shape, and never the card pool. The ordinal fit weighs every rung the same whatever its count. A copies break needs 12 copies in playsets, and a colors break four lands that fix two colors or more. A list short of that breaks on synergy instead.

## Where the grade lands

- `Deck.quality` (field 24) carries the tier, the score, the reasons, the model version, and a probability per rung.
- `Card.quality` (field 34) carries the inclusion rate per format and the cEDH signal, on `GetCards` alone.
- The deck summary ends with one sentence the code writes: the tier and the reasons (D-474).
- The revision note says so when a change lowers the tier (D-474).
- The shortlist takes the card rate as `MetaBoost`, a tenth of the score (PR-6 hook).
- A bracket 4 or 5 commander offer ranks on the cEDH signal first (D-476, OQ-48).
- The generate prompt carries a "Format shape" block (D-475). It names the land mean, the mana value mean, and the twelve most held cards of the great lists.

## The judge lane and the explain mode

`make quality-judge` reads a deck gate document, grades every deck with the stored model, asks the judge role for its tier, and writes the agreement. `go run ./cmd/quality-gate -explain <document>` prints, for free, how the model grades each deck: the detector probability against its cut, the ladder, and the six largest contributions. Runs 1 and 2 of the judge lane read 5 and 4 of 24 (D-488).

## The gate

`make quality-gate` is free. It fits the model over the local meta store and writes the document. The document holds the pair bars per format, the bracket 5 offer, and the model tables. `-write` stores the fitted model. `make meta-refresh` fills the store, over the network, with no model call.

The judge bar over the golden decks reads the next deck gate run: every summary carries the tier, and the summary judge reads it. No run happened on 2026-09-02.
