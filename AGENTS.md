# mtg-deck-builder - AGENTS

Go + Protobuf + TypeScript monorepo. Read `CLAUDE.md` for the owner's rules and `docs/design-roadmap.md` for the plan.

## Layout

- `proto/` - the one contract (buf module `mtg.v1`). Change it here only.
- `go/` - one Go module `github.com/nkramber/mtg-deck-builder/go`. `cmd/api` (Connect-RPC API, :8080 default, :8090 under `make dev`), `cmd/worker` (jobs), `internal/*` (packages), `gen/` (generated, committed). Tools: `cmd/candidates-review` (PR-6 gate document), `cmd/questions-gate` (PR-7 gate run, costs money), `cmd/questions-eval` (scores a gate run, costs money), `cmd/m5-sheet` and `cmd/m5-report` (the hand-scoring sheet), `cmd/tune-check` (accept rules for one loop iteration, free), `cmd/deck-gate` (PR-8 gate document, costs money), `cmd/revise-gate` (PR-12B gate document, costs money), `cmd/chat-probe`, `cmd/generate-probe`, and `cmd/summary-judge` (probes, each costs money). Shared packages for the tools: `internal/gatekit` (the spend and overwrite guards) and `internal/auth` (the Firebase ID-token interceptor, CORS).
- `web/` - pnpm workspace. `apps/web` (React 19 + Vite 7, dev port 5180), `packages/api-client` (generated TypeScript, committed).
- `docs/` - design roadmap, decisions, open questions, reference notes.
- `scripts/` - `doctor.sh`, `dev.sh`, `check-llm-defaults.sh`, `autotune.sh` (the tuning loop, needs `AUTOTUNE_ALLOW_UNATTENDED=1`), `autotune-fix.sh` (the fixer step the loop calls). One-off tools only. Not a source of design patterns.

## Commands

Use the Makefile. It applies the pinned versions.

```bash
make doctor        # check tools
make proto         # regenerate Go + TS from proto/ (commit the output)
make lint          # go vet, golangci-lint, eslint, tsc, and the STE check
make ste-check     # the STE check alone, free, part of make lint
make test          # go test -race, vitest
make test-smoke    # live LLM smoke test, reads .env (fails when .env is absent)
make llm-defaults-check   # warn on a roles.json or prices.json change
make proto-breaking       # buf breaking against main
make cover         # Go coverage summary
make build         # Go binaries + web bundle
make dev           # emulators (:8281, :9199) + fake GCS (:4443) + api (:8090) + worker + web (:5180)
make dev-seed      # one-shot card snapshot refresh (needs make dev)
```

CI runs on pull requests only (D-286). A changes job reads the diff, and each job runs only when its inputs changed. A weekly schedule runs govulncheck alone (D-305).

## Rules

- Never edit `go/gen/` or `web/packages/api-client/src/gen/` by hand. Run `make proto`. CI fails when the committed output is stale.
- No model id string at a call site. Models come from the role layer (`go/internal/llm`, PR-10).
- No ban list, rotation date, or Game Changers list in code or prompts. Legality comes from the card database. One exception: `internal/rules/companion_bans.json` holds the companion-only ban that Scryfall can not express (F-18, Lutri), with a source and a date.
- Every card the model names passes the rules engine before the user sees it.
- Log ids, never PII or raw prompts.
- Table-driven tests. `ctx` is the first parameter. `errors.Is` / `errors.As`. Accept interfaces, return structs.
- One concern per pull request. Squash merge into `main`.
- No AI-attribution text in any PR, branch name, commit message, or comment.
- Write docs and skills in ASD-STE100. Run `make ste-check` before you commit a `.md` file. `make lint` and CI run it too (D-264).
