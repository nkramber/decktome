// Package feedback keeps the verdicts a user gives on what the app made
// (PR-27, D-557 to D-559). One Firestore document per verdict,
// users/<uid>/feedback/<id>, holds the kind, the thumb, the ids of the
// object, the reason keys, the text, the prompt versions of the moment,
// and the time. The API writes it, and the harvest of PR-28 reads the
// collection with the owner's credentials (D-420). No RPC lists it.
package feedback

import (
	"context"
	"errors"
	"fmt"
	"slices"
	"strings"
	"time"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
	"github.com/nkramber/decktome/go/internal/gzstore"
)

// schemaVersion counts the stored shape. Version 3 fills session_gz on
// a verdict about a deck as well, because the slots that made the deck
// live on the session (D-643).
const schemaVersion = 3

// ErrNotFound reports a feedback id no document answers.
var ErrNotFound = errors.New("feedback not found")

// Item is one verdict as the store holds it. Kind and Verdict are the
// short names of the proto enums, so the harvest reads them without the
// proto: question, summary, card, deck, and up, down.
type Item struct {
	// ID is the document id of the verdict. Add answers it, and every
	// read fills it, so a harvest names the row it wrote. Add itself
	// ignores this field: the store owns the id.
	ID         string
	Kind       string
	Verdict    string
	SessionID  string
	QuestionID string
	DeckID     string
	OracleID   string
	Reasons    []string
	Text       string
	// UID is the user the verdict came from (D-596). The document sits
	// under that user, and a query over every user answers documents and
	// not paths, so the id is a field of its own as well.
	UID string
	// QuestionText is the exact wording the reader saw, and AnswerText is
	// what the reader had answered when the verdict came (D-596). The
	// question id alone reads nothing after a prompt changes. The server
	// fills both from the stored session, so neither is the client's
	// word.
	QuestionText string
	AnswerText   string
	// Prompts holds the prompt versions of the moment by role family,
	// with the keys of the eval run files: questions and generate.
	Prompts   map[string]int64
	CreatedAt time.Time
	// Session and Deck hold the object the verdict names, as it stood
	// when the reader gave the verdict (D-635). A reader deletes a
	// session or a deck, and the verdict outlives it, so a harvest that
	// reads the store finds nothing. The server fills these from the
	// object it already loaded to check the owner, so neither is the
	// client's word. Either one is nil when the verdict names no such
	// object, and both are nil on a document written before D-635.
	Session *mtgv1.Session
	Deck    *mtgv1.Deck
}

// Repo stores feedback in Firestore. The caller owns the client.
type Repo struct {
	client *firestore.Client
}

// NewRepo wraps a Firestore client.
func NewRepo(client *firestore.Client) *Repo { return &Repo{client: client} }

type stored struct {
	Schema       int64            `firestore:"schema"`
	Kind         string           `firestore:"kind"`
	Verdict      string           `firestore:"verdict"`
	UID          string           `firestore:"uid"`
	SessionID    string           `firestore:"session_id"`
	QuestionID   string           `firestore:"question_id"`
	QuestionText string           `firestore:"question_text"`
	AnswerText   string           `firestore:"answer_text"`
	DeckID       string           `firestore:"deck_id"`
	OracleID     string           `firestore:"oracle_id"`
	Reasons      []string         `firestore:"reasons"`
	Text         string           `firestore:"text"`
	Prompts      map[string]int64 `firestore:"prompts"`
	CreatedAt    time.Time        `firestore:"created_at"`
	// SessionGz and DeckGz hold the snapshot of D-635, as gzip protojson,
	// the way the session and the deck stores hold their own (D-604).
	SessionGz []byte `firestore:"session_gz"`
	DeckGz    []byte `firestore:"deck_gz"`
}

func (r *Repo) col(uid string) *firestore.CollectionRef {
	return r.client.Collection("users").Doc(uid).Collection("feedback")
}

