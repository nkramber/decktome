package pushsvc

import (
	"context"
	"testing"
	"time"

	"connectrpc.com/connect"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
)

type memStore struct{ devices map[string]bool }

func (m *memStore) Add(_ context.Context, uid, id string, _ time.Time) error {
	m.devices[uid+"/"+id] = true
	return nil
}

func (m *memStore) Remove(_ context.Context, uid, id string) error {
	delete(m.devices, uid+"/"+id)
	return nil
}

func (m *memStore) Has(_ context.Context, uid, id string) (bool, error) {
	return m.devices[uid+"/"+id], nil
}

func server(uid string) (*Server, *memStore) {
	st := &memStore{devices: map[string]bool{}}
	return New(st, func(context.Context) string { return uid }), st
}

// TestRegisterGetUnregister: a device turns on and off for its caller
// alone.
func TestRegisterGetUnregister(t *testing.T) {
	s, st := server("u1")
	ctx := t.Context()
	if _, err := s.RegisterDevice(ctx, connect.NewRequest(&mtgv1.RegisterDeviceRequest{InstallationId: "fid1"})); err != nil {
		t.Fatalf("RegisterDevice: %v", err)
	}
	if !st.devices["u1/fid1"] {
		t.Fatalf("devices = %v, want u1/fid1", st.devices)
	}
	res, err := s.GetDevice(ctx, connect.NewRequest(&mtgv1.GetDeviceRequest{InstallationId: "fid1"}))
	if err != nil || !res.Msg.GetRegistered() {
		t.Errorf("GetDevice = %v, %v, want registered", res, err)
	}
	other := New(st, func(context.Context) string { return "u2" })
	res, err = other.GetDevice(ctx, connect.NewRequest(&mtgv1.GetDeviceRequest{InstallationId: "fid1"}))
	if err != nil || res.Msg.GetRegistered() {
		t.Errorf("another user sees the device: %v, %v", res, err)
	}
	if _, err := s.UnregisterDevice(ctx, connect.NewRequest(&mtgv1.UnregisterDeviceRequest{InstallationId: "fid1"})); err != nil {
		t.Fatalf("UnregisterDevice: %v", err)
	}
	if len(st.devices) != 0 {
		t.Errorf("devices = %v, want none", st.devices)
	}
}

// TestRefusals: no user is Unauthenticated, and an id that can not be a
// Firebase Installation ID is an invalid argument.
func TestRefusals(t *testing.T) {
	s, _ := server("")
	_, err := s.RegisterDevice(t.Context(), connect.NewRequest(&mtgv1.RegisterDeviceRequest{InstallationId: "fid1"}))
	if connect.CodeOf(err) != connect.CodeUnauthenticated {
		t.Errorf("no user: %v, want Unauthenticated", err)
	}
	s, st := server("u1")
	for _, id := range []string{"", "a/b", "../x"} {
		_, err := s.RegisterDevice(t.Context(), connect.NewRequest(&mtgv1.RegisterDeviceRequest{InstallationId: id}))
		if connect.CodeOf(err) != connect.CodeInvalidArgument {
			t.Errorf("id %q: %v, want InvalidArgument", id, err)
		}
	}
	if len(st.devices) != 0 {
		t.Errorf("a refused id was stored: %v", st.devices)
	}
}
