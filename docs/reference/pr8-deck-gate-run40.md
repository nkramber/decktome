# PR-8 deck gate, a rejudge of pr8-deck-gate-run39

Run date: 2026-10-04. Card snapshot: 2026-09-04. Decks read from `pr8-deck-gate-run39.md`.

Verdict: PASS. every deck of the source passed its build bars, 0 summaries stated a false rule of the game, and 0 judge calls failed. The build rows are the rows of the source run, and no deck was built again (D-789).

- Calls: 54. Cost: $0.3743. Time: 149 seconds.

## Run

- Suite `decks`, run `pr8-deck-gate-run40`, on 2026-10-04, commit `340e581`.
- Roles: generate on `gpt-6.1-sol` (openai, effort medium), judge on `claude-sonnet-5-5` (anthropic, effort medium), repair on `gpt-6.1-sol` (openai, effort medium).
- Versions: card snapshot 2026-09-04, generate prompt version 18, plan_rubric prompt version 5, summary_judge prompt version 3, precons `5.3.0+20260923`, quality_model `20260923T202806Z`, rejudge_of `pr8-deck-gate-run39`.
- Calls: 54. Cost: $0.3743. Time: 149 seconds.
- Note: a rejudge of pr8-deck-gate-run39: the decks and the build rows are copied, and the judge rows are new (D-789)

## Decks

### 1. lifegain Commander, any card

**Summary:** Karlov leads a lifegain deck that develops early, keeps cards flowing, and builds toward sustained battlefield pressure. Use interaction to keep dangerous opposing plays in check while Karlov and the threat package establish control. Win through combat or dedicated lifegain finishers, chief among them Angel of Destiny, Felidar Sovereign, and Test of Endurance. The deck favors a steady, interactive game over an all-in combo plan, giving up some explosive speed and remaining vulnerable when its battlefield is repeatedly cleared.

The quality model grades this deck typical against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. The cards pair in ways the top lists do not, and that lowers the grade. The commander is a popular one, and that lowers the grade.

- PLAN plan_coherent=partly: Most of the cards are lifegain creatures, Angels and Clerics, but the list also holds three sweepers (Damn, Toxic Deluge, Fumigate) that clear the creature board the deck builds, and a pile of cheap fast mana and interaction pieces that do not feed lifegain.
- PLAN theme_fit=yes: The deck is led by Karlov of the Ghost Council in white-black and is full of lifegain payoffs such as Archangel of Thune, Lyra Dawnbringer, Righteous Valkyrie, Vito, Voice of the Blessed and Test of Endurance, at a reasonable bracket 3 power level.
- PLAN useful_as_built=yes: With a full mana base of dual lands and basics, plenty of cheap ramp, card draw, spot removal and several big finishers, the deck can be played as it stands, even though some pieces like Vexing Bauble and Nuka-Cola Vending Machine are weak.
- PLAN summary_honest=partly: The summary names Angel of Destiny, Felidar Sovereign and Test of Endurance as lifegain finishers and all three are in the list, but it never mentions that the deck runs three board wipes that work against its own creature plan.

### 2. aristocrats Commander, owned first

**Summary:** Denethor leads a patient aristocrats deck that turns expendable creatures into life loss and replenishes the board with end-step tokens. Develop mana and card draw first, then use sacrifice synergies, protective interaction, and removal to keep the engine working through a long game. Steady drains and larger finishers provide complementary ways to close, with combat offering another route when opponents leave an opening. The deck gives up explosive speed and a dense dedicated-drain package in favor of resilient utility creatures and broad answers.

The quality model grades this deck at the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. The mana base serves one color far better than another, and that lowers the grade. The commander does not place in cEDH events, and that lowers the grade.

- PLAN plan_coherent=partly: Sacrifice outlets, Skullclamp, drain effects and token makers support an aristocrats plan, but three board wipes, many protection spells and a pile of generic ramp and utility creatures dilute it.
- PLAN theme_fit=partly: Denethor is the requested commander in WB at a reasonable power level, but the 'from my library first' request cannot be seen in the list and the aristocrats core is thin.
- PLAN useful_as_built=partly: The deck has lots of ramp, draw and removal and a workable curve, but 26 basics with 17 Plains and 9 Swamp strain black costs like BB, and there are few real drain payoffs or ways to win.
- PLAN summary_honest=partly: The summary is wrong.

### 3. artifacts Commander, bracket 4

**Summary:** Your commander turns a wide artifact board into mana, letting you develop threats while keeping interaction ready. Cheap acceleration, draw engines, and tutors support explosive turns, with artifact-heavy combat and dedicated finishers providing multiple routes to victory. The deck gives up some resilience to artifact hate and some opening-hand consistency to fit both a strong combo plan and a substantial battlefield presence.

The quality model grades this deck good against the top lists of the format: the cards are ones the top lists of the format play, and that raises the grade. The mana base serves the colors evenly, and that raises the grade. The commander places in cEDH events, and that raises the grade.

