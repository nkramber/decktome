package auth

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"connectrpc.com/connect"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
	"github.com/nkramber/decktome/go/gen/mtg/v1/mtgv1connect"
)

// fakeVerifier accepts one token and names one user.
type fakeVerifier struct {
	token string
	uid   string
	email string
}

func (f fakeVerifier) Verify(_ context.Context, idToken string) (Identity, error) {
	if idToken == f.token {
		return Identity{UID: f.uid, Email: f.email}, nil
	}
	return Identity{}, errors.New("unknown token")
}

// fakeList allows the emails it holds. A nil error map makes every
// lookup succeed.
type fakeList struct {
	emails map[string]bool
	err    error
}

func (f fakeList) Allowed(_ context.Context, email string) (bool, error) {
	if f.err != nil {
		return false, f.err
	}
	return f.emails[email], nil
}

// echo answers the unary Check with the user id in the version field,
// and the streaming Chat with the user id as the session id.
type echo struct {
	mtgv1connect.UnimplementedHealthServiceHandler
	mtgv1connect.UnimplementedAgentServiceHandler
}

func (echo) Check(ctx context.Context, _ *connect.Request[mtgv1.CheckRequest]) (*connect.Response[mtgv1.CheckResponse], error) {
	return connect.NewResponse(&mtgv1.CheckResponse{Version: UserID(ctx)}), nil
}

func (echo) Chat(ctx context.Context, _ *connect.Request[mtgv1.ChatRequest], stream *connect.ServerStream[mtgv1.ChatResponse]) error {
	return stream.Send(&mtgv1.ChatResponse{Event: &mtgv1.ChatResponse_SessionStarted{SessionStarted: UserID(ctx)}})
}

func newServer(t *testing.T, opts ...Option) (mtgv1connect.HealthServiceClient, mtgv1connect.AgentServiceClient) {
	t.Helper()
	ic := connect.WithInterceptors(Interceptor(fakeVerifier{token: "good", uid: "u-42", email: "ann@example.com"}, opts...))
	mux := http.NewServeMux()
	mux.Handle(mtgv1connect.NewHealthServiceHandler(echo{}, ic))
	mux.Handle(mtgv1connect.NewAgentServiceHandler(echo{}, ic))
	srv := httptest.NewServer(mux)
	t.Cleanup(srv.Close)
	return mtgv1connect.NewHealthServiceClient(srv.Client(), srv.URL), mtgv1connect.NewAgentServiceClient(srv.Client(), srv.URL)
}

