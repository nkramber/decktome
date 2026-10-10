package questions

import (
	"slices"
	"testing"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
)

// pickHints offers the three commanders of the chat of D-1211, and it
// knows which names lead a deck.
type pickHints struct{}

var pickLeaders = []string{"Jace, Multiverse Architect", "Tam, the Possibility", "Tomik, Orzhov Lawmage"}

func (pickHints) Commanders(string, []string) []string { return pickLeaders }
func (pickHints) OwnedThemeCount(string) int           { return 0 }
func (pickHints) CanLead(name string) (bool, bool) {
	if hasName(pickLeaders, name) {
		return true, true
	}
	return name != "Sol Ring", true
}

// fixtureOut is a classify output of the chat of D-1211.
func fixtureOut(power string, declined []string, commanders, named []string, suggest bool) classifyOut {
	out := classifyOut{Format: "commander", Theme: "planeswalkers", PoolRule: "any_card", Power: power,
		DeclinedKeys: declined, CommanderNames: commanders, NamedCards: named}
	out.Facts.WantsSuggestion = suggest
	return out
}

// firstTurn asks the power row.
func firstTurn() turnScript {
	return turnScript{message: "Build me a Commander deck focused on planeswalkers",
		classify: fixtureOut("", nil, nil, nil, false)}
}

// toCommanderRow answers turn 1, so turn 2 asks the commander row.
func toCommanderRow() []turnScript {
	return []turnScript{firstTurn(),
		{message: "1 Exhibition, any colors, no budget", classify: fixtureOut("bracket 1", []string{"colors", "budget"}, nil, nil, false)}}
}

// toPickRow asks for three names, so turn 3 asks the pick row.
func toPickRow() []turnScript {
	return append(toCommanderRow(), turnScript{message: "Suggest three", classify: fixtureOut("", nil, nil, nil, true)})
}

func bracket(p played) int32 { return p.st.Slots.GetPower().GetBracket() }

func TestATypedBracketOptionKeepsItsBracket(t *testing.T) {
	// The classifier left the power empty and declined it. The build
	// used bracket 3 (D-1211).
	p := play(t, false, pickHints{}, []turnScript{
		firstTurn(),
		{message: "1 Exhibition. Fill the gaps from any card (buy list). Suggest three",
			classify: fixtureOut("", []string{"power"}, nil, nil, true)},
	})
	if !p.askedOn(1, "power_commander") {
		t.Fatalf("turn 1 asked %v, want the power row", p.rows[0])
	}
	if got := bracket(p); got != 1 {
		t.Errorf("bracket = %d, want 1", got)
	}
	if p.st.Ctx.Skipped["power"] {
		t.Error("the power is marked declined, and the reader answered it")
	}
}

func TestAClassifierBracketWins(t *testing.T) {
	p := play(t, false, pickHints{}, []turnScript{
		firstTurn(),
		{message: "1 Exhibition, or maybe bracket 2", classify: fixtureOut("bracket 2", nil, nil, nil, false)},
	})
	if got := bracket(p); got != 2 {
		t.Errorf("bracket = %d, want the classifier's 2", got)
	}
}

func TestANegatedBracketOptionFillsNoBracket(t *testing.T) {
	p := play(t, false, pickHints{}, []turnScript{
		firstTurn(),
		{message: "Not 1 Exhibition, something stronger", classify: fixtureOut("", nil, nil, nil, false)},
	})
	if p.st.Slots.GetPower() != nil {
		t.Errorf("power = %v, want none for a negated option", p.st.Slots.GetPower())
	}
}

func TestTwoBracketOptionsFillNoBracket(t *testing.T) {
	p := play(t, false, pickHints{}, []turnScript{
		firstTurn(),
		{message: "2 Core or 3 Upgraded, I can't decide", classify: fixtureOut("", nil, nil, nil, false)},
	})
	if p.st.Slots.GetPower() != nil {
		t.Errorf("power = %v, want none for two options", p.st.Slots.GetPower())
	}
}

