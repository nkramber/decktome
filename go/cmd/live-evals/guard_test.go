package main

import (
	"bytes"
	"context"
	"errors"
	"reflect"
	"testing"
)

// D-1158: a change of an access, rules, deploy, or live-eval path is
// held, and a change of the deck, collection, or chat code passes.
func TestProtectedOfHoldsAccessAndDeployPaths(t *testing.T) {
	files := []string{
		"go/internal/generate/fill.go",
		"go/internal/collections/parse.go",
		"web/apps/web/src/features/deck/deck-view.tsx",
		"docs/decisions.md",
		"go/internal/auth/verify.go",
		"firestore.rules",
		".github/workflows/verify.yml",
		"scripts/live-evals/eval-prompt.md",
		"scripts/live-evals.sh",
		"web/apps/web/src/features/admin/panel.tsx",
		"go/internal/users/store.go",
		"",
	}
	got := protectedOf(files)
	want := []string{"go/internal/auth/verify.go", "firestore.rules", ".github/workflows/verify.yml",
		"scripts/live-evals/eval-prompt.md", "scripts/live-evals.sh", "web/apps/web/src/features/admin/panel.tsx",
		"go/internal/users/store.go"}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("protectedOf = %v, want %v", got, want)
	}
}

// The feedback store holds the eval marks and the snapshots of D-635, so
// a steered session can not change the queue that it reads (D-1158).
func TestProtectedOfHoldsTheFeedbackStore(t *testing.T) {
	files := []string{"go/internal/feedback/evaluation.go", "go/internal/feedbacksvc/service.go",
		"proto/mtg/v1/feedback_service.proto", "web/apps/web/src/features/deck/deck-view.tsx"}
	want := files[:3]
	if got := protectedOf(files); !reflect.DeepEqual(got, want) {
		t.Fatalf("protectedOf = %v, want %v", got, want)
	}
}

// The paid-run gates and the provider cost meter are spend paths (D-1158, D-1163).
func TestProtectedOfHoldsTheSpendControls(t *testing.T) {
	files := []string{"go/internal/gatekit/spendcap.go", "go/internal/llm/prices.json",
		"go/internal/llm/roles.json", "go/internal/llm/usage.go", "go/internal/generate/generate.go"}
	want := files[:4]
	if got := protectedOf(files); !reflect.DeepEqual(got, want) {
		t.Fatalf("protectedOf = %v, want %v", got, want)
	}
}

// A prefix matches a folder, and never a longer name that shares it.
func TestProtectedOfMatchesWholeFolders(t *testing.T) {
	if got := protectedOf([]string{"go/internal/authority/x.go", "go/internal/usagestats/x.go", "scripts/live-evals-notes.md"}); len(got) != 0 {
		t.Fatalf("protectedOf = %v, want none", got)
	}
}

func TestGuardReadsTheDiffOfGitHub(t *testing.T) {
	gh := func(_ context.Context, _ string, args ...string) ([]byte, error) {
		want := []string{"api", "--paginate", "repos/{owner}/{repo}/pulls/7/files",
			"--jq", ".[] | .filename, (.previous_filename // empty)"}
		if !reflect.DeepEqual(args, want) {
			t.Fatalf("args = %v", args)
		}
		return []byte("go/internal/decks/store.go\nfirestore.rules\n"), nil
	}
	var out bytes.Buffer
	err := guard(context.Background(), gh, 7, &out)
	var g guarded
	if !errors.As(err, &g) || len(g) != 1 || g[0] != "firestore.rules" {
		t.Fatalf("guard = %v, want the rules file", err)
	}
	// A move out of a protected path lists the new path, then the old one.
	moved := func(context.Context, string, ...string) ([]byte, error) {
		return []byte("go/internal/decks/verify.go\ngo/internal/auth/verify.go\n"), nil
	}
	if err := guard(context.Background(), moved, 7, &out); !errors.As(err, &g) || g[0] != "go/internal/auth/verify.go" {
		t.Fatalf("guard of a move = %v, want the old auth path", err)
	}
	clean := func(context.Context, string, ...string) ([]byte, error) {
		return []byte("go/internal/decks/store.go\n"), nil
	}
	if err := guard(context.Background(), clean, 7, &out); err != nil {
		t.Fatalf("guard of a clean diff = %v", err)
	}
}
