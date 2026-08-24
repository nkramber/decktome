// Package llm is the seam for language-model calls.
//
// PR-0c ships the interface and the Fake provider so local mode never needs
// an API key. PR-10 adds the role-to-model layer and the real providers.
// Guardrail 3: no model id at a call site.
package llm

import (
	"context"
	"encoding/json"
	"fmt"
	"io/fs"
)

// Request is one structured-output call.
type Request struct {
	// Role names the call site's job (for example "classify", "generate").
	// PR-10 maps roles to providers and models.
	Role string
	// Input is the prompt input, serialized by the caller.
	Input string
}

// Response carries the raw structured output. The caller validates it
// against its schema. A provider never returns unvalidated prose.
type Response struct {
	// Output is a JSON document.
	Output json.RawMessage
	// Model is the resolved model id, or "fake" for the Fake provider.
	Model string
}

// Provider executes LLM requests.
type Provider interface {
	Complete(ctx context.Context, req Request) (Response, error)
}

// Fake replies from fixture files instead of a network call.
// Local mode and tests use it. The fixture for role R is R.json.
type Fake struct {
	fsys fs.FS
}

// NewFake reads fixtures from fsys, normally an embed.FS or os.DirFS.
func NewFake(fsys fs.FS) *Fake {
	return &Fake{fsys: fsys}
}

// Complete returns the fixture for the request's role.
// An absent fixture is an error, not an empty response: a silent default
// would hide a missing test case.
func (f *Fake) Complete(_ context.Context, req Request) (Response, error) {
	data, err := fs.ReadFile(f.fsys, req.Role+".json")
	if err != nil {
		return Response{}, fmt.Errorf("llm fake: no fixture for role %q: %w", req.Role, err)
	}
	if !json.Valid(data) {
		return Response{}, fmt.Errorf("llm fake: fixture for role %q is not valid JSON", req.Role)
	}
	return Response{Output: data, Model: "fake"}, nil
}
