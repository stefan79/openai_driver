package cli

import (
	"bytes"
	"context"
	"testing"

	"github.com/stefan79/openai-driver/pkg/builder"
	"github.com/stefan79/openai-driver/pkg/client"
	"github.com/stefan79/openai-driver/pkg/openai/responses/response"
	"github.com/stefan79/openai-driver/pkg/output"
)

func TestResponsesCreateCommand_NoOptions(t *testing.T) {
	model := "gpt-4o"

	mockClient := &mockClient{
		createFunc: clientCreateMockFunction,
	}

	o := ResponsesCreateOptions{
		Model: model,
	}

	if err := ResponsesCreateCommand(context.Background(), output.NewDefaultOutputter(), mockClient, builder.NewRegistry(), &o); err != nil {
		t.Fatal(err)
	}

	if receivedModel != model {
		t.Errorf("Expected model to be %s, got %s", model, receivedModel)
	}
}

func TestResponsesCreateCommand_WithPrompt(t *testing.T) {
	prompt := "What is the meaning of life?"
	var receivedPrompt string

	mockClient := &mockClient{
		createFunc: clientCreateMockFunction,
	}

	o := ResponsesCreateOptions{
		Prompt: prompt,
	}

	registry := &mockRegistry{
		TextInputFunc: func(text string) builder.ResponsesOption {
			receivedPrompt = text
			return nil
		},
	}

	if err := ResponsesCreateCommand(context.Background(), output.NewDefaultOutputter(), mockClient, registry, &o); err != nil {
		t.Fatal(err)
	}

	if len(receivedOptions) != 1 {
		t.Errorf("Expected 1 option, got %d", len(receivedOptions))
	}

	if receivedPrompt != prompt {
		t.Errorf("Expected prompt to be %s, got %s", prompt, receivedPrompt)
	}

}

func TestResponsesCreateCommand_WithFileInput(t *testing.T) {
	fileName := "test.txt"
	fileData := []byte("test")
	var (
		receivedFileName string
		receivedFileData []byte
	)

	mockClient := &mockClient{
		createFunc: clientCreateMockFunction,
	}

	registry := &mockRegistry{
		FileInputFunc: func(name string, data []byte) builder.ResponsesOption {
			receivedFileName = name
			receivedFileData = data
			return nil
		},
	}

	_, err := mockClient.createFunc(context.Background(), "gpt-4o", registry.FileInput(fileName, fileData))
	if err != nil {
		t.Fatal(err)
	}

	if len(receivedOptions) != 1 {
		t.Errorf("Expected 1 option, got %d", len(receivedOptions))
	}

	if receivedFileName != fileName {
		t.Errorf("Expected file name to be %s, got %s", fileName, receivedFileName)
	}

	if !bytes.Equal(receivedFileData, fileData) {
		t.Errorf("Expected file data to be %s, got %s", fileData, receivedFileData)
	}

}

func TestResponsesCreateCommand_WithReasoningEffort(t *testing.T) {
	effort, err := builder.ParseReasoningEffort("high")
	if err != nil {
		t.Fatal(err)
	}

	var receivedEffort builder.ResponsesReasoningEffort

	mockClient := &mockClient{
		createFunc: clientCreateMockFunction,
	}

	registry := &mockRegistry{
		ReasoningEffortFunc: func(e builder.ResponsesReasoningEffort) builder.ResponsesOption {
			receivedEffort = e
			return nil
		},
	}

	_, err = mockClient.Create(context.Background(), "gpt-4o", registry.ReasoningEffort(effort))
	if err != nil {
		t.Fatal(err)
	}

	if len(receivedOptions) != 1 {
		t.Errorf("Expected 1 option, got %d", len(receivedOptions))
	}

	if receivedEffort != effort {
		t.Errorf("Expected effort to be %s, got %s", effort, receivedEffort)
	}

}

func TestResponsesCreateCommand_WithReasoningSummary(t *testing.T) {
	summary, err := builder.ParseReasoningSummary("detailed")
	if err != nil {
		t.Fatal(err)
	}

	var receivedSummary builder.ResponsesReasoningSummary

	o := ResponsesCreateOptions{
		Summary: &summary,
	}

	mockClient := &mockClient{
		createFunc: clientCreateMockFunction,
	}

	registry := &mockRegistry{
		ReasoningSummaryFunc: func(s builder.ResponsesReasoningSummary) builder.ResponsesOption {
			receivedSummary = s
			return nil
		},
	}

	err = ResponsesCreateCommand(context.Background(), output.NewDefaultOutputter(), mockClient, registry, &o)
	if err != nil {
		t.Fatal(err)
	}

	if len(receivedOptions) != 1 {
		t.Errorf("Expected 1 option, got %d", len(receivedOptions))
	}

	if receivedSummary != summary {
		t.Errorf("Expected summary to be %s, got %s", summary, receivedSummary)
	}

}

