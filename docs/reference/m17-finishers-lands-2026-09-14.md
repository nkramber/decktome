# M-17, finishers, lands, conditions, and the theme miss, 2026-09-14

The roadmap gate of M-17 reads:

> Gate: a dated document holds every count, and the owner reads the floors before PR-52 starts.

**Verdict: the gate holds.** This document holds every count, and the owner chose the target and the floor of PR-53 (D-726). Every number here is free. No run called a provider.

**The replay found the cause of F-141.** The theme "opponent milling cards" matched no card, because `themes.json` holds a row for "mill" and no row for "milling" (F-142). So the owned-only shortlist held staple roles alone, and the mill cards had no theme signal to reach it. The same miss kept every threat off the shortlist, and it explains most of F-138 and F-140 in this deck. The owner chose PR-54 first, a fix of the theme words (D-723 to D-725).

## The method

A scratch Go test inside `go/cmd/deck-gate` made every count, and a copy sits in `.local/m17/` out of git. It reads the card snapshot of 2026-09-04 and the local meta store of 2026-09-14.

The test reads five groups of real Commander lists. No list carries a bracket, so each group stands for a bracket, as in M-16.

| Group | Source | Lists | Stands for |
|---|---|---|---|
| Precons of the repository | `precons.Decklists` | 9 | Bracket 2 |
| Precons since 2023 | MTGJSON, tier baseline, released on 2023-01-01 or later | 90 | Brackets 2 and 3 |
| EDHREC average decks | EDHREC, tier typical | 1,040 | Brackets 2 and 3 |
| TopDeck good | TopDeck, tier good, from 2026-08-01 | 7,150 | Bracket 5 |
| TopDeck top cut | TopDeck, tier great, from 2026-08-01 | 2,567 | Bracket 5 |

A list must resolve its commanders, miss 2 cards or fewer, and hold 100 cards. The size check dropped 445 good lists, 199 top-cut lists, and 1 EDHREC deck. The commander check dropped 4 EDHREC decks, 2 precons, and 1 top-cut list.

The test also reads 26 built decks. They are the 19 Commander decks of deck gate run 19, the 6 decks of bracket gate run 6, and deck `sFLbEUuKzI0QyLPft0zI`. Each table cell with three numbers gives the low quarter, the median, and the high quarter, in that order.

## The replay of the Gríma request (F-141)

The replay builds the shortlist of session `z1hshyY6Npig1FN2NuV7` with the collection of the owner, and it calls no model. The request reads Commander, the theme "opponent milling cards", Gríma, Saruman's Footman, owned only, and bracket 5. It builds the request as `deck-gate` does, and `deck-gate` builds the shortlist of the app (D-706).

The replay shortlist holds 165 cards, and the stored deck recorded 168. The local snapshot and model differ from the deployed ones, and that explains a small gap.

### The theme matched no card

- The theme words read "opponent" and "milling". The stop words removed "cards".
- `themes.json` holds a row for "mill", and it holds no row for "milling". So "milling" took the generic rule of `theme.go`: slugs such as `milling-matters`, a keyword and a subtype named Milling, and the text needle "milling".
- No such slug exists, and Oracle text writes "mills". So no signal of "milling" fired.
- The noise rule of D-411 removed the text needle "opponent".
- The result reads 0 cards on theme, and both words unmatched.

### The shortlist held staple roles alone

A card with no theme signal stays on the shortlist only when its role is a staple role (`candidates.go`). No card of the pool held a theme signal, so each card needed a staple role. The uncapped build lifts every cap and shows each card that passed that rule.

| Role | Shortlist | Uncapped build |
|---|---|---|
| Land | 40 | 71 |
| Ramp | 32 | 57 |
| Removal | 31 | 102 |
| Draw | 30 | 152 |
| Interaction | 20 | 43 |
| Wipe | 12 | 24 |
| Threat, synergy, and wincon | 0 | 0 |

Every card of the deck sat on the replay shortlist, except Island and Swamp. So the model chose from staple roles alone.

### The owned mill cards

The owned blue-black pool holds 14 cards under the tag tree of `mill-opponent`.

- One card reached the shortlist: One Ring to Rule Them All, with the role wipe. The deck did not take it.
- Two cards reached the uncapped build and fell to the role caps: Pelargir Survivor as ramp, and The Master of Lake-town as draw.
- Eleven cards never reached the uncapped build, because each held no theme signal and no staple role.

So the shortlist dropped the mill cards, and not the model.

### The lands of the replay

The land cap holds 40 lands. `capLands` fills half of the cap with the lands that make the most deck colors (D-450). Then it adds the lands of the theme, and then the mana order again.

