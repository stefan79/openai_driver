package request

type Includes string

const (
	WebSearchSources            Includes = "web_search_call.action.sources"
	CodeInterpreterCallOutputs  Includes = "code_interpreter_call.outputs"
	ComputerCallOutputImageUrls Includes = "computer_call_output.output.image_url"
	FileSearchCallResults       Includes = "file_search_call.results"
	MessageInputImageImageUrls  Includes = "message.input_image.image_url"
	MessageOutputTextLogprobs   Includes = "message.output_text.logprobs"
	ReasoningEncryptedContent   Includes = "reasoning.encrypted_content"
)

type Prompt struct {
	Id        string             `json:"id"`
	Variables *map[string]string `json:"variables,omitempty"`
	Version   *string            `json:"version,omitempty"`
}

type ReasoningEffort string

const (
	ReasoningEffortMinimal ReasoningEffort = "minimal"
	ReasoningEffortLow     ReasoningEffort = "low"
	ReasoningEffortMedium  ReasoningEffort = "medium"
	ReasoningEffortHigh    ReasoningEffort = "high"
)

type ReasoningSummary string

const (
	ReasoningSummaryAuto     ReasoningSummary = "auto"
	ReasoningSummaryConcise  ReasoningSummary = "concise"
	ReasoningSummaryDetailed ReasoningSummary = "detailed"
)

type Reasoning struct {
	Effort  *ReasoningEffort  `json:"effort,omitempty"`
	Summary *ReasoningSummary `json:"summary,omitempty"`
}

type ServiceTier string

const (
	ServiceTierAuto     ServiceTier = "auto"
	ServiceTierDefault  ServiceTier = "default"
	ServiceTierFlex     ServiceTier = "flex"
	ServiceTierPriority ServiceTier = "priority"
)

type StreamOptions struct {
	IncludeObfuscation *bool `json:"include_obfuscation,omitempty"`
}

type FormatType string

const (
	TextJsonSchema FormatType = "json_schema"
	TextText       FormatType = "text"
	TextJsonObject FormatType = "json_object"
)

type Format interface {
	formatType() FormatType
}

type TextFormat struct {
	Type FormatType `json:"type"`
}

func (t TextFormat) formatType() FormatType { return t.Type }

// Replace with JSON Schema Lib
type SchemaFormat struct {
	Type        FormatType             `json:"type"`
	Name        string                 `json:"name"`
	Schema      map[string]interface{} `json:"schema"`
	Description *string                `json:"description,omitempty"`
	Strict      *bool                  `json:"strict,omitempty"`
}

func (s SchemaFormat) formatType() FormatType { return s.Type }

type ObjectFormat struct {
	Type FormatType `json:"type"`
}

func (o ObjectFormat) formatType() FormatType { return o.Type }

type TextVerbosity string

const (
	TextVerbosityLow    TextVerbosity = "low"
	TextVerbosityMedium TextVerbosity = "medium"
	TextVerbosityHigh   TextVerbosity = "high"
)

type Text struct {
	Format    FormatType `json:"format"`
	Verbosity *string    `json:"verbosity,omitempty"`
}

type Truncaction string

const (
	TruncactionDisabled Truncaction = "disabled"
	TruncactionAuto     Truncaction = "auto"
)

// Skipped Tools for now
type ResponsesRequest struct {
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

type ToolType string

const (
	ToolTypeCodeInterpreter    ToolType = "code_interpreter"
	ToolTypeFileSearch         ToolType = "file_search"
	ToolTypeWebSearch          ToolType = "web_search"
	ToolTypeComputerUsePreview ToolType = "computer_use_preview"
	ToolTypeMCP                ToolType = "mcp"
	ToolTypeImageGeneration    ToolType = "image_generation"
)

type ToolTypleWebSearchSearchContextSize string

const (
	ToolTypleWebSearchSearchContextSizeLow    ToolTypleWebSearchSearchContextSize = "low"
	ToolTypleWebSearchSearchContextSizeMedium ToolTypleWebSearchSearchContextSize = "medium"
	ToolTypleWebSearchSearchContextSizeHigh   ToolTypleWebSearchSearchContextSize = "high"
)

// Used in /request/tools/type
type BaseToolType interface {
	ToolType() ToolType
}

// Used in /request/tools[@type=web_search]
type ToolTypeWebSearchDef struct {
	Type              ToolType                             `json:"type"`
	SearchContextSize *ToolTypleWebSearchSearchContextSize `json:"search_context_size,omitempty"`
	Filters           *ToolTypeWebSearchFiltersDef         `json:"filters,omitempty"`
	UserLocation      *ToolTypeWebSearchSearchLocationDef  `json:"user_location,omitempty"`
}

func (t ToolTypeWebSearchDef) ToolType() ToolType { return ToolTypeWebSearch }

type ToolTypeWebSearchFiltersDef struct {
	AllowedDomains []string `json:"allowed_domains,omitempty"`
}

type ToolTypeWebSearchSearchLocationDef struct {
	City     *string `json:"city,omitempty"`
	Country  *string `json:"country,omitempty"`
	Region   *string `json:"region,omitempty"`
	Timezone *string `json:"timezone,omitempty"`
	Type     *string `json:"type,omitempty"`
}
