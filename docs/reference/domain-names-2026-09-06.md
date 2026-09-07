# Product names and domains, checked on 2026-09-06

The owner asked for candidate URLs for the product on 2026-09-06, with the availability checked before any recommendation. This note holds the method, the constraints, every name checked, and the recommendation. The app calls itself "MtG Deck Builder" today, in `web/apps/web/index.html` and the shell. The owner read this note and chose decktome on 2026-09-06, and bought `decktome.com` at GoDaddy the same day (D-556). The rest of this note is the record of the search.

## 1. Method

Each check reads the RDAP record of the registry itself, which is the data every registrar search box queries. A record means someone holds the name, and an HTTP 404 means the name is free. The calibration used `google.com` and `google.app`, which both answered a record, and a nonsense name, which answered 404.

| Ending | Registry and RDAP base | Source |
|---|---|---|
| .com | Verisign, `https://rdap.verisign.com/com/v1/domain/` | IANA bootstrap, read 2026-09-06 |
| .app | Google Registry, `https://pubapi.registry.google/rdap/domain/` | IANA bootstrap, read 2026-09-06 |
| .cards, .gg | CentralNic, `https://rdap.centralnic.com/<tld>/domain/` | `https://data.iana.org/rdap/dns.json`, read 2026-09-06 |

CAUTION: a free name today is no reservation. Other people registered seven names of this family between 2025-12 and 2026-07, several on Vercel DNS, so someone else builds in this space. Buy the chosen name on the day of the choice.

CAUTION: a registry check is not a trademark check. Search the USPTO records by hand before the name goes public.

## 2. Constraints

- No Wizards of the Coast trademark in the name. The Fan Content Policy (last updated 2017-11-15, read 2026-09-06) says: "You may not incorporate any Wizards of the Coast logos and trademarks in your Fan Content without our prior, written consent." So no "Magic", no "MTG", no "Planeswalker", and no "Commander" in the name.
- No "Deckmaster" either. Deckmaster was the brand Wizards of the Coast put on the back of every card of its trading card games. Every Magic card back carries it today (MTG Wiki, "Card back", read 2026-09-06). `deckmaster.com` has a holder since 1996-05-22.
- The same policy says fan content "must be free for others (including Wizards) to view, access, share, and use without paying you anything". The invite-only app is free today. A paid tier later changes the legal ground, so the name must not depend on the policy.
- No name of an existing tool. Moxfield, Archidekt, TappedOut, Deckstats, Scryfall, EDHREC, ManaBox, Deckbox, TopDecked, Cube Cobra, Aetherhub, and Commander Spellbook belong to others.
- Both .com and .app free. People type .com, and .app fits a web app that installs on a phone (PR-25).

## 3. Prices

Porkbun, read 2026-09-06. Cloud Domains and Squarespace charge in the same range (`docs/setup-gcp.md`, section 3).

| Ending | First year | Each year after |
|---|---|---|
| .com | $11.08 | $11.08 |
| .app | $8.75 | $14.93 |
| .cards | $2.57 | $31.41 |
| .gg | $51.80 | not shown on the page |

The .app ending sits on the HSTS preload list, so every browser connects to it over HTTPS alone (registry.google, read 2026-09-06). Firebase Hosting issues the certificate, so this costs nothing.

## 4. Every name checked

Free means HTTP 404 from the registry on 2026-09-06. Taken names carry their registration date when the record showed one.