func TestATypedNameUnderTheCommanderRowLeads(t *testing.T) {
	// The base-1 replay of D-1211: the classifier read the name as a
	// card and declined every open key, so the role row asked.
	p := play(t, false, pickHints{}, append(toCommanderRow(), turnScript{
		message:  "Skip that question.. Jace, Multiverse Architect",
		classify: fixtureOut("", []string{"commander"}, nil, []string{"Jace, Multiverse Architect"}, false)}))
	if !p.askedOn(2, "commander") {
		t.Fatalf("turn 2 asked %v, want the commander row", p.rows[1])
	}
	if !slices.Equal(p.st.CommanderNames, []string{"Jace, Multiverse Architect"}) {
		t.Errorf("commander = %v, want Jace, Multiverse Architect", p.st.CommanderNames)
	}
	if p.asked("named_card_role") {
		t.Errorf("the role row asked, and the reader named the commander: %v", p.rows)
	}
}

func TestATypedNameUnderThePickRowLeads(t *testing.T) {
	// D-1213: a valid name the pick row did not offer becomes the
	// commander.
	p := play(t, false, pickHints{}, append(toPickRow(), turnScript{
		message:  "Atraxa, Praetors' Voice",
		classify: fixtureOut("", nil, nil, []string{"Atraxa, Praetors' Voice"}, false)}))
	if !p.askedOn(3, "commander_pick") {
		t.Fatalf("turn 3 asked %v, want the pick row", p.rows[2])
	}
	if !slices.Equal(p.st.CommanderNames, []string{"Atraxa, Praetors' Voice"}) {
		t.Errorf("commander = %v, want Atraxa, Praetors' Voice", p.st.CommanderNames)
	}
	if p.asked("named_card_role") {
		t.Errorf("the role row asked: %v", p.rows)
	}
}

func TestOtherWordsKeepTheRoleOpen(t *testing.T) {
	for _, msg := range []string{
		"Not Jace, Multiverse Architect",
		"Put Jace, Multiverse Architect in the 99",
		"Sol Ring",
	} {
		t.Run(msg, func(t *testing.T) {
			name := "Jace, Multiverse Architect"
			if msg == "Sol Ring" {
				name = "Sol Ring"
			}
			p := play(t, false, pickHints{}, append(toCommanderRow(), turnScript{
				message: msg, classify: fixtureOut("", nil, nil, []string{name}, false)}))
			if p.st.Ctx.CommanderSet {
				t.Errorf("commander = %v, want none", p.st.CommanderNames)
			}
		})
	}
}

func TestASkipDeclinesTheOneQuestionLeft(t *testing.T) {
	// The c4-5 replay of D-1224: the name answered the pick row, the
	// classifier declined nothing, and the budget question stayed out.
	// Turn 1 leaves the budget out, so the pick row and the budget
	// question are out together, as in the replay.
	p := play(t, false, pickHints{}, []turnScript{
		firstTurn(),
		{message: "1 Exhibition, any colors", classify: fixtureOut("bracket 1", []string{"colors"}, nil, nil, false)},
		{message: "Suggest three", classify: fixtureOut("", nil, nil, nil, true)},
		{message: "Skip that question.. Jace, Multiverse Architect",
			classify: fixtureOut("", nil, []string{"Jace, Multiverse Architect"}, nil, false)},
	})
	if !p.askedOn(3, "commander_pick") {
		t.Fatalf("turn 3 asked %v, want the pick row", p.rows[2])
	}
	for k, s := range p.st.Slots.GetSlotStates() {
		if s == mtgv1.SlotState_SLOT_STATE_ASKED {
			t.Errorf("key %s is still out after the skip", k)
		}
	}
	if !p.ready[3] {
		t.Errorf("turn 4 is not ready; rows %v", p.rows)
	}
}

func TestSkipWordsNeedOneQuestionLeft(t *testing.T) {
	cases := map[string]bool{
		"Skip that question":       true,
		"skip it":                  true,
		"Don't skip it":            false,
		"Skipper":                  false,
		"skip":                     true,
		"Skip that.":               true,
		"skip the expensive cards": false,
		"skip combo pieces":        false,
		"Jace":                     false,
	}
	for msg, want := range cases {
		if got := skipsQuestion(msg); got != want {
			t.Errorf("skipsQuestion(%q) = %v, want %v", msg, got, want)
		}
	}
}
