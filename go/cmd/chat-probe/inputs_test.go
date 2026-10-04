package main

import (
	"os"
	"path/filepath"
	"testing"

	"google.golang.org/protobuf/encoding/protojson"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
)

func TestTurnsJSONKeepsABar(t *testing.T) {
	p := filepath.Join(t.TempDir(), "m.json")
	if err := os.WriteFile(p, []byte(`["Build a deck | with a bar", "Bracket 3"]`), 0o600); err != nil {
		t.Fatal(err)
	}
	got, err := turns("ignored|text", p)
	if err != nil {
		t.Fatal(err)
	}
	if len(got) != 2 || got[0] != "Build a deck | with a bar" {
		t.Fatalf("turns = %q", got)
	}
}

func TestTurnsTextSplitsOnBar(t *testing.T) {
	got, err := turns(" a | b ", "")
	if err != nil {
		t.Fatal(err)
	}
	if len(got) != 2 || got[0] != "a" || got[1] != "b" {
		t.Fatalf("turns = %q", got)
	}
}

func TestTurnsJSONRefusesNoTurn(t *testing.T) {
	p := filepath.Join(t.TempDir(), "m.json")
	if err := os.WriteFile(p, []byte(`[]`), 0o600); err != nil {
		t.Fatal(err)
	}
	if _, err := turns("", p); err == nil {
		t.Fatal("an empty list was accepted")
	}
}

func TestOwnedFromJSONCountsCopies(t *testing.T) {
	c := &mtgv1.Collection{Entries: []*mtgv1.CollectionEntry{
		{OracleId: "o1", Quantity: 2}, {OracleId: "o1", Quantity: 1}, {OracleId: "o2", Quantity: 4},
	}}
	b, err := protojson.Marshal(c)
	if err != nil {
		t.Fatal(err)
	}
	p := filepath.Join(t.TempDir(), "c.json")
	if err := os.WriteFile(p, b, 0o600); err != nil {
		t.Fatal(err)
	}
	got, err := ownedFromJSON(p)
	if err != nil {
		t.Fatal(err)
	}
	if got["o1"] != 3 || got["o2"] != 4 {
		t.Fatalf("owned = %v", got)
	}
}
