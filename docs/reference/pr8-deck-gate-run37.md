# PR-8 deck gate, a rejudge of pr8-deck-gate-run35

Run date: 2026-09-29. Card snapshot: 2026-09-04. Decks read from `pr8-deck-gate-run35.md`.

Verdict: PASS. every deck of the source passed its build bars, 0 summaries stated a false rule of the game, and 0 judge calls failed. The build rows are the rows of the source run, and no deck was built again (D-789).

- Calls: 55. Cost: $0.3929. Time: 181 seconds.

## Run

- Suite `decks`, run `pr8-deck-gate-run37`, on 2026-09-29, commit `3cdb803`.
- Roles: generate on `gpt-5.6-terra` (openai, effort medium), judge on `claude-sonnet-5-5` (anthropic, effort medium), repair on `gpt-5.6-terra` (openai, effort medium).
- Versions: card snapshot 2026-09-04, generate prompt version 16, plan_rubric prompt version 5, summary_judge prompt version 2, precons `5.3.0+20260923`, quality_model `20260923T202806Z`, rejudge_of `pr8-deck-gate-run35`.
- Calls: 55. Cost: $0.3929. Time: 181 seconds.
- Note: a rejudge of pr8-deck-gate-run35: the decks and the build rows are copied, and the judge rows are new (D-789)

## Decks

### 1. lifegain Commander, any card

**Summary:** This lifegain deck uses steady development and card advantage to make Karlov of the Ghost Council the central battlefield threat, with protection to preserve its key pieces and broad removal to clear resistance. Its creature suite applies pressure through lifegain payoffs while Angel of Destiny, Approach of the Second Sun, and Felidar Sovereign provide distinct ways to close a game. It gives up some speed for a resilient midrange plan that is strongest once its engines and mana are established.

The quality model grades this deck below the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The mana base serves the colors evenly, and that raises the grade.

- PLAN plan_coherent=partly: Most of the white creatures and payoffs pull toward lifegain, but the ramp slots are filled with weak artifacts like Hot Dog Cart and Colossal Plow, and Karlov is only one of many threats instead of the deck's clear center.
- PLAN theme_fit=yes: It is a white-black lifegain deck led by Karlov of the Ghost Council with Archangel of Thune, Sanguine Bond, Heliod, Sun-Crowned, and many other lifegain payoffs, at a plausible bracket 3 level.
- PLAN useful_as_built=partly: The mana base of 36 lands with many dual lands is sound and there are wipes, removal and several finishers, but the ramp is made up of filler artifacts and the deck has many expensive top-end cards at six and seven mana, so it will stumble in some games.
- PLAN summary_honest=partly: The summary names Angel of Destiny, Approach of the Second Sun and Felidar Sovereign as win conditions and all three are in the list, but it says Karlov is the central battlefield threat and promises protection and card advantage when the list has only scattered small effects and a fair number of weak artifacts.

### 2. aristocrats Commander, owned first

**Summary:** Denethor turns a steady supply of creatures into sacrifice pressure, draining opponents while building a fresh Soldier force after creatures die. The deck develops through inexpensive mana and card engines, protects its important pieces, and uses broad spot interaction plus selective board resets to keep opponents from getting ahead. It closes through dedicated finishers or by grinding repeated sacrifice turns into an advantage, giving up raw speed for a board-dependent plan that wants creatures and time to establish itself.

The quality model grades this deck at the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. The mana base serves one color far better than another, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade.

- PLAN plan_coherent=partly: Many pieces are protection, removal, wipes and ramp with only a thin layer of sacrifice outlets and death payoffs, so the aristocrats plan is diluted by a generic good-stuff pile.
- PLAN theme_fit=partly: It is a Denethor deck in the right colors, but the request asked for aristocrats and the list has few sacrifice outlets, token makers, or drain effects, and the 'library first' element is only lightly reflected.
- PLAN useful_as_built=partly: It has plenty of ramp, draw, and interaction and a functional land count, but 25 Plains and Swamps skew heavily toward white for black-heavy spells, and there are few real win conditions, so it plays as a pile of good cards more than a closing deck.
- PLAN summary_honest=partly: The summary claims sacrifice pressure and a steady supply of creatures, but the list holds few fodder makers and few drain payoffs, and it never mentions that the mana base is heavily skewed toward Plains.

### 3. artifacts Commander, bracket 4

**Summary:** Urza leads a fast blue artifact deck that uses early mana, artifact tutors, and sustained card advantage to establish a dense board quickly. The deck protects its position with extensive stack interaction, answers troublesome permanents through artifact-based removal, and closes through artifact creature pressure or alternate endgames from Mechanized Production, Mirrodin Besieged, and Thassa's Oracle. It gives up broad multicolor answers and relies heavily on artifacts remaining central to its plan.

The quality model grades this deck typical against the top lists of the format: the deck makes more mana on turn four than the norm, and that raises the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The mana base serves the colors evenly, and that raises the grade.

