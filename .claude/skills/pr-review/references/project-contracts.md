# PR review: the project contracts

Part of the `pr-review` skill. Load this file when the pull request changes code, tools, or CI, or the text of a contract below.

## The contracts

Apply each row that the change touches. Record why a row does not apply when its absence can mislead the owner.

| Area | Examine |
|---|---|
| The protobuf contract | `proto/` is the one contract. `go/gen/` and `web/packages/api-client/src/gen/` come from `make proto` alone, and CI fails on stale output. |
| The model layer | No model id at a call site. Each model comes from the role layer of `go/internal/llm`. A change of `roles.json` or `prices.json` needs the owner. |
| Legality | No ban list, rotation date, or Game Changers list in code or in a prompt. Each card that the model names passes the rules engine before the user sees it. Load the `mtg-corpus` skill. |
| User data | Log ids, never PII or a raw prompt. No harvest reads `users/<uid>` (D-638). A verdict keeps the object that it names (D-635). |
| Money | A paid target has a spend guard and an overwrite guard (D-65). A new call to a provider names its cost. The reviewer never runs a paid target. |
| Go code | `ctx` is the first parameter. Use `errors.Is` and `errors.As`, and table tests. Accept interfaces, and return structs. |
| The web app | Check the phone width and the desktop width of each changed screen. A live check reads the deployed chunk, and not the screen alone. |
| CI | Each action pins a full commit SHA. A value of the event enters a script through `env`, and never through an expression. A `pull_request_target` job never runs a file of the head (D-816). |
| The deploy | A merge to `main` deploys itself (D-584). Read `docs/deploy-and-rollback.md` for a change of `cloudbuild/`, `firebase.json`, or a Dockerfile. |
| Documents | Each document follows ASD-STE100, and `make ste-check` and `make ref-check` pass. |