- PLAN plan_coherent=yes: Every card is blue or colorless and the list feeds one mono-blue artifact plan of Moxen and mana rocks, draw engines, free counterspells, artifact threats and Urza-based combo finishers.
- PLAN theme_fit=yes: It is a mono-blue Urza, Lord High Artificer deck packed with artifacts, Moxen and free counterspells, which matches a high-power bracket 4 artifact request.
- PLAN useful_as_built=partly: The deck has lots of acceleration and interaction, but 23 Islands plus utility lands and Mishra's Workshop alongside several 7-drops and bounce spells make it clunky, with few reliable ways to actually close the game.
- PLAN summary_honest=partly: The summary's claim of cheap acceleration, draw engines and tutors is mostly borne out by the Moxen, Rhystic Study and Whir of Invention, but the list has few true tutors and only a thin combat plan with several clunky 7-drops, so 'artifact-heavy combat' is overstated.

### 4. dinosaur tribal, bracket 2

**Summary:** Develop your mana early, build a herd of Dinosaurs, and use your commander's attacks to expand the battlefield. Large Dinosaur threats and dedicated finishers, chief among them Bonehoard Dracosaur and Dinosaurs on a Spaceship, support a combat-focused win. Card draw keeps the creatures coming, while removal, protective interaction, and board resets help you recover from opposition. The deck favors a straightforward creature plan over combos, giving up explosive speed and some flexibility for a slower, Dinosaur-heavy game.

The quality model grades this deck below the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. The mana base serves one color far better than another, and that lowers the grade. Many of the cards appear in no top list, and that lowers the grade.

- PLAN plan_coherent=yes: Ramp, Dinosaur creatures, green-heavy card draw, and a few removal and wipe pieces all feed one plan of building a Dinosaur board and attacking with Gishath.
- PLAN theme_fit=yes: It is a Gishath Dinosaur Commander deck in the correct three colors with a friendly, creature-centered plan that suits a new player at bracket 2.
- PLAN useful_as_built=partly: The deck has plenty of ramp, dinosaurs, and a smooth curve, but the mana base includes painful lands like Mana Confluence and City of Brass plus Horizon Canopy, and it has only 1 Plains and 1 Mountain with heavy red and white costs such as RR on Dracosaur and WW on Goring Ceratops, which is rough for a new player and for the color balance.
- PLAN summary_honest=partly: The summary names Bonehoard Dracosaur and Dinosaurs on a Spaceship as finishers and both are in the list, but it says nothing of the heavy green weighting or of the many painful and fixing-only lands, and the 'protective interaction' is thin.

### 5. blink Commander, owned first

**Summary:** Gilraen, Dúnedain Protector leads a creature-heavy blink deck that builds value through draw creatures, removal creatures, and protective support. Develop mana early, establish a useful creature board, and keep resources available for Gilraen rather than committing everything at once. Smaller creatures provide steady pressure, while the larger closing threats, chief among them Angel of Serenity, help finish the game. The deck gives up explosive speed for a slower value game, and many of its creatures prioritize utility over raw combat power.

The quality model grades this deck typical against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. The mana base serves the colors evenly, and that raises the grade. The commander does not place in cEDH events, and that lowers the grade.

This deck plan aims for 2 or more finishers, the cards that win the game, and the deck holds 1.

- PLAN plan_coherent=yes: Nearly every card is a white creature or support piece that fits a mono-white value-blink plan, with ETB creatures like Inspiring Overseer, Fiend Hunter, Palace Jailer and Flickerwisp, plus ramp, protection and removal around them.
- PLAN theme_fit=partly: The deck is a mono-white creature value deck at roughly bracket 3, but the request asked for blink and the list holds only a few true blink effects, while the Commander's color identity is respected.
- PLAN useful_as_built=partly: The mana base of 29 Plains plus utility lands and ten ramp pieces functions well, and the removal and draw are solid, but with only one finisher, and three board wipes in a creature deck, the deck has too little ability to close out games.
- PLAN summary_honest=partly: The summary names Angel of Serenity as a closing threat and that card is in the list, but it calls this a blink deck about protecting Gilraen while the list has few real blink enablers beyond Flickerwisp and Gilraen's own effects, and it never mentions the three wraths that cut against a creature board.

### 6. Modern tempo, tournament

**Summary:** This blue-red tempo deck establishes early creature pressure, then uses card selection, removal, and countermagic to keep the opponent from stabilizing. It wins through sustained attacks backed by burn rather than a large finisher or combo. The flexible interaction supports a reactive game plan, but the deck gives up explosive closing power and can lose momentum when it draws too many lands or its early threats are repeatedly answered.

The quality model grades this deck typical against the top lists of the format: the cards pair in ways the top lists do not, and that lowers the grade. The deck runs its spells as playsets, and that lowers the grade. The deck runs few spells as one copy, and that raises the grade.

