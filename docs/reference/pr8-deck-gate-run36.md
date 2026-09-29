# PR-8 deck gate, a rejudge of pr8-deck-gate-run35

Run date: 2026-09-29. Card snapshot: 2026-09-04. Decks read from `pr8-deck-gate-run35.md`.

Verdict: PASS. every deck of the source passed its build bars, 0 summaries stated a false rule of the game, and 0 judge calls failed. The build rows are the rows of the source run, and no deck was built again (D-789).

- Calls: 50. Cost: $0.3358. Time: 142 seconds.

## Run

- Suite `decks`, run `pr8-deck-gate-run36`, on 2026-09-29, commit `9fb2569`.
- Roles: generate on `gpt-5.6-terra` (openai, effort medium), judge on `claude-sonnet-5-5` (anthropic, effort medium), repair on `gpt-5.6-terra` (openai, effort medium).
- Versions: card snapshot 2026-09-04, generate prompt version 16, plan_rubric prompt version 4, summary_judge prompt version 2, precons `5.3.0+20260923`, quality_model `20260923T202806Z`, rejudge_of `pr8-deck-gate-run35`.
- Calls: 50. Cost: $0.3358. Time: 142 seconds.
- Note: a rejudge of pr8-deck-gate-run35: the decks and the build rows are copied, and the judge rows are new (D-789)

## Decks

### 1. lifegain Commander, any card

**Summary:** This lifegain deck uses steady development and card advantage to make Karlov of the Ghost Council the central battlefield threat, with protection to preserve its key pieces and broad removal to clear resistance. Its creature suite applies pressure through lifegain payoffs while Angel of Destiny, Approach of the Second Sun, and Felidar Sovereign provide distinct ways to close a game. It gives up some speed for a resilient midrange plan that is strongest once its engines and mana are established.

The quality model grades this deck below the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The mana base serves the colors evenly, and that raises the grade.

- PLAN plan_coherent=yes: Nearly every card gains life, protects the engine, or removes blockers and wins through lifegain payoffs like Sanguine Bond, Archangel of Thune, and Angel of Destiny.
- PLAN theme_fit=yes: It is a W/B lifegain deck led by Karlov, with a good number of lifegain payoffs and a bracket 3 feel (Jitte, wraths, no infinite combos).
- PLAN useful_as_built=yes: It has 38 lands plus a dozen mana rocks, a smooth curve, and several win conditions, so it plays as built although the ramp includes some weak artifacts.
- PLAN summary_honest=yes: x

### 2. aristocrats Commander, owned first

**Summary:** Denethor turns a steady supply of creatures into sacrifice pressure, draining opponents while building a fresh Soldier force after creatures die. The deck develops through inexpensive mana and card engines, protects its important pieces, and uses broad spot interaction plus selective board resets to keep opponents from getting ahead. It closes through dedicated finishers or by grinding repeated sacrifice turns into an advantage, giving up raw speed for a board-dependent plan that wants creatures and time to establish itself.

The quality model grades this deck at the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. The mana base serves one color far better than another, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade.

- PLAN plan_coherent=partly: Ramp, draw, protection, removal, and wipes are all present, but the sacrifice payoffs and fodder generators are thin, and the wraths work against a creature-based plan.
- PLAN theme_fit=partly: It is a W/B Denethor deck at reasonable power, but the aristocrats theme is weakly represented with few sac outlets, few token makers, and few drain payoffs, and the request's 'from my library first' cannot be seen in the list.
- PLAN useful_as_built=partly: The mana is stable with plenty of ramp, but 25 basics with 17 Plains against few Swamps strain the black cards, and the win conditions are few and mostly not sacrifice-based.
- PLAN summary_honest=partly: placeholder

### 3. artifacts Commander, bracket 4

**Summary:** Urza leads a fast blue artifact deck that uses early mana, artifact tutors, and sustained card advantage to establish a dense board quickly. The deck protects its position with extensive stack interaction, answers troublesome permanents through artifact-based removal, and closes through artifact creature pressure or alternate endgames from Mechanized Production, Mirrodin Besieged, and Thassa's Oracle. It gives up broad multicolor answers and relies heavily on artifacts remaining central to its plan.

The quality model grades this deck typical against the top lists of the format: the deck makes more mana on turn four than the norm, and that raises the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The mana base serves the colors evenly, and that raises the grade.

