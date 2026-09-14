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
