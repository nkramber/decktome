# mtg-deck-builder - single human entry point.
# Every target prints what it does. Versions are pinned in go/go.mod, web/package.json, and here.

GOLANGCI_LINT_VERSION := v2.13.2
GO := go -C go
BUF := .bin/buf
PNPM := pnpm --dir web

.PHONY: candidates-review questions-gate deck-gate chat-probe generate-probe summary-judge questions-eval eval-calibrate autotune m5-sheet m5-report store-check themes-check ste-check help doctor buf proto proto-check proto-breaking lint lint-go lint-web test test-repeat test-smoke llm-defaults-check cover build dev dev-docker dev-seed run-api run-worker run-web clean

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
# main branch, only refs/remotes/origin/main, so "branch=main" fails there
# with "couldn't find remote ref main". A push to main does create the
# local branch, which is why this passed on main and failed on every pull
# request (D-254). origin/main exists in both, and the fallback covers a
# clone with no remote.
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
# the prose and skips tables and code blocks. The gate and eval documents
# under docs/reference/ are machine output, and the frozen research notes
# keep their date, so both stay out of the list.
STE_FILES := $(shell git ls-files '*.md' | grep -vE '^docs/reference/(pr[0-9]|session-log|connector-syncer|wallabee)')

ste-check: ## Check every hand-written .md file against the STE rules (no cost)
	@echo "==> ste-check"
	@python3 docs/tools/ste-check.py $(STE_FILES)

# LLM_REQUIRE_KEYS is not set here. The unit tests must pass with the
# package default. Set it in a test with t.Setenv when a case needs it.
test: ## Run Go and web unit tests
	@$(GO) test -race ./...
	@$(PNPM) test

test-repeat: ## Run one Go test N times to catch flakes. Usage: make test-repeat TEST=TestCheck RUNS=25
	@[ -n "$(TEST)" ] || { echo "test-repeat: set TEST=TestName, or every test runs $(or $(RUNS),25) times."; exit 1; }
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

candidates-review: ## Write the PR-6 gate document from the local snapshot and the owner's export
	@test ! -f $(REVIEW_OUT) || ! grep -q '^Verdict:' $(REVIEW_OUT) || \
		{ echo "$(REVIEW_OUT) holds scores. Set REVIEW_OUT to a new file."; exit 1; }
	@CARDS_SNAPSHOT_DIR=$(CURDIR)/.local/gcs/mtg-local-cards/scryfall $(GO) run ./cmd/candidates-review \
		-collection internal/collections/testdata/manabox_collection.csv > $(REVIEW_OUT)
	@echo "wrote $(REVIEW_OUT)"

# GATE_OUT names the PR-7 gate document. A rerun must never overwrite a
# scored document (D-65).
GATE_OUT ?= docs/reference/pr7-question-gate.md

questions-gate: ## Write the PR-7 gate document. CAUTION: this calls the real providers and costs money
	@[ -f .env ] || { echo "questions-gate: .env is absent. Run: cp .env.example .env, then add the provider keys."; exit 1; }
	@test ! -f $(GATE_OUT) || ! grep -q '^Verdict:' $(GATE_OUT) || \
		{ echo "$(GATE_OUT) holds a verdict. Set GATE_OUT to a new file."; exit 1; }
	@set -a && . ./.env && set +a && \
		QUESTIONS_GATE=1 CARDS_SNAPSHOT_DIR=$(CURDIR)/.local/gcs/mtg-local-cards/scryfall \
		$(GO) run ./cmd/questions-gate -collection internal/collections/testdata/manabox_collection.csv > $(GATE_OUT)
	@echo "wrote $(GATE_OUT)"

# --- The PR-8 gate and the three probes (audit 2026-08-28, T-18) -------
# Each target calls a real provider and costs money. Each one has the
# same two guards as questions-gate: an env variable the command checks,
# and an output file a rerun must never overwrite (D-65).
DECK_GATE_OUT ?= docs/reference/pr8-deck-gate.md
CHAT_PROBE_OUT ?= .local/probes/chat-probe.txt
GENERATE_PROBE_OUT ?= .local/probes/generate-probe.txt
# SUMMARY_JUDGE_IN is the deck gate document the judge reads.
SUMMARY_JUDGE_IN ?= $(DECK_GATE_OUT)
SUMMARY_JUDGE_OUT ?= .local/probes/summary-judge.txt
# CHAT_PROBE_MESSAGES are the user's turns, separated by |.
CHAT_PROBE_MESSAGES ?= Build me a lifegain Commander deck from any cards.|Karlov of the Ghost Council. Bracket 3, white and black, and no budget.
GENERATE_PROBE_THEME ?= lifegain
GENERATE_PROBE_COMMANDER ?= Karlov of the Ghost Council

