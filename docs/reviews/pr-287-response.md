# The author response to the review of #287

Author provider: Claude Code.

## P2-1: Test does not check the registered flag default

Result: full merit.

Evidence: `TestDefaultPoolIsTheRuleOfTheWebApp` parsed the constant `defaultPool`. When the flag line of `flagSet` named `owned-first`, the test still passed.

Correction:

- `go/cmd/api-build/main.go`: the new function `register` declares each flag on a given set. `flagSet` calls it with `flag.CommandLine`, then parses the command line.
- `go/cmd/api-build/answers_test.go`: the test registers the flags on a new set, parses no argument, and reads the pool rule of that set (D-1161).
- `docs/design-roadmap.md`: the gate of PR-129 names the registered default.

Regression checks:

- With the registered default set to `owned-first`, `go test ./cmd/api-build -run TestDefaultPoolIsTheRuleOfTheWebApp` fails: "the default pool rule is POOL_RULE_OWNED_FIRST, want POOL_RULE_OWNED_ONLY".
- With the default restored to `defaultPool`, `go test ./cmd/api-build` passes.
