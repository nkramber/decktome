// Package access stores the requests for beta access (D-1075). One
// record holds one email, so a repeat request updates its record and
// makes no second one. The admin screen of the owner reads the records
// and decides each one (D-1076).
package access

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"slices"
	"strings"
	"time"
	"unicode/utf8"

	"cloud.google.com/go/firestore"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/nkramber/decktome/go/internal/allowlist"
)

// Collection is the top-level collection of the requests. No harvest
// reads it.
const Collection = "access_requests"

const (
	// MaxNoteRunes caps the optional note of a request (D-1074).
	MaxNoteRunes = 500
	// MaxEmailBytes is the longest address that SMTP carries.
	MaxEmailBytes = 254
	// MaxList caps one read of the admin screen.
	MaxList = 200
)

// The states of a request.
const (
	Pending   = "pending"
	Approved  = "approved"
	Dismissed = "dismissed"
)

// ErrNotFound is a decision on an email with no request.
var ErrNotFound = errors.New("access: no request holds this email")

// Request is one stored request.
type Request struct {
	Email     string    `firestore:"email"`
	Note      string    `firestore:"note"`
	Status    string    `firestore:"status"`
	Count     int64     `firestore:"count"`
	CreatedAt time.Time `firestore:"created_at"`
	UpdatedAt time.Time `firestore:"updated_at"`
}

// Key is the document id of an email. The id is a hash, because an
// address can hold a slash, and a slash breaks a document path.
func Key(email string) string {
	sum := sha256.Sum256([]byte(allowlist.Normalize(email)))
	return hex.EncodeToString(sum[:16])
}

// Clean checks one request and answers the stored form: the email in
// its normal form, and the note trimmed.
func Clean(email, note string) (string, string, error) {
	email = allowlist.Normalize(email)
	at := strings.LastIndex(email, "@")
	if at < 1 || at == len(email)-1 || len(email) > MaxEmailBytes || strings.ContainsAny(email, " \t\r\n") {
		return "", "", errors.New("enter one email address")
	}
	note = strings.TrimSpace(note)
	if utf8.RuneCountInString(note) > MaxNoteRunes {
		return "", "", fmt.Errorf("the note can hold %d characters at most", MaxNoteRunes)
	}
	return email, note, nil
}

// Repo is the Firestore store of the requests.
type Repo struct {
	client *firestore.Client
}

// NewRepo builds the store.
func NewRepo(client *firestore.Client) *Repo { return &Repo{client: client} }

// Put stores one request in its clean form. A new email makes a pending
// record, and created is true. A repeat request counts one more, and a
// note that is not empty replaces the note. A repeat never changes the
// state, so a dismissed request stays dismissed.
func (r *Repo) Put(ctx context.Context, email, note string, now time.Time) (created bool, err error) {
	ref := r.client.Collection(Collection).Doc(Key(email))
	err = r.client.RunTransaction(ctx, func(ctx context.Context, tx *firestore.Transaction) error {
		created = false
		snap, err := tx.Get(ref)
		if status.Code(err) == codes.NotFound {
			created = true
			return tx.Create(ref, Request{Email: email, Note: note, Status: Pending, Count: 1, CreatedAt: now, UpdatedAt: now})
		}
		if err != nil {
			return err
		}
		updates := []firestore.Update{
			{Path: "count", Value: firestore.Increment(1)},
			{Path: "updated_at", Value: now},
		}
		if note != "" {
			updates = append(updates, firestore.Update{Path: "note", Value: note})
		}
		return tx.Update(snap.Ref, updates)
	})
	return created, err
}

// List reads the requests in one state, the newest first. The sort runs
// here, so the query needs no composite index.
func (r *Repo) List(ctx context.Context, state string) ([]Request, error) {
	docs, err := r.client.Collection(Collection).Where("status", "==", state).Documents(ctx).GetAll()
	if err != nil {
		return nil, err
	}
	out := make([]Request, 0, len(docs))
	for _, d := range docs {
		var req Request
		if err := d.DataTo(&req); err != nil {
			return nil, fmt.Errorf("access: %s: %w", d.Ref.ID, err)
		}
		out = append(out, req)
	}
	slices.SortFunc(out, func(a, b Request) int { return b.CreatedAt.Compare(a.CreatedAt) })
	if len(out) > MaxList {
		out = out[:MaxList]
	}
	return out, nil
}

// Get reads the request of one email, or ErrNotFound.
func (r *Repo) Get(ctx context.Context, email string) (Request, error) {
	snap, err := r.client.Collection(Collection).Doc(Key(email)).Get(ctx)
	if status.Code(err) == codes.NotFound {
		return Request{}, ErrNotFound
	}
	if err != nil {
		return Request{}, err
	}
	var req Request
	if err := snap.DataTo(&req); err != nil {
		return Request{}, err
	}
	return req, nil
}

// Decide sets the state of the request of one email.
func (r *Repo) Decide(ctx context.Context, email, state string, now time.Time) error {
	_, err := r.client.Collection(Collection).Doc(Key(email)).Update(ctx, []firestore.Update{
		{Path: "status", Value: state},
		{Path: "updated_at", Value: now},
	})
	if status.Code(err) == codes.NotFound {
		return ErrNotFound
	}
	return err
}