- The shortlist kept the 36 owned lands that make both deck colors, and 4 lands that make one deck color.
- Of the 36 lands, 2 count as untapped duals: Command Tower and City of Brass. Sunken Hollow enters untapped on a condition. 11 lands enter tapped, and 22 lands have a condition or a cost on their colored mana.
- Sunken Hollow ranked fourth and reached the shortlist, and the model left it out. The model took Thriving Isle, Thriving Moor, Exotic Orchard, Plaza of Heroes, and Spire of Industry.
- No fetch land reached the shortlist. Scryfall lists no produced mana for a fetch land, so it ranks under every land that makes a color.
- Marsh Flats, Fabled Passage, Evolving Wilds, Terramorphic Expanse, Ash Barrens, and Reliquary Tower fell under the cap.

### The theme probe

The probe builds a Commander shortlist in the colors of Gríma for each theme, with any card allowed, at bracket 3. The count reads the cards of the pool with a theme signal.

| Theme | Cards on theme | Theme | Cards on theme |
|---|---|---|---|
| milling | 0 | mill | 823 |
| opponent milling cards | 0 | mill my opponents | 1,035 |
| blinking | 0 | blink | 144 |
| flickering | 0 | flicker | 144 |
| reanimating | 0 | reanimator | 1,453 |
| sacrificing | 14 | sacrifice | 698 |
| stealing | 0 | theft | 417 |
| gaining life | 0 | lifegain | 868 |
| token | 43 | tokens | 1,185 |

Three more themes matched no card: "draining", "wheeling", and "self-mill". The theme "drain" matched 6 cards.

## Lands in real and built decks

Each land takes one class, and the class reads the deck colors:

- **Basic**: a land with the supertype Basic.
- **Fetch**: a land whose text searches the library for a land.
- **Untapped dual**: a land that makes two or more deck colors and enters untapped. A shock land counts here.
- **Enters untapped on a condition**: a dual whose text holds "enters tapped" with "unless" or "if you don't". A check land, a fast land, and a battle land count here.
- **Mana with a condition or a cost**: a dual whose text holds "spend this mana only", "activate only if", "could produce", "among", or `{1}, {T}: Add`. Exotic Orchard, Plaza of Heroes, and Opal Palace count here.
- **Tapped dual**: a dual that enters tapped with no condition.
- **Utility**: each other nonbasic land. It makes one deck color or none.

The two classes with a condition together make the "land with a condition" of the plan. A deck of one color holds no dual under this rule, so the tables read decks of two colors and more.

### Decks of two colors

| Class | Precons since 2023 (24) | EDHREC average (387) | TopDeck good (2,023) | TopDeck top cut (706) |
|---|---|---|---|---|
| All lands | 37, 38, 38 | 35, 35, 36 | 25, 27, 28 | 25, 27, 28 |
| Basic | 16, 18, 23 | 19, 21, 22 | 2, 2, 4 | 2, 2, 3 |
| Fetch | 1, 2, 2 | 0, 1, 1 | 4, 6, 7 | 5, 6, 7 |
| Untapped dual | 1, 2, 3 | 3, 4, 4 | 6, 8, 9 | 6, 8, 9 |
| Enters untapped on a condition | 2, 2, 4 | 4, 4, 5 | 1, 1, 2 | 1, 1, 2 |
| Mana with a condition or a cost | 1, 3, 3 | 1, 1, 2 | 0, 1, 2 | 0, 1, 2 |
| Tapped dual | 2, 4, 6 | 2, 2, 3 | 0, 0, 1 | 0, 0, 1 |
| Utility | 3, 4, 7 | 2, 3, 4 | 5, 7, 9 | 5, 7, 9 |

### Decks of three colors or more

| Class | Precons since 2023 (64) | EDHREC average (165) | TopDeck good (4,305) | TopDeck top cut (1,568) |
|---|---|---|---|---|
| All lands | 37, 38, 38 | 35, 36, 36 | 25, 27, 28 | 25, 26, 28 |
| Basic | 11, 14, 16 | 13, 15, 16 | 0, 0, 1 | 0, 0, 0 |
| Fetch | 1, 2, 3 | 1, 3, 3 | 7, 8, 9 | 7, 8, 9 |
| Untapped dual | 1, 3, 5 | 4, 6, 7 | 9, 10, 12 | 9, 11, 12 |
| Enters untapped on a condition | 4, 5, 8 | 3, 6, 7 | 0, 1, 1 | 0, 1, 1 |
| Mana with a condition or a cost | 2, 4, 4 | 1, 1, 2 | 1, 1, 2 | 1, 1, 2 |
| Tapped dual | 5, 6, 8 | 2, 3, 4 | 0, 0, 0 | 0, 0, 0 |
| Utility | 1, 3, 5 | 1, 2, 3 | 3, 4, 6 | 3, 4, 6 |

