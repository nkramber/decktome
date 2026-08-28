---
name: mtg-corpus
description: Magic: The Gathering terminology, formats, deck-construction rules, ban-list snapshot, archetypes, and the clarifying-question catalog for the deck-builder agent. Load before you reason about any MtG request.
---

# MtG corpus skill

This skill is the knowledge base for the deck-builder agent. It is also the seed of the app's own corpus. Facts carry a date. Rules and ban lists change. Check the date before you trust a fact.

Snapshot date: **2026-08-23**. Section 11 revised 2026-08-28: five rows retired (A-6 of the 2026-08-28 audit), after the PR-7 dogfood runs of 2026-08-24 (D-66, D-67, D-68). Sources: Scryfall API and bulk data, Wizards of the Coast announcements, mtgcommander.net, and the Commander Format Panel.

Reference files in `references/`:
- `scryfall-catalogs.md` - all keyword abilities, keyword actions, ability words, and type lists.
- `scryfall-oracle-tags.md` - 765 community theme tags with 40 or more cards (of 4,522 total).

## 1. Core vocabulary

| Term | Meaning |
|---|---|
| Card | One Magic card. Identified by its English name. A card can have many printings. |
| Oracle text | The current official rules text of a card. Scryfall gives it per `oracle_id`. |
| Printing | One version of a card in one set. Scryfall gives it per `id` (Scryfall ID). |
| Mana | The resource that pays for spells. Five colors: White (W), Blue (U), Black (B), Red (R), Green (G). Colorless (C) is not a color. |
| Mana cost | The symbols in the top corner. The mana value (MV, old name CMC) is the total. |
| Color identity | The card's colors plus every mana symbol in its rules text. Commander uses this. |
| Land | A card that makes mana. Basic lands: Plains, Island, Swamp, Mountain, Forest (and Wastes). |
| Spell | Any nonland card when cast. Types: creature, instant, sorcery, artifact, enchantment, planeswalker, battle, kindred. |
| Permanent | A card that stays on the battlefield: land, creature, artifact, enchantment, planeswalker, battle. |
| Instant speed | Can be cast at any time you have priority. Instants, and cards with flash. |
| Sorcery speed | Only on your main phase with an empty stack. |
| Library | Your deck during the game. Graveyard: discard pile. Exile: removed from the game. |
| Life total | 20 in most formats. 40 in Commander. 30 in paper Brawl multiplayer, 25 in 1v1 and on Arena. |
| Mulligan | Redraw your opening hand. The London mulligan: draw seven, then put N cards on the bottom. |
| Turn structure | Untap, upkeep, draw, main 1, combat, main 2, end step. |
| Stack | Spells and abilities resolve last-in first-out. |
| Priority | The right to act. Passes between players. |
| Legendary | Only one permanent with that name per controller. |
| Token | A permanent that is not a card. Made by an effect. |
| Counter | A marker on a permanent or player (+1/+1, loyalty, poison). |
| Trigger | An ability that starts when an event happens ("When", "Whenever", "At"). |
| ETB / LTB | Enters the battlefield / leaves the battlefield. |
| Dies | Goes from the battlefield to the graveyard. |

## 2. Formats

The app builds three formats: Commander, Standard, and Modern (D-155). It declines every other format by name and asks which of the three to build instead. It names no substitute, because no measurement supports one.

The app must know the format before it builds. A "60-card" request has many possible formats. Ask.

### 2.1 Constructed formats the app builds (60-card style)

| Format | Deck size | Copies | Card pool (2026-08-24) | Sideboard |
|---|---|---|---|---|
| Standard | min 60 | max 4 per name (basic lands unlimited) | Sets from Wilds of Eldraine (WOE, 2023-09-08) to The Hobbit (HOB, 2026-08-14). 18 sets. The Big Score (BIG) is legal with Outlaws of Thunder Junction and is not counted as a separate set. No rotation in 2026. Next rotation: first set of 2027. Six sets leave then: WOE, LCI, MKM, OTJ, BLB, DSK. Verified 2026-08-24 on the Scryfall sets API: 19 paper core or expansion sets have a full Standard-legal card list, and BIG is one of them. | 15 |
| Modern | min 60 | max 4 | Eighth Edition (2003-07) and Mirrodin forward. Modern Horizons sets included. | 15 |

Scryfall `legalities` keys: `standard, future, historic, timeless, gladiator, pioneer, modern, legacy, pauper, vintage, penny, commander, oathbreaker, standardbrawl, brawl, competitivebrawl, alchemy, paupercommander, duel, oldschool, premodern, predh, tlr`. Values: `legal, not_legal, banned, restricted`.