- PLAN plan_coherent=yes: Cheap threats like Ragavan, Ledger Shredder and Faerie Mastermind are backed by Bolt, Discharge, Counterspell and cantrips, which all serve one tempo plan.
- PLAN theme_fit=partly: It is a blue-red tempo deck for Modern with Ragavan and Ledger Shredder, but it has no Delver of Secrets or other flip creature, so the Delver part of the request is not met.
- PLAN useful_as_built=partly: The mana base of 24 lands with dual lands is sound and the curve is very low, but only 12 creatures and no reach beyond burn leave few threats to carry the game, and the deck counts 4 Mishra's Bauble among its spells.
- PLAN summary_honest=yes: The summary claims early creature pressure backed by burn and countermagic with no large finisher, and the list matches this with 12 creatures, 8 burn spells, 6 counters and no combo or big threat.

### 7. Modern burn, casual

**Summary:** This mono-red burn deck combines direct damage with a substantial creature plan. Apply pressure with early spells, remove blockers when necessary, and finish with burn or attacks from your larger threats. Draw support helps sustain pressure into the midgame. The deck trades the explosive speed of leaner burn builds for a stronger battlefield presence, so its opening turns are less aggressive and creature removal can disrupt its damage engines.

The quality model grades this deck typical against the top lists of the format: the cards pair in ways the top lists do not, and that lowers the grade. The deck runs few spells as one copy, and that raises the grade. Few of the cards are ones the top lists of the format play, and that lowers the grade.

- PLAN plan_coherent=partly: Burn spells, Torbran, Thermo-Alchemist and Sunspine Lynx push toward a burn-with-creatures plan, but Hazoret (which wants an empty hand) sits with four-mana Lambholt cards and card-draw spells that fill the hand, and so the pieces pull against each other a little.
- PLAN theme_fit=yes: Every spell is mono-red, the deck is built on burn spells with a creature plan, and it uses only Modern cards in line with a casual mono-red burn request.
- PLAN useful_as_built=partly: The deck can be played as it stands with 24 lands, a good number of burn spells and a mix of threats, but 24 lands is heavy for a curve topping at four mana, the weak Tectonic Giant and Lambholt cards dilute the burn, and the card draw is slow.
- PLAN summary_honest=yes: The summary claims a creature-heavy burn deck with Lightning Bolt, Boltwave and Stomp as damage, Risk Factor and Artist's Talent as draw, and larger threats like Tectonic Giant and Hazoret, and the list carries each of these and admits it is slower than lean burn.

### 8. Modern lifegain, FNM

**Summary:** This white-black lifegain deck develops a creature board, keeps cards flowing, and wins through steady combat pressure backed by larger midgame threats. Removal helps keep opposing creatures from taking over, while the sideboard offers disruption, protection, and board resets for different matchups. The deck favors a sustained creature game over fast combo finishes, giving up some speed and extensive disruption for a stronger battlefield presence.

The quality model grades this deck typical against the top lists of the format: the cards pair in ways the top lists do not, and that lowers the grade. The deck runs few spells as one copy, and that raises the grade. Many of the cards appear in no top list, and that lowers the grade.

- PLAN plan_coherent=yes: The deck pairs lifegain enablers like Guide of Souls, Ocelot Pride, Enduring Innocence and Elenda with lifelink or drain finishers like Archangel of Thune, Sheoldred and Gisela, all pointing at one creature-based lifegain plan.
- PLAN theme_fit=yes: It is a white-black deck built around lifegain payoffs and enablers in Modern, with a power level fitting FNM.
- PLAN useful_as_built=partly: The mana base of dual lands is solid for WB and the curve is low with plenty of threats, but 5 Swamp and 3 Plains plus 18 or so lands of heavy WW and BB costs is workable, while the deck has only a few removal spells and no sideboard, and several cards are odd for Modern, so it plays but with gaps.
- PLAN summary_honest=partly: The claim of a creature board with removal and larger midgame threats matches the list, but the summary mentions a sideboard offering disruption, protection and board resets while the listed cards show no sideboard, and removal is thin with only Solitude and Gumdrop Poisoner.

### 9. Standard midrange, FNM

**Summary:** This black-green midrange deck develops creature pressure, uses removal to keep opponents from taking over, and leans on draw to keep presenting threats in longer games. It wins through creature combat, with protection supporting its board and several creatures contributing to the draw or removal plan. It favors sustained pressure over explosive starts, so very fast opponents can force it onto the defensive before its midrange engine gets established.

The quality model grades this deck at the precon baseline against the top lists of the format: the cards pair in ways the top lists do not, and that lowers the grade. Few of the cards are ones the top lists of the format play, and that lowers the grade. The curve sits high for the format, and that lowers the grade.