### Built decks at brackets 4 and 5

| Deck | Bracket | Colors | Lands | Basic | Fetch | Untapped dual | Enters untapped on a condition | Mana with a condition or a cost | Tapped dual | Utility |
|---|---|---|---|---|---|---|---|---|---|---|
| Run 6, Urza, Lord High Artificer | 4 | 1 | 34 | 29 | 0 | 0 | 0 | 0 | 0 | 5 |
| Run 6, Korvold, Fae-Cursed King | 4 | 3 | 33 | 17 | 0 | 8 | 5 | 3 | 0 | 0 |
| Run 6, Prosper, Tome-Bound | 4 | 2 | 33 | 14 | 0 | 9 | 3 | 6 | 0 | 1 |
| Run 6, Kinnan, Bonder Prodigy | 5 | 2 | 31 | 13 | 9 | 6 | 3 | 0 | 0 | 0 |
| Run 6, Yuriko, the Tiger's Shadow | 5 | 2 | 30 | 15 | 0 | 8 | 3 | 4 | 0 | 0 |
| Run 6, Najeela, the Blade-Blossom | 5 | 5 | 31 | 21 | 0 | 7 | 0 | 2 | 1 | 0 |
| `sFLbEUuKzI0QyLPft0zI`, Gríma, Saruman's Footman | 5 | 2 | 33 | 26 | 0 | 2 | 0 | 3 | 2 | 0 |

- The median TopDeck list plays 2 basic lands in two colors and none in three colors or more. The built decks at bracket 5 play 13 to 26.
- The TopDeck lists play 6 to 8 fetch lands at the median. Of the built decks at brackets 4 and 5, only Kinnan plays one or more.
- In the TopDeck lists, untapped duals lead at the median, and fetch lands and utility lands follow. Tapped duals come last.
- The precons play 4 to 6 tapped duals at the median, and the TopDeck lists play none.

## Finishers in real and built decks

### The tags

M-17 reads nine Tagger slugs: alternate-win-condition, burn-player-each, drain-life, mill-opponent, poison-opponents, extra-combat-phase, gives-double-strike, damage-multiplier, and overrun. An evasive creature of power 5 or more also counts. A card is evasive when it sits under the Tagger tree of `evasion`.

The tree of a slug holds its child tags too, and two trees count cards that win no game:

- The tree of `mill-opponent` holds Ragavan, Nimble Pilferer, Brain Freeze, and Grinding Station. The top-cut lists play them in 47, 33, and 12 percent. The three cards sit under the child tag `mill-any`.
- The child tag `blood-artist-ability` of `drain-life` holds Blood Artist and Zulaport Cutthroat. Blood Artist sits in 10 percent and Zulaport Cutthroat in 7 percent of EDHREC average decks, and each drains as a finisher.
- The slug `alternate-win-condition` holds Thassa's Oracle, and 31 percent of top-cut lists play it. So the count at bracket 5 reads a combo finisher most of the time.

So the count reads the nine parent slugs with no child tags, and it adds `blood-artist-ability`. This document calls that set the curated count.

### The counts

| Group | Tag trees | Curated count | Curated count with evasive creatures | Lists with one or more, with evasive creatures |
|---|---|---|---|---|
| Precons of the repository (9) | 2, 2, 3 | 1, 2, 3 | 1, 3, 4 | 100% |
| Precons since 2023 (90) | 1, 2, 4 | 1, 2, 3 | 2, 3, 5 | 96% |
| EDHREC average (1,040) | 1, 2, 5 | 1, 2, 4 | 1, 3, 6 | 88% |
| TopDeck good (7,150) | 1, 2, 3 | 0, 1, 1 | 1, 1, 2 | 78% |
| TopDeck top cut (2,567) | 1, 2, 3 | 0, 1, 1 | 1, 1, 2 | 78% |

### The built decks

The curated count with evasive creatures reads:

- Deck gate run 19, 19 Commander decks: prompts 12 and 17 hold 0, and prompts 3, 5, 22, and 23 hold 1. The other 13 decks hold 2 to 10.
- Bracket gate run 6: Urza holds 0, Yuriko 1, Kinnan 3, Prosper 4, Najeela 5, and Korvold 9. The one card of Yuriko is Thassa's Oracle.
- Deck `sFLbEUuKzI0QyLPft0zI` holds 1: Grave Venerations. It drains 1 life when a creature of the deck dies, and the deck holds 3 creatures.

### The target and the floor of PR-53 (D-726)

