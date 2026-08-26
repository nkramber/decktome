package main

import (
	"context"
	"io"
	"log/slog"
	"os"
	"regexp"
	"sort"
	"strings"

	mtgv1 "github.com/nkramber/mtg-deck-builder/go/gen/mtg/v1"
	"github.com/nkramber/mtg-deck-builder/go/internal/cards"
	"github.com/nkramber/mtg-deck-builder/go/internal/tune"
)

// The eval invents card facts of its own, and a false one changes a
// verdict. Eval run 14 refused three questions on two such claims:
//
//	"'Ran and Shaw' is not a real Magic card or commander option."
//	"'Quina, Qu Gourmet' is not a valid Magic card option."
//
// Both cards are real. Ran and Shaw is a Legendary Creature - Dragon,
// mono-red, and the request was for a red dragon deck. Quina, Qu Gourmet
// is a Legendary Creature - Qu, mono-green. Both are legal in Commander.
// The calibration of 2026-08-26 showed that claude-sonnet-5 invents card
// facts as well, so no second judge fixes this class (D-149).
//
// The check is deterministic and it costs nothing. Every claim of this
// shape names the card, and the snapshot answers whether it exists.

// denialSigns mark a claim that a card does not exist. The list is narrow
// on purpose. A reason that says a card is wrong for the deck, or off
// color, makes no claim about its existence and must reach the report.
var denialSigns = []string{
	"is not a real", "is not a valid", "is not an actual",
	"is not a magic card", "is not a legal magic card",
	"not a real magic card", "not a valid magic card",
	"does not exist", "doesn't exist", "do not exist",
	"nonexistent", "non-existent", "not an existing",
	"fictional", "made-up", "made up", "invented card",
	"is not a card",
}

// quotedName reads a card name the reason puts in quotes. Both false
// claims of run 14 quoted the name. The bounds keep a whole sentence out.
var quotedName = regexp.MustCompile(`["'\x{201C}\x{2018}]([^"'\x{201C}\x{201D}\x{2018}\x{2019}]{3,60})["'\x{201D}\x{2019}]`)

// cardChecker answers whether a card name is real, from the local
// snapshot and never from memory.
type cardChecker struct{ idx *cards.Index }

// newCardChecker reads CARDS_SNAPSHOT_DIR. A missing snapshot disables
// the check and never fails the run: the eval still reports, and the
// document says the check did not run.
func newCardChecker() (cardChecker, string) {
	dir := os.Getenv("CARDS_SNAPSHOT_DIR")
	if dir == "" {
		return cardChecker{}, "no card snapshot: the eval card check did not run"
	}
	quiet := slog.New(slog.NewTextHandler(io.Discard, nil))
	idx, err := cards.LoadIndex(context.Background(), cards.DirStore{Root: dir}, quiet)
	if err != nil || idx == nil {
		return cardChecker{}, "the card snapshot did not load: the eval card check did not run"
	}
	return cardChecker{idx: idx}, ""
}

// deniesExistence reports whether the reason claims a card is not real.
// A reason that says a card is wrong for the deck, or off color, makes no
// such claim and must reach the report untouched.
func deniesExistence(reason string) bool {
	low := strings.ToLower(reason)
	for _, s := range denialSigns {
		if strings.Contains(low, s) {
			return true
		}
	}
	return false
}

// quotedNames reads every quoted name of the reason, in order.
func quotedNames(reason string) []string {
	var out []string
	for _, m := range quotedName.FindAllStringSubmatch(reason, -1) {
		if name := strings.TrimSpace(m[1]); name != "" {
			out = append(out, name)
		}
	}
	return out
}

// deniedRealCard returns the card the reason wrongly calls unreal.
func (c cardChecker) deniedRealCard(reason string) (string, bool) {
	if c.idx == nil || reason == "" || !deniesExistence(reason) {
		return "", false
	}
	for _, name := range quotedNames(reason) {
		if _, ok := c.idx.ByName(name); ok {
			return name, true
		}
	}
	return "", false
}