| Name | .com | .app | .cards | .gg |
|---|---|---|---|---|
| binderscribe | free | free | not checked | not checked |
| binderquill | free | free | not checked | not checked |
| binderfirst | free | free | not checked | not checked |
| bindersmith | free | free | not checked | not checked |
| bindersage | free | free | not checked | not checked |
| binderseer | free | free | not checked | not checked |
| bindermage | free | free | not checked | not checked |
| binderlore | free | free | not checked | not checked |
| bindertome | free | free | not checked | not checked |
| binderbrewer | free | free | not checked | not checked |
| binderwarden | free | free | not checked | not checked |
| brewwarden | free | free | not checked | not checked |
| brewherald | free | free | not checked | not checked |
| openinghand | taken 2025-06-24, pending delete | free | not checked | not checked |
| curvebrew | free | free | free | free |
| brewquill | free | free | free | free |
| binderpilot | free | free | free | free |
| binderwright | free | free | free | free |
| decktome | free | free | free | free |
| binderdeck | free | free | free | free |
| curvewright | free | free | not checked | not checked |
| ownedbrew | free | free | not checked | not checked |
| ownedfirst | free | free | not checked | not checked |
| binderbrew | taken 2026-05-15 | free | not checked | not checked |
| binderforge | taken 2024-10-03 | free | | |
| brewwright | taken 2025-10-30 | free | | |
| brewtutor | taken 2011-06-16 | free | | |
| brewcurve | taken 2023-08-05 | free | | |
| havebrew | taken 2023-06-13 | free | | |
| yourbinder | taken 2022-07-12, in redemption | free | | |
| brewlist | taken 2006-01-09 | free | | |
| curvecraft | taken 2010-03-05 | free | | |
| brewloom | taken 2025-07-13, in redemption | free | | |
| brewfoundry | taken 2012-11-19 | free | | |
| deckfoundry | taken 2012-07-12 | free | | |
| binderfoundry | taken 2026-05-13 | free | | |
| brewkeeper | taken 1997-01-06 | free | | |
| binderbuilt | taken 2017-09-24 | free | | |
| binderstack | taken 2025-08-30 | free | | |
| deckquill | taken 2026-02-05 | free | | |
| brewcodex | taken 2026-03-01 | free | | |
| brewtome | taken 2020-08-12 | free | | |
| binderfolio | taken 2026-03-24 | free | | |
| deckwright | taken 2000-09-19 | taken 2026-07-05 | | |
| brewbinder | taken 2025-08-28 | taken 2026-07-21 | | |
| brewsmith | taken 1999-10-06 | taken 2020-08-29 | | |
| deckpilot | taken 2010-03-27 | taken 2026-03-04 | | |
| decktutor | taken 2009-10-19 | taken 2026-02-04 | | |
| manabinder | taken 2016-07-26 | taken 2020-04-16 | | |
| deckloom | taken 2025-12-05 | taken 2026-04-28 | | |
| brewstack | taken 2012-04-19 | taken 2025-12-02 | | |
| deckcodex | taken 2025-05-03 | taken 2025-05-03 | | |
| deckfolio | taken 2011-01-08 | taken 2025-10-06 | | |
| brewscribe | taken 2023-05-02 | free | | |
| deckscribe | taken 2025-11-10 | free | | |
| manascribe | taken 2023-05-28 | free | | |
| spellscribe | taken 2025-05-14 | free | | |
| deckwarden | taken 2025-12-28 | free | | |
| manaquill | taken 2026-06-10 | free | | |
| onthecurve | taken 2003-07-14 | free | | |
| brewbound | taken 2010-01-21 | free | | |
| deckbound | taken 2014-02-25 | free | | |
| manabound | taken 2018-06-19 | free | | |
| brewsheet | taken 2011-06-14 | free | | |
| curvelist | taken 2022-01-05 | free | | |
| brewsage | taken 2015-03-22 | taken 2026-03-01 | | |
| decksage | taken 2025-03-22 | taken 2025-10-15 | | |
| brewseer | taken 2023-08-15 | free | | |
| deckseer | taken 2025-05-13 | taken 2026-06-23 | | |
| deckie | taken 2012-05-07 | taken 2026-06-23 | free | free |
| decki | taken 2000-07-19 | taken 2025-08-01 | free | free |
| deckly | taken 2009-01-15 | taken 2024-12-17 | | |
| deckee | taken 2013-07-08 | taken 2018-05-08 | | |
| deckio | taken 2015-12-25 | taken 2025-12-23 | | |
| deckit | taken 2001-06-11 | taken 2025-12-06 | | |
| deckmaster | taken 1996-05-22 | not checked, a Wizards brand | | |
| decksmith | taken 1996-03-27 | not checked | | |
| brewpilot | taken 2019-12-22 | not checked | | |
| cardpool | taken 2003-05-29 | not checked | | |
| deckmate | taken 2000-02-27 | not checked | | |
| brewmate | taken 2005-11-21 | not checked | | |
| thebinder | taken 2010-07-17 | not checked | | |
| sleeved | taken 2003-04-04 | not checked | | |

