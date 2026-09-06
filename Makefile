# mtg-deck-builder - single human entry point.
# Every target prints what it does. Versions are pinned in go/go.mod, web/package.json, and here.

# Every recipe runs under bash with pipefail, so a command that fails
# inside a pipe (for example before a tee) fails the target.
SHELL := bash
.SHELLFLAGS := -o pipefail -c

GOLANGCI_LINT_VERSION := v2.13.2
GO := go -C go
BUF := .bin/buf
PNPM := pnpm --dir web

.PHONY: candidates-review questions-gate deck-gate bracket-gate revise-gate chat-probe generate-probe summary-judge questions-eval eval-calibrate autotune m5-sheet m5-report store-check themes-check ste-check help doctor buf proto proto-check proto-breaking lint lint-go lint-web test test-repeat test-smoke llm-defaults-check cover build dev dev-docker dev-seed run-api run-worker run-web clean

help: ## Show this help
	@grep -E '^[a-zA-Z0-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "  %-14s %s\n", $$1, $$2}'

doctor: ## Check that every required tool is installed
	@./scripts/doctor.sh

$(BUF): go/go.mod go/go.sum ## Build the pinned buf binary from go/go.mod
	@mkdir -p .bin && $(GO) build -o ../$(BUF) github.com/bufbuild/buf/cmd/buf

buf: $(BUF)

proto: $(BUF) ## Generate Go and TypeScript code from proto/ (output is committed)
	@echo "==> buf lint"
	@$(BUF) lint
	@echo "==> buf generate"
	@$(BUF) generate
	@echo "==> go mod tidy"
	@$(GO) mod tidy

proto-check: proto ## Fail if generated code differs from the committed code
	@git diff --exit-code --stat -- go/gen web/packages/api-client/src/gen go/go.mod go/go.sum \
		&& test -z "$$(git ls-files --others --exclude-standard -- go/gen web/packages/api-client/src/gen)" \
		|| (echo "Generated code is stale. Run: make proto && git add -A" && exit 1)

# The ref to compare against. A CI checkout of a pull request has no local
# main branch, only refs/remotes/origin/main (D-254). origin/main exists
# in both, and the fallback covers a clone with no remote.
PROTO_BASE = $(shell git rev-parse --verify --quiet origin/main >/dev/null && echo origin/main || echo main)

proto-breaking: $(BUF) ## Fail on a breaking proto change against the main branch (needs fetch-depth 0 in CI)
	@echo "==> buf breaking against $(PROTO_BASE)"
	@$(BUF) breaking --against '.git#branch=$(PROTO_BASE)'

lint: lint-go lint-web ste-check ## Lint Go, TypeScript, and the docs

lint-go: ## Lint Go (vet + golangci-lint, built from source with the local toolchain)
	@echo "==> go vet"
	@$(GO) vet ./...
	@echo "==> golangci-lint"
	@$(GO) run github.com/golangci/golangci-lint/v2/cmd/golangci-lint@$(GOLANGCI_LINT_VERSION) run ./... -c .golangci.yml

lint-web: ## Lint and typecheck TypeScript
	@echo "==> web lint + typecheck"
	@$(PNPM) lint
	@$(PNPM) typecheck

# Every .md file follows ASD-STE100 (CLAUDE.md rule 2). The script reads
# the prose and skips tables and code blocks. Tracked and untracked files
# are both checked, so a new document is checked before its first commit.
# Dated records are exempt on purpose (D-304): the gate and eval documents
# (docs/reference/pr[0-9]*) and the session logs are machine output or a
# record of one day, and the connector-syncer and wallabee notes are
# frozen copies of another repo. The audits (docs/audit-*) are a record of
# one day, so they are exempt too. Test fixtures under testdata are data.
STE_FILES := $(shell (git ls-files '*.md'; git ls-files --others --exclude-standard '*.md') | sort -u | grep -vE '^docs/reference/(pr[0-9]|session-log|connector-syncer|wallabee)|^docs/audit-|/testdata/')

