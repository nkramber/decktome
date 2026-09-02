# The bracket profile: sources and facts (2026-09-02)

This note holds every external fact PR-14A rests on, with its source and the date it was read. `go/internal/profile/bands.json`, `go/internal/rules/brackets.json`, and `go/internal/profile/karsten.go` carry the numbers. This note says where they came from.

## 1. The bracket rules

Source: the Wizards of the Coast article "Introducing Commander Brackets (Beta)" of 2025-02-11, and its update of 2025-10-21. Both were read on 2026-09-02 through a fetch summary, not the raw page.

| Bracket | Name | Turns | Game Changers | Mass land denial | Extra turns | Two-card infinite combos |
|---|---|---|---|---|---|---|
| 1 | Exhibition | 9 or more | none | none | none | none |
| 2 | Core | 8 or more | none | none | low quantities, never chained | none |
| 3 | Upgraded | 6 or more | up to three | none | low quantities, never chained | none that is cheap and in about the first six turns |
| 4 | Optimized | 4 or more | no limit | no limit | no limit | no limit |
| 5 | cEDH | any | no limit | no limit | no limit | no limit |

The 2025-10-21 update removed every tutor limit. The article defines mass land denial by its effect. Such a card destroys, exiles, bounces, keeps tapped, or changes the mana of four or more lands per player, with no replacement.

## 2. The Commander Spellbook endpoint

Source: `https://backend.commanderspellbook.com/schema/?format=json` and the backend source `backend/spellbook/models/variant.py` on GitHub (SpaceCowMedia/commander-spellbook-backend, MIT license, last commit to that file 2026-09-01). Both were read on 2026-09-02.

`POST /estimate-bracket` takes a JSON body with `commanders` and `main`, each a list of `{card, quantity}`, up to 600 main cards. It also takes a plain-text deck list. It answers `bracketTag`, a `cards` list with `gameChanger`, `massLandDenial`, `extraTurn`, and `banned` per card, and a `combos` list. A combo carries `relevant`, `definitelyTwoCard`, `arguablyTwoCard`, `speed`, `extraTurn`, `massLandDenial`, `lock`, `skipTurns`, and `controlAllOpponents`.

The speed scale: 5 needs no mana, 4 needs four or less, 3 six or less, 2 eight or less, and 1 more. An uncertain minimum adds one.

The endpoint's own bracket rule, from `estimate_bracket`:

- Banned: any banned card.
- Ruthless (4): more than three Game Changers, or two or more extra-turn cards, or an extra-turn combo. Also any mass land denial, a combo that controls every opponent, or a relevant two-card combo at speed 4 or more.
- Spicy (3): a near two-card combo at speed 4 or more, a lock, a skip-turns combo, or a combo that controls some opponents.
- Powerful (3): a relevant two-card combo at speed 3 or more, or any Game Changer.
- Oddball (2): a borderline two-card combo at speed 3 or more.
- Core (2): one extra-turn card, or a relevant two-card combo at speed 2 or more.
- Exhibition (1): the rest.

A live call on 2026-09-02 with an anonymous `POST` answered in under a second with no rate-limit header. The owner read the terms and allowed the call at 90 requests a minute on the app's side (D-459).

The app's content rules in `brackets.json` follow the endpoint's thresholds:

- No mass land denial through bracket 3.
- No extra-turn card at bracket 1, and one at brackets 2 and 3.
- No extra-turn combo through bracket 3.
- A two-card combo cap of speed 0 at bracket 1, speed 2 at bracket 2, and speed 3 at bracket 3.

## 3. The Karsten source tables

Source: Frank Karsten, "How Many Sources Do You Need to Consistently Cast Your Spells? A 2022 Update", TCGplayer, 2022-08-02, read 2026-09-02 through the TCGplayer content API. The consistency target is 89 percent plus the mana value. The 60-card column assumes 25 lands, and the 99-card column 41 lands with the free mulligan and the draw on turn one.

| Cost | 60 cards | 99 cards |
|---|---|---|
| C | 14 | 19 |
| 1C | 13 | 19 |
| 2C | 12 | 18 |
| 3C | 10 | 16 |
| 4C | 9 | 15 |
| 5C | 9 | 14 |
| CC | 21 | 30 |
| 1CC | 18 | 28 |
| 2CC | 16 | 26 |
| 3CC | 15 | 23 |
| 4CC | 13 | 22 |
| 5CC | 12 | 20 |
| CCC | 23 | 36 |
| 1CCC | 21 | 33 |
| 2CCC | 19 | 30 |
| 3CCC | 17 | 28 |
| 4CCC | 16 | 26 |
| CCCC | 24 | 39 |
| 1CCCC | 22 | 36 |

The article counts a fetch land as a full source of every color it can find. It counts a mana rock as three-fourths of a source for spells of mana value three or more. The app counts a rock as three-fourths for every spell, which is one step simpler.

## 4. The Karsten land formulas

Source: Frank Karsten, "How Many Lands Do You Need in Your Deck? An Updated Analysis", TCGplayer, 2022-07-29, read 2026-09-02 through the content API.

- 60 cards: lands = 19.59 + 1.90 × the average mana value − 0.28 × the cheap draw or ramp spells. Add 0.27 with a companion.
- 99 cards: lands = 31.42 + 3.13 × the average mana value − 0.28 × the cheap draw or ramp spells. The 1.35 cut accounts for the free mulligan and the draw on turn one.

Source: Frank Karsten, "What's an Optimal Mana Curve and Land/Ramp Count for Commander?", TCGplayer, 2022-07-15, read 2026-09-02. The optimal lists by commander cost:

| Commander cost | Mana rocks | Lands |
|---|---|---|
| 2 | Sol Ring | 42 |
| 3 | Sol Ring | 42 |
| 4 | Sol Ring and 7 Signets | 39 |
| 5 | Sol Ring and 8 Signets | 39 |
| 6 | Sol Ring and 9 Signets | 38 |

The mulligan rule of that article has three steps. Keep a first seven with three to five lands. Keep a second seven with two to five. Keep a six or a five with two to four lands. The goldfish simulation follows it.

## 5. The tag slugs

Source: the `oracle_tags` file of the snapshot of 2026-09-02, counted by direct taggings. `mass-land-denial` holds 114 cards, `extra-turn` 64, `mana-rock` 109, `mana-dork` 456, and `land-ramp` 494. The tutor set is `tutor` less `tutor-land`, so a basic land search is ramp and not a tutor.

## 6. What no source states

CAUTION: five band groups of `bands.json` are this session's first values. They are the average mana value, the tapped and colorless land caps, the tutor and fast mana caps, and the goldfish floors. No published table states them.
