package request

// Skipped Tools for now
type RequestDef struct {
	Background         *bool              `json:"background,omitempty"`
	ConversationId     *string            `json:"conversation,omitempty"`
	Includes           *[]Includes        `json:"includes,omitempty"`
	Input              []InputDef         `json:"input"`
	Instructions       *string            `json:"instructions,omitempty"`
	MaxOutputTokens    *int               `json:"max_output_tokens,omitempty"`
	MaxToolCalls       *int               `json:"max_tool_calls,omitempty"`
	MetaData           *map[string]string `json:"meta_data,omitempty"`
	Model              *string            `json:"model,omitempty"`
	ParallelToolCalls  *bool              `json:"parallel_tool_calls,omitempty"`
	PreviousResponseId *string            `json:"previous_response_id,omitempty"`
	Prompt             *Prompt            `json:"prompt,omitempty"`
	PromptCacheKey     *string            `json:"prompt_cache_key,omitempty"`
	Reasoning          *Reasoning         `json:"reasoning,omitempty"`
	SafetyIdentifier   *string            `json:"safety_identifier,omitempty"`
	ServiceTier        *string            `json:"service_tier,omitempty"`
	Store              *bool              `json:"store,omitempty"`
	Stream             *bool              `json:"stream,omitempty"`
	StreamOptions      *StreamOptions     `json:"stream_options,omitempty"`
	Temperature        *float64           `json:"temperature,omitempty"`
	Text               *Text              `json:"text,omitempty"`
	TopLogProbs        *int               `json:"top_logprobs,omitempty"`
	TopP               *int               `json:"top_p,omitempty"`
	Truncation         *Truncaction       `json:"truncation,omitempty"`
	Tools              *[]BaseToolType    `json:"tools,omitempty"`
}
