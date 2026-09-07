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
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	mtgv1 "github.com/nkramber/mtg-deck-builder/go/gen/mtg/v1"
)

// schemaVersion counts the stored shape.
const schemaVersion = 1

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
	Schema     int64            `firestore:"schema"`
	Kind       string           `firestore:"kind"`
	Verdict    string           `firestore:"verdict"`
	SessionID  string           `firestore:"session_id"`
	QuestionID string           `firestore:"question_id"`
	DeckID     string           `firestore:"deck_id"`
	OracleID   string           `firestore:"oracle_id"`
	Reasons    []string         `firestore:"reasons"`
	Text       string           `firestore:"text"`
	Prompts    map[string]int64 `firestore:"prompts"`
	CreatedAt  time.Time        `firestore:"created_at"`
}

func (r *Repo) col(uid string) *firestore.CollectionRef {
	return r.client.Collection("users").Doc(uid).Collection("feedback")
}

// Add writes one verdict under the user and answers its id.
func (r *Repo) Add(ctx context.Context, uid string, item Item) (string, error) {
	doc := r.col(uid).NewDoc()
	_, err := doc.Create(ctx, stored{
		Schema:     schemaVersion,
		Kind:       item.Kind,
		Verdict:    item.Verdict,
		SessionID:  item.SessionID,
		QuestionID: item.QuestionID,
		DeckID:     item.DeckID,
		OracleID:   item.OracleID,
		Reasons:    item.Reasons,
		Text:       item.Text,
		Prompts:    item.Prompts,
		CreatedAt:  item.CreatedAt.UTC(),
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
	return Item{
		Kind:       s.Kind,
		Verdict:    s.Verdict,
		SessionID:  s.SessionID,
		QuestionID: s.QuestionID,
		DeckID:     s.DeckID,
		OracleID:   s.OracleID,
		Reasons:    s.Reasons,
		Text:       s.Text,
		Prompts:    s.Prompts,
		CreatedAt:  s.CreatedAt,
	}, nil
}

// kindNames are the short names the store writes. An enum value off the
// table has no name, and the API refuses it.
var kindNames = map[mtgv1.FeedbackKind]string{
	mtgv1.FeedbackKind_FEEDBACK_KIND_QUESTION: "question",
	mtgv1.FeedbackKind_FEEDBACK_KIND_SUMMARY:  "summary",
	mtgv1.FeedbackKind_FEEDBACK_KIND_CARD:     "card",
	mtgv1.FeedbackKind_FEEDBACK_KIND_DECK:     "deck",
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
