package collections

import (
	"io"
	"strings"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
)

// The Moxfield reader (PR-34, D-647, F-91).
//
// Every fact below comes from a real export, `testdata/moxfield_collection.csv`,
// written by Moxfield on 2026-09-10 and read the same day. The
// documentation of the format disagrees with itself on two points, and
// the file settles both.
//
//   - `Edition` holds a lowercase set code, "fin" and "msh", and never
//     a set name. One guide describes the set name, and that guide
//     describes the import format and not the export.
//   - `Condition` holds the long name, "Near Mint", and never "NM".
//
// The header of that file is:
//
//	Count, Tradelist Count, Name, Edition, Condition, Language, Foil,
//	Tags, Last Modified, Collector Number, Alter, Proxy, Purchase Price
//
// Four columns are read and the rest are ignored. `Tradelist Count` is
// what the reader will trade away and not what they own. `Alter` and
// `Proxy` read False on every row of the export, so no row of this repo
// has ever shown what a True does.

// moxfieldColumns are the key columns that identify the format. Count,
// Name, and Edition are the three every export writes, and no other
// format of this app writes all three.
var moxfieldColumns = [][]string{
	{"Count", "Name", "Edition"},
}

// moxfieldFinish maps the Foil cell. The export writes an empty cell,
// "foil", or "etched", and the fixture holds 2526, 663, and 4 of them.
var moxfieldFinish = map[string]mtgv1.Finish{
	"":       mtgv1.Finish_FINISH_NORMAL,
	"normal": mtgv1.Finish_FINISH_NORMAL,
	"foil":   mtgv1.Finish_FINISH_FOIL,
	"etched": mtgv1.Finish_FINISH_ETCHED,
}

// moxfieldCondition maps the Condition cell. "Near Mint" and "Played"
// are verified against the export. The rest are the other rungs of the
// Moxfield condition list, and no export of this repo has shown one, so
// they are unverified. An unknown value is reported and never defaulted
// in silence.
var moxfieldCondition = map[string]mtgv1.Condition{
	"":                      mtgv1.Condition_CONDITION_NEAR_MINT,
	"mint":                  mtgv1.Condition_CONDITION_MINT,
	"near mint":             mtgv1.Condition_CONDITION_NEAR_MINT,
	"excellent":             mtgv1.Condition_CONDITION_EXCELLENT,
	"good":                  mtgv1.Condition_CONDITION_GOOD,
	"good (lightly played)": mtgv1.Condition_CONDITION_LIGHT_PLAYED,
	"lightly played":        mtgv1.Condition_CONDITION_LIGHT_PLAYED,
	"played":                mtgv1.Condition_CONDITION_PLAYED,
	"heavily played":        mtgv1.Condition_CONDITION_POOR,
	"damaged":               mtgv1.Condition_CONDITION_POOR,
	"poor":                  mtgv1.Condition_CONDITION_POOR,
}

// moxfieldLanguage maps the Language cell onto the code vocabulary the
// rest of this package reads.
//
// This mapping is not a convenience. Resolve refuses every row whose
// language is not "en" (D-23), and Moxfield writes "English". A file
// that kept the long name would lose all 3192 English rows of the
// fixture as non-English.
//
// "English" and "Japanese" are verified against the export. The rest
// are the other languages the Moxfield format offers, and the codes are
// the Scryfall codes the card index uses.
var moxfieldLanguage = map[string]string{
	"english":             "en",
	"spanish":             "es",
	"french":              "fr",
	"german":              "de",
	"italian":             "it",
	"portuguese":          "pt",
	"japanese":            "ja",
	"korean":              "ko",
	"russian":             "ru",
	"simplified chinese":  "zhs",
	"traditional chinese": "zht",
}

// ParseMoxfieldCSV reads a Moxfield collection export. It returns the
// parsed rows and the rows it could not parse. It fails only on a
// broken header.
func ParseMoxfieldCSV(r io.Reader) ([]Row, []*mtgv1.UnresolvedRow, error) {
	return csvRows(r, "moxfield", moxfieldColumns, func(get func(string) string) (Row, mtgv1.UnresolvedReason) {
		qty, ok := parseQuantity(get("Count"))
		if !ok {
			return Row{}, mtgv1.UnresolvedReason_UNRESOLVED_REASON_BAD_ROW
		}
		finish, ok := moxfieldFinish[strings.ToLower(get("Foil"))]
		if !ok {
			return Row{}, mtgv1.UnresolvedReason_UNRESOLVED_REASON_UNKNOWN_VALUE
		}
		condition, ok := moxfieldCondition[strings.ToLower(get("Condition"))]
		if !ok {
			return Row{}, mtgv1.UnresolvedReason_UNRESOLVED_REASON_UNKNOWN_VALUE
		}
		language, ok := moxfieldLanguageOf(get("Language"))
		if !ok {
			return Row{}, mtgv1.UnresolvedReason_UNRESOLVED_REASON_UNKNOWN_VALUE
		}
		return Row{
			Name:      get("Name"),
			SetCode:   strings.ToLower(get("Edition")),
			Collector: get("Collector Number"),
			Quantity:  qty,
			Finish:    finish,
			Condition: condition,
			Language:  language,
		}, mtgv1.UnresolvedReason_UNRESOLVED_REASON_UNSPECIFIED
	})
}

// moxfieldLanguageOf reads a Language cell as a code. An empty cell
// reads English, which is what the export writes for a card with no
// language of its own.
func moxfieldLanguageOf(v string) (string, bool) {
	if v == "" {
		return "en", true
	}
	code, ok := moxfieldLanguage[strings.ToLower(v)]
	return code, ok
}
