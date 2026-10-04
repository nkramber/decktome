package gcpenv

import (
	"bytes"
	"encoding/json"
	"log/slog"
	"net/http"
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
		runJob    string
		wantKeys  []string
		wantValue string
	}{
		{name: "local keeps slog keys", wantKeys: []string{"level", "msg"}, wantValue: "WARN"},
		{name: "cloud run renames", kService: "api", wantKeys: []string{"severity", "message"}, wantValue: "WARNING"},
		{name: "cloud run job renames", runJob: "mtg-meta", wantKeys: []string{"severity", "message"}, wantValue: "WARNING"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Setenv("K_SERVICE", tt.kService)
			t.Setenv("CLOUD_RUN_JOB", tt.runJob)
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

// TestTraceAttr is D-1117: the trace of the request reaches the log line
// under the key that Cloud Logging reads, and a bad id is refused.
func TestTraceAttr(t *testing.T) {
	const id = "105445aa7843bc8bf206b12000100000"
	tests := []struct {
		name   string
		header string
		value  string
		want   string
	}{
		{name: "cloud trace context", header: "X-Cloud-Trace-Context", value: id + "/1;o=1", want: "projects/p-1/traces/" + id},
		{name: "cloud trace context with no span", header: "X-Cloud-Trace-Context", value: id, want: "projects/p-1/traces/" + id},
		{name: "traceparent", header: "traceparent", value: "00-" + id + "-00f067aa0ba902b7-01", want: "projects/p-1/traces/" + id},
		{name: "not hex", header: "X-Cloud-Trace-Context", value: "zz5445aa7843bc8bf206b12000100000/1", want: ""},
		{name: "too short", header: "X-Cloud-Trace-Context", value: "abc/1", want: ""},
		{name: "all zero", header: "traceparent", value: "00-00000000000000000000000000000000-00f067aa0ba902b7-01", want: ""},
		{name: "absent", want: ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := http.Header{}
			if tt.header != "" {
				h.Set(tt.header, tt.value)
			}
			a := TraceAttr("p-1")(h)
			if tt.want == "" {
				if a.Key != "" {
					t.Errorf("trace = %v, want none", a)
				}
				return
			}
			if a.Key != TraceKey || a.Value.String() != tt.want {
				t.Errorf("trace = %v, want %s=%s", a, TraceKey, tt.want)
			}
		})
	}
	t.Setenv("K_SERVICE", "api")
	var buf bytes.Buffer
	h := http.Header{}
	h.Set("X-Cloud-Trace-Context", id+"/1;o=1")
	NewLogger(&buf).LogAttrs(t.Context(), slog.LevelInfo, "rpc", TraceAttr("p-1")(h))
	var row map[string]any
	if err := json.Unmarshal(buf.Bytes(), &row); err != nil {
		t.Fatalf("not JSON: %v", err)
	}
	if row[TraceKey] != "projects/p-1/traces/"+id {
		t.Errorf("the Cloud Run logger lost the trace: %s", buf.String())
	}
}
