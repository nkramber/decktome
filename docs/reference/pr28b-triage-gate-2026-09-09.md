# PR-28b gate, 2026-09-09

The roadmap gate of PR-28b has two items.

1. A dry triage over a fixture of ten items, the owner's own past complaints, names the expected class of each.
2. The question gate reads a "must not ask" expectation and fails a conversation that fires the row.

**Verdict: PASS.** Both items hold. Every step is free: no lane of this gate calls a model.

## Item 1: the dry triage over ten items

`go run ./cmd/feedback-triage -root .. -in internal/triage/testdata/fixture.jsonl -dry` reads the fixture and routes every verdict. The run reads PASS and it calls no model.

The fixture is `go/internal/triage/testdata/fixture.jsonl`. Each item carries the shape the dialog of PR-27 sends: a kind, a verdict, the reason keys the reader checked, the reader's own text, and the snapshot of the object it names (D-635). The ten cover every artifact writer and every routing path.

| # | Kind and reason | Expected class | Route | Artifact |
|---|---|---|---|---|
| f01 | question / already_answered | Q1, asked again | the reason key | conversation, with `must_not_ask` |
| f02 | question / bad_options | Q3, bad options | the reason key | conversation |
| f03 | summary / false_claim | S1, false summary | the reason key | summary case |
| f04 | card / off_theme | C1, off theme | the reason key | deck prompt, with `must_not_include` |
| f05 | card / illegal | C2, illegal | the reason key | defect row, no case |
| f06 | card / unwanted_buy, with text | C3, unwanted buy | the judge, it meets D-37 | none until the judge answers |
| f07 | deck / wrong_power, no text | D3, wrong power | the reason key, it meets D-459 | bracket prompt |
| f08 | deck / bad_mana and too_little_interaction | two classes | the judge, the reasons cross | none until the judge answers |
| f09 | chat / stuck | X1, the chat stopped | the reason key | defect row, no case |
| f10 | deck / thumbs up | keep case | the verdict word | deck prompt, no assertion |

The run answers every row. `TestTheDryTriageNamesTheClassOfEveryFixtureItem` reads the same fixture and the same table, so `make verify` measures item 1 on every pull request.

Two items want the judge, and eight do not. The two are the design of D-643: a reader who argues with a rule the owner set, and a reader who checked reasons of two classes. Eight of ten cost nothing.

### What the run wrote

`f01`, the headline artifact, reads:

```json
{
  "id": 110,
  "name": "Q1: asked again",
  "collection": true,
  "messages": [
    "Build me a lifegain deck from my library.",
    "Commander, white and black.",
    "Bracket 3, and Karlov of the Ghost Council."
  ],
  "note": "From a reader's thumbs down, 2026-09-08. asked again. Reasons: already_answered. The reader wrote: I already told you the bracket in my first message.",
  "expect": {
    "budget": "50 to buy",
    "colors": "WB",
    "format": "commander",
    "pool_rule": "owned_first",
    "power": "bracket 3",
    "theme": "lifegain"
  },
  "must_not_ask": ["power_commander"]
}
```

The row comes from the question id, `q3-power_commander`, and no model call reads it. The expectation comes from the slots of the stored session through `questions.SlotWords`, which is the renderer the gate reads, so a case can not name a word the gate does not read.

### The gaps the run named

The fixture uses invented Oracle ids, and no card index holds them. Every case named that gap, and the document printed it. That is the design: a case says what it could not fill, and a person finishes it on the pull request.

## Item 2: the "must not ask" expectation

`TestAForbiddenRowFailsTheGate` in `go/cmd/questions-gate/main_test.go` runs the gate document writer over one conversation that names `power_commander` as a row it must never ask.

- The row fires: the verdict reads FAIL, the document names `must_not_ask: power_commander fired at turn 1`, and the run records the bar `109/never_asked_power_commander` at zero.
- The row does not fire: the verdict reads PASS, and the same bar reads one with the detail "the row never fired".

The control matters. A bar that has never passed and a bar that has never failed are both untested (D-234), so this test fires both sides.

Three more tests hold the expectation:

- `TestAForbiddenRowThatFiresIsAMiss` reads the check itself. A row that fires twice is one miss, and the miss names the first turn.
- `TestAForbiddenRowMustExist` refuses a row id the catalog does not hold. A typo makes an expectation that can never fail, and the gate would spend a run to learn nothing.
- `TestEveryForbiddenRowOfTheGateFileExists` reads the shipped `conversations.json`, so a bad row id fails a free test and never a paid run.

## The other bars this change adds

The deck gate reads two assertions of a case prompt. `TestACaseAssertionFailsTheRun` and `TestTheOwnershipAssertionFailsTheRun` fire both, and `TestAPromptWithNoAssertionReadsNone` proves the 25 golden prompts are untouched.

`TestTheClassKeysMatchTheStore` reads the reason keys of the store against the class table, in both directions. A key the dialog gains later fails a free test, which is how F-90 stays fixed.

`TestTheHarvestNeverReadsTheUserRecord` covers the triage now. No package of the feedback loop may import the user record, because that record holds the reader's email (D-559, D-638).

`TestADeckVerdictKeepsTheSessionThatBuiltIt` reads the store change of D-643, and `TestTwoSnapshotsStayInsideOneDocument` reads the guard under it: two snapshots that pass the room of one Firestore document drop the session and keep the deck.

## Cost

Nothing. The dry triage calls no model, and every test above runs under `make verify`.

A live triage costs a judge call for each verdict the reason keys can not place. On this fixture that is 2 of 10.

## What this gate does not measure

The live judge lane. No run of it exists yet, because the three real verdicts on record predate the snapshot of D-635 and carry no context. The judge prompt and its schema are in `go/internal/triage/judge.go`, and `make feedback-triage TRIAGE_OUT=<document>` runs the lane. The first live run belongs with the first real harvest that carries a snapshot.
