package rpcerr

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"connectrpc.com/connect"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
	"github.com/nkramber/decktome/go/gen/mtg/v1/mtgv1connect"
)

type stubHealth struct {
	mtgv1connect.UnimplementedHealthServiceHandler
	err error
}

func (s stubHealth) Check(context.Context, *connect.Request[mtgv1.CheckRequest]) (*connect.Response[mtgv1.CheckResponse], error) {
	if s.err != nil {
		return nil, s.err
	}
	return connect.NewResponse(&mtgv1.CheckResponse{}), nil
}

// TestInterceptorMasksInternalAlone: an Internal error reaches the caller
// as Message, and its text reaches the log alone. Every other code keeps
// its message.
func TestInterceptorMasksInternalAlone(t *testing.T) {
	const detail = "decode: proto: cannot parse invalid wire-format data"
	tests := []struct {
		name     string
		err      error
		wantCode connect.Code
		wantMsg  string
		wantLog  bool
	}{
		{"internal", connect.NewError(connect.CodeInternal, errors.New(detail)), connect.CodeInternal, Message, true},
		{"wrapped internal", fmt.Errorf("read: %w", connect.NewError(connect.CodeInternal, errors.New(detail))), connect.CodeInternal, Message, true},
		{"not found", connect.NewError(connect.CodeNotFound, errors.New("no such deck")), connect.CodeNotFound, "no such deck", false},
		{"invalid argument", connect.NewError(connect.CodeInvalidArgument, errors.New("a token is required")), connect.CodeInvalidArgument, "a token is required", false},
		{"no error", nil, 0, "", false},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			var logged bytes.Buffer
			mux := http.NewServeMux()
			mux.Handle(mtgv1connect.NewHealthServiceHandler(stubHealth{err: tc.err},
				connect.WithInterceptors(Interceptor(slog.New(slog.NewTextHandler(&logged, nil))))))
			srv := httptest.NewServer(mux)
			defer srv.Close()
			client := mtgv1connect.NewHealthServiceClient(srv.Client(), srv.URL)

			_, err := client.Check(context.Background(), connect.NewRequest(&mtgv1.CheckRequest{}))
			if tc.err == nil {
				if err != nil {
					t.Fatalf("err = %v, want none", err)
				}
				return
			}
			var ce *connect.Error
			if !errors.As(err, &ce) {
				t.Fatalf("err = %v, want a Connect error", err)
			}
			if ce.Code() != tc.wantCode || ce.Message() != tc.wantMsg {
				t.Fatalf("got %v %q, want %v %q", ce.Code(), ce.Message(), tc.wantCode, tc.wantMsg)
			}
			hasLog := strings.Contains(logged.String(), detail) && strings.Contains(logged.String(), mtgv1connect.HealthServiceCheckProcedure)
			if hasLog != tc.wantLog {
				t.Fatalf("log %q, want the detail and procedure logged: %v", logged.String(), tc.wantLog)
			}
		})
	}
}

type fakeConn struct {
	connect.StreamingHandlerConn
	spec connect.Spec
}

func (c fakeConn) Spec() connect.Spec { return c.spec }

// TestInterceptorMasksAStreamError: a streaming handler that ends with an
// Internal error ends with Message, and a NotFound passes unchanged.
func TestInterceptorMasksAStreamError(t *testing.T) {
	const detail = "firestore: deadline"
	tests := []struct {
		name    string
		err     error
		wantMsg string
	}{
		{"internal", connect.NewError(connect.CodeInternal, errors.New(detail)), Message},
		{"not found", connect.NewError(connect.CodeNotFound, errors.New("no such chat")), "no such chat"},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			var logged bytes.Buffer
			wrap := Interceptor(slog.New(slog.NewTextHandler(&logged, nil))).WrapStreamingHandler(
				func(context.Context, connect.StreamingHandlerConn) error { return tc.err })
			err := wrap(context.Background(), fakeConn{spec: connect.Spec{Procedure: "/mtg.v1.AgentService/Chat"}})
			var ce *connect.Error
			if !errors.As(err, &ce) || ce.Code() != connect.CodeOf(tc.err) || ce.Message() != tc.wantMsg {
				t.Fatalf("err = %v, want %v %q", err, connect.CodeOf(tc.err), tc.wantMsg)
			}
			if strings.Contains(logged.String(), detail) != (tc.wantMsg == Message) {
				t.Fatalf("log %q", logged.String())
			}
		})
	}
}
