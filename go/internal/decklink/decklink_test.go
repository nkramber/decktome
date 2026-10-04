package decklink

import (
	"errors"
	"strings"
	"testing"
)

func TestParse(t *testing.T) {
	for _, tc := range []struct {
		raw  string
		want Link
		err  error
	}{
		{"https://archidekt.com/decks/3045150/modern_artifact_burn", Link{Site: Archidekt, Host: "archidekt.com", DeckID: 3045150}, nil},
		{"www.Archidekt.com/decks/7031486", Link{Site: Archidekt, Host: "archidekt.com", DeckID: 7031486}, nil},
		{"https://moxfield.com/decks/VaR9P-HceECgmgm55DC7ow", Link{Site: Moxfield, Host: "moxfield.com"}, nil},
		{"www.moxfield.com/decks/VaR9P-HceECgmgm55DC7ow/primer", Link{Site: Moxfield, Host: "moxfield.com"}, nil},
		{"https://tappedout.net/mtg-decks/some-deck/", Link{Site: Unknown, Host: "tappedout.net"}, nil},
		{"http://www.mtggoldfish.com/deck/123", Link{Site: Unknown, Host: "mtggoldfish.com"}, nil},
		{"https://archidekt.com/folders/1", Link{}, ErrNotLink},
		{"https://moxfield.com/users/someone", Link{}, ErrNotLink},
		{"https://moxfield.com/decks/", Link{}, ErrNotLink},
		{"1 Sol Ring", Link{}, ErrNotLink},
		{"hello", Link{}, ErrNotLink},
		{"javascript:alert(1)", Link{}, ErrNotLink},
		{"ftp://tappedout.net/x", Link{}, ErrNotLink},
		{"https://user@tappedout.net/x", Link{}, ErrNotLink},
		{"", Link{}, ErrNotLink},
		{"https://tappedout.net/" + strings.Repeat("a", maxBytes), Link{}, ErrNotLink},
	} {
		got, err := Parse(tc.raw)
		if !errors.Is(err, tc.err) || got != tc.want {
			t.Errorf("Parse(%q) = %+v, %v; want %+v, %v", tc.raw, got, err, tc.want, tc.err)
		}
	}
}

// TestSteps is D-1103: Archidekt needs none, Moxfield has its exact
// menu, and any other site has the general steps.
func TestSteps(t *testing.T) {
	if s := (Link{Site: Archidekt}).Steps(); s != nil {
		t.Errorf("Archidekt steps = %v", s)
	}
	mox := strings.Join((Link{Site: Moxfield}).Steps(), " ")
	if !strings.Contains(mox, "Select More, then Export.") || !strings.Contains(mox, "Select Copy for Arena.") {
		t.Errorf("Moxfield steps = %q", mox)
	}
	other := Link{Site: Unknown, Host: "tappedout.net"}
	if strings.Join(other.Steps(), " ") == mox || len(other.Steps()) == 0 {
		t.Errorf("general steps = %v", other.Steps())
	}
	if other.Name() != "tappedout.net" || (Link{Site: Moxfield}).Name() != "Moxfield" {
		t.Errorf("names = %q, %q", other.Name(), (Link{Site: Moxfield}).Name())
	}
}
