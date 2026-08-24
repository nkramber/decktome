package health

import (
	"context"
	"testing"

	"connectrpc.com/connect"

	mtgv1 "github.com/nkramber/mtg-deck-builder/go/gen/mtg/v1"
)

func TestCheck(t *testing.T) {
	tests := []struct {
		name    string
		version string
	}{
		{name: "dev build", version: "dev"},
		{name: "tagged build", version: "v0.1.0"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := New(tt.version)
			res, err := s.Check(context.Background(), connect.NewRequest(&mtgv1.CheckRequest{}))
			if err != nil {
				t.Fatalf("Check returned error: %v", err)
			}
			if got := res.Msg.GetStatus(); got != "ok" {
				t.Errorf("status = %q, want %q", got, "ok")
			}
			if got := res.Msg.GetVersion(); got != tt.version {
				t.Errorf("version = %q, want %q", got, tt.version)
			}
		})
	}
}