- PLAN plan_coherent=partly: The black-green creature, removal and draw pieces broadly fit a midrange plan, but the Rooms, Diresight and Massacre Girl pull toward value and sweeping while the creature count is low and thin on pressure.
- PLAN theme_fit=yes: It is a black-green midrange deck in Standard with creatures, removal, and card draw, which is what the request asked for at FNM level.
- PLAN useful_as_built=partly: The mana base of 24 lands with dual lands is stable and the removal is good, but the curve is clunky with many three-drops, only 2 one-drops, and a thin set of ways to close out the game.
- PLAN summary_honest=partly: The summary claims creature pressure and a win through creature combat, but the list has only 14 creatures with just 2 Llanowar Elves as ramp, and several of them are low-impact or expensive like the 6-mana Rottenmouth Viper.

### 10. Standard aggro, tournament

**Summary:** Build an attacking board, use removal and interaction to keep the pressure on, and win through combat. Creature-based draw support helps sustain the assault when the first wave is not enough, while Stingerback Terror supplies a substantial finishing threat. The sideboard offers more disruption, removal, and staying power for slower matchups. This deck favors a direct creature-based game plan over late-game inevitability, so keep applying pressure rather than settling into a prolonged exchange of resources.

The quality model grades this deck at the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. The cards pair in ways the top lists do not, and that lowers the grade. Many of the cards appear in no top list, and that raises the grade.

- PLAN plan_coherent=partly: Cheap red creatures, Boros Charm and Case of the Gateway Express push toward aggro, but 4 copies of the six-mana Twinmaw Stormbrood and a few four-drops pull against a low curve.
- PLAN theme_fit=yes: The deck is a Standard red-white list built on Boros lands and cheap red and white creatures, which matches the request for red-white aggro.
- PLAN useful_as_built=partly: The 28-land mana base with Sacred Foundry, Inspiring Vantage and Sunbillow Verge is stable, but that many lands with only 32 spells and few creatures leaves the deck short on threats, and the six-drop Dragons slow an aggro plan.
- PLAN summary_honest=partly: The summary calls Stingerback Terror a finishing threat and the deck does run four, but it also says the sideboard offers disruption and removal when the list shows no sideboard, and it calls Reckless Lackey and Fugitive Codebreaker draw support when they are mostly cheap attackers.

### 11. Commander with a locked card

**Summary:** Karlov of the Ghost Council leads a sacrifice deck that develops expendable creatures, establishes value engines, and turns its board into sustained pressure. Use early ramp and draw to keep developing, hold interaction for threats to the engine, and commit larger creatures once you can profit from losing them. Win through accumulated sacrifice payoffs or creature combat, with several finishers providing a stronger late game. The deck gives up some immediate aggression for staying power, and hands full of payoffs without enough creatures can develop slowly.

The quality model grades this deck at the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. The mana base serves one color far better than another, and that lowers the grade. The cards pair in ways the top lists do not, and that lowers the grade.

- PLAN plan_coherent=yes: Nearly every card is a sacrifice outlet, a death-trigger payoff, a fodder creature, or a ramp and draw piece feeding them, so the list pulls toward one aristocrats plan around Karlov.
- PLAN theme_fit=yes: It is a Commander deck in white-black with Karlov as commander and a clear sacrifice theme, which is what was requested.
- PLAN useful_as_built=partly: The deck has ramp, outlets, payoffs and draw, but the 13 Swamps and 6 Plains alongside many dual lands leave white cards like Mondrak and Aven Interrupter shaky, and Damn and Toxic Deluge cut against a creature-based plan, so it plays but unevenly.
- PLAN summary_honest=partly: The summary's claim of sacrifice payoffs and several finishers checks out against Zulaport Cutthroat, Grave Pact, Dictate of Erebos, Razaketh and Meathook Massacre, but it says interaction is held up while the list has three sweepers and little cheap fodder, which it only hints at.

### 12. Commander on a budget

**Summary:** A budget-conscious, bracket-2 mono-white token deck built around Adeline's combat pressure. Repeatable token engines and anthem effects build a formidable army, while steady card advantage supports longer games. Flexible removal and protective interaction help preserve momentum, with evasive threats and an alternate victory route providing ways to close.

The quality model grades this deck below the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. Many of the cards appear in no top list, and that lowers the grade. The mana base serves the colors evenly, and that raises the grade.

- PLAN plan_coherent=yes: Nearly every card is a white token maker, anthem, token-doubler, or protection piece that supports Adeline's go-wide attack plan, with ramp and draw feeding that plan.
- PLAN theme_fit=yes: It is a mono-white token deck led by Adeline, Resplendent Cathar, with many token makers like Hero of Bladehold, Felidar Retreat, and Gideon, Ally of Zendikar, and its power looks consistent with bracket 2.
- PLAN useful_as_built=partly: The mana base of 37 lands plus many rocks is steady and the curve is playable, but running Hour of Reckoning, Austere Command, and Vanquish the Horde in a token deck, with some clunky top-end cards, makes it clumsier than it should be.
- PLAN summary_honest=partly: The claim of an alternate victory route checks out through Halo Fountain and Windbrisk Heights, but 'flexible removal' is overstated because several removal slots are weak or situational, and the summary skips the three costly wraths that cut against a go-wide plan.

