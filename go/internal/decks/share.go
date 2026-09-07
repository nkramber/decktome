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
// read and needs no index. The token itself is never stored.
type storedShare struct {
	UID       string    `firestore:"uid"`
	DeckID    string    `firestore:"deck_id"`
	CreatedAt time.Time `firestore:"created_at"`
}

func (r *Repo) shareDoc(hash string) *firestore.DocumentRef {
	return r.client.Collection("shares").Doc(hash)
}

// Share records the hash of a new token for one deck. A deck that had a
// link loses the old one: its share document goes, so the old link
// answers NotFound (D-315). The deck's own document keeps the hash, and
// the packed proto marks the deck as shared.
func (r *Repo) Share(ctx context.Context, uid, id, hash string) error {
	if hash == "" {
		return errors.New("decks: a share needs a token hash")
	}
	doc := r.doc(uid, id)
	return r.client.RunTransaction(ctx, func(_ context.Context, tx *firestore.Transaction) error {
		sd, d, err := readStored(tx, doc)
		if err != nil {
			return err
		}
		old := sd.ShareTokenHash
		d.Shared = true
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

// LookupShare answers the deck a token hash opens, or ErrNotFound.
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

// restore packs a deck back into its stored row, with the create time
// and the share hash the row already had, so a write moves nothing in
// the list and drops no link.
func restore(d *mtgv1.Deck, sd storedDeck) (storedDeck, error) {
	payload, err := gzstore.MarshalProto(d)
	if err != nil {
		return storedDeck{}, err
	}
	if len(payload) > gzstore.MaxStoredBytes {
		return storedDeck{}, ErrTooLarge
	}
	updated := toStored(d, payload)
	updated.CreatedAt = sd.CreatedAt
	updated.ShareTokenHash = sd.ShareTokenHash
	return updated, nil
}
