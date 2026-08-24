package llm

import (
	"context"
	"testing"
	"testing/fstest"
)

func TestFakeComplete(t *testing.T) {
	fsys := fstest.MapFS{
		"classify.json": {Data: []byte(`{"format":"commander"}`)},
		"broken.json":   {Data: []byte(`{not json`)},
	}
	tests := []struct {
		name    string
		role    string
		wantErr bool
		want    string
	}{
		{name: "fixture found", role: "classify", want: `{"format":"commander"}`},
		{name: "fixture absent", role: "generate", wantErr: true},
		{name: "fixture invalid", role: "broken", wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := NewFake(fsys)
			res, err := f.Complete(context.Background(), Request{Role: tt.role})
			if tt.wantErr {
				if err == nil {
					t.Fatal("want error, got nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("Complete: %v", err)
			}
			if string(res.Output) != tt.want {
				t.Errorf("output = %s, want %s", res.Output, tt.want)
			}
			if res.Model != "fake" {
				t.Errorf("model = %q, want %q", res.Model, "fake")
			}
		})
	}
}