### 13. owned first, and the commander is not owned

**Summary:** An Orzhov lifegain deck built around Karlov's growing combat presence and creature control. Small value creatures, Food, and lifelinking attackers sustain the engine, while efficient removal and protective interaction help maintain momentum. Angel-led attacks and dedicated late-game payoffs provide closing power without relying on an infinite combo.

The quality model grades this deck at the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. The cards pair in ways the top lists do not, and that lowers the grade. The mana base serves one color far better than another, and that lowers the grade.

- PLAN plan_coherent=partly: Lifegain pieces such as Angel of Vitality, Aerith, Exemplar of Light and Karlov point one way, but three board wipes, a heavy ramp suite and many loose instants pull away from a creature-based lifegain plan.
- PLAN theme_fit=yes: It is a Karlov-led Orzhov Commander deck with many lifegain payoffs and gainers, and the White-Black identity matches the request.
- PLAN useful_as_built=partly: It has plenty of ramp, removal and Angel finishers, but 14 Plains against 9 Swamps with many colorless-heavy utility lands and a thin Black presence make the mana uneven, and the amount of ramp with few lifegain engines leaves the deck short of a clear win path.
- PLAN summary_honest=partly: The summary claims Food and lifelinking attackers sustain the engine, but the list carries only Lembas as Food and few real lifelinkers, though the Angels it names (Lyra, Angel of Invention) are present.

### 14. the user delegates the commander

**Summary:** Atarka, World Render leads a Dragon deck that develops mana early, establishes a threatening board, and turns Dragon attacks into double-striking pressure. Lower-cost threats keep the deck active before its larger finishers arrive, while draw support and inexpensive interaction help sustain the attack. Lathliss, Dragon Queen, Utvara Hellkite, and Wrathful Red Dragon provide additional ways to close the game. The deck favors tribal combat over elaborate combo lines, giving up some speed and resilience in exchange for a strong Dragon theme and explosive attacking turns.

The quality model grades this deck at the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. The commander does not place in cEDH events, and that lowers the grade. The mana base serves one color far better than another, and that lowers the grade.

- PLAN plan_coherent=partly: Most cards are red Dragons or Dragon support, but the list adds off-plan pieces such as Faithless Looting, Pyroblast, Red Elemental Blast, Legolas's Quick Reflexes and three wipes that pull against a board of Dragons.
- PLAN theme_fit=yes: The deck is a Commander list led by Atarka, a Dragon commander in red and green, and it holds many Dragons that fit a tribal request.
- PLAN useful_as_built=partly: The deck has plenty of lands, ramp and Dragons to play, but it has heavy red costs against a mana base built from many Mountains, several painful lands and a few green sources, and the situational interaction like Pyroblast and Red Elemental Blast does little.
- PLAN summary_honest=partly: The summary says Atarka turns Dragon attacks into double-striking pressure, and Atarka is in the list as a 7-mana commander, but it hides that the deck is almost entirely red with a thin green side and carries wipes like Blasphemous Act.

### 15. delegated commander, owned first

**Summary:** Astarion pairs a lifegain plan with a creature-heavy board, using its end-step choice to grow your life cushion or press an opponent toward defeat. Angels and other threats supply the main attacking force, backed by dedicated finishers for longer games. Card draw, removal, and protective interaction help sustain that pressure. The deck favors steady development and board-based wins over fast combos, so rebuilding after repeated board wipes can take time.

The quality model grades this deck below the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. Many of the cards appear in no top list, and that lowers the grade. The commander does not place in cEDH events, and that lowers the grade.

- PLAN plan_coherent=partly: Lifegain creatures, Angels, and removal point the same way, but three board wipes and a pile of equipment and artifacts pull against the creature-heavy board the summary describes.
- PLAN theme_fit=yes: x
- PLAN useful_as_built=yes: x
- PLAN summary_honest=yes: x

### 16. a tight budget, owned first

**Summary:** Denethor leads a steady aristocrats deck that develops mana, builds a supply of creatures, and turns sacrifices into accumulating advantage. Keep the engine stocked with draw, use interaction to preserve important pieces, and finish through repeated life-loss pressure or a larger creature board. Board resets give you a way back when opponents pull ahead. The deck favors resilient, incremental progress over explosive turns, so it gives up some speed and depends on keeping creatures and payoff pieces available.

The quality model grades this deck at the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. The commander does not place in cEDH events, and that lowers the grade. Few decks lead with the commander, and that raises the grade.

