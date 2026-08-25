package questions

import (
	"testing"

	mtgv1 "github.com/nkramber/mtg-deck-builder/go/gen/mtg/v1"
)

// step is one turn of a scripted conversation. want is the row ids the
// planner must ask, in ask order. fill names the keys the user's answers
// close. set applies the facts the answers reveal, for example the format.
type step struct {
	want []string
	fill []string
	set  func(*Context)
}

type conversation struct {
	name  string
	ctx   Context
	steps []step
}

func newCtx(words string) Context {
	return Context{
		Format: mtgv1.FormatId_FORMAT_ID_UNSPECIFIED,
		Filled: map[string]bool{}, Asked: map[string]bool{}, Words: words,
	}
}

func commander(c *Context) { c.Format = mtgv1.FormatId_FORMAT_ID_COMMANDER }

// conversations are the PR-7 gate cases. Each one must reach a complete
// slot set in at most four turns, with no repeated question. Eight of the
// twelve come from the dogfood runs of 2026-08-24.
func conversations() []conversation {
	var cs []conversation

	c1 := conversation{name: "lifegain with a collection"}
	c1.ctx = newCtx("build me a lifegain deck")
	c1.ctx.HasCollection, c1.ctx.Theme = true, "lifegain"
	c1.steps = []step{
		{want: []string{"format", "theme", "colors"}, fill: []string{"format", "theme", "colors"}, set: commander},
		{want: []string{"commander", "power_commander", "pool"}, fill: []string{"power", "pool_rule"},
			set: func(c *Context) { c.OwnedMode, c.BuyList, c.Suggested = true, true, true }},
		{want: []string{"commander_pick", "budget"}, fill: []string{"commander", "budget"},
			set: func(c *Context) { c.CommanderSet = true }},
	}
	cs = append(cs, c1)

	c2 := conversation{name: "blink with a thin library"}
	c2.ctx = newCtx("i want a blink deck from my library")
	c2.ctx.HasCollection, c2.ctx.ThinTheme, c2.ctx.Theme = true, true, "blink"
	c2.steps = []step{
		{want: []string{"format", "theme", "colors"}, fill: []string{"format", "theme", "colors"}, set: commander},
		{want: []string{"commander", "power_commander", "pool_thin"}, fill: []string{"power", "pool_rule", "commander"},
			set: func(c *Context) { c.OwnedMode, c.BuyList, c.CommanderSet = true, true, true }},
		{want: []string{"budget"}, fill: []string{"budget"}},
	}
	cs = append(cs, c2)

	c3 := conversation{name: "a named card, role unknown"}
	c3.ctx = newCtx("build around grist, the hunger tide")
	c3.ctx.HasCollection, c3.ctx.NamedCard = true, true
	c3.steps = []step{
		{want: []string{"format", "theme_card_named", "colors"}, fill: []string{"format", "theme", "colors"}, set: commander},
		{want: []string{"named_card_role", "power_commander", "pool"},
			fill: []string{"named_card_role", "commander", "power", "pool_rule"},
			set:  func(c *Context) { c.CommanderSet = true }},
		{want: []string{"locked"}, fill: []string{"locked"}},
	}
	cs = append(cs, c3)

	c4 := conversation{name: "fun and janky"}
	c4.ctx = newCtx("make me something fun and janky")
	c4.ctx.HasCollection = true
	c4.steps = []step{
		{want: []string{"format", "theme", "jank"}, fill: []string{"format", "theme", "jank"},
			set: func(c *Context) { commander(c); c.Theme = "sacrifice" }},
		{want: []string{"commander", "power_commander", "colors"}, fill: []string{"power", "colors"},
			set: func(c *Context) { c.Suggested = true }},
		{want: []string{"commander_pick", "pool"}, fill: []string{"commander", "pool_rule"},
			set: func(c *Context) { c.CommanderSet = true }},
	}
	cs = append(cs, c4)

	c5 := conversation{name: "60-card anything goes"}
	c5.ctx = newCtx("60-card anything goes")
	c5.steps = []step{
		{want: []string{"format", "theme", "house_rules"}, fill: []string{"format", "theme", "house_rules"},
			set: func(c *Context) {
				c.Format, c.HouseFormat, c.Theme = mtgv1.FormatId_FORMAT_ID_MODERN, true, "dragons"
			}},
		{want: []string{"house_format_limits", "power_sixty", "colors"},
			fill: []string{"house_format_limits", "power", "colors", "budget"}},
	}
	cs = append(cs, c5)

	c6 := conversation{name: "the strongest deck, no collection"}
	c6.ctx = newCtx("build the strongest deck possible i own nothing")
	c6.ctx.BuyList = true
	c6.steps = []step{
		{want: []string{"format", "theme", "colors"}, fill: []string{"format", "theme", "colors"},
			set: func(c *Context) {
				c.Format, c.PowerCompetitive, c.Theme = mtgv1.FormatId_FORMAT_ID_MODERN, true, "best deck"
			}},
		{want: []string{"power_sixty", "budget", "meta"}, fill: []string{"power", "budget", "meta"}},
	}
	cs = append(cs, c6)

	c7 := conversation{name: "FNM on Friday"}
	c7.ctx = newCtx("i need a deck for fnm on friday, something competitive")
	c7.ctx.BuyList, c7.ctx.Deadline = true, true
	c7.steps = []step{
		{want: []string{"format_store", "theme", "colors"}, fill: []string{"format", "theme", "colors"},
			set: func(c *Context) {
				c.Format, c.PowerCompetitive, c.Theme = mtgv1.FormatId_FORMAT_ID_PIONEER, true, "best deck"
			}},
		{want: []string{"power_sixty", "budget", "meta"}, fill: []string{"power", "budget", "meta"}},
		{want: []string{"acquisition"}, fill: []string{"acquisition"}},
	}
	cs = append(cs, c7)

	c8 := conversation{name: "mill for a playgroup"}
	c8.ctx = newCtx("make a mill deck for my playgroup, budget 100")
	c8.ctx.HasCollection, c8.ctx.Theme, c8.ctx.BudgetAmbiguous = true, "mill", true
	c8.steps = []step{
		{want: []string{"format", "theme", "colors"}, fill: []string{"format", "theme", "colors"},
			set: func(c *Context) { commander(c); c.TwoPlans = true }},
		{want: []string{"commander", "power_commander", "pool"}, fill: []string{"power", "pool_rule", "commander"},
			set: func(c *Context) { c.OwnedMode, c.CommanderSet = true, true }},
		{want: []string{"table_tolerance", "budget_scope", "plan_choice"},
			fill: []string{"table_tolerance", "budget_scope", "budget", "plan_variant"}},
	}
	cs = append(cs, c8)

	c9 := conversation{name: "a commander the library does not hold"}
	c9.ctx = newCtx("brago blink deck from my library")
	c9.ctx.HasCollection, c9.ctx.OwnedMode, c9.ctx.CommanderNotOwned = true, true, true
	c9.ctx.Theme, c9.ctx.Filled["format"], c9.ctx.Filled["theme"] = "blink", true, true
	c9.ctx.Format = mtgv1.FormatId_FORMAT_ID_COMMANDER
	c9.steps = []step{
		{want: []string{"commander_not_owned", "power_commander", "colors"},
			fill: []string{"commander", "power", "colors"}, set: func(c *Context) { c.CommanderSet, c.BuyList = true, true }},
		{want: []string{"pool", "budget"}, fill: []string{"pool_rule", "budget"}},
	}
	cs = append(cs, c9)

	c10 := conversation{name: "no strong commander in the library"}
	c10.ctx = newCtx("lifegain from my collection")
	c10.ctx.HasCollection, c10.ctx.OwnedMode, c10.ctx.WeakCommanderPool = true, true, true
	c10.ctx.Theme, c10.ctx.Format = "lifegain", mtgv1.FormatId_FORMAT_ID_COMMANDER
	c10.ctx.Filled["format"], c10.ctx.Filled["theme"], c10.ctx.Filled["colors"] = true, true, true
	c10.steps = []step{
		{want: []string{"commander_weak_pool", "power_commander", "pool"},
			fill: []string{"commander", "power", "pool_rule", "budget"}, set: func(c *Context) { c.CommanderSet = true }},
	}
	cs = append(cs, c10)

	c11 := conversation{name: "another version after a build"}
	c11.ctx = newCtx("give me another version")
	c11.ctx.AfterBuild = true
	c11.ctx.Format = mtgv1.FormatId_FORMAT_ID_COMMANDER
	for _, k := range []string{"format", "theme", "colors", "commander", "power", "pool_rule", "budget"} {
		c11.ctx.Filled[k] = true
	}
	c11.ctx.CommanderSet = true
	c11.steps = []step{
		{want: []string{"variance"}, fill: []string{"plan_variant"}},
	}
	cs = append(cs, c11)

	c12 := conversation{name: "a frozen run asks nothing"}
	c12.ctx = newCtx("switch to owned-only")
	c12.ctx.Frozen, c12.ctx.HasCollection = true, true
	c12.steps = []step{{want: nil}}
	cs = append(cs, c12)

	return cs
}

