package decks

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"time"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Pending is one stored deck that no live eval read (D-1132).
type Pending struct {
	UID         string
	ID          string
	SessionID   string
	Name        string
	RevisedFrom string
	Imported    bool
	CreatedAt   time.Time
}

// evalFields are the flat fields Unevaluated reads. It never inflates a
// deck.
var evalFields = []string{"session_id", "name", "created_at", "imported", "revised_from_deck_id", "has_been_evaluated"}

// Unevaluated lists every deck of every user that holds no true eval
// mark, oldest first. A deck written before D-1132 holds no mark, and it
// counts as unread: Firestore can not query an absent field, so the list
// reads the flat fields of each deck. It lists the user documents by
// reference alone, so it never reads the user record (D-638).
func (r *Repo) Unevaluated(ctx context.Context) ([]Pending, error) {
	var out []Pending
	users := r.client.Collection("users").DocumentRefs(ctx)
	for {
		user, err := users.Next()
		if errors.Is(err, iterator.Done) {
			break
		}
		if err != nil {
			return nil, err
		}
		it := r.col(user.ID).Select(evalFields...).Documents(ctx)
		for {
			snap, err := it.Next()
			if errors.Is(err, iterator.Done) {
				break
			}
			if err != nil {
				it.Stop()
				return nil, err
			}
			var sd storedDeck
			if err := snap.DataTo(&sd); err != nil {
				it.Stop()
				return nil, fmt.Errorf("deck %s/%s: %w", user.ID, snap.Ref.ID, err)
			}
			if sd.HasBeenEvaluated {
				continue
			}
			out = append(out, Pending{
				UID: user.ID, ID: snap.Ref.ID, SessionID: sd.SessionID, Name: sd.Name,
				RevisedFrom: sd.RevisedFromDeckID, Imported: sd.Imported, CreatedAt: sd.CreatedAt,
			})
		}
		it.Stop()
	}
	sort.SliceStable(out, func(i, j int) bool { return out[i].CreatedAt.Before(out[j].CreatedAt) })
	return out, nil
}

// MarkEvaluated records that a live eval read one deck (D-1132). It
// writes the two eval fields alone, so it never races a rename or a
// share. A deck that the user deleted answers ErrNotFound.
func (r *Repo) MarkEvaluated(ctx context.Context, uid, id string, at time.Time) error {
	if uid == "" || id == "" {
		return errors.New("decks: a mark needs a user and a deck id")
	}
	_, err := r.doc(uid, id).Update(ctx, []firestore.Update{
		{Path: "has_been_evaluated", Value: true},
		{Path: "evaluated_at", Value: at.UTC()},
	})
	if status.Code(err) == codes.NotFound {
		return ErrNotFound
	}
	return err
}

// ClearEvaluated removes the eval mark of one deck, so the deck waits for
// a live eval again (D-1168). The owner runs it after a fault of the
// harness consumed the mark. A deck that the user deleted answers
// ErrNotFound.
func (r *Repo) ClearEvaluated(ctx context.Context, uid, id string) error {
	if uid == "" || id == "" {
		return errors.New("decks: an unmark needs a user and a deck id")
	}
	_, err := r.doc(uid, id).Update(ctx, []firestore.Update{
		{Path: "has_been_evaluated", Value: false},
		{Path: "evaluated_at", Value: firestore.Delete},
	})
	if status.Code(err) == codes.NotFound {
		return ErrNotFound
	}
	return err
}