ste-check: ## Check every hand-written .md file against the STE rules (no cost)
	@echo "==> ste-check"
	@python3 docs/tools/ste-check.py $(STE_FILES)

# LLM_REQUIRE_KEYS is not set here. The unit tests must pass with the
# package default. Set it in a test with t.Setenv when a case needs it.
test: ## Run Go and web unit tests
	@$(GO) test -race ./...
	@$(PNPM) test

test-repeat: ## Run one Go test N times to catch flakes. Usage: make test-repeat TEST=TestCheck RUNS=25
	@[ -n "$(TEST)" ] || { echo "test-repeat: set TEST=TestName. Usage: make test-repeat TEST=TestCheck RUNS=$(or $(RUNS),25)"; exit 1; }
	@$(GO) test -race -run '$(TEST)' -count=$(or $(RUNS),25) ./...

# Keys are required here on purpose (LLM_REQUIRE_KEYS keeps its default).
test-smoke: ## Run the live LLM smoke test (needs OPENAI_API_KEY and ANTHROPIC_API_KEY, read from .env)
	@[ -f .env ] || { echo "test-smoke: .env is absent. Run: cp .env.example .env, then add the provider keys."; exit 1; }
	@set -a && . ./.env && set +a && \
		LLM_SMOKE=1 $(GO) test -race -run TestSmokeLiveProviders -v -count=1 ./internal/llm/

llm-defaults-check: ## Warn when roles.json or prices.json differ from the merge base (never fails)
	@./scripts/check-llm-defaults.sh

# REVIEW_OUT names the gate document. A rerun must never overwrite a
# document that already holds scores (D-65). Raise the run number instead.
REVIEW_OUT ?= docs/reference/pr6-candidate-review-run2.md
# REVIEW_SCORED matches a filled "On theme" cell of the summary table, or
# a filled total line. The tool writes both empty.
REVIEW_SCORED := ^\| [0-9]+ \|[^|]*\|[^|]*\|[[:space:]]*[^[:space:]|]|^Total on theme: [0-9]

candidates-review: ## Write the PR-6 gate document from the local snapshot and the test collection (no model calls, no cost)
	@test ! -f $(REVIEW_OUT) || ! grep -qE '$(REVIEW_SCORED)' $(REVIEW_OUT) || \
		{ echo "$(REVIEW_OUT) holds scores. Set REVIEW_OUT to a new file."; exit 1; }
	@CARDS_SNAPSHOT_DIR=$(CURDIR)/.local/gcs/mtg-local-cards/scryfall $(GO) run ./cmd/candidates-review \
		-collection internal/collections/testdata/manabox_collection.csv -out $(abspath $(REVIEW_OUT))
	@echo "wrote $(REVIEW_OUT)"

# GATE_OUT names the PR-7 gate document. A rerun must never overwrite a
# scored document (D-65).
GATE_OUT ?= docs/reference/pr7-question-gate.md
# GATE_RUN is the run file of PR-15, named after the document (D-65).
GATE_RUN ?= docs/reference/eval/$(notdir $(basename $(GATE_OUT))).jsonl
# GATE_ARGS passes flags to the gate, for example -only 109 for one probe.
GATE_ARGS ?=

questions-gate: ## Write the PR-7 gate document. CAUTION: this calls the real providers and costs money
	@[ -f .env ] || { echo "questions-gate: .env is absent. Run: cp .env.example .env, then add the provider keys."; exit 1; }
	@test ! -f $(GATE_OUT) || ! grep -q '^Verdict:' $(GATE_OUT) || \
		{ echo "$(GATE_OUT) holds a verdict. Set GATE_OUT to a new file."; exit 1; }
	@test ! -f $(GATE_RUN) || { echo "$(GATE_RUN) exists. Set GATE_RUN to a new file."; exit 1; }
	@set -a && . ./.env && set +a && \
		QUESTIONS_GATE=1 CARDS_SNAPSHOT_DIR=$(CURDIR)/.local/gcs/mtg-local-cards/scryfall \
		$(GO) run ./cmd/questions-gate -collection internal/collections/testdata/manabox_collection.csv \
		-run-out $(abspath $(GATE_RUN)) $(GATE_ARGS) > $(GATE_OUT)
	@echo "wrote $(GATE_OUT) and $(GATE_RUN)"

