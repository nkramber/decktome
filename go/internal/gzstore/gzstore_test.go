package gzstore

import (
	"bytes"
	"compress/gzip"
	"errors"
	"strings"
	"testing"

	"google.golang.org/protobuf/proto"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
)

func TestJSONRoundTrip(t *testing.T) {
	want := map[string]int32{"a": 1, "b": 2}
	payload, err := MarshalJSON(want)
	if err != nil {
		t.Fatal(err)
	}
	var got map[string]int32
	if err := UnmarshalJSON(payload, &got); err != nil {
		t.Fatal(err)
	}
	if got["a"] != 1 || got["b"] != 2 {
		t.Errorf("got %v", got)
	}
}

func TestProtoRoundTrip(t *testing.T) {
	want := &mtgv1.Deck{Id: "d1", Name: "a deck", Cards: []*mtgv1.DeckCard{{Name: "Plains", Count: 30}}}
	payload, err := MarshalProto(want)
	if err != nil {
		t.Fatal(err)
	}
	var got mtgv1.Deck
	if err := UnmarshalProto(payload, &got); err != nil {
		t.Fatal(err)
	}
	if !proto.Equal(want, &got) {
		t.Errorf("got %v", &got)
	}
}

// TestEmptyPayloadOpens: a document from before the field opens as a
// zero value.
func TestEmptyPayloadOpens(t *testing.T) {
	var got map[string]int32
	if err := UnmarshalJSON(nil, &got); err != nil {
		t.Fatal(err)
	}
	if len(got) != 0 {
		t.Errorf("got %v", got)
	}
	var d mtgv1.Deck
	if err := UnmarshalProto(nil, &d); err != nil {
		t.Fatal(err)
	}
}

// TestInflateLimit: a payload past MaxInflatedBytes is refused, and one
// at the limit is not.
func TestInflateLimit(t *testing.T) {
	over, err := Marshal(bytes.Repeat([]byte("a"), MaxInflatedBytes+1))
	if err != nil {
		t.Fatal(err)
	}
	if _, err := Unmarshal(over); !errors.Is(err, ErrTooLarge) {
		t.Errorf("err = %v, want ErrTooLarge", err)
	}
	at, err := Marshal(bytes.Repeat([]byte("a"), MaxInflatedBytes))
	if err != nil {
		t.Fatal(err)
	}
	raw, err := Unmarshal(at)
	if err != nil || len(raw) != MaxInflatedBytes {
		t.Errorf("a payload at the limit: len %d, err %v", len(raw), err)
	}
}

func TestBadGzipIsAnError(t *testing.T) {
	if _, err := Unmarshal([]byte("not gzip")); err == nil {
		t.Error("bad gzip opened")
	}
	var buf bytes.Buffer
	_ = gzip.NewWriter(&buf).Close()
	if raw, err := Unmarshal(buf.Bytes()); err != nil || len(raw) != 0 {
		t.Errorf("an empty gzip stream: %q, %v", raw, err)
	}
}

func TestValidID(t *testing.T) {
	for id, want := range map[string]bool{
		"abc":                             true,
		"sess-1":                          true,
		"":                                false,
		".":                               false,
		"..":                              false,
		"a/b":                             false,
		"../other":                        false,
		strings.Repeat("x", MaxIDBytes):   true,
		strings.Repeat("x", MaxIDBytes+1): false,
	} {
		if got := ValidID(id); got != want {
			t.Errorf("ValidID(%q) = %v, want %v", id, got, want)
		}
	}
}
