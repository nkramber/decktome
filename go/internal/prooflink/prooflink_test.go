package prooflink

import (
	"strings"
	"testing"
)

func TestNewCodeHasTheFormOfACode(t *testing.T) {
	seen := map[string]bool{}
	for range 200 {
		code, err := NewCode()
		if err != nil {
			t.Fatal(err)
		}
		if !Valid(code) {
			t.Fatalf("NewCode answered %q, which Valid refuses", code)
		}
		if seen[code] {
			t.Fatalf("NewCode repeated %q", code)
		}
		seen[code] = true
	}
}

func TestValidRefusesAWrongForm(t *testing.T) {
	for _, code := range []string{"", "abc", "abcdefghijk", "abcde/ghij", "abcde ghij", "abcdé1234"} {
		if Valid(code) {
			t.Errorf("Valid(%q) = true", code)
		}
	}
	if !Valid("Ab3dEf9kQ2") {
		t.Error("Valid refused a good code")
	}
}

func TestKeyHidesTheCode(t *testing.T) {
	code := "Ab3dEf9kQ2"
	key := Key(code)
	if strings.Contains(key, code) || len(key) != 32 {
		t.Fatalf("Key(%q) = %q", code, key)
	}
	if Key(code) != key || Key("Ab3dEf9kQ3") == key {
		t.Fatal("Key is not a stable function of the code")
	}
}
