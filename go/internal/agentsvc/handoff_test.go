package agentsvc

import (
	"os"
	"regexp"
	"strings"
	"testing"
)

// Every field of the Slots message is either read by the build or named
// below as deliberately unread. A new field with neither fails the test
// (D-243).

// unreadSlots names the Slots fields the build does not read, and why.
// A field leaves this map when the build starts reading it.
var unreadSlots = map[string]string{
	// slot_states is the question layer's own bookkeeping. The build reads
	// the values, not whether they were asked.
	"slot_states": "the question layer's bookkeeping, not a build input",

	// locked_oracle_ids is the proto's copy. The build reads the same
	// cards from the private state, which holds the names the user wrote
	// (D-242).
	"locked_oracle_ids": "the build reads the locked cards from the private state",
	// commander_oracle_ids is read through the private state as well,
	// because the state holds the names and the index resolves them.
	"commander_oracle_ids": "the build reads the commanders from the private state",
}

// slotFieldRe reads a field name from a proto message body.
var slotFieldRe = regexp.MustCompile(`(?m)^\s+(?:repeated\s+|map<[^>]+>\s+)?[A-Za-z0-9_.<>, ]+?\s+([a-z_]+)\s*=\s*\d+;`)

func slotsFields(t *testing.T) []string {
	t.Helper()
	raw, err := os.ReadFile("../../../proto/mtg/v1/session.proto")
	if err != nil {
		t.Fatalf("read the proto: %v", err)
	}
	body := string(raw)
	i := strings.Index(body, "message Slots {")
	if i < 0 {
		t.Fatal("no Slots message in the proto")
	}
	j := strings.Index(body[i:], "\n}")
	if j < 0 {
		t.Fatal("the Slots message does not end")
	}
	var out []string
	for _, m := range slotFieldRe.FindAllStringSubmatch(body[i:i+j], -1) {
		out = append(out, m[1])
	}
	if len(out) == 0 {
		t.Fatal("no fields parsed from the Slots message")
	}
	return out
}

// TestEverySlotIsReadOrNamed is D-243. It fails when a slot reaches
// neither the build nor the list above.
func TestEverySlotIsReadOrNamed(t *testing.T) {
	src, err := os.ReadFile("build.go")
	if err != nil {
		t.Fatalf("read build.go: %v", err)
	}
	build := string(src)
	for _, f := range slotsFields(t) {
		getter := "Get" + camel(f) + "("
		read := strings.Contains(build, getter)
		_, named := unreadSlots[f]
		switch {
		case read && named:
			t.Errorf("slot %q is read by the build and still listed as unread. Remove it from unreadSlots.", f)
		case !read && !named:
			t.Errorf("slot %q reaches neither the build nor unreadSlots. Either read it in build.go or say why it is unread.", f)
		}
	}
}

// camel turns a proto field name into its Go getter name.
func camel(s string) string {
	var b strings.Builder
	up := true
	for _, r := range s {
		if r == '_' {
			up = true
			continue
		}
		if up && r >= 'a' && r <= 'z' {
			r -= 'a' - 'A'
		}
		up = false
		b.WriteRune(r)
	}
	return b.String()
}
