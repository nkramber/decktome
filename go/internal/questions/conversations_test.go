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
// first fourteen come from the dogfood runs of 2026-08-24. Every catalog
// row fires in at least one of the thirty.
func conversations() []conversation {
	var cs []conversation

	c1 := conversation{name: "lifegain with a collection"}
	c1.ctx = newCtx("build me a lifegain deck")
	c1.ctx.HasCollection, c1.ctx.Theme = true, "lifegain"
	c1.steps = []step{
		{want: []string{"format", "theme", "colors"}, fill: []string{"format", "theme", "colors"}, set: commander},
		{want: []string{"commander", "power_commander", "pool"}, fill: []string{"power", "pool_rule"},
			set: func(c *Context) { c.OwnedMode, c.BuyList, c.Suggested = true, true, true }},
		{want: []string{"commander_pick", "budget"}, fill: []string{"commander", "commander_pick", "budget"},
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

	// The named card becomes the commander, so no card is locked. The
	// locked row must stay silent here: the live run of 2026-08-24 asked
	// the user to keep or cut a list that held only their commander (D-70).
	c3 := conversation{name: "a named card, role unknown"}
	c3.ctx = newCtx("build around grist, the hunger tide")
	c3.ctx.HasCollection, c3.ctx.NamedCard = true, true
	c3.steps = []step{
		{want: []string{"format", "theme_card_named", "colors"}, fill: []string{"format", "theme", "colors"}, set: commander},
		{want: []string{"named_card_role", "power_commander", "pool"},
			fill: []string{"named_card_role", "commander", "power", "pool_rule"},
			set:  func(c *Context) { c.CommanderSet = true }},
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
		{want: []string{"commander_pick", "pool"}, fill: []string{"commander", "commander_pick", "pool_rule"},
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
				// The agent fills the tournament step here, because the
				// user named none, and the confirm row asks about it.
				c.PowerInferred = true
			}},
		{want: []string{"budget", "meta"}, fill: []string{"budget", "meta"}},
	}
	cs = append(cs, c6)

	c7 := conversation{name: "FNM on Friday"}
	c7.ctx = newCtx("i need a deck for fnm on friday, something competitive")
	c7.ctx.BuyList = true
	c7.steps = []step{
		{want: []string{"format_store", "theme", "colors"}, fill: []string{"format", "theme", "colors"},
			set: func(c *Context) {
				c.Format, c.PowerCompetitive, c.Theme = mtgv1.FormatId_FORMAT_ID_MODERN, true, "best deck"
				// The agent fills the tournament step here, because the
				// user named none, and the confirm row asks about it.
				c.PowerInferred = true
			}},
		{want: []string{"budget", "meta"}, fill: []string{"budget", "meta"}},
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
		{want: []string{"budget_scope", "plan_choice"},
			fill: []string{"budget_scope", "budget", "plan_variant"}},
	}
	cs = append(cs, c8)

	c9 := conversation{name: "a commander the library does not hold"}
	c9.ctx = newCtx("brago blink deck from my library")
	c9.ctx.HasCollection, c9.ctx.OwnedMode = true, true
	c9.ctx.Theme, c9.ctx.Filled["format"], c9.ctx.Filled["theme"] = "blink", true, true
	c9.ctx.Format = mtgv1.FormatId_FORMAT_ID_COMMANDER
	c9.steps = []step{
		// The not-owned row is retired (D-226), so the base commander row
		// asks and the rules engine reports ownership after the build.
		{want: []string{"commander", "power_commander", "colors"},
			fill: []string{"commander", "power", "colors"}, set: func(c *Context) { c.CommanderSet, c.BuyList = true, true }},
		{want: []string{"pool", "budget"}, fill: []string{"pool_rule", "budget"}},
	}
	cs = append(cs, c9)

	// The weak-pool row is retired (D-232), so a thin library asks the
	// plain commander question and PR-8 reports the shortfall with the
	// deck.
	c10 := conversation{name: "no strong commander in the library"}
	c10.ctx = newCtx("lifegain from my collection")
	c10.ctx.HasCollection, c10.ctx.OwnedMode = true, true
	c10.ctx.Theme, c10.ctx.Format = "lifegain", mtgv1.FormatId_FORMAT_ID_COMMANDER
	c10.ctx.Filled["format"], c10.ctx.Filled["theme"], c10.ctx.Filled["colors"] = true, true, true
	c10.steps = []step{
		{want: []string{"commander", "power_commander", "pool"},
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
	// The variance row is retired with PR-9 (D-256). A message after a
	// build asks nothing: every slot is settled, so the session is ready
	// and the build runs again.
	c11.steps = []step{{want: nil}}
	cs = append(cs, c11)

	// The other half of D-70: a named card that is not the commander is a
	// locked card, and the locked row fires for it.
	c13 := conversation{name: "a card to keep that is not the commander"}
	c13.ctx = newCtx("karlov lifegain deck, and keep sanguine bond")
	c13.ctx.HasCollection, c13.ctx.NamedCard, c13.ctx.LockedCard = true, true, true
	c13.ctx.CommanderSet, c13.ctx.Theme = true, "lifegain"
	c13.ctx.Format = mtgv1.FormatId_FORMAT_ID_COMMANDER
	for _, k := range []string{"format", "theme", "colors", "commander", "commander_pick", "named_card_role"} {
		c13.ctx.Filled[k] = true
	}
	c13.steps = []step{
		{want: []string{"power_commander", "pool", "locked"},
			fill: []string{"power", "pool_rule", "locked"}},
	}
	cs = append(cs, c13)

	// D-73: the pick row repeats until the user takes a commander. Every
	// round must name three others, which the hint source enforces.
	c14 := conversation{name: "the user says none, then picks"}
	c14.ctx = newCtx("lifegain commander deck")
	c14.ctx.Theme, c14.ctx.Format = "lifegain", mtgv1.FormatId_FORMAT_ID_COMMANDER
	for _, k := range []string{"format", "theme", "colors"} {
		c14.ctx.Filled[k] = true
	}
	c14.steps = []step{
		{want: []string{"commander", "power_commander"}, fill: []string{"power"},
			set: func(c *Context) { c.Suggested = true }},
		// The user answers "none of those", which retires the names. The
		// row asks again only for that reason now: three names the user
		// has not seen (D-163).
		{want: []string{"commander_pick"}, set: func(c *Context) { c.OfferChanged = true }},
		{want: []string{"commander_pick"}, fill: []string{"commander", "commander_pick"},
			set: func(c *Context) { c.CommanderSet = true }},
	}
	cs = append(cs, c14)

	// The freeze is retired (D-241). A slot change after a build is an
	// ordinary turn: it updates the slot and asks nothing new, because
	// every other slot is already settled.
	c12 := conversation{name: "a slot change after a build"}
	c12.ctx = newCtx("switch to owned-only")
	c12.ctx.HasCollection, c12.ctx.AfterBuild = true, true
	c12.ctx.Format = mtgv1.FormatId_FORMAT_ID_COMMANDER
	for _, k := range []string{"format", "theme", "colors", "commander", "power", "pool_rule", "budget"} {
		c12.ctx.Filled[k] = true
	}
	c12.ctx.CommanderSet = true
	c12.steps = []step{{want: nil}}
	cs = append(cs, c12)

	// Conversations 15 to 30 widen the gate to the 30 the roadmap asks
	// for. Each one exercises a trigger the first fourteen do not reach.

	c15 := conversation{name: "standard at the store, no library"}
	c15.ctx = newCtx("i need a standard deck for my local store")
	c15.steps = []step{
		{want: []string{"format_store", "theme", "colors"}, fill: []string{"format", "theme", "colors"},
			set: func(c *Context) { c.Format, c.Theme = mtgv1.FormatId_FORMAT_ID_STANDARD, "aggro" }},
		{want: []string{"power_sixty"}, fill: []string{"power"}},
	}
	cs = append(cs, c15)

	// The competitive theme row fires only when the user asks for power in
	// the first message. Conversations 6 and 7 learn it one turn later.
	c16 := conversation{name: "the strongest modern deck"}
	c16.ctx = newCtx("i want the strongest modern deck, money is no object")
	c16.ctx.PowerCompetitive, c16.ctx.BuyList = true, true
	c16.steps = []step{
		{want: []string{"format", "theme_competitive", "colors"}, fill: []string{"format", "theme", "colors"},
			set: func(c *Context) {
				c.Format, c.Theme = mtgv1.FormatId_FORMAT_ID_MODERN, "best deck"
				c.PowerInferred = true
			}},
		{want: []string{"budget", "meta"}, fill: []string{"budget", "meta"}},
	}
	cs = append(cs, c16)

	// A set commander fills the color slot, so the card-pool question can
	// still fire. Without that rule the pool row waits forever (D-67).
	c17 := conversation{name: "precon upgrade at bracket 2"}
	c17.ctx = newCtx("upgrade my atraxa precon, we play bracket 2")
	c17.ctx.HasCollection, c17.ctx.CommanderSet, c17.ctx.Precon = true, true, true
	c17.ctx.Format, c17.ctx.Theme = mtgv1.FormatId_FORMAT_ID_COMMANDER, "superfriends"
	for _, k := range []string{"format", "theme", "commander", "commander_pick", "named_card_role", "colors"} {
		c17.ctx.Filled[k] = true
	}
	c17.steps = []step{
		{want: []string{"power_commander", "pool_precon"}, fill: []string{"power", "pool_rule"}},
	}
	cs = append(cs, c17)

	c18 := conversation{name: "burn on a budget"}
	c18.ctx = newCtx("standard burn deck, as cheap as possible")
	c18.ctx.BuyList = true
	c18.steps = []step{
		{want: []string{"format", "theme", "colors"}, fill: []string{"format", "theme", "colors"},
			set: func(c *Context) { c.Format, c.Theme = mtgv1.FormatId_FORMAT_ID_STANDARD, "burn" }},
		{want: []string{"power_sixty", "budget"}, fill: []string{"power", "budget"}},
	}
	cs = append(cs, c18)

	c19 := conversation{name: "a big library, no theme"}
	c19.ctx = newCtx("i have a big collection, build me something good")
	c19.ctx.HasCollection = true
	c19.steps = []step{
		{want: []string{"format", "theme", "colors"}, fill: []string{"format", "theme", "colors"},
			set: func(c *Context) { commander(c); c.Theme = "tokens" }},
		{want: []string{"commander", "power_commander", "pool"},
			fill: []string{"commander", "commander_pick", "power", "pool_rule"},
			set:  func(c *Context) { c.CommanderSet = true }},
	}
	cs = append(cs, c19)

	// "Casual" alone no longer routes to house rules (D-78). The gate run
	// of 2026-08-25 asked this user about house rules, and they answered
	// "casual means low power, not a house format".
	c20 := conversation{name: "dinosaur tribal for a child"}
	c20.ctx = newCtx("dinosaur deck for my kid, keep it casual")
	c20.ctx.HasCollection = true
	c20.steps = []step{
		{want: []string{"format", "theme", "colors"}, fill: []string{"format", "theme", "colors"},
			set: func(c *Context) { commander(c); c.Theme = "dinosaurs" }},
		{want: []string{"commander", "power_commander", "pool"},
			fill: []string{"commander", "commander_pick", "power", "pool_rule"},
			set:  func(c *Context) { c.CommanderSet = true }},
	}
	cs = append(cs, c20)

	c21 := conversation{name: "modern with proxies"}
	c21.ctx = newCtx("we proxy everything at our table")
	c21.steps = []step{
		{want: []string{"format", "theme", "colors"}, fill: []string{"format", "theme", "colors"},
			set: func(c *Context) {
				c.Format, c.Theme, c.HouseFormat = mtgv1.FormatId_FORMAT_ID_MODERN, "shops", true
				// The second message is what raises the house rules. The
				// word "proxy" no longer does, because a proxy user has no
				// budget and names no legality (D-111).
				c.Words += " any card, no ban list. call it vintage"
			}},
		{want: []string{"house_rules", "power_sixty"}, fill: []string{"house_rules", "power"}},
		{want: []string{"house_format_limits"}, fill: []string{"house_format_limits"}},
	}
	cs = append(cs, c21)

	c22 := conversation{name: "extra turns commander"}
	c22.ctx = newCtx("i want an extra turns deck")
	c22.ctx.HasCollection = true
	c22.steps = []step{
		{want: []string{"format", "theme", "colors"}, fill: []string{"format", "theme", "colors"},
			set: func(c *Context) { commander(c); c.Theme = "extra turns" }},
		{want: []string{"commander", "power_commander", "pool"},
			fill: []string{"commander", "commander_pick", "power", "pool_rule"},
			set:  func(c *Context) { c.CommanderSet = true }},
	}
	cs = append(cs, c22)

	c23 := conversation{name: "land destruction"}
	c23.ctx = newCtx("land destruction commander deck")
	c23.ctx.HasCollection, c23.ctx.Theme = true, "land destruction"
	c23.steps = []step{
		{want: []string{"format", "theme", "colors"}, fill: []string{"format", "theme", "colors"}, set: commander},
		{want: []string{"commander", "power_commander", "pool"},
			fill: []string{"commander", "commander_pick", "power", "pool_rule"},
			set:  func(c *Context) { c.CommanderSet = true }},
	}
	cs = append(cs, c23)

	c24 := conversation{name: "stax, owned only"}
	c24.ctx = newCtx("stax deck from only the cards i own")
	c24.ctx.HasCollection, c24.ctx.OwnedMode, c24.ctx.Theme = true, true, "stax"
	c24.steps = []step{
		{want: []string{"format", "theme", "colors"}, fill: []string{"format", "theme", "colors"}, set: commander},
		{want: []string{"commander", "power_commander", "pool"},
			fill: []string{"commander", "commander_pick", "power", "pool_rule"},
			set:  func(c *Context) { c.CommanderSet = true }},
	}
	cs = append(cs, c24)

	c25 := conversation{name: "poison in a tournament"}
	c25.ctx = newCtx("infect deck for a modern tournament")
	c25.ctx.PowerCompetitive, c25.ctx.BuyList = true, true
	c25.steps = []step{
		{want: []string{"format_store", "theme_competitive", "colors"}, fill: []string{"format", "theme", "colors"},
			set: func(c *Context) {
				c.Format, c.Theme = mtgv1.FormatId_FORMAT_ID_MODERN, "poison"
				c.PowerInferred = true
			}},
		{want: []string{"budget", "meta"}, fill: []string{"budget", "meta"}},
	}
	cs = append(cs, c25)

	c26 := conversation{name: "modern for an event"}
	c26.ctx = newCtx("modern deck for an event")
	c26.ctx.BuyList = true
	c26.steps = []step{
		{want: []string{"format_store", "theme", "colors"}, fill: []string{"format", "theme", "colors"},
			set: func(c *Context) { c.Format, c.Theme = mtgv1.FormatId_FORMAT_ID_MODERN, "delver" }},
		{want: []string{"power_sixty", "budget"}, fill: []string{"power", "budget"}},
	}
	cs = append(cs, c26)

	// The named card goes into the 99, so it stays a locked card (D-70),
	// and the commander row still has work to do.
	c27 := conversation{name: "a card for the 99"}
	c27.ctx = newCtx("build around grist but not as my commander")
	c27.ctx.HasCollection, c27.ctx.NamedCard, c27.ctx.LockedCard = true, true, true
	c27.steps = []step{
		{want: []string{"format", "theme_card_named", "colors"}, fill: []string{"format", "theme", "colors"},
			set: func(c *Context) { commander(c); c.Theme = "sacrifice" }},
		{want: []string{"named_card_role", "power_commander", "pool"},
			fill: []string{"named_card_role", "power", "pool_rule"}},
		{want: []string{"commander", "locked"}, fill: []string{"commander", "commander_pick", "locked"},
			set: func(c *Context) { c.CommanderSet = true }},
	}
	cs = append(cs, c27)

	c28 := conversation{name: "reanimator with two plans"}
	c28.ctx = newCtx("reanimator commander deck")
	c28.ctx.HasCollection, c28.ctx.TwoPlans, c28.ctx.Theme = true, true, "reanimator"
	c28.steps = []step{
		{want: []string{"format", "theme", "colors"}, fill: []string{"format", "theme", "colors"}, set: commander},
		{want: []string{"commander", "power_commander", "pool"},
			fill: []string{"commander", "commander_pick", "power", "pool_rule"},
			set:  func(c *Context) { c.CommanderSet = true }},
		{want: []string{"plan_choice"}, fill: []string{"plan_variant"}},
	}
	cs = append(cs, c28)

	c29 := conversation{name: "two hundred dollars"}
	c29.ctx = newCtx("commander deck, i can spend 200 dollars, i have a library")
	c29.ctx.HasCollection, c29.ctx.BudgetAmbiguous, c29.ctx.BuyList = true, true, true
	c29.ctx.Theme = "artifacts"
	c29.steps = []step{
		{want: []string{"format", "theme", "colors"}, fill: []string{"format", "theme", "colors"}, set: commander},
		{want: []string{"commander", "power_commander", "pool"},
			fill: []string{"commander", "commander_pick", "power", "pool_rule"},
			set:  func(c *Context) { c.CommanderSet = true }},
		{want: []string{"budget"}, fill: []string{"budget"}},
		{want: []string{"budget_scope"}, fill: []string{"budget_scope"}},
	}
	cs = append(cs, c29)

	c30 := conversation{name: "another version, same plan"}
	c30.ctx = newCtx("give me the same deck with other cards")
	c30.ctx.AfterBuild, c30.ctx.HasCollection, c30.ctx.CommanderSet = true, true, true
	c30.ctx.Format = mtgv1.FormatId_FORMAT_ID_COMMANDER
	for _, k := range []string{"format", "theme", "colors", "commander", "power", "pool_rule", "budget"} {
		c30.ctx.Filled[k] = true
	}
	c30.steps = []step{
		{want: nil},
	}
	cs = append(cs, c30)

	// D-99: a request for another game gets one question and no others.
	c31 := conversation{name: "a request we can not serve"}
	c31.ctx = newCtx("can you build me a yu-gi-oh deck")
	c31.ctx.OutOfScope = true
	c31.steps = []step{
		{want: []string{"out_of_scope"}, fill: []string{"scope"},
			set: func(c *Context) { c.OutOfScope = false }},
		{want: []string{"format", "theme", "colors"}, fill: []string{"format", "theme", "colors"},
			set: func(c *Context) { commander(c); c.Theme = "dragons" }},
		{want: []string{"commander", "power_commander"},
			fill: []string{"commander", "commander_pick", "power"},
			set:  func(c *Context) { c.CommanderSet = true }},
	}
	cs = append(cs, c31)

	// D-112: a request for two decks gets one question and no others. The
	// app builds one deck at a time, and it says so.
	c32 := conversation{name: "two decks at once"}
	c32.ctx = newCtx("i want two decks, one commander and one modern")
	c32.ctx.TwoDecks = true
	c32.steps = []step{
		{want: []string{"one_deck"}, fill: []string{"deck_count", "format", "theme"},
			set: func(c *Context) { commander(c); c.Theme = "dragons" }},
		{want: []string{"commander", "power_commander", "colors"},
			fill: []string{"commander", "commander_pick", "power", "colors"},
			set:  func(c *Context) { c.CommanderSet = true }},
	}
	cs = append(cs, c32)

	// D-112: a format this app does not build gets named, with the
	// nearest format it does build. Gate run 13 offered Brawl instead.
	c33 := conversation{name: "a format we do not build"}
	c33.ctx = newCtx("i want a brawl deck for arena")
	c33.ctx.UnsupportedFormat = true
	c33.steps = []step{
		{want: []string{"format_unsupported", "theme", "colors"},
			fill: []string{"format", "theme", "colors"},
			set: func(c *Context) {
				commander(c)
				c.Theme, c.UnsupportedFormat = "dragons", false
			}},
		{want: []string{"commander", "power_commander"},
			fill: []string{"commander", "commander_pick", "power"},
			set:  func(c *Context) { c.CommanderSet = true }},
	}
	cs = append(cs, c33)

	// D-146: Historic and Timeless name no nearest format. The pool
	// measurement of 2026-08-26 put Pioneer nearest for both, and not
	// Modern and Legacy, so the owner chose to offer no substitute.
	c33b := conversation{name: "an unsupported format with no substitute"}
	c33b.ctx = newCtx("i want a historic deck")
	c33b.ctx.UnsupportedFormat = true
	c33b.ctx.NoNearFormat = true
	c33b.steps = []step{
		{want: []string{"format_unsupported_open", "theme", "colors"},
			fill: []string{"format", "theme", "colors"},
			set: func(c *Context) {
				commander(c)
				c.Theme = "dragons"
				c.UnsupportedFormat, c.NoNearFormat = false, false
			}},
		{want: []string{"commander", "power_commander"},
			fill: []string{"commander", "commander_pick", "power"},
			set:  func(c *Context) { c.CommanderSet = true }},
	}
	cs = append(cs, c33b)

	// D-129: a card that can not lead a deck is not a commander. Probe 41
	// named Lightning Bolt, and every run accepted it in silence.
	c34 := conversation{name: "a commander that can not lead"}
	c34.ctx = newCtx("commander deck with lightning bolt as my commander")
	c34.ctx.Format, c34.ctx.Theme = mtgv1.FormatId_FORMAT_ID_COMMANDER, "burn"
	c34.ctx.CommanderIllegal = true
	// The card can only sit in the 99, so the role question is settled.
	for _, k := range []string{"format", "theme", "named_card_role"} {
		c34.ctx.Filled[k] = true
	}
	c34.steps = []step{
		{want: []string{"commander_illegal", "power_commander", "colors"},
			fill: []string{"commander_illegal", "power", "colors"},
			set: func(c *Context) {
				c.CommanderIllegal, c.Suggested = false, true
			}},
		{want: []string{"commander_pick"}, fill: []string{"commander", "commander_pick", "named_card_role"},
			set: func(c *Context) { c.CommanderSet = true }},
	}
	cs = append(cs, c34)

	return cs
}

// TestConversations is the deterministic half of the PR-7 gate: a complete
// slot set in at most four turns, with no repeated question. The other
// half of the gate (at least 25 of 30 conversations use catalog questions
// only) needs the model in the loop, because only the model invents a
// question. Every question below comes from the catalog.
// TestGateSize holds the count. The other half of the gate (at least 25
// of the 30 use catalog questions only) needs the model in the loop, and
// `cmd/questions-gate` runs it.
func TestGateSize(t *testing.T) {
	if n := len(conversations()); n < MinGateSize {
		t.Errorf("%d scripted conversations, the gate needs %d", n, MinGateSize)
	}
}

// TestEveryRowFires proves the gate reaches every catalog row. A row no
// conversation reaches is a row no test covers.
func TestEveryRowFires(t *testing.T) {
	c := load(t)
	fired := map[string]bool{}
	for _, conv := range conversations() {
		ctx := conv.ctx
		for _, st := range conv.steps {
			for _, r := range c.Plan(ctx) {
				fired[r.ID] = true
				ctx.Asked[r.ID] = true
			}
			for _, k := range st.fill {
				ctx.Filled[k] = true
			}
			if st.set != nil {
				st.set(&ctx)
			}
		}
	}
	for _, r := range c.Rows {
		if !fired[r.ID] {
			t.Errorf("no gate conversation reaches row %q", r.ID)
		}
	}
}

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
					row, _ := c.Row(id)
					if everAsked[id] && !row.Repeat {
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
