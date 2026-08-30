// Package decks stores built decks in Firestore (D-245). A kept deck is
// what `GetDeck`, `ListDecks`, `Export`, and the staleness job read, and
// its id on the session is what makes `Context.AfterBuild` true.
//
// The shape follows internal/sessions: one document per deck under the
// user, the proto stored as gzip protojson so a proto change reads back
// with no migration, and flat fields for the list query.
package decks

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"

	mtgv1 "github.com/nkramber/mtg-deck-builder/go/gen/mtg/v1"
	"github.com/nkramber/mtg-deck-builder/go/internal/gzstore"
)

// schemaVersion counts the stored shape, not the proto.
const schemaVersion = 1

// ErrNotFound reports a deck id no document answers.
var ErrNotFound = errors.New("deck not found")

// ErrTooLarge reports a deck that does not fit one document.
var ErrTooLarge = errors.New("deck too large for one document (max 900 KiB gzip)")

// Repo stores decks in Firestore. The caller owns the client.
type Repo struct {
	client *firestore.Client
}

// NewRepo wraps a Firestore client.
func NewRepo(client *firestore.Client) *Repo { return &Repo{client: client} }

// storedDeck is the document shape. The flat fields answer the list
// query without inflating every deck.
type storedDeck struct {
	SessionID          string    `firestore:"session_id"`
	Name               string    `firestore:"name"`
	FormatID           int64     `firestore:"format_id"`
	CreatedAt          time.Time `firestore:"created_at"`
	LegalityAsOf       string    `firestore:"legality_as_of"`
	BuyCostUSD         float64   `firestore:"buy_cost_usd"`
	Stale              bool      `firestore:"stale"`
	Favorite           bool      `firestore:"favorite"`
	PowerBracket       int64     `firestore:"power_bracket"`
	PowerSixtyStep     int64     `firestore:"power_sixty_step"`
	CardCount          int64     `firestore:"card_count"`
	CommanderOracleIDs []string  `firestore:"commander_oracle_ids"`
	CommanderNames     []string  `firestore:"commander_names"`
	DeckGz             []byte    `firestore:"deck_gz"`
	SchemaVersion      int64     `firestore:"schema_version"`
}

// listFields are the flat fields List reads. The list never inflates a
// deck: it answers from these alone.
//
// A document written before PR-17 holds none of the five fields that PR-17
// added. Firestore reads an absent field as the zero value, so an older
// deck lists with no favorite mark, no power, a count of zero, and no
// commander. A rename or a favorite write fills them for that deck.
var listFields = []string{
	"session_id", "name", "format_id", "created_at", "legality_as_of", "buy_cost_usd", "stale",
	"favorite", "power_bracket", "power_sixty_step", "card_count", "commander_oracle_ids", "commander_names",
}

func (r *Repo) col(uid string) *firestore.CollectionRef {
	return r.client.Collection("users").Doc(uid).Collection("decks")
}

func (r *Repo) doc(uid, id string) *firestore.DocumentRef { return r.col(uid).Doc(id) }

// NewID reserves a deck id without a write. The build needs the id before
// it stores the deck, because the deck carries its own id.
func (r *Repo) NewID(uid string) string { return r.col(uid).NewDoc().ID }

// Put writes one deck. A build writes a deck once: a re-roll is a new
// deck with a new id (D-18). Update writes the two fields the user owns
// after that (PR-17).
func (r *Repo) Put(ctx context.Context, uid string, d *mtgv1.Deck) error {
	if d.GetId() == "" {
		return errors.New("decks: a deck needs an id")
	}
	if uid == "" {
		return errors.New("decks: a deck needs a user")
	}
	payload, err := gzstore.MarshalProto(d)
	if err != nil {
		return err
	}
	if len(payload) > gzstore.MaxStoredBytes {
		return ErrTooLarge
	}
	_, err = r.doc(uid, d.GetId()).Set(ctx, toStored(d, payload))
	return err
}

// toStored builds the document from a deck and its packed payload.
func toStored(d *mtgv1.Deck, payload []byte) storedDeck {
	created := d.GetCreatedAt().AsTime()
	if d.GetCreatedAt() == nil {
		created = time.Now().UTC()
	}
	return storedDeck{
		SessionID:          d.GetSessionId(),
		Name:               d.GetName(),
		FormatID:           int64(d.GetFormat().GetId()),
		CreatedAt:          created,
		LegalityAsOf:       d.GetLegalityAsOf(),
		BuyCostUSD:         d.GetBuyCostUsd(),
		Stale:              d.GetStale(),
		Favorite:           d.GetFavorite(),
		PowerBracket:       int64(d.GetPower().GetBracket()),
		PowerSixtyStep:     int64(d.GetPower().GetSixtyStep()),
		CardCount:          int64(CardCount(d)),
		CommanderOracleIDs: d.GetCommanderOracleIds(),
		CommanderNames:     CommanderNames(d),
		DeckGz:             payload,
		SchemaVersion:      schemaVersion,
	}
}