- JUDGE [unknown]: "The commander does not place in cEDH events, and that lowers the grade.". This is a statement about the quality model's grading inputs, not a rule of the game, so it is not a rules claim. It is listed only for completeness.
- PLAN plan_coherent=yes: Nearly every card feeds one aristocrats plan of cheap creatures, sacrifice outlets and drain payoffs like Blood Artist, Zulaport Cutthroat and Falkenrath Noble, backed by Skullclamp, ramp and draw.
- PLAN theme_fit=partly: It is a white-black aristocrats Commander deck led by Denethor with a casual power level, but nothing in the list shows it was built from the person's own library first or kept under 25 dollars, since it holds Sol Ring, Swords to Plowshares and Marsh Flats.
- PLAN useful_as_built=yes: The deck has 39 lands plus about eleven mana rocks, plenty of cheap creatures, removal, draw and several drain finishers, so it plays smoothly as built.
- PLAN summary_honest=yes: The summary claims board wipes as a way back and the list carries Austere Command, Martial Coup and Dusk // Dawn, and it also admits the deck is slow and incremental, which fits the list.

### 17. upgrade a precon, owned first

**Summary:** Build a wide Goblin board, then use your commander to spread targeted spells across your creatures and sustain an explosive storm turn. The Goblin synergy package supplies steady pressure, while Great Train Heist and Assault on Osgiliath provide finishing plays. Early mana development helps establish the commander before committing to a big turn. The deck favors coordinated creature-and-spell turns over standalone threats, leaving it vulnerable when opponents repeatedly clear the board or disrupt its setup.

The quality model grades this deck typical against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. The mana base serves the colors evenly, and that raises the grade. The deck holds less interaction than the norm, and that lowers the grade.

- PLAN plan_coherent=partly: Most of the list is red Goblins, token makers, rituals and Impact Tremors/Grapeshot finishers that fit a Goblin storm plan, but Blasphemous Act, Vandalblast, and a pile of cheap cantrips and filler pull against a wide board.
- PLAN theme_fit=yes: It is a mono-red Zada, Hedron Grinder Goblin deck built around token makers, rituals and burn finishers, which matches a Goblin Storm upgrade at bracket 3.
- PLAN useful_as_built=partly: With 24 Mountains plus utility lands and Sol Ring, Arcane Signet and rituals the mana is steady and there are real win routes, but many weak cantrips and low-impact spells make the turns inconsistent.
- PLAN summary_honest=partly: The summary says Great Train Heist and Assault on Osgiliath are finishing plays and both are in the list, but it never mentions Impact Tremors, Grapeshot, or Goblin Bombardment, and the sweepers it hints at as vulnerability are actually in the deck.

### 18. upgrade a precon, any card

**Summary:** This five-color Turtle deck develops its mana, establishes a cast of Mutants, Ninjas, and Turtles, then turns successful attacks into growing creatures and fresh cards through its commander. Keep attacking to build an advantage, using removal and sweepers when opponents threaten to take over. Raphael, the Muscle, Dimension X Pizzasaur, and Everything Pizza supply finishing threats within the precon's character-driven plan. The deck favors sustained creature combat over compact combo wins, so it gives up some speed and remains dependent on maintaining a board and finding profitable attacks.

The quality model grades this deck at the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. Many of the cards appear in no top list, and that lowers the grade. The land count sits above the norm of the format, and that lowers the grade.

- PLAN plan_coherent=partly: The deck leans on a Turtle and Mutant creature plan with ramp and removal, but many cards are loosely related, such as Fast Forward, Game Over, Lessons from Life, and Electric Seaweed, and the five-color spread dilutes focus.
- PLAN theme_fit=yes: This is a five-color Heroes in a Half Shell deck packed with Turtle, Mutant, and Ninja characters and Turtle Power cards, which matches an upgrade of the precon at bracket 3.
- PLAN useful_as_built=partly: The deck has plenty of ramp, fixing, and creatures and can be played, but the many lands, clunky sweepers, and filler such as Coin of Mastery, Arcade Cabinet, and Exploding Barrel weaken it, and it has no clear way to close out games.
- PLAN summary_honest=partly: The summary names Raphael, Dimension X Pizzasaur, and Everything Pizza as finishing threats and they are in the list, but Everything Pizza is a 2-mana Food artifact and not a real finisher, and the claimed sweepers are only Blasphemous Act at {8}{R} and Vanquish the Horde at {6}{W}{W}.

### 19. the Hobbit family, two colours

**Summary:** Thranduil leads an Elf-focused midrange deck that develops mana, builds a creature presence, and uses legendary Elves to keep cards flowing. Removal and battlefield resets help it reach a longer game, where its Elf board and finishers—chief among them Witch-king of Angmar and Troll of Khazad-dûm—provide the closing pressure. It favors thematic creature development over explosive wins, while the basic-heavy mana base makes color-intensive draws less dependable.

The quality model grades this deck below the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. Many of the cards appear in no top list, and that lowers the grade. The cards pair in ways the top lists do not, and that lowers the grade.

