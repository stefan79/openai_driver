package cli

import (
	"context"
	"fmt"
	"strings"

	"github.com/stefan79/openai-driver/pkg/builder"
	"github.com/stefan79/openai-driver/pkg/client"
	"github.com/stefan79/openai-driver/pkg/openai/responses/response"
	"github.com/stefan79/openai-driver/pkg/output"
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

func NewClient(outputter output.Outputter, o *ResponsesCreateOptions) (client.Client, error) {
	return client.NewClient(outputter, o.OpenAIAPIKey, o.BaseUrl, o.Proxy)
}

func ResponsesCreateCommand(ctx context.Context, outputter output.Outputter, client client.Client, registry builder.ResponseRegistry, o *ResponsesCreateOptions) error {
	outputter.VMessage(fmt.Sprintf("Using Model: %s", o.Model))

	options := []builder.ResponsesOption{}
	if o.Prompt != "" {
		outputter.VVVMessage(fmt.Sprintf("Prompt: %s", o.Prompt))
		options = append(options, registry.TextInput(o.Prompt))
	}
	if o.Effort != nil {
		outputter.VVVMessage(fmt.Sprintf("Effort: %v", o.Effort))
		options = append(options, registry.ReasoningEffort(*o.Effort))
	}
	if o.Summary != nil {
		outputter.VVVMessage(fmt.Sprintf("Summary: %v", o.Summary))
		options = append(options, registry.ReasoningSummary(*o.Summary))
	}
	if o.FileName != nil {
		outputter.VVVMessage(fmt.Sprintf("File: %s", *o.FileName))
		options = append(options, registry.FileInput(*o.FileName, o.FileData))
	}
	if o.WebSearch {
		outputter.VVVMessage("Web Search: true")
		options = append(options, registry.WebSearch())
	}
	if o.WebSearchContextSize != nil {
		outputter.VVVMessage(fmt.Sprintf("Web Search Context Size: %v", o.WebSearchContextSize))
		options = append(options, registry.WebSearchContextSize(*o.WebSearchContextSize))
	}
	r, err := client.Create(ctx, o.Model, options...)
	if err != nil {
		return fmt.Errorf("Error creating response: %v\n", err)
	}
	outputs, err := r.SelectOutput(response.SelectMessage)
	if err != nil {
		return fmt.Errorf("Error selecting output: %v\n", err)
	}
	for _, output := range outputs {
		outputter.Output(fmt.Sprintf("Output: %s \n", strings.Join(response.SerializeText(&output), "\n")))
	}
	return nil
}
