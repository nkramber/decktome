package dispatch

import (
	"context"
	"log/slog"
	"sync/atomic"
	"testing"
)

func TestLocalDispatch(t *testing.T) {
	tests := []struct {
		name     string
		tasks    []Task
		wantRuns int32
	}{
		{
			name:     "two distinct tasks run twice",
			tasks:    []Task{{Queue: "q", Name: "a"}, {Queue: "q", Name: "b"}},
			wantRuns: 2,
		},
		{
			name:     "duplicate names run once",
			tasks:    []Task{{Queue: "q", Name: "a"}, {Queue: "q", Name: "a"}},
			wantRuns: 1,
		},
		{
			name:     "unnamed tasks never dedup",
			tasks:    []Task{{Queue: "q"}, {Queue: "q"}},
			wantRuns: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var runs atomic.Int32
			d := NewLocal(func(_ context.Context, _ Task) error {
				runs.Add(1)
				return nil
			}, slog.Default())
			for _, task := range tt.tasks {
				if err := d.Dispatch(context.Background(), task); err != nil {
					t.Fatalf("Dispatch: %v", err)
				}
			}
			d.Wait()
			if got := runs.Load(); got != tt.wantRuns {
				t.Errorf("runs = %d, want %d", got, tt.wantRuns)
			}
		})
	}
}
