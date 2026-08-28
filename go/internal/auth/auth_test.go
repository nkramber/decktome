package auth

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"connectrpc.com/connect"

	mtgv1 "github.com/nkramber/mtg-deck-builder/go/gen/mtg/v1"
	"github.com/nkramber/mtg-deck-builder/go/gen/mtg/v1/mtgv1connect"
)

// fakeVerifier accepts one token and names one user.
type fakeVerifier struct {
	token string
	uid   string
}

func (f fakeVerifier) Verify(_ context.Context, idToken string) (string, error) {
	if idToken == f.token {
		return f.uid, nil
	}
	return "", errors.New("unknown token")
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
	ic := connect.WithInterceptors(Interceptor(fakeVerifier{token: "good", uid: "u-42"}, opts...))
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
