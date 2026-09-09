package triage

import (
	"encoding/json"
	"fmt"
	"strings"

	"google.golang.org/protobuf/encoding/protojson"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
	"github.com/nkramber/decktome/go/internal/harvest"
	"github.com/nkramber/decktome/go/internal/questions"
)

// The artifact writers (PR-28b). Each class writes a case in the shape
// of the file that owns it, and the case joins that file. The owner
// reads it on the pull request, which is the accept step (D-642).
//
// A case carries the reader's own words, and never an email (D-559). It
// carries no uid and no verdict id either: the harvest holds those, and
// a gate file has no use for them.

// The files a case joins, from the repo root.
const (
	TargetConversations  = "go/cmd/questions-gate/conversations.json"
	TargetDeckPrompts    = "go/cmd/deck-gate/prompts.json"
	TargetBracketPrompts = "go/cmd/bracket-gate/prompts.json"
)

// Case is one artifact the triage wrote.
type Case struct {
	Class    string   `json:"class"`
	Artifact Artifact `json:"artifact"`
	// Target is the file the body joins, or empty for a defect.
	Target string `json:"target,omitempty"`
	// Body is the case in the shape of that file.
	Body json.RawMessage `json:"body,omitempty"`
	// Gaps name what the verdict could not fill. A person reads them on
	// the pull request and finishes the case by hand.
	Gaps []string `json:"gaps,omitempty"`
	// Detail is the defect in words, for a class that writes no case.
	Detail string `json:"detail,omitempty"`
}

// Namer reads a card name from an Oracle id. A nil Namer leaves every
// card name out, and the case then names the gap.
type Namer func(oracleID string) (string, bool)

// conversationCase mirrors one conversation of the question gate. The
// gate owns the shape, and this package can not import a command, so it
// holds a mirror.
//
// Each gate package pins its own mirror. TestTheCaseShapeMatchesTheGateFile
// in cmd/questions-gate, and TestTheCaseShapeMatchesThePromptFile in
// cmd/deck-gate and cmd/bracket-gate, decode a case of this writer into
// the real struct with an unknown field refused. So a renamed tag fails
// a free test, and never in a paid run.
type conversationCase struct {
	ID         int               `json:"id"`
	Name       string            `json:"name"`
	Collection bool              `json:"collection,omitempty"`
	Messages   []string          `json:"messages"`
	Note       string            `json:"note,omitempty"`
	Expect     map[string]string `json:"expect,omitempty"`
	MustNotAsk []string          `json:"must_not_ask,omitempty"`
}

// deckPromptCase mirrors one prompt of the deck gate.
type deckPromptCase struct {
	ID             int      `json:"id"`
	Name           string   `json:"name"`
	Format         string   `json:"format"`
	Theme          string   `json:"theme,omitempty"`
	Colors         []string `json:"colors,omitempty"`
	Commander      string   `json:"commander,omitempty"`
	Bracket        int32    `json:"bracket,omitempty"`
	Power          string   `json:"power,omitempty"`
	Pool           string   `json:"pool,omitempty"`
	Collection     bool     `json:"collection,omitempty"`
	Budget         float64  `json:"budget,omitempty"`
	Sets           []string `json:"sets,omitempty"`
	MustNotInclude []string `json:"must_not_include,omitempty"`
	MustOwnAll     bool     `json:"must_own_all,omitempty"`
	Plan           string   `json:"plan"`
}

// bracketPromptCase mirrors one prompt of the bracket gate.
type bracketPromptCase struct {
	ID        int    `json:"id"`
	Bracket   int32  `json:"bracket"`
	Theme     string `json:"theme"`
	Commander string `json:"commander"`
	Plan      string `json:"plan"`
}

// CaseOf writes the artifact of one route. id is the number the case
// takes in its file, and name reads card names, or is nil. A route that
// still needs the judge writes no case.
func CaseOf(r Route, id int, name Namer) (Case, error) {
	if r.Need == NeedJudge {
		return Case{}, fmt.Errorf("triage %s: the route still needs the judge", r.Record.ID)
	}
	if r.Owner {
		return Case{Class: r.Class.ID, Artifact: ADefect,
			Detail: "the reader argues with " + r.Decision + ", so it takes a row in docs/owner-questions.md and no case"}, nil
	}
	if r.Keep {
		return keepCase(r, id, name)
	}
	switch r.Class.Artifact {
	case AConversation:
		return conversationOf(r, id, name)
	case ADeckPrompt, ASummaryCase:
		return deckPromptOf(r, id, name)
	case ABracketPrompt:
		return bracketPromptOf(r, id, name)
	case ADefect:
		return defectOf(r), nil
	}
	return Case{}, fmt.Errorf("triage %s: no class, so no artifact", r.Record.ID)
}

