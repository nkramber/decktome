// Command mark-verified lists the accounts of the deployed app whose email
// is not proved, and marks the accounts the owner names as proved (D-903).
// The invite gate refuses an email that is not proved, so the accounts
// made before that rule need this one step before the deploy of the rule.
//
// Usage:
//
//	PROJECT_ID=my-project go run ./cmd/mark-verified
//	PROJECT_ID=my-project go run ./cmd/mark-verified -uid u1,u2
//	PROJECT_ID=my-project go run ./cmd/mark-verified -uid u1,u2 -apply
//
// With no -uid it lists each account that is not proved, and whether the
// invite list holds its email. With -uid and no -apply it names what it
// would change and changes nothing. The owner confirms each person before
// the -apply run. The credentials come from
// `gcloud auth application-default login`.
package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
	"time"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go/v4"
	fbauth "firebase.google.com/go/v4/auth"
	"google.golang.org/api/iterator"

	"github.com/nkramber/decktome/go/internal/allowlist"
	"github.com/nkramber/decktome/go/internal/gcpenv"
)

// account is the part of one Firebase account this command reads.
type account struct {
	UID       string
	Email     string
	Verified  bool
	Created   time.Time
	LastLogIn time.Time
}

// accounts reads and changes the Firebase accounts.
type accounts interface {
	List(ctx context.Context) ([]account, error)
	MarkVerified(ctx context.Context, uid string) error
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	uids := flag.String("uid", "", "the comma-separated user ids to mark as proved")
	apply := flag.Bool("apply", false, "change the named accounts; without it the run changes nothing")
	flag.Parse()
	ctx := context.Background()
	project, err := gcpenv.ProjectID()
	if err != nil {
		return err
	}
	app, err := firebase.NewApp(ctx, &firebase.Config{ProjectID: project})
	if err != nil {
		return fmt.Errorf("firebase app: %w", err)
	}
	client, err := app.Auth(ctx)
	if err != nil {
		return fmt.Errorf("firebase auth: %w", err)
	}
	fs, err := firestore.NewClient(ctx, project)
	if err != nil {
		return fmt.Errorf("firestore: %w", err)
	}
	defer func() { _ = fs.Close() }()
	_, _ = fmt.Fprintf(os.Stdout, "project %s\n", project)
	return mark(ctx, os.Stdout, firebaseAccounts{client}, allowlist.FromFirestore(fs), splitIDs(*uids), *apply)
}

// mark lists the accounts that are not proved when ids is empty. Else it
// marks each named account as proved, and it refuses the whole run when
// one id names no account. A dry run changes nothing.
func mark(ctx context.Context, w io.Writer, acc accounts, list *allowlist.List, ids []string, apply bool) error {
	all, err := acc.List(ctx)
	if err != nil {
		return fmt.Errorf("list accounts: %w", err)
	}
	if len(ids) == 0 {
		var open []account
		for _, a := range all {
			if !a.Verified {
				open = append(open, a)
			}
		}
		sort.Slice(open, func(i, j int) bool { return open[i].Created.Before(open[j].Created) })
		_, _ = fmt.Fprintf(w, "%d of %d accounts are not proved\n", len(open), len(all))
		for _, a := range open {
			invited, err := list.Allowed(ctx, a.Email)
			if err != nil {
				return fmt.Errorf("read the invite list: %w", err)
			}
			_, _ = fmt.Fprintf(w, "%s\t%s\tinvited=%v\tcreated=%s\tlast=%s\n", a.UID, a.Email, invited, day(a.Created), day(a.LastLogIn))
		}
		return nil
	}
	byUID := map[string]account{}
	for _, a := range all {
		byUID[a.UID] = a
	}
	var picked []account
	for _, id := range ids {
		a, ok := byUID[id]
		if !ok {
			return fmt.Errorf("no account has the id %q, so nothing changed", id)
		}
		picked = append(picked, a)
	}
	for _, a := range picked {
		switch {
		case a.Verified:
			_, _ = fmt.Fprintf(w, "%s\t%s\talready proved\n", a.UID, a.Email)
		case !apply:
			_, _ = fmt.Fprintf(w, "%s\t%s\twould mark as proved (dry run; add -apply)\n", a.UID, a.Email)
		default:
			if err := acc.MarkVerified(ctx, a.UID); err != nil {
				return fmt.Errorf("mark %s: %w", a.UID, err)
			}
			_, _ = fmt.Fprintf(w, "%s\t%s\tmarked as proved\n", a.UID, a.Email)
		}
	}
	return nil
}

func splitIDs(s string) []string {
	var out []string
	for _, id := range strings.Split(s, ",") {
		if id = strings.TrimSpace(id); id != "" {
			out = append(out, id)
		}
	}
	return out
}

func day(t time.Time) string {
	if t.IsZero() {
		return "-"
	}
	return t.UTC().Format(time.DateOnly)
}

// firebaseAccounts is accounts over the Firebase Admin SDK.
type firebaseAccounts struct{ client *fbauth.Client }

func (f firebaseAccounts) List(ctx context.Context) ([]account, error) {
	var out []account
	it := f.client.Users(ctx, "")
	for {
		u, err := it.Next()
		if errors.Is(err, iterator.Done) {
			return out, nil
		}
		if err != nil {
			return nil, err
		}
		a := account{UID: u.UID, Email: u.Email, Verified: u.EmailVerified}
		if m := u.UserMetadata; m != nil {
			a.Created = time.UnixMilli(m.CreationTimestamp)
			a.LastLogIn = time.UnixMilli(m.LastLogInTimestamp)
		}
		out = append(out, a)
	}
}

func (f firebaseAccounts) MarkVerified(ctx context.Context, uid string) error {
	_, err := f.client.UpdateUser(ctx, uid, (&fbauth.UserToUpdate{}).EmailVerified(true))
	return err
}