- PLAN plan_coherent=yes: Nearly every card is a blue or colorless artifact, artifact tutor, artifact ramp piece, counterspell, or artifact-based finisher, so the list works toward one plan of building a big artifact board with Urza.
- PLAN theme_fit=yes: It is a mono-blue artifact deck led by Urza, Lord High Artificer, and it has fast mana, Force of Will, Fierce Guardianship, and tutors that suit bracket 4.
- PLAN useful_as_built=partly: The deck has plentiful free mana and interaction, but 34 lands including 25 Islands plus many rocks is heavy on mana, several rainbow lands add nothing over Islands in a mono-blue deck and some hurt with pain, and a few cards like Myr Enforcer, Sojourner's Companion, and Braided Net are weak, so it plays but is unpolished.
- PLAN summary_honest=yes: The summary claims Mechanized Production, Mirrodin Besieged, and Thassa's Oracle as alternate endgames, and all three are in the list, along with the mana artifacts, tutors like Transmute Artifact and Whir of Invention, and many counterspells it describes.

### 4. dinosaur tribal, bracket 2

**Summary:** This is a straightforward Dinosaur combat deck: build mana, develop Dinosaur support, then bring Gishath, Sun's Avatar into combat to turn its trigger into a large board of Dinosaurs. The main finish is an overwhelming attack backed by large Dinosaur threats and finishers, with removal and a few reset buttons available when the board gets crowded. It favors committing creatures and pressing combat over a highly reactive game plan, so careful sequencing and protecting Gishath matter.

The quality model grades this deck below the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The land count sits above the norm of the format, and that lowers the grade.

- PLAN plan_coherent=yes: Nearly every card is a Dinosaur, Dinosaur ramp or tutor, or protection for Gishath, so the whole list supports one plan of building mana and attacking with Dinosaurs.
- PLAN theme_fit=yes: This is a Dinosaur Commander deck led by a Dinosaur commander with plenty of tribal support, and it has no infinite combos or extreme cards, which suits bracket 2.
- PLAN useful_as_built=partly: The deck is playable and has lots of ramp and Dinosaurs, but a new player must handle a 35-land, three-color base with painlands, fetches, and Mana Confluence and City of Brass, and the many expensive three-color cards make it clunky and awkward to sequence.
- PLAN summary_honest=partly: The summary's claim of a Dinosaur combat deck with Gishath's trigger is borne out by the many Dinosaurs, but it says nothing of the 35-land base with heavy painland and fetch usage, and its 'reset buttons' amount to Vandalblast, Forerunner of the Empire and Raging Swordtooth, which are not real board wipes.

### 5. blink Commander, owned first

**Summary:** Gilraen anchors a creature-focused blink plan, repeatedly setting up favorable returns while the deck develops mana, cards, equipment, and a protected board. It wins by turning its growing creature force sideways, with Angel of Serenity chief among its closing threats, while removal and wipes clear away resistance. The deck gives up multicolor flexibility for a very consistent Plains-based mana base and relies on creatures staying relevant on the battlefield.

The quality model grades this deck typical against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The mana base serves the colors evenly, and that raises the grade.

This deck plan aims for 2 or more finishers, the cards that win the game, and the deck holds 1.

- PLAN plan_coherent=partly: The deck pairs a creature blink plan with heavy artifact, equipment, removal and wipe packages, but it holds only about 20 creatures, several of them blink-irrelevant, and wipes such as Austere Command and Vanquish the Horde work against its own board.
- PLAN theme_fit=partly: Gilraen is a white commander who fits a mono-white blink deck, but there are few blink enablers beyond Flickerwisp, Ennis and Together Forever, and the deck reads more as white equipment and artifact good-stuff than as a blink deck, so the request is only partly met.
- PLAN useful_as_built=partly: The deck is castable, with 36 Plains and about ten mana rocks, but that is too much mana for the list, which has only about 20 creatures and one real finisher, so it will flood and lack a reliable way to win.
- PLAN summary_honest=partly: The summary says Angel of Serenity is chief among the closing threats, and it is the one true finisher in the list, but the summary implies several closers and a working blink engine when only Flickerwisp and a few enters-the-battlefield creatures actually blink.

### 6. Modern tempo, tournament

**Summary:** This blue-red tempo deck establishes pressure with Ragavan, Nimble Pilferer, Ledger Shredder, and Faerie Mastermind, then protects that pressure with efficient interaction and removal. Preordain and Mishra's Bauble keep the deck moving, while Temporal Mastery and Temporal Trespass reinforce its spell-focused plan. It gives up broader late-game threats and sweeping answers in exchange for a fast, focused game built around trading efficiently and staying ahead.

The quality model grades this deck typical against the top lists of the format: the cards pair in ways the top lists do not, and that lowers the grade. The deck runs few spells as one copy, and that raises the grade. The deck runs its spells as playsets, and that lowers the grade.

