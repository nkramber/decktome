package llm

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/anthropics/anthropic-sdk-go"
	anthropicopt "github.com/anthropics/anthropic-sdk-go/option"
)

// AnthropicName is the provider key in roles.json.
const AnthropicName = "anthropic"

// Anthropic calls the Messages API with a JSON Schema output format (D-22:
// the judge provider). The system prompt carries a cache breakpoint. The
// SDK's own retries are off: the Client owns the budget. The request omits
// the thinking field, so a model with adaptive thinking runs it (D-4).
type Anthropic struct {
	client anthropic.Client
}

// NewAnthropic builds the adapter. timeout bounds one attempt. extra
// options follow the defaults, so a test can set a base URL.
func NewAnthropic(apiKey string, timeout time.Duration, extra ...anthropicopt.RequestOption) *Anthropic {
	opts := []anthropicopt.RequestOption{
		anthropicopt.WithAPIKey(apiKey),
		anthropicopt.WithMaxRetries(0),
		anthropicopt.WithRequestTimeout(timeout),
	}
	opts = append(opts, extra...)
	return &Anthropic{client: anthropic.NewClient(opts...)}
}

// Name implements Provider.
func (a *Anthropic) Name() string { return AnthropicName }

// Complete implements Provider.
func (a *Anthropic) Complete(ctx context.Context, call Call) (Response, error) {
	schema, err := schemaMap(call.Schema)
	if err != nil {
		return Response{}, newErr(ClassTerminal, AnthropicName, call.Model, 0, err)
	}
	params := anthropic.MessageNewParams{
		Model:     call.Model,
		MaxTokens: int64(call.MaxOutputTokens),
		Messages:  []anthropic.MessageParam{anthropic.NewUserMessage(anthropic.NewTextBlock(call.Input))},
		OutputConfig: anthropic.OutputConfigParam{
			Format: anthropic.JSONOutputFormatParam{Schema: schema},
		},
	}
	if call.Instructions != "" {
		params.System = []anthropic.TextBlockParam{{
			Text:         call.Instructions,
			CacheControl: anthropic.NewCacheControlEphemeralParam(),
		}}
	}
	if call.Effort != "" {
		params.OutputConfig.Effort = anthropic.OutputConfigEffort(call.Effort)
	}

	resp, err := a.client.Messages.New(ctx, params)
	if err != nil {
		return Response{}, classifyAnthropic(ctx, err, call.Model)
	}
	out := Response{Model: resp.Model}
	if resp.JSON.Usage.Valid() {
		u := resp.Usage
		out.Usage = &Usage{
			InputTokens:       u.InputTokens + u.CacheReadInputTokens + u.CacheCreationInputTokens,
			CachedInputTokens: u.CacheReadInputTokens,
			CacheWriteTokens:  u.CacheCreationInputTokens,
			OutputTokens:      u.OutputTokens,
			ReasoningTokens:   u.OutputTokensDetails.ThinkingTokens,
		}
	}
	switch resp.StopReason {
	case anthropic.StopReasonEndTurn, anthropic.StopReasonStopSequence:
		// The model finished. Read the output below.
	case anthropic.StopReasonMaxTokens:
		return out, newErr(ClassTruncation, AnthropicName, call.Model, 0,
			fmt.Errorf("output hit the cap of %d tokens", call.MaxOutputTokens))
	case anthropic.StopReasonRefusal:
		return out, newErr(ClassRefusal, AnthropicName, call.Model, 0,
			fmt.Errorf("refusal %q: %s", resp.StopDetails.Category, resp.StopDetails.Explanation))
	default:
		// model_context_window_exceeded, tool_use, pause_turn, or a new
		// value. A higher cap can not fix any of them.
		return out, newErr(ClassTerminal, AnthropicName, call.Model, 0,
			fmt.Errorf("stop reason %q", resp.StopReason))
	}
	var sb strings.Builder
	for _, block := range resp.Content {
		// Thinking blocks and every other block type are not output.
		if t, ok := block.AsAny().(anthropic.TextBlock); ok {
			sb.WriteString(t.Text)
		}
	}
	text := strings.TrimSpace(sb.String())
	if text == "" {
		return out, newErr(ClassTerminal, AnthropicName, call.Model, 0, errors.New("empty output"))
	}
	out.Output = json.RawMessage(text)
	return out, nil
}

func classifyAnthropic(ctx context.Context, err error, model string) error {
	var apierr *anthropic.Error
	if errors.As(err, &apierr) {
		class := ClassTerminal
		switch {
		case apierr.StatusCode == http.StatusTooManyRequests,
			apierr.StatusCode == http.StatusRequestTimeout,
			apierr.StatusCode == http.StatusConflict,
			apierr.StatusCode >= 500:
			class = ClassTransient
		}
		e := newErr(class, AnthropicName, model, apierr.StatusCode, err)
		if class == ClassTransient {
			e = e.withRetryAfter(apierr.Response)
		}
		return e
	}
	return classifyTransport(ctx, err, AnthropicName, model)
}
