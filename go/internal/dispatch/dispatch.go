// Package dispatch abstracts background-job dispatch.
//
// Production uses Cloud Tasks. Local mode has no Cloud Tasks emulator (roadmap
// F-6), so Local runs the handler in a goroutine instead. Handlers must be
// idempotent in both modes: Cloud Tasks delivers at least once.
package dispatch

import (
	"context"
	"log/slog"
	"sync"
)

// Task names one unit of background work.
type Task struct {
	// Queue selects the queue (production) or is a label (local).
	Queue string
	// Name deduplicates: two dispatches with one name run once.
	Name string
	// Payload is an opaque body the handler decodes.
	Payload []byte
}

// Handler runs one task. A nil error acknowledges the task.
type Handler func(ctx context.Context, task Task) error

// Dispatcher enqueues tasks for later execution.
type Dispatcher interface {
	// Dispatch enqueues the task. It returns after the enqueue, not after
	// the task runs.
	Dispatch(ctx context.Context, task Task) error
}

// maxSeen caps the dedup set. Local mode runs for one dev session, so a
// cap that resets the set is enough. Cloud Tasks keeps names for an
// hour after the task completes, and this set is the local stand-in.
const maxSeen = 10000

// Local runs each task in a goroutine, at most once per task name.
// It is the local-mode replacement for Cloud Tasks.
type Local struct {
	handler Handler
	logger  *slog.Logger

	mu   sync.Mutex
	seen map[string]struct{}
	wg   sync.WaitGroup
}

// NewLocal returns a Local dispatcher that sends every task to handler.
func NewLocal(handler Handler, logger *slog.Logger) *Local {
	return &Local{handler: handler, logger: logger, seen: make(map[string]struct{})}
}

// Dispatch runs the task in a goroutine. A task name that was dispatched
// before is dropped, which copies the Cloud Tasks named-task dedup behavior.
func (l *Local) Dispatch(ctx context.Context, task Task) error {
	if task.Name != "" {
		l.mu.Lock()
		if _, dup := l.seen[task.Name]; dup {
			l.mu.Unlock()
			l.logger.Info("dispatch: duplicate task dropped", "queue", task.Queue, "name", task.Name)
			return nil
		}
		if len(l.seen) >= maxSeen {
			l.logger.Warn("dispatch: dedup set full, reset", "size", len(l.seen))
			clear(l.seen)
		}
		l.seen[task.Name] = struct{}{}
		l.mu.Unlock()
	}
	l.wg.Add(1)
	go func() {
		defer l.wg.Done()
		if err := l.handler(context.WithoutCancel(ctx), task); err != nil {
			l.logger.Error("dispatch: task failed", "queue", task.Queue, "name", task.Name, "err", err)
		}
	}()
	return nil
}

// Wait blocks until every dispatched task returned. Tests use it.
func (l *Local) Wait() {
	l.wg.Wait()
}