- PLAN plan_coherent=partly: The cheap creatures, burn and counterspells form a coherent tempo core, but Temporal Mastery and Temporal Trespass are seven- and eleven-mana spells that pull against a deck built to win fast on one to two mana.
- PLAN theme_fit=partly: The list is blue-red tempo with cheap threats, burn and counterspells for Modern, but it carries no Delver of Secrets or Delver-style flip payoff, and the expensive Temporal spells are off-theme for the request.
- PLAN useful_as_built=partly: The 26 lands, cheap threats and interaction play smoothly, but only 12 creatures and a few uncastable-in-practice top-end spells leave the deck thin on threats and holding dead cards.
- PLAN summary_honest=partly: The summary correctly lists Ragavan, Ledger Shredder, Faerie Mastermind and Preordain, but it calls Temporal Mastery and Temporal Trespass part of a spell-focused plan when they are expensive cards with little support, and it never mentions that there are no Delver of Secrets in the list.

### 7. Modern burn, casual

**Summary:** This mono-red burn deck applies early pressure with efficient removal and spell-focused synergy creatures, then keeps the cards coming to sustain its assault. Its threats give it a board-based route to victory when direct burn alone is not enough, with Hazoret the Fervent chief among them as a closing threat. The deck gives up flexibility against specialized strategies in exchange for a focused, consistent red mana base and a direct aggressive game plan.

The quality model grades this deck at the precon baseline against the top lists of the format: the cards pair in ways the top lists do not, and that lowers the grade. The deck runs few spells as one copy, and that raises the grade. Many of the cards appear in no top list, and that lowers the grade.

- PLAN plan_coherent=partly: The burn spells, Eidolon of the Great Revel and Thermo-Alchemist point toward a spell-based burn plan, but a pile of four-drop creatures like Barret Wallace, Champion of the Path and Hazoret pulls it toward midrange, and Hazoret wants an empty hand that Browbeat and Risk Factor do not reliably deliver.
- PLAN theme_fit=partly: The deck is mono-red and legal for Modern, but a burn deck should be full of cheap burn spells and this one has only six, with the rest of the deck made up of creatures, so it reads more like red midrange than the burn deck that was requested.
- PLAN useful_as_built=no: With 24 Mountains, only six real burn spells and a top-heavy cluster of 12 four-drops, the deck floods often, has too little reach to close games, and lacks the cheap interaction to survive until its expensive creatures matter.
- PLAN summary_honest=partly: The summary says Hazoret the Fervent is a closing threat and that the deck has efficient removal, which the list does confirm with four Hazoret and Lightning Bolts, but it calls the removal a large part of the deck when only six cheap burn spells are present and it never admits the deck is missing most of the burn.

### 8. Modern lifegain, FNM

**Summary:** This white-black lifegain deck establishes its early board with Guide of Souls and Ocelot Pride, then turns that foundation into pressure through Enduring Tenacity, Twinblade Paladin, and its larger threats. Solitude and Murderous Rider keep opposing threats from taking over while Lembas and Enduring Innocence help sustain the hand. It wins by building a resilient lifegain-based board and closing with its powerful creatures; the tradeoff is that several of its strongest threats ask the deck to reach the middle and later stages of the game.

The quality model grades this deck below the precon baseline against the top lists of the format: the cards pair in ways the top lists do not, and that lowers the grade. The deck runs few spells as one copy, and that raises the grade. The deck runs its spells as playsets, and that lowers the grade.

- PLAN plan_coherent=partly: Guide of Souls, Ocelot Pride, Enduring Tenacity and Twinblade Paladin do work together as a lifegain plan, but Lembas is a weak draw piece, Altar of the Pantheon is off-plan, and the deck has only 2 Murderous Rider and 2 Archangel with almost no cheap interaction.
- PLAN theme_fit=yes: The deck is white-black, leans on lifegain payoffs like Guide of Souls, Ocelot Pride, Enduring Tenacity, Sheoldred and Archangel of Thune, and is set at a reasonable FNM power level for Modern.
- PLAN useful_as_built=partly: The 23 lands with a full set of dual lands give a solid mana base and the one-drops and Sheoldred give a workable curve, but four Lembas and two Altar of the Pantheon are low-impact slots and the top end of four Solitude plus Archangels is somewhat clunky.
- PLAN summary_honest=partly: The summary says Guide of Souls and Ocelot Pride open the board and Solitude and Murderous Rider handle threats, which the list bears out, but it leaves out Sheoldred, the Apocalypse (a main threat) and its claim of resilience overstates a deck with almost no cheap interaction.

### 9. Standard midrange, FNM

**Summary:** This black-green midrange deck builds its mana around Overgrown Tomb and Wastewood Verge, uses Llanowar Elves to reach its heavier creatures, and keeps cards flowing through Darkstar Augur and Phyrexian Arena. Goldvein Hydra and Chomping Changeling apply steady pressure, while Vein Ripper, Massacre Wurm, and Vaultborn Tyrant provide heavier finishes. Bitter Triumph and Maelstrom Pulse clear the way, and Snakeskin Veil plus Not Dead After All help the creature plan survive resistance. It gives up a broad range of alternate angles in favor of committing to creatures, protection, and direct answers.

The quality model grades this deck below the precon baseline against the top lists of the format: the cards pair in ways the top lists do not, and that lowers the grade. Few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade.

