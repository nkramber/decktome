// Command api-build builds one deck over the deployed API alone, with
// no browser and no GUI (D-778).
//
// It signs in with an email and a password, imports a collection,
// starts a chat, answers every question the agent asks, and reads the
// finished deck back out of storage. A session can then prove the whole
// flow on the deployed app, which the live check of #196 could not: the
// sandbox refuses every list under users/<uid> (D-638).
//
// CAUTION: this calls the deployed API, and the deployed API calls the
// real providers. One run costs money. API_BUILD=1 is required, so it
// can not run by accident.
//
// Usage:
//
//	API_BUILD=1 API_BUILD_EMAIL=... API_BUILD_PASSWORD=... \
//	  go run ./cmd/api-build -prompt "Build me a lifegain Commander deck."
package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	"connectrpc.com/connect"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
	"github.com/nkramber/decktome/go/gen/mtg/v1/mtgv1connect"
	"github.com/nkramber/decktome/go/internal/gatekit"
)

// defaults name the deployed project. Both values already sit in
// cloudbuild/web.yaml, and the built web app serves them to every
// reader, so neither is a secret. The email and the password are not
// here, and they never join a file of this repository (D-639).
const (
	defaultBaseURL = "https://mtg-api-qk2ackpb3q-uc.a.run.app"
	defaultAPIKey  = "AIzaSyCe_Ef4Heqzeguk-IJ5mCNLEH19UJXL9Do"
	defaultCSV     = "internal/collections/testdata/manabox_collection.csv"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, "api-build: "+err.Error())
		os.Exit(1)
	}
}

// options hold every flag of one run.
type options struct {
	base     string
	csv      string
	name     string
	prompt   string
	answers  string
	poolRule string
	maxTurns int
	timeout  time.Duration
	ready    time.Duration
	cleanup  bool
}

func run() error {
	var o options
	flagSet(&o)
	if err := gatekit.SpendGuard("API_BUILD"); err != nil {
		return err
	}
	return build(context.Background(), o, credentials{
		endpoint: identityToolkit,
		apiKey:   envOr("FIREBASE_API_KEY", defaultAPIKey),
		email:    os.Getenv("API_BUILD_EMAIL"),
		password: os.Getenv("API_BUILD_PASSWORD"),
	})
}

// credentials name the sign-in. The endpoint is a field so a test can
// serve it, and no test ever reaches the deployed project.
type credentials struct {
	endpoint string
	apiKey   string
	email    string
	password string
}

// build runs the whole flow: the sign-in, the import, the chat, the
// read-back, and the cleanup. It is the body of one run.
func build(ctx context.Context, o options, cred credentials) error {
	p, err := parsePlan(o.answers)
	if err != nil {
		return err
	}
	rule, err := parsePoolRule(o.poolRule)
	if err != nil {
		return err
	}
	csv, err := os.ReadFile(o.csv)
	if err != nil {
		return fmt.Errorf("the collection file: %w", err)
	}

	acct, err := signIn(ctx, newHTTPClient("", 30*time.Second), cred.endpoint,
		cred.apiKey, cred.email, cred.password)
	if err != nil {
		return err
	}
	// The email never reaches the output. The uid names the account, and
	// the repository is public (D-639).
	fmt.Printf("signed in    uid %s at %s\n", acct.LocalID, o.base)

	hc := newHTTPClient(acct.IDToken, o.timeout)
	colls := mtgv1connect.NewCollectionServiceClient(hc, o.base)
	agent := mtgv1connect.NewAgentServiceClient(hc, o.base)
	decks := mtgv1connect.NewDeckServiceClient(hc, o.base)
	health := mtgv1connect.NewHealthServiceClient(hc, o.base)

	// A cold instance answers an import with "card database not loaded
	// yet". The import runs after the snapshot is in memory.
	if err := waitReady(ctx, health, o.ready, func(f string, a ...any) { fmt.Printf(f, a...) }); err != nil {
		return err
	}

	start := time.Now()
	imported, err := colls.ImportCollection(ctx, connect.NewRequest(&mtgv1.ImportCollectionRequest{
		Name:    o.name,
		Source:  mtgv1.ImportSource_IMPORT_SOURCE_MANABOX_CSV,
		Content: csv,
	}))
	if err != nil {
		return explain("the import", err)
	}
	coll := imported.Msg.GetCollection()
	rep := imported.Msg.GetReport()
	// The summary is filled by the import, so the count is right whether
	// the answer carries the entries or leaves them out (D-392).
	fmt.Printf("collection   %s, %d rows, %d cards, %d rows resolved, %d unresolved\n",
		coll.GetId(), coll.GetSummary().GetRowCount(), coll.GetCardCount(),
		rep.GetResolvedCount(), len(rep.GetUnresolved()))

	built, err := chat(ctx, agent, o, coll.GetId(), rule, p)
	if err != nil {
		return err
	}
	fmt.Printf("built        deck %s in session %s after %d turns, %.0fs\n",
		built.deckID, built.sessionID, built.turns, time.Since(start).Seconds())

	// The read-back is the point of D-778. A deck in the stream proves
	// the build. Only a deck that comes back out of storage proves that
	// the reader would find it.
	stored, err := decks.GetDeck(ctx, connect.NewRequest(&mtgv1.GetDeckRequest{DeckId: built.deckID}))
	if err != nil {
		return explain("the stored deck", err)
	}
	listed, err := decks.ListDecks(ctx, connect.NewRequest(&mtgv1.ListDecksRequest{SessionId: built.sessionID}))
	if err != nil {
		return explain("the deck list", err)
	}
	session, err := agent.GetSession(ctx, connect.NewRequest(&mtgv1.GetSessionRequest{SessionId: built.sessionID}))
	if err != nil {
		return explain("the stored session", err)
	}
	report(stored.Msg.GetDeck(), len(listed.Msg.GetDecks()), session.Msg.GetSession(), p)

	if o.cleanup {
		if err := remove(ctx, agent, decks, colls, built, coll.GetId()); err != nil {
			return err
		}
		fmt.Println("cleanup      the deck, the session, and the collection are deleted")
	} else {
		// D-780 keeps what the run made, so a later read finds the deck.
		fmt.Printf("kept         deck %s, session %s, collection %s (D-780). Pass -cleanup to delete them\n",
			built.deckID, built.sessionID, coll.GetId())
	}
	return nil
}