// Add writes one verdict under the user and answers its id.
func (r *Repo) Add(ctx context.Context, uid string, item Item) (string, error) {
	sessionGz, deckGz, err := snapshotOf(item)
	if err != nil {
		return "", err
	}
	doc := r.col(uid).NewDoc()
	_, err = doc.Create(ctx, stored{
		Schema:       schemaVersion,
		Kind:         item.Kind,
		Verdict:      item.Verdict,
		UID:          uid,
		SessionID:    item.SessionID,
		QuestionID:   item.QuestionID,
		QuestionText: item.QuestionText,
		AnswerText:   item.AnswerText,
		DeckID:       item.DeckID,
		OracleID:     item.OracleID,
		Reasons:      item.Reasons,
		Text:         item.Text,
		Prompts:      item.Prompts,
		CreatedAt:    item.CreatedAt.UTC(),
		SessionGz:    sessionGz,
		DeckGz:       deckGz,
	})
	if err != nil {
		return "", fmt.Errorf("feedback: %w", err)
	}
	return doc.ID, nil
}

// Get reads one verdict of the user, or ErrNotFound.
func (r *Repo) Get(ctx context.Context, uid, id string) (Item, error) {
	snap, err := r.col(uid).Doc(id).Get(ctx)
	if status.Code(err) == codes.NotFound {
		return Item{}, ErrNotFound
	}
	if err != nil {
		return Item{}, err
	}
	var s stored
	if err := snap.DataTo(&s); err != nil {
		return Item{}, fmt.Errorf("feedback %s: %w", id, err)
	}
	return itemOf(id, s), nil
}

// snapshotOf encodes the objects the verdict names (D-635). A snapshot
// that passes the Firestore room of one document drops, because a
// verdict the store refuses helps nobody: the words of the reader are
// worth more than the context.
func snapshotOf(item Item) (sessionGz, deckGz []byte, err error) {
	if item.Session != nil {
		if sessionGz, err = gzstore.MarshalProto(item.Session); err != nil {
			return nil, nil, fmt.Errorf("feedback: %w", err)
		}
		if len(sessionGz) > gzstore.MaxStoredBytes {
			sessionGz = nil
		}
	}
	if item.Deck != nil {
		if deckGz, err = gzstore.MarshalProto(item.Deck); err != nil {
			return nil, nil, fmt.Errorf("feedback: %w", err)
		}
		if len(deckGz) > gzstore.MaxStoredBytes {
			deckGz = nil
		}
	}
	// A verdict on a deck carries both snapshots since D-643, and
	// Firestore caps one document at 1 MiB. The deck is the object the
	// verdict names, so the session is the one that drops when the two
	// together pass the room of one document.
	if len(sessionGz)+len(deckGz) > gzstore.MaxStoredBytes {
		sessionGz = nil
	}
	return sessionGz, deckGz, nil
}

// itemOf reads a stored document as an Item. A snapshot that does not
// open leaves its field nil, so one bad blob never hides a verdict.
func itemOf(id string, s stored) Item {
	var sess *mtgv1.Session
	if len(s.SessionGz) > 0 {
		var m mtgv1.Session
		if err := gzstore.UnmarshalProto(s.SessionGz, &m); err == nil {
			sess = &m
		}
	}
	var deck *mtgv1.Deck
	if len(s.DeckGz) > 0 {
		var m mtgv1.Deck
		if err := gzstore.UnmarshalProto(s.DeckGz, &m); err == nil {
			deck = &m
		}
	}
	return Item{
		ID:           id,
		Session:      sess,
		Deck:         deck,
		Kind:         s.Kind,
		Verdict:      s.Verdict,
		UID:          s.UID,
		SessionID:    s.SessionID,
		QuestionID:   s.QuestionID,
		QuestionText: s.QuestionText,
		AnswerText:   s.AnswerText,
		DeckID:       s.DeckID,
		OracleID:     s.OracleID,
		Reasons:      s.Reasons,
		Text:         s.Text,
		Prompts:      s.Prompts,
		CreatedAt:    s.CreatedAt,
	}
}

