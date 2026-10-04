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

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
)

// Pending is one down verdict (D-1149) or one general note (D-1156) that
// no live eval read.
type Pending struct {
	UID       string
	ID        string
	Kind      string
	SessionID string
	DeckID    string
	Screen    string
	CreatedAt time.Time
}

// evalFields are the flat fields the two queues read. They never read a
// snapshot of a session or a deck.
var evalFields = []string{"uid", "kind", "verdict", "session_id", "deck_id", "screen", "created_at", "has_been_evaluated"}

// UnevaluatedDown lists every down verdict of every user that holds no
// true eval mark, oldest first (D-1149). The index "verdict ascending,
// created_at descending" serves the query, as it serves Since.
func (r *Repo) UnevaluatedDown(ctx context.Context) ([]Pending, error) {
	return unevaluated(ctx, r.client.CollectionGroup("feedback").
		Where("verdict", "==", "down").
		OrderBy("created_at", firestore.Desc))
}

// UnevaluatedNotes lists every general note of every user that holds no
// true eval mark, oldest first (D-1156). A note stores the verdict "".
// The index "kind ascending, verdict ascending, created_at descending"
// serves the query.
func (r *Repo) UnevaluatedNotes(ctx context.Context) ([]Pending, error) {
	return unevaluated(ctx, r.client.CollectionGroup("feedback").
		Where("kind", "==", KindName(mtgv1.FeedbackKind_FEEDBACK_KIND_GENERAL)).
		Where("verdict", "==", "").
		OrderBy("created_at", firestore.Desc))
}

func unevaluated(ctx context.Context, q firestore.Query) ([]Pending, error) {
	it := q.Select(evalFields...).Documents(ctx)
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
			DeckID: s.DeckID, Screen: s.Screen, CreatedAt: s.CreatedAt})
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