- PLAN plan_coherent=partly: The deck pairs creatures, removal and protection in black-green, but 4 Not Dead After All and 4 Snakeskin Veil sit beside only 22 creature slots of mixed weight, and the top end of six-drops and a seven-drop with only 2 Llanowar Elves pulls against the small early game.
- PLAN theme_fit=yes: It is a black-green creature midrange deck with removal and big finishers in Standard, fitting the request at a casual FNM level.
- PLAN useful_as_built=partly: With 24 lands and a good black-green fixing base it should function, but the heavy BBB six-drops, few early plays besides Changeling and Augur, and situational protection spells make it clunky and inconsistent.
- PLAN summary_honest=partly: The summary says Llanowar Elves help reach heavier creatures, but the list has only 2 Elves, so the ramp claim is thin, though Overgrown Tomb, Wastewood Verge, Darkstar Augur and Phyrexian Arena all appear as stated.

### 10. Standard aggro, tournament

**Summary:** This red-white aggro deck commits threats early, backs its attacks with removal and flexible interaction, and uses its draw cards to keep pressure on after the first exchange. Warleader's Call is chief among the cards tying the creature plan together. It aims to win by maintaining a fast board presence and clearing the way rather than by relying on a slow, high-end finish.

The quality model grades this deck below the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. The cards pair in ways the top lists do not, and that lowers the grade. The curve sits high for the format, and that lowers the grade.

- PLAN plan_coherent=partly: The red-white colors and Warleader's Call point toward a creature-based aggro plan, but only a dozen or so creatures are spread among 4-drops, a 6-drop, and non-creature enchantments, so the cards pull in slightly different directions.
- PLAN theme_fit=yes: The deck is a red-white creature and removal deck built on Boros lands for Standard, which matches the request for red-white aggro.
- PLAN useful_as_built=partly: The 22-land Boros mana base is smooth and there is removal and card flow, but the threat count is low, the top end includes four six-mana Beasties, and only four two-drop creatures make the aggro plan sluggish.
- PLAN summary_honest=partly: The summary calls Warleader's Call the card tying the creature plan together, and four copies are in the list, but the claim of committing threats early is weakened by Bedhead Beastie at six mana and by having few cheap creatures, and it does not mention the 22 lands and high curve.

### 11. Commander with a locked card

**Summary:** This is an Orzhov sacrifice deck built to turn a broad creature board into steady pressure, with Karlov of the Ghost Council as a growing centerpiece. Sacrifice outlets, death-focused payoffs, and creature-based value keep the board working through exchanges, while high-impact threats and dedicated closing cards end drawn-out games. It gives up some speed to retain a deep creature package and several expensive top-end plays.

The quality model grades this deck at the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. Few lands enter tapped, and that lowers the grade.

- PLAN plan_coherent=yes: Nearly every card is a sacrifice outlet, a death-trigger payoff, a token or creature engine, or a sweeper that feeds the aristocrats plan, and the commander sits inside that plan.
- PLAN theme_fit=yes: It is a Commander deck led by Karlov of the Ghost Council in white and black, and its cards are dedicated to sacrificing creatures for value.
- PLAN useful_as_built=yes: With 24 basics on top of a healthy set of dual lands, Sol Ring, two altars, and cheap outlets, the deck can be played as it stands, although the manabase is more colorless-light than it needs to be.
- PLAN summary_honest=yes: The summary promises sacrifice outlets and death payoffs, and the list carries Viscera Seer, Ashnod's Altar, Phyrexian Altar, Yawgmoth, Zulaport Cutthroat and Grave Pact, so the claim checks out.

### 12. Commander on a budget

**Summary:** Adeline, Resplendent Cathar drives a wide white token strategy, turning attacks into an expanding Human force and using the deck’s token-focused support to make that board matter. The deck develops with plentiful mana and card flow, keeps opposing boards in check with targeted answers and resets, then closes through combat pressure or dedicated finishing cards such as Halo Fountain, Luck Bobblehead, and Sword of Body and Mind. It favors steady board development over explosive early-game acceleration, so it is most effective when it can keep creatures on the table and attack repeatedly.

The quality model grades this deck below the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. The deck makes less mana on turn four than the norm, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade.

- PLAN plan_coherent=yes: Nearly every card is white and supports going wide with tokens and attacking, from Hero of Bladehold, Warren Warleader and Felidar Retreat to Intangible Virtue and the wraths that reset the board.
- PLAN theme_fit=yes: It is a mono-white token deck led by Adeline in Commander with budget-friendly cards, and the list has no expensive staples, which fits the bracket 2 and under 100 dollars request.
- PLAN useful_as_built=partly: With 37 lands and many mana rocks it will always cast its spells, but the ramp is weak (Coin of Mastery, Currency Converter, Druidic Satchel) and many cards are filler, so it plays smoothly but a bit clunkily.
- PLAN summary_honest=yes: The summary names Halo Fountain, Luck Bobblehead and Sword of Body and Mind as finishers, and all three are in the list as wincon cards, and its admission of no explosive early acceleration fits the list's mostly 3+ mana ramp.

