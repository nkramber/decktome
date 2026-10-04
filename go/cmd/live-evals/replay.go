package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
)

// replayInput turns a bundle into the inputs of `make chat-probe`: one
// text for each turn of the reader, and the stored collection. The eval
// session replays the chat of the reader on the base code and on its fix
// with them (D-1144). It reads local files alone and costs nothing.
func replayInput(bundle, out string) error {
	var s mtgv1.Session
	if err := readEnvelope(filepath.Join(bundle, "session.json"), &s); err != nil {
		return fmt.Errorf("session: %w", err)
	}
	msgs := replayTurns(&s)
	if len(msgs) == 0 {
		return errors.New("the session holds no turn of the reader")
	}
	if err := os.MkdirAll(out, 0o700); err != nil {
		return err
	}
	b, err := json.MarshalIndent(msgs, "", "  ")
	if err != nil {
		return err
	}
	if err := os.WriteFile(filepath.Join(out, "messages.json"), b, 0o600); err != nil {
		return err
	}
	var c mtgv1.Collection
	switch err := readEnvelope(filepath.Join(bundle, "collection.json"), &c); {
	case errors.Is(err, os.ErrNotExist):
		return nil // a session with no collection replays in the any-cards mode
	case err != nil:
		return fmt.Errorf("collection: %w", err)
	}
	cb, err := protojson.Marshal(&c)
	if err != nil {
		return err
	}
	return os.WriteFile(filepath.Join(out, "collection.json"), cb, 0o600)
}

// replayTurns gives one text for each turn: the answers to the questions
// of the turn before, then the message of the reader. A choice becomes
// the text of its option, because the probe answers in text.
func replayTurns(s *mtgv1.Session) []string {
	asked := map[string]*mtgv1.Question{}
	var out []string
	for _, t := range s.GetTurns() {
		var parts []string
		for _, a := range t.GetAnswers() {
			if text := replayAnswer(asked[a.GetQuestionId()], a); text != "" {
				parts = append(parts, text)
			}
		}
		if m := strings.TrimSpace(t.GetUserMessage()); m != "" {
			parts = append(parts, m)
		}
		for _, q := range t.GetQuestions() {
			asked[q.GetId()] = q
		}
		if len(parts) > 0 {
			out = append(out, strings.Join(parts, ". "))
		}
	}
	return out
}

func replayAnswer(q *mtgv1.Question, a *mtgv1.Answer) string {
	switch {
	case a.GetDeclined():
		return "Skip that question."
	case strings.TrimSpace(a.GetText()) != "":
		return strings.TrimSpace(a.GetText())
	case a.OptionIndex != nil && q != nil && int(a.GetOptionIndex()) < len(q.GetOptions()):
		return q.GetOptions()[a.GetOptionIndex()]
	}
	return ""
}

// readEnvelope reads the data of one bundle file, inside its notices.
func readEnvelope(path string, m proto.Message) error {
	b, err := os.ReadFile(path) // #nosec G304 -- the script names the bundle.
	if err != nil {
		return err
	}
	var env envelope
	if err := json.Unmarshal(b, &env); err != nil {
		return err
	}
	if len(env.Data) == 0 {
		return errors.New("no data")
	}
	return protojson.UnmarshalOptions{DiscardUnknown: true}.Unmarshal(env.Data, m)
}
