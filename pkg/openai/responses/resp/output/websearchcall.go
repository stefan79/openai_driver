package output

// Used here /response/output[@type=web_search_call]/
type TypeWebSearchCallDef struct {
	Type   TypeEnum                    `json:"type"`
	Id     string                      `json:"id"`
	Status string                      `json:"string"`
	Action BaseTypeWebSearchCallAction `json:"action"`
}

// Used here /response/output[@type=web_search_call]/action/type
type TypeWebSearchCallActionTypeEnum string

const (
	TypeWebSearchCallActionTypeSearch   TypeWebSearchCallActionTypeEnum = "search"
	TypeWebSearchCallActionTypeOpenPage TypeWebSearchCallActionTypeEnum = "open_page"
	TypeWebSearchCallActionTypeFind     TypeWebSearchCallActionTypeEnum = "find"
)

// Used here /response/output[@type=web_search_call]/action
type BaseTypeWebSearchCallAction interface {
	actionType() TypeWebSearchCallActionTypeEnum
}

// Used here /response/output[@type=web_search_call]/action[@type=search]
type TypeWebSearchCallActionTypeSearchDef struct {
	Type    TypeWebSearchCallActionTypeEnum           `json:"type"`
	Query   string                                    `json:"query"`
	Sources []TypeWebSearchCallActionTypeSearchSource `json:"sources"`
}

func (t TypeWebSearchCallActionTypeSearchDef) actionType() TypeWebSearchCallActionTypeEnum {
	return TypeWebSearchCallActionTypeSearch
}

// Used here /response/output[@type=web_search_call]/action[@type=search]/sources
type TypeWebSearchCallActionTypeSearchSource struct {
	Type string `json:"type"`
	Url  string `json:"url"`
}

// Used here /response/output[@type=web_search_call]/action[@type=open_page]
type TypeWebSearchCallActionTypeOpenPageDef struct {
	Type TypeWebSearchCallActionTypeEnum `json:"type"`
	Url  string                          `json:"url"`
}

func (t TypeWebSearchCallActionTypeOpenPageDef) actionType() TypeWebSearchCallActionTypeEnum {
	return TypeWebSearchCallActionTypeOpenPage
}

// Used here /response/output[@type=web_search_call]/action[@type=find]
type TypeWebSearchCallActionTypeFindDef struct {
	Type    TypeWebSearchCallActionTypeEnum `json:"type"`
	Pattern string                          `json:"pattern"`
	Url     string                          `json:"url"`
}
