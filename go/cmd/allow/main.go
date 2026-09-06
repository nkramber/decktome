// Command allow puts one email on the invite list of the deployed app, or
// takes one off (D-314, D-420). It writes the one Firestore document
// config/allowlist through the caller's own credentials, so an invite
// needs no deploy.
//
// Usage:
//
//	PROJECT_ID=my-project go run ./cmd/allow -email ann@example.com
//	PROJECT_ID=my-project go run ./cmd/allow -email ann@example.com -remove
//
// The credentials come from `gcloud auth application-default login`, or
// from the emulator when FIRESTORE_EMULATOR_HOST is set.
package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"cloud.google.com/go/firestore"

	"github.com/nkramber/mtg-deck-builder/go/internal/allowlist"
	"github.com/nkramber/mtg-deck-builder/go/internal/gcpenv"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	email := flag.String("email", "", "the email to invite")
	remove := flag.Bool("remove", false, "take the email off the list instead")
	flag.Parse()
	if *email == "" {
		return fmt.Errorf("allow: -email is required")
	}
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
	if *remove {
		if err := allowlist.Remove(ctx, client, *email); err != nil {
			return err
		}
		fmt.Printf("removed %s from %s/%s in project %s\n", allowlist.Normalize(*email), allowlist.Collection, allowlist.Doc, project)
		return nil
	}
	if err := allowlist.Add(ctx, client, *email); err != nil {
		return err
	}
	fmt.Printf("invited %s in %s/%s of project %s\n", allowlist.Normalize(*email), allowlist.Collection, allowlist.Doc, project)
	return nil
}
