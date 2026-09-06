// Package usage keeps the monthly spend of each user, for the cap of
// D-421. One Firestore document per user and month,
// users/<uid>/usage/<YYYY-MM>, holds the cost in USD and the calls. The
// API adds the cost of every turn to it and reads it before a turn. A
// session delete leaves the ledger alone, so a user can not reset the
// cap by deleting chats.
package usage

import (
	"context"
	"fmt"
	"time"

	"cloud.google.com/go/firestore"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Repo is the ledger over Firestore.
type Repo struct {
	client *firestore.Client
}

// NewRepo wraps a Firestore client.
func NewRepo(client *firestore.Client) *Repo { return &Repo{client: client} }

// Month names the ledger document of a time, in UTC: 2026-09.
func Month(t time.Time) string { return t.UTC().Format("2006-01") }

// ResetDate is the first day of the month after the one named, the day
// the cap resets (D-421).
func ResetDate(month string) (string, error) {
	t, err := time.Parse("2006-01", month)
	if err != nil {
		return "", fmt.Errorf("usage: bad month %q: %w", month, err)
	}
	return t.AddDate(0, 1, 0).Format("2006-01-02"), nil
}

type stored struct {
	CostUSD   float64   `firestore:"cost_usd"`
	Calls     int64     `firestore:"calls"`
	UpdatedAt time.Time `firestore:"updated_at"`
}

func (r *Repo) doc(uid, month string) *firestore.DocumentRef {
	return r.client.Collection("users").Doc(uid).Collection("usage").Doc(month)
}

// Spent reads the cost a user spent in a month, in USD. No document is
// zero.
func (r *Repo) Spent(ctx context.Context, uid, month string) (float64, error) {
	snap, err := r.doc(uid, month).Get(ctx)
	if err != nil {
		if status.Code(err) == codes.NotFound {
			return 0, nil
		}
		return 0, err
	}
	var s stored
	if err := snap.DataTo(&s); err != nil {
		return 0, fmt.Errorf("usage: %w", err)
	}
	return s.CostUSD, nil
}

// Add sums the cost and the calls of one turn onto the month. The write
// is an increment, so two turns that land together both count.
func (r *Repo) Add(ctx context.Context, uid, month string, costUSD float64, calls int64, at time.Time) error {
	if costUSD <= 0 && calls <= 0 {
		return nil
	}
	_, err := r.doc(uid, month).Set(ctx, map[string]any{
		"cost_usd":   firestore.Increment(costUSD),
		"calls":      firestore.Increment(calls),
		"updated_at": at.UTC(),
	}, firestore.MergeAll)
	return err
}