# --- The PR-8 gate, the revise gate, and the three probes (T-18) --------
# Each target calls a real provider and costs money. Each one has the
# same two guards as questions-gate: an env variable the command checks,
# and an output file a rerun must never overwrite (D-65). A gate writes
# its document first and then exits 1 on a FAIL verdict, so the file is
# complete when the target fails.
DECK_GATE_OUT ?= docs/reference/pr8-deck-gate.md
# DECK_GATE_ARGS passes flags to the gate, for example -only 19,20 to run
# the set prompts of PR-17B alone.
DECK_GATE_ARGS ?=
CHAT_PROBE_OUT ?= .local/probes/chat-probe.txt
GENERATE_PROBE_OUT ?= .local/probes/generate-probe.txt
# SUMMARY_JUDGE_IN is the deck gate document the judge reads.
SUMMARY_JUDGE_IN ?= $(DECK_GATE_OUT)
SUMMARY_JUDGE_OUT ?= .local/probes/summary-judge.txt
# CHAT_PROBE_MESSAGES are the user's turns, separated by |.
CHAT_PROBE_MESSAGES ?= Build me a lifegain Commander deck from any cards.|Karlov of the Ghost Council. Bracket 3, white and black, and no budget.
GENERATE_PROBE_THEME ?= lifegain
GENERATE_PROBE_COMMANDER ?= Karlov of the Ghost Council

# DECK_GATE_RUN is the run file of PR-15: the header and the rows of the
# document, as JSONL, named after the document. The command refuses an
# existing file (D-65).
DECK_GATE_RUN ?= docs/reference/eval/$(notdir $(basename $(DECK_GATE_OUT))).jsonl

deck-gate: ## Write the PR-8 deck gate document. CAUTION: calls a real provider and costs money
	@[ -f .env ] || { echo "deck-gate: .env is absent."; exit 1; }
	@test ! -f $(DECK_GATE_OUT) || ! grep -q '^Verdict:' $(DECK_GATE_OUT) || \
		{ echo "$(DECK_GATE_OUT) holds a verdict. Set DECK_GATE_OUT to a new file."; exit 1; }
	@test ! -f $(DECK_GATE_RUN) || { echo "$(DECK_GATE_RUN) exists. Set DECK_GATE_RUN to a new file."; exit 1; }
	@set -a && . ./.env && set +a && \
		DECK_GATE=1 CARDS_SNAPSHOT_DIR=$(CURDIR)/.local/gcs/mtg-local-cards/scryfall \
		$(GO) run ./cmd/deck-gate -collection internal/collections/testdata/manabox_collection.csv \
		-run-out $(abspath $(DECK_GATE_RUN)) $(DECK_GATE_ARGS) > $(DECK_GATE_OUT)
	@echo "wrote $(DECK_GATE_OUT) and $(DECK_GATE_RUN)"

# BRACKET_GATE_OUT is the PR-14A gate document, and BRACKET_GATE_ARGS
# passes flags, for example -only 7,8,9 for the bracket 3 prompts alone.
BRACKET_GATE_OUT ?= docs/reference/pr14a-bracket-gate.md
BRACKET_GATE_ARGS ?=
# BRACKET_GATE_RUN is the run file of PR-15, named after the document.
BRACKET_GATE_RUN ?= docs/reference/eval/$(notdir $(basename $(BRACKET_GATE_OUT))).jsonl

