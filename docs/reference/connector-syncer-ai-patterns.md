# Reference: how connector-syncer uses AI

Status: research note. Verified 2026-08-23 against the local checkout at `/Users/nate/Repos/connector-syncer`. Line numbers can move. Check them before you cite them in the design doc.

Purpose: connector-syncer (Wallabee) is a Python RAG app for engineering documents. This note lists the AI mechanisms it uses and the lessons we take into the deck builder. It is not a full architecture description.

## 1. Provider layer

The app has one provider seam: `libs/clients/openai_wrapper.py`. Two other wrappers copy its signature: `libs/clients/anthropic_wrapper.py` and `libs/clients/openrouter_wrapper.py`. The Anthropic wrapper gets structured output through one forced tool call. It validates the tool input with the same pydantic model. Call sites see one interface (`.output_parsed`) and do not know the provider.

Key mechanisms:

| Mechanism | Where | What it does |
|---|---|---|
| Request budget | `openai_wrapper.py:49-114` | One object holds max attempts and an absolute deadline. Every fallback candidate shares it. The SDK's hidden retries are off. |
| Three retry classes | `openai_wrapper.py:412-516` | Truncation: retry once at once, with a higher token cap. Transient: exponential backoff. All other errors: raise now. |
| Escalation bound | `openai_wrapper.py:117-122` | The higher cap is `min(model max, max(8 x caller cap, 8192))`. A small call can not become a huge call. |
| None-guard | `document_summarizer.py:180-192` | `output_parsed` is None on refusal or incomplete output. A named guard turns this into a typed failure. |
| Reasoning effort | `openai_wrapper.py:343` | The `reasoning` field is sent only for models in an explicit set. |

Lesson for us: classify each error by one question. "Can a retry help?" Truncation, rate limit, and schema errors need three different answers.

## 2. Role-to-model configuration

`libs/models/llm_roles.py` is a frozen pydantic model. It has one field per call-site role (`routing_model`, `summarization_model`, `synthesis_model`, and more). Each field has a default `ModelId`. Each field is overridable by an environment variable.

The module docstring records the A/B evidence behind each default.

`libs/models/model_routing.py` holds the registries: model ids, reasoning-effort set, max output tokens per model, tiers, and fallback chains. `libs/services/model_orchestration_service.py:225-229` falls back to the next model only on availability errors. A schema or parse error raises. It does not hide behind another model.

`scripts/check_model_default_changes.py` runs in CI. It diffs model defaults between the merge base and the PR head. It warns and never fails. `docs/model-bakeoffs.md` gives the staged ladder: freeze prompts and datasets, then compare models only.

Lesson for us: the owner chose a provider-agnostic role layer. This file is the template. Put no model string at a call site.

## 3. Token and cost accounting

`libs/utils/llm_usage.py:67-122` reads usage from three response shapes (OpenAI Responses, OpenAI Chat, Anthropic Messages). It never raises.

`LLMUsageAccumulator` (125-209) is one object per user flow. It reports `None` for every field when nothing was recorded. A null row means "not instrumented". A zero row means "measured zero". These are different facts.

`libs/models/product_analytics_event.py:38-51` rejects raw `query`, `answer`, and `email` keys at construction time. Raw content goes to a separate short-retention table.

Lesson for us: instrument cost on day one. The roadmap doc says every cost claim without M-1 was an estimate.

## 4. Prompt construction

Prompts live in three tiers:
1. Module constants next to the consumer (for example `libs/models/llm_query_response.py:197`).
2. A runtime control plane: `libs/models/prompt_config.py` gives versioned prompts per environment, with `change_reason` and `eval_status`.
3. A fallback registry: `libs/services/prompt_default_registry.py`.

