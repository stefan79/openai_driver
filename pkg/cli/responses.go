package cli

import (
	"context"
	"fmt"
	"strings"

	"github.com/stefan79/openai-driver/pkg/builder"
	"github.com/stefan79/openai-driver/pkg/client"
	"github.com/stefan79/openai-driver/pkg/openai/responses/resp"
)

type ResponsesCreateOptions struct {
	OpenAIAPIKey         string
	Proxy                *string
	BaseUrl              string
	Model                string
	Output               string
	Prompt               string
	Effort               *builder.ResponsesReasoningEffort
	Summary              *builder.ResponsesReasoningSummary
	FileName             *string
	FileData             []byte
	WebSearch            bool
	WebSearchContextSize *builder.WebSearchContextSize
}

func NewClient(o *ResponsesCreateOptions) (client.Client, error) {
	return client.NewClient(o.OpenAIAPIKey, o.BaseUrl, o.Proxy)
}

func ResponsesCreateCommand(ctx context.Context, client client.Client, registry builder.ResponseRegistry, o *ResponsesCreateOptions) error {
	options := []builder.ResponsesOption{}
	if o.Prompt != "" {
		options = append(options, registry.TextInput(o.Prompt))
	}
	if o.Effort != nil {
		options = append(options, registry.ReasoningEffort(*o.Effort))
	}
	if o.Summary != nil {
		options = append(options, registry.ReasoningSummary(*o.Summary))
	}
	if o.FileName != nil {
		options = append(options, registry.FileInput(*o.FileName, o.FileData))
	}
	if o.WebSearch {
		options = append(options, registry.WebSearch())
	}
	if o.WebSearchContextSize != nil {
		options = append(options, registry.WebSearchContextSize(*o.WebSearchContextSize))
	}
	r, err := client.Create(ctx, o.Model, options...)
	if err != nil {
		return fmt.Errorf("Error creating response: %v\n", err)
	}
	outputs, err := r.SelectOutput(resp.SelectMessage)
	if err != nil {
		return fmt.Errorf("Error selecting output: %v\n", err)
	}
	for _, output := range outputs {
		fmt.Printf("%s\n", strings.Join(resp.SerializeText(&output), "\n"))
	}
	return nil
}
