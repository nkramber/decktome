---
name: mtg-corpus
description: Magic: The Gathering terminology, formats, deck-construction rules, ban-list snapshot, archetypes, and the clarifying-question catalog for the deck-builder agent. Load before you reason about any MtG request.
---

# MtG corpus skill

This skill is the knowledge base for the deck-builder agent. It is also the seed of the app's own corpus. Facts carry a date. Rules and ban lists change. Check the date before you trust a fact.

Snapshot date: **2026-08-23**. Sources: Scryfall API and bulk data, Wizards of the Coast announcements, mtgcommander.net, and the Commander Format Panel.

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
| Life total | 20 in most formats. 40 in Commander. 30 in Brawl multiplayer. |
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

The app must know the format before it builds. A "60-card" request has many possible formats. Ask.

### 2.1 Constructed formats (60-card style)

| Format | Deck size | Copies | Card pool (2026-08-23) | Sideboard |
|---|---|---|---|---|
| Standard | min 60 | max 4 per name (basic lands unlimited) | Sets from Foundations (FDN, 2024-11-15) to The Hobbit (2026-08-14). 12 sets. No rotation in 2026. Next rotation: first set of 2027 (Bloomburrow and Duskmourn leave). | 15 |
| Pioneer | min 60 | max 4 | Return to Ravnica (2012-10) forward. | 15 |
| Modern | min 60 | max 4 | Eighth Edition (2003-07) and Mirrodin forward. Modern Horizons sets included. | 15 |
| Legacy | min 60 | max 4 | All sets. Banned list. | 15 |
| Vintage | min 60 | max 4 | All sets. Banned list plus a restricted list (max 1 copy). | 15 |
| Pauper | min 60 | max 4 | Only cards printed at common (in a paper or MTGO set). | 15 |
| Historic, Timeless, Alchemy | min 60 | max 4 | MTG Arena only. Digital-only cards exist. | 15 |

Scryfall `legalities` keys: `standard, future, historic, timeless, gladiator, pioneer, modern, legacy, pauper, vintage, penny, commander, oathbreaker, standardbrawl, brawl, competitivebrawl, alchemy, paupercommander, duel, oldschool, premodern, predh, tlr`. Values: `legal, not_legal, banned, restricted`.

Legal Oracle-card counts on 2026-08-23: commander 31,830 · vintage 31,690 · legacy 31,672 · modern 22,450 · pioneer 14,817 · pauper 10,793 · standard 4,887.

### 2.2 Commander (EDH)

- Exactly 100 cards, the commander included.
- Singleton: no two cards with the same English name, except basic lands and cards that say otherwise (for example Relentless Rats).
- The commander is a legendary creature, or a card that says it can be your commander. Partner, Partner with, Friends forever, Background, and Doctor's companion allow two commanders.
- Every card must fit the commander's color identity.
- 40 life. Four players is the normal table. 21 combat damage from one commander kills a player.
- The commander goes to the command zone when it would leave. It costs 2 more each time you cast it from there (the "commander tax").
- Rules and ban list: the Commander Format Panel (from 2024-09, owned by Wizards of the Coast with a community panel). Site: mtgcommander.net.
- Commander is a social format. Power level matters more than in any other format. Always ask about the bracket.

### 2.3 Commander brackets (official, beta from 2025-02, updated 2026-02-09)

| Bracket | Name | Expected game | Game Changers | Other limits |
|---|---|---|---|---|
| 1 | Exhibition | 9+ turns | none | No infinite combos, no mass land denial, no extra-turn chains. |
| 2 | Core | 8+ turns | none | No infinite combos, no efficient tutors. Precon level. |
| 3 | Upgraded | 6+ turns | max 3 | No mass land denial. Late-game combos only. |
| 4 | Optimized | 4+ turns | unlimited | Mass land denial, stax, two-card combos allowed. |
| 5 | cEDH | any | unlimited | Only the ban list applies. Plays the best strategy, not a theme. |

Game Changers: 53 cards on 2026-08-23. Scryfall flags them with `game_changer: true`. Examples: Rhystic Study, Cyclonic Rift, Smothering Tithe, Thassa's Oracle, Demonic Tutor, Vampiric Tutor, Ancient Tomb, The One Ring, Gaea's Cradle, Force of Will. Note: Mana Crypt is banned, so it is not a Game Changer.

### 2.4 Other multiplayer or digital formats

Brawl (Arena, 100-card singleton with a commander, Historic pool), Standard Brawl (60-card), Oathbreaker (planeswalker commander, 60 cards), Pauper Commander (common creature commander), Duel Commander (1v1, 20 life). Support these later.

