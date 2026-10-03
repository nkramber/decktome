// Package prooflink stores the links that prove an email (D-1081). A
// link is a random code of 10 characters. The store keeps a hash of the
// code alone, so a read of the store gives no working link. A link works
// one time, for 3 days (D-1082).
package prooflink

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"
	"strings"
	"time"

	"cloud.google.com/go/firestore"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Collection is the top-level collection of the links. No harvest reads
// it.
const Collection = "proof_links"

const (
	// CodeLen is the length of a code. 62 symbols in 10 places give
	// about 59 bits, and a limit per client address bounds a guess.
	CodeLen = 10
	// Life is how long a link works (D-1082).
	Life = 72 * time.Hour
	// Cooldown is the shortest time between two sends for one account.
	Cooldown = time.Minute
)

// alphabet holds the symbols of a code. Each one is safe in a path.
const alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"

var (
	// ErrUnknown is a code that the store does not hold: a code that
	// never existed, or a link that already worked.
	ErrUnknown = errors.New("prooflink: unknown code")
	// ErrExpired is a link older than Life. The open removes it.
	ErrExpired = errors.New("prooflink: expired")
	// ErrTooSoon is a send within Cooldown of the last send.
	ErrTooSoon = errors.New("prooflink: too soon")
)

// Link is one stored link.
type Link struct {
	UID       string    `firestore:"uid"`
	Email     string    `firestore:"email"`
	CreatedAt time.Time `firestore:"created_at"`
	ExpiresAt time.Time `firestore:"expires_at"`
}

// NewCode answers a random code of CodeLen symbols.
func NewCode() (string, error) {
	out := make([]byte, CodeLen)
	limit := big.NewInt(int64(len(alphabet)))
	for i := range out {
		n, err := rand.Int(rand.Reader, limit)
		if err != nil {
			return "", fmt.Errorf("prooflink: %w", err)
		}
		out[i] = alphabet[n.Int64()]
	}
	return string(out), nil
}

// Valid answers whether a text has the form of a code. The open checks
// the form before it reads the store.
func Valid(code string) bool {
	if len(code) != CodeLen {
		return false
	}
	for i := range len(code) {
		if !strings.ContainsRune(alphabet, rune(code[i])) {
			return false
		}
	}
	return true
}

// Key is the document id of a code: a hash, so the store holds no code.
func Key(code string) string {
	sum := sha256.Sum256([]byte(code))
	return hex.EncodeToString(sum[:16])
}

// Repo is the Firestore store of the links.
type Repo struct {
	client *firestore.Client
}

// NewRepo builds the store.
func NewRepo(client *firestore.Client) *Repo { return &Repo{client: client} }

// Put stores the link of one code. It removes each older link of the
// same account, so one account holds one working link. A send within
// Cooldown of the last one answers ErrTooSoon and stores nothing.
func (r *Repo) Put(ctx context.Context, code string, link Link) error {
	ref := r.client.Collection(Collection).Doc(Key(code))
	q := r.client.Collection(Collection).Where("uid", "==", link.UID)
	return r.client.RunTransaction(ctx, func(_ context.Context, tx *firestore.Transaction) error {
		old, err := tx.Documents(q).GetAll()
		if err != nil {
			return err
		}
		for _, d := range old {
			var prev Link
			if err := d.DataTo(&prev); err != nil {
				return fmt.Errorf("prooflink: %s: %w", d.Ref.ID, err)
			}
			if link.CreatedAt.Sub(prev.CreatedAt) < Cooldown {
				return ErrTooSoon
			}
		}
		for _, d := range old {
			if err := tx.Delete(d.Ref); err != nil {
				return err
			}
		}
		return tx.Create(ref, link)
	})
}

// Take reads the link of one code and removes it, so a link works one
// time. An expired link answers ErrExpired, and the take removes it too.
func (r *Repo) Take(ctx context.Context, code string, now time.Time) (Link, error) {
	ref := r.client.Collection(Collection).Doc(Key(code))
	var link Link
	err := r.client.RunTransaction(ctx, func(_ context.Context, tx *firestore.Transaction) error {
		snap, err := tx.Get(ref)
		if status.Code(err) == codes.NotFound {
			return ErrUnknown
		}
		if err != nil {
			return err
		}
		if err := snap.DataTo(&link); err != nil {
			return err
		}
		return tx.Delete(ref)
	})
	if err != nil {
		return Link{}, err
	}
	if !now.Before(link.ExpiresAt) {
		return Link{}, ErrExpired
	}
	return link, nil
}
