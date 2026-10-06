package main

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/nkramber/decktome/go/internal/livespend"
)

func TestSpendReadsEachSessionLog(t *testing.T) {
	dir := t.TempDir()
	event := func(text string) string {
		raw, _ := json.Marshal(map[string]any{"type": "user", "message": map[string]any{"content": []any{
			map[string]any{"type": "tool_result", "content": text},
		}}})
		return string(raw) + "\n"
	}
	logs := map[string]string{
		"session-1.log": event(livespend.Marker + `{"id":"1111111111111111","event":"end","target":"chat-probe","usd":0.01,"measured":true}`),
		"session-2.log": event(livespend.Marker + `{"id":"2222222222222222","event":"end","target":"revise-gate","usd":0.5,"measured":true}`),
		"other.log":     event(livespend.Marker + `{"id":"3333333333333333","event":"end","target":"x","usd":9,"measured":true}`),
	}
	for name, text := range logs {
		if err := os.WriteFile(filepath.Join(dir, name), []byte(text), 0o644); err != nil {
			t.Fatal(err)
		}
	}
	var out bytes.Buffer
	file := strings.NewReader(`{"target":"chat-probe","usd":0.30}` + "\n")
	if err := spend(dir, 3, file, &out); err != nil {
		t.Fatal(err)
	}
	var got spendReport
	if err := json.Unmarshal(out.Bytes(), &got); err != nil {
		t.Fatal(err)
	}
	if got.Runs != 2 || got.Charged != 0.51 || got.Refused != 1 || len(got.FileOnly) != 0 {
		t.Fatalf("report = %+v, want 2 runs, $0.51, 1 refused line", got)
	}
}
