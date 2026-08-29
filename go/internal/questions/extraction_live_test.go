package questions

import (
	"context"
	"encoding/json"
	"io"
	"log/slog"
	"os"
	"testing"

	"github.com/nkramber/mtg-deck-builder/go/internal/llm"
)

// TestLiveFormatExtraction guards D-92, which no offline test can catch.
// A schema enum on the format field suppresses it: the model answers
// "unknown" for a message that names the format outright, and it fills
// every free-text field around it. A free string extracts what the enum
// misses.
//
// The failure is invisible offline, because a fake provider answers
// whatever the test tells it to. Only a real call shows it. Run this
// after any change to the classify schema or the format rule.
//
//	set -a && . ./.env && set +a && QUESTIONS_LIVE=1 \
//	  go test ./internal/questions -run TestLiveFormatExtraction -v -count=1
//
// TestLiveDecline guards D-93. A decline is the only way a user hands a
// choice back, and the schema must carry a field for it. An offline
// test can not prove the model emits the
// field, so this one asks the real provider.
func TestLiveDecline(t *testing.T) {
	if os.Getenv("QUESTIONS_LIVE") != "1" {
		t.Skip("set QUESTIONS_LIVE=1 and the API keys to run the live decline check")
	}
	env := func(k string) string {
		if k == llm.EnvRequireKeys {
			return "1"
		}
		return os.Getenv(k)
	}
	client, err := llm.NewFromEnv(env, slog.New(slog.NewTextHandler(io.Discard, nil)))
	if err != nil {
		t.Fatalf("client: %v", err)
	}
	// The color and bracket questions are both out.
	st := NewState(false)
	st.Slots.Theme = "sacrifice"
	st.Close("theme")
	st.Close("format")
	st.MarkAsked("colors", "colors", "colors")
	st.MarkAsked("power_commander", "power", "power")

	cases := []struct {
		message  string
		declined []string
	}{
		{"Any colors are fine.", []string{"colors"}},
		{"You decide.", []string{"colors", "power"}},
		{"Bracket 3.", nil},
		{"White and black.", nil},
	}
	// A decline must reach the key the user named, and no other.
	exact := map[string]int{"Any colors are fine.": 1, "You decide.": 2}
	for _, tc := range cases {
		t.Run(tc.message, func(t *testing.T) {
			input, err := json.Marshal(map[string]any{
				"message":     tc.message,
				"slots_known": st.Slots.SlotStates,
				"theme":       st.Slots.Theme,
				"open_keys":   openKeys(st),
			})
			if err != nil {
				t.Fatal(err)
			}
			res, err := client.Complete(context.Background(), llm.RoleClassify, llm.Request{
				Instructions: classifyInstructions,
				Input:        string(input),
				SchemaName:   "slot_fill",
				Schema:       json.RawMessage(classifySchema),
			}, nil)
			if err != nil {
				t.Fatalf("classify: %v", err)
			}
			var out classifyOut
			if err := json.Unmarshal(res.Output, &out); err != nil {
				t.Fatalf("output: %v", err)
			}
			t.Logf("declined=%v closed=%v colors=%v power=%q", out.DeclinedKeys, out.ClosedKeys, out.Colors, out.Power)
			for _, want := range tc.declined {
				var found bool
				for _, got := range out.DeclinedKeys {
					if got == want {
						found = true
					}
				}
				if !found {
					t.Errorf("declined_keys = %v, want it to hold %q", out.DeclinedKeys, want)
				}
			}
			if tc.declined == nil && len(out.DeclinedKeys) > 0 {
				t.Errorf("a real answer was read as a decline: %v", out.DeclinedKeys)
			}
			if want, ok := exact[tc.message]; ok && len(out.DeclinedKeys) != want {
				t.Errorf("declined %v, want %d key(s): a decline must not reach a slot the user did not mention", out.DeclinedKeys, want)
			}
		})
	}
}

