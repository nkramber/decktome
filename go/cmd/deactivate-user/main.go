// Command deactivate-user closes one account of the deployed app, and it
// keeps every record of the user (REV-036, D-941). It marks the user
// record closed, so the API refuses the user and the share links answer
// NotFound. Then it deletes the Firebase Auth user, so no new token
// comes. It reads and prints the plan alone, unless -confirm is set.
//
// Usage:
//
//	PROJECT_ID=my-project go run ./cmd/deactivate-user -uid abc123
//	PROJECT_ID=my-project go run ./cmd/deactivate-user -uid abc123 -confirm
//
// The credentials come from `gcloud auth application-default login`, or
// from the emulators when FIRESTORE_EMULATOR_HOST and
// FIREBASE_AUTH_EMULATOR_HOST are set.
package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"time"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go/v4"
	fbauth "firebase.google.com/go/v4/auth"

	"github.com/nkramber/decktome/go/internal/gcpenv"
	"github.com/nkramber/decktome/go/internal/gzstore"
	"github.com/nkramber/decktome/go/internal/users"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	uid := flag.String("uid", "", "the user id to close")
	confirm := flag.Bool("confirm", false, "close the account; without it the command prints the plan alone")
	flag.Parse()
	ctx := context.Background()
	project, err := gcpenv.ProjectID()
	if err != nil {
		return err
	}
	fs, err := firestore.NewClient(ctx, project)
	if err != nil {
		return fmt.Errorf("firestore: %w", err)
	}
	defer func() { _ = fs.Close() }()
	app, err := firebase.NewApp(ctx, &firebase.Config{ProjectID: project})
	if err != nil {
		return fmt.Errorf("firebase: %w", err)
	}
	ac, err := app.Auth(ctx)
	if err != nil {
		return fmt.Errorf("firebase auth: %w", err)
	}
	fmt.Fprintf(os.Stdout, "project %s\n", project)
	return deactivate(ctx, users.NewRepo(fs), authUsers{ac}, *uid, *confirm, time.Now(), os.Stdout)
}

// Records reads and closes the user record.
type Records interface {
	Get(ctx context.Context, uid string) (users.Record, error)
	Deactivate(ctx context.Context, uid string, at time.Time) (time.Time, error)
}

// AuthUsers deletes the Firebase Auth user. A user that is gone already
// answers ErrNoAuthUser.
type AuthUsers interface {
	Delete(ctx context.Context, uid string) error
}

// ErrNoAuthUser is the answer for an Auth user that does not exist.
var ErrNoAuthUser = errors.New("no Firebase Auth user has that id")

type authUsers struct{ c *fbauth.Client }

func (a authUsers) Delete(ctx context.Context, uid string) error {
	err := a.c.DeleteUser(ctx, uid)
	if fbauth.IsUserNotFound(err) {
		return ErrNoAuthUser
	}
	return err
}

// deactivate prints the plan, and with confirm it closes the record
// first, so the API refuses the user before the Auth delete. It deletes
// no record of the user (D-941).
func deactivate(ctx context.Context, recs Records, auth AuthUsers, uid string, confirm bool, now time.Time, w io.Writer) error {
	if uid == "" {
		return errors.New("deactivate-user: -uid is required")
	}
	if !gzstore.ValidID(uid) {
		return fmt.Errorf("deactivate-user: %q is no valid user id", uid)
	}
	rec, err := recs.Get(ctx, uid)
	if err != nil {
		return err
	}
	switch {
	case rec.DeactivatedAt != nil:
		fmt.Fprintf(w, "user %s is closed since %s\n", uid, rec.DeactivatedAt.UTC().Format(time.RFC3339))
	case rec.CreatedAt.IsZero():
		fmt.Fprintf(w, "user %s has no user record. The close writes one that holds the mark\n", uid)
	default:
		fmt.Fprintf(w, "user %s, %s, created %s\n", uid, rec.Email, rec.CreatedAt.UTC().Format("2006-01-02"))
	}
	fmt.Fprintln(w, "plan: mark users/"+uid+" closed, then delete the Firebase Auth user. Every record of the user stays.")
	if !confirm {
		fmt.Fprintln(w, "dry run: nothing changed. Add -confirm to close the account.")
		return nil
	}
	at, err := recs.Deactivate(ctx, uid, now)
	if err != nil {
		return fmt.Errorf("the record was not marked, and nothing changed: %w", err)
	}
	fmt.Fprintf(w, "marked closed at %s. The API refuses the user within %s\n", at.Format(time.RFC3339), users.ClosedTTL)
	switch err := auth.Delete(ctx, uid); {
	case errors.Is(err, ErrNoAuthUser):
		fmt.Fprintln(w, "the Firebase Auth user was gone already")
	case err != nil:
		return fmt.Errorf("the record is closed, and the Auth delete failed, so run the command again: %w", err)
	default:
		fmt.Fprintln(w, "deleted the Firebase Auth user")
	}
	fmt.Fprintln(w, "note: the email stays on the invite list. Run make disallow to stop a new sign-up with it.")
	return nil
}
