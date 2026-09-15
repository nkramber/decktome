# PR-54: theme words and the theme row (2026-09-14)

This note holds every count of PR-54 (F-141 to F-144, D-723 to D-725, D-728 to D-731). Each count reads the local card snapshot `20260904T210157` and calls no model. The "before" counts read `main` at `c7a4ff4`. The "after" counts read the branch `pr54-theme-words`.

## What changed

- A theme word finds its row by name, by an alias, or as the singular of a plural row name (D-724). A type row joins no plural index (D-731).
- Each base form of a word tries the same three things. The base forms are the singular and the stem of an "-ing" form, so "milling" finds mill.
- Two neighbor words join when their hyphenated form finds a row. So "extra turn" finds extra-turns, and "land destruction" finds land-destruction.
- Two rows are new. The superfriends row reads the card type Planeswalker (D-729). The land-destruction row reads the tags `removal-land` and `mass-land-denial`.
- The theme row asks for the theme again before the build when no word of the theme matches a card (D-725, D-728).

## The theme probe

The probe builds a Commander shortlist with any card allowed and no color limit. The count reads the cards of the pool with a theme signal.

| Theme | Before | After | Row |
|---|---|---|---|
| milling | 0 | 1,260 | mill |
| opponent milling cards | 0 | 1,260 | mill |
| blinking | 0 | 377 | blink |
| flickering | 0 | 377 | flicker |
| reanimating | 0 | 2,433 | reanimator |
| sacrificing | 27 | 1,413 | sacrifice |
| stealing | 0 | 679 | theft |
| gaining life | 0 | 2,645 | lifegain |
| wheeling | 0 | 333 | wheel |
| ramping | 0 | 3,644 | ramp |
| storming | 0 | 772 | storm |
| clones | 0 | 956 | copy |
| flyers | 0 | 5,118 | flying |
| counterspells | 0 | 11,294 | control |
| self-mill | 0 | 4,755 | graveyard |
| prison | 2 | 4,804 | stax |
| going wide | 1 | 3,361 | go-wide |
| superfriends | 0 | 645 | superfriends |
| token | 216 | 3,870 | tokens |
| artifact | 1,025 | 4,041 | artifacts |

Three probe words still match no card: "lifedrain", "draining", and "anime". The theme row asks about each of them. A singular type word such as "dragon" keeps its count, because it keeps the generic rule (D-731).

## The counted gate conversations

Before PR-54, four counted gate conversations named a theme that matched no card. With the new row and no alias, each of them would get a question that its script never answers.

| Conversation | Theme | Before | After |
|---|---|---|---|
| 17, 62, and 82 | superfriends | 0 | 645 |
| 71 | weenie | 0 | 1,020 |

After PR-54, the theme of each counted gate conversation matches cards. The source for the alias "weenie" is the MTG Wiki page "White Weenie", read 2026-09-14: "an aggro archetype that uses relatively cheap… efficient white creatures". Conversation 110, "a theme no card matches", is new, and its theme "anime" matches no card on purpose.

Five more themes of the gate now find a row. Question gate run 51 reads their conversations again.

| Conversations | Theme | Before | After |
|---|---|---|---|
| 19 | token | 110 | 1,680 |
| 29, 73, 74, 96, and 103 | artifact | 311 to 1,025 | 1,090 to 4,041 |
| 23 and 57 | land destruction | 1 | 287 |
| 34 and 61 | prison | 2 | 4,804 |
| 22 | extra turn | 46 | 46 |

## The Gríma replay

The replay builds the shortlist of session `z1hshyY6Npig1FN2NuV7` with the collection of the owner, as M-17 did. The request reads Commander, the theme "opponent milling cards", Gríma, Saruman's Footman, owned only, and bracket 5.

| Count | Before | After |
|---|---|---|
| Cards on theme | 0 | 823 |
| Owned cards on theme | 0 | 56 |
| Shortlist | 165 | 204 |
| Owned mill cards on the shortlist | 1 of 14 | 14 of 14 |
| Threat, synergy, and wincon cards | 0 | 39 |

Before PR-54, the one owned mill card on the shortlist was One Ring to Rule Them All, as a wipe with no theme signal. After PR-54, the shortlist also holds Laboratory Maniac as a wincon. The collection held 3,040 entries, 2,235 cards, and 6,030 copies on 2026-09-14. The export stays out of the repository.

## The gate prompts

The theme words of the gate prompts find the same signals before and after PR-54.

- A scratch build of each prompt theme gives the same shortlist on both commits, for all 43 prompts of the deck, bracket, and revise gates.
- `make deck-gate-dry` gives the same output on both commits, for all 25 prompts.

## The type rows (D-731, F-144)

The first build of the plural rule sent a singular type word, such as "zombie", to its type row. The table counts the Commander-legal cards that the generic rule matched and the type row did not match. The type text column counts the cards whose text names the type in a phrase, such as "Zombie creature token".

| Word | Cards lost | Type text | Other |
|---|---|---|---|
| spirit | 190 | 144 | 46 |
| zombie | 167 | 81 | 86 |
| angel | 129 | 30 | 99 |
| cat | 128 | 13 | 115 |
| human | 125 | 32 | 93 |
| dragon | 76 | 28 | 48 |
| goblin | 53 | 35 | 18 |
| knight | 49 | 38 | 11 |
| elf | 45 | 2 | 43 |

The type row gained no card for any word. The owner kept the generic rule for a singular type word (D-731). So "zombie" and "zombies" still read two different lists, and F-144 records the gap of the type rows.

## Question gate run 51

Run 51 ran on `cab4f8e` (D-730). It read PASS: 75 of 75 counted conversations used catalog questions only, and every expected slot held. It cost $0.1958 over 1,276 seconds and 682 calls.

The theme row asked in 8 conversations.

| Conversation | Theme that the classifier wrote | Right to ask |
|---|---|---|
| 110 | anime | yes, as the conversation intends |
| 108, a probe | animals | yes, because no card holds the word |
| 4 | fun and janky | no, because a jank word names a power |
| 16 and 63 | the strongest Modern deck | no, because the request names no theme |
| 68 | the strongest Modern deck possible | no |
| 37, a probe | competitive Modern deck | no |
| 72 | Good stuff | no, because the deck has no theme |

Against run 50, each of the 7 conversations beside conversation 110 asked one more question. Conversations whose theme did not change moved too, by one or two questions in both directions. So a move of one question in one run is noise.

The fix adds format names, jank words, and "stuff" to the stop words (F-145). The unit test fails on the old match for all 5 phrases and passes on the new. Question gate run 52 measures the fix (D-732).