### 13. owned first, and the commander is not owned

**Summary:** Karlov of the Ghost Council is the focal point of a white-black lifegain deck that develops mana, gains cards through a mix of creatures and equipment, and turns steady life gains into a dangerous commander. Protective equipment and instant-speed safeguards help Karlov stay active while targeted answers clear the way. The deck closes through Karlov's accumulated pressure, evasive Angel threats, and finishers such as Frodo, Sauron's Bane, Lyra Dawnbringer, and Grave Venerations; in exchange, it leans heavily on creatures and artifacts remaining in play to sustain its momentum.

The quality model grades this deck at the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The land count sits above the norm of the format, and that lowers the grade.

- PLAN plan_coherent=partly: The white-black lifegain shell with equipment, removal and Angels mostly points one way, but the many equipment and artifact-ramp pieces, three board wipes that undo the creature plan, and only a handful of real lifegain payoffs leave the pieces pulling apart.
- PLAN theme_fit=partly: Karlov of the Ghost Council leads a white-black deck as asked, but lifegain is thin: Angel of Vitality, Lyra and a few incidental sources carry it, while much of the list is generic ramp, equipment and removal.
- PLAN useful_as_built=partly: The deck has plenty of ramp, draw, cheap removal and protection, so it functions, but 25 basics plus 12 nonbasics and about nine mana rocks make it flood-prone, and it has too few creatures and lifegain engines to convert that mana into a reliable win.
- PLAN summary_honest=partly: The summary names Frodo, Lyra Dawnbringer and Grave Venerations as finishers and all three are in the list, but it presents a lifegain deck built around steady gains when the list has few lifegain sources and reads more as a generic equipment and removal pile.

### 14. the user delegates the commander

**Summary:** Atarka, World Render leads a Gruul Dragon deck that ramps into a dense flight of threatening Dragons, then turns attacks into decisive double-strike combat steps. Tribal cost support helps deploy the expensive core, while draw, protection, removal, and sweepers keep the pressure sustained. The primary wins come from overwhelming combat with Atarka and the deck’s chief Dragon finishers, with the tradeoff that the deck remains creature-centered and can need its mana development to line up before its heaviest threats take over.

The quality model grades this deck below the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The land count sits above the norm of the format, and that lowers the grade.

- PLAN plan_coherent=yes: Ramp artifacts and lands feed a heavy red Dragon suite led by Atarka, with Dragon-cost reducers, protection equipment and burn all supporting one Dragon-pile plan.
- PLAN theme_fit=yes: It is a Dragon-tribal Commander deck led by a Dragon commander, Atarka, in Gruul colors, packed with Dragons and Dragon support at a plausible bracket 3 power level.
- PLAN useful_as_built=partly: The deck has plenty of ramp and Dragons, but the 12 Mountains plus many painlands and taplands for a heavy RRR-costed red suite with light green is workable, while the large land count and many mana rocks leave it flood-prone with few real draw engines.
- PLAN summary_honest=partly: The claim of ramping into a dense flight of Dragons is true, but 'sweepers' overstates a list whose only wipe-like effects are Breath Weapon and Draconic Intervention, and the deck is far more red than the Gruul framing implies.

### 15. delegated commander, owned first

**Summary:** A black-white lifegain deck built to turn incremental healing into an overwhelming board of creatures, resilient equipment carriers, and powerful flying finishers. It combines efficient removal, protective tools, and several board resets to maintain control while its life-total advantages become decisive pressure.

The quality model grades this deck at the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The land count sits above the norm of the format, and that lowers the grade.

- PLAN plan_coherent=partly: Lifegain payoffs, angels, removal and wipes point roughly one way, but the many equipment pieces, four board wipes in a creature deck, and a few stray cards like Canyon Crawler and Rooftop Percher pull in different directions.
- PLAN theme_fit=partly: It is a black-white lifegain Commander deck with a chosen commander, but Astarion at six mana is a clumsy pick and only a handful of cards actually gain life, and nothing shows it was built from the person's library first.
- PLAN useful_as_built=partly: It has ramp, removal, and flying finishers like Lyra and Angel of Invention, but the 32 basics split 18 Plains to 14 Swamp with WW-heavy costs, many low-impact artifacts, and a light lifegain count make it play unevenly.
- PLAN summary_honest=partly: The summary claims efficient removal and several board resets, which the list carries (Swords, Infernal Grasp, Austere Command, Fumigate, Vanquish the Horde), but it never mentions the 32 basics with a Plains-heavy split against black cards like Canyon Crawler.

### 16. a tight budget, owned first

**Summary:** This is a patient Orzhov aristocrats deck that develops creatures and support pieces, then uses Denethor’s sacrifice ability to turn creature deaths into life swings while building an end-step Soldier force. The finishing cards give the deck ways to close once the board has been established, while the removal, protection, and board resets help it survive longer games. It gives up explosive speed for a library-first mana base and an incremental, board-centered plan.