// correct drops a refusal that rests on a false card claim. The
// inaccurate fault goes, because the snapshot refutes it. A verdict that
// held another fault keeps that fault and stays bad: "duplicate,
// inaccurate" on a repeated question is still a duplicate.
func (c cardChecker) correct(v tune.Verdict) (tune.Verdict, string, bool) {
	name, ok := c.deniedRealCard(v.Reason)
	if !ok {
		return v, "", false
	}
	kept := make([]string, 0, len(v.Faults))
	for _, f := range v.Faults {
		if strings.EqualFold(strings.TrimSpace(f), "inaccurate") {
			continue
		}
		kept = append(kept, f)
	}
	v.Faults = kept
	// Nothing is left against the question, so it was warranted.
	if len(kept) == 0 {
		v.Warranted = "yes"
		v.CatalogAction = "none"
	}
	v.Reason = strings.TrimSpace(v.Reason) +
		" [D-149: the snapshot holds " + name + ", so the claim that it is not a real card is false.]"
	return v, name, true
}

// writeCardCheck reports what the card check refuted. A run with nothing
// to report still says so, because silence reads as "the check did not
// run" (D-149).
func writeCardCheck(w io.Writer, note string, corrected map[string]string) {
	_, _ = io.WriteString(w, "\n## The card check\n\n")
	if note != "" {
		_, _ = io.WriteString(w, note+"\n")
		return
	}
	if len(corrected) == 0 {
		_, _ = io.WriteString(w,
			"The check ran against the local snapshot. No refusal rested on a card name the snapshot refutes.\n")
		return
	}
	_, _ = io.WriteString(w, "The eval called these cards unreal, and the snapshot holds every one. "+
		"The `inaccurate` fault was dropped from each verdict (D-149).\n\n| Conversation and card | Row |\n|---|---|\n")
	keys := make([]string, 0, len(corrected))
	for k := range corrected {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		_, _ = io.WriteString(w, "| "+k+" | `"+corrected[k]+"` |\n")
	}
}

// cardFact is what the eval is told about one card a question names.
// The eval invents facts about a card it does not know, and every miss so
// far came from a crossover set: Quina, Qu Gourmet from Final Fantasy,
// Ran and Shaw from Avatar, Cloak and Dagger, Entwined and Shadow the
// Hedgehog from Marvel and Sonic. It called four of them unreal or
// inapplicable, and all four are real and legal.
//
// The check of D-149 refutes a claim after the model makes it, and it
// reads existence alone. This block prevents the claim instead, and it
// carries the color identity as well, so an applicability claim has the
// same ground under it (D-152).
type cardFact struct {
	Name          string   `json:"name"`
	TypeLine      string   `json:"type_line"`
	ColorIdentity []string `json:"color_identity"`
	CommanderList bool     `json:"legal_as_a_commander"`
}

// colorLetter names a color the way Scryfall does.
var colorLetter = map[mtgv1.Color]string{
	mtgv1.Color_COLOR_W: "W", mtgv1.Color_COLOR_U: "U", mtgv1.Color_COLOR_B: "B",
	mtgv1.Color_COLOR_R: "R", mtgv1.Color_COLOR_G: "G",
}

// facts reads every card the questions of one conversation name, and
// answers what the snapshot holds about each one.
//
// It reads the options, because that is where the pick row puts the
// commander names, and every invented claim so far was about a name the
// pick row offered. A string that the index does not hold is left out:
// the eval is told about real cards, and nothing is claimed about the
// rest.
func (c cardChecker) facts(conv tune.Conversation) []cardFact {
	if c.idx == nil {
		return nil
	}
	seen := map[string]bool{}
	var out []cardFact
	for _, q := range conv.Questions {
		for _, opt := range q.Options {
			name := strings.TrimSpace(opt)
			if name == "" || seen[strings.ToLower(name)] {
				continue
			}
			card, ok := c.idx.ByName(name)
			if !ok {
				continue
			}
			seen[strings.ToLower(name)] = true
			f := cardFact{
				Name:          card.GetName(),
				TypeLine:      card.GetTypeLine(),
				CommanderList: card.GetCanBeCommander() || card.GetIsBackground(),
			}
			for _, col := range card.GetColorIdentity() {
				if l, ok := colorLetter[col]; ok {
					f.ColorIdentity = append(f.ColorIdentity, l)
				}
			}
			out = append(out, f)
		}
	}
	return out
}
