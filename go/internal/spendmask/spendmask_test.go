package spendmask

import (
	"context"
	"net/http"
	"testing"

	"connectrpc.com/connect"
	"google.golang.org/protobuf/proto"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
	"github.com/nkramber/decktome/go/internal/auth"
)

func spend() *mtgv1.Usage {
	return &mtgv1.Usage{Calls: 6, InputTokens: 7957, OutputTokens: 484, CostUsd: 0.0009, Priced: true}
}

// deckEvent is a chat event whose deck holds the spend of its build.
func deckEvent() *mtgv1.ChatResponse {
	return &mtgv1.ChatResponse{Event: &mtgv1.ChatResponse_Deck{
		Deck: &mtgv1.Deck{Name: "Tokens", Build: &mtgv1.BuildMetrics{Shortlist: 295, Usage: spend()}},
	}}
}

// usageEvent is the chat event that holds the session total alone.
func usageEvent() *mtgv1.ChatResponse {
	return &mtgv1.ChatResponse{Event: &mtgv1.ChatResponse_Usage{Usage: spend()}}
}

// D-1148: a copy loses each spend message, and the message of the
// service keeps its spend.
func TestMaskedClearsEachDepthOfACopy(t *testing.T) {
	m := deckEvent()
	got := Masked(m).(*mtgv1.ChatResponse)
	if got.GetDeck().GetBuild().GetUsage() != nil {
		t.Fatalf("the copy kept a spend: %v", got)
	}
	if got.GetDeck().GetBuild().GetShortlist() != 295 || got.GetDeck().GetName() != "Tokens" {
		t.Fatalf("the copy lost a field that is not spend: %v", got)
	}
	if !proto.Equal(m, deckEvent()) {
		t.Fatalf("the message of the service changed: %v", m)
	}
}

func TestHasReadsAListAndAnEmptyMessage(t *testing.T) {
	list := &mtgv1.ListSessionsResponse{Sessions: []*mtgv1.SessionSummary{{Name: "a"}, {Name: "b", Usage: spend()}}}
	if !Has(list) {
		t.Fatal("Has missed the spend of the second summary")
	}
	masked := Masked(list).(*mtgv1.ListSessionsResponse)
	if Has(masked) || masked.GetSessions()[1].GetName() != "b" {
		t.Fatalf("the masked list: %v", masked)
	}
	if Has(&mtgv1.ChatResponse{Event: &mtgv1.ChatResponse_Status{Status: "x"}}) {
		t.Fatal("Has found a spend in a message with none")
	}
}

func unary(ctx context.Context, t *testing.T) *mtgv1.Session {
	t.Helper()
	sent := &mtgv1.Session{Id: "s1", Usage: spend()}
	next := func(context.Context, connect.AnyRequest) (connect.AnyResponse, error) {
		res := connect.NewResponse(sent)
		res.Header().Set("X-Test", "kept")
		return res, nil
	}
	res, err := Interceptor().WrapUnary(next)(ctx, connect.NewRequest(&mtgv1.GetSessionRequest{}))
	if err != nil {
		t.Fatal(err)
	}
	if res.Header().Get("X-Test") != "kept" {
		t.Fatal("the mask dropped a header")
	}
	if sent.GetUsage() == nil {
		t.Fatal("the mask changed the message of the service")
	}
	return res.Any().(*mtgv1.Session)
}

// D-1148: a caller without the admin claim gets no spend, and the admin
// gets it.
func TestUnaryMasksAllButTheAdmin(t *testing.T) {
	if got := unary(context.Background(), t); got.GetUsage() != nil || got.GetId() != "s1" {
		t.Fatalf("a reader got %v", got)
	}
	if got := unary(auth.WithAdmin(context.Background(), true), t); got.GetUsage().GetCalls() != 6 {
		t.Fatalf("the admin got %v", got)
	}
}

type fakeConn struct {
	connect.StreamingHandlerConn
	sent []any
}

func (c *fakeConn) Send(m any) error {
	c.sent = append(c.sent, m)
	return nil
}

func (c *fakeConn) ResponseHeader() http.Header { return http.Header{} }

// stream sends a status, a deck, and a session total through the mask.
func stream(ctx context.Context, t *testing.T) (*mtgv1.ChatResponse, []any) {
	t.Helper()
	sent := deckEvent()
	conn := &fakeConn{}
	next := func(_ context.Context, c connect.StreamingHandlerConn) error {
		for _, m := range []proto.Message{&mtgv1.ChatResponse{Event: &mtgv1.ChatResponse_Status{Status: "x"}}, sent, usageEvent()} {
			if err := c.Send(m); err != nil {
				return err
			}
		}
		return nil
	}
	if err := Interceptor().WrapStreamingHandler(next)(ctx, conn); err != nil {
		t.Fatal(err)
	}
	return sent, conn.sent
}

// D-1148: a reader gets the status and the deck with no spend, and no
// event of the session total. The admin gets all three as sent.
func TestStreamMasksAllButTheAdmin(t *testing.T) {
	sent, got := stream(context.Background(), t)
	if len(got) != 2 {
		t.Fatalf("a reader got %d events, want 2: %v", len(got), got)
	}
	for _, m := range got {
		if Has(m.(proto.Message)) {
			t.Fatalf("a reader got a spend: %v", m)
		}
	}
	if !proto.Equal(sent, deckEvent()) {
		t.Fatal("the mask changed the message of the service")
	}
	_, got = stream(auth.WithAdmin(context.Background(), true), t)
	if len(got) != 3 || got[2].(*mtgv1.ChatResponse).GetUsage().GetCalls() != 6 {
		t.Fatalf("the admin got %v", got)
	}
}
