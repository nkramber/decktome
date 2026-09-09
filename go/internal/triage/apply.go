package triage

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// Apply writes the cases into the files that own them (D-642). The owner
// reads the change on the pull request, which is the accept step, so no
// staging file sits between the triage and the gate.
//
// Every write is a text insert and never a re-encode. conversations.json
// holds 108 conversations, and a whole re-encode would put every one of
// them into the diff, which would hide the one line that matters.

// tailRe matches the whitespace that runs to the close of the array and
// the close of the document. A new case goes at the start of that run,
// so the tail of the file comes through the write unchanged.
var tailRe = regexp.MustCompile(`\s*\]\s*\}\s*$`)

// idRe reads every "id": <n> of a gate file.
var idRe = regexp.MustCompile(`"id"\s*:\s*(\d+)`)

// NextID answers one past the highest id the file holds, so an appended
// case never takes an id the file already uses.
func NextID(path string) (int, error) {
	raw, err := os.ReadFile(path) // #nosec G304 -- the caller names a file of this repo.
	if err != nil {
		return 0, fmt.Errorf("triage: %s: %w", path, err)
	}
	next := 1
	for _, m := range idRe.FindAllSubmatch(raw, -1) {
		n, err := strconv.Atoi(string(m[1]))
		if err != nil {
			continue
		}
		if n >= next {
			next = n + 1
		}
	}
	return next, nil
}

// Append writes one case body into the file that owns it. The body joins
// the array as the last element, indented the way its neighbours are.
func Append(path string, body json.RawMessage) error {
	raw, err := os.ReadFile(path) // #nosec G304 -- the caller names a file of this repo.
	if err != nil {
		return fmt.Errorf("triage: %s: %w", path, err)
	}
	loc := tailRe.FindIndex(raw)
	if loc == nil {
		return fmt.Errorf("triage: %s does not end with an array and a document, so no case can join it", path)
	}
	var buf bytes.Buffer
	if err := json.Indent(&buf, body, "    ", "  "); err != nil {
		return fmt.Errorf("triage: %s: %w", path, err)
	}
	head := raw[:loc[0]]
	// An empty array takes no comma. Every gate file holds cases today,
	// so this is the shape of a file somebody makes later.
	sep := ",\n    "
	if len(head) == 0 || head[len(head)-1] == '[' {
		sep = "\n    "
	}
	out := append([]byte{}, head...)
	out = append(out, []byte(sep+buf.String())...)
	out = append(out, raw[loc[0]:]...)
	// The writer edits text and never re-encodes, so it can not lean on
	// the encoder to keep the file valid. A file that no longer parses
	// never reaches the disk.
	if !json.Valid(out) {
		return fmt.Errorf("triage: the case would leave %s unparseable, so nothing was written", path)
	}
	// The file is a source file of this repo and not a secret.
	if err := os.WriteFile(path, out, 0o644); err != nil { // #nosec G306
		return fmt.Errorf("triage: %s: %w", path, err)
	}
	return nil
}

// oqRe reads every OQ number of the owner-questions file.
var oqRe = regexp.MustCompile(`OQ-(\d+)`)

// waitsHeading is the table a new row joins. Every row under it waits
// for the owner, and the tuning loop refuses to decide one (D-133).
const waitsHeading = "## Waits on a decision"

// NextOQ answers one past the highest OQ number the file holds.
func NextOQ(path string) (int, error) {
	raw, err := os.ReadFile(path) // #nosec G304 -- the caller names a file of this repo.
	if err != nil {
		return 0, fmt.Errorf("triage: %s: %w", path, err)
	}
	next := 1
	for _, m := range oqRe.FindAllSubmatch(raw, -1) {
		n, err := strconv.Atoi(string(m[1]))
		if err != nil {
			continue
		}
		if n >= next {
			next = n + 1
		}
	}
	return next, nil
}

// AppendOwnerQuestion writes one row into the decision queue. A class
// that meets an owner decision makes no fix (D-558), so the row is the
// whole artifact.
func AppendOwnerQuestion(path, number, question, whyYou, blocks string) error {
	raw, err := os.ReadFile(path) // #nosec G304 -- the caller names a file of this repo.
	if err != nil {
		return fmt.Errorf("triage: %s: %w", path, err)
	}
	text := string(raw)
	head := strings.Index(text, waitsHeading)
	if head < 0 {
		return fmt.Errorf("triage: %s holds no %q heading", path, waitsHeading)
	}
	// The row goes at the head of the table, under the two header lines,
	// so the newest question is the first one the owner reads.
	sep := strings.Index(text[head:], "|---|---|---|---|")
	if sep < 0 {
		return fmt.Errorf("triage: %s holds no table under %q", path, waitsHeading)
	}
	at := head + sep + len("|---|---|---|---|")
	row := fmt.Sprintf("\n| %s | %s | %s | %s |", number, oneLine(question), oneLine(whyYou), oneLine(blocks))
	out := text[:at] + row + text[at:]
	if err := os.WriteFile(path, []byte(out), 0o644); err != nil { // #nosec G306
		return fmt.Errorf("triage: %s: %w", path, err)
	}
	return nil
}

// oneLine keeps a row on one line and its cells apart. A reader's own
// words may hold a newline or a pipe, and either one would break the
// table.
func oneLine(s string) string {
	s = strings.ReplaceAll(s, "|", "/")
	s = strings.Join(strings.Fields(s), " ")
	return strings.TrimSpace(s)
}