- PLAN plan_coherent=yes: Every card is blue or colorless and feeds one artifact plan of fast mana, tutors, card draw, counterspells, and artifact-creature or alternate wincons under Urza.
- PLAN theme_fit=yes: It is a mono-blue Urza artifact deck with Moxen, Force of Will, and tutors, which matches a bracket 4 high-power request.
- PLAN useful_as_built=partly: Cheap mana and artifacts make it functional, but 25 Islands plus 8 nonbasic lands with painful rainbow lands, on top of the many free mana rocks, adds flood risk, while several win conditions are expensive or clunky (Myr Enforcer, Sojourner's Companion) and a few slots (Lion's Eye Diamond, Welding Jar) are weak here.
- PLAN summary_honest=partly: The summary matches most of the list, but it says 'extensive stack interaction' and 'artifact-based removal' while several removal slots are weak or odd, and it leaves out the 26-land count, the heavy fast-mana package, and the rainbow lands that do nothing in a mono-blue deck.

### 4. dinosaur tribal, bracket 2

**Summary:** This is a straightforward Dinosaur combat deck: build mana, develop Dinosaur support, then bring Gishath, Sun's Avatar into combat to turn its trigger into a large board of Dinosaurs. The main finish is an overwhelming attack backed by large Dinosaur threats and finishers, with removal and a few reset buttons available when the board gets crowded. It favors committing creatures and pressing combat over a highly reactive game plan, so careful sequencing and protecting Gishath matter.

The quality model grades this deck below the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The land count sits above the norm of the format, and that lowers the grade.

- PLAN plan_coherent=yes: Ramp, Dinosaur support, and many Dinosaur threats all feed the Gishath and Dinosaur combat plan, with only a few stray cards like Faithless Looting and Vandalblast.
- PLAN theme_fit=yes: It is a Dinosaur Commander deck led by Gishath in its own colors with a wide Dinosaur suite and mostly low-power cards suited to bracket 2, though the expensive fetch and dual lands are less friendly to a new player.
- PLAN useful_as_built=partly: There is enough ramp, threats, and protection to function, but about 37 lands in a three-color deck with many painlands and tapped lands leaves too few action cards and makes it clunky for a new player, and Gishath costs eight.
- PLAN summary_honest=partly: The Dinosaur combat plan and the removal are real, but the 'few reset buttons' are mostly Vandalblast and pinger-style creatures, and the summary hides a very heavy land count and a greedy three-color mana base.

### 5. blink Commander, owned first

**Summary:** Gilraen anchors a creature-focused blink plan, repeatedly setting up favorable returns while the deck develops mana, cards, equipment, and a protected board. It wins by turning its growing creature force sideways, with Angel of Serenity chief among its closing threats, while removal and wipes clear away resistance. The deck gives up multicolor flexibility for a very consistent Plains-based mana base and relies on creatures staying relevant on the battlefield.

The quality model grades this deck typical against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The mana base serves the colors evenly, and that raises the grade.

This deck plan aims for 2 or more finishers, the cards that win the game, and the deck holds 1.

- PLAN plan_coherent=partly: Mono-white with equipment, protection, and removal is loosely coherent, but the deck has few real blink enablers and many unrelated artifacts, so the cards do not clearly serve one blink plan.
- PLAN theme_fit=partly: It is bracket-3-appropriate mono-white with Gilraen, but the request was for a blink deck and there is little blink, and nothing shows it was built from the person's library first.
- PLAN useful_as_built=partly: With 36 Plains plus about 10 mana rocks the deck will flood, and it has only one real finisher with a thin creature count, so it plays but often sputters.
- PLAN summary_honest=no: The summary claims a repeated blink plan and lists Angel of Serenity among several closers, but the list has only Flickerwisp and a few incidental flicker effects, and it holds just one finisher.

### 6. Modern tempo, tournament

**Summary:** This blue-red tempo deck establishes pressure with Ragavan, Nimble Pilferer, Ledger Shredder, and Faerie Mastermind, then protects that pressure with efficient interaction and removal. Preordain and Mishra's Bauble keep the deck moving, while Temporal Mastery and Temporal Trespass reinforce its spell-focused plan. It gives up broader late-game threats and sweeping answers in exchange for a fast, focused game built around trading efficiently and staying ahead.

The quality model grades this deck typical against the top lists of the format: the cards pair in ways the top lists do not, and that lowers the grade. The deck runs few spells as one copy, and that raises the grade. The deck runs its spells as playsets, and that lowers the grade.

- PLAN plan_coherent=partly: The cheap creatures, counters and burn form a coherent tempo core, but Temporal Mastery (7 mana) and Temporal Trespass (11 mana) pull against a low-curve plan and are named as reinforcing it.
- PLAN theme_fit=yes: It is a blue-red Modern deck with cheap threats, counterspells and burn, which matches a tempo request, though it lacks Delver of Secrets itself.
- PLAN useful_as_built=partly: The core is playable, but 28 lands (too many for this curve) and dead 7- and 11-mana cards leave it flooding with too few threats and win conditions.
- PLAN summary_honest=no: It calls the Temporal spells reinforcement for a fast, focused game, though they are near-uncastable in this deck, and it omits that the deck lists 28 lands, plus Mishra's Bauble, with almost no real Delver threats.