// Find reads one verdict by its id, over every user (D-600).
//
// A collection group query needs a composite index, and an index of a
// deployed project is a deploy of its own. This walks the users instead,
// so a reader of one id needs no index at all. The user list is short,
// and one document read answers each user.
func (r *Repo) Find(ctx context.Context, id string) (Item, string, error) {
	// DocumentRefs and not Documents. The app writes
	// users/<uid>/feedback/<id> and never gives users/<uid> a field of
	// its own, so that parent is a missing document. Documents leaves a
	// missing document out, and this read then finds nothing at all.
	iter := r.client.Collection("users").DocumentRefs(ctx)
	for {
		ref, err := iter.Next()
		if errors.Is(err, iterator.Done) {
			break
		}
		if err != nil {
			return Item{}, "", fmt.Errorf("feedback: users: %w", err)
		}
		uid := ref.ID
		item, err := r.Get(ctx, uid, id)
		if errors.Is(err, ErrNotFound) {
			continue
		}
		if err != nil {
			return Item{}, "", err
		}
		if item.UID == "" {
			item.UID = uid
		}
		return item, uid, nil
	}
	return Item{}, "", ErrNotFound
}

// Down reads the newest verdicts of every user, over one collection
// group query (D-596). The documents sit under each user, and a group
// query reads them all at once, so no caller walks the user list.
//
// A down verdict is what the fixer of PR-28 reads, and `down` alone is
// the common case. An empty verdict reads both.
func (r *Repo) Down(ctx context.Context, verdict string, limit int) ([]Item, error) {
	q := r.client.CollectionGroup("feedback").OrderBy("created_at", firestore.Desc).Limit(limit)
	if verdict != "" {
		q = r.client.CollectionGroup("feedback").
			Where("verdict", "==", verdict).
			OrderBy("created_at", firestore.Desc).Limit(limit)
	}
	snaps, err := q.Documents(ctx).GetAll()
	if err != nil {
		return nil, fmt.Errorf("feedback: %w", err)
	}
	out := make([]Item, 0, len(snaps))
	for _, snap := range snaps {
		var s stored
		if err := snap.DataTo(&s); err != nil {
			return nil, fmt.Errorf("feedback %s: %w", snap.Ref.ID, err)
		}
		// A document written before the uid field reads its user from the
		// path: users/<uid>/feedback/<id>.
		if s.UID == "" {
			s.UID = snap.Ref.Parent.Parent.ID
		}
		out = append(out, itemOf(snap.Ref.ID, s))
	}
	return out, nil
}

// Since reads every verdict the store holds after t, oldest first, both
// verdicts together (PR-28a).
//
// It runs one query per verdict and merges them. A collection group
// query needs an index for its shape, and a single-field index of that
// scope is not automatic: the store holds no index for created_at
// alone, in either direction. It does hold "verdict ascending,
// created_at descending", which serves an equality on the verdict with
// a range and an order on the time. So two indexed queries cost nothing
// on the deployed project, and one unindexed query costs an index
// deploy. Find carries the note on why that matters.
//
// A zero t reads every verdict the store holds. A limit bounds what the
// caller writes and never what this reads, and it keeps the oldest rows
// after t, so a harvest advances its watermark one chunk at a time and
// leaves no gap under it.
func (r *Repo) Since(ctx context.Context, t time.Time, limit int) ([]Item, error) {
	var out []Item
	for _, verdict := range []string{"down", "up"} {
		q := r.client.CollectionGroup("feedback").
			Where("verdict", "==", verdict).
			OrderBy("created_at", firestore.Desc)
		if !t.IsZero() {
			q = r.client.CollectionGroup("feedback").
				Where("verdict", "==", verdict).
				Where("created_at", ">", t.UTC()).
				OrderBy("created_at", firestore.Desc)
		}
		snaps, err := q.Documents(ctx).GetAll()
		if err != nil {
			return nil, fmt.Errorf("feedback: %w", err)
		}
		for _, snap := range snaps {
			var st stored
			if err := snap.DataTo(&st); err != nil {
				return nil, fmt.Errorf("feedback %s: %w", snap.Ref.ID, err)
			}
			// A document written before the uid field reads its user from
			// the path: users/<uid>/feedback/<id>.
			if st.UID == "" {
				st.UID = snap.Ref.Parent.Parent.ID
			}
			out = append(out, itemOf(snap.Ref.ID, st))
		}
	}
	// Oldest first, so a harvest document tells the story in order. The
	// id breaks a tie, so two verdicts of one instant read the same way
	// on every run.
	slices.SortStableFunc(out, func(a, b Item) int {
		if !a.CreatedAt.Equal(b.CreatedAt) {
			return a.CreatedAt.Compare(b.CreatedAt)
		}
		return strings.Compare(a.ID, b.ID)
	})
	// The oldest, and never the newest. The watermark of a harvest is the
	// newest time it wrote, so a chunk must be the oldest rows after the
	// floor. A chunk of the newest rows moves the floor past every row
	// under it, and no later harvest ever reads those (PR-28a, F-88).
	//
	// The read itself takes no limit for the same reason: a descending
	// read of n answers the newest n, and the oldest rows after the floor
	// are the ones this must not miss.
	if limit > 0 && len(out) > limit {
		out = out[:limit]
	}
	return out, nil
}

