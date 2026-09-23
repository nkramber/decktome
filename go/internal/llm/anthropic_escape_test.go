package llm

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	anthropicopt "github.com/anthropics/anthropic-sdk-go/option"
)

func TestDecodeLiteralEscapes(t *testing.T) {
	cases := []struct {
		name, raw, want string
	}{
		{"em dash", `{"why":"two combos \\u2014 Heliod"}`, "two combos — Heliod"},
		{"card name", `{"why":"Bartolom\\u00e9 and Thr\\u00f3r's Map"}`, "Bartolomé and Thrór's Map"},
		{"upper hex", `{"why":"a \\u00E9"}`, "a é"},
		{"surrogate pair", `{"why":"\\ud83d\\ude00"}`, "\U0001F600"},
		{"real escape", `{"why":"a — b"}`, "a — b"},
		{"escaped backslash", `{"why":"a\\b \\n"}`, `a\b \n`},
		{"quote", `{"why":"\"x\\u2014\""}`, "\"x—\""},
		{"short hex", `{"why":"a \\u20zz"}`, `a \u20zz`},
		{"end of text", `{"why":"a \\u20"}`, `a \u20`},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var out struct {
				Why string `json:"why"`
			}
			if err := json.Unmarshal(decodeLiteralEscapes([]byte(tc.raw)), &out); err != nil {
				t.Fatal(err)
			}
			if out.Why != tc.want {
				t.Errorf("why = %q, want %q", out.Why, tc.want)
			}
		})
	}
}

func TestDecodeLiteralEscapesKeepsCleanOutput(t *testing.T) {
	raw := []byte(`{"format":"commander","why":"a — b"}`)
	if got := decodeLiteralEscapes(raw); &got[0] != &raw[0] {
		t.Errorf("clean output was copied: %s", got)
	}
}

func TestAnthropicDecodesLiteralEscapes(t *testing.T) {
	st := &stub{}
	srv := httptest.NewServer(http.HandlerFunc(st.handler))
	defer srv.Close()
	prov := NewAnthropic("sk-ant-test", time.Second, anthropicopt.WithBaseURL(srv.URL))
	st.set(http.StatusOK, strings.Replace(anthropicOK, `{\"format\":\"commander\"}`,
		`{\"why\":\"combos \\\\u2014 Heliod\"}`, 1))
	res, err := prov.Complete(context.Background(), testCall())
	if err != nil {
		t.Fatal(err)
	}
	var out struct {
		Why string `json:"why"`
	}
	if err := json.Unmarshal(res.Output, &out); err != nil {
		t.Fatal(err)
	}
	if out.Why != "combos — Heliod" {
		t.Errorf("why = %q", out.Why)
	}
}
