package candidates

import (
	"sort"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
	"github.com/nkramber/decktome/go/internal/cards"
	"github.com/nkramber/decktome/go/internal/profile"
)

// PowerRate is how a bracket 4 or 5 shortlist reads the top-list rate
// (D-704, D-707, D-710, D-712). Weight is the weight of the rate in the
// score. Keep is the rate at which a card with no theme signal stays on
// the list and keeps its full score. Pin lets a power card at the keep
// rate skip the cap of its role and add to the total. The zero value keeps
// the default, and any other value reads as given, with a weight or a keep
// rate of zero read as none.
type PowerRate struct {
	Weight float64
	Keep   float64
	Pin    bool
}

// DefaultPowerRate is the rate a bracket 4 or 5 shortlist reads. The free
// sweep of PR-45b set it: with the pin, the weight of PR-14B meets every
// power floor of the bracket gate prompts, and more weight only drops
// on-theme cards (D-711).
var DefaultPowerRate = PowerRate{Weight: rateWeight, Keep: 0.3, Pin: true}

func (r PowerRate) withDefaults() PowerRate {
	if r == (PowerRate{}) {
		return DefaultPowerRate
	}
	if r.Weight <= 0 {
		r.Weight = rateWeight
	}
	if r.Keep <= 0 {
		r.Keep = noKeep
	}
	return r
}

// rateWeight is the weight of the top-list rate in the score of every
// other request (PR-14B).
const rateWeight = 0.1

// noKeep is a keep rate no card reaches, because a rate is at most 1.
const noKeep = 2

// powerRequest reports a request whose bracket holds power floors: a
// Commander deck at bracket 4 or 5 (D-704).
func powerRequest(req Request) bool {
	return req.Format == mtgv1.FormatId_FORMAT_ID_COMMANDER && req.Bracket >= 4
}

// powerScan is how one request reads power: whether its bracket holds
// power floors, the weight of the rate, the keep rate, and the reader of
// the floors a card counts toward (D-704, D-707, D-709).
type powerScan struct {
	on     bool
	weight float64
	rate   PowerRate
	of     func(*mtgv1.Card) []string
}

// powerScanOf reads the power of a request. A request below bracket 4
// reads the rate at the weight of PR-14B and keeps no card by it.
func powerScanOf(idx *cards.Index, req Request) powerScan {
	if !powerRequest(req) {
		return powerScan{weight: rateWeight}
	}
	rate := req.PowerRate.withDefaults()
	return powerScan{on: true, weight: rate.Weight, rate: rate, of: profile.PowerOf(idx.Tags())}
}

// FinisherTarget is the finisher count a Commander bracket asks for
// (D-726). Brackets 1 to 4 read the median of the precons since 2023 and
// of the EDHREC average decks. Bracket 5 reads the median of the TopDeck
// top cut, where one combo finisher closes the game.
func FinisherTarget(bracket int32) int {
	if bracket >= 5 {
		return 1
	}
	return 3
}

// promoteFinishers gives the role wincon to the best finishers of a
// Commander shortlist, up to the target of the bracket, and it pins each
// one (D-726, D-741). Every other finisher keeps the role it earned, so
// the wincon cap of 15 drops no card. The list holds the cards it held
// before, and the target count of them now reads wincon.
//
// The promotion stops at the target on purpose. A pinned card skips the
// cap of its role and adds to the total (F-132, D-712), and the snapshot
// of 2026-09-04 holds 1,857 finishers.
//
// An owned mode keeps the owned cards and drops the rest after this
// step, so the owned finishers take the role first (D-742). Without that
// order a promotion lands on a card the mode drops, and the shortlist
// reads under the target of its bracket.
func promoteFinishers(cs []Candidate, req Request, mode mtgv1.PoolRule, target int, finishers map[string]bool) {
	if req.Format != mtgv1.FormatId_FORMAT_ID_COMMANDER || target <= 0 || len(finishers) == 0 {
		return
	}
	n := 0
	if mode != mtgv1.PoolRule_POOL_RULE_ANY_CARD {
		n = promoteUpTo(cs, target, n, finishers, func(c Candidate) bool { return c.Owned > 0 })
	}
	// Owned-only drops every unowned card, so it takes no second pass.
	if mode == mtgv1.PoolRule_POOL_RULE_OWNED_ONLY {
		return
	}
	promoteUpTo(cs, target, n, finishers, func(Candidate) bool { return true })
}