// kindNames are the short names the store writes. An enum value off the
// table has no name, and the API refuses it.
var kindNames = map[mtgv1.FeedbackKind]string{
	mtgv1.FeedbackKind_FEEDBACK_KIND_QUESTION: "question",
	mtgv1.FeedbackKind_FEEDBACK_KIND_SUMMARY:  "summary",
	mtgv1.FeedbackKind_FEEDBACK_KIND_CARD:     "card",
	mtgv1.FeedbackKind_FEEDBACK_KIND_DECK:     "deck",
	mtgv1.FeedbackKind_FEEDBACK_KIND_CHAT:     "chat",
}

var verdictNames = map[mtgv1.FeedbackVerdict]string{
	mtgv1.FeedbackVerdict_FEEDBACK_VERDICT_UP:   "up",
	mtgv1.FeedbackVerdict_FEEDBACK_VERDICT_DOWN: "down",
}

// KindName is the short name of a kind, or "" for one the store does
// not know.
func KindName(k mtgv1.FeedbackKind) string { return kindNames[k] }

// VerdictName is the short name of a verdict, or "" for one the store
// does not know.
func VerdictName(v mtgv1.FeedbackVerdict) string { return verdictNames[v] }

// reasonKeys are the prefilled reasons per kind, section 2.4 of the
// design note. The web app holds the labels, and the harvest of PR-28
// reads the keys. The two sides pin the same list in a test each.
var reasonKeys = map[mtgv1.FeedbackKind][]string{
	mtgv1.FeedbackKind_FEEDBACK_KIND_QUESTION: {"already_answered", "not_applicable", "bad_options", "unclear"},
	mtgv1.FeedbackKind_FEEDBACK_KIND_SUMMARY:  {"false_claim", "misses_plan", "too_long_or_vague"},
	mtgv1.FeedbackKind_FEEDBACK_KIND_CARD:     {"off_theme", "illegal", "unwanted_buy", "wrong_printing", "wrong_power"},
	mtgv1.FeedbackKind_FEEDBACK_KIND_DECK:     {"off_spec", "bad_mana", "too_little_interaction", "wrong_power", "too_many_to_buy"},
	// A chat that stops belongs to no question, so it has a kind and a
	// reason set of its own (D-594).
	mtgv1.FeedbackKind_FEEDBACK_KIND_CHAT: {"stuck", "ignored_request", "wrong_questions", "no_deck", "error"},
}

// Reasons lists the reason keys of a kind, in the order the dialog
// shows them. An unknown kind has none.
func Reasons(k mtgv1.FeedbackKind) []string {
	return append([]string(nil), reasonKeys[k]...)
}

// KnownReason reports whether key is a reason of the kind.
func KnownReason(k mtgv1.FeedbackKind, key string) bool {
	for _, r := range reasonKeys[k] {
		if r == key {
			return true
		}
	}
	return false
}
