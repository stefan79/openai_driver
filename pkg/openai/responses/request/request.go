package request

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