### 7. Modern burn, casual

**Summary:** This mono-red burn deck applies early pressure with efficient removal and spell-focused synergy creatures, then keeps the cards coming to sustain its assault. Its threats give it a board-based route to victory when direct burn alone is not enough, with Hazoret the Fervent chief among them as a closing threat. The deck gives up flexibility against specialized strategies in exchange for a focused, consistent red mana base and a direct aggressive game plan.

The quality model grades this deck at the precon baseline against the top lists of the format: the cards pair in ways the top lists do not, and that lowers the grade. The deck runs few spells as one copy, and that raises the grade. Many of the cards appear in no top list, and that lowers the grade.

- PLAN plan_coherent=partly: Bolt, Eidolon and Thermo-Alchemist form a burn core, but four-drop creatures like Hazoret (which wants an empty hand), Barret Wallace and Champion of the Path pull toward a midrange plan, and Browbeat plus Eidolon pulls both ways.
- PLAN theme_fit=yes: It is a mono-red Modern deck built around burn spells, Eidolon and Thermo-Alchemist, which matches the request.
- PLAN useful_as_built=no: With 24 Mountains and only 12 real burn spells, the deck floods badly, has too few ways to deal damage and a top-heavy curve of 4-drops, so it will not play as a functioning burn deck.
- PLAN summary_honest=partly: It correctly names Hazoret, burn and spell-synergy creatures, but 'efficient removal' and 'early pressure' overstate a list with only 6 cheap burn spells and a top-heavy creature suite.

### 8. Modern lifegain, FNM

**Summary:** This white-black lifegain deck establishes its early board with Guide of Souls and Ocelot Pride, then turns that foundation into pressure through Enduring Tenacity, Twinblade Paladin, and its larger threats. Solitude and Murderous Rider keep opposing threats from taking over while Lembas and Enduring Innocence help sustain the hand. It wins by building a resilient lifegain-based board and closing with its powerful creatures; the tradeoff is that several of its strongest threats ask the deck to reach the middle and later stages of the game.

The quality model grades this deck below the precon baseline against the top lists of the format: the cards pair in ways the top lists do not, and that lowers the grade. The deck runs few spells as one copy, and that raises the grade. The deck runs its spells as playsets, and that lowers the grade.

- JUDGE [unknown]: "The quality model grades this deck below the precon baseline against the top lists of the format: the cards pair in ways the top lists do not, and that lowers the grade.". This describes how a grading model scores the deck, not a rule of the game. It is not a rules claim, so it is listed only for completeness.
- PLAN plan_coherent=yes: Cheap lifegain enablers (Guide of Souls, Ocelot Pride, Lembas, Enduring creatures) feed lifegain payoffs (Twinblade Paladin, Archangel of Thune, Sheoldred), backed by removal, so the cards serve one plan.
- PLAN theme_fit=yes: It is a white-black lifegain deck legal for Modern, built from playsets at a power level suited to FNM.
- PLAN useful_as_built=yes: The 24-land mana base of shocklands, Caves of Koilos, Shattered Sanctum and basics supports both colors, and the curve has cheap one-drops, four-drops and Solitude, so it plays as it stands, though it has few cheap interaction spells.
- PLAN summary_honest=partly: The summary describes real cards and admits the top-end tradeoff, but it never mentions Sheoldred, Archangel of Thune or Altar of the Pantheon and it calls the deck resilient though it has little interaction beyond Solitude and 2 Riders.

### 9. Standard midrange, FNM

**Summary:** This black-green midrange deck builds its mana around Overgrown Tomb and Wastewood Verge, uses Llanowar Elves to reach its heavier creatures, and keeps cards flowing through Darkstar Augur and Phyrexian Arena. Goldvein Hydra and Chomping Changeling apply steady pressure, while Vein Ripper, Massacre Wurm, and Vaultborn Tyrant provide heavier finishes. Bitter Triumph and Maelstrom Pulse clear the way, and Snakeskin Veil plus Not Dead After All help the creature plan survive resistance. It gives up a broad range of alternate angles in favor of committing to creatures, protection, and direct answers.

The quality model grades this deck below the precon baseline against the top lists of the format: the cards pair in ways the top lists do not, and that lowers the grade. Few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade.

