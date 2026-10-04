# The author response to the review of #285

Author provider: Claude Code.

## P2-1: The guard allows changes to feedback storage and service paths

Result: full merit.

Evidence: `protectedPrefixes` and `protectedFiles` of `go/cmd/live-evals/guard.go` held no path of the feedback store. `go/internal/feedback` holds the eval marks of D-1149 and D-1156, and the queue of the live evals reads them. So a steered session can change the queue that it reads, and the guard of D-1158 then sends a ready notice. The scope rule of D-1157 puts the live evals themselves out of scope.

Correction: the guard holds `go/internal/feedback/`, `go/internal/feedbacksvc/`, and `proto/mtg/v1/feedback_service.proto`. D-1158 and the caution of `scripts/live-evals/eval-prompt.md` name the feedback store.

Regression check: `TestProtectedOfHoldsTheFeedbackStore` gives the three paths and a deck path. It fails on the old list, which returned no path. `go test ./cmd/live-evals` passes.

## P2-2: Review-provider attribution remains in pull request metadata

Result: partial merit.

Rules: hard rule 6 forbids AI-attribution text in a pull request, a branch name, a commit message, and a comment. D-811 lets the review record and the `Author provider` line of the hand-off name a provider. D-811 also says that hard rule 6 still covers each commit, body, branch, and comment.

The commit message. Full merit. The squash message of this repo holds each commit message (`COMMIT_MESSAGES`), so the text "Codex P2-1 on #285" goes to `main`. #287 cites its finding with no provider name, and this commit now uses that form. The owner permitted a rebase onto `c8a70af` with a new message. `git log --format=%B c8a70af..HEAD` holds no provider name.

The body. Full merit. The `## Review` section of `.github/pull_request_template.md` asks for the state of the review record. It does not need the provider name. The body now names "the review record" and "the review round" in the three lines that named Codex.

The comments. No merit. The three comments of the author name the command `make codex-review PR=285`. They name no reviewer. The command is a technical name of hard rule 10 of `CLAUDE.md`. A command name does not make a provider the author of the work.

Regression check: read the body with `gh pr view 285 --json body`, and run `git log --format=%B c8a70af..HEAD`. Neither holds "Codex", "OpenAI", or a co-author line. The body names "Claude" only in the paths `CLAUDE.md` and `.claude/`.

## P2-3: The guard allows changes to paid-run controls

Result: full merit.

Evidence: `protectedPrefixes` of `go/cmd/live-evals/guard.go` held neither path. `go/internal/gatekit/` holds the permission and the spend cap of each paid target. `go/internal/llm/` holds the cost meter, `prices.json`, and `roles.json`. D-1158 names the spend as a protected area.

Correction: the guard holds `go/internal/gatekit/` and `go/internal/llm/`. The caution of `scripts/live-evals/eval-prompt.md` names the spend controls. The user cap of D-1109 sits in `go/cmd/api/main.go` and `go/internal/agentsvc/service.go`. The owner chose to keep those files open, because a product fix often changes them (D-1162).

Regression check: `TestProtectedOfHoldsTheSpendControls` gives four spend paths and a deck path. It fails on the old list, which returned no path. `go test ./cmd/live-evals` passes.