// TestLiveOutOfScope guards D-99. "Can you build me a Yu-Gi-Oh deck?"
// must get a decline and not "Which Yu-Gi-Oh format would you like?".
// The row is useless unless the classifier raises the fact, and only a
// real call proves that.
func TestLiveOutOfScope(t *testing.T) {
	if os.Getenv("QUESTIONS_LIVE") != "1" {
		t.Skip("set QUESTIONS_LIVE=1 and the API keys to run the live scope check")
	}
	env := func(k string) string {
		if k == llm.EnvRequireKeys {
			return "1"
		}
		return os.Getenv(k)
	}
	client, err := llm.NewFromEnv(env, slog.New(slog.NewTextHandler(io.Discard, nil)))
	if err != nil {
		t.Fatalf("client: %v", err)
	}
	st := NewState(false)
	cases := []struct {
		message string
		want    bool
	}{
		{"Can you build me a Yu-Gi-Oh deck?", true},
		{"Build me a Pokemon deck for my league.", true},
		{"Build me a lifegain Commander deck.", false},
		{"I want a Modern burn deck.", false},
		// A Magic request that names another game is still in scope.
		{"A Magic deck that feels like a Yu-Gi-Oh deck.", false},
	}
	for _, tc := range cases {
		t.Run(tc.message, func(t *testing.T) {
			input, err := json.Marshal(map[string]any{
				"message":     tc.message,
				"slots_known": st.Slots.SlotStates,
				"theme":       st.Slots.Theme,
				"open_keys":   openKeys(st),
			})
			if err != nil {
				t.Fatal(err)
			}
			res, err := client.Complete(context.Background(), llm.RoleClassify, llm.Request{
				Instructions: classifyInstructions,
				Input:        string(input),
				SchemaName:   "slot_fill",
				Schema:       json.RawMessage(classifySchema),
			}, nil)
			if err != nil {
				t.Fatalf("classify: %v", err)
			}
			var out classifyOut
			if err := json.Unmarshal(res.Output, &out); err != nil {
				t.Fatalf("output: %v", err)
			}
			if out.Facts.OutOfScope != tc.want {
				t.Errorf("out_of_scope = %v, want %v", out.Facts.OutOfScope, tc.want)
			}
		})
	}
}

func TestLiveFormatExtraction(t *testing.T) {
	if os.Getenv("QUESTIONS_LIVE") != "1" {
		t.Skip("set QUESTIONS_LIVE=1 and the API keys to run the live extraction check")
	}
	env := func(k string) string {
		if k == llm.EnvRequireKeys {
			return "1"
		}
		return os.Getenv(k)
	}
	client, err := llm.NewFromEnv(env, slog.New(slog.NewTextHandler(io.Discard, nil)))
	if err != nil {
		t.Fatalf("client: %v", err)
	}
	// The state a real turn 2 carries: the format question is out, and
	// the theme and colors are already answered.
	st := NewState(true)
	st.Slots.Theme = "lifegain"
	st.Close("theme")
	st.Close("colors")
	st.MarkAsked("format", "format", "format")

	cases := []struct{ message, want string }{
		{"Commander, and white and black is right.", "commander"},
		{"I want a Modern burn deck.", "modern"},
		{"The format is Modern.", "modern"},
		{"Pauper. Burn, mono red.", "pauper"},
		// The negative matters as much: no format named, no format set.
		{"Build me a lifegain deck. I have a collection.", ""},
	}
	for _, tc := range cases {
		t.Run(tc.message, func(t *testing.T) {
			input, err := json.Marshal(map[string]any{
				"message":     tc.message,
				"slots_known": st.Slots.SlotStates,
				"theme":       st.Slots.Theme,
				"open_keys":   openKeys(st),
			})
			if err != nil {
				t.Fatal(err)
			}
			res, err := client.Complete(context.Background(), llm.RoleClassify, llm.Request{
				Instructions: classifyInstructions,
				Input:        string(input),
				SchemaName:   "slot_fill",
				Schema:       json.RawMessage(classifySchema),
			}, nil)
			if err != nil {
				t.Fatalf("classify: %v", err)
			}
			var out classifyOut
			if err := json.Unmarshal(res.Output, &out); err != nil {
				t.Fatalf("output: %v", err)
			}
			got := ""
			if _, ok := formatIDs[slotWord(out.Format)]; ok {
				got = slotWord(out.Format)
			}
			if got != tc.want {
				t.Errorf("format = %q (raw %q), want %q", got, out.Format, tc.want)
			}
		})
	}
}
