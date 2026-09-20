# The payoff shape of the commander, the Gríma evidence, 2026-09-20

This document holds the free measurement that OQ-83 to OQ-86 need. The owner asked whether the build reads the payoff shape of the commander. The review of the Gríma deck of 2026-09-14 raised the four questions. `docs/reference/deck-review-grima-2026-09-14.md` maps every suggestion of that review.

**Every number here is free.** No paid target ran. The sources are the card snapshot of 2026-09-04, the ManaBox export of the owner of 2026-09-02 (D-756), and the code of `main` at `5d1d7e0`.

## The method

A scratch Go test inside `go/cmd/deck-gate` made every count. A copy sits in `.local/oq84/`, out of git, because the repository is public (D-639). The test replays the request of M-17 and M-18: Commander, the theme "opponent milling cards", Gríma, Saruman's Footman, owned only, and bracket 5.

The replay reads the same inputs as M-18, and it reproduces its counts. So the two documents compare.

| Input | Value |
|---|---|
| Export rows | 2,956 |
| Rows that no card resolved | 1 |
| Collection entries | 2,848 |
| Cards, copies counted | 5,335 |
| Distinct oracle ids | 2,178 |
| Shortlist cards | 201 |

## The commander

The snapshot of 2026-09-04 holds this card.

| Field | Value |
|---|---|
| Name | Gríma, Saruman's Footman |
| Type | Legendary Creature — Human Advisor |
| Mana cost | `{2}{U}{B}` |
| Power and toughness | 1/4 |
| Color identity | blue and black |
| Keywords | none |

The Oracle text holds two lines:

- "Gríma can't be blocked."
- "Whenever Gríma deals combat damage to a player, that player exiles cards from the top of their library until they exile an instant or sorcery card. You may cast that card without paying its mana cost. Then that player puts the exiled cards that weren't cast this way on the bottom of their library in a random order."

Two facts of that text drive OQ-83 and OQ-84:

- The trigger fires one time for each combat damage event to a player. A second hit needs a second combat, a second attacker, or double strike. More power on Gríma adds nothing to the trigger.
- Gríma holds the strongest evasion in the game. A card that grants flying, menace, shadow, or the unblockable line adds nothing to it.

## OQ-83, what the build reads about the commander

**The build reads the name and the color identity of the commander, and nothing else.**

`candidates.Request.CommanderOracleIDs` has two uses in `go/internal/candidates/candidates.go`. Line 306 excludes the commander from the 99 candidates. Line 1189 excludes it from the pair offer. No score term, no role rule, and no cap reads the card.

The model prompt carries no more. `go/internal/generate/input.go` line 43 writes one sentence: "The commander is {name}. It is chosen, and it is not one of the cards you list." The shortlist omits the commander (D-302). Each shortlist line carries the name, the type line, the job, and the power marks. No line carries Oracle text.

So the model must recall the card from its training data. The app gives it no text of the commander at all. The color identity reaches the build through `Request.Colors`, which the caller derives from the card.

## OQ-84, the evasion the commander already holds

A coarse text match read every shortlist card for an evasion word beside a grant phrase. It found 10 cards of the 201. A hand read of the 10 gives this classification.

| Class | Count | Cards |
|---|---|---|
| The card grants evasion, and Gríma can take it | 3 | Escape Tunnel, "My Precious // Allure of Power", Waterbender Ascension |
| The card grants evasion, and Gríma can not take it | 1 | Turtle Lair |
| The card holds the keyword itself, and grants nothing | 6 | Dawnhand Eulogist, Ingenious Prodigy, "Massacre Girl, Known Killer", "The Rise of Sozin // Fire Lord Sozin", Voracious Fell Beast, Witch-king of Angmar |

**The answer: 3 cards of 201 grant evasion that Gríma already holds. No card of the shortlist has that grant as its only job.** Each of the three carries a second mode that the deck can use:

- Escape Tunnel searches a basic land. Its second mode gives a creature of power 2 or less an unblockable turn, and Gríma has power 1.
- "My Precious // Allure of Power" grants hexproof beside the unblockable line, and its Adventure half draws two cards.
- Waterbender Ascension draws a card for each four hits, and Gríma hits.

Turtle Lair can not target Gríma at all. Its grant names a Ninja or a Turtle, and Gríma is a Human Advisor. So the card is a dead mode for this commander, and a discount on evasion does not find it.

CAUTION: 6 of the 10 flagged cards carry the keyword on themselves. A discount that reads the keyword alone, and not the grant, punishes all six. Four of them are the threats and the removal of the shortlist.

## OQ-85, the effect classes of the shortlist

The same replay counted coarse effect classes over the 201 shortlist cards. Each count is a floor, because a text match misses a card that words the effect in another way.

| Effect class | Cards |
|---|---|
| Equipment | 12 |
| Combat damage trigger | 11 |
| Counterspell | 8 |
| Static pump on a carrier | 6 |
| Extra turn | 1 |
| Trigger copier | 0 |
| Extra combat | 0 |
| Double strike | 0 |

The review named two classes: "seven counterspells are fine, and eight pump equipment on a trigger that fires once are not". The shortlist holds 8 counterspells and 12 Equipment, and 6 of the Equipment give a static pump.

The owned pool is the limit here, not the cap. The pool holds no trigger copier, no extra combat, and no double strike card that reaches the shortlist. Those three classes are the ones that fit the payoff shape of Gríma best. F-157 records the thin owned pool.

## OQ-86, the power estimate

This question needs no new measurement. The review read the deck as bracket 3 by the rules and 2.5 by power, and the build aimed at bracket 5. `docs/reference/deck-review-grima-2026-09-14.md` holds the counts: 0 of 4 tutors, 3 of 6 fast mana, and 1 of 8 Game Changers against the floors of bracket 5.

The profile holds those floor counts today, and the bracket judge reads a bracket. No number that a reader sees separates the rules bracket from the power of the deck.

## What the evidence supports

Each of the four questions rests on a true premise. The owner read this document and answered all four on 2026-09-20.

| Question | The premise | The answer |
|---|---|---|
| OQ-83 | The build reads no ability of the commander | D-771. The prompt carries the Oracle text of the commander. F-159 records the gap. |
| OQ-84 | The build discounts no redundant evasion grant | D-772. The build adds no discount. Three cards of 201 carry such a grant, and each carries a second mode. |
| OQ-85 | No cap counts an effect class | D-773. A cap waits for the shape work, because the class weights need the shape. |
| OQ-86 | No reader-facing number shows the power apart from the rules | D-774. A deck shows the floor counts of its bracket, in its own pull request. |