// outcome is what one chat produced.
type outcome struct {
	sessionID string
	deckID    string
	turns     int
}

// chat runs the conversation from the first prompt to a finished deck.
// It answers every question the agent asks, so no reader is needed.
func chat(ctx context.Context, agent mtgv1connect.AgentServiceClient, o options,
	collectionID string, rule mtgv1.PoolRule, p plan) (*outcome, error) {
	out := &outcome{}
	req := &mtgv1.ChatRequest{
		CollectionId: collectionID,
		Message:      o.prompt,
		PoolRule:     rule,
	}
	for out.turns < o.maxTurns {
		out.turns++
		fmt.Printf("\n--- turn %d\n", out.turns)
		asked, deck, err := turn(ctx, agent, req, out)
		if err != nil {
			return nil, err
		}
		if deck != nil {
			out.deckID = deck.GetId()
			if out.deckID == "" {
				return nil, errors.New("the deck reached the client with no id, so no read can find it")
			}
			return out, nil
		}
		if len(asked) == 0 {
			return nil, fmt.Errorf("turn %d asked nothing and built nothing, so the run has no next move", out.turns)
		}
		answers, err := answerAll(asked, p)
		if err != nil {
			return nil, err
		}
		for i, a := range answers {
			fmt.Printf("  A [%s] %s\n", asked[i].GetSlot(), describe(asked[i], a))
		}
		// A turn of answers alone is a valid turn, and the service takes
		// it with an empty message.
		req = &mtgv1.ChatRequest{SessionId: out.sessionID, Answers: answers}
	}
	return nil, fmt.Errorf("no deck after %d turns: raise -max-turns, or answer more slots with -answers", o.maxTurns)
}

// turn sends one request and reads the whole stream it answers with.
func turn(ctx context.Context, agent mtgv1connect.AgentServiceClient,
	req *mtgv1.ChatRequest, out *outcome) ([]*mtgv1.Question, *mtgv1.Deck, error) {
	stream, err := agent.Chat(ctx, connect.NewRequest(req))
	if err != nil {
		return nil, nil, explain("the chat", err)
	}
	defer func() { _ = stream.Close() }()

	var asked []*mtgv1.Question
	var deck *mtgv1.Deck
	for stream.Receive() {
		switch e := stream.Msg().GetEvent().(type) {
		case *mtgv1.ChatResponse_SessionStarted:
			out.sessionID = e.SessionStarted
			fmt.Printf("  session %s\n", e.SessionStarted)
		case *mtgv1.ChatResponse_Question:
			asked = append(asked, e.Question)
			fmt.Printf("  Q [%s] %s\n", e.Question.GetSlot(), e.Question.GetText())
		case *mtgv1.ChatResponse_Status:
			fmt.Printf("  status: %s\n", e.Status)
		case *mtgv1.ChatResponse_Deck:
			deck = e.Deck
		case *mtgv1.ChatResponse_Failure:
			return nil, nil, fmt.Errorf("the agent failed with %s: %s", e.Failure.GetCode(), e.Failure.GetMessage())
		}
	}
	if err := stream.Err(); err != nil {
		return nil, nil, explain("the chat stream", err)
	}
	return asked, deck, nil
}

