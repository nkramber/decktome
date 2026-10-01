package rules

import (
	"slices"
	"strings"
	"testing"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
	"github.com/nkramber/decktome/go/internal/stale"
)

// banSource reads the test index, with one card banned in one format.
type banSource struct {
	id, key string
}

func (b banSource) ByOracleID(id string) (*mtgv1.Card, bool) {
	card, ok := testIndex.ByOracleID(id)
	if !ok || id != b.id {
		return card, ok
	}
	out := &mtgv1.Card{Name: card.GetName(), TypeLine: card.GetTypeLine(), Legalities: map[string]mtgv1.LegalityStatus{}}
	for k, v := range card.GetLegalities() {
		out.Legalities[k] = v
	}
	out.Legalities[b.key] = mtgv1.LegalityStatus_LEGALITY_STATUS_BANNED
	return out, true
}

// statusWord is the Scryfall word of a legality status.
func statusWord(s mtgv1.LegalityStatus) string {
	switch s {
	case mtgv1.LegalityStatus_LEGALITY_STATUS_LEGAL:
		return "legal"
	case mtgv1.LegalityStatus_LEGALITY_STATUS_RESTRICTED:
		return "restricted"
	case mtgv1.LegalityStatus_LEGALITY_STATUS_BANNED:
		return "banned"
	}
	return "not_legal"
}

// checkSyntheticBans is the gate of I-1 on one good golden deck. It bans
// each card of the deck in turn, and it asks two things of each ban:
//
//   - stale.Find names the banned card alone, and the rules engine
//     blocks that card as banned_card, so the pass and the engine agree.
//   - stale.Classify names the case of D-1008: a rebuild for the
//     commander or a tenth of the nonland copies, and a patch otherwise.
//
// The rules engine checks the legality of the rerun deck at each build,
// so the second half of the gate, a legal deck, holds at the build.
func checkSyntheticBans(t *testing.T, d *mtgv1.Deck) {
	t.Helper()
	key := testCfg.Formats[d.GetFormat().GetId().String()].ScryfallKey
	if key == "" {
		return
	}
	legal := stale.Legalities{}
	ids := append(slices.Clone(d.GetCommanderOracleIds()), d.GetCompanionOracleId())
	lands := map[string]bool{}
	for _, list := range [][]*mtgv1.DeckCard{d.GetCards(), d.GetSideboard()} {
		for _, dc := range list {
			ids = append(ids, dc.GetOracleId())
		}
	}
	for _, dc := range d.GetCards() {
		if card, ok := testIndex.ByOracleID(dc.GetOracleId()); ok && strings.Contains(card.GetTypeLine(), "Land") {
			dc.Role = mtgv1.CardRole_CARD_ROLE_LAND
			lands[dc.GetOracleId()] = true
		}
	}
	for _, id := range ids {
		if card, ok := testIndex.ByOracleID(id); ok && id != "" {
			legal[id] = map[string]string{key: statusWord(card.GetLegalities()[key])}
		}
	}
	if got := stale.Find(d, key, legal); len(got) > 0 {
		t.Fatalf("a legal golden deck reads stale: %v", got)
	}
	var nonland int
	copies := map[string]int{}
	for _, dc := range d.GetCards() {
		copies[dc.GetOracleId()] += int(dc.GetCount())
		if !lands[dc.GetOracleId()] {
			nonland += int(dc.GetCount())
		}
	}
	seen := map[string]bool{}
	for _, id := range ids {
		if id == "" || seen[id] {
			continue
		}
		seen[id] = true
		ban := stale.Legalities{}
		for k, v := range legal {
			ban[k] = v
		}
		ban[id] = map[string]string{key: "banned"}
		got := stale.Find(d, key, ban)
		if !slices.Equal(got, []string{id}) {
			t.Errorf("ban of %s: Find = %v", id, got)
			continue
		}
		res := testCfg.Validate(Input{Deck: d, Cards: banSource{id: id, key: key}})
		var blocked []string
		for _, f := range res.GetFindings() {
			if f.GetCode() == CodeBannedCard {
				blocked = append(blocked, f.GetOracleId())
			}
		}
		if !slices.Equal(blocked, []string{id}) {
			t.Errorf("ban of %s: the engine blocks %v", id, blocked)
		}
		want := mtgv1.RerunCase_RERUN_CASE_PATCH
		switch {
		case slices.Contains(d.GetCommanderOracleIds(), id):
			want = mtgv1.RerunCase_RERUN_CASE_REBUILD
		case !lands[id] && nonland > 0 && float64(copies[id]) >= stale.ShareLimit*float64(nonland):
			want = mtgv1.RerunCase_RERUN_CASE_REBUILD
		}
		if v := stale.Classify(d, got); v.Case != want {
			t.Errorf("ban of %s: case %v, want %v (%s)", id, v.Case, want, v.Reason)
		}
	}
	t.Logf("synthetic bans: %d", len(seen))
}