// TestConversations is the deterministic half of the PR-7 gate: a complete
// slot set in at most four turns, with no repeated question. The other
// half of the gate (at least 25 of 30 conversations use catalog questions
// only) needs the model in the loop, because only the model invents a
// question. Every question below comes from the catalog.
func TestConversations(t *testing.T) {
	c := load(t)
	for _, conv := range conversations() {
		t.Run(conv.name, func(t *testing.T) {
			ctx := conv.ctx
			if len(conv.steps) > 4 {
				t.Fatalf("%d turns, the gate allows 4", len(conv.steps))
			}
			everAsked := map[string]bool{}
			for i, st := range conv.steps {
				got := ids(c.Plan(ctx))
				if !equal(got, st.want) {
					t.Fatalf("turn %d asked %v, want %v", i+1, got, st.want)
				}
				if len(got) > MaxPerTurn {
					t.Fatalf("turn %d asked %d questions, max is %d", i+1, len(got), MaxPerTurn)
				}
				for _, id := range got {
					if everAsked[id] {
						t.Fatalf("turn %d repeated question %q", i+1, id)
					}
					everAsked[id] = true
					ctx.Asked[id] = true
				}
				for _, k := range st.fill {
					ctx.Filled[k] = true
				}
				if st.set != nil {
					st.set(&ctx)
				}
			}
			if left := c.Plan(ctx); len(left) > 0 {
				t.Errorf("slots still open after %d turns: %v", len(conv.steps), ids(left))
			}
		})
	}
}

func equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
