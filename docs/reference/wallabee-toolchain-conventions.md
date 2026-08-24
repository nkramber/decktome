# Reference: toolchain conventions from wallabee-ui and connector-syncer

Status: research note. Verified 2026-08-23 against local checkouts. These are conventions we can inherit for a Go + Protobuf + TypeScript monorepo.

## wallabee-ui (TypeScript frontend)

- React 19, Vite 7, TypeScript 5.8, react-router 8, Tailwind 4, Radix/shadcn primitives.
- pnpm 9.2.0 pinned in `package.json`. Node 24.14.0 pinned in `.nvmrc`.
- Workspace: `apps/web`, `packages/api-types`, `packages/eslint-config`, `packages/tsconfig`, `functions`.
- Generated API types are committed. Nobody edits them by hand. `apps/web/scripts/normalize-openapi.mjs` canonicalizes the spec first. Without it, generated diffs had thousands of phantom lines.
- State in three tiers: TanStack Query for server state, Zustand stores for client state, React context for cross-cutting providers.
- Streaming: `apps/web/src/api/streaming.ts` is one pure SSE primitive with backoff, `AbortSignal`, and a request id. The chat thread (`features/search/components/search-thread.tsx`) consumes it.
- Tests: Vitest with jsdom and `jest-axe`, Storybook 10, Playwright visual and E2E. Dead code check with `knip`.
- CI: a `verify` matrix (`static`, `unit`, `artifacts`, `api-contract`) plus one fan-in `build` job. The `api-contract` job has a hard gate (committed spec vs committed types) and an advisory check (live dev spec vs committed types).
- Architecture boundaries are lint rules, not prose. `packages/eslint-config/web-app.js:139-200` gives each feature an explicit allowlist of sibling features it can import.
- `CLAUDE.md` is one line: `@AGENTS.md`.

## connector-syncer (Python backend)

- Makefile is the human entry point. Guard targets print the fix command. Test scope is variable-driven (`SERVICE=`, `LIBS=`).
- CI uses `dorny/paths-filter` per service and gates steps, not jobs. Filters always include shared libs, dependency manifests, and the workflow file.
- Policy scripts run as CI steps and have their own unit tests.
- `scripts/check_model_default_changes.py` is advisory only.
- House rule: no AI-attribution text in any PR, branch, commit, or comment.

## go-crash-course

This repo is a personal Go learning course. It has no protobuf, gRPC, Connect, or buf. It gives no tooling preference for the new monorepo. The owner's Go style there: `ctx` first, table-driven tests, `errors.Is/As`, `internal/` and `cmd/` layout, small interfaces.

## Conventions to inherit

1. Commit generated contract code. CI regenerates and diffs it. Fail with a copy-paste fix command.
2. Split contract checks: hard gate for committed source vs committed output, advisory for deployed drift.
3. Pin every generator and plugin version. `buf.gen.yaml` with pinned plugins is the proto equivalent.
4. One leaf package for generated TypeScript, with no build step.
5. Enforce import boundaries in the linter. Unit-test the policy scripts.
6. Named `verify:*` jobs in a `fail-fast: false` matrix, with one fan-in gate job.
7. Path filters per module.
8. Pin toolchain versions: `packageManager`, `.nvmrc`, `go` directive, buf, golangci-lint.
9. Makefile as the single entry point: `proto`, `lint`, `test`, `test-repeat`, `cover`, `run-<svc>`.
10. `CLAUDE.md` includes `AGENTS.md`. AGENTS.md holds commands, never-edit rules, and architecture rules.
11. Explicit test pyramid with an E2E budget. Go: table-driven unit tests, `-race` in CI.
12. Pre-commit hooks that run the same checks as CI.
