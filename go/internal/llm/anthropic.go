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
	"github.com/anthropics/anthropic-sdk-go/option"
)

// AnthropicName is the provider key in roles.json.
const AnthropicName = "anthropic"

// Anthropic calls the Messages API with a JSON Schema output format (D-22:
// the judge provider). The system prompt carries a cache breakpoint. The
// SDK's own retries are off: the Client owns the budget.
type Anthropic struct {
	client anthropic.Client
}

// NewAnthropic builds the adapter. timeout bounds one attempt.
func NewAnthropic(apiKey string, timeout time.Duration) *Anthropic {
	return &Anthropic{client: anthropic.NewClient(
		option.WithAPIKey(apiKey),
		option.WithMaxRetries(0),
		option.WithRequestTimeout(timeout),
	)}
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
		return Response{}, classifyAnthropic(err, call.Model)
	}
	out := Response{Model: resp.Model, Usage: &Usage{
		InputTokens:       resp.Usage.InputTokens + resp.Usage.CacheReadInputTokens + resp.Usage.CacheCreationInputTokens,
		CachedInputTokens: resp.Usage.CacheReadInputTokens,
		OutputTokens:      resp.Usage.OutputTokens,
	}}
	if !resp.JSON.Usage.Valid() {
		out.Usage = nil
	}
	switch resp.StopReason {
	case anthropic.StopReasonMaxTokens, anthropic.StopReasonModelContextWindowExceeded:
		return out, newErr(ClassTruncation, AnthropicName, call.Model, 0,
			fmt.Errorf("output hit the cap of %d tokens (%s)", call.MaxOutputTokens, resp.StopReason))
	case anthropic.StopReasonRefusal:
		return out, newErr(ClassRefusal, AnthropicName, call.Model, 0,
			fmt.Errorf("refusal %q: %s", resp.StopDetails.Category, resp.StopDetails.Explanation))
	}
	var sb strings.Builder
	for _, block := range resp.Content {
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

func classifyAnthropic(err error, model string) error {
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
		return newErr(class, AnthropicName, model, apierr.StatusCode, err)
	}
	return classifyTransport(err, AnthropicName, model)
}