A web search on 2026-09-06 found no product or company named curvebrew, brewquill, binderpilot, binderwright, or decktome. "Binder deck" reads as a physical product, and an iOS app named "Card Binder: MTG Manager" exists, so binderdeck is the weakest of the six.

The second round of 2026-09-06 checked 30 more names on .com and .app. Every "deck" compound of that round belongs to someone else, most since 2025. The "binder" family is open: 13 of its names are free on both endings. A web search found no product named binderscribe, binderquill, or binderfirst.

Four names have a soft collision. Mage Hand Press sells a "Binder" class for D&D, so bindermage reads as theirs. "Binder lore" is a D&D term. A developer tool named Warden installs with `brew install warden`, so brewwarden reads as that. Old Herald Brewery in Illinois sits near brewherald.

`openinghand.com` is in the pending-delete state on 2026-09-06, so the registry releases it within about five days of that state. A drop-catch service takes an order for it, and `openinghand.app` is free now. The opening hand is the first seven cards of a game, and every player knows the phrase.

The owner asked for the short forms deckie, decki, and deckmaster on 2026-09-06. Every short "deck" form has a holder on .com and on .app: deckie, decki, deckly, deckee, deckio, and deckit. Four of them went on .app in 2025 or 2026. The .gg and .cards endings of deckie and decki are free.

Two collisions sit beside them. Decky Loader is the plugin loader of the Steam Deck, and it sounds the same as deckie. DECKEE is a boating app. Deckmaster is out as a Wizards brand, and section 2 says why.

Five short dotted forms on 2026-09-06: `binder.cards`, `brew.cards`, `curve.cards`, and `binder.build` answered 404, and `brew.build` belongs to someone since 2024-12-16. CAUTION: a registry answers 404 for a premium name too. A registrar then asks a premium price for it, often hundreds of dollars a year. Read the price on the registrar page before you count one of these as free.

## 5. Names with their own identity

The owner asked on 2026-09-06 for names that stray from the "deck builder" theme and build an identity of their own. Moxfield, Archidekt, and Scryfall are names of that kind. Three rounds checked 60 such names on .com and .app, of five kinds:

- slang that doubles as a brand,
- invented fantasy words that fit the gold and the Cinzel face,
- rare English words with the right texture,
- divination words,
- coinages with the -mancer and -wright endings.

The read: the shelf of real fantasy words is empty, and it empties fast. `gildwright.com` went on 2026-08-30 and `hollowmere.app` on 2026-09-05. Every plain English word and every common compound of this kind belongs to someone. The names that survive are coinages, and the root "sigil" is the one that stays free. A sigil is a magical sign, and the five mana symbols are the sigils of the game.

Free on both .com and .app on 2026-09-06, with the web search of the same day:

| Name | Reads as | Search |
|---|---|---|
| sigilwright | a maker of sigils | clean, one name-generator hit |
| sigilhold | a keep of sigils | clean |
| sigilfold | a fold of sigils | clean |
| sleevewright | a maker of sleeved decks | clean |
| sleevefold | | clean |
| sigilary | the adjective of sigil, spelled with one l | a story title and a handle carry it |
| curvefold | | "curved folding" is an origami term |
| quillwyrm | a wyrm with a quill | a developer uses it as a handle on GitHub and YouTube |
| curvemancer | a wizard of the curve | taken as a name: CurveMancer is a Unity spline tool |

Free on .app alone: keepseven, grimalkin, runewell, glyphwright, gildwright, lanternfall, curveout, brewhold, quillhold, curvery, sigilfold, sigilhold, sigilry, cartomancy, brewmancer, curvemancer, sortileger, augurist, tapwright, sleevewright, sleevefold, curvewell, curvestone, sigilary, brewvane, cardessa. `lanternfall.com` is in its redemption period, so a return to the pool is possible.

