package main

import (
	"bytes"
	"context"
	"errors"
	"strings"
	"testing"
	"time"

	"github.com/nkramber/decktome/go/internal/users"
)

type fakeRecords struct {
	rec     users.Record
	markErr error
	steps   *[]string
}

func (f *fakeRecords) Get(context.Context, string) (users.Record, error) { return f.rec, nil }

func (f *fakeRecords) Deactivate(_ context.Context, _ string, at time.Time) (time.Time, error) {
	*f.steps = append(*f.steps, "mark")
	return at, f.markErr
}

type fakeAuth struct {
	err   error
	steps *[]string
}

func (f fakeAuth) Delete(context.Context, string) error {
	*f.steps = append(*f.steps, "auth")
	return f.err
}

// TestDeactivate is REV-036 (D-941): a dry run changes nothing, and a
// close marks the record before the Auth delete, and deletes no record.
func TestDeactivate(t *testing.T) {
	now := time.Date(2026, 9, 25, 12, 0, 0, 0, time.UTC)
	tests := []struct {
		name      string
		uid       string
		confirm   bool
		markErr   error
		authErr   error
		wantSteps string
		wantErr   string
		wantOut   string
	}{
		{name: "a dry run changes nothing", uid: "u1", wantSteps: "", wantOut: "dry run"},
		{name: "a close marks, then deletes the Auth user", uid: "u1", confirm: true, wantSteps: "mark auth", wantOut: "deleted the Firebase Auth user"},
		{name: "an Auth user that is gone", uid: "u1", confirm: true, authErr: ErrNoAuthUser, wantSteps: "mark auth", wantOut: "gone already"},
		{name: "a failed Auth delete", uid: "u1", confirm: true, authErr: errors.New("auth down"), wantSteps: "mark auth", wantErr: "run the command again"},
		{name: "a failed mark stops before Auth", uid: "u1", confirm: true, markErr: errors.New("store down"), wantSteps: "mark", wantErr: "nothing changed"},
		{name: "no uid", uid: "", confirm: true, wantErr: "-uid is required"},
		{name: "a path is no uid", uid: "u1/decks", confirm: true, wantErr: "no valid user id"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var steps []string
			recs := &fakeRecords{rec: users.Record{Email: "ann@example.com", CreatedAt: now}, markErr: tt.markErr, steps: &steps}
			var out bytes.Buffer
			err := deactivate(context.Background(), recs, fakeAuth{err: tt.authErr, steps: &steps}, tt.uid, tt.confirm, now, &out)
			if tt.wantErr != "" {
				if err == nil || !strings.Contains(err.Error(), tt.wantErr) {
					t.Fatalf("err = %v, want %q", err, tt.wantErr)
				}
			} else if err != nil {
				t.Fatal(err)
			}
			if got := strings.Join(steps, " "); got != tt.wantSteps {
				t.Errorf("steps = %q, want %q", got, tt.wantSteps)
			}
			if tt.wantOut != "" && !strings.Contains(out.String(), tt.wantOut) {
				t.Errorf("output lacks %q:\n%s", tt.wantOut, out.String())
			}
		})
	}
}