func TestInterceptor(t *testing.T) {
	tests := []struct {
		name     string
		header   string
		fallback string
		wantUID  string
		wantCode connect.Code
	}{
		{name: "valid token", header: "Bearer good", wantUID: "u-42"},
		{name: "valid token, case of scheme", header: "bearer good", wantUID: "u-42"},
		{name: "valid token with a fallback set", header: "Bearer good", fallback: "local-dev", wantUID: "u-42"},
		{name: "bad token", header: "Bearer nope", wantCode: connect.CodeUnauthenticated},
		{name: "bad token with a fallback set is still refused", header: "Bearer nope", fallback: "local-dev", wantCode: connect.CodeUnauthenticated},
		{name: "empty bearer", header: "Bearer ", fallback: "local-dev", wantCode: connect.CodeUnauthenticated},
		{name: "bare bearer with a fallback is refused", header: "Bearer", fallback: "local-dev", wantCode: connect.CodeUnauthenticated},
		{name: "bare bearer without a fallback", header: "Bearer", wantCode: connect.CodeUnauthenticated},
		{name: "tab after the scheme is refused", header: "Bearer\tgood", fallback: "local-dev", wantCode: connect.CodeUnauthenticated},
		{name: "scheme glued to a word is another scheme", header: "Bearerx good", fallback: "local-dev", wantUID: "local-dev"},
		{name: "no token, no fallback", wantCode: connect.CodeUnauthenticated},
		{name: "no token, fallback", fallback: "local-dev", wantUID: "local-dev"},
		{name: "other scheme, fallback", header: "Basic abc", fallback: "local-dev", wantUID: "local-dev"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var opts []Option
			if tt.fallback != "" {
				opts = append(opts, WithFallback(tt.fallback))
			}
			health, agent := newServer(t, opts...)

			req := connect.NewRequest(&mtgv1.CheckRequest{})
			if tt.header != "" {
				req.Header().Set("Authorization", tt.header)
			}
			res, err := health.Check(context.Background(), req)
			if codeOf(err) != tt.wantCode {
				t.Fatalf("unary code = %v, want %v: %v", connect.CodeOf(err), tt.wantCode, err)
			}
			if err == nil && res.Msg.GetVersion() != tt.wantUID {
				t.Errorf("unary uid = %q, want %q", res.Msg.GetVersion(), tt.wantUID)
			}

			sreq := connect.NewRequest(&mtgv1.ChatRequest{Message: "hi"})
			if tt.header != "" {
				sreq.Header().Set("Authorization", tt.header)
			}
			stream, err := agent.Chat(context.Background(), sreq)
			if err != nil {
				t.Fatalf("chat: %v", err)
			}
			var got string
			for stream.Receive() {
				got = stream.Msg().GetSessionStarted()
			}
			err = stream.Err()
			_ = stream.Close()
			if codeOf(err) != tt.wantCode {
				t.Fatalf("stream code = %v, want %v: %v", connect.CodeOf(err), tt.wantCode, err)
			}
			if err == nil && got != tt.wantUID {
				t.Errorf("stream uid = %q, want %q", got, tt.wantUID)
			}
		})
	}
}

// codeOf is connect.CodeOf with a zero code for a nil error.
func codeOf(err error) connect.Code {
	if err == nil {
		return 0
	}
	return connect.CodeOf(err)
}

func TestRejectAll(t *testing.T) {
	if _, err := RejectAll().Verify(context.Background(), "anything"); !errors.Is(err, ErrNotConfigured) {
		t.Errorf("err = %v", err)
	}
}

func TestUserID(t *testing.T) {
	if got := UserID(context.Background()); got != "" {
		t.Errorf("empty context gave %q", got)
	}
	if got := UserID(WithUserID(context.Background(), "u-1")); got != "u-1" {
		t.Errorf("got %q", got)
	}
}

func TestCORS(t *testing.T) {
	tests := []struct {
		name      string
		allowed   string
		origin    string
		method    string
		wantAllow string
		wantCode  int
	}{
		{name: "same origin, no header", allowed: "", origin: "", method: http.MethodPost, wantCode: http.StatusOK},
		{name: "empty list refuses every origin", allowed: "", origin: "https://app.example", method: http.MethodPost, wantCode: http.StatusOK},
		{name: "listed origin", allowed: "https://app.example, https://dev.example", origin: "https://app.example", method: http.MethodPost, wantAllow: "https://app.example", wantCode: http.StatusOK},
		{name: "listed origin preflight", allowed: "https://app.example", origin: "https://app.example", method: http.MethodOptions, wantAllow: "https://app.example", wantCode: http.StatusNoContent},
		{name: "unlisted origin preflight reaches the handler", allowed: "https://app.example", origin: "https://evil.example", method: http.MethodOptions, wantCode: http.StatusOK},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := CORS(ParseOrigins(tt.allowed), http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
				w.WriteHeader(http.StatusOK)
			}))
			req := httptest.NewRequest(tt.method, "/x", nil)
			if tt.origin != "" {
				req.Header.Set("Origin", tt.origin)
			}
			rec := httptest.NewRecorder()
			h.ServeHTTP(rec, req)
			if rec.Code != tt.wantCode {
				t.Errorf("code = %d, want %d", rec.Code, tt.wantCode)
			}
			if got := rec.Header().Get("Access-Control-Allow-Origin"); got != tt.wantAllow {
				t.Errorf("allow origin = %q, want %q", got, tt.wantAllow)
			}
		})
	}
}