- PLAN plan_coherent=partly: The Elf creatures, ramp, and card draw all point at a midrange Elf plan, but the black package of Witch-king, Troll, Languish, and Gnashing of Teeth wipes and finishers pulls away from the Elf board the green-blue core builds.
- PLAN theme_fit=yes: It is a three-color Thranduil, the Elvenking deck built entirely from Hobbit and Hobbit commander set cards with a heavy Elf count, which matches the request.
- PLAN useful_as_built=partly: It has plenty of ramp, draw, and removal and a curve that lands, but wipes sit awkwardly with an Elf creature deck, the finishers are few and clunky, and heavy BB, GG, and UU costs on 30 basics plus few fixing lands make it inconsistent.
- PLAN summary_honest=yes: The summary names Witch-king of Angmar and Troll of Khazad-dûm as finishers, and both are in the list tagged as wincons, and it admits the basic-heavy mana base (30 basics) makes color-intensive draws less dependable.

### 20. the Hobbit family, a delegated commander

**Summary:** Your commander leads a Dragon-focused deck that develops mana, builds up Treasure for powerful attacks, and uses removal and protective artifacts to keep its key creatures relevant. Dwarves, Goblins, and Equipment supply the supporting battlefield presence before Cavern-Hoard Dragon, Desert Were-Worm, and Smaug, the Great Calamity // Spew Flame close the game. The plan favors steady mana and combat over explosive combinations, but gives up a dense Dragon roster and plentiful card draw for its thematic supporting cast.

The quality model grades this deck at the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. The mana base serves the colors evenly, and that raises the grade. The cards pair in ways the top lists do not, and that lowers the grade.

- PLAN plan_coherent=partly: Treasure, Dwarves, Goblins, and Equipment all sit in mono-red and loosely feed a Smaug-led plan, but the list is a pile of Hobbit cards with only four Dragons and no clear engine tying the Equipment, tribes, and Treasure together.
- PLAN theme_fit=partly: It is a mono-red deck led by Smaug the Magnificent with every card from the Hobbit sets, which matches the request for a Hobbit commander, but the list has very few actual Dragons for a deck asked to be about dragons.
- PLAN useful_as_built=partly: The deck has a stable mono-red mana base of 38 lands plus ramp and enough removal, but with only a few Dragon finishers at six to nine mana and many low-impact Equipment and small creatures, it is playable yet slow and unfocused in winning.
- PLAN summary_honest=yes: The summary says the deck gives up a dense Dragon roster, and the list confirms it with only Smaug the Magnificent, Cavern-Hoard Dragon, Desert Were-Worm, and Smaug, the Great Calamity as Dragons, while the named Dwarf, Goblin, and Equipment support is plainly present.

### 21. the Hobbit family, mana from outside

**Summary:** Smaug leads a creature-heavy midrange deck that develops its mana, keeps cards flowing, and uses removal to make room for combat pressure. Goblins and thematic Equipment support the early and middle turns, while the Dragon finishers and Witch-king provide the closing threats. Smaug's noncombat-damage Treasure engine can help fuel the late game. The deck favors a broad Hobbit-themed battlefield over a compact combo finish, giving up some speed and protection in exchange for sustained creature pressure.

The quality model grades this deck at the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. Many of the cards appear in no top list, and that lowers the grade. The cards pair in ways the top lists do not, and that lowers the grade.

- PLAN plan_coherent=partly: Most of the list is Goblins, Equipment and Dragons in black and red, but a wipe package (Languish, Gnashing of Teeth, Desolation of Smaug) works against a creature-heavy pressure plan and there are many small cards with no shared payoff.
- PLAN theme_fit=yes: It is a black-red Smaug the Impenetrable deck built almost entirely from Hobbit-set cards, with lands and rocks from outside the set as the request allowed.
- PLAN useful_as_built=partly: The mana base of 22 basics plus fetches and duals is stable and there is ramp, draw and removal, but the ramp is thin and the cheap Goblins are low-impact, so the deck has few strong ways to close games.
- PLAN summary_honest=partly: placeholder

### 22. a set family and a card from outside it

**Summary:** An Elf-centered Sultai deck built around legendary-creature value, graveyard utility, and steady combat pressure. Mana acceleration and card advantage sustain development, while interaction and protective equipment support powerful evasive finishers. The supplied shortlist cannot satisfy the multicolor-land requirement under the singleton restriction, so full compliance is not possible.

The quality model grades this deck at the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. Many of the cards appear in no top list, and that lowers the grade. The cards pair in ways the top lists do not, and that lowers the grade.

- JUDGE [unknown]: "The supplied shortlist cannot satisfy the multicolor-land requirement under the singleton restriction, so full compliance is not possible.". The sentence mentions the singleton restriction, but it mainly reports a builder target (multicolor-land count) and the shortlist falling short. That is a builder limit, not a rule of the game. It does imply the singleton rule applies, which is true in Commander, but it states nothing false.
- PLAN plan_coherent=partly: Most cards support an Elf-tribal, legendary-creature value plan with ramp and card draw, but the equipment, Orc and Wraith finishers, and Languish sweepers pull in other directions and dilute the focus.
- PLAN theme_fit=yes: The deck is led by Thranduil, the Elvenking in his Sultai identity, draws almost entirely from the Hobbit sets, includes the requested Sol Ring, and leans into Elves as the commander suggests.
- PLAN useful_as_built=partly: With 29 lands plus plenty of ramp and draw it functions, but the mana base is mostly basics with few fixing lands for three colors, the threats are mostly small Elves, and the win conditions are modest.
- PLAN summary_honest=partly: The summary's Elf-centered Sultai plan with ramp, draw, interaction and protective equipment matches the list, but the claim of 'graveyard utility' is not visible among the cards, and the note about a multicolor-land shortfall is only loosely supported by the seven nonbasic lands.

