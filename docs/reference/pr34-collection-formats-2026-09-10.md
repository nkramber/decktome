# PR-34 gate, 2026-09-10

The roadmap gate of PR-34 reads:

> One fixture per platform, parsed to the same row shape the ManaBox lane produces. A file of an unknown shape reports every row it failed to read, and it drops none in silence. The upload screen names every format it takes.

**Verdict: PASS**, for the one platform the owner named. Moxfield reads, on a real export.

Every step of this gate is free. No lane of it calls a model.

## The fixture is a real export

`go/internal/collections/testdata/moxfield_collection.csv` is the owner's own Moxfield collection, exported 2026-09-10 and read the same day. The roadmap item asked for a real export and not a column list, and that was the right call: **the documentation of this format disagrees with itself on two points, and both of them decide whether a row resolves.**

| Point | What the guides said | What the export holds |
|---|---|---|
| `Edition` | A set name, "Ikoria: Lair of Behemoths" | A lowercase set code, `fin` and `msh`. 125 distinct, none longer than four characters, none with a space. |
| `Condition` | `NM`, `SP`, `MP`, `HP` | The long name, `Near Mint` and `Played`. |

The guide that describes a set name and the short condition describes the Moxfield **import** format. The app reads the **export**, and the two are not the same file.

The header of the real export:

```
"Count","Tradelist Count","Name","Edition","Condition","Language","Foil",
"Tags","Last Modified","Collector Number","Alter","Proxy","Purchase Price"
```

## The fault that would have lost every card

`Resolve` refuses a row whose language is not `en` (D-23). Moxfield writes `English`.

A parser that passed the cell through unchanged would have reported **all 3192 English cards of that export as non-English**, and the reader would have read that their whole collection failed. The parser maps the name onto the code, and `TestTheMoxfieldLanguageIsACode` reads the counts back off the real file.

This is the clearest argument for the real-export rule. No column list would have shown it.

## The measurement

`ParseMoxfieldCSV` over the real export:

| Measure | Value |
|---|---|
| Rows under the header | 3193 |
| Rows parsed | 3193 |
| Rows unread | 0 |
| Distinct sets | 125 |
| Normal, foil, etched | 2526, 663, 4 |
| Near mint, played | 3192, 1 |
| English, Japanese | 3192, 1 |

`Resolve` over those rows, against the card snapshot of 2026-09-04:

| Measure | Value |
|---|---|
| Entries | 3035 |
| Cards owned | 5884 |
| Unresolved | 1 |

The one unresolved row is a Japanese printing of Helpful Hunter, reported as `UNRESOLVED_REASON_NON_ENGLISH`. D-23 refuses a non-English row by design, so that is the app working and not a fault of this parser.

Rows fall from 3193 to 3035 entries because the resolver merges rows that name one printing. A reader who holds the same card in two binders writes two rows.

## F-93, the gap a resolution count hid

3192 of 3193 rows resolve, and that number hides a fault.

Moxfield writes no Scryfall id column, so a row of it resolves on the set code and the collector number. `buildEntry` filled the entry's id from the row, and the row had none. **Every entry of a Moxfield collection carried an empty printing id.**

Three readers of that field then did nothing, in silence.

- The binder reads the printing for the art and the price of what the reader owns (D-299, F-60). It fell back to the default printing for all 3035 entries.
- `OwnedPrintings` drops an entry with no id. It covered 0 oracle ids of the collection.
- The binder tile keys itself on `scryfallId-finish-condition`. Every normal near-mint tile of the collection shared one key.

`Index.PrintingBySetCollector` answers the printing the pair names, and the entry takes it. It holds no index of its own: it walks the printings of the one card, and a card holds 225 at the most.

Measured over the real export, after the fix:

| Measure | Before | After |
|---|---|---|
| Entries with a printing id | 0 | 3035 |
| Printings with art | 0 | 3009 |
| Printings with a price | 0 | 3035 |
| Oracle ids `OwnedPrintings` covers | 0 | 2232 |

The 26 entries whose printing carries no art are two-faced cards, which hold their art on the faces. `collectionsvc` already falls back for those (F-60).

The ManaBox path is untouched: its rows carry a Scryfall id, so the new branch never fires for them.

## The detection

`collections.Detect` reads the header row and names the format. The reader drops a file and never says which app wrote it (D-647).

| Test | What it holds |
|---|---|
| `TestDetectNamesTheFormat` | Eight real shapes: three ManaBox headers, a byte order mark, three Arena lists, and a Moxfield header. Every detected format then parses, so the detector can never name a format the parser refuses. |
| `TestDetectRefusesAFileItDoesNotKnow` | A spreadsheet of something else, an empty file, and prose. The refusal names every format the app does read. |
| `TestDetectNeverGuessesBetweenTwoFormats` | A second format that claims the ManaBox columns. The tie is refused with both names. |
| `TestTheDetectorAndTheParserReadOneTable` | Every key column set of the signature table detects, and the parser takes it. |
| `TestTheRealExportDetectsAsMoxfield` | The real Moxfield export detects as Moxfield, and the real ManaBox export still detects as ManaBox. |

## F-92, the fault detection uncovered

`upload-dialog.tsx` sent `IMPORT_SOURCE_MANABOX_CSV` on every upload, hardcoded, and the reader never chose. The app has read an Arena list since D-15, and **no reader could ever upload one**: the ManaBox parser read the header, found no key column, and refused the whole file.

The upload names no format now. `TestImportDetectsTheFormat` uploads an Arena list with no source and reads a collection back, and `TestANamedSourceBeatsTheDetector` holds the other half: a caller that names a format means it (D-591).

## The shape a new platform takes

One CSV walker serves every format. `csvRows` reads the header, checks the key columns, walks the records, and reports what it could not read. A platform brings two things and never a copy of that loop:

1. One entry in the signature table of `detect.go`.
2. One row builder, which reads the cells of that format into a `Row`.

The Moxfield reader is 134 lines, and most of them are the three value vocabularies.

## What this gate does not measure

**The other five platforms.** The owner named Moxfield alone for the first pass. Archidekt, Deckbox, Delver Lens, TCGplayer, and Helvault wait for the owner's word and a real export each.

**The unverified condition rungs.** The export holds `Near Mint` and `Played`. The other rungs of the Moxfield condition list are mapped from its documentation and no export of this repo has shown one. `moxfield.go` marks them as unverified, and an unknown value is reported and never defaulted in silence.

**`Alter` and `Proxy`.** Both read `False` on every row of the export, so no file of this repo has shown what a `True` does. The parser ignores both columns.

`Proxy` is the one that matters, and it is **OQ-80** now. No other format this app reads carries such a column, so a proxy is information the app has never held. A proxy sits in the binder and plays at a table that allows one, and the reader never bought the card. So owned-only either builds with it or refuses it, and owned-first either leaves it off the buy list or prices it. The same card is owned for play and unowned for money, and the app holds one meaning of owned for both.

The owner parked it on 2026-09-10 rather than decide it on no evidence. The parser keeps dropping the column, so a later answer needs a re-upload from any reader who holds a proxy.

`Alter` is a smaller thing: an altered card is the card, with other art. The ManaBox format carries an `Altered` column and this app has always ignored it.
