// Package allowlist holds the invited emails of the deployed app (D-314,
// D-420). The list is one Firestore document, config/allowlist, with one
// field: emails, a list of lower-case addresses. The interceptor reads
// the list through Allowed, and a cache of one minute keeps a turn from
// a Firestore read. `make allow EMAIL=...` writes the document through
// Add, so an invite needs no deploy.
package allowlist

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"sync"
	"time"

	"cloud.google.com/go/firestore"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Collection and Doc name the one document.
const (
	Collection = "config"
	Doc        = "allowlist"
	// Field is the list field of the document.
	Field = "emails"
	// TTL is how long a read list serves before the next read (D-420).
	TTL = time.Minute
)

// Fetch reads the emails of the list. Firestore is the one production
// source, and tests give a function.
type Fetch func(ctx context.Context) ([]string, error)

// List answers Allowed from a cached read.
type List struct {
	fetch Fetch
	now   func() time.Time
	ttl   time.Duration

	mu      sync.Mutex
	emails  map[string]bool
	readAt  time.Time
	haveOne bool
}

// New builds a list over a fetch, with the TTL of D-420.
func New(fetch Fetch) *List {
	return &List{fetch: fetch, now: time.Now, ttl: TTL}
}

// WithClock replaces time.Now (tests).
func (l *List) WithClock(now func() time.Time) *List {
	l.now = now
	return l
}

// Allowed reports whether the email is on the list. An empty email is
// never on it. A read error answers the error, and a list read inside
// the last minute serves without a read.
func (l *List) Allowed(ctx context.Context, email string) (bool, error) {
	key := Normalize(email)
	if key == "" {
		return false, nil
	}
	l.mu.Lock()
	defer l.mu.Unlock()
	if !l.haveOne || l.now().Sub(l.readAt) >= l.ttl {
		emails, err := l.fetch(ctx)
		if err != nil {
			return false, err
		}
		l.emails = make(map[string]bool, len(emails))
		for _, e := range emails {
			if k := Normalize(e); k != "" {
				l.emails[k] = true
			}
		}
		l.readAt, l.haveOne = l.now(), true
	}
	return l.emails[key], nil
}

// Normalize is the key of an email: lower case, trimmed. Two spellings of
// one address are one invite.
func Normalize(email string) string { return strings.ToLower(strings.TrimSpace(email)) }

// FromFirestore builds the production list over the one document.
func FromFirestore(client *firestore.Client) *List {
	return New(func(ctx context.Context) ([]string, error) {
		snap, err := client.Collection(Collection).Doc(Doc).Get(ctx)
		if err != nil {
			if isNotFound(err) {
				// No document is an empty list: nobody is invited yet.
				return nil, nil
			}
			return nil, err
		}
		var doc struct {
			Emails []string `firestore:"emails"`
		}
		if err := snap.DataTo(&doc); err != nil {
			return nil, fmt.Errorf("allowlist: %w", err)
		}
		return doc.Emails, nil
	})
}

// Add puts one email on the list, and it creates the document on the
// first invite. The write is a set union, so a repeated invite changes
// nothing.
func Add(ctx context.Context, client *firestore.Client, email string) error {
	key := Normalize(email)
	if key == "" || !strings.Contains(key, "@") {
		return errors.New("allowlist: an invite needs an email address")
	}
	_, err := client.Collection(Collection).Doc(Doc).Set(ctx, map[string]any{
		Field: firestore.ArrayUnion(key),
	}, firestore.MergeAll)
	return err
}

// Remove takes one email off the list.
func Remove(ctx context.Context, client *firestore.Client, email string) error {
	key := Normalize(email)
	if key == "" {
		return errors.New("allowlist: an email is needed")
	}
	_, err := client.Collection(Collection).Doc(Doc).Set(ctx, map[string]any{
		Field: firestore.ArrayRemove(key),
	}, firestore.MergeAll)
	return err
}

func isNotFound(err error) bool { return status.Code(err) == codes.NotFound }