### 23. two set families at once

**Summary:** Kíli the Resourceful turns a steady procession of Dwarves and Equipment into cards, while artifacts and legendary permanents build toward an enduring story. Develop mana early, establish an attacker, and use the free equip opportunity to keep combat pressure moving. The deck wins through its creature and Equipment package, with Angel of the Ruins providing a top-end finisher. Draw engines and protective interaction help sustain that pressure, but the deck relies on its commander and battlefield rather than a compact combo, so repeated board clears can slow it considerably.

The quality model grades this deck typical against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. The mana base serves the colors evenly, and that raises the grade. The commander does not place in cEDH events, and that lowers the grade.

This deck plan aims for 2 or more finishers, the cards that win the game, and the deck holds 1.

- PLAN plan_coherent=yes: Mono-white Dwarves, Equipment and artifacts with cheap ramp, draw, and equip-focused threats all feed one equipment-and-creature plan under Kíli.
- PLAN theme_fit=yes: Every card is from the Hobbit or Bloomburrow sets, the commander is a legendary Dwarf from The Hobbit chosen by the builder, and the deck is a plausible bracket 3 build.
- PLAN useful_as_built=partly: Thirty-one Plains plus ten ramp pieces make the mana smooth and the curve is low, but only one clear finisher, three wipes that hurt its own creatures, and several marginal cards leave the deck short of ways to close the game.
- PLAN summary_honest=partly: The claim that Angel of the Ruins is the top-end finisher checks out as a seven-mana artifact creature in the list, but the summary glosses over that the deck holds only one real finisher and that three board wipes sit awkwardly in a creature-based plan.

### 24. a 60-card deck from one set

**Summary:** This mono-red Bloomburrow aggro deck applies early creature pressure, with a Mouse-focused core and supporting attackers. Combat support and removal keep attacks productive, while draw spells help sustain momentum and larger threats provide a finish when games run longer. It favors consistent red mana and proactive play over defensive flexibility, leaving it vulnerable to repeated board clears and opponents that establish stronger late-game engines.

The quality model grades this deck at the precon baseline against the top lists of the format: the cards pair in ways the top lists do not, and that lowers the grade. Many of the cards appear in no top list, and that lowers the grade. Few of the cards are ones the top lists of the format play, and that lowers the grade.

- PLAN plan_coherent=yes: Nearly every card is a cheap red creature or a combat trick or burn spell that backs an early Mouse-led attack, so the cards all serve one aggro plan.
- PLAN theme_fit=yes: It is a mono-red aggro deck built entirely from Bloomburrow cards, matching the request for a casual Modern deck.
- PLAN useful_as_built=yes: Twenty-four lands with a low curve topping out at two five-drops, mono-red mana, and a real mix of threats, tricks and removal make it playable as it stands, though 24 lands is a touch heavy for such a low curve.
- PLAN summary_honest=yes: The summary claims a Mouse-focused core with early pressure, and the list carries 12 Mice (Heartfire Hero, Manifold Mouse, Emberheart Challenger, Roughshod Duo) plus Hired Claw and removal as described.

### 25. use no card of an owned precon

**Summary:** Captain America leads a white Avengers team that develops its mana, keeps cards flowing, and pressures opponents with Heroes. Protect Captain America, maintain the team's presence, and close with finishers, chief among them The Sentry, Golden Guardian. The deck favors sustained creature combat over explosive combos and trades some superhero flavor for practical support. The Avengers Assemble precon inventory was not supplied, so exclusion of its cards cannot be confirmed.

The quality model grades this deck at the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. The mana base serves the colors evenly, and that raises the grade. The commander does not place in cEDH events, and that lowers the grade.

- PLAN plan_coherent=partly: Mono-white Heroes with removal, protection and ramp points one way, but the heavy artifact ramp, three wraths and a long list of one-shot instants pull against a creature-combat plan.
- PLAN theme_fit=yes: It is a mono-white Captain America commander deck full of Avengers and Marvel Hero creatures, matching the superhero theme, and the builder chose the commander as asked.
- PLAN useful_as_built=partly: With 35 lands, 10 mana rocks and cheap interaction it will function, but 31 Plains plus a few utility lands leaves limited card advantage and few ways to win, and the three wraths work against a creature plan.
- PLAN summary_honest=partly: The claim that The Sentry, Golden Guardian is a chief finisher is shaky because it is a 4-mana Human Hero in a list that otherwise has few real finishers, though the summary does admit the precon exclusion could not be verified.

