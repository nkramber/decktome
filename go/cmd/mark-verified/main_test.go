package main

import (
	"bytes"
	"context"
	"slices"
	"strings"
	"testing"
	"time"

	"github.com/nkramber/decktome/go/internal/allowlist"
)

type fakeAccounts struct {
	all    []account
	marked []string
}

func (f *fakeAccounts) List(context.Context) ([]account, error) { return f.all, nil }

func (f *fakeAccounts) MarkVerified(_ context.Context, uid string) error {
	f.marked = append(f.marked, uid)
	return nil
}

// TestMarkChangesTheNamedAccountsAlone is D-903: a list run and a dry run
// change nothing, an apply run marks the named accounts alone, and an id
// with no account stops the run before any change.
func TestMarkChangesTheNamedAccountsAlone(t *testing.T) {
	day := time.Date(2026, 9, 1, 0, 0, 0, 0, time.UTC)
	list := allowlist.New(func(context.Context) ([]string, error) { return []string{"ann@example.com"}, nil })
	for _, tc := range []struct {
		name    string
		ids     []string
		apply   bool
		marked  []string
		wantErr string
		want    []string
	}{
		{"a list run names the accounts that are not proved", nil, false, nil, "", []string{"2 of 3 accounts are not proved", "u1\tann@example.com\tinvited=true", "u2\tbob@example.com\tinvited=false"}},
		{"a dry run changes nothing", []string{"u1"}, false, nil, "", []string{"would mark as proved"}},
		{"an apply run marks the named account", []string{"u1"}, true, []string{"u1"}, "", []string{"u1\tann@example.com\tmarked as proved"}},
		{"a proved account stays as it is", []string{"u3"}, true, nil, "", []string{"already proved"}},
		{"an unknown id stops the run", []string{"u1", "nobody"}, true, nil, `no account has the id "nobody"`, nil},
	} {
		t.Run(tc.name, func(t *testing.T) {
			acc := &fakeAccounts{all: []account{
				{UID: "u2", Email: "bob@example.com", Created: day.AddDate(0, 0, 1)},
				{UID: "u1", Email: "ann@example.com", Created: day},
				{UID: "u3", Email: "cy@example.com", Verified: true, Created: day},
			}}
			var out bytes.Buffer
			err := mark(context.Background(), &out, acc, list, tc.ids, tc.apply)
			if tc.wantErr != "" {
				if err == nil || !strings.Contains(err.Error(), tc.wantErr) {
					t.Fatalf("err = %v, want %q", err, tc.wantErr)
				}
			} else if err != nil {
				t.Fatal(err)
			}
			if !slices.Equal(acc.marked, tc.marked) {
				t.Errorf("marked = %v, want %v", acc.marked, tc.marked)
			}
			for _, w := range tc.want {
				if !strings.Contains(out.String(), w) {
					t.Errorf("output lacks %q:\n%s", w, out.String())
				}
			}
		})
	}
}