bracket-gate: ## Write the PR-14A bracket gate document. CAUTION: calls a real provider and costs money
	@[ -f .env ] || { echo "bracket-gate: .env is absent."; exit 1; }
	@test ! -f $(BRACKET_GATE_OUT) || ! grep -q '^Verdict:' $(BRACKET_GATE_OUT) || \
		{ echo "$(BRACKET_GATE_OUT) holds a verdict. Set BRACKET_GATE_OUT to a new file."; exit 1; }
	@test ! -f $(BRACKET_GATE_RUN) || { echo "$(BRACKET_GATE_RUN) exists. Set BRACKET_GATE_RUN to a new file."; exit 1; }
	@set -a && . ./.env && set +a && \
		BRACKET_GATE=1 CARDS_SNAPSHOT_DIR=$(CURDIR)/.local/gcs/mtg-local-cards/scryfall \
		$(GO) run ./cmd/bracket-gate -run-out $(abspath $(BRACKET_GATE_RUN)) $(BRACKET_GATE_ARGS) > $(BRACKET_GATE_OUT)
	@echo "wrote $(BRACKET_GATE_OUT) and $(BRACKET_GATE_RUN)"

# The trimmed snapshot of D-521 serves the free dry-run lane of the deck
# gate. TestTrimmedSnapshotBuildsEveryShortlist runs the same lane in CI.
TRIM_SNAPSHOT := $(CURDIR)/go/cmd/deck-gate/testdata/snapshot

deck-gate-dry: ## Build every deck gate shortlist over the trimmed snapshot of the repo, no provider call (D-521)
	@CARDS_SNAPSHOT_DIR=$(TRIM_SNAPSHOT)/scryfall \
		$(GO) run ./cmd/deck-gate -dry -collection internal/collections/testdata/manabox_collection.csv

deck-gate-trim: ## Rewrite the trimmed snapshot from the local store, after a prompt or a fixture changes (D-521)
	@rm -rf $(TRIM_SNAPSHOT)
	@CARDS_SNAPSHOT_DIR=$(CURDIR)/.local/gcs/mtg-local-cards/scryfall \
		$(GO) run ./cmd/deck-gate -dry -trim $(TRIM_SNAPSHOT) -collection internal/collections/testdata/manabox_collection.csv

REVISE_GATE_OUT ?= docs/reference/pr12b-revise-gate.md
# REVISE_GATE_RUN is the run file of PR-15, named after the document.
REVISE_GATE_RUN ?= docs/reference/eval/$(notdir $(basename $(REVISE_GATE_OUT))).jsonl

revise-gate: ## Write the PR-12B revise gate document. CAUTION: calls a real provider and costs money
	@[ -f .env ] || { echo "revise-gate: .env is absent."; exit 1; }
	@test ! -f $(REVISE_GATE_OUT) || ! grep -q '^Verdict:' $(REVISE_GATE_OUT) || \
		{ echo "$(REVISE_GATE_OUT) holds a verdict. Set REVISE_GATE_OUT to a new file."; exit 1; }
	@test ! -f $(REVISE_GATE_RUN) || { echo "$(REVISE_GATE_RUN) exists. Set REVISE_GATE_RUN to a new file."; exit 1; }
	@set -a && . ./.env && set +a && \
		REVISE_GATE=1 CARDS_SNAPSHOT_DIR=$(CURDIR)/.local/gcs/mtg-local-cards/scryfall \
		$(GO) run ./cmd/revise-gate -run-out $(abspath $(REVISE_GATE_RUN)) > $(REVISE_GATE_OUT)
	@echo "wrote $(REVISE_GATE_OUT) and $(REVISE_GATE_RUN)"

chat-probe: ## Drive the real Chat RPC to a deck. CAUTION: calls the real providers and costs money
	@[ -f .env ] || { echo "chat-probe: .env is absent."; exit 1; }
	@test ! -f $(CHAT_PROBE_OUT) || { echo "$(CHAT_PROBE_OUT) exists. Set CHAT_PROBE_OUT to a new file."; exit 1; }
	@mkdir -p $(dir $(CHAT_PROBE_OUT))
	@set -a && . ./.env && set +a && \
		CHAT_PROBE=1 CARDS_SNAPSHOT_DIR=$(CURDIR)/.local/gcs/mtg-local-cards/scryfall \
		$(GO) run ./cmd/chat-probe -messages "$(CHAT_PROBE_MESSAGES)" | tee $(CHAT_PROBE_OUT)
	@echo "wrote $(CHAT_PROBE_OUT)"

