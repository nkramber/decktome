# mtg-deck-builder - single human entry point.
# Every target prints what it does. Versions are pinned in go/go.mod, web/package.json, and here.

GOLANGCI_LINT_VERSION := v2.6.2
GO := go -C go
BUF := .bin/buf
PNPM := pnpm --dir web

.PHONY: candidates-review questions-gate questions-eval eval-calibrate autotune m5-sheet m5-report store-check themes-check help doctor buf proto proto-check proto-breaking lint lint-go lint-web test test-repeat test-smoke llm-defaults-check cover build dev dev-docker dev-seed run-api run-worker run-web clean

help: ## Show this help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "  %-14s %s\n", $$1, $$2}'

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

proto-breaking: $(BUF) ## Fail on a breaking proto change against the main branch (needs the main ref, fetch-depth 0 in CI)
	@echo "==> buf breaking against main"
	@$(BUF) breaking --against '.git#branch=main'

lint: lint-go lint-web ## Lint Go and TypeScript

lint-go: ## Lint Go (vet + golangci-lint, built from source with the local toolchain)
	@echo "==> go vet"
	@$(GO) vet ./...
	@echo "==> golangci-lint"
	@$(GO) run github.com/golangci/golangci-lint/v2/cmd/golangci-lint@$(GOLANGCI_LINT_VERSION) run ./... -c .golangci.yml

lint-web: ## Lint and typecheck TypeScript
	@echo "==> web lint + typecheck"
	@$(PNPM) lint
	@$(PNPM) typecheck

# LLM_REQUIRE_KEYS is not set here. The unit tests must pass with the
# package default. Set it in a test with t.Setenv when a case needs it.
test: ## Run Go and web unit tests
	@$(GO) test -race ./...
	@$(PNPM) test

test-repeat: ## Run one Go test N times to catch flakes. Usage: make test-repeat TEST=TestCheck RUNS=25
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
	@CARDS_SNAPSHOT_DIR=../.local/gcs/mtg-local-cards/scryfall $(GO) run ./cmd/candidates-review \
		-collection internal/collections/testdata/manabox_collection.csv > $(REVIEW_OUT)
	@echo "wrote $(REVIEW_OUT)"

# GATE_OUT names the PR-7 gate document. A rerun must never overwrite a
# scored document (D-65).
GATE_OUT ?= docs/reference/pr7-question-gate.md

questions-gate: ## Write the PR-7 gate document. CAUTION: this calls the real providers and costs money
	@test ! -f $(GATE_OUT) || ! grep -q '^Verdict:' $(GATE_OUT) || \
		{ echo "$(GATE_OUT) holds a verdict. Set GATE_OUT to a new file."; exit 1; }
	@set -a && . ./.env && set +a && \
		QUESTIONS_GATE=1 CARDS_SNAPSHOT_DIR=$(CURDIR)/.local/gcs/mtg-local-cards/scryfall \
		$(GO) run ./cmd/questions-gate -collection internal/collections/testdata/manabox_collection.csv > $(GATE_OUT)
	@echo "wrote $(GATE_OUT)"

# M5_OUT names the scoring sheet. A rerun must never overwrite a sheet
# the owner has scored.
# M5_OUT stays the version-1 sheet, because `make m5-report` reads it and
# the owner is still scoring it. The scored-sheet guard below refuses to
# write over it, so building the version-2 sheet needs a new name:
#
#   M5_OUT=docs/reference/pr7-m5-scoring-run14.md make m5-sheet
M5_OUT ?= docs/reference/pr7-m5-scoring.md

# M5_RUNS names the gate documents the sheet may read. Only a run whose
# engine matches the current code belongs here (D-96). A run made before
# a defect was fixed measures the defect, not the catalog.
#
# Runs 1 to 13 are all held back now. The fixes of 2026-08-25 changed the
# catalog rows, the classify and ask prompts, and the conversation set
# (D-104 to D-116). A sheet built from them would measure the old engine.
# The next sheet reads run 14 alone.
M5_RUNS ?= ../docs/reference/pr7-question-gate-run14.md

