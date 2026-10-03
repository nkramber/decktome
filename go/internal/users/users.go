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
// The first verified call of a user writes the record, so a user who
// signs in and makes nothing has one too (D-1092).
//
// The server writes this record and the client never does. The email
// comes from the verified token (D-596), and the Firestore rules deny
// every client read.
package users

import (
	"context"
	"errors"
	"sync"
	"time"

	"cloud.google.com/go/firestore"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// schemaVersion names the shape of the stored document. Version 2 moves
// last_seen_at on any activity and keeps the newest creation in
// last_creation_at (D-1093). In version 1, last_seen_at was the newest
// creation.
const schemaVersion = 2

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
	// DecksImported counts the lists a user brought, apart from the decks
	// the app built (PR-70, D-852).
	DecksImported Counter = "total_decks_imported"
)

// counters lists every field the backfill and the reader know. A field
// off this list is not a counter, and Note refuses it.
var counters = map[Counter]bool{
	DecksCreated: true, DeckRevisions: true, CollectionsUpload: true,
	SessionsStarted: true, FeedbackUp: true, FeedbackDown: true,
	DecksImported: true,
}

// creations lists the counters that move last_creation_at: the decks and
// the chats (D-1093). An upload and a verdict move last_seen_at alone.
var creations = map[Counter]bool{
	DecksCreated: true, DeckRevisions: true, DecksImported: true,
	SessionsStarted: true,
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
	// LastSeenAt is the newest activity of the user: a verified call or a
	// creation (D-1093). A call moves it at most once in VisitTTL.
	LastSeenAt time.Time `firestore:"last_seen_at"`
	// LastCreationAt is the newest deck or chat the user made (D-1093). A
	// user who made none has the zero time.
	LastCreationAt time.Time `firestore:"last_creation_at"`

	DecksCreated      int64 `firestore:"total_decks_created"`
	DeckRevisions     int64 `firestore:"total_deck_revisions"`
	CollectionsUpload int64 `firestore:"total_collections_uploaded"`
	SessionsStarted   int64 `firestore:"total_sessions_started"`
	FeedbackUp        int64 `firestore:"total_feedback_up"`
	FeedbackDown      int64 `firestore:"total_feedback_down"`
	DecksImported     int64 `firestore:"total_decks_imported"`

	// DeactivatedAt closes the account (D-941). The API refuses the user,
	// and the share links of the user answer NotFound. Every record of
	// the user stays. Nil is an open account.
	DeactivatedAt *time.Time `firestore:"deactivated_at,omitempty"`
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
// the time, and sets the creation date on the first write alone. A deck
// or a chat also moves last_creation_at (D-1093).
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
	if creations[c] {
		first["last_creation_at"] = at
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
	if creations[c] {
		update["last_creation_at"] = at
	}
	// An empty email never writes over an address the record holds.
	if email != "" {
		update["email"] = email
	}
	_, err := r.doc(uid).Set(ctx, update, firestore.MergeAll)
	return err
}

// Touch records one verified call of uid at the given time (D-1092). It
// makes the record with zero counts when none exists, and it never moves
// last_seen_at back. A record of version 1 first gets its last_seen_at
// as last_creation_at, because that field held the newest creation then
// (D-1094).
func (r *Repo) Touch(ctx context.Context, uid, email string, at time.Time) error {
	if uid == "" {
		return errors.New("users: a record needs a user")
	}
	at = at.UTC()
	return r.client.RunTransaction(ctx, func(_ context.Context, tx *firestore.Transaction) error {
		ref := r.doc(uid)
		snap, err := tx.Get(ref)
		if status.Code(err) == codes.NotFound {
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
			return tx.Create(ref, first)
		}
		if err != nil {
			return err
		}
		var cur Record
		if err := snap.DataTo(&cur); err != nil {
			return err
		}
		data := map[string]any{"schema": int64(schemaVersion)}
		upgrade(cur, data)
		if at.After(cur.LastSeenAt) {
			data["last_seen_at"] = at
		}
		if cur.CreatedAt.IsZero() {
			data["created_at"] = at
		}
		if email != "" {
			data["email"] = email
		}
		return tx.Set(ref, data, firestore.MergeAll)
	})
}

// upgrade adds to data the move of a version 1 record to version 2: the
// last_seen_at of version 1 is the newest creation (D-1094). A record of
// version 2, or one with no last_seen_at, needs nothing.
func upgrade(cur Record, data map[string]any) {
	if cur.Schema < 2 && cur.LastCreationAt.IsZero() && !cur.LastSeenAt.IsZero() {
		data["last_creation_at"] = cur.LastSeenAt
	}
}

// Get reads one record. A user with no record reads the zero value and
// no error. A user who has made no verified call since D-1092 has none.
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
		upgrade(cur, data)
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
			FeedbackDown: cur.FeedbackDown, DecksImported: cur.DecksImported,
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

// Deactivate closes the account of uid at the given time, and keeps
// every record (D-941). It writes the one field, and a second call keeps
// the first time. A uid with no record gets one that holds the mark.
func (r *Repo) Deactivate(ctx context.Context, uid string, at time.Time) (time.Time, error) {
	ref := r.doc(uid)
	var when time.Time
	err := r.client.RunTransaction(ctx, func(_ context.Context, tx *firestore.Transaction) error {
		snap, err := tx.Get(ref)
		if err != nil && status.Code(err) != codes.NotFound {
			return err
		}
		if err == nil {
			var rec Record
			if err := snap.DataTo(&rec); err != nil {
				return err
			}
			if rec.DeactivatedAt != nil {
				when = *rec.DeactivatedAt
				return nil
			}
		}
		when = at.UTC()
		return tx.Set(ref, map[string]any{"deactivated_at": when}, firestore.MergeAll)
	})
	return when, err
}

// Deactivated says whether the account of uid is closed. A uid with no
// record is open.
func (r *Repo) Deactivated(ctx context.Context, uid string) (bool, error) {
	rec, err := r.Get(ctx, uid)
	if err != nil {
		return false, err
	}
	return rec.DeactivatedAt != nil, nil
}

// ClosedTTL is how long one read of a closed mark serves, the TTL of the
// invite list (D-420, D-941).
const ClosedTTL = time.Minute

// ClosedCache answers Closed for each uid from a read at most ClosedTTL
// old, so a request reads the store once a minute for each user.
type ClosedCache struct {
	read func(ctx context.Context, uid string) (bool, error)
	now  func() time.Time
	mu   sync.Mutex
	seen map[string]closedRead
}

type closedRead struct {
	closed bool
	at     time.Time
}

// NewClosedCache wraps a read of the closed mark, such as
// Repo.Deactivated.
func NewClosedCache(read func(ctx context.Context, uid string) (bool, error), now func() time.Time) *ClosedCache {
	if now == nil {
		now = time.Now
	}
	return &ClosedCache{read: read, now: now, seen: map[string]closedRead{}}
}

// Closed says whether the account of uid is closed. A failed read is an
// error, and the cache keeps nothing from it.
func (c *ClosedCache) Closed(ctx context.Context, uid string) (bool, error) {
	c.mu.Lock()
	got, ok := c.seen[uid]
	c.mu.Unlock()
	if ok && c.now().Sub(got.at) < ClosedTTL {
		return got.closed, nil
	}
	closed, err := c.read(ctx, uid)
	if err != nil {
		return false, err
	}
	c.mu.Lock()
	c.seen[uid] = closedRead{closed: closed, at: c.now()}
	c.mu.Unlock()
	return closed, nil
}

// VisitTTL is the shortest time between two writes of one user's
// last_seen_at by one server (D-1093).
const VisitTTL = 5 * time.Minute

// visitTimeout bounds one write, because the write runs inside the
// request it records.
const visitTimeout = 2 * time.Second

// Visits records the verified calls of each user through Repo.Touch, at
// most once in VisitTTL for each user (D-1092, D-1093).
type Visits struct {
	touch func(ctx context.Context, uid, email string, at time.Time) error
	now   func() time.Time
	mu    sync.Mutex
	seen  map[string]time.Time
}

// NewVisits wraps a write of one visit, such as Repo.Touch.
func NewVisits(touch func(ctx context.Context, uid, email string, at time.Time) error, now func() time.Time) *Visits {
	if now == nil {
		now = time.Now
	}
	return &Visits{touch: touch, now: now, seen: map[string]time.Time{}}
}

// Seen records one verified call of uid. It answers no error: a record
// that does not write must never fail the call it records. A failed
// write keeps nothing, so the next call tries again.
func (v *Visits) Seen(ctx context.Context, uid, email string) {
	if uid == "" {
		return
	}
	at := v.now()
	v.mu.Lock()
	last, ok := v.seen[uid]
	v.mu.Unlock()
	if ok && at.Sub(last) < VisitTTL {
		return
	}
	ctx, cancel := context.WithTimeout(ctx, visitTimeout)
	defer cancel()
	if err := v.touch(ctx, uid, email, at); err != nil {
		return
	}
	v.mu.Lock()
	v.seen[uid] = at
	v.mu.Unlock()
}
