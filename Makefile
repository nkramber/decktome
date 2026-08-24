# mtg-deck-builder - single human entry point.
# Every target prints what it does. Versions are pinned in go/go.mod, web/package.json, and here.

GOLANGCI_LINT_VERSION := v2.6.2
GO := go -C go
BUF := .bin/buf
PNPM := pnpm --dir web

.PHONY: help doctor buf proto proto-check lint test test-repeat cover build dev dev-docker dev-seed run-api run-worker run-web clean

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

lint: ## Lint Go and TypeScript
	@echo "==> go vet"
	@$(GO) vet ./...
	@echo "==> golangci-lint"
	@$(GO) run github.com/golangci/golangci-lint/v2/cmd/golangci-lint@$(GOLANGCI_LINT_VERSION) run ./... -c .golangci.yml
	@echo "==> web lint + typecheck"
	@$(PNPM) lint
	@$(PNPM) typecheck

test: ## Run Go and web unit tests
	@$(GO) test -race ./...
	@$(PNPM) test

test-repeat: ## Run one Go test N times to catch flakes. Usage: make test-repeat TEST=TestCheck RUNS=25
	@$(GO) test -race -run '$(TEST)' -count=$(or $(RUNS),25) ./...

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

dev-docker: ## Start the emulators, fake GCS, and API in containers (Compose)
	@docker compose up --build

clean: ## Remove build outputs
	@rm -rf go/bin go/coverage.out web/apps/web/dist