The quality model grades this deck at the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The deck makes less mana on turn four than the norm, and that lowers the grade.

- PLAN plan_coherent=partly: Many pieces are sacrifice and death-trigger creatures, but a large share of the list is generic goodstuff such as ten mana rocks, three wraths, and a pile of removal, and the wraths work against a creature-based aristocrats board.
- PLAN theme_fit=partly: The deck is Orzhov with a creature-and-sacrifice flavor at a low power level, but the request asks for a library-first deck and the list is mostly generic staples with few real sacrifice payoffs and outlets.
- PLAN useful_as_built=partly: It has 27 basics plus utility lands and plenty of mana rocks, so it will cast its spells, but the curve is shallow on real sacrifice outlets and payoffs and the wraths clash with the creature plan, so it plays as a pile of good cards rather than a working aristocrats deck.
- PLAN summary_honest=partly: The summary says Denethor turns creature deaths into life swings and builds an end-step Soldier force, but the list carries no clear engine for that; Denethor is just a commander on the list and the Soldier claim rests on a few incidental cards, so the plan it describes is thinly supported.

### 17. upgrade a precon, owned first

**Summary:** Zada Goblin Storm builds a wide Goblin board and turns creature-targeted spells into broad bursts of cards, pressure, and momentum. The deck wins by converting that developed board into a decisive attack, with Great Train Heist and Collective Inferno chief among its closing cards, while Goblin-linked removal clears resistance. It gives up some consistency to retain the precon's varied Goblin package and remains most effective when it has both creatures and a stocked hand.

The quality model grades this deck typical against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The deck makes less mana on turn four than the norm, and that lowers the grade.

- PLAN plan_coherent=partly: Most of the list is Goblin tokens and Goblin creatures that feed a wide board, but many one-mana cantrips and rituals, plus Blasphemous Act and Vandalblast, pull toward a spell-storm and board-wipe plan that fights the go-wide creature plan.
- PLAN theme_fit=yes: It is a mono-red Zada, Hedron Grinder Goblin deck with a full Goblin package and Sol Ring, Skullclamp, and other staples layered on, which is a reasonable bracket 3 upgrade of the Goblin Storm precon.
- PLAN useful_as_built=partly: With 27 Mountains plus 10 utility lands the mana is stable, but the many low-impact cantrips and situational cards like Frontline Heroism and Haze of Rage dilute the threats, and the deck makes little mana by turn four.
- PLAN summary_honest=partly: The summary says Great Train Heist and Collective Inferno close the game and that the deck has a wide Goblin board, which the list does carry, but it never mentions the heavy filler of cheap cantrips or the Blasphemous Act and Vandalblast wipes that undercut that board.

### 18. upgrade a precon, any card

**Summary:** This deck keeps Turtle Power’s broad Mutant, Ninja, and Turtle creature core while giving its mana base a cleaner route to all five colors. Build a board of themed creatures, send them into combat, and let Heroes in a Half Shell turn successful hits into larger attackers and more cards. Raphael, the Muscle, Dimension X Pizzasaur, and Everything Pizza are chief among the cards that can close a game. The deck gives up some of the precon’s slower lands to improve early development, while retaining its varied character-driven creature package and big finishing turns.

The quality model grades this deck below the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The land count sits above the norm of the format, and that lowers the grade.

- PLAN plan_coherent=partly: The creature-and-combat plan around Heroes in a Half Shell is clear, but the list also includes a wrath package (Blasphemous Act, Vanquish the Horde, Wave Goodbye) and several vague synergy pieces that pull against a board of attackers.
- PLAN theme_fit=yes: It is a five-color Heroes in a Half Shell Commander deck full of Mutant Ninja Turtle characters, which matches an upgrade of the Turtle Power precon.
- PLAN useful_as_built=partly: The mana base is strong with many dual lands, Sol Ring, Arcane Signet and Chromatic Lantern, but it runs about 40 lands and has several low-impact cards like Coin of Mastery, Arcade Cabinet and Level Up, so it plays smoothly yet sometimes flat.
- PLAN summary_honest=partly: The summary says Raphael, the Muscle, Dimension X Pizzasaur and Everything Pizza can close games, and all three are in the list, but Everything Pizza is a 2-mana Food artifact that is a weak finisher and the summary does not mention the sweepers or the 40 lands.

### 19. the Hobbit family, two colours

**Summary:** Thranduil leads an Elf-centered Sultai deck that develops its mana, fills the battlefield with Elf creatures, and turns legendary Elf arrivals into fresh cards while using Elf cards in the graveyard as a resource. The deck controls key opposing pieces with a broad mix of removal, disruption, and sweepers, then closes through its large marked finishers, chief among them the two Witch-kings. It gives up some raw speed for a creature-heavy, theme-driven plan that needs its mana and Elf presence to stay established.

The quality model grades this deck at the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. Many of the cards appear in no top list, and that lowers the grade.

