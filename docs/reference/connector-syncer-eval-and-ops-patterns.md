# Reference: evaluation and operations patterns in connector-syncer

Status: research note. Verified 2026-08-23 against the local checkout. Line numbers can move. This note lists what we can reuse for the deck builder's quality gates and operations.

## 1. The eval service (`apps/sage`)

Sage is a nightly batch job, not an API. Its data model has one fact table per domain. Every report is a view. `evals.eval_runs` has one row per run (judge model, prompt version, config, code version). `evals.eval_results` is long format: one row per item and metric. A new metric adds rows, not columns.

Five suites exist. Three are "golden lanes" with hand-verified reference data. One is reference-free LLM judging of live traffic. One is a "dead role" tripwire that checks each small LLM role still works.

## 2. Mechanisms to copy

| Mechanism | Where | Why it matters for us |
|---|---|---|
| Corpus fingerprint | `apps/sage/golden_evals/compare.py:66-91` | Each run records dataset version, environment, model, effort, schema version. Baseline selection filters on these. A run from a different epoch is never compared. |
| Record resolved values | `apps/sage/golden_evals/bod_answerer.py:59-89` | The model id in the eval row is the one that ran, not the one requested. A moved default breaks the epoch instead of shifting the series. |
| Metric suffixes | `apps/sage/golden_evals/scorer.py:53-75` | `_info` and `_lenient` rows are reported but never gate. |
| Observe-only is not pass | `compare.py:55-64` | A suite without gate metrics renders a different headline. An empty failure list means "not evaluated". |
| Named flips | `compare.py:328-380` | The report names the exact items that regressed, not only the aggregate delta. |
| Dead-instrument tripwire | `apps/sage/golden_roles/runtime.py:172-199` | Each probe emits `<role>_llm_invoked`. A role without this row is dead, even if the metric is green. |
| Hallucination traps | `apps/sage/docs/golden-evals-design.md:63-79` | Facts known to be absent are free deterministic tests. Any confident answer is a fabrication. |
| Pure lib + thin CLI | `apps/sage/golden_evals/rerank_ab.py` and `scripts/eval/golden_rerank_ab.py` | Deterministic scoring in a unit-tested package. I/O and cloud clients in the script. Resumable arms keyed by (item, arm). |
| Label-gated sweep | `.github/workflows/golden-eval.yml` | An expensive eval runs only on a PR label. A missing corpus skips green with a notice. A broken corpus fails loudly. |
| Four-tier ladder | `apps/sage/docs/golden-evals-design.md:1-60` | Tier 0: CI, $0, determinism only. Tier 1: nightly on prod. Tier 2: offline sweep on a seeded corpus. Tier 3: full E2E. Same dataset, same runner, same metric names. |
| Bake-off protocol | `docs/model-bakeoffs.md` | Freeze prompts, dataset, thresholds, and instrumentation. Compare models only. Cost cap $5. Verdicts: PASS, FAIL, INCONCLUSIVE. |

## 3. Deck-builder equivalents

The deck builder has cheaper deterministic tests than a document RAG app. We can check most quality claims without a judge model:
- Legality: every card legal in the format on the query date. Deterministic from Scryfall data.
- Deck size and copy limits: deterministic.
- Color identity (Commander): deterministic.
- Ownership: every card in the collection, or marked as an acquisition. Deterministic.
- Mana curve, land count, color sources: deterministic heuristics with published ranges.
- Theme fit: partially deterministic through Scryfall Oracle tags and keywords.
- Usefulness and variance: needs a judge or a human. This is the hard part.

Hallucination traps map directly: a prompt for a card that does not exist, a banned card, or an off-color card. Any deck that includes it is a fabrication. The name-check after every LLM call is our "Not Found" guard.

## 4. Operations patterns

- Flash (API) and Sloth (worker) split. Both dispatch Cloud Tasks. `libs/clients/tasks.py:13-40` gives an application lease that outlives the dispatch deadline by a fixed grace. Task names are deterministic so racing producers collide by name.
- Queue retry budgets are provisioned by a dry-run-by-default script (`scripts/provision_sloth_task_queues.sh`). It prints the current state first and refuses prod without an explicit flag.
- Alert policies ship in shadow mode with no notifications until history exists (`scripts/provision_sloth_parse_alerts.py`).
- Analytics event types are a registry with a validator (`libs/models/product_analytics_event.py:18-51`). Prohibited keys (raw query, answer, email) are rejected at construction.
- The E2E project allowlist is a code constant, not a settings field. Environment-overridable settings are a fail-open path (`libs/config/e2e_allowlist.py`).
- OTEL collector turns spans into RED metrics with low-cardinality dimensions only (`otel-collector-config.yaml`).
- The MCP server exposes six Cloud Run log tools and denies prod access by default.

## 5. Document style

The roadmap style has nine parts: a status header with verification dates, a thesis, lessons, a system map, and a cost model. Then a defect register with a status legend and numbered guardrails. Then a phased plan with a gate and a plain-English paragraph on each entry. Then strict-order sequencing and open questions. Refutations are recorded, not deleted. The `design-doc-style` skill holds the template.