Legal Oracle-card counts on 2026-08-23, for the three formats the app builds: commander 31,830 · modern 22,450 · standard 4,887.

### 2.2 Commander (EDH)

- Exactly 100 cards, the commander included.
- Singleton: no two cards with the same English name, except basic lands and cards that say otherwise (for example Relentless Rats).
- The commander is a legendary creature card, a legendary Vehicle card, or a legendary Spacecraft card with a power box (CR 903.3, 2026-08-07). A card that says it can be your commander also qualifies. Only the front face of a double-faced card counts.
- Read the card as it is in the command zone, not as it is on the battlefield. A card that is a creature card outside the battlefield can be your commander, even when the printed type line names no creature. Example: Grist, the Hunger Tide reads `Legendary Planeswalker - Grist`, and Scryfall's ruling of 2021-06-18 says it can be your commander. A type-line test alone gives the wrong answer for this class of card.
- Two commanders: Partner, Partner with, Partner—[text] (Friends forever, Father & son, Survivors, Character select: equal text only), Choose a Background plus a Background, and Doctor's companion with a Time Lord Doctor. A Background alone can not be a commander.
- Every card must fit the commander's color identity.
- 40 life. Four players is the normal table. 21 combat damage from one commander kills a player.
- The commander goes to the command zone when it would leave. It costs 2 more each time you cast it from there (the "commander tax").
- Rules and ban list: the Commander Format Panel (from 2024-09, owned by Wizards of the Coast with a community panel). Site: mtgcommander.net.
- Commander is a social format. Power level matters more than in any other format. Always ask about the bracket.

### 2.3 Commander brackets (official, beta from 2025-02-11, revised 2025-10-21, list update 2026-02-09)

Source: the Wizards article of 2025-10-21 and its infographic. The 2025-10-21 revision removed every tutor limit and the tie between Bracket 2 and precons.

| Bracket | Name | Expected game | Game Changers | Other limits |
|---|---|---|---|---|
| 1 | Exhibition | 9+ turns | none (thematic exceptions) | No mass land denial. No extra turns. No two-card infinite combos. |
| 2 | Core | 8+ turns | none | No mass land denial. No chained extra turns. No two-card infinite combos. |
| 3 | Upgraded | 6+ turns | max 3 | No mass land denial. No chained or looped extra turns, and few of them. No two-card combos in the first six or so turns. |
| 4 | Optimized | 4+ turns | unlimited | Only the ban list applies. |
| 5 | cEDH | any | unlimited | Only the ban list applies. Plays the best strategy, not a theme. |

The commander counts toward the Game Changer limit. The Wizards bracket article of 2025-02-11 says a Game Changer commander counts as one of the three at Bracket 3. Such a commander can not play in Brackets 1 and 2. The engine counts the command zone (verified 2026-08-26, source in section 13).

Game Changers: 53 cards on 2026-08-24 (Scryfall `is:gamechanger`). The list changed on 2025-04-22, 2025-10-21, and 2026-02-09. Scryfall flags them with `game_changer: true`. Examples: Rhystic Study, Cyclonic Rift, Smothering Tithe, Thassa's Oracle, Demonic Tutor, Vampiric Tutor, Ancient Tomb, The One Ring, Gaea's Cradle, Force of Will. Note: Mana Crypt is banned, so it is not a Game Changer.

### 2.4 Formats the app does not build

The app declines each of these by name (D-112, D-155). Seven of them have no substitute, and six have a nearest format (see the last paragraph). In every case it asks which of Commander, Standard, or Modern to build instead.

- **Pioneer** (Return to Ravnica forward), **Legacy** (all sets, ban list), **Vintage** (all sets, ban list plus a restricted list), **Pauper** (commons only). Removed 2026-08-26.
- **Historic**, **Timeless**, **Alchemy**: MTG Arena only, with digital-only cards.
- **Brawl** (Arena, 100-card singleton with a commander), **Standard Brawl** (60-card), **Oathbreaker** (planeswalker commander, 60 cards), **Pauper Commander** (an uncommon creature, Vehicle, or Spacecraft as commander, 99 commons, 30 life), **Duel Commander** (1v1, 20 life), **Canadian Highlander**.

Brawl, Standard Brawl, Oathbreaker, Duel Commander, and Canadian Highlander name Commander as the nearest format (D-192), and Alchemy names Standard. Those five are singleton formats of the same shape, and Alchemy is Standard with the Arena-only rebalanced cards.

