package main

import (
	"strings"
	"testing"

	"github.com/nkramber/decktome/go/internal/meta"
	"github.com/nkramber/decktome/go/internal/precons"
)

// TestResolvePreconsIgnoresBasicLands is D-523 for the gate: a binder
// that holds every nonbasic printing of the product owns it, so the
// prompt resolves. With no basic land test the same binder falls short,
// and a missing nonbasic card fails the prompt either way.
func TestResolvePreconsIgnoresBasicLands(t *testing.T) {
	card := func(name, oracle string, n int) meta.PreconCard {
		return meta.PreconCard{Name: name, Count: n, OracleID: oracle, ScryfallID: oracle + "-p", SetCode: "msc", Number: "1"}
	}
	tbl := precons.NewTable("v1", []meta.Precon{{
		Name: "Avengers Assemble", Code: "MSC", Type: "Commander Deck", ReleaseDate: "2026-06-26",
		Commanders: []meta.PreconCard{card("Captain America, Team Leader", "o-cap", 1)},
		Cards:      []meta.PreconCard{card("Sol Ring", "o-sol", 1), card("Plains", "o-plains", 6)},
	}})
	isBasic := func(id string) bool { return id == "o-plains" }
	noBasics := map[string]int32{"o-cap-p": 1, "o-sol-p": 1}

	products, names, err := resolvePrecons(tbl, []string{"Avengers Assemble"}, noBasics, isBasic)
	if err != nil || len(products) != 1 || names[0] != "Avengers Assemble" {
		t.Fatalf("a binder with no basic lands owns the product: %v %v %v", products, names, err)
	}
	if _, _, err := resolvePrecons(tbl, []string{"Avengers Assemble"}, noBasics, nil); err == nil || !strings.Contains(err.Error(), "does not hold Avengers Assemble whole") {
		t.Errorf("with no basic land test every printing counts: %v", err)
	}
	noSol := map[string]int32{"o-cap-p": 1, "o-plains-p": 6}
	if _, _, err := resolvePrecons(tbl, []string{"Avengers Assemble"}, noSol, isBasic); err == nil {
		t.Error("a missing nonbasic card still fails the prompt")
	}
	if _, _, err := resolvePrecons(tbl, []string{"Justice League"}, noBasics, isBasic); err == nil || !strings.Contains(err.Error(), "names no product") {
		t.Errorf("an unknown name fails the prompt: %v", err)
	}
}