### 2.5 "Anything goes"

Not a defined format. The owner's own meaning: any card, no ban list. Other users mean Vintage, or "Modern but with proxies", or "kitchen table". **Always ask what the user means.** Then record the answer as the session's house rules.

## 3. Ban list snapshot (2026-08-23)

Source of truth at run time: Scryfall `legalities` (updated within a day of each announcement). Never store a ban list in a prompt. Look it up. The facts below are for orientation only.

- 2026-08-10 announcement: Standard banned Badgermole Cub, Stormchaser's Talent, Gran-Gran. Legacy banned The Fantasticar. Vintage restricted The Fantasticar. Next announcement: **2026-10-12**.
- 2026-06-29: Legacy banned Candelabra of Tawnos.
- 2026-05-18: Pioneer banned Cori-Steel Cutter. Alchemy banned Sewer-veillance Cam.
- 2026-03-23: Historic banned Food Chain.
- Commander 2026-02-09: Biorhythm unbanned (added to Game Changers). Lutri, the Spellchaser unbanned but "banned as a companion" (a new category). Still banned: Sundering Titan, Iona, Griselbrand, Mana Crypt, Jeweled Lotus, Dockside Extortionist, Nadu, and about 35 more.
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

The Scryfall Oracle tags file maps most themes to real card lists. Example: tag `lifegain` has 881 direct cards and 3,374 with children (2026-08-23). Use the tag hierarchy to expand a theme word into candidate cards.

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

Color sources: for a two-color 60-card deck, aim for at least 12 sources of each main color. Use the Frank Karsten tables for exact counts. Commander with three or more colors needs many dual lands and fixing.

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
| Precon | A preconstructed deck sold by Wizards. Bracket 2 by design. |
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

The full list of 223 keyword abilities, 79 keyword actions, and 69 ability words is in `references/scryfall-catalogs.md`. The most common evergreen keywords, with Oracle-card counts on 2026-08-23:
- Flying 3,361 · Trample 1,055 · Vigilance 767 · Haste 699 · Flash 636
- Reach 449 · Menace 417 · First strike 413 · Lifelink 401 · Deathtouch 360
- Defender 314 · Ward 217 · Protection 213 · Hexproof · Indestructible · Double strike

Deciduous mechanics: Scry, Surveil, Mill, Fight, Food, Treasure, Clue, Blood, Map, Equip, Crew, Cycling, Kicker, Flashback, Landfall, Transform, Adventure, Day/Night, Monarch, Initiative.

## 9. ManaBox export format

ManaBox exports a CSV. Columns (from the MtgCsvHelper mapping, verified 2026-08-23): `Name, Set code, Set name, Collector number, Foil, Rarity, Quantity, ManaBox ID, Scryfall ID, Purchase price, Misprint, Altered, Condition, Language, Purchase price currency`. A whole-collection export adds the binder or list name. Values: `Foil` is `normal`, `foil`, or `etched`. `Condition` is `mint`, `near_mint`, `excellent`, `good`, `light_played`, `played`, `poor`. `Language` is a code such as `en`, `ja`, `zh_CN`.

Language: the app supports English only (D-23). Rows with another language code are reported to the user and skipped.

The `Scryfall ID` column is the join key. It identifies one printing. Map it to `oracle_id` to count copies of one card across printings. Fallback when the ID is missing: `Set code` plus `Collector number`, then `Name` plus `Set name`.

ManaBox also exports decks as text in the MTG Arena format: `4 Lightning Bolt (STA) 42`. The app should import both.

## 10. Scryfall facts for the app

- API base: `https://api.scryfall.com`. Requests need a real `User-Agent` and an `Accept` header.
- Rate limits: `/cards/search`, `/cards/named`, `/cards/collection`: 2 per second. Other endpoints: 10 per second. HTTP 429 blocks you for 30 seconds. Repeated overload gets a ban.
- Bulk data (daily, no rate limit on `*.scryfall.io`): `oracle_cards` (24.5 MB gz, one card per Oracle ID, 38,626 rows), `default_cards` (77.5 MB gz, every printing in English), `all_cards` (392 MB gz), `rulings` (5.4 MB gz), `oracle_tags` (5.9 MB gz, 4,522 tags), `art_tags`, `unique_artwork`.
- Update cadence: prices once per day. Gameplay data less often. Download bulk once per day.
- Images: `image_uris` keys `small, normal, large, png, art_crop, border_crop`. Double-faced cards have `card_faces[].image_uris`. Show the artist and copyright. Do not crop, skew, or watermark.
- Data license: Wizards Fan Content Policy through Scryfall. No paywall on card data. No repackaging without added value. Do not imply Scryfall endorsement.
- Prices: `prices.usd`, `usd_foil`, `usd_etched` are TCGplayer near-mint market estimates, once per day. No condition tiers. The app shows a 7-day rolling average with outliers removed, labeled "NM market estimate" (D-17).
- Card fields the builder needs: `name, oracle_id, id, mana_cost, cmc, colors, color_identity, type_line, oracle_text, keywords, legalities, game_changer, edhrec_rank, penny_rank, rarity, set, collector_number, prices, card_faces, layout, produced_mana, power, toughness, loyalty`.

