// Command users-backfill seeds the user record from what each user
// already holds (D-638).
//
// The counters count creations, and they started at zero on the day the
// record shipped. Every user who came before that undercounts. This
// reads the decks, the collections, the sessions, and the verdicts each
// user holds and writes those numbers once.
//
// It never lowers a count the record already holds, so a creation that
// lands during the run survives it. A deck the reader deleted is gone,
// so a backfilled count is a floor and not a history.
//
// Usage:
//
//	PROJECT_ID=decktome-prod go run ./cmd/users-backfill -dry
//	PROJECT_ID=decktome-prod go run ./cmd/users-backfill
//
// CAUTION: the output names users. Keep it off any shared page.
package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"os"
	"time"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"

	"github.com/nkramber/decktome/go/internal/gcpenv"
	"github.com/nkramber/decktome/go/internal/users"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	dry := flag.Bool("dry", false, "count and write nothing")
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
	fmt.Fprintf(os.Stderr, "project      %s\n", project)

	repo := users.NewRepo(client)
	// DocumentRefs and not Documents. The app writes
	// users/<uid>/<kind>/<id> and gave the parent no field until D-638,
	// so a plain list leaves an older user out.
	iter := client.Collection("users").DocumentRefs(ctx)
	seen := 0
	for {
		ref, err := iter.Next()
		if errors.Is(err, iterator.Done) {
			break
		}
		if err != nil {
			return fmt.Errorf("users: %w", err)
		}
		seen++
		counts := map[users.Counter]int64{}
		var oldest, newest time.Time

		for kind, counter := range map[string]users.Counter{
			"decks":       users.DecksCreated,
			"collections": users.CollectionsUpload,
			"sessions":    users.SessionsStarted,
		} {
			n, first, last, err := walk(ctx, ref.Collection(kind))
			if err != nil {
				return fmt.Errorf("%s of %s: %w", kind, ref.ID, err)
			}
			// A revised deck names the deck it came from, and walk counts
			// it apart.
			if kind == "decks" {
				var revisions int64
				revisions, err = countRevisions(ctx, ref.Collection(kind))
				if err != nil {
					return fmt.Errorf("revisions of %s: %w", ref.ID, err)
				}
				counts[users.DeckRevisions] = revisions
				n -= revisions
			}
			counts[counter] = n
			oldest = earlier(oldest, first)
			newest = later(newest, last)
		}

		up, down, first, last, err := verdicts(ctx, ref.Collection("feedback"))
		if err != nil {
			return fmt.Errorf("feedback of %s: %w", ref.ID, err)
		}
		counts[users.FeedbackUp], counts[users.FeedbackDown] = up, down
		oldest = earlier(oldest, first)
		newest = later(newest, last)

		fmt.Fprintf(os.Stderr, "user %s  decks %d, revisions %d, collections %d, chats %d, up %d, down %d\n",
			ref.ID, counts[users.DecksCreated], counts[users.DeckRevisions],
			counts[users.CollectionsUpload], counts[users.SessionsStarted], up, down)
		if *dry {
			continue
		}
		// The email is not here to read: it lives in Firebase Auth and on
		// the record a later turn writes. The backfill leaves it alone.
		if err := repo.Seed(ctx, ref.ID, "", counts, oldest, newest); err != nil {
			return fmt.Errorf("seed %s: %w", ref.ID, err)
		}
	}
	if *dry {
		fmt.Fprintf(os.Stderr, "%d user(s) read, and -dry wrote nothing\n", seen)
		return nil
	}
	fmt.Fprintf(os.Stderr, "%d user(s) seeded\n", seen)
	return nil
}

// walk counts the documents of one collection and names the oldest and
// the newest created_at it holds.
func walk(ctx context.Context, col *firestore.CollectionRef) (n int64, first, last time.Time, err error) {
	snaps, err := col.Documents(ctx).GetAll()
	if err != nil {
		return 0, time.Time{}, time.Time{}, err
	}
	for _, snap := range snaps {
		n++
		if t, ok := snap.Data()["created_at"].(time.Time); ok {
			first, last = earlier(first, t), later(last, t)
		}
	}
	return n, first, last, nil
}

// countRevisions counts the decks that name a deck they came from.
func countRevisions(ctx context.Context, col *firestore.CollectionRef) (int64, error) {
	snaps, err := col.Documents(ctx).GetAll()
	if err != nil {
		return 0, err
	}
	var n int64
	for _, snap := range snaps {
		if id, _ := snap.Data()["revised_from_deck_id"].(string); id != "" {
			n++
		}
	}
	return n, nil
}

// verdicts counts the thumbs up and the thumbs down of one user.
func verdicts(ctx context.Context, col *firestore.CollectionRef) (up, down int64, first, last time.Time, err error) {
	snaps, err := col.Documents(ctx).GetAll()
	if err != nil {
		return 0, 0, time.Time{}, time.Time{}, err
	}
	for _, snap := range snaps {
		switch v, _ := snap.Data()["verdict"].(string); v {
		case "up":
			up++
		case "down":
			down++
		}
		if t, ok := snap.Data()["created_at"].(time.Time); ok {
			first, last = earlier(first, t), later(last, t)
		}
	}
	return up, down, first, last, nil
}

func earlier(a, b time.Time) time.Time {
	if b.IsZero() || (!a.IsZero() && a.Before(b)) {
		return a
	}
	return b
}

func later(a, b time.Time) time.Time {
	if b.After(a) {
		return b
	}
	return a
}
