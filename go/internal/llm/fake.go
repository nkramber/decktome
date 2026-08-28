package llm

import (
	"context"
	"encoding/json"
	"fmt"
	"io/fs"
	"sort"
	"strings"
	"sync"
)

// FakeName is the provider key for both fakes.
const FakeName = "fake"

// Fake replies from fixture files instead of a network call.
// Local mode without API keys uses it. The fixture for role R is R.json.
type Fake struct {
	fsys fs.FS
}

// NewFake reads fixtures from fsys, normally an embed.FS or os.DirFS.
func NewFake(fsys fs.FS) *Fake {
	return &Fake{fsys: fsys}
}

// Name implements Provider.
func (f *Fake) Name() string { return FakeName }

// Roles lists the roles a fixture file exists for, in name order.
func (f *Fake) Roles() []string {
	entries, err := fs.ReadDir(f.fsys, ".")
	if err != nil {
		return nil
	}
	var out []string
	for _, e := range entries {
		if name, ok := strings.CutSuffix(e.Name(), ".json"); ok && !e.IsDir() {
			out = append(out, name)
		}
	}
	sort.Strings(out)
	return out
}

// Complete returns the fixture for the call's role.
// An absent fixture is an error, not an empty response: a silent default
// would hide a missing test case. The error names the role and lists the
// fixtures present, so a no-key local run says which roles it can serve
// (A-11).
func (f *Fake) Complete(_ context.Context, call Call) (Response, error) {
	data, err := fs.ReadFile(f.fsys, string(call.Role)+".json")
	if err != nil {
		return Response{}, newErr(ClassTerminal, FakeName, call.Model, 0,
			fmt.Errorf("the fixture fake cannot serve role %q, it has fixtures for %s only: %w",
				call.Role, strings.Join(f.Roles(), ", "), err))
	}
	if !json.Valid(data) {
		return Response{}, newErr(ClassTerminal, FakeName, call.Model, 0,
			fmt.Errorf("fixture for role %q is not valid JSON", call.Role))
	}
	return Response{Output: data, Model: call.Model}, nil
}

// Step is one scripted provider answer.
type Step struct {
	Output json.RawMessage
	Err    error
	Usage  *Usage
}

// Script is a Provider that answers from a fixed sequence and records every
// Call. Tests use it to drive the retry classes.
type Script struct {
	mu    sync.Mutex
	steps []Step
	Calls []Call
}

// NewScript makes a Script from steps in order.
func NewScript(steps ...Step) *Script { return &Script{steps: steps} }

// Name implements Provider.
func (s *Script) Name() string { return FakeName }

// Complete pops the next step. Running past the script is a terminal error.
func (s *Script) Complete(_ context.Context, call Call) (Response, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.Calls = append(s.Calls, call)
	if len(s.steps) == 0 {
		return Response{}, newErr(ClassTerminal, FakeName, call.Model, 0, fmt.Errorf("script exhausted after %d calls", len(s.Calls)))
	}
	st := s.steps[0]
	s.steps = s.steps[1:]
	if st.Err != nil {
		return Response{Usage: st.Usage}, st.Err
	}
	return Response{Output: st.Output, Model: call.Model, Usage: st.Usage}, nil
}
