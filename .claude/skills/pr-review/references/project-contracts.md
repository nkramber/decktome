# PR review: project contracts

Part of the `pr-review` skill. Load this file when the pull request changes code, tools, or CI, or the text of a contract in one of the areas below.

## Project contracts

Apply each relevant row. Record why an area does not apply when its omission can mislead a reviewer.

| Area | Required examination |
|---|---|
| Proto contract | `proto/` is the one contract. The generated code is committed and current, and `make proto-breaking` passes against `main`. |
| Errors | An error keeps its context and reaches the caller. A silent default, an empty `catch`, or a swallowed error is a defect. |
| Stored data | Firestore documents and GCS objects stay readable by the deployed code. A verdict keeps the object it names (D-635). No harvest reads `users/<uid>` (D-638). |
| Public repository | No email, personal address, key, or secret in a file, an issue, or a pull request (D-639). |
| Spend | A paid target costs money. It asks the owner first, and its guard refuses a document that holds a result (D-65). A change to a model default passes `make llm-defaults-check`. |
| MtG facts | Each card name, rule, legality, and ban-list fact has a source and a date (hard rules 4 and 7). |
| Web | The web tests run under the Node version of `.nvmrc`. A screen change keeps the look of the reference design. |
| Deploy | A merge to `main` deploys itself on Cloud Build. The change deploys and rolls back as `docs/deploy-and-rollback.md` says. |
| CI boundaries | Each action pins a full commit SHA. A value of an event enters a script through env alone. A `pull_request_target` job never runs a file of the head. |
| Documents | ASD-STE100, each cited id and path resolves, and the byte budgets hold (`make lint`). |

Do not bring back an earlier contract that a later decision superseded.