- PLAN plan_coherent=partly: The creatures, removal, and card draw all fit a midrange shell, but the protection tricks (four Snakeskin Veil, four Not Dead After All) sit on top of only about 18 creatures, and only 2 Llanowar Elves support six-mana finishers.
- PLAN theme_fit=yes: It is a black-green creature midrange deck in Standard with a reasonable power level for FNM, using both colors and midrange roles.
- PLAN useful_as_built=partly: With 24 lands, 4 Overgrown Tomb, and 4 Verge, the mana is fine, but the deck has 4 Not Dead After All and 4 Snakeskin Veil with few ways to use them well, only 2 Llanowar Elves, and 6-drops with BBB costs, so it will stumble often though it is playable.
- PLAN summary_honest=partly: The named cards are all in the list, but the summary presents the deck as coherent and steady while it leaves out the heavy top end and thin ramp, and it ignores the many protection spells that lack good targets.

### 10. Standard aggro, tournament

**Summary:** This red-white aggro deck commits threats early, backs its attacks with removal and flexible interaction, and uses its draw cards to keep pressure on after the first exchange. Warleader's Call is chief among the cards tying the creature plan together. It aims to win by maintaining a fast board presence and clearing the way rather than by relying on a slow, high-end finish.

The quality model grades this deck below the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. The cards pair in ways the top lists do not, and that lowers the grade. The curve sits high for the format, and that lowers the grade.

- PLAN plan_coherent=partly: Warleader's Call and the token-ish creatures loosely support a go-wide plan, but the deck has few cheap threats, a top-end Bedhead Beastie at six mana, and several spot enchantments that do not push an aggressive plan.
- PLAN theme_fit=yes: The deck is red-white and uses Boros lands and cards, and it is built for Standard as requested, although it plays more midrange than a fast aggro deck.
- PLAN useful_as_built=partly: With 28 lands and a mana base that is fine, the deck can be played, but only 16 creatures and a six-drop leave it short of early pressure and clean ways to close games.
- PLAN summary_honest=partly: placeholder

### 11. Commander with a locked card

**Summary:** This is an Orzhov sacrifice deck built to turn a broad creature board into steady pressure, with Karlov of the Ghost Council as a growing centerpiece. Sacrifice outlets, death-focused payoffs, and creature-based value keep the board working through exchanges, while high-impact threats and dedicated closing cards end drawn-out games. It gives up some speed to retain a deep creature package and several expensive top-end plays.

The quality model grades this deck at the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. Few lands enter tapped, and that lowers the grade.

- PLAN plan_coherent=yes: Nearly every card is a sacrifice outlet, a death-trigger payoff, or a fodder creature, and the Orzhov colors match the commander's plan of growing from creature deaths.
- PLAN theme_fit=yes: It is a Karlov-led Orzhov sacrifice deck in Commander, with the commander in the command zone and the list made of sacrifice outlets and payoffs.
- PLAN useful_as_built=partly: The 40 or so lands with Sol Ring and two altars are enough mana, and the outlets, drain payoffs, and wipes give it ways to win, but heavy black weighting and the uncounted-for basics leave too little white for its white cards, and the count of real ramp beyond Sol Ring is thin.
- PLAN summary_honest=yes: The summary describes an Orzhov sacrifice deck with outlets, death payoffs, creature value, big threats, and closers, all of which the list carries, and it admits the slower, top-heavy shape.

### 12. Commander on a budget

**Summary:** Adeline, Resplendent Cathar drives a wide white token strategy, turning attacks into an expanding Human force and using the deck’s token-focused support to make that board matter. The deck develops with plentiful mana and card flow, keeps opposing boards in check with targeted answers and resets, then closes through combat pressure or dedicated finishing cards such as Halo Fountain, Luck Bobblehead, and Sword of Body and Mind. It favors steady board development over explosive early-game acceleration, so it is most effective when it can keep creatures on the table and attack repeatedly.

The quality model grades this deck below the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. The deck makes less mana on turn four than the norm, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade.