deck-gate: ## Write the PR-8 deck gate document. CAUTION: calls a real provider and costs money
	@[ -f .env ] || { echo "deck-gate: .env is absent."; exit 1; }
	@test ! -f $(DECK_GATE_OUT) || ! grep -q '^Verdict:' $(DECK_GATE_OUT) || \
		{ echo "$(DECK_GATE_OUT) holds a verdict. Set DECK_GATE_OUT to a new file."; exit 1; }
	@set -a && . ./.env && set +a && \
		DECK_GATE=1 CARDS_SNAPSHOT_DIR=$(CURDIR)/.local/gcs/mtg-local-cards/scryfall \
		$(GO) run ./cmd/deck-gate -collection internal/collections/testdata/manabox_collection.csv > $(DECK_GATE_OUT)
	@echo "wrote $(DECK_GATE_OUT)"

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
# the owner has scored.
# M5_OUT stays the version-1 sheet, because `make m5-report` reads it and
# the owner is still scoring it. The scored-sheet guard below refuses to
# write over it, so building the version-2 sheet needs a new name:
#
#   M5_OUT=docs/reference/pr7-m5-scoring-run18.md make m5-sheet
M5_OUT ?= docs/reference/pr7-m5-scoring.md

# M5_RUNS names the gate documents the sheet may read. Only a run whose
# engine matches the current code belongs here (D-96). A run made before
# a defect was fixed measures the defect, not the catalog.
#
# Runs 1 to 13 are all held back now. The fixes of 2026-08-25 changed the
# catalog rows, the classify and ask prompts, and the conversation set
# (D-104 to D-116). A sheet built from them would measure the old engine.
# The next sheet reads the latest gate document alone.
# Paths are repo-relative. The recipe makes them absolute for the Go
# tool, which runs from go/. An absolute path also works.
M5_RUNS ?= docs/reference/pr7-question-gate-run24.md

# The guard reads every field the owner fills, in any case, and free text
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
# document it writes. The eval calls a real provider, so it costs money.
# A 66-conversation run on the cost tier is about eleven cents (D-133).
#
# The defaults name the latest gate run. Its eval document and its JSON
# summary already exist, so the guards below refuse a run with the
# defaults. That is correct: set EVAL_OUT and EVAL_JSON to new names for
# a rerun. EVAL_JSON is the baseline that tune-check compares against, so
# a rerun must never overwrite it (D-65).
EVAL_RUN ?= docs/reference/pr7-question-gate-run24.md
EVAL_OUT ?= docs/reference/pr7-question-eval-run24.md
EVAL_JSON ?= .local/tune/run24.json
EVAL_BUDGET ?= 0.50

questions-eval: ## Score every question of a gate run. CAUTION: calls a real provider and costs money
	@[ -f .env ] || { echo "questions-eval: .env is absent. Run: cp .env.example .env, then add the provider keys."; exit 1; }
	@test -f $(EVAL_RUN) || { echo "no gate document at $(EVAL_RUN). Set EVAL_RUN."; exit 1; }
	@test ! -f $(EVAL_OUT) || { echo "$(EVAL_OUT) exists. Set EVAL_OUT to a new file."; exit 1; }
	@test ! -f $(EVAL_JSON) || { echo "$(EVAL_JSON) exists. Set EVAL_JSON to a new file."; exit 1; }
	@mkdir -p $(dir $(EVAL_JSON))
	@set -a && . ./.env && set +a && QUESTIONS_EVAL=1 \
		CARDS_SNAPSHOT_DIR=$(CURDIR)/.local/gcs/mtg-local-cards/scryfall \
		$(GO) run ./cmd/questions-eval -in ../$(EVAL_RUN) -out ../$(EVAL_OUT) \
		-json ../$(EVAL_JSON) -budget $(EVAL_BUDGET)
	@echo "wrote $(EVAL_OUT)"

# The eval role runs on the model that also writes the questions, which
# the owner chose for cost (D-133). This target measures what that costs
# in judgment: it scores a sample twice and compares the two verdicts.
CALIBRATE_N ?= 12
CALIBRATE_MODEL ?= claude-sonnet-5
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
	@echo "The loop edits code and pushes with nobody watching."
	@echo "docs/reference/autotune-readme.md holds every command."
	@echo "docs/reference/autotune-design.md holds the reasons and the limits."
	@echo
	@echo "  AUTOTUNE_ALLOW_UNATTENDED=1 AUTOTUNE_FIXER_CMD=... scripts/autotune.sh --budget 3.00"

store-check: ## Run the session and deck stores against the local Firestore emulator (needs `firebase emulators:start --only firestore`)
	@nc -z 127.0.0.1 8281 2>/dev/null || \
		{ echo "no Firestore emulator on :8281. Start one: firebase emulators:start --only firestore --project mtg-local"; exit 1; }
	@FIRESTORE_EMULATOR_HOST=127.0.0.1:8281 $(GO) test ./internal/sessions ./internal/decks -count=1

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