// keepCase writes the case a thumbs up makes: the same shape the kind's
// down classes write, with no assertion. A change must not flip it.
func keepCase(r Route, id int, name Namer) (Case, error) {
	switch r.Record.Kind {
	case "question", "chat":
		return conversationOf(r, id, name)
	case "summary", "card", "deck":
		return deckPromptOf(r, id, name)
	}
	return Case{}, fmt.Errorf("triage %s: kind %q has no keep case", r.Record.ID, r.Record.Kind)
}

// conversationOf writes a golden conversation from the stored session.
// The reader's own messages are the conversation, and the row the
// verdict names becomes the "must not ask" expectation for the two
// classes that ask for one (D-643).
func conversationOf(r Route, id int, name Namer) (Case, error) {
	rec := r.Record
	sess, err := SessionOf(rec)
	if err != nil {
		return Case{}, err
	}
	c := conversationCase{ID: id, Name: caseName(r), Note: noteOf(r)}
	var gaps []string
	if sess == nil {
		gaps = append(gaps, "the verdict carries no session, so the case holds no message and no expectation")
	} else {
		c.Collection = sess.GetCollectionId() != ""
		for _, t := range sess.GetTurns() {
			if msg := strings.TrimSpace(t.GetUserMessage()); msg != "" {
				c.Messages = append(c.Messages, msg)
			}
		}
		c.Expect = expectOf(sess, name)
		if len(c.Messages) == 0 {
			gaps = append(gaps, "the session holds no message of the reader, so the case asks nothing")
		}
		if _, ok := c.Expect["commander"]; !ok && len(sess.GetSlots().GetCommanderOracleIds()) > 0 {
			gaps = append(gaps, "the commander is an Oracle id and no card index named it")
		}
	}
	// Q1 and Q2 are the two classes whose fix is a row that must stop
	// firing. Q3 and Q4 want a better wording or better options on a row
	// that should still fire, and the gate reads no expectation of that
	// shape yet.
	if r.Class.ID == "Q1" || r.Class.ID == "Q2" {
		if row := RowOf(rec.QuestionID); row != "" {
			c.MustNotAsk = []string{row}
		} else {
			gaps = append(gaps, "the verdict names no question id, so no row could be forbidden")
		}
	}
	if r.Class.ID == "Q3" || r.Class.ID == "Q4" {
		gaps = append(gaps, "the fix is the wording or the options of row "+orUnknown(RowOf(rec.QuestionID))+", and the gate reads no expectation of that shape")
	}
	if r.Class.ID == "X3" {
		gaps = append(gaps, "a chat verdict names no question, so a person names the row that must not fire")
	}
	body, err := json.Marshal(c)
	if err != nil {
		return Case{}, fmt.Errorf("triage %s: %w", rec.ID, err)
	}
	return Case{Class: classID(r), Artifact: AConversation, Target: TargetConversations, Body: body, Gaps: gaps}, nil
}

// deckPromptOf writes a deck gate prompt from the deck the reader judged
// and the session that built it. The deck holds the format, the power,
// and the commander, and the slots of the session hold the theme, the
// pool rule, the budget, and the sets (D-643).
func deckPromptOf(r Route, id int, name Namer) (Case, error) {
	rec := r.Record
	deck, err := DeckOf(rec)
	if err != nil {
		return Case{}, err
	}
	sess, err := SessionOf(rec)
	if err != nil {
		return Case{}, err
	}
	p := deckPromptCase{ID: id, Name: caseName(r)}
	var gaps []string
	if deck == nil {
		gaps = append(gaps, "the verdict carries no deck, so the prompt holds no format and no commander")
	} else {
		p.Format = strings.ToLower(formatWord(deck.GetFormat().GetId()))
		if b, ok := deck.GetPower().GetLevel().(*mtgv1.PowerLevel_Bracket); ok {
			p.Bracket = b.Bracket
		}
		if s, ok := deck.GetPower().GetLevel().(*mtgv1.PowerLevel_SixtyStep); ok {
			p.Power = strings.ToLower(strings.TrimPrefix(s.SixtyStep.String(), "SIXTY_STEP_"))
		}
		p.Commander = strings.Join(cardNames(deck.GetCommanderOracleIds(), name), " + ")
		if p.Commander == "" && len(deck.GetCommanderOracleIds()) > 0 {
			gaps = append(gaps, "the commander is an Oracle id and no card index named it")
		}
	}
	if sess == nil {
		gaps = append(gaps, "the verdict carries no session, so the prompt holds no theme, no pool rule, and no budget")
	} else {
		words := expectOf(sess, name)
		p.Theme = words["theme"]
		p.Pool = words["pool_rule"]
		p.Collection = sess.GetCollectionId() != ""
		p.Budget = sess.GetSlots().GetBudgetUsd()
		p.Sets = sess.GetSlots().GetSetCodes()
		if p.Theme == "" {
			gaps = append(gaps, "the session names no theme")
		}
	}
	// The assertion is the fault the reader met. A card the reader called
	// off theme must not come back, and a reader who did not want to buy
	// asks for a deck they own whole. C2 and C4 write a defect and never
	// reach here: a rules fault and a printing fault take a code fix, and
	// no gate of this repo measures either one yet.
	switch r.Class.ID {
	case "C1":
		if n := cardName(rec.OracleID, name, deck); n != "" {
			p.MustNotInclude = []string{n}
		} else {
			gaps = append(gaps, "the card is an Oracle id and no card index named it, so the prompt forbids nothing")
		}
	case "C3", "D5":
		p.MustOwnAll = true
		if p.Pool == "" {
			p.Pool = "owned_first"
		}
		p.Collection = true
	}
	p.Plan = planOf(r, p)
	body, err := json.Marshal(p)
	if err != nil {
		return Case{}, fmt.Errorf("triage %s: %w", rec.ID, err)
	}
	art := r.Class.Artifact
	if r.Keep {
		art = ADeckPrompt
	}
	return Case{Class: classID(r), Artifact: art, Target: TargetDeckPrompts, Body: body, Gaps: gaps}, nil
}

