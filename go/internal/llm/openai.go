package llm

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/openai/openai-go/v3"
	openaiopt "github.com/openai/openai-go/v3/option"
	"github.com/openai/openai-go/v3/responses"
	"github.com/openai/openai-go/v3/shared"
)

// OpenAIName is the provider key in roles.json.
const OpenAIName = "openai"

// OpenAI calls the Responses API with strict JSON Schema output (D-21).
// Prompt caching is automatic above 1,024 tokens. CacheKey routes a
// session's calls to the same cache. The SDK's own retries are off: the
// Client owns the budget.
type OpenAI struct {
	client openai.Client
}

// NewOpenAI builds the adapter. timeout bounds one attempt. extra options
// follow the defaults, so a test can set a base URL.
func NewOpenAI(apiKey string, timeout time.Duration, extra ...openaiopt.RequestOption) *OpenAI {
	opts := []openaiopt.RequestOption{
		openaiopt.WithAPIKey(apiKey),
		openaiopt.WithMaxRetries(0),
		openaiopt.WithRequestTimeout(timeout),
	}
	opts = append(opts, extra...)
	return &OpenAI{client: openai.NewClient(opts...)}
}

// Name implements Provider.
func (o *OpenAI) Name() string { return OpenAIName }

// Complete implements Provider.
func (o *OpenAI) Complete(ctx context.Context, call Call) (Response, error) {
	schema, err := schemaMap(call.Schema)
	if err != nil {
		return Response{}, newErr(ClassTerminal, OpenAIName, call.Model, 0, err)
	}
	params := responses.ResponseNewParams{
		Model:           call.Model,
		Input:           responses.ResponseNewParamsInputUnion{OfString: openai.String(call.Input)},
		MaxOutputTokens: openai.Int(int64(call.MaxOutputTokens)),
		Store:           openai.Bool(false),
		Text: responses.ResponseTextConfigParam{
			Format: responses.ResponseFormatTextConfigUnionParam{
				OfJSONSchema: &responses.ResponseFormatTextJSONSchemaConfigParam{
					Name:   call.SchemaName,
					Schema: schema,
					Strict: openai.Bool(true),
				},
			},
		},
	}
	if call.Instructions != "" {
		params.Instructions = openai.String(call.Instructions)
	}
	if call.Effort != "" {
		params.Reasoning = shared.ReasoningParam{Effort: shared.ReasoningEffort(call.Effort)}
	}
	if call.CacheKey != "" {
		params.PromptCacheKey = openai.String(call.CacheKey)
	}

	resp, err := o.client.Responses.New(ctx, params)
	if err != nil {
		return Response{}, classifyOpenAI(ctx, err, call.Model)
	}
	out := Response{Model: resp.Model}
	if resp.JSON.Usage.Valid() {
		out.Usage = &Usage{
			InputTokens:       resp.Usage.InputTokens,
			CachedInputTokens: resp.Usage.InputTokensDetails.CachedTokens,
			OutputTokens:      resp.Usage.OutputTokens,
			ReasoningTokens:   resp.Usage.OutputTokensDetails.ReasoningTokens,
		}
	}
	switch resp.Status {
	case responses.ResponseStatusCompleted:
		// Read the output below.
	case responses.ResponseStatusIncomplete:
		if resp.IncompleteDetails.Reason == "max_output_tokens" {
			return out, newErr(ClassTruncation, OpenAIName, call.Model, 0,
				fmt.Errorf("output hit the cap of %d tokens", call.MaxOutputTokens))
		}
		return out, newErr(ClassRefusal, OpenAIName, call.Model, 0,
			fmt.Errorf("incomplete: %s", resp.IncompleteDetails.Reason))
	case responses.ResponseStatusFailed:
		return out, newErr(ClassTerminal, OpenAIName, call.Model, 0,
			fmt.Errorf("status failed: %s: %s", resp.Error.Code, resp.Error.Message))
	case responses.ResponseStatusQueued, responses.ResponseStatusInProgress:
		// A synchronous call answered before it finished. A new call can
		// finish, so the class is transient.
		return out, newErr(ClassTransient, OpenAIName, call.Model, 0, fmt.Errorf("status %q", resp.Status))
	case responses.ResponseStatusCancelled:
		return out, newErr(ClassTerminal, OpenAIName, call.Model, 0, errors.New("status cancelled"))
	default:
		return out, newErr(ClassTerminal, OpenAIName, call.Model, 0, fmt.Errorf("status %q", resp.Status))
	}
	for _, item := range resp.Output {
		for _, c := range item.Content {
			if c.Type == "refusal" {
				return out, newErr(ClassRefusal, OpenAIName, call.Model, 0, errors.New(c.Refusal))
			}
		}
	}
	text := resp.OutputText()
	if text == "" {
		return out, newErr(ClassTerminal, OpenAIName, call.Model, 0, errors.New("empty output"))
	}
	out.Output = json.RawMessage(text)
	return out, nil
}

func classifyOpenAI(ctx context.Context, err error, model string) error {
	var apierr *openai.Error
	if errors.As(err, &apierr) {
		class := ClassTerminal
		switch {
		case apierr.StatusCode == http.StatusTooManyRequests,
			apierr.StatusCode == http.StatusRequestTimeout,
			apierr.StatusCode == http.StatusConflict,
			apierr.StatusCode >= 500:
			class = ClassTransient
		}
		e := newErr(class, OpenAIName, model, apierr.StatusCode, err)
		if class == ClassTransient {
			e = e.withRetryAfter(apierr.Response)
		}
		return e
	}
	return classifyTransport(ctx, err, OpenAIName, model)
}

// classifyTransport handles errors with no HTTP status. When the caller's
// ctx has ended, the budget is spent: ClassBudget. Every other failure
// (the per-attempt timeout, a connection reset) can succeed on retry.
func classifyTransport(ctx context.Context, err error, provider, model string) error {
	if ctx.Err() != nil {
		return newErr(ClassBudget, provider, model, 0, err)
	}
	return newErr(ClassTransient, provider, model, 0, err)
}

func schemaMap(raw json.RawMessage) (map[string]any, error) {
	var m map[string]any
	if err := json.Unmarshal(raw, &m); err != nil {
		return nil, fmt.Errorf("schema must be a JSON object: %w", err)
	}
	return m, nil
}
