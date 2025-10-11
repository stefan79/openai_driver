package output

// Used here /response/output[@type=reasoning]/
type TypeReasoningDef struct {
	Type    TypeEnum                `json:"type"`
	Id      string                  `json:"id"`
	Status  StatusEnum              `json:"string"`
	Summary []any                   `json:"summary"`
	Content TypeReasoningContentDef `json:"content"`
}

func (o TypeReasoningDef) OutputType() TypeEnum {
	return TypeReasoning
}

// Used here /response/output[@type=reasoning]/content
type TypeReasoningContentDef struct {
	Type string `json:"type"`
	Text string `json:"query"`
}

// Used here /response/output[@type=reasoning]/summary
type TypeReasoningSummaryDef struct {
	Type string `json:"type"`
	Text string `json:"query"`
}
