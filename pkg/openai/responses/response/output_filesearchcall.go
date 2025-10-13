package response

// Used at /response/output[@type=file_search_call]/status
type OutputTypeFileSearchCallStatusEnum string

const (
	OutputTypeFileSearchCallStatusInProgress OutputTypeFileSearchCallStatusEnum = "in_progress"
	OutputTypeFileSearchCallStatusSearching  OutputTypeFileSearchCallStatusEnum = "searching"
	OutputTypeFileSearchCallStatusIncomplete OutputTypeFileSearchCallStatusEnum = "incomplete"
	TypeFileSearchCallStatusFailed           OutputTypeFileSearchCallStatusEnum = "failed"
)

// Used at /response/output[@type=file_search_call]
type OutputTypeFileSearchCallDef struct {
	Type    OutputTypeEnum                     `json:"type"`
	Id      string                             `json:"id"`
	Queries []string                           `json:"queries"`
	Status  OutputTypeFileSearchCallStatusEnum `json:"status"`
	Results []TypeFileSearchCallResultDef      `json:"results"`
}

func (t *OutputTypeFileSearchCallDef) OutputType() OutputTypeEnum {
	return OutputTypeFileSearchCall
}

// Used at /response/output[@type=file_search_call]/results
type TypeFileSearchCallResultDef struct {
	Attributes map[string]string `json:"attributes"`
	FileId     string            `json:"file_id"`
	Filename   string            `json:"filename"`
	Score      float64           `json:"score"`
	Text       string            `json:"text"`
}