generate-probe: ## Build one deck with the real generate role. CAUTION: calls a real provider and costs money
	@[ -f .env ] || { echo "generate-probe: .env is absent."; exit 1; }
	@test ! -f $(GENERATE_PROBE_OUT) || { echo "$(GENERATE_PROBE_OUT) exists. Set GENERATE_PROBE_OUT to a new file."; exit 1; }
	@mkdir -p $(dir $(GENERATE_PROBE_OUT))
	@set -a && . ./.env && set +a && \
		GENERATE_PROBE=1 CARDS_SNAPSHOT_DIR=$(CURDIR)/.local/gcs/mtg-local-cards/scryfall \
		$(GO) run ./cmd/generate-probe -theme "$(GENERATE_PROBE_THEME)" -commander "$(GENERATE_PROBE_COMMANDER)" | tee $(GENERATE_PROBE_OUT)
	@echo "wrote $(GENERATE_PROBE_OUT)"

summary-judge: ## Judge every deck summary of a deck gate document (F-26). CAUTION: calls a real provider and costs money
	@[ -f .env ] || { echo "summary-judge: .env is absent."; exit 1; }
	@test -f $(SUMMARY_JUDGE_IN) || { echo "no deck gate document at $(SUMMARY_JUDGE_IN). Set SUMMARY_JUDGE_IN."; exit 1; }
	@test ! -f $(SUMMARY_JUDGE_OUT) || { echo "$(SUMMARY_JUDGE_OUT) exists. Set SUMMARY_JUDGE_OUT to a new file."; exit 1; }
	@mkdir -p $(dir $(SUMMARY_JUDGE_OUT))
	@set -a && . ./.env && set +a && \
		SUMMARY_JUDGE=1 $(GO) run ./cmd/summary-judge -in $(abspath $(SUMMARY_JUDGE_IN)) | tee $(SUMMARY_JUDGE_OUT)
	@echo "wrote $(SUMMARY_JUDGE_OUT)"
# --- end of the PR-8 gate and probe targets ------------------------------

# M5_OUT names the scoring sheet. A rerun must never overwrite a sheet
# that holds scores (D-65). `make m5-report` reads this sheet, so a new
# sheet needs a new name:
#
#   M5_OUT=docs/reference/pr7-m5-scoring-run18.md make m5-sheet
M5_OUT ?= docs/reference/pr7-m5-scoring.md

# M5_RUNS names the gate documents the sheet may read. Only a run whose
# engine matches the current code belongs here (D-96): a run made before
# a prompt or catalog change measures the old engine, not the catalog.
# Paths are repo-relative. The recipe makes them absolute for the Go
# tool, which runs from go/. An absolute path also works.
M5_RUNS ?= docs/reference/pr7-question-gate-run24.md

# The guard reads every field a scorer fills, in any case, and free text
# counts. filled_slot is left out because the generator pre-fills it.
M5_SCORED := \| (catalog_enough|invented_better|right_slot|faults|catalog_action) \|[[:space:]]*[^[:space:]|]

m5-sheet: ## Build the M-5 scoring sheet from the gate documents (no model calls, no cost)
	@test ! -f $(M5_OUT) || ! grep -qEi '$(M5_SCORED)' $(M5_OUT) || \
		{ echo "$(M5_OUT) holds scores. Set M5_OUT to a new file."; exit 1; }
	@$(GO) run ./cmd/m5-sheet $(abspath $(M5_RUNS)) > $(M5_OUT)
	@echo "wrote $(M5_OUT)"

m5-report: ## Read the scored M-5 sheet and compute the thresholds (no model calls, no cost)
	@$(GO) run ./cmd/m5-report $(abspath $(M5_OUT))

