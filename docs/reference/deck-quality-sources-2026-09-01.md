# Deck quality sources (2026-09-01)

This note records the facts the deck quality model of PR-14 rests on (D-413 to D-416). A session verified each one on 2026-09-01 against the live site. Re-verify before the slice starts.

## MTGO decklists, the official source (D-5)

- `https://www.mtgo.com/decklists` lists events per month for eight formats: Standard, Modern, Pioneer, Vintage, Legacy, Pauper, Limited, Duel Commander, Premodern, and Contraption. The event types are League, Challenge, and Showcase. The archive runs from 2015.
- A month page lists the event pages. On 2026-09-02 the September page held 21 events over two days, among them three Modern Challenges, one Standard Challenge, and daily Modern and Standard Leagues.
- An event page renders from an embedded object, `window.MTGO.decklists.data`. Its keys are `brackets`, `decklists`, `description`, `event_id`, `final_rank`, `format`, `inplayoffs`, `player_count`, `site_name`, `standings`, `starttime`, `type`, `url`, and `winloss`.
- A Modern Challenge 32 page of 2026-09-01 held 32 decklists. Each one carries `player`, `main_deck`, and `sideboard_deck`. A card row carries `qty` and `card_attributes` with `card_name`, `cardset`, `rarity`, `color`, and `card_type`.
- `standings` carries one row per player with `rank`, `score`, `opponentmatchwinpercentage`, `gamewinpercentage`, and `eliminated`. The placement comes from there.
- One event page is about 330 KB of HTML.

CAUTION: a page reader breaks when the markup changes. The worker stores the raw page, so a new parse needs no new fetch.

## MTGTop8, an aggregator (D-5)

- A format page shows the archetype shares over "Last 2 Weeks", "Last 5 Days", and "Last Major Events (2 Months)", the most played cards, and the events with a star rating.
- The formats hold Standard, Pioneer, Modern, Legacy, Vintage, Pauper, cEDH, Duel Commander, and more. The events run back to 2011.
- Each archetype page links decklists with a placement. No API or export exists.
- The page carries a copyright note: the card information is the property of Wizards of the Coast.

## The cEDH Decklist Database

- `https://cedh-decklist-database.com/` curates competitive Commander lists in four sections: "Competitive Decks", "Brewer's Corner", "Database", which holds viable lists and "Historic" lists, and an "OUTDATED" status.
- Every list links to Moxfield. No export exists, and no reuse terms are visible.

CAUTION: the lists live on Moxfield, and a fetch of Moxfield's terms answered 403 on 2026-09-01. The terms are unverified (OQ-49). The database gives the tier and the commander, and the cards wait on the answer.

## EDHREC (D-5)

- The commanders page ranks commanders by deck count, for example 52,420 decks for the first one on 2026-09-01. No documented API or export exists.

## Topdeck.gg, the cEDH API (D-417)

- `https://topdeck.gg/docs/tournaments-v2` documents the API. The base is `https://topdeck.gg/api`, every request carries a key in the `Authorization` header, and the key comes free from `/developers`.
- The limit is 100 requests a minute, and a request over it answers 429. The terms of the API say: "Any project using the API must include a visible credit and link back to TopDeck.gg."
- `POST /v2/tournaments` searches completed tournaments by game, format, and date range. `GET /v2/tournaments/{TID}` returns the tournament, its standings, and its rounds. `GET /v2/tournaments/{TID}/standings` returns each player's placement with `wins`, `draws`, `losses`, and `winRate`.
- A decklist appears when the tournament ended or the organizer showed the decks, as a structured `deckObj`. The formats are case-sensitive, for example `EDH`.
- The owner declined Topdeck.gg on 2026-09-01 (D-415) and added it the same day (D-417). EDHTop16 aggregates it, so the plan reads Topdeck.gg alone.

## Moxfield

- A fetch of the terms by a session answered 403 on 2026-09-01. The owner read them the same day and reported that they allow the fetch (D-419). No session has read the text, so the fact is the owner's.
