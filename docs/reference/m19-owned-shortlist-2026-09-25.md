# M-19: the owned-only shortlist of the thumbs down of 2026-09-24

Read 2026-09-25 (F-174, D-881, D-950, D-951). No model call ran, and no paid target ran.

## The input

- **The verdict.** Card thumbs down `SM8yuktY7wiZA6xI8R8a` of 2026-09-24, reason `wrong_power`, on Island. The reader wrote: "Too many basic Islands (14) for a bracket 4 commander deck."
- **The deck.** Deck `3QqYHgVamrFPV8nhtzkj` of the snapshot (D-635). Commander Hope Estheim, W/U, bracket 4, pool rule owned-only, theme "Lifelink and counters". It holds 33 lands: 14 Island, 10 Plains, and 9 nonbasic lands.
- **The collection.** The deployed collection has the id `h-4bc6ef3e…`, a SHA-256 of the file. No local export has that hash, so the deployed file is newer. The replay reads the ManaBox export of 2026-09-02. It matches 71 of the 77 owned counts of the deck, and all 11 land counts. The export of 2026-08-30 gives the same result. Both exports stay out of git (D-756).
- **The cards.** The local snapshot of 2026-09-04 and the local quality model. The deployed build read the snapshot of 2026-09-24.

## The method

A scratch test in package `generate` ran the deployed request of `agentsvc` through `candidates.Build`: the colors, the theme, the commander, owned-only, and bracket 4. It also ran the build with no caps. Then it built the pool as `FromListOwned` builds it, and ran the guards of the mana pass over the stored deck. The test stays in `.local/m19`, outside git.

## The result

| Step | What it did with the owned lands |
|---|---|
| The owned pool | 70 nonbasic lands, legal, inside W/U |
| The shortlist | 40 of them, the cap of the land role. It holds every owned land of classes 0 to 3, and 28 of the 30 lands of D-881 |
| The model | 11 lands: all 6 owned lands of classes 0 to 2, and Exotic Orchard, Irrigated Farmland, and Tranquil Cove |
| `swapBasics` | No swap. Its list reads classes 0 to 2 alone, and the deck holds each owned land of those classes |
| `fillFixing` | No fill. The deck holds 9 fixing lands, and the floor of two colors at bracket 4 is 9 (D-799) |
| `balanceBasics` | No trade. White has 25 sources for a need of 23, and blue has 29 for a need of 19. Both colors meet their need, so the capped ratios can not rise |

The classes come from `profile.LandClassOf`. Class 0 is an untapped dual, 1 a fetch, and 2 untapped on a condition. Class 3 makes mana on a condition, 4 is a tapped dual, and 5 is other. The six owned lands of classes 0 to 2 are City of Brass, Command Tower, Fabled Passage, Glacial Fortress, Port Town, and Prairie Stream.

**The count of D-881 was wide.** It counted each owned land whose Scryfall `produced_mana` holds W and U. That list holds 19 lands of class 3, such as Villainous Hideout, Secluded Courtyard, and Plaza of Heroes. Each one makes colored mana only for some spells, or for {1}. 10 lands of class 4 enter tapped, such as Path of Ancestry, Thriving Isle, and Crossroads Village. Baxter Building reads class 5. None of the 30 is an untapped dual land.

**So no step dropped a better owned land.** The collection holds no better W/U land than the deck holds. The split of the basics is the fault that the reader named: blue holds 10 sources over its need, and white holds 2.

## The options the owner read

| Option | Effect on this deck | Risk |
|---|---|---|
| A. Balance the basics past the need | 5 trades, 15 Plains and 9 Island. White 1.30, blue 1.26. The band score stays at 1.75 | The deck keeps 24 basic lands |
| B. The swap takes class 3 lands | All 8 lands of the fill list keep each guard | The profile counts Villainous Hideout as a whole W and U source |
| C. A floor of 11 at bracket 4 | The fill adds Opal Palace and Villainous Hideout for 2 Islands | The same risk as B, and it overrides the floors of D-799 |
| D. No fix | No change | The split stays |

The owner chose A, in this pull request (D-951).

## The fix over deck gate run 29

`make manapass-check` read the 25 decks of `docs/reference/pr8-deck-gate-run29.md`, before and after the fix. It calls no provider.

- 11 decks change, and 14 keep their basics.
- No deck gains an off-band feature.
- The worst color rises on 10 decks. Deck 2 ends on the same basics after one more step.
- The largest moves: deck 16 goes from W 12, B 11 to W 16, B 7. Deck 14 goes from R 11, G 11 to R 14, G 8. Deck 18, a precon upgrade of five colors, goes from W 2, U 2, B 2, R 2, G 3 to W 6, U 1, B 1, R 1, G 2.