// report prints what the deployed API stored, and it is the evidence of
// the run.
func report(d *mtgv1.Deck, listed int, s *mtgv1.Session, p plan) {
	fmt.Printf("\n=== the deployed API built and stored a deck ===\n")
	fmt.Printf("deck         %s, %q\n", d.GetId(), d.GetName())
	fmt.Printf("format       %v, %d commanders, %d main cards\n",
		d.GetFormat().GetId(), len(d.GetCommanderOracleIds()), countCards(d))
	if b := d.GetPower().GetBracket(); b != 0 {
		fmt.Printf("bracket      %d\n", b)
	}
	fmt.Printf("listing      %d deck(s) under the session\n", listed)
	fmt.Printf("session      %d turn(s) stored, %d slot state(s)\n",
		len(s.GetTurns()), len(s.GetSlots().GetSlotStates()))
	if len(p) > 0 {
		fmt.Printf("plan         %s\n", strings.Join(p.slots(), ", "))
	}
	for _, f := range d.GetValidation().GetFindings() {
		fmt.Printf("  [%s] %s: %s\n",
			strings.TrimPrefix(f.GetSeverity().String(), "SEVERITY_"), f.GetCode(), f.GetMessage())
	}
}

// countCards adds the copies of every main-deck card.
func countCards(d *mtgv1.Deck) int {
	n := 0
	for _, c := range d.GetCards() {
		n += int(c.GetCount())
	}
	return n
}

// remove deletes what the run made, newest first, so no orphan is left
// when one call fails.
func remove(ctx context.Context, agent mtgv1connect.AgentServiceClient,
	decks mtgv1connect.DeckServiceClient, colls mtgv1connect.CollectionServiceClient,
	built *outcome, collectionID string) error {
	if _, err := decks.DeleteDeck(ctx, connect.NewRequest(&mtgv1.DeleteDeckRequest{DeckId: built.deckID})); err != nil {
		return explain("the deck delete", err)
	}
	if _, err := agent.DeleteSession(ctx, connect.NewRequest(&mtgv1.DeleteSessionRequest{SessionId: built.sessionID})); err != nil {
		return explain("the session delete", err)
	}
	if _, err := colls.DeleteCollection(ctx, connect.NewRequest(&mtgv1.DeleteCollectionRequest{CollectionId: collectionID})); err != nil {
		return explain("the collection delete", err)
	}
	return nil
}

// explain names the call that failed, and it turns the invite refusal
// into the command that fixes it (D-314).
func explain(what string, err error) error {
	var ce *connect.Error
	if errors.As(err, &ce) && ce.Meta().Get("Deck-Tome-Refusal") == "not-invited" {
		return fmt.Errorf("%s: the account is not on the invite list. Run: make allow EMAIL=<the check account> PROJECT_ID=decktome-prod", what)
	}
	return fmt.Errorf("%s: %w", what, err)
}

// parsePoolRule reads the -pool flag.
func parsePoolRule(s string) (mtgv1.PoolRule, error) {
	switch strings.TrimSpace(strings.ToLower(s)) {
	case "owned-first":
		return mtgv1.PoolRule_POOL_RULE_OWNED_FIRST, nil
	case "owned-only":
		return mtgv1.PoolRule_POOL_RULE_OWNED_ONLY, nil
	case "any":
		return mtgv1.PoolRule_POOL_RULE_ANY_CARD, nil
	case "", "ask":
		return mtgv1.PoolRule_POOL_RULE_UNSPECIFIED, nil
	default:
		return 0, fmt.Errorf("-pool %q is not owned-first, owned-only, any, or ask", s)
	}
}

// envOr reads a variable, and it falls back to the committed default.
func envOr(name, fallback string) string {
	if v := os.Getenv(name); v != "" {
		return v
	}
	return fallback
}

// flagSet declares every flag and parses the command line.
func flagSet(o *options) {
	flag.StringVar(&o.base, "base", envOr("API_BASE_URL", defaultBaseURL), "the API origin. decktome.com serves the web app alone, and it rewrites nothing to the API")
	flag.StringVar(&o.csv, "collection", defaultCSV, "the ManaBox CSV to import")
	flag.StringVar(&o.name, "name", "api-build check", "the name of the imported collection")
	flag.StringVar(&o.prompt, "prompt", "Build me a lifegain Commander deck from the cards I own.", "the first message")
	flag.StringVar(&o.answers, "answers", "", "the prepared answers: slot=text;slot=#2;slot=decline. A slot with no entry takes the first option of a closed question, and declines the rest")
	flag.StringVar(&o.poolRule, "pool", "owned-first", "the card pool: owned-first, owned-only, any, or ask")
	flag.IntVar(&o.maxTurns, "max-turns", 8, "the turns the run takes before it gives up")
	flag.DurationVar(&o.timeout, "timeout", 15*time.Minute, "the limit of one streamed turn")
	flag.DurationVar(&o.ready, "ready", 3*time.Minute, "how long to wait for the API to load the card snapshot")
	flag.BoolVar(&o.cleanup, "cleanup", false, "delete the deck, the session, and the collection at the end (D-780)")
	flag.Parse()
}
