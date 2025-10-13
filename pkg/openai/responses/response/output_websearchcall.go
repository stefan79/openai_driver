package response

import (
	"encoding/json"

	"github.com/stefan79/openai-driver/pkg/util"
)

var outputWebSearchCallRegistry = util.NewRegistry[OutputBaseTypeWebSearchCallAction]()

func init() {
	outputWebSearchCallRegistry.Register("search", func() OutputBaseTypeWebSearchCallAction {
		return &OutputTypeWebSearchCallActionTypeSearchDef{}
	})
	outputWebSearchCallRegistry.Register("open_page", func() OutputBaseTypeWebSearchCallAction {
		return &OutputTypeWebSearchCallActionTypeOpenPageDef{}
	})
	outputWebSearchCallRegistry.Register("find", func() OutputBaseTypeWebSearchCallAction {
		return &OutputTypeWebSearchCallActionTypeFindDef{}
	})

}

// Used here /response/output[@type=web_search_call]/
type OutputTypeWebSearchCallDef struct {
	Type   OutputTypeEnum                    `json:"type"`
	Id     string                            `json:"id"`
	Status string                            `json:"string"`
	Action OutputBaseTypeWebSearchCallAction `json:"action"`
}

func (o OutputTypeWebSearchCallDef) OutputType() OutputTypeEnum {
	return OutputTypeWebSearchCall
}

func (o *OutputTypeWebSearchCallDef) UnmarshalJSON(data []byte) error {
	var raw map[string]json.RawMessage
	errHandler := util.NewErrorHandler()
	errHandler.Push("content", json.Unmarshal(data, &raw))

	// If we couldn't parse the initial JSON, return early
	if err := errHandler.Error(); err != nil {
		return err
	}

	util.MapField(raw, "type", &o.Type, errHandler)
	util.MapField(raw, "id", &o.Id, errHandler)
	util.MapField(raw, "status", &o.Status, errHandler)

	if action, ok := raw["action"]; ok && len(action) > 0 {
		var err error
		o.Action, err = util.UnmarshalWithRegistry(action, outputWebSearchCallRegistry, "type")
		errHandler.Push("content", err)
	}

	return errHandler.Error()
}

// Used here /response/output[@type=web_search_call]/action/type
type OutputTypeWebSearchCallActionTypeEnum string

const (
	TypeWebSearchCallActionTypeSearch   OutputTypeWebSearchCallActionTypeEnum = "search"
	TypeWebSearchCallActionTypeOpenPage OutputTypeWebSearchCallActionTypeEnum = "open_page"
	TypeWebSearchCallActionTypeFind     OutputTypeWebSearchCallActionTypeEnum = "find"
)

// Used here /response/output[@type=web_search_call]/action
type OutputBaseTypeWebSearchCallAction interface {
	ActionType() OutputTypeWebSearchCallActionTypeEnum
}

// Used here /response/output[@type=web_search_call]/action[@type=search]
type OutputTypeWebSearchCallActionTypeSearchDef struct {
	Type    OutputTypeWebSearchCallActionTypeEnum           `json:"type"`
	Query   string                                          `json:"query"`
	Sources []OutputTypeWebSearchCallActionTypeSearchSource `json:"sources"`
}

func (t OutputTypeWebSearchCallActionTypeSearchDef) ActionType() OutputTypeWebSearchCallActionTypeEnum {
	return TypeWebSearchCallActionTypeSearch
}

// Used here /response/output[@type=web_search_call]/action[@type=search]/sources
type OutputTypeWebSearchCallActionTypeSearchSource struct {
	Type string `json:"type"`
	Url  string `json:"url"`
}

// Used here /response/output[@type=web_search_call]/action[@type=open_page]
type OutputTypeWebSearchCallActionTypeOpenPageDef struct {
	Type OutputTypeWebSearchCallActionTypeEnum `json:"type"`
	Url  string                                `json:"url"`
}

func (t OutputTypeWebSearchCallActionTypeOpenPageDef) ActionType() OutputTypeWebSearchCallActionTypeEnum {
	return TypeWebSearchCallActionTypeOpenPage
}

// Used here /response/output[@type=web_search_call]/action[@type=find]
type OutputTypeWebSearchCallActionTypeFindDef struct {
	Type    OutputTypeWebSearchCallActionTypeEnum `json:"type"`
	Pattern string                                `json:"pattern"`
	Url     string                                `json:"url"`
}

func (t OutputTypeWebSearchCallActionTypeFindDef) ActionType() OutputTypeWebSearchCallActionTypeEnum {
	return TypeWebSearchCallActionTypeFind
}
