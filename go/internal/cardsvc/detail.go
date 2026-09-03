package cardsvc

import (
	"context"
	"errors"
	"sort"

	"connectrpc.com/connect"
	"google.golang.org/protobuf/proto"

	mtgv1 "github.com/nkramber/mtg-deck-builder/go/gen/mtg/v1"
)

var errUnknownCard = errors.New("the card database holds no card with this oracle_id")

// GetRulings answers the rulings of one card, oldest first, with the
// snapshot date (PR-20). A card the index does not know is NotFound. A
// snapshot with no rulings file answers an empty list and has_rulings
// false, so the client can say "no data" and not "no ruling".
func (s *Server) GetRulings(_ context.Context, req *connect.Request[mtgv1.GetRulingsRequest]) (*connect.Response[mtgv1.GetRulingsResponse], error) {
	idx, err := s.ready()
	if err != nil {
		return nil, err
	}
	id := req.Msg.GetOracleId()
	if _, ok := idx.ByOracleID(id); !ok {
		return nil, connect.NewError(connect.CodeNotFound, errUnknownCard)
	}
	res := &mtgv1.GetRulingsResponse{AsOf: idx.AsOf.UTC().Format("2006-01-02"), HasRulings: idx.HasRulings()}
	for _, r := range idx.Rulings(id) {
		res.Rulings = append(res.Rulings, &mtgv1.Ruling{PublishedAt: r.PublishedAt, Comment: r.Comment, Source: r.Source})
	}
	return connect.NewResponse(res), nil
}

// GetPrintings answers every playable printing of one card with its
// price, newest set first and then by collector number (PR-20). The
// index rows are shared, so the answer carries copies.
func (s *Server) GetPrintings(_ context.Context, req *connect.Request[mtgv1.GetPrintingsRequest]) (*connect.Response[mtgv1.GetPrintingsResponse], error) {
	idx, err := s.ready()
	if err != nil {
		return nil, err
	}
	id := req.Msg.GetOracleId()
	if _, ok := idx.ByOracleID(id); !ok {
		return nil, connect.NewError(connect.CodeNotFound, errUnknownCard)
	}
	rows := idx.PrintingsOf(id)
	out := make([]*mtgv1.Printing, 0, len(rows))
	for _, p := range rows {
		out = append(out, proto.Clone(p).(*mtgv1.Printing))
	}
	sets := idx.Sets()
	released := func(code string) string {
		if info, ok := sets.Get(code); ok {
			return info.ReleasedAt
		}
		return ""
	}
	sort.SliceStable(out, func(i, j int) bool {
		ri, rj := released(out[i].GetSetCode()), released(out[j].GetSetCode())
		if ri != rj {
			return ri > rj
		}
		if out[i].GetSetCode() != out[j].GetSetCode() {
			return out[i].GetSetCode() < out[j].GetSetCode()
		}
		return collectorLess(out[i].GetCollectorNumber(), out[j].GetCollectorNumber())
	})
	return connect.NewResponse(&mtgv1.GetPrintingsResponse{Printings: out, PriceAsOf: idx.AsOf.UTC().Format("2006-01-02")}), nil
}

// collectorLess orders collector numbers by their leading number, so
// "10" follows "9", and by the text after it.
func collectorLess(a, b string) bool {
	na, ra := leadingNumber(a)
	nb, rb := leadingNumber(b)
	if na != nb {
		return na < nb
	}
	return ra < rb
}

func leadingNumber(s string) (int, string) {
	n := 0
	i := 0
	for i < len(s) && s[i] >= '0' && s[i] <= '9' {
		n = n*10 + int(s[i]-'0')
		i++
	}
	return n, s[i:]
}
