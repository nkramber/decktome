package main

import "testing"

// TestWithAdminKeepsTheOtherClaims is D-1076.
func TestWithAdminKeepsTheOtherClaims(t *testing.T) {
	in := map[string]any{"tier": "beta"}
	got := withAdmin(in, true)
	if got["admin"] != true || got["tier"] != "beta" {
		t.Errorf("set: %v", got)
	}
	if _, ok := in["admin"]; ok {
		t.Error("the input changed")
	}
	got = withAdmin(got, false)
	if _, ok := got["admin"]; ok || got["tier"] != "beta" {
		t.Errorf("remove: %v", got)
	}
	if got := withAdmin(nil, true); got["admin"] != true {
		t.Errorf("nil claims: %v", got)
	}
}