// bracketPromptOf writes a bracket gate prompt. The judge of that gate
// names the bracket a deck plays at, which is the bar a power complaint
// asks for.
func bracketPromptOf(r Route, id int, name Namer) (Case, error) {
	rec := r.Record
	deck, err := DeckOf(rec)
	if err != nil {
		return Case{}, err
	}
	sess, err := SessionOf(rec)
	if err != nil {
		return Case{}, err
	}
	p := bracketPromptCase{ID: id}
	var gaps []string
	if deck == nil {
		gaps = append(gaps, "the verdict carries no deck, so the prompt holds no bracket and no commander")
	} else {
		if b, ok := deck.GetPower().GetLevel().(*mtgv1.PowerLevel_Bracket); ok {
			p.Bracket = b.Bracket
		}
		p.Commander = strings.Join(cardNames(deck.GetCommanderOracleIds(), name), " + ")
		if p.Commander == "" {
			gaps = append(gaps, "the prompt names no commander, and the bracket gate needs one")
		}
	}
	if p.Bracket == 0 {
		gaps = append(gaps, "the deck names no bracket, so a person sets the one the reader played at")
	}
	if sess != nil {
		p.Theme = strings.TrimSpace(sess.GetSlots().GetTheme())
	}
	if p.Theme == "" {
		gaps = append(gaps, "the session names no theme, and the bracket gate needs one")
	}
	p.Plan = fmt.Sprintf("A %s Commander deck at bracket %d. %s", orUnknown(p.Theme), p.Bracket, readerSaid(rec))
	body, err := json.Marshal(p)
	if err != nil {
		return Case{}, fmt.Errorf("triage %s: %w", rec.ID, err)
	}
	return Case{Class: classID(r), Artifact: ABracketPrompt, Target: TargetBracketPrompts, Body: body, Gaps: gaps}, nil
}

// defectOf writes the row of a class that no gate measures. The fix is a
// code fix, and a person writes the test.
func defectOf(r Route) Case {
	rec := r.Record
	var b strings.Builder
	fmt.Fprintf(&b, "%s, %s. The fix lives in %s. ", r.Class.ID, r.Class.Name, r.Class.Where)
	if rec.OracleID != "" {
		fmt.Fprintf(&b, "Card %s. ", rec.OracleID)
	}
	b.WriteString(readerSaid(rec))
	return Case{Class: r.Class.ID, Artifact: ADefect, Detail: strings.TrimSpace(b.String())}
}

// SessionOf reads the session snapshot of a verdict, or nil when it
// carries none (D-635).
func SessionOf(rec harvest.Record) (*mtgv1.Session, error) {
	if len(rec.Session) == 0 {
		return nil, nil
	}
	var s mtgv1.Session
	if err := protojson.Unmarshal(rec.Session, &s); err != nil {
		return nil, fmt.Errorf("triage %s: session: %w", rec.ID, err)
	}
	return &s, nil
}

// DeckOf reads the deck snapshot of a verdict, or nil when it carries
// none (D-635).
func DeckOf(rec harvest.Record) (*mtgv1.Deck, error) {
	if len(rec.Deck) == 0 {
		return nil, nil
	}
	var d mtgv1.Deck
	if err := protojson.Unmarshal(rec.Deck, &d); err != nil {
		return nil, fmt.Errorf("triage %s: deck: %w", rec.ID, err)
	}
	return &d, nil
}