// CardCount sums count over the main deck. The sideboard and the upgrades
// stay out, and a commander counts when the card list holds it.
func CardCount(d *mtgv1.Deck) int32 {
	var n int32
	for _, c := range d.GetCards() {
		n += c.GetCount()
	}
	return n
}

// CommanderNames reads the names of the deck's commanders from its card
// list. A commander the card list omits contributes no name, so a search
// by commander misses that deck.
func CommanderNames(d *mtgv1.Deck) []string {
	ids := d.GetCommanderOracleIds()
	if len(ids) == 0 {
		return nil
	}
	byID := make(map[string]string, len(d.GetCards()))
	for _, c := range d.GetCards() {
		byID[c.GetOracleId()] = c.GetName()
	}
	var out []string
	for _, id := range ids {
		if name := byID[id]; name != "" {
			out = append(out, name)
		}
	}
	return out
}

// Get reads one deck.
func (r *Repo) Get(ctx context.Context, uid, id string) (*mtgv1.Deck, error) {
	snap, err := r.doc(uid, id).Get(ctx)
	if status.Code(err) == codes.NotFound {
		return nil, ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	var sd storedDeck
	if err := snap.DataTo(&sd); err != nil {
		return nil, err
	}
	var d mtgv1.Deck
	if err := gzstore.UnmarshalProto(sd.DeckGz, &d); err != nil {
		return nil, err
	}
	return &d, nil
}

// Filter narrows a listing (PR-17). The zero value keeps every deck.
type Filter struct {
	// Format keeps one format. UNSPECIFIED keeps them all.
	Format mtgv1.FormatId
	// Favorite keeps the favorites, or the rest. Nil keeps them all.
	Favorite *bool
	// Query keeps a deck whose name or commander name holds this text,
	// without regard to case. Empty keeps them all.
	Query string
	// PowerBracket keeps one Commander bracket, 1 to 5. Zero keeps them
	// all. PowerLevel is a oneof, so the power filter has two arms
	// (OQ-47), and a deck matches when either arm matches.
	PowerBracket int32
	// PowerSixtyStep keeps one step of the 60-card scale. UNSPECIFIED
	// keeps them all.
	PowerSixtyStep mtgv1.SixtyStep
}

// keep reports whether one stored row passes the filter.
func (f Filter) keep(sd storedDeck) bool {
	if f.Format != mtgv1.FormatId_FORMAT_ID_UNSPECIFIED && mtgv1.FormatId(sd.FormatID) != f.Format {
		return false
	}
	if f.Favorite != nil && sd.Favorite != *f.Favorite {
		return false
	}
	if f.PowerBracket > 0 && int32(sd.PowerBracket) != f.PowerBracket { //nolint:gosec // PowerBracket was an int32 at Put
		return false
	}
	if f.PowerSixtyStep != mtgv1.SixtyStep_SIXTY_STEP_UNSPECIFIED && mtgv1.SixtyStep(sd.PowerSixtyStep) != f.PowerSixtyStep {
		return false
	}
	if f.Query == "" {
		return true
	}
	q := strings.ToLower(f.Query)
	if strings.Contains(strings.ToLower(sd.Name), q) {
		return true
	}
	for _, name := range sd.CommanderNames {
		if strings.Contains(strings.ToLower(name), q) {
			return true
		}
	}
	return false
}

// List reads the user's decks, newest first, from the flat fields alone.
// A listed deck carries its id, name, session, format, power, count,
// commanders, cost, legality date, staleness, and favorite mark, and no
// cards: Get reads the whole deck.
//
// The filter runs in Go over the rows Firestore returns, not as a query.
// One read serves every filter, and no composite index has to exist. The
// scan bounds the rows the read takes, and 0 reads them all. A user with
// more decks than the scan needs a search index, and PR-17 has none.
func (r *Repo) List(ctx context.Context, uid string, f Filter, scan int) ([]*mtgv1.Deck, error) {
	q := r.col(uid).Select(listFields...).OrderBy("created_at", firestore.Desc)
	if scan > 0 {
		q = q.Limit(scan)
	}
	it := q.Documents(ctx)
	defer it.Stop()
	var out []*mtgv1.Deck
	for {
		snap, err := it.Next()
		if errors.Is(err, iterator.Done) {
			return out, nil
		}
		if err != nil {
			return nil, err
		}
		var sd storedDeck
		if err := snap.DataTo(&sd); err != nil {
			return nil, fmt.Errorf("deck %s: %w", snap.Ref.ID, err)
		}
		if !f.keep(sd) {
			continue
		}
		out = append(out, storedToProto(snap.Ref.ID, sd))
	}
}

// storedToProto builds the list view of a deck from its flat fields.
func storedToProto(id string, sd storedDeck) *mtgv1.Deck {
	return &mtgv1.Deck{
		Id:                 id,
		Name:               sd.Name,
		SessionId:          sd.SessionID,
		Format:             &mtgv1.Format{Id: mtgv1.FormatId(sd.FormatID)}, //nolint:gosec // FormatID was an enum at Put
		Power:              storedPower(sd),
		LegalityAsOf:       sd.LegalityAsOf,
		BuyCostUsd:         sd.BuyCostUSD,
		Stale:              sd.Stale,
		Favorite:           sd.Favorite,
		CardCount:          int32(sd.CardCount), //nolint:gosec // CardCount was an int32 at Put
		CommanderOracleIds: sd.CommanderOracleIDs,
		CreatedAt:          timestamppb.New(sd.CreatedAt),
	}
}

// storedPower rebuilds the power oneof from the two flat fields. A deck
// with neither one carries no power, the same as a deck of a build that
// set none.
func storedPower(sd storedDeck) *mtgv1.PowerLevel {
	switch {
	case sd.PowerBracket > 0:
		return &mtgv1.PowerLevel{Level: &mtgv1.PowerLevel_Bracket{Bracket: int32(sd.PowerBracket)}} //nolint:gosec // PowerBracket was an int32 at Put
	case sd.PowerSixtyStep > 0:
		return &mtgv1.PowerLevel{Level: &mtgv1.PowerLevel_SixtyStep{SixtyStep: mtgv1.SixtyStep(sd.PowerSixtyStep)}} //nolint:gosec // PowerSixtyStep was an enum at Put
	default:
		return nil
	}
}

// Delete removes one deck for good (PR-17). A deck id no document answers
// is ErrNotFound, so a second delete does not read as a success.
func (r *Repo) Delete(ctx context.Context, uid, id string) error {
	doc := r.doc(uid, id)
	return r.client.RunTransaction(ctx, func(_ context.Context, tx *firestore.Transaction) error {
		if _, err := tx.Get(doc); err != nil {
			if status.Code(err) == codes.NotFound {
				return ErrNotFound
			}
			return err
		}
		return tx.Delete(doc)
	})
}

// Update writes the two fields the user owns: the name and the favorite
// mark (PR-17). A nil pointer leaves that field as it is. The write
// carries both the flat field and the packed proto, so a later Get reads
// the same value the list shows.
func (r *Repo) Update(ctx context.Context, uid, id string, name *string, favorite *bool) (*mtgv1.Deck, error) {
	doc := r.doc(uid, id)
	var out *mtgv1.Deck
	err := r.client.RunTransaction(ctx, func(_ context.Context, tx *firestore.Transaction) error {
		snap, err := tx.Get(doc)
		if status.Code(err) == codes.NotFound {
			return ErrNotFound
		}
		if err != nil {
			return err
		}
		var sd storedDeck
		if err := snap.DataTo(&sd); err != nil {
			return err
		}
		var d mtgv1.Deck
		if err := gzstore.UnmarshalProto(sd.DeckGz, &d); err != nil {
			return err
		}
		if name != nil {
			d.Name = *name
		}
		if favorite != nil {
			d.Favorite = *favorite
		}
		payload, err := gzstore.MarshalProto(&d)
		if err != nil {
			return err
		}
		if len(payload) > gzstore.MaxStoredBytes {
			return ErrTooLarge
		}
		updated := toStored(&d, payload)
		// The stored document keeps the create time it already had. A
		// rename must not move the deck to the top of the list.
		updated.CreatedAt = sd.CreatedAt
		if err := tx.Set(doc, updated); err != nil {
			return err
		}
		out = &d
		return nil
	})
	if err != nil {
		return nil, err
	}
	return out, nil
}
