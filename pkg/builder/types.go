package builder

import "driver/pkg/openai/responses"

type ResponsesOption func(*responses.ResponsesRequest)

type registry struct {
}

type ResponseRegistry interface {
	FileInput(name string, data []byte) ResponsesOption
	TextInput(text string) ResponsesOption
	ReasoningEffort(e Effort) ResponsesOption
	ReasoningSummary(s Summary) ResponsesOption
	WebSearch() ResponsesOption
	WebSearchContextSize(s WebSearchContextSize) ResponsesOption
}

func NewRegistry() ResponseRegistry {
	return &registry{}
}

func (r *registry) FileInput(name string, data []byte) ResponsesOption {
	return responsesWithFileInput(name, data)
}

func (r *registry) TextInput(text string) ResponsesOption {
	return responsesWithTextInput(text)
}

func (r *registry) ReasoningEffort(e Effort) ResponsesOption {
	return responsesWithReasoningEffort(e)
}

func (r *registry) ReasoningSummary(s Summary) ResponsesOption {
	return responsesWithReasoningSummary(s)
}

func (r *registry) WebSearch() ResponsesOption {
	return responsesWithWebSearch()
}

func (r *registry) WebSearchContextSize(s WebSearchContextSize) ResponsesOption {
	return responsesWithWebSearchContextSize(s)
}