// promoteUpTo gives the role wincon to each finisher that want reads,
// from the best down, until the list holds target of them. It answers
// the new count.
func promoteUpTo(cs []Candidate, target, n int, finishers map[string]bool, want func(Candidate) bool) int {
	for i := range cs {
		if n >= target {
			return n
		}
		if cs[i].Role == mtgv1.CardRole_CARD_ROLE_WINCON && cs[i].Pinned {
			continue
		}
		if finishers[cs[i].Card.GetOracleId()] && want(cs[i]) {
			cs[i].Role = mtgv1.CardRole_CARD_ROLE_WINCON
			cs[i].Pinned = true
			n++
		}
	}
	return n
}

// pinPower pins each power card whose rate reaches the keep rate, when the
// request pins (D-710). A pinned card skips the cap of its role.
func pinPower(cs []Candidate, pw powerScan) {
	if !pw.on || !pw.rate.Pin {
		return
	}
	for i := range cs {
		if cs[i].Rate >= pw.rate.Keep && len(pw.of(cs[i].Card)) > 0 {
			cs[i].Pinned = true
		}
	}
}

// capRole keeps the n best cards of one role, in score order. A pinned
// power card skips the cap and does not count against it (D-710).
func capRole(cs []Candidate, n int) []Candidate {
	out := make([]Candidate, 0, n)
	for _, c := range cs {
		switch {
		case c.Pinned:
			out = append(out, c)
		case n > 0:
			out = append(out, c)
			n--
		}
	}
	return out
}

// capPinnedLands is capLands with the pinned lands set aside. A Game
// Changer land counts toward a power floor, so a pinned land skips the
// land cap and does not count against it, as a pinned card of every other
// role does (F-131, D-710). The result keeps the order of the bucket.
func capPinnedLands(cs []Candidate, n int) []Candidate {
	rest := make([]Candidate, 0, len(cs))
	for _, c := range cs {
		if !c.Pinned {
			rest = append(rest, c)
		}
	}
	keep := make(map[*mtgv1.Card]bool, n)
	for _, c := range capLands(rest, n) {
		keep[c.Card] = true
	}
	out := make([]Candidate, 0, n)
	for _, c := range cs {
		if c.Pinned || keep[c.Card] {
			out = append(out, c)
		}
	}
	return out
}

// reservePerFloor is how many reserve cards each power floor keeps. The
// highest floor is 8 Game Changers, so a deck that holds none still
// reads a full answer (D-709).
const reservePerFloor = 12

// topReserve keeps the highest-rate cards of each power floor, by rate,
// then by play, then by name (D-709). A card with no rate is no answer,
// and a card that counts toward two floors is kept once.
func topReserve(cs []Candidate, of func(*mtgv1.Card) []string) []Candidate {
	sort.SliceStable(cs, func(i, j int) bool {
		if cs[i].Rate != cs[j].Rate {
			return cs[i].Rate > cs[j].Rate
		}
		if cs[i].Pop != cs[j].Pop {
			return cs[i].Pop > cs[j].Pop
		}
		return cs[i].Card.GetName() < cs[j].Card.GetName()
	})
	taken := map[string]int{}
	var out []Candidate
	for _, c := range cs {
		if c.Rate <= 0 {
			break
		}
		keys := of(c.Card)
		keep := false
		for _, key := range keys {
			keep = keep || taken[key] < reservePerFloor
		}
		if !keep {
			continue
		}
		for _, key := range keys {
			taken[key]++
		}
		out = append(out, c)
	}
	return out
}