- PLAN plan_coherent=partly: The green-blue Elf core with ramp and draw hangs together, but the black half of the deck is a pile of removal, sweepers and Orcs, and Languish, Gnashing of Teeth and Raise the Palisade fight the wide Elf board.
- PLAN theme_fit=yes: The deck is led by Thranduil, the Elvenking in his Sultai colors, uses only Hobbit-set cards, and leans on Elf creatures, which matches the request.
- PLAN useful_as_built=yes: With 35 lands, several Elf mana creatures and rocks, plenty of card draw and a cheap curve, the deck can be played as it stands, and it has enough removal and big finishers to win.
- PLAN summary_honest=partly: The summary claims Elf cards in the graveyard as a resource, but the list has very few graveyard payoffs beyond Haunt of the Dead Marshes, and Witch-king of Angmar and Witch-king, Bringer of Ruin are not Elves.

### 20. the Hobbit family, a delegated commander

**Summary:** A mono-red Hobbit Dragon deck centered on building Treasures, equipping resilient creatures, and leveraging combat to create overwhelming pressure. Smaug turns accumulated wealth into direct damage, while Dragons, giants, and a wide supporting cast provide multiple routes to a decisive finish. Removal and sweepers keep opposing boards manageable while legendary artifacts reinforce the deck’s adventurous Middle-earth character.

The quality model grades this deck at the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The mana base serves the colors evenly, and that raises the grade.

- PLAN plan_coherent=yes: Every card is mono-red and Hobbit-set, and the Treasure makers, equipment, Dragons, and sweepers all push toward Smaug turning Treasures into damage and pressure.
- PLAN theme_fit=yes: It is a mono-red Dragon-led Commander deck built entirely from the Hobbit sets with Smaug the Magnificent as the chosen commander, matching the request.
- PLAN useful_as_built=yes: With 36 lands, cheap ramp, card draw, removal, sweepers and Dragon finishers, the deck can be played as it stands, though a few top-end Dragons are expensive and the mono-red mana base is trivially consistent.
- PLAN summary_honest=partly: The summary claims Smaug turns accumulated wealth into direct damage and that Treasures are built, which the Treasure makers and Dragon's Desire support, but 'a wide supporting cast' and 'equipping resilient creatures' overstate a list that is many small vanilla-ish creatures and 20-plus equipment and artifacts of mixed value.

### 21. the Hobbit family, mana from outside

**Summary:** This Rakdos Smaug deck develops its mana early, deploys Smaug the Impenetrable and other large threats, then presses combat while using focused removal and broad resets to clear resistance. Smaug is chief among the finishers, with Dragons and other legendary threats providing several ways to end a stalled game. It gives up some flexibility for a committed battlefield plan and depends on its mana development to bring its expensive threats online quickly.

The quality model grades this deck at the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The deck makes less mana on turn four than the norm, and that lowers the grade.

- PLAN plan_coherent=partly: The deck aims at a Rakdos Dragon and legendary-threat plan, but about a dozen equipment pieces, a few ramp cards that do little, and loose draw cards leave the cards without one clear route to winning.
- PLAN theme_fit=yes: It is a Rakdos Commander deck led by Smaug the Impenetrable, built almost entirely from Hobbit-set cards, with lands and rocks from outside the sets as the request allowed.
- PLAN useful_as_built=partly: With 39 lands including many fetches, the mana is stable and the deck has removal, wipes and draw, but the quality note says it makes less mana on turn four than typical, and the equipment-heavy filler leaves it short of creatures and reliable ways to win.
- PLAN summary_honest=partly: The summary calls Smaug the Impenetrable a finisher and lists Dragons among the ways to win, and the list does hold Smaug Wicked Worm, Smaug the Magnificent and Desert Were-Worm, but it never mentions the pile of equipment or the weak early mana that its own note about mana development leaves unexplained.

### 22. a set family and a card from outside it

**Summary:** This is a Sultai Elf deck centered on Thranduil, the Elvenking, using legendary Elves and a full Elf creature base to develop pressure while keeping mana and cards flowing. Targeted answers, protective tools, and sweepers give it room to stabilize, then Troll of Khazad-dûm and the two Witch-kings provide its strongest closing threats. It gives up some raw speed for a creature-centric plan that benefits from a developed board and a stocked graveyard.

The quality model grades this deck at the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. Many of the cards appear in no top list, and that lowers the grade.

- PLAN plan_coherent=partly: The deck has a clear Elf creature core around Thranduil, but the Elf count is modest, the wipes (Languish, Gnashing of Teeth, Raise the Palisade) work against a creature-centric plan, and much of the artifact and equipment package does little for it.
- PLAN theme_fit=yes: It is a Sultai Thranduil, the Elvenking commander deck drawn almost entirely from the Hobbit sets, and it keeps the requested Sol Ring.
- PLAN useful_as_built=partly: With roughly 38 lands plus Mox Amber, Sol Ring, and Signet, the mana is plentiful and colors are covered, but the game plan is diffuse and many cards are low-impact, so it plays as a loose pile.
- PLAN summary_honest=partly: The summary says a full Elf creature base, but the list has only about 15 Elves next to many non-Elf artifacts, spells, and removal, so the base is not full.

