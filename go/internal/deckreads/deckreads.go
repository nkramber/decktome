// Package deckreads records the deck reads of each user for one hour
// (D-1107). An import keeps the link of an Archidekt deck only after a
// read of the same deck and the same text, so a stored link always names
// the list that the server read. The store keeps the deck id and a hash
// of the text, and no list.
package deckreads

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"time"

	"cloud.google.com/go/firestore"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Collection is the top-level collection of the reads, one document for
// each user. No harvest reads it.
const Collection = "deck_reads"

const (
	// Life is how long a read lets an import keep its link.
	Life = time.Hour
	// MaxReads bounds the reads of one document. The oldest read goes
	// first.
	MaxReads = 20
)

// Read is one deck read of a user.
type Read struct {
	DeckID   int64     `firestore:"deck_id"`
	TextHash string    `firestore:"text_hash"`
	ReadAt   time.Time `firestore:"read_at"`
}

// record is the document of one user.
type record struct {
	Reads []Read `firestore:"reads"`
}

// Hash is the hash of a list text.
func Hash(text string) string {
	sum := sha256.Sum256([]byte(text))
	return hex.EncodeToString(sum[:])
}

// Keep answers the reads younger than Life, and a new read when add is
// not nil. It keeps the newest MaxReads.
func Keep(reads []Read, add *Read, now time.Time) []Read {
	out := make([]Read, 0, len(reads)+1)
	for _, r := range reads {
		if now.Sub(r.ReadAt) < Life {
			out = append(out, r)
		}
	}
	if add != nil {
		out = append(out, *add)
	}
	if len(out) > MaxReads {
		out = out[len(out)-MaxReads:]
	}
	return out
}

// Matches answers whether a read younger than Life has the deck and the
// text.
func Matches(reads []Read, deckID int64, text string, now time.Time) bool {
	hash := Hash(text)
	for _, r := range Keep(reads, nil, now) {
		if r.DeckID == deckID && r.TextHash == hash {
			return true
		}
	}
	return false
}

// Repo is the Firestore store of the reads.
type Repo struct {
	client *firestore.Client
}

// NewRepo builds the store.
func NewRepo(client *firestore.Client) *Repo { return &Repo{client: client} }

// Add records one read of a user. It removes each read older than Life.
func (r *Repo) Add(ctx context.Context, uid string, deckID int64, text string, now time.Time) error {
	ref := r.client.Collection(Collection).Doc(uid)
	read := Read{DeckID: deckID, TextHash: Hash(text), ReadAt: now}
	return r.client.RunTransaction(ctx, func(_ context.Context, tx *firestore.Transaction) error {
		var rec record
		snap, err := tx.Get(ref)
		switch {
		case status.Code(err) == codes.NotFound:
		case err != nil:
			return err
		default:
			if err := snap.DataTo(&rec); err != nil {
				return err
			}
		}
		return tx.Set(ref, record{Reads: Keep(rec.Reads, &read, now)})
	})
}

// Has answers whether a user read the deck with the text less than Life
// ago.
func (r *Repo) Has(ctx context.Context, uid string, deckID int64, text string, now time.Time) (bool, error) {
	snap, err := r.client.Collection(Collection).Doc(uid).Get(ctx)
	if status.Code(err) == codes.NotFound {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	var rec record
	if err := snap.DataTo(&rec); err != nil {
		return false, err
	}
	return Matches(rec.Reads, deckID, text, now), nil
}