func TestResponsesCreateCommand_WithWebSearch(t *testing.T) {
	var receivedWebSearch bool

	mockClient := &mockClient{
		createFunc: clientCreateMockFunction,
	}

	registry := &mockRegistry{
		WebSearchFunc: func() builder.ResponsesOption {
			receivedWebSearch = true
			return nil
		},
	}

	o := ResponsesCreateOptions{
		WebSearch: true,
	}

	err := ResponsesCreateCommand(context.Background(), output.NewDefaultOutputter(), mockClient, registry, &o)
	if err != nil {
		t.Fatal(err)
	}

	if len(receivedOptions) != 1 {
		t.Errorf("Expected 1 option, got %d", len(receivedOptions))
	}

	if !receivedWebSearch {
		t.Errorf("Expected web search to be true, got false")
	}

}

func TestResponsesCreateCommand_WithWebSearchContextSize(t *testing.T) {
	size := builder.WebSearchContextSize("small")
	var receivedWebSearchContextSize builder.WebSearchContextSize

	mockClient := &mockClient{
		createFunc: clientCreateMockFunction,
	}

	registry := &mockRegistry{
		WebSearchContextSizeFunc: func(s builder.WebSearchContextSize) builder.ResponsesOption {
			receivedWebSearchContextSize = s
			return nil
		},
	}

	o := ResponsesCreateOptions{
		WebSearchContextSize: &size,
	}

	err := ResponsesCreateCommand(context.Background(), output.NewDefaultOutputter(), mockClient, registry, &o)
	if err != nil {
		t.Fatal(err)
	}

	if len(receivedOptions) != 1 {
		t.Errorf("Expected 1 option, got %d", len(receivedOptions))
	}

	if receivedWebSearchContextSize != size {
		t.Errorf("Expected web search context size to be %s, got %s", size, receivedWebSearchContextSize)
	}

}

var (
	receivedModel   string
	receivedOptions []builder.ResponsesOption
)

type mockRegistry struct {
	TextInputFunc            func(text string) builder.ResponsesOption
	ReasoningEffortFunc      func(e builder.ResponsesReasoningEffort) builder.ResponsesOption
	ReasoningSummaryFunc     func(s builder.ResponsesReasoningSummary) builder.ResponsesOption
	WebSearchFunc            func() builder.ResponsesOption
	WebSearchContextSizeFunc func(s builder.WebSearchContextSize) builder.ResponsesOption
	FileInputFunc            func(name string, data []byte) builder.ResponsesOption
}

func (m *mockRegistry) TextInput(text string) builder.ResponsesOption {
	return m.TextInputFunc(text)
}

func (m *mockRegistry) ReasoningEffort(e builder.ResponsesReasoningEffort) builder.ResponsesOption {
	return m.ReasoningEffortFunc(e)
}

func (m *mockRegistry) ReasoningSummary(s builder.ResponsesReasoningSummary) builder.ResponsesOption {
	return m.ReasoningSummaryFunc(s)
}

func (m *mockRegistry) WebSearch() builder.ResponsesOption {
	return m.WebSearchFunc()
}

func (m *mockRegistry) WebSearchContextSize(s builder.WebSearchContextSize) builder.ResponsesOption {
	return m.WebSearchContextSizeFunc(s)
}

func (m *mockRegistry) FileInput(name string, data []byte) builder.ResponsesOption {
	return m.FileInputFunc(name, data)
}

func clientCreateMockFunction(ctx context.Context, model string, options ...builder.ResponsesOption) (*response.ResponseDef, error) {
	receivedModel = model
	receivedOptions = options
	return &response.ResponseDef{}, nil
}

type mockClient struct {
	createFunc       func(ctx context.Context, model string, options ...builder.ResponsesOption) (*response.ResponseDef, error)
	cancelFunc       func(ctx context.Context, responseId string) error
	retrieveFunc     func(ctx context.Context, responseId string) (*response.ResponseDef, error)
	createStreamFunc func(ctx context.Context, model string, options ...builder.ResponsesOption) (client.StreamReader, error)
}

func (m *mockClient) Create(ctx context.Context, model string, options ...builder.ResponsesOption) (*response.ResponseDef, error) {
	return m.createFunc(ctx, model, options...)
}

func (m *mockClient) Cancel(ctx context.Context, responseId string) error {
	return m.cancelFunc(ctx, responseId)
}

func (m *mockClient) Retrieve(ctx context.Context, responseId string) (*response.ResponseDef, error) {
	return m.retrieveFunc(ctx, responseId)
}

func (m *mockClient) CreateStream(ctx context.Context, model string, options ...builder.ResponsesOption) (client.StreamReader, error) {
	return m.createStreamFunc(ctx, model, options...)
}
