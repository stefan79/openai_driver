package responses

import "driver/pkg/openai"

type Includes string

type InputContentType string

const (
	InputContentTypeText  InputContentType = "input_text"
	InputContentTypeFile  InputContentType = "input_file"
	InputContentTypeImage InputContentType = "input_image"
	InputContentTypeAudio InputContentType = "input_audio"
)

type InputContent interface {
	contentType() InputContentType
}

// TextContent represents text input
type TextInputContent struct {
	Type InputContentType `json:"type"`
	Text string           `json:"text"`
}

func (t TextInputContent) contentType() InputContentType { return InputContentTypeText }

// FileContent represents file input
type FileInputContent struct {
	Type     InputContentType    `json:"type"`
	FileData *openai.Base64Bytes `json:"file_data,omitempty"`
	FileId   *string             `json:"file_id,omitempty"`
	FileUrl  *string             `json:"file_url,omitempty"`
	FileName *string             `json:"filename,omitempty"`
}

func (t FileInputContent) contentType() InputContentType { return InputContentTypeFile }

type ImageDetail string

const (
	ImageDetailAuto ImageDetail = "auto"
	ImageDetailHigh ImageDetail = "high"
	ImageDetailLow  ImageDetail = "low"
)

// ImageContent represents image input
type ImageInputContent struct {
	Type        InputContentType `json:"type"`
	ImageDetail ImageDetail      `json:"detail"`
	FileId      *string          `json:"file_id,omitempty"`
	ImageUrl    *string          `json:"image_url,omitempty"`
}

func (t ImageInputContent) contentType() InputContentType { return InputContentTypeImage }

type AudioFormat string

const (
	AudioFormatMP3 AudioFormat = "mp3"
	AudioFormatWAV AudioFormat = "wav"
)

type AudioInput struct {
	Data   []byte      `json:"data"`
	Format AudioFormat `json:"format"`
}

// AudioContent represents audio input
type AudioInputContent struct {
	Type       InputContentType `json:"type"`
	AudioInput AudioInput       `json:"audio_input"`
}

func (t AudioInputContent) contentType() InputContentType { return InputContentTypeAudio }

const (
	WebSearchSources            Includes = "web_search_call.action.sources"
	CodeInterpreterCallOutputs  Includes = "code_interpreter_call.outputs"
	ComputerCallOutputImageUrls Includes = "computer_call_output.output.image_url"
	FileSearchCallResults       Includes = "file_search_call.results"
	MessageInputImageImageUrls  Includes = "message.input_image.image_url"
	MessageOutputTextLogprobs   Includes = "message.output_text.logprobs"
	ReasoningEncryptedContent   Includes = "reasoning.encrypted_content"
)

type Input struct {
	Role    Role           `json:"role"`
	Content []InputContent `json:"content"`
}

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
	Input              []Input            `json:"input"`
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
}
