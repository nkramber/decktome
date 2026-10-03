// Command grant-admin sets the custom claim admin: true on one Firebase
// account, or takes it off (D-1076). The claim opens the admin screen,
// and it works after the next sign-in of the account. Other custom
// claims of the account stay.
//
// Usage:
//
//	PROJECT_ID=my-project go run ./cmd/grant-admin -email ann@example.com
//	PROJECT_ID=my-project go run ./cmd/grant-admin -email ann@example.com -remove
//
// The credentials come from `gcloud auth application-default login`.
package main

import (
	"context"
	"flag"
	"fmt"
	"maps"
	"os"

	firebase "firebase.google.com/go/v4"

	"github.com/nkramber/decktome/go/internal/auth"
	"github.com/nkramber/decktome/go/internal/gcpenv"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	email := flag.String("email", "", "the email of the account")
	remove := flag.Bool("remove", false, "take the claim off instead")
	flag.Parse()
	if *email == "" {
		return fmt.Errorf("grant-admin: -email is required")
	}
	ctx := context.Background()
	project, err := gcpenv.ProjectID()
	if err != nil {
		return err
	}
	app, err := firebase.NewApp(ctx, &firebase.Config{ProjectID: project})
	if err != nil {
		return fmt.Errorf("firebase: %w", err)
	}
	client, err := app.Auth(ctx)
	if err != nil {
		return fmt.Errorf("firebase auth: %w", err)
	}
	user, err := client.GetUserByEmail(ctx, *email)
	if err != nil {
		return fmt.Errorf("grant-admin: %w", err)
	}
	claims := withAdmin(user.CustomClaims, !*remove)
	if err := client.SetCustomUserClaims(ctx, user.UID, claims); err != nil {
		return fmt.Errorf("grant-admin: %w", err)
	}
	state := "set"
	if *remove {
		state = "removed"
	}
	fmt.Printf("%s the admin claim on uid %s in project %s. It works after the next sign-in.\n", state, user.UID, project)
	return nil
}

// withAdmin answers the claims with the admin claim set or removed, and
// every other claim kept.
func withAdmin(claims map[string]any, admin bool) map[string]any {
	out := maps.Clone(claims)
	if out == nil {
		out = map[string]any{}
	}
	if admin {
		out[auth.AdminClaim] = true
	} else {
		delete(out, auth.AdminClaim)
	}
	return out
}