The owner chose these numbers. The target is the count that the prompt asks for. The floor is the count under which the deck earns a finding.

| Bracket | Target | Floor | Source |
|---|---|---|---|
| 1 to 4 | 3 | 2 | The target reads the median of the precons since 2023 and of the EDHREC average decks. The floor reads the low quarter of the precons since 2023. |
| 5 | 1 | 1 | Both numbers read the median of the TopDeck top cut. |

- At least a quarter of the EDHREC average decks hold 1 or fewer, so they read under the floor of 2.
- About 1 in 5 top-cut lists hold 0, so they read under the floor of 1.
- No real group stands for bracket 4.
- Under these floors, 7 of the 22 built decks at brackets 1 to 4 read under the floor. None of the 4 built decks at bracket 5 reads under it.
- Deck `sFLbEUuKzI0QyLPft0zI` is a bracket 5 deck and holds 1, so it meets the floor. A floor alone does not catch this deck.

## Card conditions (F-140)

### The rule

A rules-text rule marks a card that needs one of three things:

- **Creatures**: an Equipment, a Vehicle, or text that holds "creature you control", "creatures you control", "choose a creature type", "sacrifice a creature", or "untapped creature".
- **Tokens**: text that holds "created a token", "whenever you create", "token you control", "tokens you control", "for each token", or "whenever a token".
- **Artifacts**: text that holds "artifact you control", "artifacts you control", "control an artifact", "three or more artifacts", "affinity for artifacts", "for each artifact", "improvise", or "metalcraft".

Each class has its enablers in the 99:

- **Creatures**: the creature cards, and the cards that make creature tokens.
- **Tokens**: the cards that make a token.
- **Artifacts**: the artifact cards, and the cards that make artifact tokens.

A condition fails when the deck holds fewer enablers than the line of its class. The line is the lowest low decile over the four large real groups. Each decile reads only the lists that play a card of the class.

| Class | Precons since 2023 | EDHREC average | TopDeck good | TopDeck top cut | Line |
|---|---|---|---|---|---|
| Creatures | 28 (90 lists) | 22 (1,034 lists) | 11 (7,142 lists) | 13 (2,565 lists) | 11 |
| Tokens | 11 (31 lists) | 7 (262 lists) | 6 (551 lists) | 7 (186 lists) | 6 |
| Artifacts | 12 (33 lists) | 11 (312 lists) | 14 (5,300 lists) | 14 (1,881 lists) | 11 |

### The built decks

Of the 26 built decks, only deck `sFLbEUuKzI0QyLPft0zI` reads under a line. It holds 8 creature enablers: 3 creature cards and 5 cards that make creature tokens. The rule marks 19 of its cards as cards that need creatures.

The rule counts the commander as no enabler, and no creature can block Gríma. So several of the 19 cards serve Gríma itself: Lightning Greaves, Swiftfoot Boots, Champion's Helm, Darksteel Plate, Magic Damper, and Octopus Form. The Equipment of the deck can also go on Gríma.

The four cards that the review named read this way:

- Skullclamp draws two cards when the equipped creature dies, so it needs a creature that the deck can lose.
- Springleaf Drum taps an untapped creature, so Gríma can pay that cost only on a turn that Gríma does not attack.
- Kindred Discovery can draw a card when Gríma attacks, because Gríma is a Human Advisor.
- Idol of Oblivion can draw a card after Coin of Mastery makes a Treasure token, and the deck holds Coin of Mastery.

The owner chose a record, and a replay after PR-54 reads the conditions again (D-727).

## Limits

- Tagger tags are community data (F-5). A tag can miss a finisher, or it can hold a card that wins no game.
- No real list carries a bracket, and no real group stands for bracket 4.
- The land classes and the condition rule read rules text. A card with rare text can take the wrong class.
- The replay reads the local card snapshot of 2026-09-04 and the local model `20260914T154223Z`. The deployed app read a snapshot of 2026-09-14 and the model `20260914T070904Z`.
- The web app sent the pool rule with the first message (`session-page.tsx`), so no card pool question went out. The thin-theme question of D-63 never had a turn.

## What the plan takes from M-17

- **PR-54 comes first** (D-723 to D-725). A theme word finds its row through an alias or a word-form rule. A theme that matches no card gets a question before the build.
- **PR-52 keeps its shape** (D-720). The swap has real targets in the owned pool of this session: Sunken Hollow, Marsh Flats, and Fabled Passage. The class counts of the TopDeck lists above are information for its rank.
- **PR-53 reads the curated count**, with the target and the floor of D-726.
- **F-140 stays a record**, and the replay after PR-54 reads it again (D-727).
