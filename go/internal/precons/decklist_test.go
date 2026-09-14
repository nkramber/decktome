package precons

import "testing"

// TestDecklistsLeaveTheCommanderOut is M-15: each precon reads as its
// commander and 99 rows, and no row names the commander.
func TestDecklistsLeaveTheCommanderOut(t *testing.T) {
	lists, err := Decklists()
	if err != nil {
		t.Fatal(err)
	}
	if len(lists) == 0 {
		t.Fatal("no precon decklist")
	}
	for _, d := range lists {
		n := 0
		for _, r := range d.Rows {
			n += r.Quantity
			if r.Name == d.Commander {
				t.Errorf("%s: the rows name the commander %q", d.Slug, d.Commander)
			}
		}
		if d.Commander == "" || n != 99 {
			t.Errorf("%s: commander %q and %d cards, want a commander and 99", d.Slug, d.Commander, n)
		}
		if d.Name == "" {
			t.Errorf("%s: no display name", d.Slug)
		}
	}
}
