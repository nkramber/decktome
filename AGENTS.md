# decktome - AGENTS

Go + Protobuf + TypeScript monorepo. Read `CLAUDE.md` for the owner's rules and `docs/design-roadmap.md` for the plan.

## Layout

- `proto/` - the one contract (buf module `mtg.v1`). Change it here only.
- `go/` - one Go module `github.com/nkramber/decktome/go`. `cmd/api` (Connect-RPC API, :8080 default, :8090 under `make dev`), `cmd/worker` (jobs), `internal/*` (packages), `gen/` (generated, committed). Tools: `cmd/candidates-review` (PR-6 gate document), `cmd/questions-gate` (PR-7 gate run, costs money), `cmd/questions-eval` (scores a gate run, costs money), `cmd/m5-sheet` and `cmd/m5-report` (the hand-scoring sheet), `cmd/tune-check` (accept rules for one loop iteration, free), `cmd/deck-gate` (PR-8 gate document, costs money), `cmd/bracket-gate` (PR-14A gate document, costs money), `cmd/revise-gate` (PR-12B gate document, costs money), `cmd/chat-probe`, `cmd/generate-probe`, and `cmd/summary-judge` (probes, each costs money), `cmd/quality-gate` (PR-14B gate document, free, and its `-judge` lane costs money), and `cmd/eval` (PR-15: compare, check, baseline, and import are free, and sweep runs the paid suites under a cap). Shared packages for the tools: `internal/gatekit` (the spend and overwrite guards), `internal/evalrun` (the run files of PR-15), and `internal/auth` (the Firebase ID-token interceptor, CORS).
- `web/` - pnpm workspace. `apps/web` (React 19 + Vite 7, dev port 5180), `packages/api-client` (generated TypeScript, committed).
- `docs/` - design roadmap, decisions, open questions, reference notes, `SESSION-HANDOFF.md` (the resume point), `setup.md`, and `setup-second-mac.md`.
- `scripts/` - `doctor.sh`, `dev.sh`, `check-llm-defaults.sh`, `autotune.sh` (the tuning loop, needs `AUTOTUNE_ALLOW_UNATTENDED=1`), `autotune-fix.sh` (the fixer step the loop calls). One-off tools only. Not a source of design patterns.

## Commands

Use the Makefile. It applies the pinned versions.

```bash
make doctor        # check tools
make proto         # regenerate Go + TS from proto/ (commit the output)
make lint          # go vet, golangci-lint, eslint, tsc, and the STE check
make ste-check     # the STE check alone, free, part of make lint
make ref-check     # every cited id and path resolves, free, part of make lint (D-753)
make test          # go test -race, vitest
make test-smoke    # live LLM smoke test, reads .env (fails when .env is absent)
make llm-defaults-check   # warn on a roles.json or prices.json change
make proto-breaking       # buf breaking against main
make cover         # Go coverage summary
make build         # Go binaries + web bundle
make dev           # emulators (:8281, :9199) + fake GCS (:4443) + api (:8090) + worker + web (:5180)
make dev-seed      # one-shot card snapshot refresh (needs make dev)
make eval-check    # compare every eval baseline with its newest run, free (PR-15)
make deck-gate-dry # build every deck gate shortlist over the trimmed snapshot, free (D-521)
make allow EMAIL=... PROJECT_ID=...  # invite one email to the deployed app (D-420)
make quality-gate  # the PR-14B gate document from the local meta store, free
make meta-refresh  # read the deck list sources into the meta store, network, free
make themes-check  # theme slugs and the commander ranking against the snapshot
make verify        # every check of the verify workflow, on this machine, free (D-578)
make where         # the branch, the tree, and the state of the pull request of the branch
make hooks         # install the pre-commit hook that refuses a commit on main (D-585)
make pr-check      # the body and the diff of a pull request against D-747
make lifecycle-check  # the skill frontmatter and the wiring, part of make lint (D-748)
make context-budget   # the byte limits of the start read and the lists of paid targets (D-749)
make pipefail-check   # each piped recipe of the Makefile sets pipefail, part of make lint (F-160)
make ruleset-check    # the live ruleset of main against .github/rulesets (D-828)
```

Fifteen targets and two loop scripts spend money: `make codex-review`, `make questions-gate`, `make questions-eval`, `make eval-calibrate`, `make deck-gate`, `make bracket-gate`, `make sixty-gate`, `make revise-gate`, `make chat-probe`, `make generate-probe`, `make summary-judge`, `make quality-judge`, `make test-smoke`, `make feedback-triage`, `make api-build`, `scripts/autotune.sh`, and `scripts/feedback-loop.sh`. Ask the owner before each run. `docs/reference/paid-targets.md` holds the cost, the flags, and the guards of each one.

CI runs `verify` on each pull request (D-639). A change of documents alone skips the six heavy jobs when the code under it passed (D-818). The `pr-contract` workflow reads the body and the diff of each pull request against D-747. The `review-gate` workflow reads the review record.

The ruleset of `main` requires `review-gate`, `pr-contract`, and each pull request job of `verify` (D-815, D-828). `.github/rulesets/` holds that ruleset. A weekly schedule runs govulncheck alone (D-305).

## Review

Codex reviews each pull request after the Gitar pass (D-811). The author session starts the review with `make codex-review PR=<number>` (D-823). Load `.claude/skills/pr-review/SKILL.md` and `.claude/skills/one-pr-one-session/SKILL.md` before the review. The review writes `docs/reviews/pr-<number>.md` and updates the hand-off in one commit (D-827). The `review-gate` check reads its verdict and its head. A later commit of documents alone keeps the approval, and Gitar still reviews it (D-837).

The reviewer never replies to Gitar, never merges, and never pushes to `main`. The third open round of one finding stops the loop, and the owner decides (D-826).

## Rules

- Never edit `go/gen/` or `web/packages/api-client/src/gen/` by hand. Run `make proto`. CI fails when the committed output is stale.
- No model id string at a call site. Models come from the role layer (`go/internal/llm`, PR-10).
- No ban list, rotation date, or Game Changers list in code or prompts. Legality comes from the card database. One exception: `internal/rules/companion_bans.json` holds the companion-only ban that Scryfall can not express (F-18, Lutri), with a source and a date.
- Every card the model names passes the rules engine before the user sees it.
- Log ids, never PII or raw prompts.
- Table-driven tests. `ctx` is the first parameter. `errors.Is` / `errors.As`. Accept interfaces, return structs.
- One concern per pull request. Squash merge into `main`, the one merge method that the repository allows. After the Codex approval, the author session asks the owner with a summary in four sections inside the question: What, How, CI, and Codex review. It turns on the auto-merge after the owner confirms (D-828, D-834, D-836).
- One pull request per clean session. The pull request carries its own documents and hand-off, and no pull request records an earlier merge. Load `.claude/skills/one-pr-one-session/SKILL.md` for all work on a pull request (D-746 to D-748).
- No AI-attribution text in any PR, branch name, commit message, or comment.
- Write docs and skills in ASD-STE100. Run `make ste-check` before you commit a `.md` file. `make lint` and CI run it too (D-264).
- Cite an id that a register defines, and a path that exists. `make ref-check` fails on either one (D-753).
