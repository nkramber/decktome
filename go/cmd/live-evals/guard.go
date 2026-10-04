package main

import (
	"context"
	"fmt"
	"io"
	"strconv"
	"strings"
)

// protectedPrefixes and protectedFiles are the paths that a live-eval
// pull request never changes (D-1158): access, accounts, spend (D-1163), the rules
// and the deploy, the agent rules, the feedback store that holds the eval
// marks, and the live evals themselves. The
// scope rule of the prompt (D-1157) keeps a session away from them, and
// this list holds when a reader text fools the session.
var protectedPrefixes = []string{
	".claude/",
	".github/",
	"cloudbuild/",
	"docker/",
	"go/cmd/live-evals/",
	"go/internal/access/",
	"go/internal/adminsvc/",
	"go/internal/allowlist/",
	"go/internal/auth/",
	"go/internal/authblock/",
	"go/internal/feedback/",
	"go/internal/feedbacksvc/",
	"go/internal/gatekit/",
	"go/internal/invitesvc/",
	"go/internal/llm/",
	"go/internal/prooflink/",
	"go/internal/proofsvc/",
	"go/internal/ratelimit/",
	"go/internal/spendmask/",
	"go/internal/usage/",
	"go/internal/users/",
	"scripts/live-evals/",
	"web/apps/web/src/features/admin/",
	"web/apps/web/src/features/auth/",
}

var protectedFiles = map[string]bool{
	".firebaserc":                         true,
	"AGENTS.md":                           true,
	"CLAUDE.md":                           true,
	"docs/tools/live_evals_runfs.py":      true,
	"firebase.json":                       true,
	"firestore.rules":                     true,
	"proto/mtg/v1/admin_service.proto":    true,
	"proto/mtg/v1/feedback_service.proto": true,
	"proto/mtg/v1/invite_service.proto":   true,
	"proto/mtg/v1/proof_service.proto":    true,
	"scripts/live-evals.sh":               true,
	"scripts/live-evals-launchd.sh":       true,
	"start-live-evals":                    true,
}

// guarded is the answer of `guard` for a pull request that changes a
// protected path. main exits 4 on it, so the loop script holds the pull
// request for the owner.
type guarded []string

func (g guarded) Error() string { return "protected paths: " + strings.Join(g, ", ") }

// protectedOf answers each protected path of a list of changed files.
func protectedOf(files []string) []string {
	var hit []string
	for _, f := range files {
		f = strings.TrimSpace(f)
		if f == "" {
			continue
		}
		if protectedFiles[f] {
			hit = append(hit, f)
			continue
		}
		for _, p := range protectedPrefixes {
			if strings.HasPrefix(f, p) {
				hit = append(hit, f)
				break
			}
		}
	}
	return hit
}

// guard reads the changed files of a pull request on GitHub, and never
// the word of the eval session. A rename lists the old path too, so a
// move out of a protected path is held.
func guard(ctx context.Context, gh runner, pr int, out io.Writer) error {
	raw, err := gh(ctx, "gh", "api", "--paginate", "repos/{owner}/{repo}/pulls/"+strconv.Itoa(pr)+"/files",
		"--jq", ".[] | .filename, (.previous_filename // empty)")
	if err != nil {
		return fmt.Errorf("gh api pulls files: %w", err)
	}
	if hit := protectedOf(strings.Split(string(raw), "\n")); len(hit) > 0 {
		return guarded(hit)
	}
	_, err = fmt.Fprintln(out, "no protected path")
	return err
}