- PLAN plan_coherent=yes: Nearly every card is white and supports a token-and-attack plan with token makers, anthems, doublers, sweepers and finishers all feeding Adeline's wide board.
- PLAN theme_fit=yes: It is a mono-white token Commander deck led by Adeline, with budget-friendly and bracket 2-appropriate cards and no obvious infinite or fast-mana pieces.
- PLAN useful_as_built=partly: The 30 Plains plus 7 utility lands give a stable mana base and plenty of playable threats, but several ramp and draw artifacts are low-impact filler and the deck has too many clunky top-end cards, so it will feel sluggish.
- PLAN summary_honest=partly: The token plan, the card flow and the named finishers are all really in the list, but the summary skips over the weak ramp (Coin of Mastery, Currency Converter, Druidic Satchel and Collector's Vault are marginal) and calls the mana 'plentiful' though the deck makes less mana on turn four than the norm.

### 13. owned first, and the commander is not owned

**Summary:** Karlov of the Ghost Council is the focal point of a white-black lifegain deck that develops mana, gains cards through a mix of creatures and equipment, and turns steady life gains into a dangerous commander. Protective equipment and instant-speed safeguards help Karlov stay active while targeted answers clear the way. The deck closes through Karlov's accumulated pressure, evasive Angel threats, and finishers such as Frodo, Sauron's Bane, Lyra Dawnbringer, and Grave Venerations; in exchange, it leans heavily on creatures and artifacts remaining in play to sustain its momentum.

The quality model grades this deck at the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The land count sits above the norm of the format, and that lowers the grade.

- PLAN plan_coherent=partly: Lifegain with Karlov is present through Angels and lifegain creatures, but the list is padded with nine mana rocks, a large equipment pile, and three wraths that work against a creature-and-equipment plan.
- PLAN theme_fit=partly: It is a W/B Karlov deck at a plausible bracket 3, but lifegain payoffs are thin and the deck reads more as generic good stuff, with no sign of a library-first choice.
- PLAN useful_as_built=partly: It has 36+ lands plus many rocks and is heavily color-weighted to white, so it will flood, and it has few real lifegain engines and only a handful of finishers, though removal and mana are solid.
- PLAN summary_honest=partly: x

### 14. the user delegates the commander

**Summary:** Atarka, World Render leads a Gruul Dragon deck that ramps into a dense flight of threatening Dragons, then turns attacks into decisive double-strike combat steps. Tribal cost support helps deploy the expensive core, while draw, protection, removal, and sweepers keep the pressure sustained. The primary wins come from overwhelming combat with Atarka and the deck’s chief Dragon finishers, with the tradeoff that the deck remains creature-centered and can need its mana development to line up before its heaviest threats take over.

The quality model grades this deck below the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The land count sits above the norm of the format, and that lowers the grade.

- PLAN plan_coherent=yes: Ramp, Dragon-cost support, and a heavy red Dragon suite all push toward casting big Dragons and attacking, with Atarka as an extra double-strike finisher.
- PLAN theme_fit=yes: It is a Gruul Dragon Commander deck led by a Dragon commander, with about 25 Dragons and Dragon-support cards that suit a bracket 3 request.
- PLAN useful_as_built=partly: It has plenty of ramp, draw, and Dragon threats, but 38 lands plus many mana rocks is flood-prone, 12 Mountains against a heavy RR/RRR curve with only a few Forest sources is workable yet clunky, and the multiple Orbs and filler artifacts are weak.
- PLAN summary_honest=yes: placeholder

### 15. delegated commander, owned first

**Summary:** A black-white lifegain deck built to turn incremental healing into an overwhelming board of creatures, resilient equipment carriers, and powerful flying finishers. It combines efficient removal, protective tools, and several board resets to maintain control while its life-total advantages become decisive pressure.

The quality model grades this deck at the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The land count sits above the norm of the format, and that lowers the grade.

- PLAN plan_coherent=partly: Lifegain payoffs (Angel of Vitality, Lyra, Minwu, Exemplar) and equipment fit together, but four board wipes, a mass of loose equipment, and a few off-plan cards (Canyon Crawler, Rooftop Percher) pull against the creature-based plan.
- PLAN theme_fit=partly: It is a WB lifegain deck led by a commander the builder chose, but the request said 'from my library first' and nothing in the list shows that, and Astarion at 6 mana is an odd lifegain pick.
- PLAN useful_as_built=partly: It has plenty of ramp, removal, and flying finishers, but 14 Swamps with a white-heavy spell list (WW and WWW costs) and roughly 39 lands plus rocks give an unbalanced mana base and flood risk.
- PLAN summary_honest=partly: The summary describes lifegain, equipment carriers, flyers, removal, and wipes that are all in the list, but it hides a weak lifegain core, a mana base of 18 Plains and 14 Swamp that is clumsy for a white-heavy deck, and the fact that the deck is over-ramped.

### 16. a tight budget, owned first

**Summary:** This is a patient Orzhov aristocrats deck that develops creatures and support pieces, then uses Denethor’s sacrifice ability to turn creature deaths into life swings while building an end-step Soldier force. The finishing cards give the deck ways to close once the board has been established, while the removal, protection, and board resets help it survive longer games. It gives up explosive speed for a library-first mana base and an incremental, board-centered plan.

The quality model grades this deck at the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The deck makes less mana on turn four than the norm, and that lowers the grade.

- PLAN plan_coherent=partly: Most cards are Orzhov creatures, sacrifice outlets and removal that loosely support an aristocrats plan, but the deck also carries three board wipes, about ten mana rocks (including off-plan ones like Astral Cornucopia and Chromatic Lantern), and generic removal, so the pieces pull in different directions.
- PLAN theme_fit=partly: The deck is a legal Orzhov aristocrats build at a low power level, but the request asked for cards from the user's library first and the 25-dollar budget, and Sol Ring, Skullclamp, Swords to Plowshares, Austere Command and Infernal Grasp make it doubtful the deck respects that limit or bracket 2.
- PLAN useful_as_built=partly: With 27 basics plus 10 utility lands and about ten rocks, the mana is plentiful and consistent, and the deck has removal and plenty of creatures to play, but it lacks reliable sacrifice outlets and drain payoffs, so it may struggle to close games.
- PLAN summary_honest=partly: The summary describes an aristocrats deck with Denethor's sacrifice ability and an end-step Soldier force, and it does name the wipes and removal, but it glosses over the heavy rock package and the mostly incidental sacrifice payoffs, and it never mentions the 'library first' claim beyond calling the mana base library-first.

### 17. upgrade a precon, owned first

**Summary:** Zada Goblin Storm builds a wide Goblin board and turns creature-targeted spells into broad bursts of cards, pressure, and momentum. The deck wins by converting that developed board into a decisive attack, with Great Train Heist and Collective Inferno chief among its closing cards, while Goblin-linked removal clears resistance. It gives up some consistency to retain the precon's varied Goblin package and remains most effective when it has both creatures and a stocked hand.

The quality model grades this deck typical against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The deck makes less mana on turn four than the norm, and that lowers the grade.

- PLAN plan_coherent=partly: The Goblin token and creature core hangs together, but the many cheap cantrips, Zada-targeting spells, Grapeshot and Past in Flames pull toward a storm plan that the mostly mono-red Goblin board only partly supports.
- PLAN theme_fit=yes: It is a mono-red Goblin Zada deck that keeps the precon's Goblin and Zada-targeting core while adding staples like Sol Ring, Skullclamp and Krenko.
- PLAN useful_as_built=partly: It has 38 or so lands and rocks with plenty of Goblins and token makers, but the many low-impact filler spells and slow mana on turn four make it clunky, and Great Train Heist is a weak finisher.
- PLAN summary_honest=partly: placeholder

### 18. upgrade a precon, any card

**Summary:** This deck keeps Turtle Power’s broad Mutant, Ninja, and Turtle creature core while giving its mana base a cleaner route to all five colors. Build a board of themed creatures, send them into combat, and let Heroes in a Half Shell turn successful hits into larger attackers and more cards. Raphael, the Muscle, Dimension X Pizzasaur, and Everything Pizza are chief among the cards that can close a game. The deck gives up some of the precon’s slower lands to improve early development, while retaining its varied character-driven creature package and big finishing turns.

The quality model grades this deck below the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The land count sits above the norm of the format, and that lowers the grade.

- PLAN plan_coherent=partly: Most cards are Mutant/Ninja/Turtle creatures feeding Heroes in a Half Shell's combat plan, but a pile of loose five-color goodstuff, several sweepers, and a couple of vehicle/artifact pieces pull away from a single attack-focused plan.
- PLAN theme_fit=yes: It keeps Heroes in a Half Shell as commander and the TMNT creature core, in five colors, as an upgrade of the Turtle Power precon.
- PLAN useful_as_built=yes: With dual lands, Sol Ring, Signet, Lantern, and Cultivate, the five-color mana is reliable, the curve is low, and there are enough creatures and finishers to play the deck as it stands.
- PLAN summary_honest=partly: The creature-and-combat plan and named finishers do appear in the list, but the summary calls the deck cleaner and quicker while the land count sits well above the norm, and it never mentions the sweepers or the lack of ramp beyond a few rocks.

### 19. the Hobbit family, two colours

**Summary:** Thranduil leads an Elf-centered Sultai deck that develops its mana, fills the battlefield with Elf creatures, and turns legendary Elf arrivals into fresh cards while using Elf cards in the graveyard as a resource. The deck controls key opposing pieces with a broad mix of removal, disruption, and sweepers, then closes through its large marked finishers, chief among them the two Witch-kings. It gives up some raw speed for a creature-heavy, theme-driven plan that needs its mana and Elf presence to stay established.

The quality model grades this deck at the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. Many of the cards appear in no top list, and that lowers the grade.

- PLAN plan_coherent=partly: The Elf creatures, ramp, and draw pieces fit one Elf plan, but the heavy black removal, three sweepers, and Witch-king finishers pull against a wide board of Elf creatures.
- PLAN theme_fit=yes: It is a Sultai Thranduil, the Elvenking deck built entirely from Hobbit and Hobbit commander set cards with an Elf focus, which fits the request.
- PLAN useful_as_built=partly: With 35 lands, a dozen ramp pieces, and plenty of draw it should run smoothly, but the mana is a heavy three-color spread with many basics and few fixers, and the finishers are thin and the sweepers clash with the creature plan.
- PLAN summary_honest=partly: The Elf, ramp, removal, and Witch-king claims all match the list, but the summary says nothing of the thin Elf count in black and the sweepers that undo its own board, and it calls the Witch-kings 'chief' finishers when the deck has few real ways to win.

### 20. the Hobbit family, a delegated commander

**Summary:** A mono-red Hobbit Dragon deck centered on building Treasures, equipping resilient creatures, and leveraging combat to create overwhelming pressure. Smaug turns accumulated wealth into direct damage, while Dragons, giants, and a wide supporting cast provide multiple routes to a decisive finish. Removal and sweepers keep opposing boards manageable while legendary artifacts reinforce the deck’s adventurous Middle-earth character.

The quality model grades this deck at the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The mana base serves the colors evenly, and that raises the grade.

- PLAN plan_coherent=partly: Treasure making, equipment, and Dragons all appear, but the list is a pile of about 16 equipment pieces and vanilla-ish Goblins and Dwarves with only a few cards actually feeding Smaug's Treasure payoff.
- PLAN theme_fit=yes: It is a mono-red Dragon commander, Smaug the Magnificent, with every card from the Hobbit sets, matching the request.
- PLAN useful_as_built=partly: With 36 lands and mono-red mana it casts its spells, but the deck has many low-impact equipment and Goblin cards and few reliable ways to win, so it plays sluggishly.
- PLAN summary_honest=partly: Treasure, equipment, Dragons and sweepers are all present, but the claim of a plan that turns treasure into damage and the description of multiple routes to a finish overstate a deck with only a handful of real finishers.

### 21. the Hobbit family, mana from outside

**Summary:** This Rakdos Smaug deck develops its mana early, deploys Smaug the Impenetrable and other large threats, then presses combat while using focused removal and broad resets to clear resistance. Smaug is chief among the finishers, with Dragons and other legendary threats providing several ways to end a stalled game. It gives up some flexibility for a committed battlefield plan and depends on its mana development to bring its expensive threats online quickly.

The quality model grades this deck at the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The deck makes less mana on turn four than the norm, and that lowers the grade.

- PLAN plan_coherent=partly: Ramp into big Dragons and legends with removal and wipes is a clear shape, but a large pile of situational equipment (Sting, Glamdring, both Andúrils, Lances, Mattock) has few creatures to carry it and pulls away from the ramp-and-finish plan.
- PLAN theme_fit=yes: It is a Rakdos Commander deck led by Smaug the Impenetrable, built almost entirely from Hobbit-set cards with off-set lands allowed, and it fits a mid-power bracket 3 game.
- PLAN useful_as_built=partly: The 37-land base with fixing is stable, but few real ramp pieces, about a dozen weak equipment cards, and a thin creature count leave the deck slow and short on ways to win.
- PLAN summary_honest=partly: The summary describes ramp, big threats, removal and wipes accurately, but calling Smaug a chief finisher and promising fast mana overstates a deck that makes less mana on turn four and carries many low-impact equipment cards.

### 22. a set family and a card from outside it

**Summary:** This is a Sultai Elf deck centered on Thranduil, the Elvenking, using legendary Elves and a full Elf creature base to develop pressure while keeping mana and cards flowing. Targeted answers, protective tools, and sweepers give it room to stabilize, then Troll of Khazad-dûm and the two Witch-kings provide its strongest closing threats. It gives up some raw speed for a creature-centric plan that benefits from a developed board and a stocked graveyard.

The quality model grades this deck at the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. Many of the cards appear in no top list, and that lowers the grade.

- PLAN plan_coherent=partly: There is a real Elf-tribal core led by Thranduil, but the deck also packs three sweepers, many equipment pieces and artifact-draw cards, and a Witch-king top end that pull against a wide Elf board.
- PLAN theme_fit=yes: It is a Sultai Thranduil deck built from the Hobbit sets with Sol Ring kept in, as the request asked, and it looks like a bracket 3 deck.
- PLAN useful_as_built=partly: It has about 27 basics plus 9 nonbasic lands and lots of ramp and draw, but the three-color mana with many double-pip cards in each color, the sweepers that hurt its own creatures, and only a few real finishers make it clunky.
- PLAN summary_honest=partly: placeholder

### 23. two set families at once

**Summary:** Kíli the Resourceful leads a mono-white Dwarf and Equipment deck that develops through artifact mana, cheap gear, and a steady flow of creatures. Establish an enduring story early, use Kíli to turn each turn's first equip into efficient pressure, and keep cards flowing as Dwarves and Equipment enter. The deck wins by building one or more well-equipped attackers, with Angel of the Ruins, Helm of the Host, and Sunscorch Regent serving as major closing threats. It gives up broad color access and relies on its white mana base, artifacts, and creature board to carry the game.

The quality model grades this deck typical against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The mana base serves the colors evenly, and that raises the grade.

This deck plan aims for 2 or more finishers, the cards that win the game, and the deck holds 1.

- PLAN plan_coherent=partly: Equipment, artifact mana and cheap Dwarves broadly point at one Kíli equip plan, but the list also holds sweepers, spot removal and assorted value cards, and it has only a thin creature count to carry all those pieces.
- PLAN theme_fit=yes: Every card comes from the Hobbit and Bloomburrow sets, the deck is mono-white with a fitting Hobbit commander, and the Dwarf and Equipment plan suits a bracket 3 build.
- PLAN useful_as_built=partly: With 28 Plains and 8 utility lands plus a lot of cheap mana rocks it will cast its spells, but it has very few creatures to carry equipment, wipes that work against its own plan, and only one true finisher.
- PLAN summary_honest=partly: The summary describes the Dwarf and Equipment plan accurately, but it names Angel of the Ruins, Helm of the Host and Sunscorch Regent as major closers when Helm is a slow equipment, and it says nothing about the heavy pile of sweepers and the single real finisher.

### 24. a 60-card deck from one set

**Summary:** This mono-red Bloomburrow aggro deck wins by building an attacking force of Mice, Lizards, and Raccoons, then using removal to keep its pressure pointed at the opponent. Emberheart Challenger and Hearthborn Battler reinforce the creature core, while Dragonhawk, Fate's Tempest is chief among the deck's threats. It gives up broad answers and a deep long game in exchange for a focused, consistent attack plan.

The quality model grades this deck at the precon baseline against the top lists of the format: the cards pair in ways the top lists do not, and that lowers the grade. Many of the cards appear in no top list, and that lowers the grade. Few of the cards are ones the top lists of the format play, and that lowers the grade.

- PLAN plan_coherent=partly: The deck is all red creatures and burn, but 24 Mountains with a top-heavy curve of 4- and 5-drops pulls against an aggro plan, and the spells hold little synergy with the Mouse/Lizard/Raccoon mix.
- PLAN theme_fit=yes: It is mono-red, made entirely of Bloomburrow cards, and built around creatures and burn, which matches a casual aggro request in Modern.
- PLAN useful_as_built=partly: It is castable and has plenty of creatures and removal, but 24 lands for a deck with only 2 one-drops and many 4-5 drops is flood-prone and slow, with weak reach to close games.
- PLAN summary_honest=partly: The summary is broadly accurate about creatures and removal, but calling Dragonhawk the chief threat while playing only 2 copies, and claiming a 'focused, consistent' attack, hides the clunky curve and heavy land count.

### 25. use no card of an owned precon

**Summary:** Captain America leads a mono-white superhero squad built around deploying Heroes, protecting the team, and using Equipment to turn steady attacks into pressure. The deck develops its board with S.H.I.E.L.D. agents and Avengers-adjacent heroes, keeps cards flowing through its creature and Equipment pieces, and closes with major superhero threats such as The Sentry, Golden Guardian and Origin of Spider-Man. It gives up some flexibility in exchange for a focused white mana base and a creature-forward plan that can be vulnerable when the board is repeatedly cleared.

The quality model grades this deck at the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. The deck makes less mana on turn four than the norm, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade.

- PLAN plan_coherent=partly: Heroes and Equipment in mono-white is a real plan, but three board wipes and a pile of generic ramp rocks in a 35-Plains deck pull against the creature-forward Hero plan.
- PLAN theme_fit=partly: It is a mono-white Avengers Commander deck led by Captain America with many Marvel Heroes, but many filler non-Marvel cards dilute the theme, and nothing shows the deck was built from the player's library or kept clear of the Avengers Assemble precon.
- PLAN useful_as_built=partly: With 35 Plains and 10 ramp artifacts the mana works, but Chromatic Lantern, Astral Cornucopia and Bender's Waterskin add little, the deck is land-heavy with few threats, and the finishers are thin.
- PLAN summary_honest=partly: The Heroes, Equipment and named threats are all present, but the summary skips the heavy artifact ramp and three board wipes, and it says 'protecting the team' while the deck also contains sweepers that hurt that team.