### 2.5 "Anything goes"

Not a format the app offers, and not a format the user can select (D-155). The owner's own meaning: any card, no ban list. Other users mean "Modern but with proxies", or "kitchen table". **Always ask what the user means.** Then record the answer as the session's house rules, which ride on top of Commander, Standard, or Modern.

A user who names no format at all and asks for no ban list gets `FORMAT_ID_HOUSE`. It builds 60 cards with 4 copies and a 15-card sideboard until the user says otherwise.

## 3. Ban list snapshot (2026-08-24)

Source of truth at run time: Scryfall `legalities` (updated within a day of each announcement). Never store a ban list in a prompt. Look it up. The facts below are for orientation only.

- 2026-08-10 announcement: Standard banned Badgermole Cub, Stormchaser's Talent, Gran-Gran. Legacy banned The Fantasticar. Vintage restricted The Fantasticar. Next announcement: **2026-10-12**.
- 2026-06-29: Legacy banned Candelabra of Tawnos. Pauper banned Seeker of Skybreak. Brawl banned Force of Will, Subtlety, Wash Away, Ugin's Labyrinth, Time Warp, Temporal Manipulation.
- 2026-05-18: Pioneer banned Cori-Steel Cutter. Modern banned Phlage, Titan of Fire's Fury and Lotus Field, and unbanned Violent Outburst and Umezawa's Jitte. Legacy banned Undercity Informer. Pauper unbanned Bonder's Ornament. Alchemy banned Sewer-veillance Cam.
- 2026-03-23: Historic banned Food Chain.
- Commander 2026-02-09: Biorhythm unbanned (added to Game Changers). Lutri, the Spellchaser unbanned but "banned as a companion" (a new category). Still banned: Sundering Titan, Iona, Griselbrand, Mana Crypt, Jeweled Lotus, Dockside Extortionist, Nadu, and 35 more named cards (42 named in total) plus three categories (ante, Conspiracy, offensive cards).
- Announcement cadence: Wizards posts on a published schedule, about every 6 to 8 weeks. Commander updates come 1 to 3 times per year.

## 4. Color names

| Colors | Name | Colors | Name |
|---|---|---|---|
| WU | Azorius | WUB | Esper |
| UB | Dimir | UBR | Grixis |
| BR | Rakdos | BRG | Jund |
| RG | Gruul | RGW | Naya |
| GW | Selesnya | GWU | Bant |
| WB | Orzhov | WBG | Abzan |
| UR | Izzet | URW | Jeskai |
| BG | Golgari | BGU | Sultai |
| RW | Boros | RWB | Mardu |
| GU | Simic | GUR | Temur |
| Four colors | named by the missing color: "non-white" is also "Glint-Eye" or "Chaos" in casual use. Five colors: "five-color" or "WUBRG". | | |

Mono-colored: mono-white, mono-blue, and so on. Colorless: no colored mana symbols at all.

## 5. Archetypes and strategies

| Archetype | What it does | Typical signs |
|---|---|---|
| Aggro | Wins fast with cheap creatures and burn. | Low curve (MV 1-3), 18-22 lands in 60. |
| Midrange | Efficient threats and removal. Wins the long game against aggro, pressures control. | Curve 2-5, 24 lands. |
| Control | Answers everything, then wins with few threats. | Counterspells, board wipes, card draw, 25-27 lands. |
| Combo | Assembles specific cards for a win, often in one turn. | Tutors, card selection, protection. |
| Tempo | Cheap threats plus cheap interaction to stay ahead on time. | Flash creatures, bounce, counters. |
| Ramp | Extra mana early, big spells fast. | Mana dorks, mana rocks, land search. |
| Stax | Slows every player with taxing or lock pieces. | Bracket 4+ only in Commander. |
| Prison | Locks the opponent out of the game. | Bracket 4+. |
| Reanimator | Puts big creatures into the graveyard, then returns them cheaply. | Discard outlets, reanimation spells. |
| Storm | Casts many spells in one turn, then a payoff. | Rituals, cantrips, storm cards. |
| Voltron | One creature made huge with auras or equipment. | Commander-common. |
| Aristocrats | Sacrifices own creatures for value. | Sacrifice outlets, death triggers. |
| Tokens (go-wide) | Many small creatures plus anthems. | Token makers, overrun effects. |
| Kindred (tribal) | One creature type and its lords. Elves, Goblins, Zombies, Dragons, and more. | The word "Kindred" replaced "Tribal" in 2024. |
| Spellslinger | Instants and sorceries with payoffs. | Prowess, magecraft, copy effects. |
| Blink (flicker) | Exiles and returns own creatures for repeated ETB value. | |
| Lifegain | Gains life and uses it as a resource or payoff. | Ajani's Pridemate style cards, lifelink, drain. |
| Mill | Puts cards from the opponent's library into the graveyard. | |
| Landfall / lands | Lands entering as the engine. | |
| +1/+1 counters | Counters as the theme. | |
| Group hug / chaos | Casual Commander themes. Ask the bracket. | |
| Superfriends | Many planeswalkers. | |
| Wheels | Effects that make each player discard and draw seven. | |
| Enchantress | Enchantment-based card draw. | |
| Artifacts / affinity | Artifact synergies. | |
| Big mana / Tron | Assembles special lands for huge mana. | Modern Tron. |
| Burn | Direct damage to the face. | Red. |
| Hatebears | Small creatures with taxing effects. | |

