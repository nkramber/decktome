package feedback

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

// Pending is one down verdict that no live eval read (D-1149).
type Pending struct {
	UID       string
	ID        string
	Kind      string
	SessionID string
	DeckID    string
	CreatedAt time.Time
}

// evalFields are the flat fields UnevaluatedDown reads. It never reads a
// snapshot of a session or a deck.
var evalFields = []string{"uid", "kind", "verdict", "session_id", "deck_id", "created_at", "has_been_evaluated"}

// UnevaluatedDown lists every down verdict of every user that holds no
// true eval mark, oldest first (D-1149). The index "verdict ascending,
// created_at descending" serves the query, as it serves Since.
func (r *Repo) UnevaluatedDown(ctx context.Context) ([]Pending, error) {
	it := r.client.CollectionGroup("feedback").
		Where("verdict", "==", "down").
		OrderBy("created_at", firestore.Desc).
		Select(evalFields...).Documents(ctx)
	defer it.Stop()
	var out []Pending
	for {
		snap, err := it.Next()
		if errors.Is(err, iterator.Done) {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("feedback: %w", err)
		}
		var s stored
		if err := snap.DataTo(&s); err != nil {
			return nil, fmt.Errorf("feedback %s: %w", snap.Ref.ID, err)
		}
		if s.HasBeenEvaluated {
			continue
		}
		// A document written before the uid field reads its user from the
		// path: users/<uid>/feedback/<id>.
		if s.UID == "" {
			s.UID = snap.Ref.Parent.Parent.ID
		}
		out = append(out, Pending{UID: s.UID, ID: snap.Ref.ID, Kind: s.Kind, SessionID: s.SessionID,
			DeckID: s.DeckID, CreatedAt: s.CreatedAt})
	}
	sort.SliceStable(out, func(i, j int) bool { return out[i].CreatedAt.Before(out[j].CreatedAt) })
	return out, nil
}

// MarkEvaluated records that a live eval read one verdict (D-1149). It
// writes the two eval fields alone. A verdict that no longer exists
// answers ErrNotFound.
func (r *Repo) MarkEvaluated(ctx context.Context, uid, id string, at time.Time) error {
	if uid == "" || id == "" {
		return errors.New("feedback: a mark needs a user and a verdict id")
	}
	_, err := r.col(uid).Doc(id).Update(ctx, []firestore.Update{
		{Path: "has_been_evaluated", Value: true},
		{Path: "evaluated_at", Value: at.UTC()},
	})
	if status.Code(err) == codes.NotFound {
		return ErrNotFound
	}
	return err
}
