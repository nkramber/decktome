# The author response to the review of #285

Author provider: Claude Code.

## P2-1: The guard allows changes to feedback storage and service paths

Result: full merit.

Evidence: `protectedPrefixes` and `protectedFiles` of `go/cmd/live-evals/guard.go` held no path of the feedback store. `go/internal/feedback` holds the eval marks of D-1149 and D-1156, and the queue of the live evals reads them. So a steered session can change the queue that it reads, and the guard of D-1158 then sends a ready notice. The scope rule of D-1157 puts the live evals themselves out of scope.

Correction: the guard holds `go/internal/feedback/`, `go/internal/feedbacksvc/`, and `proto/mtg/v1/feedback_service.proto`. D-1158 and the caution of `scripts/live-evals/eval-prompt.md` name the feedback store.

Regression check: `TestProtectedOfHoldsTheFeedbackStore` gives the three paths and a deck path. It fails on the old list, which returned no path. `go test ./cmd/live-evals` passes.