The Scryfall Oracle tags file maps most themes to real card lists. Example: tag `lifegain` has 881 direct cards and 3,374 with children (bulk file of 2026-08-23, method not recorded, unverified). A live `otag:lifegain` search returns 2,596 cards (2026-08-24). Use the tag hierarchy to expand a theme word into candidate cards.

## 6. Card roles in a deck

Every deck needs cards in these roles. The ratios differ by format and archetype. These are guide numbers, not rules.

| Role | 60-card guide | Commander (99) guide |
|---|---|---|
| Lands | 20-27 (see archetype) | 34-38 |
| Ramp (rocks, dorks, land search) | 0-8 | 8-12 |
| Card draw / advantage | 4-8 | 8-12 |
| Targeted removal | 4-10 | 6-10 |
| Board wipes | 0-4 (control) | 2-5 |
| Threats / win conditions | 8-20 | 10-20 |
| Interaction (counters, protection) | 0-12 | 4-10 |
| Synergy / theme pieces | rest | 20-35 |

Mana curve: most decks want the most cards at MV 2 and 3, fewer at 1 and 4, few at 5+. Commander decks average MV 2.8 to 3.5. Aggro 60-card decks average under 2.5.

Color sources: for a two-color 60-card deck, aim for 12 to 14 sources of each main color. Use the Frank Karsten tables for exact counts. Commander with three or more colors needs many dual lands and fixing.

Sideboard (60-card formats): 15 cards. Answer the expected metagame. Not used in Commander.

## 7. Slang glossary

| Term | Meaning |
|---|---|
| Bolt | Lightning Bolt, or any 3 damage for one mana. "Bolt test": does a creature die to it? |
| Wrath / sweeper / board wipe | Destroys all creatures. |
| Tutor | Searches the library for a card. |
| Cantrip | A cheap spell that draws a card. |
| Dork | A creature that makes mana. Rock: an artifact that makes mana. |
| Fixing | Lands or spells that give the right colors. |
| Curve out | Play one spell per turn on curve. |
| Flood / screw | Too many lands / too few lands. |
| Top-deck | Draw the card you needed. |
| Chump block | Block with a creature that dies for no value. |
| Race | Both players attack, nobody blocks. |
| Hate piece / hate card | A card that stops a strategy. Graveyard hate, artifact hate. |
| Sac outlet | A permanent that lets you sacrifice creatures at will. |
| Pump | An effect that raises power and toughness. |
| Bounce | Return to hand. |
| Blink / flicker | Exile and return. |
| Value | Card advantage from one card over time. |
| Tempo | Time advantage. |
| Wincon | The card or combination that wins. |
| Precon | A preconstructed deck sold by Wizards. Most fit Bracket 2, but the bracket is no longer tied to precons (2025-10-21). |
| Pod | A Commander table of players. |
| Rule 0 | The pre-game talk about power and house rules in Commander. |
| Pubstomp | Bring a strong deck to a weak table. Avoid it. |
| Salt | Player frustration. "Salty cards" cause it: land destruction, extra turns, stax. |
| Jank | Weak but fun cards. |
| Staple | A card that most decks in its colors play. |
| Cut | Remove a card from the deck. "What do I cut?" is the most common question. |
| Sol Ring | The one-mana artifact that every Commander deck runs. Not a Game Changer. |
| Boros, Golgari, ... | Guild names for color pairs (see section 4). |
| MV / CMC | Mana value. The old name is converted mana cost. |
| ETB, LTB, dies, cast trigger | Trigger types. |
| Poison | Ten poison counters lose the game. |
| Commander damage | 21 combat damage from one commander loses. |
| cEDH | Competitive Commander. Bracket 5. |
| FNM | Friday Night Magic. Store-level competitive play. |
| LGS | Local game store. |
| Meta / metagame | The set of decks people play right now, and their shares. |
| Tier 1 / tier 2 | The strongest decks in the meta. |
| Netdeck | Copy a list from the internet. |
| Brew | Make an original deck. |
| Budget | A deck under a price cap. Ask for the cap. |
| Proxy | A stand-in for a card you do not own. Legal only in casual play. |

