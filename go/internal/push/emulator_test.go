package push

import (
	"context"
	"errors"
	"fmt"
	"os"
	"slices"
	"strings"
	"testing"
	"time"

	"cloud.google.com/go/firestore"
)

// The device store talks to Firestore, and the cap and the removal can
// not be proven without it. These tests need the local emulator:
//
//	firebase emulators:start --only firestore --project mtg-local
//	FIRESTORE_EMULATOR_HOST=127.0.0.1:8281 go test ./internal/push -count=1

func emulatorRepo(t *testing.T) *Repo {
	t.Helper()
	if os.Getenv("FIRESTORE_EMULATOR_HOST") == "" {
		t.Skip("set FIRESTORE_EMULATOR_HOST to run the device store against the emulator")
	}
	client, err := firestore.NewClient(t.Context(), "mtg-local")
	if err != nil {
		t.Fatalf("firestore: %v", err)
	}
	t.Cleanup(func() { _ = client.Close() })
	return NewRepo(client)
}

func uniqueUID() string { return "u-" + time.Now().UTC().Format("150405.000000000") }

// TestAddHasRemove: a device the user adds is held, a second add keeps
// one entry, and a remove ends it. A remove of a missing device is no
// error.
func TestAddHasRemove(t *testing.T) {
	r := emulatorRepo(t)
	ctx := context.Background()
	uid := uniqueUID()
	at := time.Date(2026, 9, 29, 10, 0, 0, 0, time.UTC)
	for range 2 {
		if err := r.Add(ctx, uid, "fid1", at); err != nil {
			t.Fatalf("Add: %v", err)
		}
	}
	if ids, err := r.IDs(ctx, uid); err != nil || !slices.Equal(ids, []string{"fid1"}) {
		t.Fatalf("IDs = %v, %v, want [fid1]", ids, err)
	}
	if ok, err := r.Has(ctx, uid, "fid1"); err != nil || !ok {
		t.Errorf("Has = %v, %v, want true", ok, err)
	}
	if ok, err := r.Has(ctx, uniqueUID()+"x", "fid1"); err != nil || ok {
		t.Errorf("another user holds the device: %v, %v", ok, err)
	}
	if err := r.Remove(ctx, uid, "fid1"); err != nil {
		t.Fatalf("Remove: %v", err)
	}
	if err := r.Remove(ctx, uid, "fid1"); err != nil {
		t.Errorf("a second Remove: %v", err)
	}
	if ok, _ := r.Has(ctx, uid, "fid1"); ok {
		t.Error("the removed device is still held")
	}
	if err := r.Add(ctx, uid, "a/b", at); !errors.Is(err, ErrBadID) {
		t.Errorf("Add of a path = %v, want ErrBadID", err)
	}
}

// TestOneOwnerForEachID is the review finding of #258: the ID belongs
// to the browser, so a second account that registers it takes it from
// the first. A late remove by the first account leaves the second one.
func TestOneOwnerForEachID(t *testing.T) {
	r := emulatorRepo(t)
	ctx := context.Background()
	a, b := uniqueUID()+"a", uniqueUID()+"b"
	id := "fid" + strings.ReplaceAll(uniqueUID(), ".", "")
	at := time.Date(2026, 9, 29, 10, 0, 0, 0, time.UTC)
	if err := r.Add(ctx, a, id, at); err != nil {
		t.Fatalf("Add a: %v", err)
	}
	if err := r.Add(ctx, b, id, at.Add(time.Minute)); err != nil {
		t.Fatalf("Add b: %v", err)
	}
	if ids, _ := r.IDs(ctx, a); len(ids) != 0 {
		t.Errorf("the first account keeps %v after the second registered the ID", ids)
	}
	if ok, _ := r.Has(ctx, b, id); !ok {
		t.Fatal("the second account does not hold the ID")
	}
	if err := r.Remove(ctx, a, id); err != nil {
		t.Fatalf("a late Remove of a: %v", err)
	}
	if ok, _ := r.Has(ctx, b, id); !ok {
		t.Error("a late remove by the first account took the ID from the second")
	}
	if err := r.Add(ctx, a, id, at.Add(2*time.Minute)); err != nil {
		t.Fatalf("Add a again: %v", err)
	}
	if ok, _ := r.Has(ctx, b, id); ok {
		t.Error("the ID stayed with the second account after the first took it back")
	}
}

// TestAddKeepsTheNewestDevices: a device past MaxDevices removes the
// oldest registration.
func TestAddKeepsTheNewestDevices(t *testing.T) {
	r := emulatorRepo(t)
	ctx := context.Background()
	uid := uniqueUID()
	start := time.Date(2026, 9, 29, 10, 0, 0, 0, time.UTC)
	for i := range MaxDevices + 2 {
		if err := r.Add(ctx, uid, fmt.Sprintf("fid%02d", i), start.Add(time.Duration(i)*time.Minute)); err != nil {
			t.Fatalf("Add %d: %v", i, err)
		}
	}
	ids, err := r.IDs(ctx, uid)
	if err != nil {
		t.Fatalf("IDs: %v", err)
	}
	if len(ids) != MaxDevices {
		t.Fatalf("devices = %d, want %d", len(ids), MaxDevices)
	}
	for _, old := range []string{"fid00", "fid01"} {
		if slices.Contains(ids, old) {
			t.Errorf("the oldest device %s stayed", old)
		}
	}
}
