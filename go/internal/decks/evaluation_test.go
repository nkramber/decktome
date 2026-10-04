package decks

import (
	"errors"
	"testing"
	"time"
)

// A rewrite must keep the eval mark, or a rename, a share, or the stale
// pass would send a read deck to the live evals again (D-1132).
func TestRestoreKeepsTheEvalMark(t *testing.T) {
	at := time.Date(2026, 10, 4, 2, 0, 0, 0, time.UTC)
	sd := storedDeck{CreatedAt: at.Add(-time.Hour), ShareTokenHash: "h", HasBeenEvaluated: true, EvaluatedAt: &at}
	got, err := restore(sampleDeck("d-1", at), sd)
	if err != nil {
		t.Fatalf("restore: %v", err)
	}
	if !got.HasBeenEvaluated || got.EvaluatedAt == nil || !got.EvaluatedAt.Equal(at) {
		t.Errorf("restore dropped the eval mark: %v %v", got.HasBeenEvaluated, got.EvaluatedAt)
	}
	if !got.CreatedAt.Equal(sd.CreatedAt) || got.ShareTokenHash != "h" {
		t.Errorf("restore dropped the create time or the share hash")
	}
}

// A new deck and a revision start unread: toStored writes no mark.
func TestANewRowIsUnread(t *testing.T) {
	sd := toStored(sampleDeck("d-2", time.Now()), nil)
	if sd.HasBeenEvaluated || sd.EvaluatedAt != nil {
		t.Errorf("a new row holds an eval mark: %v %v", sd.HasBeenEvaluated, sd.EvaluatedAt)
	}
}

func TestEmulatorEvalMark(t *testing.T) {
	repo, done := emulatorRepo(t)
	defer done()
	ctx := t.Context()

	uid := "u-eval-" + repo.NewID("seed")
	created := time.Now().UTC().Truncate(time.Second)
	first, second := repo.NewID(uid), repo.NewID(uid)
	if err := repo.Put(ctx, uid, sampleDeck(first, created)); err != nil {
		t.Fatalf("put: %v", err)
	}
	rev := sampleDeck(second, created.Add(time.Minute))
	rev.RevisedFromDeckId = first
	if err := repo.Put(ctx, uid, rev); err != nil {
		t.Fatalf("put revision: %v", err)
	}

	pending := func() []Pending {
		t.Helper()
		all, err := repo.Unevaluated(ctx)
		if err != nil {
			t.Fatalf("unevaluated: %v", err)
		}
		var mine []Pending
		for _, p := range all {
			if p.UID == uid {
				mine = append(mine, p)
			}
		}
		return mine
	}
	got := pending()
	if len(got) != 2 || got[0].ID != first || got[1].ID != second || got[1].RevisedFrom != first {
		t.Fatalf("unevaluated = %+v, want the build then its revision", got)
	}

	if err := repo.MarkEvaluated(ctx, uid, first, created); err != nil {
		t.Fatalf("mark: %v", err)
	}
	name := "renamed"
	if _, err := repo.Update(ctx, uid, first, &name, nil, false); err != nil {
		t.Fatalf("update: %v", err)
	}
	if err := repo.Share(ctx, uid, first, "tok-"+first, "hash-"+first); err != nil {
		t.Fatalf("share: %v", err)
	}
	if got := pending(); len(got) != 1 || got[0].ID != second {
		t.Fatalf("after the mark, a rename, and a share: unevaluated = %+v, want the revision alone", got)
	}

	if err := repo.MarkEvaluated(ctx, uid, "no-such-deck", created); !errors.Is(err, ErrNotFound) {
		t.Errorf("mark of a missing deck = %v, want ErrNotFound", err)
	}
}