### 23. two set families at once

**Summary:** Kíli the Resourceful leads a mono-white Dwarf and Equipment deck that develops through artifact mana, cheap gear, and a steady flow of creatures. Establish an enduring story early, use Kíli to turn each turn's first equip into efficient pressure, and keep cards flowing as Dwarves and Equipment enter. The deck wins by building one or more well-equipped attackers, with Angel of the Ruins, Helm of the Host, and Sunscorch Regent serving as major closing threats. It gives up broad color access and relies on its white mana base, artifacts, and creature board to carry the game.

The quality model grades this deck typical against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The mana base serves the colors evenly, and that raises the grade.

This deck plan aims for 2 or more finishers, the cards that win the game, and the deck holds 1.

- PLAN plan_coherent=partly: Most cards feed a white Equipment and Dwarf plan around Kíli, but 28 Plains beside seven utility lands, three sweepers, and several unrelated draw and removal cards leave the deck spread across small goals with few payoffs.
- PLAN theme_fit=yes: Every card is marked as in the Hobbit and Bloomburrow sets, the commander is a mono-white Hobbit Dwarf legend, and the Equipment plan fits a bracket 3 deck the person asked to have picked for them.
- PLAN useful_as_built=partly: The mana base is stable and the ramp and equipment are cheap, but 28 Plains plus artifact mana is a lot of mana, the deck has only one real finisher and few Dwarves for Kíli, and its late game is thin.
- PLAN summary_honest=partly: The summary names Angel of the Ruins, Helm of the Host and Sunscorch Regent as closers, and all three are in the list, but it leaves out that there are only a few real finishers and that the Dwarf count is too low to support the Dwarf payoffs it promises.

### 24. a 60-card deck from one set

**Summary:** This mono-red Bloomburrow aggro deck wins by building an attacking force of Mice, Lizards, and Raccoons, then using removal to keep its pressure pointed at the opponent. Emberheart Challenger and Hearthborn Battler reinforce the creature core, while Dragonhawk, Fate's Tempest is chief among the deck's threats. It gives up broad answers and a deep long game in exchange for a focused, consistent attack plan.

The quality model grades this deck at the precon baseline against the top lists of the format: the cards pair in ways the top lists do not, and that lowers the grade. Many of the cards appear in no top list, and that lowers the grade. Few of the cards are ones the top lists of the format play, and that lowers the grade.

- PLAN plan_coherent=partly: The red creatures and burn all push toward aggro, but the curve is top-heavy for an aggro deck with 14 creatures at four or five mana against only six one-drops, and the many cheap tricks and card draw spells do not tie tightly to the creature core.
- PLAN theme_fit=yes: The deck is mono-red, built entirely from Bloomburrow cards, and leans on creatures and burn in the way a casual Modern aggro deck would.
- PLAN useful_as_built=partly: With 24 Mountains and a heavy set of four- and five-drops the deck is playable and never has color trouble, but it floods easily, has few cheap plays, and the mix of weak tricks and card draw gives it a slow, inconsistent start for aggro.
- PLAN summary_honest=partly: The summary says Dragonhawk, Fate's Tempest is chief among the threats, which is true of the list as a top-end creature, but it calls Dragonhawk chief when it has only 2 copies against 4 each of several other threats, and it says the deck wins with Mice, Lizards, and Raccoons, which the list does carry.

### 25. use no card of an owned precon

**Summary:** Captain America leads a mono-white superhero squad built around deploying Heroes, protecting the team, and using Equipment to turn steady attacks into pressure. The deck develops its board with S.H.I.E.L.D. agents and Avengers-adjacent heroes, keeps cards flowing through its creature and Equipment pieces, and closes with major superhero threats such as The Sentry, Golden Guardian and Origin of Spider-Man. It gives up some flexibility in exchange for a focused white mana base and a creature-forward plan that can be vulnerable when the board is repeatedly cleared.

The quality model grades this deck at the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. The deck makes less mana on turn four than the norm, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade.

- PLAN plan_coherent=partly: The creature and Equipment core supports a Hero plan, but three board wipes (Austere Command, Fumigate, Vanquish the Horde) and about ten mana rocks in a creature deck pull against it, and 35 Plains leave many slots on filler.
- PLAN theme_fit=partly: It is a mono-white Captain America deck with many Marvel Heroes, which fits the Avengers superhero theme, but a lot of the non-Hero filler and generic Equipment dilutes the theme, and nothing in the list shows it drew on the person's own library or avoided the precon.
- PLAN useful_as_built=partly: The deck has 35 Plains, ten rocks, removal and a real curve of creatures, so it functions, but that many lands and rocks with only a handful of real finishers means flooding and few ways to win, and colorless rocks like Chromatic Lantern add little to a mono-color deck.
- PLAN summary_honest=partly: The summary names The Sentry, Golden Guardian and Origin of Spider-Man as finishers, and both are in the list, but Origin of Spider-Man is a cheap Saga and not a real major threat, and the summary does not mention the ten mana rocks or the three wipes that work against its creature-forward plan.