# The guard reads every field the owner fills, in any case, and free text
# counts. filled_slot is left out because the generator pre-fills it.
M5_SCORED := \| (catalog_enough|invented_better|right_slot|faults|catalog_action) \|[[:space:]]*[^[:space:]|]

m5-sheet: ## Build the M-5 scoring sheet from the gate documents (no model calls, no cost)
	@test ! -f $(M5_OUT) || ! grep -qEi '$(M5_SCORED)' $(M5_OUT) || \
		{ echo "$(M5_OUT) holds scores. Set M5_OUT to a new file."; exit 1; }
	@$(GO) run ./cmd/m5-sheet $(M5_RUNS) > $(M5_OUT)
	@echo "wrote $(M5_OUT)"

m5-report: ## Read the scored M-5 sheet and compute the thresholds (no model calls, no cost)
	@$(GO) run ./cmd/m5-report ../$(M5_OUT)

# EVAL_RUN names the gate document the eval scores, and EVAL_OUT the
# document it writes. The eval calls a real provider, so it costs money.
# A 66-conversation run on the cost tier is about eleven cents (D-133).
EVAL_RUN ?= docs/reference/pr7-question-gate-run14.md
EVAL_OUT ?= docs/reference/pr7-question-eval-run14.md
EVAL_JSON ?= .local/tune/run14.json
EVAL_BUDGET ?= 0.50

questions-eval: ## Score every question of a gate run. CAUTION: calls a real provider and costs money
	@test -f $(EVAL_RUN) || { echo "no gate document at $(EVAL_RUN). Set EVAL_RUN."; exit 1; }
	@test ! -f $(EVAL_OUT) || { echo "$(EVAL_OUT) exists. Set EVAL_OUT to a new file."; exit 1; }
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

eval-calibrate: ## Measure the eval model against a stronger one on a sample. CAUTION: costs money
	@test -f $(EVAL_RUN) || { echo "no gate document at $(EVAL_RUN). Set EVAL_RUN."; exit 1; }
	@mkdir -p .local/tune
	@set -a && . ./.env && set +a && QUESTIONS_EVAL=1 \
		$(GO) run ./cmd/questions-eval -in ../$(EVAL_RUN) -n $(CALIBRATE_N) -holdout 0 \
		-json ../.local/tune/calibrate-base.json -out /dev/null -budget $(EVAL_BUDGET)
	@set -a && . ./.env && set +a && QUESTIONS_EVAL=1 \
		LLM_EVAL_PROVIDER=$(CALIBRATE_PROVIDER) LLM_EVAL_MODEL=$(CALIBRATE_MODEL) \
		$(GO) run ./cmd/questions-eval -in ../$(EVAL_RUN) -n $(CALIBRATE_N) -holdout 0 \
		-json ../.local/tune/calibrate-strong.json -out /dev/null -budget $(EVAL_BUDGET)
	@$(GO) run ./cmd/tune-check -agree \
		-next ../.local/tune/calibrate-base.json -prev ../.local/tune/calibrate-strong.json

autotune: ## Print how to start the overnight tuning loop. It never starts one
	@echo "The loop edits code and pushes with nobody watching."
	@echo "docs/reference/autotune-readme.md holds every command."
	@echo "docs/reference/autotune-design.md holds the reasons and the limits."
	@echo
	@echo "  AUTOTUNE_ALLOW_UNATTENDED=1 AUTOTUNE_FIXER_CMD=... scripts/autotune.sh --budget 3.00"

store-check: ## Run the session store against the local Firestore emulator (needs `firebase emulators:start --only firestore`)
	@nc -z 127.0.0.1 8281 2>/dev/null || \
		{ echo "no Firestore emulator on :8281. Start one: firebase emulators:start --only firestore --project mtg-local"; exit 1; }
	@FIRESTORE_EMULATOR_HOST=127.0.0.1:8281 $(GO) test ./internal/sessions -count=1

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
		go -C go run ./cmd/worker -once

dev-docker: ## Start the emulators, fake GCS, and API in containers (Compose). Seed: docker compose --profile seed run --rm worker
	@docker compose up --build

clean: ## Remove build outputs
	@rm -rf go/bin go/coverage.out web/apps/web/dist
