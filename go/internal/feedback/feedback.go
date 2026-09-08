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
	"time"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
)

// schemaVersion counts the stored shape.
const schemaVersion = 2

// ErrNotFound reports a feedback id no document answers.
var ErrNotFound = errors.New("feedback not found")

// Item is one verdict as the store holds it. Kind and Verdict are the
// short names of the proto enums, so the harvest reads them without the
// proto: question, summary, card, deck, and up, down.
type Item struct {
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
}

func (r *Repo) col(uid string) *firestore.CollectionRef {
	return r.client.Collection("users").Doc(uid).Collection("feedback")
}

// Add writes one verdict under the user and answers its id.
func (r *Repo) Add(ctx context.Context, uid string, item Item) (string, error) {
	doc := r.col(uid).NewDoc()
	_, err := doc.Create(ctx, stored{
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
	return itemOf(s), nil
}

// itemOf reads a stored document as an Item.
func itemOf(s stored) Item {
	return Item{
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
		out = append(out, itemOf(s))
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