# EVAL_RUN names the gate document the eval scores, and EVAL_OUT the
# document it writes. The eval calls a real provider on the cost tier, so
# it costs money (D-133).
#
# The defaults name the latest gate run. Its eval document and its JSON
# summary already exist, so the guards below refuse a run with the
# defaults. That is correct: set EVAL_OUT and EVAL_JSON to new names for
# a rerun. EVAL_JSON is the baseline that tune-check compares against, so
# a rerun must never overwrite it (D-65). The command checks both paths
# again before its first call.
EVAL_RUN ?= docs/reference/pr7-question-gate-run24.md
EVAL_OUT ?= docs/reference/pr7-question-eval-run24.md
EVAL_JSON ?= .local/tune/run24.json
EVAL_BUDGET ?= 0.50
# EVAL_ROWS is the run file of PR-15, named after the eval document.
# EVAL_RUN is the gate document the eval reads, so this one is not
# called EVAL_RUN.
EVAL_ROWS ?= docs/reference/eval/$(notdir $(basename $(EVAL_OUT))).jsonl

questions-eval: ## Score every question of a gate run. CAUTION: calls a real provider and costs money
	@[ -f .env ] || { echo "questions-eval: .env is absent. Run: cp .env.example .env, then add the provider keys."; exit 1; }
	@test -f $(EVAL_RUN) || { echo "no gate document at $(EVAL_RUN). Set EVAL_RUN."; exit 1; }
	@test ! -f $(EVAL_OUT) || { echo "$(EVAL_OUT) exists. Set EVAL_OUT to a new file."; exit 1; }
	@test ! -f $(EVAL_JSON) || { echo "$(EVAL_JSON) exists. Set EVAL_JSON to a new file."; exit 1; }
	@test ! -f $(EVAL_ROWS) || { echo "$(EVAL_ROWS) exists. Set EVAL_ROWS to a new file."; exit 1; }
	@mkdir -p $(dir $(EVAL_JSON))
	@set -a && . ./.env && set +a && QUESTIONS_EVAL=1 \
		CARDS_SNAPSHOT_DIR=$(CURDIR)/.local/gcs/mtg-local-cards/scryfall \
		$(GO) run ./cmd/questions-eval -in $(abspath $(EVAL_RUN)) -out $(abspath $(EVAL_OUT)) \
		-json $(abspath $(EVAL_JSON)) -run-out $(abspath $(EVAL_ROWS)) -budget $(EVAL_BUDGET)
	@echo "wrote $(EVAL_OUT) and $(EVAL_ROWS)"

# The eval role runs on the model that also writes the questions, for
# cost (D-133). This target measures what that costs in judgment: it
# scores a sample twice and compares the two verdicts.
CALIBRATE_N ?= 12
CALIBRATE_MODEL ?= claude-opus-5
CALIBRATE_PROVIDER ?= anthropic
# CALIBRATE_OUT is the stem of the two summaries this target writes:
# <stem>-base.json and <stem>-strong.json. A paid result is never
# overwritten (D-65). Set CALIBRATE_OUT to a new stem for a rerun.
CALIBRATE_OUT ?= .local/tune/calibrate-run24
CALIBRATE_BASE := $(CALIBRATE_OUT)-base.json
CALIBRATE_STRONG := $(CALIBRATE_OUT)-strong.json

eval-calibrate: ## Measure the eval model against a stronger one on a sample. CAUTION: costs money
	@[ -f .env ] || { echo "eval-calibrate: .env is absent. Run: cp .env.example .env, then add the provider keys."; exit 1; }
	@test -f $(EVAL_RUN) || { echo "no gate document at $(EVAL_RUN). Set EVAL_RUN."; exit 1; }
	@test ! -f $(CALIBRATE_BASE) || { echo "$(CALIBRATE_BASE) exists. Set CALIBRATE_OUT to a new stem."; exit 1; }
	@test ! -f $(CALIBRATE_STRONG) || { echo "$(CALIBRATE_STRONG) exists. Set CALIBRATE_OUT to a new stem."; exit 1; }
	@mkdir -p $(dir $(CALIBRATE_OUT))
	@set -a && . ./.env && set +a && QUESTIONS_EVAL=1 \
		$(GO) run ./cmd/questions-eval -in $(abspath $(EVAL_RUN)) -n $(CALIBRATE_N) -holdout 0 \
		-json $(abspath $(CALIBRATE_BASE)) -out /dev/null -budget $(EVAL_BUDGET)
	@set -a && . ./.env && set +a && QUESTIONS_EVAL=1 \
		LLM_EVAL_PROVIDER=$(CALIBRATE_PROVIDER) LLM_EVAL_MODEL=$(CALIBRATE_MODEL) \
		$(GO) run ./cmd/questions-eval -in $(abspath $(EVAL_RUN)) -n $(CALIBRATE_N) -holdout 0 \
		-json $(abspath $(CALIBRATE_STRONG)) -out /dev/null -budget $(EVAL_BUDGET)
	@$(GO) run ./cmd/tune-check -agree \
		-next $(abspath $(CALIBRATE_BASE)) -prev $(abspath $(CALIBRATE_STRONG))

