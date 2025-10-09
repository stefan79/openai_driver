package output

// Used at /response/output[@type=file_search_call]/status
type TypeFileSearchCallStatusEnum string

const (
	TypeFileSearchCallStatusInProgress TypeFileSearchCallStatusEnum = "in_progress"
	TypeFileSearchCallStatusSearching  TypeFileSearchCallStatusEnum = "searching"
	TypeFileSearchCallStatusIncomplete TypeFileSearchCallStatusEnum = "incomplete"
	TypeFileSearchCallStatusFailed     TypeFileSearchCallStatusEnum = "failed"
)

// Used at /response/output[@type=file_search_call]
type TypeFileSearchCallDef struct {
	Type    TypeEnum                      `json:"type"`
	Id      string                        `json:"id"`
	Queries []string                      `json:"queries"`
	Status  TypeFileSearchCallStatusEnum  `json:"status"`
	Results []TypeFileSearchCallResultDef `json:"results"`
}

func (t TypeFileSearchCallDef) outputType() TypeEnum {
	return TypeFileSearchCall
}

// Used at /response/output[@type=file_search_call]/results
type TypeFileSearchCallResultDef struct {
	Attributes map[string]string `json:"attributes"`
	FileId     string            `json:"file_id"`
	Filename   string            `json:"filename"`
	Score      float64           `json:"score"`
	Text       string            `json:"text"`
}