// expectOf renders the slots of a stored session in the words a golden
// expectation uses. The renderer is the gate's own (questions.SlotWords),
// so a case can not name a word the gate does not read. An empty value
// is dropped: an expectation names what a slot must end with, and a key
// with no value asserts nothing.
func expectOf(sess *mtgv1.Session, name Namer) map[string]string {
	commanders := cardNames(sess.GetSlots().GetCommanderOracleIds(), name)
	locked := cardNames(sess.GetSlots().GetLockedOracleIds(), name)
	words := questions.SlotWords(sess.GetSlots(), commanders, lockedOnly(locked, commanders), sess.GetCollectionId() != "")
	out := map[string]string{}
	for k, v := range words {
		if strings.TrimSpace(v) != "" {
			out[k] = v
		}
	}
	return out
}

// lockedOnly drops a locked card that is also the commander (D-70).
func lockedOnly(locked, commanders []string) []string {
	var out []string
	for _, n := range locked {
		if !contains(commanders, n) {
			out = append(out, n)
		}
	}
	return out
}

func contains(list []string, s string) bool {
	for _, v := range list {
		if strings.EqualFold(v, s) {
			return true
		}
	}
	return false
}

// cardNames reads names for Oracle ids. An id the namer does not hold is
// dropped, and the caller names the gap.
func cardNames(ids []string, name Namer) []string {
	if name == nil {
		return nil
	}
	var out []string
	for _, id := range ids {
		if n, ok := name(id); ok {
			out = append(out, n)
		}
	}
	return out
}

// cardName reads one card name. It tries the index, then the deck the
// reader judged, which carries the name beside the id.
func cardName(oracleID string, name Namer, deck *mtgv1.Deck) string {
	if oracleID == "" {
		return ""
	}
	if name != nil {
		if n, ok := name(oracleID); ok {
			return n
		}
	}
	for _, dc := range append(append([]*mtgv1.DeckCard{}, deck.GetCards()...), deck.GetSideboard()...) {
		if dc.GetOracleId() == oracleID {
			return dc.GetName()
		}
	}
	return ""
}

// caseName is the one-line name the gate document shows.
func caseName(r Route) string {
	if r.Keep {
		return fmt.Sprintf("keep: a reader liked this %s", r.Record.Kind)
	}
	return fmt.Sprintf("%s: %s", r.Class.ID, r.Class.Name)
}

// classID is the class of a case, and "keep" for a thumbs up.
func classID(r Route) string {
	if r.Keep {
		return "keep"
	}
	return r.Class.ID
}

// noteOf says why the case exists, in the reader's own words. A gate
// file carries the note, so the next reader knows what the case pins.
func noteOf(r Route) string {
	if r.Keep {
		return "A reader gave this a thumbs up. A change must not flip it."
	}
	return fmt.Sprintf("From a reader's thumbs down, %s. %s %s",
		r.Record.CreatedAt.Format("2006-01-02"), r.Class.Name+".", readerSaid(r.Record))
}

// readerSaid writes what the reader checked and typed, in their words.
func readerSaid(rec harvest.Record) string {
	var parts []string
	if len(rec.Reasons) > 0 {
		parts = append(parts, "Reasons: "+strings.Join(rec.Reasons, ", ")+".")
	}
	if t := strings.TrimSpace(rec.Text); t != "" {
		parts = append(parts, "The reader wrote: "+t)
	}
	return strings.Join(parts, " ")
}

// planOf writes the prose plan a deck gate prompt carries, in the shape
// the other prompts use.
func planOf(r Route, p deckPromptCase) string {
	var b strings.Builder
	fmt.Fprintf(&b, "A %s %s deck", orUnknown(p.Theme), formatTitle(p.Format))
	if p.Commander != "" {
		fmt.Fprintf(&b, " led by %s", p.Commander)
	}
	if p.Bracket > 0 {
		fmt.Fprintf(&b, ", at bracket %d", p.Bracket)
	}
	if p.MustOwnAll {
		b.WriteString(", from my library only")
	}
	b.WriteString(". ")
	b.WriteString(readerSaid(r.Record))
	return strings.TrimSpace(b.String())
}

// formatTitle writes a format the way a reader says it in a plan.
func formatTitle(word string) string {
	switch word {
	case "commander":
		return "Commander"
	case "standard":
		return "Standard"
	case "modern":
		return "Modern"
	case "house":
		return "house-rules"
	}
	return "(unknown format)"
}

func formatWord(f mtgv1.FormatId) string {
	switch f {
	case mtgv1.FormatId_FORMAT_ID_COMMANDER:
		return "commander"
	case mtgv1.FormatId_FORMAT_ID_STANDARD:
		return "standard"
	case mtgv1.FormatId_FORMAT_ID_MODERN:
		return "modern"
	case mtgv1.FormatId_FORMAT_ID_HOUSE:
		return "house"
	}
	return ""
}

func orUnknown(s string) string {
	if strings.TrimSpace(s) == "" {
		return "(unknown)"
	}
	return s
}
