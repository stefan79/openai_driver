package response

// Used here /response/output[@type=reasoning]/
type OutputTypeReasoningDef struct {
	Type    OutputTypeEnum                `json:"type"`
	Id      string                        `json:"id"`
	Status  OutputStatusEnum              `json:"string"`
	Summary []any                         `json:"summary"`
	Content OutputTypeReasoningContentDef `json:"content"`
}

func (o *OutputTypeReasoningDef) OutputType() OutputTypeEnum {
	return OutputTypeReasoning
}

// Used here /response/output[@type=reasoning]/content
type OutputTypeReasoningContentDef struct {
	Type string `json:"type"`
	Text string `json:"query"`
}

// Used here /response/output[@type=reasoning]/summary
type OutputTypeReasoningSummaryDef struct {
	Type string `json:"type"`
	Text string `json:"query"`
}
