# Precon data: the MTGJSON deck files (2026-09-01)

This note records the facts the precon exclusion slice rests on (D-407 to D-409). A session verified each one on 2026-09-01 against the live MTGJSON endpoints. Re-verify before the slice starts.

## The deck list

- `https://mtgjson.com/api/v5/DeckList.json` lists every deck product. On 2026-09-01 it held 3,013 entries at version `5.3.0+20260901`.
- Each entry carries `code`, `fileName`, `name`, `releaseDate`, `source`, and `type`.
- The counts per type, on that day:

| Type | Entries |
|---|---|
| Secret Lair Drop | 723 |
| Jumpstart | 570 |
| Theme Deck | 220 |
| MTGO Redemption | 197 |
| Commander Deck | 190 |
| Intro Pack | 167 |
| Deck Builder's Toolkit | 120 |
| Arena Starter Deck | 101 |
| Bundle Land Pack | 89 |
| Box Set | 71 |
| Duel Deck | 52 |
| Welcome Deck | 50 |
| Planeswalker Deck | 41 |
| Event Deck | 26 |
| Challenger Deck | 22 |
| Starter Deck | 20 |

- The newest release date in the list was 2026-08-18. The newest Commander decks were the Marvel Super Heroes Commander decks of 2026-06-26, set code `MSC`. They are Avengers Assemble, Doom Prevails, The Fantastic Four, and Wakanda Forever. Each one has a Collector's Edition entry of its own.

## One deck file

- `https://mtgjson.com/api/v5/decks/<fileName>.json` holds one deck. The keys are `code`, `commander`, `displayCommander`, `mainBoard`, `name`, `planes`, `releaseDate`, `schemes`, `sealedProductUuids`, `sideBoard`, `source`, `sourceSetCodes`, `tokens`, and `type`.
- Each card carries `count`, `isFoil`, `finishes`, `setCode`, `number`, `uuid`, and `identifiers`. The identifiers hold `scryfallId` and `scryfallOracleId`.
- `AvengersAssemble_MSC.json` holds 89 main-board rows, 99 copies, one commander (Captain America, Team Leader), and no card without a Scryfall id.

CAUTION: the deck endpoint answers 403 to a request with no user agent. The worker must send one.

CAUTION: MTGJSON also offers every deck file in one archive. A session did not measure its size, so the fetch plan is unverified.

## What the slice adds

- A `precons.jsonl.gz` file beside `sets.json.gz` in the snapshot (D-377 shape). A row holds the product name, the set code, the type, the release date, and the cards as printing ids with counts. About two megabytes gzipped for every product.
- `candidates.Request.ExcludeOracleIDs`, read by the commander pool and the 99, and a block in the rules check.
- An ownership check over the collection, on the printings and their counts (D-408).
- A classify fact and a catalog row for "not from my precons" and "not from precon X", and the chat names what it excluded (D-390).
- The nine embedded lists of D-247 become a cross-check test, then leave.