## 8. Keywords

The full list of 223 keyword abilities, 79 keyword actions, and 69 ability words is in `references/scryfall-catalogs.md`. The most common evergreen keywords, with Oracle-card counts from the snapshot of 2026-08-24 (the `keywords` field of each object in `oracle_cards`, art-series, token, and emblem layouts left out):
- Flying 3,361 · Trample 1,055 · Vigilance 768 · Haste 699 · Flash 636
- Reach 449 · Menace 417 · First strike 413 · Lifelink 401 · Deathtouch 360
- Defender 314 · Ward 217 · Protection 213 · Double strike 132 · Indestructible 116 · Hexproof 105

Deciduous mechanics: Scry, Surveil, Mill, Fight, Food, Treasure, Clue, Blood, Map, Equip, Crew, Cycling, Kicker, Flashback, Landfall, Transform, Adventure, Day/Night, Monarch, Initiative.

## 9. ManaBox export format

ManaBox exports a CSV. A real whole-collection export (verified 2026-08-24 against the owner's file) has 18 columns: `Binder Name, Binder Type, Name, Set code, Set name, Collector number, Foil, Rarity, Quantity, ManaBox ID, Scryfall ID, Purchase price, Misprint, Altered, Condition, Language, Purchase price currency, Added`. A single-list export can drop the binder columns. Parse by header name, never by position.

Values: `Foil` is `normal`, `foil`, or `etched`. `Condition` is `mint`, `near_mint`, `excellent`, `good`, `light_played`, `played`, `poor`. `Language` is a code such as `en`, `ja`, `zh_CN`.

Language: the app supports English only (D-23). Rows with another language code are reported to the user and skipped.

The `Scryfall ID` column is the join key. It identifies one printing. Map it to `oracle_id` to count copies of one card across printings. Fallback when the ID is missing: `Set code` plus `Collector number`, then `Name` plus `Set name`.

ManaBox also exports decks as text in the MTG Arena format: `4 Lightning Bolt (STA) 42`. The app should import both.

## 10. Scryfall facts for the app

- API base: `https://api.scryfall.com`. Requests need a real `User-Agent` and an `Accept` header.
- Rate limits (2026-08-24): `/cards/search`, `/cards/named`, `/cards/random`, `/cards/collection`: 2 per second. `/cards/manifest`: 10 per minute. Other endpoints: 10 per second. HTTP 429 blocks you for 30 seconds. Repeated overload gets a ban.
- Bulk data (collected every 12 to 24 hours, no rate limit on `*.scryfall.io`): `oracle_cards` (24.5 MB gz, one card per Oracle ID, 38,626 rows), `default_cards` (77.5 MB gz, every printing in English), `all_cards` (392 MB gz), `rulings` (5.4 MB gz), `oracle_tags` (5.9 MB gz, 4,522 tags), `art_tags`, `unique_artwork`.
- Update cadence: prices once per day. Gameplay data less often. Download bulk once per day. Bulk prices are stale after 24 hours.
- Images: `image_uris` keys `small, normal, large, png, art_crop, border_crop`. Double-faced cards have `card_faces[].image_uris`. Show the artist and copyright. Do not crop, skew, or watermark.
- Data license: Wizards Fan Content Policy through Scryfall. No paywall on card data. No repackaging without added value. Do not imply Scryfall endorsement.
- Prices: `prices.usd`, `usd_foil`, `usd_etched` are TCGplayer near-mint market estimates, once per day. No condition tiers. The app shows a 7-day rolling average with outliers removed, labeled "NM market estimate" (D-17).
- Card fields the builder needs: `name, oracle_id, id, mana_cost, cmc, colors, color_identity, type_line, oracle_text, keywords, legalities, game_changer, edhrec_rank, penny_rank, rarity, set, collector_number, prices, card_faces, layout, produced_mana, power, toughness, loyalty`.

## 11. Clarifying-question catalog

The agent asks only what the prompt did not answer. Never ask more than three questions in one turn. Order by impact.

Neither owned-commander row survives. The not-owned row and the weak-pool row both sat on the commander key. A delegated commander fills that key, so a user who says "you pick" silences them. That user is the one who needs them. PR-8 reports both after the build, from the library it actually read (D-226, D-232).

Five more rows retired on 2026-08-28 (A-6 of the audit): theme (card named), jank, meta, plan choice, and locked cards. No function and no planned PR read their answers. The card names for locked cards survive: the classifier names them, and the build keeps every one (D-242). A jank word routes to Power, and the meta row returns if the sideboard builder opens.

| Slot | Ask when the slot is empty | Example question |
|---|---|---|
| Out of scope | The user asked for something this app does not build, such as a deck for another card game (D-99). Ask this alone, before every other row. The row closes when the user then asks for a Magic deck: a filled format, theme, color, or commander slot is the answer. | "I build Magic: The Gathering decks only. Would you like one instead?" |
| One deck at a time | The user asked for more than one deck (D-112). Ask this alone. Read one message for it, and never the whole conversation. The row closes when the next message names one deck. | "I build one deck at a time. Which deck do you want first?" |
| Format (not supported) | The user named a format this app does not build, such as Brawl (D-112). The row asks again only for another unsupported format. The same format twice gets one sentence (D-210). | "I do not build {bad_format}. The nearest format I build is {near_format}. Shall I use that?" |
| Format (no substitute) | The user named an unsupported format with no nearest format to offer (D-146, D-155). Historic, Timeless, Pioneer, Legacy, Vintage, Pauper, and Pauper Commander are the seven. The row asks again only for another unsupported format (D-210). | "I do not build {bad_format}. Which format should I build instead: Commander, Standard, or Modern?" |
| Format | Always, unless stated. Ask this first. Every other slot depends on it. The app builds three (D-155). | "Which format: Commander, Standard, or Modern?" |
| Format (store event) | The user names FNM, an LGS, a store, or an event. The row names the three formats as the ones it builds, and never as the ones the event runs (D-213). | "I build Standard, Modern, and Commander. Which one does your event run?" |
| Theme or plan | The prompt gives only a format. | "What should the deck do: a creature type, a mechanic, or a play style?" |
| Theme (competitive) | Power is FNM or tournament-meta. | "Do you want a named tier-one deck, or the best deck under your budget?" |
| Named card role | The user named one card, and the format is Commander or empty. The row closes when the user says the card is not the commander, for example "build around X, but not as my commander" (D-70). A card the index says can not lead a deck has a settled role, so the row does not fire (D-220). | "Do you want {card} as your commander, or as one card in the 99?" |
| Commander | The format is Commander and no commander is given. | "Which commander do you want? Name one, or I suggest three." |
| Commander (pick) | The user asked the agent to name a commander (D-71). The row asks again only with names the user has not seen: a refusal retires the old ones (D-73, D-80), and the color check drops the names the colors exclude (D-153). The same three names never go out twice (D-163). | "Which one do you want: {A}, {B}, or {C}? Say 'none' and I name three more." |
| Commander (can not lead) | The user named a commander that can not lead a deck, such as Lightning Bolt (D-129). The card index answers it. A legal commander named later closes the row. | "{bad_commander} can not lead a deck. Shall I suggest a commander instead?" |
| Power (Commander) | Always for Commander, unless the user names a bracket. "cEDH" names bracket 5 (D-164). Name no table: the user may build a deck as a gift (D-109). | "Which power bracket should the deck target? 2 is the core level, near a precon, 3 is upgraded, 4 is high power." |
| Power (60-card) | The user asked for no strong deck. Ask again when the user names a step and also says competitive, strong, best, or serious. Those words conflict with the named step. | "How strong should this be: casual, FNM level, or tournament-meta?" |
| Colors | The user gave no preference. Never ask when a commander is set. The color identity fills this slot. "Colorless" is an answer, and it closes the slot (D-165). State no fact about which colors are strongest (D-108). | "Any color preference?" |
| Card pool (precon) | The user asked to upgrade a precon (D-113). This row replaces the row below, and it names the precon. The build keeps 85 percent of the precon's cards, and it reads the product name from the user's own words (D-218, D-247). | "Should I build from your {precon} precon first, use only cards from it, or ignore it for a fully optimized deck?" |
| Card pool | A collection is attached (D-37), and the format, the colors, and the theme are filled (D-67). | "Build from your library first, only your library, or ignore it for a fully optimized deck?" |
| Card pool (thin theme) | `ThinTheme` is set (D-63). This row replaces the row above. | "Your library holds {n} {theme} cards. I want 30 or more. Build owned-first with a buy list, or use the whole pool?" |
| Budget | The user mentions cost, a buy list is needed, or the pool mode is any-card. A session with no collection always needs a buy list, because the user owns nothing to build from (D-168). "Money is no object" answers the row (D-168). | "Is there a budget for cards to buy?" |
| Budget scope | A collection is attached and the user named one number, and the message did not say which the cap covers. A phrase that names the buy list answers it, and a proxy rule is not a budget at all (D-253). The answer is stored, and the build reads it: a cap on the cards to buy checks what the user must acquire, and a cap on the whole deck checks every copy (D-238). | "Is that a cap on the cards you buy, or on the whole deck value?" |
| House rules | "Anything goes", "kitchen table", or "no ban list". "Casual" alone does not fire this row (D-78). "Proxy" and "whatever" do not fire it either (D-111). A negation stops every trigger word. The row names no format, because Vintage is not one the app builds (D-155). The answer is stored in the user's own words, and the build copies it to the deck's format (A-6). | "When you say anything goes, do you mean any card with no ban list?" |
| House format limits | House rules set a house format, and the user answered the house-rules question (D-81). The row goes out word for word: a rewrite reads as three questions (D-162). The row names no list of limits, because a list reads as one question for each item (D-212). | "Inside your house format, do the normal 60-card deck limits hold?" |

The agent does not ask the user to confirm a power step it inferred. A user who asks for the strongest deck gives the answer, and a question about it repeats the answer (D-216). The agent fills the tournament step, closes the slot, and marks the step as inferred. The plan states the step, and the user can change it.

Do not ask where the user buys, or by what date they need the cards. The app can not act on either answer. It holds no store stock and no delivery times, and Scryfall gives a price estimate, not availability (D-87).

Do not ask whether the table accepts mill, land destruction, extra turns, or stax. The power level answers the same need (D-110). The salt list of section 7 stays as reference material, and it drives no question.

An out-of-scope request gets one question and no others. Gate run 11 of 2026-08-25 asked "Which Yu-Gi-Oh format would you like?", because the catalog held no way to decline (D-99).

Ask order (from the PR-7 dogfood runs, 2026-08-24): format, theme, house rules, commander, power, colors, card pool, budget. Ask the card pool after the format, the colors, and the theme (D-67). A special row beats its general row: ask "Format (store event)" before "Format", and "Card pool (precon)" before "Card pool". Never ask a slot that another slot already fills. `internal/questions` holds this order as data.

Word routing: "anything goes", "kitchen table", and "no ban list" route to House rules (D-3). "Casual" alone routes to Power, not to House rules. The gate run of 2026-08-25 asked a parent about house rules for a child's deck, and the parent answered "casual means low power, not a house format" (D-78). "Strongest", "competitive", "best", and "serious" route to Power.

An occasion routes nowhere. "For an event" and "at my store" name a place or a happening, and they name no power step (D-219). "Janky", "fun", "silly", and "meme" route to Power, because the jank row retired (A-6). Do not route a jank word to House rules. House rules cover legality.

"Proxy" routes to Budget, and not to House rules (D-111). A user who proxies every card has no budget, so the agent asks no budget question. The word says nothing about which cards are legal. "Whatever" routes nowhere. Gate runs 11 to 13 read "whatever is winning" and "whatever you think is best" as house rules, six times.

Negation rule (D-111): a negator before a trigger word stops that trigger. "No proxies" is not a proxy user. The negators are "no", "not", "never", "without", and the short negative verb forms. One shape is exempt: "not as my commander" denies the role of a card, and it names the Commander format.

Format inference (D-116): read the format from an adjective, such as "a Commander deck" or "a Modern burn deck". "EDH" means Commander. A message with "my commander", "in the 99", "bracket 3", or "my precon" means Commander, even with no format word. Gate run 11 asked conversation 23 for the format. The user had written "A land destruction Commander deck."

A message that names a format this app does not build names no format. That holds even when the message holds a format word. "Duel Commander" and "Pauper Commander" are not Commander. The unsupported-format row declines them (D-112).

Question wording rules (D-109, D-116): add no clause that only repeats a value the user gave. Keep a clause that narrows the question. State no fact about the game inside a question. Presume no table, no playgroup, and no event that the user did not name.

Slot rules: a slot stays open through the question phase. A later answer replaces an earlier one, and the agent states the change. A slot change after a build is an ordinary turn, and the next build reads the new value (D-241). A changed format retires every question that is out (D-125, D-126). A retired question leaves the asked state, so the session can report ready, and the row that asked it does not ask again.

Commander rules: a name the user gives as the commander closes every commander row (D-71), the "can not lead" row included. The three rows ask one thing in different words. The pick row is the one exception to the no-repeat rule. It asks again only when the names on the table change. The same three names stay on the table until the user asks for others, and a retired name never comes back (D-73, D-80). The row never sends one list twice (D-163).

A superlative such as "buy the best lifegain commander" hands the choice to the agent, as "you pick" does (D-147, D-167).

Locked-card rule: a card that becomes the commander is not a locked card (D-70). The locked row retired on 2026-08-28 (A-6). The classifier names the locked cards, the session state holds them, and the build keeps every one (D-242). No question asks whether the deck may cut one.

Question source rule (D-25): use a catalog question when one fits the empty slot. Compute a gap score: how well the best catalog question matches the slot and the user's words. When the score is below the threshold, invent a question and log it with the score. The owner scores each invented question on the six-field rubric (D-66). Invented questions that repeat become catalog candidates.

Decline rule (D-93): a user can hand any slot back. "Any colors are fine", "you decide", and "surprise me" are declines. A decline names no value, and the slot goes to the skipped state. The agent does not ask again, and the generator applies the default below. The agent must ask the question before the user can decline it.

Default answers when the user says "you decide": format Commander (the most played format in 2026), bracket 2 to 3, colors from the collection's strongest overlap with the theme. Pool mode: owned-first when a collection is attached, any-card when none is (D-37). A user without a collection never gets the card-pool question.

## 12. Validation checklist (deterministic, run after every build)

1. Deck size matches the format.
2. Copy limits respected (4, singleton by name, per-card overrides such as Seven Dwarves). No restricted list applies, because Vintage left the app (D-155).
3. Every card legal in the format on the query date, per Scryfall legalities.
4. Commander: every card inside the color identity. Commander eligible. A two-commander pair is a valid pairing.
5. Commander: Game Changer count within the bracket. Bracket 1-2: zero.
6. Every card name exists in the card database. No invented names.
7. Ownership (owned modes only, D-37): every card in the collection with enough copies, or listed as an acquisition. Basic lands are exempt. Owned-first gives a warning, owned-only a block. In any-card mode this check is off, and ownership marks are information.
8. Sideboard size (15 max in 60-card formats, none in Commander). House format (D-3) skips legality. With a house format the agent states "no legality check applies" in place of a legality date.
9. Companion: the card has the companion keyword and is not banned as a companion. In 60-card formats it sits in the sideboard. In Commander it counts as a 101st card.
10. Land count within the guide range for the archetype (engine advisory).

Model-side checks (PR-8, not the engine):
- Color sources within the guide range.
- Mana curve within the guide range.
- Each role at its minimum count.
- A stated plan that every card serves.

## 13. Sources

- Scryfall API docs: https://scryfall.com/docs/api (fetched 2026-08-23).
- Banned and Restricted 2026-08-10: https://magic.wizards.com/en/news/announcements/banned-and-restricted-august-10-2026
- Commander B&R 2026-02-09: https://magic.wizards.com/en/news/announcements/commander-banned-and-restricted-february-9-2026
- Commander rules: https://mtgcommander.net/index.php/rules/ Note: its ban page still lists Biorhythm as banned on 2026-08-26. Wizards and Scryfall are the authority, and Biorhythm is unbanned (2026-02-09).
- Commander brackets, Game Changer commander rule: https://magic.wizards.com/en/news/announcements/introducing-commander-brackets-beta (2025-02-11, read 2026-08-26).
- Pauper Commander rules: https://pdhhomebase.com/rules (read 2026-08-26).
- Game Changers list: https://playgroup.gg/commander/game-changers (2026-08-24 update)
- Commander brackets revision: https://magic.wizards.com/en/news/announcements/commander-brackets-beta-update-october-21-2025
- Banned and Restricted 2026-05-18 and 2026-06-29: https://magic.wizards.com/en/news/announcements/banned-and-restricted-may-18-2026 and .../banned-and-restricted-june-29-2026
- Comprehensive Rules 2026-08-07: https://media.wizards.com/2026/downloads/MagicCompRules%2020260819.txt
- Standard sets: https://draftsim.com/mtg-standard-rotation/ (2026-08)
- ManaBox import/export: https://www.manabox.app/guides/collection/import-export/
- ManaBox CSV columns: https://github.com/StepKie/MtgCsvHelper (appsettings.json)
- Metagame pages (August 2026): mtgdecks.net, aetherhub.com, mtggoldfish.com, mtgtop8.com, mtgo.com/decklists
