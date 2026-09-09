// Package users keeps one record per user at users/<uid> (D-638).
//
// That document was a missing parent until now: the app wrote
// users/<uid>/sessions/<id> and gave the parent no field of its own, so
// a plain list left it out. The record makes it real, and it answers
// who a user is and what they have done without a read of every child
// collection.
//
// Every counter counts a creation and never a current holding. A reader
// deletes a deck, and the count of decks they ever made does not fall.
// Count the children when the current number is the question.
//
// The server writes this record and the client never does. The email
// comes from the verified token (D-596), and the Firestore rules deny
// every client read.
package users

import (
	"context"
	"errors"
	"time"

	"cloud.google.com/go/firestore"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// schemaVersion names the shape of the stored document.
const schemaVersion = 1

// Counter names one field a creation raises.
type Counter string

// The counters. Each one counts a creation, for the life of the user.
const (
	DecksCreated      Counter = "total_decks_created"
	DeckRevisions     Counter = "total_deck_revisions"
	CollectionsUpload Counter = "total_collections_uploaded"
	SessionsStarted   Counter = "total_sessions_started"
	FeedbackUp        Counter = "total_feedback_up"
	FeedbackDown      Counter = "total_feedback_down"
)

// counters lists every field the backfill and the reader know. A field
// off this list is not a counter, and Note refuses it.
var counters = map[Counter]bool{
	DecksCreated: true, DeckRevisions: true, CollectionsUpload: true,
	SessionsStarted: true, FeedbackUp: true, FeedbackDown: true,
}

// ErrBadCounter reports a field that names no counter.
var ErrBadCounter = errors.New("users: that field is no counter")

// Record is the user document.
type Record struct {
	Schema int64 `firestore:"schema"`
	// Email is the address of the verified token, and never a word of
	// the client (D-638). It is the third home of an address, beside
	// Firebase Auth and the invite list. No harvest reads this record.
	Email string `firestore:"email"`
	// CreatedAt is the first time the app saw the user, and one write
	// sets it for good.
	CreatedAt time.Time `firestore:"created_at"`
	// LastSeenAt is the newest creation the user made. A read of a page
	// does not move it, so it reads as the last time they did something.
	LastSeenAt time.Time `firestore:"last_seen_at"`

	DecksCreated      int64 `firestore:"total_decks_created"`
	DeckRevisions     int64 `firestore:"total_deck_revisions"`
	CollectionsUpload int64 `firestore:"total_collections_uploaded"`
	SessionsStarted   int64 `firestore:"total_sessions_started"`
	FeedbackUp        int64 `firestore:"total_feedback_up"`
	FeedbackDown      int64 `firestore:"total_feedback_down"`
}

// Repo stores the records. The caller owns the client.
type Repo struct {
	client *firestore.Client
}

// NewRepo wraps a Firestore client.
func NewRepo(client *firestore.Client) *Repo { return &Repo{client: client} }

func (r *Repo) doc(uid string) *firestore.DocumentRef {
	return r.client.Collection("users").Doc(uid)
}

// Note records one creation: it raises the counter, writes the email and
// the time, and sets the creation date on the first write alone.
//
// It makes the document with Create first and ignores an answer that
// says the document exists. Create writes CreatedAt exactly once, and no
// transaction and no read stand between the caller and the write.
func (r *Repo) Note(ctx context.Context, uid, email string, c Counter, at time.Time) error {
	if uid == "" {
		return errors.New("users: a record needs a user")
	}
	if !counters[c] {
		return ErrBadCounter
	}
	at = at.UTC()
	first := map[string]any{
		"schema":       int64(schemaVersion),
		"created_at":   at,
		"last_seen_at": at,
	}
	for name := range counters {
		first[string(name)] = int64(0)
	}
	if email != "" {
		first["email"] = email
	}
	// A document that exists keeps its creation date, and this answer is
	// the ordinary one after the first call.
	if _, err := r.doc(uid).Create(ctx, first); err != nil && status.Code(err) != codes.AlreadyExists {
		return err
	}
	update := map[string]any{
		"schema":       int64(schemaVersion),
		"last_seen_at": at,
		string(c):      firestore.Increment(int64(1)),
	}
	// An empty email never writes over an address the record holds.
	if email != "" {
		update["email"] = email
	}
	_, err := r.doc(uid).Set(ctx, update, firestore.MergeAll)
	return err
}

// Get reads one record. A user with no record reads the zero value and
// no error, because a user who has made nothing has no document.
func (r *Repo) Get(ctx context.Context, uid string) (Record, error) {
	snap, err := r.doc(uid).Get(ctx)
	if status.Code(err) == codes.NotFound {
		return Record{}, nil
	}
	if err != nil {
		return Record{}, err
	}
	var rec Record
	if err := snap.DataTo(&rec); err != nil {
		return Record{}, err
	}
	return rec, nil
}

// Seed writes the counts a backfill measured, and it never lowers a
// count that is already higher: a creation between the read and the
// write must survive it (D-638).
func (r *Repo) Seed(ctx context.Context, uid, email string, counts map[Counter]int64, createdAt, lastSeen time.Time) error {
	if uid == "" {
		return errors.New("users: a record needs a user")
	}
	for c := range counts {
		if !counters[c] {
			return ErrBadCounter
		}
	}
	return r.client.RunTransaction(ctx, func(_ context.Context, tx *firestore.Transaction) error {
		ref := r.doc(uid)
		snap, err := tx.Get(ref)
		if err != nil && status.Code(err) != codes.NotFound {
			return err
		}
		var cur Record
		if snap != nil && snap.Exists() {
			if err := snap.DataTo(&cur); err != nil {
				return err
			}
		}
		data := map[string]any{"schema": int64(schemaVersion)}
		if email != "" {
			data["email"] = email
		}
		if cur.CreatedAt.IsZero() && !createdAt.IsZero() {
			data["created_at"] = createdAt.UTC()
		}
		if !lastSeen.IsZero() && lastSeen.After(cur.LastSeenAt) {
			data["last_seen_at"] = lastSeen.UTC()
		}
		held := map[Counter]int64{
			DecksCreated: cur.DecksCreated, DeckRevisions: cur.DeckRevisions,
			CollectionsUpload: cur.CollectionsUpload,
			SessionsStarted:   cur.SessionsStarted, FeedbackUp: cur.FeedbackUp,
			FeedbackDown: cur.FeedbackDown,
		}
		for c, n := range counts {
			if n > held[c] {
				data[string(c)] = n
			}
		}
		return tx.Set(ref, data, firestore.MergeAll)
	})
}

// Noter is what a service needs of this package. A nil Noter records
// nothing, so a service runs without a record store and every test that
// wires none keeps working.
type Noter interface {
	Note(ctx context.Context, uid, email string, c Counter, at time.Time) error
}

// NoteQuietly records one creation through a Noter that may be nil. It
// answers no error to the caller: a record that does not write must
// never fail the deck, the upload, or the verdict it counts.
func NoteQuietly(ctx context.Context, n Noter, uid, email string, c Counter, at time.Time) {
	if n == nil || uid == "" {
		return
	}
	_ = n.Note(ctx, uid, email, c, at)
}
