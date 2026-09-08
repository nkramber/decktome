// Command feedback reads the newest verdicts of every user, for the
// owner who answers for the app (D-596).
//
// The documents sit under each user, at users/<uid>/feedback/<id>. One
// collection group query reads them all, so no caller walks the user
// list and no second store holds a copy.
//
// Usage:
//
//	PROJECT_ID=decktome-prod go run ./cmd/feedback
//	PROJECT_ID=decktome-prod go run ./cmd/feedback -verdict up -limit 100
//	PROJECT_ID=decktome-prod go run ./cmd/feedback -json
//
// CAUTION: the output holds what a reader wrote. Keep it off any shared
// page.
package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"

	"cloud.google.com/go/firestore"

	"github.com/nkramber/decktome/go/internal/feedback"
	"github.com/nkramber/decktome/go/internal/gcpenv"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	verdict := flag.String("verdict", "down", `the verdict to read: "down", "up", or "" for both`)
	limit := flag.Int("limit", 50, "how many verdicts to read, newest first")
	asJSON := flag.Bool("json", false, "print the verdicts as JSON")
	id := flag.String("id", "", "read one verdict by its id, over every user")
	flag.Parse()

	ctx := context.Background()
	project, err := gcpenv.ProjectID()
	if err != nil {
		return err
	}
	client, err := firestore.NewClient(ctx, project)
	if err != nil {
		return fmt.Errorf("firestore: %w", err)
	}
	defer func() { _ = client.Close() }()

	repo := feedback.NewRepo(client)
	var items []feedback.Item
	if *id != "" {
		// One id needs no index, so this reads without the collection
		// group query (D-600).
		item, uid, err := repo.Find(ctx, *id)
		if err != nil {
			return fmt.Errorf("feedback %s: %w", *id, err)
		}
		fmt.Printf("feedback     %s\nuser         %s\n", *id, uid)
		items = []feedback.Item{item}
	} else {
		var err error
		items, err = repo.Down(ctx, *verdict, *limit)
		if err != nil {
			return err
		}
	}
	if *asJSON {
		enc := json.NewEncoder(os.Stdout)
		enc.SetIndent("", "  ")
		return enc.Encode(items)
	}
	fmt.Printf("%d verdict(s) in %s\n\n", len(items), project)
	for _, it := range items {
		fmt.Printf("%s  %s/%s  user %s\n", it.CreatedAt.Format("2006-01-02 15:04"), it.Kind, it.Verdict, it.UID)
		if it.SessionID != "" {
			fmt.Printf("  session   %s\n", it.SessionID)
		}
		if it.DeckID != "" {
			fmt.Printf("  deck      %s\n", it.DeckID)
		}
		if it.QuestionText != "" {
			fmt.Printf("  asked     %s\n", it.QuestionText)
		}
		if it.AnswerText != "" {
			fmt.Printf("  answered  %s\n", it.AnswerText)
		}
		if len(it.Reasons) > 0 {
			fmt.Printf("  reasons   %s\n", strings.Join(it.Reasons, ", "))
		}
		if it.Text != "" {
			fmt.Printf("  said      %s\n", it.Text)
		}
		fmt.Println()
	}
	return nil
}
