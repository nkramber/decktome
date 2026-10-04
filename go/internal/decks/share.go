package decks

import (
	"context"
	"errors"
	"time"

	"cloud.google.com/go/firestore"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
	"github.com/nkramber/decktome/go/internal/gzstore"
)

// storedShare is the document at shares/{token hash}: the deck a share
// link opens (D-315). The hash is the document id, so a lookup is one
// read and needs no index. The token sits in the deck alone (D-1061).
type storedShare struct {
	UID       string    `firestore:"uid"`
	DeckID    string    `firestore:"deck_id"`
	CreatedAt time.Time `firestore:"created_at"`
}

func (r *Repo) shareDoc(hash string) *firestore.DocumentRef {
	return r.client.Collection("shares").Doc(hash)
}

// Share records a new token and its hash for one deck. A deck that had a
// link loses the old one: its share document goes, so the old link
// answers NotFound (D-315). The deck's own document keeps the hash, and
// the packed proto marks the deck as shared and holds the token, so its
// owner reads the link again (D-1061).
func (r *Repo) Share(ctx context.Context, uid, id, token, hash string) error {
	if token == "" || hash == "" {
		return errors.New("decks: a share needs a token and its hash")
	}
	doc := r.doc(uid, id)
	return r.client.RunTransaction(ctx, func(_ context.Context, tx *firestore.Transaction) error {
		sd, d, err := readStored(tx, doc)
		if err != nil {
			return err
		}
		old := sd.ShareTokenHash
		d.Shared = true
		d.ShareToken = token
		updated, err := restore(d, sd)
		if err != nil {
			return err
		}
		updated.ShareTokenHash = hash
		if err := tx.Set(doc, updated); err != nil {
			return err
		}
		if old != "" && old != hash {
			if err := tx.Delete(r.shareDoc(old)); err != nil {
				return err
			}
		}
		return tx.Set(r.shareDoc(hash), storedShare{UID: uid, DeckID: id, CreatedAt: time.Now().UTC()})
	})
}

// Revoke ends the link of one deck. A deck with no link changes nothing.
func (r *Repo) Revoke(ctx context.Context, uid, id string) error {
	doc := r.doc(uid, id)
	return r.client.RunTransaction(ctx, func(_ context.Context, tx *firestore.Transaction) error {
		sd, d, err := readStored(tx, doc)
		if err != nil {
			return err
		}
		if sd.ShareTokenHash == "" {
			return nil
		}
		old := sd.ShareTokenHash
		d.Shared = false
		d.ShareToken = ""
		updated, err := restore(d, sd)
		if err != nil {
			return err
		}
		updated.ShareTokenHash = ""
		if err := tx.Set(doc, updated); err != nil {
			return err
		}
		return tx.Delete(r.shareDoc(old))
	})
}

// LookupShare answers the deck a token hash opens, or ErrNotFound. The
// deck must hold the same hash. A share document that its deck no longer
// names opens nothing, so a revoke or a delete always ends the link
// (D-315, REV-007).
func (r *Repo) LookupShare(ctx context.Context, hash string) (uid, id string, err error) {
	if hash == "" {
		return "", "", ErrNotFound
	}
	snap, err := r.shareDoc(hash).Get(ctx)
	if status.Code(err) == codes.NotFound {
		return "", "", ErrNotFound
	}
	if err != nil {
		return "", "", err
	}
	var s storedShare
	if err := snap.DataTo(&s); err != nil {
		return "", "", err
	}
	if s.UID == "" || s.DeckID == "" {
		return "", "", ErrNotFound
	}
	deck, err := r.doc(s.UID, s.DeckID).Get(ctx)
	if status.Code(err) == codes.NotFound {
		return "", "", ErrNotFound
	}
	if err != nil {
		return "", "", err
	}
	if h, err := deck.DataAt("share_token_hash"); err != nil || h != hash {
		return "", "", ErrNotFound
	}
	return s.UID, s.DeckID, nil
}

// readStored reads one deck document inside a transaction, as the stored
// row and the packed proto.
func readStored(tx *firestore.Transaction, doc *firestore.DocumentRef) (storedDeck, *mtgv1.Deck, error) {
	var sd storedDeck
	snap, err := tx.Get(doc)
	if status.Code(err) == codes.NotFound {
		return sd, nil, ErrNotFound
	}
	if err != nil {
		return sd, nil, err
	}
	if err := snap.DataTo(&sd); err != nil {
		return sd, nil, err
	}
	d := &mtgv1.Deck{}
	if err := gzstore.UnmarshalProto(sd.DeckGz, d); err != nil {
		return sd, nil, err
	}
	return sd, d, nil
}

// restore packs a deck back into its stored row, with the fields that
// the store alone holds, so a write moves nothing in the list and drops
// no link and no eval mark.
func restore(d *mtgv1.Deck, sd storedDeck) (storedDeck, error) {
	payload, err := gzstore.MarshalProto(d)
	if err != nil {
		return storedDeck{}, err
	}
	if len(payload) > gzstore.MaxStoredBytes {
		return storedDeck{}, ErrTooLarge
	}
	updated := toStored(d, payload)
	keepStored(&updated, sd)
	return updated, nil
}

// keepStored copies the fields that the deck message does not carry from
// the row that a write replaces: the create time, the share hash
// (D-315), and the eval mark (D-1132).
func keepStored(updated *storedDeck, sd storedDeck) {
	updated.CreatedAt = sd.CreatedAt
	updated.ShareTokenHash = sd.ShareTokenHash
	updated.HasBeenEvaluated = sd.HasBeenEvaluated
	updated.EvaluatedAt = sd.EvaluatedAt
}