autotune: ## Print how to start the overnight tuning loop. It never starts one
	@echo "The loop edits code and commits to its own branch with nobody watching. It pushes only with --push."
	@echo "docs/reference/autotune-readme.md holds every command."
	@echo "docs/reference/autotune-design.md holds the reasons and the limits."
	@echo
	@echo "  AUTOTUNE_ALLOW_UNATTENDED=1 AUTOTUNE_FIXER_CMD=... scripts/autotune.sh --budget 3.00"

store-check: ## Run the session, deck, and collection stores against the local Firestore emulator (needs `firebase emulators:start --only firestore`)
	@nc -z 127.0.0.1 8281 2>/dev/null || \
		{ echo "no Firestore emulator on :8281. Start one: firebase emulators:start --only firestore --project mtg-local"; exit 1; }
	@FIRESTORE_EMULATOR_HOST=127.0.0.1:8281 $(GO) test ./internal/sessions ./internal/decks ./internal/collections -count=1

# --- The eval harness of PR-15 (free) -----------------------------------
# EVAL_DIR holds the run files the gates write and baselines.json. The
# check reads every suite of the baselines against its newest run and
# names the flips. It calls no provider, so CI runs it (Tier 0).
EVAL_DIR ?= docs/reference/eval
EVAL_MARGIN ?= 0

eval-check: ## Compare every baseline of the eval harness with its newest run (PR-15, no cost)
	@$(GO) run ./cmd/eval check -dir $(abspath $(EVAL_DIR)) -margin $(EVAL_MARGIN)

# QUALITY_GATE_OUT is the PR-14B gate document. The run is free: it
# fits the model over the stored lists and calls no provider.
QUALITY_GATE_OUT ?= docs/reference/pr14b-quality-gate.md
QUALITY_GATE_ARGS ?=
# QUALITY_GATE_RUN is the run file of PR-15, named after the document.
QUALITY_GATE_RUN ?= docs/reference/eval/$(notdir $(basename $(QUALITY_GATE_OUT))).jsonl

quality-gate: ## Write the PR-14B quality gate document from the local meta store (no model calls, no cost)
	@test ! -f $(QUALITY_GATE_OUT) || ! grep -q '^Verdict:' $(QUALITY_GATE_OUT) || \
		{ echo "$(QUALITY_GATE_OUT) holds a verdict. Set QUALITY_GATE_OUT to a new file."; exit 1; }
	@test ! -f $(QUALITY_GATE_RUN) || { echo "$(QUALITY_GATE_RUN) exists. Set QUALITY_GATE_RUN to a new file."; exit 1; }
	@CARDS_SNAPSHOT_DIR=$(CURDIR)/.local/gcs/mtg-local-cards/scryfall \
		$(GO) run ./cmd/quality-gate -run-out $(abspath $(QUALITY_GATE_RUN)) $(QUALITY_GATE_ARGS) > $(QUALITY_GATE_OUT)
	@echo "wrote $(QUALITY_GATE_OUT) and $(QUALITY_GATE_RUN)"

