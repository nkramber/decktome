# mtg-deck-builder - AGENTS

Go + Protobuf + TypeScript monorepo. Read `CLAUDE.md` for the owner's rules and `docs/design-roadmap.md` for the plan.

## Layout

- `proto/` - the one contract (buf module `mtg.v1`). Change it here only.
- `go/` - one Go module `github.com/nkramber/mtg-deck-builder/go`. `cmd/api` (Connect-RPC API, :8080), `cmd/worker` (jobs), `internal/*` (packages), `gen/` (generated, committed).
- `web/` - pnpm workspace. `apps/web` (React 19 + Vite 7, dev port 5180), `packages/api-client` (generated TypeScript, committed).
- `docs/` - design roadmap, decisions, open questions, reference notes.
- `scripts/` - `doctor.sh`, `dev.sh`. One-off tools only. Not a source of design patterns.

## Commands

Use the Makefile. It applies the pinned versions.

```bash
make doctor        # check tools
make proto         # regenerate Go + TS from proto/ (commit the output)
make lint          # go vet, golangci-lint, eslint, tsc
make test          # go test -race, vitest
make dev           # api (:8080) + worker + web (:5180)
```

## Rules

- Never edit `go/gen/` or `web/packages/api-client/src/gen/` by hand. Run `make proto`. CI fails when the committed output is stale.
- No model id string at a call site. Models come from the role layer (`go/internal/llm`, PR-10).
- No ban list, rotation date, or Game Changers list in code or prompts. Legality comes from the card database.
- Every card the model names passes the rules engine before the user sees it.
- Log ids, never PII or raw prompts.
- Table-driven tests. `ctx` is the first parameter. `errors.Is` / `errors.As`. Accept interfaces, return structs.
- One concern per PR. Squash merge into `main`.
- No AI-attribution text in any PR, branch name, commit message, or comment.
- Docs and skills are written in ASD-STE100. Run `python3 docs/tools/ste-check.py <files>` before you commit a `.md` file.