Safety mechanisms:
- `rag_service.py:66-73` HTML-escapes every interpolated user or document string. This blocks prompt injection through XML-tagged sections.
- `generic_search_agent.py:583-614` wraps each retrieved chunk in an XML tag with ids. Plain-text labels leaked back into answers as fake citations.
- `document_summarizer.py:195-205` serializes pydantic state to JSON before it goes into a prompt. One site used a raw Python repr. It was a defect (Q2 in the roadmap).
- `rag_service.py:945-1005` drops model citations that point at chunks that do not exist. Then it remaps every dependent index.
- `libs/models/term_utils.py:20-218` defines one canonical "Not Found" value and careful matching for absence phrases.

Lesson for us: our equivalent of a hallucinated citation is a card name that does not exist, or a card that is not legal. The normalizer must check every card against the card database after every LLM call.

## 5. Agentic patterns

Two architectures exist in `libs/agents/`.

A. Config-driven search agent (`generic_search_agent.py`). An `AgentConfig` holds the system prompt, templates, a response type, `max_tool_calls`, and a `ModelConfig`. The `ModelConfig` separates a cheap search model from a strong analysis model. 

`SearchState` (151-237) holds the multi-turn state: queries, chunks keyed by id, thoughts, messages, and token counters. 

Tool definitions (365-477) are strict JSON Schema with `maxLength`, `enum`, and `pattern` constraints. The loop (804-922) caps tool calls. It feeds malformed tool JSON back to the model as an error string. It adds a nudge when the model repeats a query.

B. Repository-tool agents (`app_intelligence_agent.py`, `schema_intelligence_agent.py`). Each tool requires an `agent_thought` string. The UI shows it as a progress step. `_get_tools_for_mode` (505-530) gives the model only the tools that are legal for the classified mode. 

`run()` is a generator that yields SSE events. Output has three fallback levels: strict parse, then JSON-repair reprompt, then a deterministic answer from repository data with no LLM. The schema agent keeps a guided question workflow across turns (1903-1962). It records each answer and summarizes prior Q&A so the next turn continues.

Other mechanisms: `DeepSearchBudget` (43-83) is one deadline and one cancel event shared by all phases. A heartbeat thread (977-1031) re-emits a progress event every 10 seconds during a long call. `DeepSearchCheckpoint` (86-110) saves partial results after each tool call.

Lesson for us: the deck builder is a turn-based question workflow. The schema agent's "record the answer, summarize, continue" pattern is the direct model. The mode-gated tools pattern maps to "format is known, so only expose format-legal tools".

## 6. Embeddings and retrieval

`libs/clients/pinecone_wrapper.py` uses a strategy object per tenant type. Namespace is the account id. Every filter includes a run id. Dense and sparse lanes are fused with Reciprocal Rank Fusion (`chunk_retriever_service.py:3646-3728`, k=60). 

The sparse lane is optional. Its failure does not fail the search. The reranker payload has two funnels: max 100 candidates with a per-lane floor of 20, then a 200k character budget.

`libs/models/chunk.py:40-74` marks retrieval-only fields `Field(exclude=True)`. They can not leak into storage or the API.

Lesson for us: a card database of about 35,000 Oracle cards is small. We may not need a vector store for card search at first. Scryfall Oracle tags and keyword catalogs give a structured theme index. Decide this in the design doc (open question).

## 7. Patterns we take

1. One request budget per logical call, shared across fallbacks.
2. Three retry classes: truncation, transient, terminal.
3. Fallback to another model only on availability errors.
4. One frozen role-to-model object. No model strings at call sites. Defaults carry their evidence.
5. CI warns on model-default changes. A written bake-off protocol gates the swap.
6. Usage accounting that reports null, not zero, when not instrumented.
7. Strict JSON Schema on tool arguments and on final output. An `agent_thought` on every tool.
8. Malformed tool calls go back to the model as tool output.
9. Three-level output degradation, with a no-LLM deterministic path tried first.
10. Normalize model output against ground truth after every call.
11. Escape everything you interpolate. Exclude internal fields at the model boundary.
12. Heartbeats for long calls. Checkpoints for long flows.
13. Small explicit loop caps (5 to 8 tool calls).
14. Cross-turn question workflows: record, summarize, continue.