## 11. Clarifying-question catalog

The agent asks only what the prompt did not answer. Never ask more than three questions in one turn. Order by impact.

| Slot | Ask when the slot is empty | Example question |
|---|---|---|
| Format | Always, unless stated. | "Which format: Commander, Standard, Modern, or something else?" |
| Commander | Format is Commander and no commander given. | "Do you have a commander in mind, or should I suggest three from your collection?" |
| Power level | Always for Commander (bracket). For 60-card, unless "casual" or "FNM" is clear. | "Which bracket does your table play? 2 is precon level, 3 is upgraded, 4 is high power." |
| Colors | Theme does not imply colors. | "Any color preference? Lifegain is strongest in white and black." |
| Theme or plan | Prompt gives only a format. | "What should the deck do: a creature type, a mechanic, or a play style?" |
| Card pool | A collection is attached (D-37). | "Build from your library first, only your library, or ignore it for a fully optimized deck?" Also ask when the collection is too thin for the plan. |
| Budget | User mentions cost, or a buy list is needed. | "Is there a budget for cards to buy?" |
| House rules | "Anything goes", "casual", "kitchen table". | "What does anything-goes mean at your table: any card with no ban list, or Vintage rules?" |
| Meta | Power is competitive. | "Is this for a specific event or local meta? I can tune the sideboard to it." |
| Variance | User asks for "another version". | "Same plan with different cards, or a different plan in the same colors?" |
| Locked cards | User names cards. | "Should I keep all of those, or can I cut some if they do not fit?" |

Question source rule (D-25): use a catalog question when one fits the empty slot. Compute a gap score: how well the best catalog question matches the slot and the user's words. When the score is below the threshold, invent a question and log it with the score. Invented questions that repeat become catalog candidates.

Default answers when the user says "you decide": format Commander (the most played format in 2026), bracket 2 to 3, colors from the collection's strongest overlap with the theme. Pool mode: owned-first when a collection is attached, any-card when none is (D-37). A user without a collection never gets the card-pool question.

## 12. Validation checklist (deterministic, run after every build)

1. Deck size matches the format.
2. Copy limits respected (4, singleton, restricted 1).
3. Every card legal in the format on the query date, per Scryfall legalities.
4. Commander: every card inside the color identity. Commander eligible.
5. Commander: Game Changer count within the bracket. Bracket 1-2: zero.
6. Every card name exists in the card database. No invented names.
7. Ownership (owned modes only, D-37): every card in the collection with enough copies, or listed as an acquisition. In any-card mode this check is off, and ownership marks are information.
8. Land count and color sources within the guide range for the archetype.
9. Mana curve within the guide range.
10. Each role (ramp, draw, removal, wincon) has at least the minimum count.
11. The deck has a stated plan, and every card serves it or a role.

## 13. Sources

- Scryfall API docs: https://scryfall.com/docs/api (fetched 2026-08-23).
- Banned and Restricted 2026-08-10: https://magic.wizards.com/en/news/announcements/banned-and-restricted-august-10-2026
- Commander B&R 2026-02-09: https://magic.wizards.com/en/news/announcements/commander-banned-and-restricted-february-9-2026
- Commander rules: https://mtgcommander.net/index.php/rules/
- Game Changers list: https://playgroup.gg/commander/game-changers (2026-08-24 update)
- Standard sets: https://draftsim.com/mtg-standard-rotation/ (2026-08)
- ManaBox import/export: https://www.manabox.app/guides/collection/import-export/
- ManaBox CSV columns: https://github.com/StepKie/MtgCsvHelper (appsettings.json)
- Metagame pages (August 2026): mtgdecks.net, aetherhub.com, mtggoldfish.com, mtgtop8.com, mtgo.com/decklists