Taken on both endings, with the .com date: snapkeep (2013), seneschal (1995), inkwyrm (2017), bookwyrm (1998), emberwell (2020), wardstone (1999), loreweave (2026-03-02), spellweave (2019), palimpsest (1999), hollowmere (2025-04-06), lorekeep (2014), tutorium (1998), manafold (2021), brewmark (2006), emberfold (2023), inkfold (2011), quillmark (2002), curvia (2020), quillory (2025-12-15), lorium (2014), cartomancer (2014), sortilege (2008), chartula (2000), tapstone (2001), deckora (2009), sigilla (2007). Taken on .com alone: keepseven (2025-09-27), grimalkin (1995), runewell (2011), glyphwright (2024-12-03), gildwright (2026-08-30), lanternfall (2024), curveout (2016), brewhold (2025-10-27), quillhold (2026-06-27), curvery (2020), brassquill (2026-01-28), lorefold (2008), curvium (2019), tomery (2019), cartomancy (2001), brewmancer (2021), sortileger (2005), augurist (1999), tapwright (2026-02-22), curvewell (2014), curvestone (2019), sigilry (2017), brewvane (2023), cardessa (2014).

The identity pick: **sigilwright**. It is a coined word with a clear reading, a maker of sigils, and it looks right in capitals in the Cinzel face. Behind it, **sigilhold** as a name for a keep, and **sleevewright** for a maker of sleeved decks. The rest are weaker: quillwyrm is a person's handle, curvemancer is a product, and sigilary, curvefold, and sleevefold say little.

## 6. The recommendation

Thirty-one names are free on both .com and .app on 2026-09-06. Section 4 holds 22 that describe the product, and section 5 holds 9 with an identity of their own. These are the ones to examine, in the order of this reader's preference.

1. **curvebrew**. "Curve" is the mana curve and "brew" is a deck you build yourself. Both words belong to every Magic player and to no company. It is two syllables, it spells as it sounds, and all four endings are free.
2. **sigilwright**, the identity pick of section 5. It describes nothing and belongs to nobody, which is what a brand wants. The mana symbols are the sigils of the game.
3. **binderscribe**. A scribe writes for you, and the binder is what the product reads first. It names the agent and the owned-first rule in one word, and it fits the quill of the look.
4. **binderquill** and **brewquill**. Both fit the look of the app: the Cinzel display face, the gold, and the grimoire texture. Brewquill says less about decks to a stranger, and binderquill says more.
5. **binderfirst**. It is the pool rule of the product as a name. It is plain and honest, and it reads as a policy more than a brand.
6. **binderpilot**. It says what the product does: an agent that flies your binder. "Pilot" is the word every AI product uses in 2026, which dates it.
7. **bindersmith**, **binderwright**, and **curvewright**. A smith or a wright is a maker, as in deckwright, which someone else holds. All three read well and need a spelling on the phone.
8. **openinghand**, the wildcard. The phrase is the best of the list, and the .com waits on a drop catch.
9. **decktome**, **bindertome**, **binderlore**, **bindersage**, **binderseer**, **binderwarden**, **binderbrewer**, **binderdeck**, **ownedbrew**, **ownedfirst**, **brewwarden**, and **brewherald**. Free on both endings, and each one is weaker than the names above, by section 4 or by length.

Buy the .com and the .app of the chosen name together, about $26 the first year and about $26 a year after that. Point the .app at the .com, or the other way. Leave .cards and .gg alone unless a squatter matters: $31.41 and $51.80 a year buy nothing the product needs.

## 7. Sources

- Verisign RDAP, `https://rdap.verisign.com/com/v1/domain/`, read 2026-09-06.
- Google Registry RDAP, `https://pubapi.registry.google/rdap/domain/`, read 2026-09-06.
- CentralNic RDAP through the IANA bootstrap file, `https://data.iana.org/rdap/dns.json`, read 2026-09-06.
- Porkbun price pages, `https://porkbun.com/tld/com`, `/app`, `/cards`, `/gg`, read 2026-09-06.
- Google Registry, `https://www.registry.google/domains/app/`, the HSTS preload statement, read 2026-09-06.
- Wizards of the Coast Fan Content Policy, `https://company.wizards.com/en/legal/fancontentpolicy`, last updated 2017-11-15, read 2026-09-06.
