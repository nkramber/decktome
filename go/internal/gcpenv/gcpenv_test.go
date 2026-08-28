package gcpenv

import (
	"bytes"
	"encoding/json"
	"testing"
)

func TestProjectID(t *testing.T) {
	keys := []string{"PROJECT_ID", "GOOGLE_CLOUD_PROJECT", "FIRESTORE_EMULATOR_HOST", "STORAGE_EMULATOR_HOST", "CARDS_SNAPSHOT_DIR"}
	tests := []struct {
		name    string
		env     map[string]string
		want    string
		wantErr bool
	}{
		{name: "nothing set", wantErr: true},
		{name: "firestore emulator", env: map[string]string{"FIRESTORE_EMULATOR_HOST": "localhost:8081"}, want: LocalProject},
		{name: "storage emulator", env: map[string]string{"STORAGE_EMULATOR_HOST": "http://localhost:4443"}, want: LocalProject},
		{name: "snapshot dir", env: map[string]string{"CARDS_SNAPSHOT_DIR": "/tmp/x"}, want: LocalProject},
		{name: "google project", env: map[string]string{"GOOGLE_CLOUD_PROJECT": "real-proj"}, want: "real-proj"},
		{name: "PROJECT_ID wins", env: map[string]string{"GOOGLE_CLOUD_PROJECT": "real-proj", "PROJECT_ID": "explicit"}, want: "explicit"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, k := range keys {
				t.Setenv(k, "")
			}
			for k, v := range tt.env {
				t.Setenv(k, v)
			}
			got, err := ProjectID()
			if (err != nil) != tt.wantErr {
				t.Fatalf("err = %v, wantErr %v", err, tt.wantErr)
			}
			if got != tt.want {
				t.Errorf("project = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestNewLogger(t *testing.T) {
	tests := []struct {
		name      string
		kService  string
		wantKeys  []string
		wantValue string
	}{
		{name: "local keeps slog keys", wantKeys: []string{"level", "msg"}, wantValue: "WARN"},
		{name: "cloud run renames", kService: "api", wantKeys: []string{"severity", "message"}, wantValue: "WARNING"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Setenv("K_SERVICE", tt.kService)
			var buf bytes.Buffer
			NewLogger(&buf).Warn("hello", "k", "v")
			var row map[string]any
			if err := json.Unmarshal(buf.Bytes(), &row); err != nil {
				t.Fatalf("not JSON: %v: %s", err, buf.String())
			}
			for _, k := range tt.wantKeys {
				if _, ok := row[k]; !ok {
					t.Errorf("key %q missing in %s", k, buf.String())
				}
			}
			if row[tt.wantKeys[0]] != tt.wantValue {
				t.Errorf("%s = %v, want %s", tt.wantKeys[0], row[tt.wantKeys[0]], tt.wantValue)
			}
			if row["k"] != "v" {
				t.Errorf("attr lost: %s", buf.String())
			}
		})
	}
}
