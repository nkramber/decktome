// Package deckreads records the deck reads of each user for one hour
// (D-1107). An import keeps the link of an Archidekt deck only after a
// read of the same deck and the same text, so a stored link always names
// the list that the server read. The store keeps the deck id and a hash
// of the text, and no list.
//
// Each read is one document, so no count of reads drops a read before
// its hour ends. A Firestore TTL policy on expire_at deletes a document
// after its hour, typically within 24 hours (D-1113).
package deckreads

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"time"

	"cloud.google.com/go/firestore"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Collection is the top-level collection of the reads. No harvest reads
// it.
const Collection = "deck_reads"

// TTLField is the field of the TTL policy of Collection (D-1113).
const TTLField = "expire_at"

// Life is how long a read lets an import keep its link.
const Life = time.Hour

// Read is one deck read of a user.
type Read struct {
	UID      string    `firestore:"uid"`
	DeckID   int64     `firestore:"deck_id"`
	TextHash string    `firestore:"text_hash"`
	ReadAt   time.Time `firestore:"read_at"`
	ExpireAt time.Time `firestore:"expire_at"`
}

// Hash is the hash of a list text.
func Hash(text string) string {
	sum := sha256.Sum256([]byte(text))
	return hex.EncodeToString(sum[:])
}

// Key is the document id of one read: a hash of the user, the deck, and
// the text. A second read of the same list replaces the first one.
func Key(uid string, deckID int64, text string) string {
	sum := sha256.Sum256([]byte(uid + "\x00" + strconv.FormatInt(deckID, 10) + "\x00" + Hash(text)))
	return hex.EncodeToString(sum[:])
}

// New is the read of a deck and a text by a user at a time.
func New(uid string, deckID int64, text string, now time.Time) Read {
	return Read{UID: uid, DeckID: deckID, TextHash: Hash(text), ReadAt: now, ExpireAt: now.Add(Life)}
}

// Matches answers whether a read has the user, the deck, and the text,
// and its hour has not ended.
func (r Read) Matches(uid string, deckID int64, text string, now time.Time) bool {
	return r.UID == uid && r.DeckID == deckID && r.TextHash == Hash(text) && now.Before(r.ExpireAt)
}

// Repo is the Firestore store of the reads.
type Repo struct {
	client *firestore.Client
}

// NewRepo builds the store.
func NewRepo(client *firestore.Client) *Repo { return &Repo{client: client} }

// Add records one read of a user.
func (r *Repo) Add(ctx context.Context, uid string, deckID int64, text string, now time.Time) error {
	_, err := r.client.Collection(Collection).Doc(Key(uid, deckID, text)).Set(ctx, New(uid, deckID, text, now))
	return err
}

// Has answers whether a user read the deck with the text less than Life
// ago. A read whose hour ended answers false before the TTL policy
// deletes it.
func (r *Repo) Has(ctx context.Context, uid string, deckID int64, text string, now time.Time) (bool, error) {
	snap, err := r.client.Collection(Collection).Doc(Key(uid, deckID, text)).Get(ctx)
	if status.Code(err) == codes.NotFound {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	var read Read
	if err := snap.DataTo(&read); err != nil {
		return false, err
	}
	return read.Matches(uid, deckID, text, now), nil
}