# QUALITY_JUDGE_IN is the deck gate document the tier judge lane reads,
# and QUALITY_JUDGE_OUT the judge document it writes.
QUALITY_JUDGE_IN ?= docs/reference/pr8-deck-gate-run13b.md
QUALITY_JUDGE_OUT ?= docs/reference/pr14b-quality-judge.md
# QUALITY_JUDGE_RUN is the run file of PR-15, named after the document.
QUALITY_JUDGE_RUN ?= docs/reference/eval/$(notdir $(basename $(QUALITY_JUDGE_OUT))).jsonl

quality-judge: ## Judge the tier of every deck of a deck gate document (PR-14B). CAUTION: calls a real provider and costs money
	@[ -f .env ] || { echo "quality-judge: .env is absent."; exit 1; }
	@test -f $(QUALITY_JUDGE_IN) || { echo "no deck gate document at $(QUALITY_JUDGE_IN). Set QUALITY_JUDGE_IN."; exit 1; }
	@test ! -f $(QUALITY_JUDGE_OUT) || ! grep -q '^Verdict:' $(QUALITY_JUDGE_OUT) || \
		{ echo "$(QUALITY_JUDGE_OUT) holds a verdict. Set QUALITY_JUDGE_OUT to a new file."; exit 1; }
	@test ! -f $(QUALITY_JUDGE_RUN) || { echo "$(QUALITY_JUDGE_RUN) exists. Set QUALITY_JUDGE_RUN to a new file."; exit 1; }
	@set -a && . ./.env && set +a && \
		QUALITY_JUDGE=1 CARDS_SNAPSHOT_DIR=$(CURDIR)/.local/gcs/mtg-local-cards/scryfall \
		$(GO) run ./cmd/quality-gate -judge $(abspath $(QUALITY_JUDGE_IN)) -run-out $(abspath $(QUALITY_JUDGE_RUN)) > $(QUALITY_JUDGE_OUT)
	@echo "wrote $(QUALITY_JUDGE_OUT) and $(QUALITY_JUDGE_RUN)"

# META_ARGS passes flags to the worker's meta job, for example
# -meta-months 3 -meta-pages 50 for a short first read.
META_ARGS ?=

meta-refresh: ## Read the deck list sources into the local meta store and fit the quality model (network, no model calls, no cost)
	@[ -f .env ] && set -a && . ./.env && set +a; \
		CARDS_SNAPSHOT_DIR=$(CURDIR)/.local/gcs/mtg-local-cards/scryfall \
		$(GO) run ./cmd/worker -meta $(META_ARGS)

themes-check: ## Check the theme slugs and the commander ranking against the local snapshot
	@CARDS_SNAPSHOT_DIR=$(CURDIR)/.local/gcs/mtg-local-cards/scryfall $(GO) test ./internal/candidates -run 'TestThemeSlugsExist|TestCommanderQualitySnapshot' -count=1

cover: ## Go coverage report
	@$(GO) test -coverprofile=coverage.out ./... && $(GO) tool cover -func=coverage.out | tail -1

build: ## Build the Go binaries and the web app
	@$(GO) build -o bin/api ./cmd/api
	@$(GO) build -o bin/worker ./cmd/worker
	@$(PNPM) build

run-api: ## Run the API on :8080 (make dev uses :8090)
	@$(GO) run ./cmd/api

run-worker: ## Run the worker
	@$(GO) run ./cmd/worker

run-web: ## Run the Vite dev server on :5180
	@$(PNPM) dev

dev: ## Start the local stack: emulators, fake GCS, API, worker, web. No cloud credentials.
	@./scripts/dev.sh

dev-seed: ## Download the Scryfall snapshot into the local stack (network, ~110 MB)
	@echo "==> one-shot card snapshot refresh (needs make dev running for fake GCS)"
	@PROJECT_ID=mtg-local CARDS_BUCKET=mtg-local-cards \
		STORAGE_EMULATOR_HOST=http://127.0.0.1:4443 \
		$(GO) run ./cmd/worker -once

dev-docker: ## Start the emulators, fake GCS, and API in containers (Compose). Seed: docker compose --profile seed run --rm worker
	@docker compose up --build

clean: ## Remove build outputs
	@rm -rf go/bin go/coverage.out web/apps/web/dist
