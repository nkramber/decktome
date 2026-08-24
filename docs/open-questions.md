# Open questions

Questions not yet asked, or asked and not yet answered. Move each answered question to `decisions.md`. Ask in small batches when the design reaches the point that needs the answer.

## Not yet asked

| # | Question | Why it matters | Ask when |
|---|---|---|---|
| OQ-1 | Deck export targets: ManaBox text, Moxfield, Arena, plain text, PDF? | Sets the export module scope. | Before the UI section is final. |
| OQ-2 | Should the app store the user's full collection, or only a hash and a summary? | Privacy, storage cost, and re-import flow. | Before the data model section. |
| OQ-3 | Price data: show prices (Scryfall gives daily USD/EUR/TIX)? Which currency? | Budget questions need it. | Before the acquisition list design. |
| OQ-4 | How much variance between two builds of the same prompt? A seed the user can reuse? | The owner asked for variance. A seed makes it reproducible. | Before the generation section. |
| OQ-5 | Should the agent explain each card choice (one line per card) in the deck view? | UX and token cost. | Before the UI section. |
| OQ-6 | Does the user ever play the deck inside the app (goldfish test, sample hands)? | A sample-hand simulator is cheap and useful. Scope decision. | Phase 2 planning. |
| OQ-7 | Which LLM providers must the role layer support on day one? | Adapter count. | Before the provider section is final. |
| OQ-8 | Judge model for eval: same provider or a different one? | Bake-off independence. | Eval section. |
| OQ-9 | Multi-language card names (ManaBox `Language` column)? | Non-English collections need name mapping through Scryfall ID. | Import section. |
| OQ-10 | Terms-of-use check on MTGGoldfish, MTGTop8, Aetherhub, EDHREC: who does it and when? | D-5 allows them with a legal check. | Before the meta ingester ships. |
| OQ-11 | Domain name and GCP project ids (dev, prod)? | Environment config. | Phase 0. |
| OQ-12 | Should the agent's questions come from a fixed catalog only, or can the model invent questions? | Predictability versus flexibility. | Agent section. |

## Asked, waiting

(none)