func TestBearer(t *testing.T) {
	for _, tc := range []struct {
		in      string
		token   string
		present bool
	}{
		{"Bearer abc", "abc", true},
		{"bearer abc", "abc", true},
		{"Bearer", "", true},
		{"Bearer ", "", true},
		{"Bearer\tabc", "", true},
		{"Basic abc", "", false},
		{"", "", false},
		{"Bearerx abc", "", false},
	} {
		token, present := bearer(tc.in)
		if token != tc.token || present != tc.present {
			t.Errorf("bearer(%q) = %q, %v, want %q, %v", tc.in, token, present, tc.token, tc.present)
		}
	}
}

// TestPublicProcedurePassesWithNoToken is D-315: a procedure on the
// public list needs no bearer token and carries no user, and every other
// procedure of the service keeps the check.
func TestPublicProcedurePassesWithNoToken(t *testing.T) {
	in := Interceptor(RejectAll(), WithPublic("/mtg.v1.DeckService/GetSharedDeck"))
	var seen string
	next := in.WrapUnary(func(ctx context.Context, _ connect.AnyRequest) (connect.AnyResponse, error) {
		seen = UserID(ctx)
		return nil, nil
	})
	call := func(procedure string) error {
		req := connect.NewRequest(&struct{}{})
		_, err := next(context.Background(), &specRequest{Request: req, spec: connect.Spec{Procedure: procedure}})
		return err
	}
	if err := call("/mtg.v1.DeckService/GetSharedDeck"); err != nil {
		t.Fatalf("public procedure: %v", err)
	}
	if seen != "" {
		t.Errorf("a public call carries a user %q", seen)
	}
	if err := call("/mtg.v1.DeckService/GetDeck"); connect.CodeOf(err) != connect.CodeUnauthenticated {
		t.Errorf("another procedure: %v, want Unauthenticated", err)
	}
}

// specRequest gives a request the procedure a server sees.
type specRequest struct {
	*connect.Request[struct{}]
	spec connect.Spec
}

func (r *specRequest) Spec() connect.Spec { return r.spec }

// TestAllowlistRefusesAnEmailOffTheList is D-314: a verified user whose
// email is not on the list gets PermissionDenied with one sentence, a
// user on the list gets in, a list that can not be read answers
// Unavailable, and the fallback user of local mode never meets the list.
func TestAllowlistRefusesAnEmailOffTheList(t *testing.T) {
	call := func(t *testing.T, header string, opts ...Option) (string, error) {
		t.Helper()
		health, _ := newServer(t, opts...)
		req := connect.NewRequest(&mtgv1.CheckRequest{})
		if header != "" {
			req.Header().Set("Authorization", header)
		}
		res, err := health.Check(context.Background(), req)
		if err != nil {
			return "", err
		}
		return res.Msg.GetVersion(), nil
	}
	if uid, err := call(t, "Bearer good", WithAllowlist(fakeList{emails: map[string]bool{"ann@example.com": true}})); err != nil || uid != "u-42" {
		t.Errorf("a listed email must get in: %q %v", uid, err)
	}
	_, err := call(t, "Bearer good", WithAllowlist(fakeList{emails: map[string]bool{"bob@example.com": true}}))
	if codeOf(err) != connect.CodePermissionDenied || !strings.Contains(err.Error(), "invited users alone") {
		t.Errorf("an email off the list: %v", err)
	}
	if _, err := call(t, "Bearer good", WithAllowlist(fakeList{err: errors.New("firestore down")})); codeOf(err) != connect.CodeUnavailable {
		t.Errorf("a list that can not be read: %v", err)
	}
	if uid, err := call(t, "", WithFallback("local-dev"), WithAllowlist(fakeList{})); err != nil || uid != "local-dev" {
		t.Errorf("the fallback user never meets the list: %q %v", uid, err)
	}
}
