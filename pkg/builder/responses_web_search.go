package builder

import (
	"fmt"

	"github.com/stefan79/openai-driver/pkg/openai/responses"
)

type WebSearchContextSize string

const (
	WebSearchEffortLow    WebSearchContextSize     = "low"
	WebSearchEffortMedium WebSearchContextSize     = "medium"
	EffortHigh            ResponsesReasoningEffort = "high"
)

func responsesWithWebSearch() ResponsesOption {
	return func(r *responses.ResponsesRequest) {
		if r.Tools == nil {
			ts := make([]responses.BaseToolType, 0)
			r.Tools = &ts
		}
		var ws *responses.ToolTypeWebSearchDef
		for _, t := range *r.Tools {
			if t.ToolType() == responses.ToolTypeWebSearch {
				ws = t.(*responses.ToolTypeWebSearchDef)
			}
		}
		if ws == nil {
			ws = &responses.ToolTypeWebSearchDef{
				Type: responses.ToolTypeWebSearch,
			}
			*r.Tools = append(*r.Tools, ws)
		}
	}
}

func ParseWebSearchContextSize(i string) (WebSearchContextSize, error) {
	switch i {
	case "low":
		return WebSearchEffortLow, nil
	case "medium":
		return WebSearchEffortMedium, nil
	default:
		return "", fmt.Errorf("invalid context size: %s", i)
	}
}

func responsesWithWebSearchContextSize(e WebSearchContextSize) ResponsesOption {
	return func(r *responses.ResponsesRequest) {
		if r.Tools == nil {
			ts := make([]responses.BaseToolType, 0)
			r.Tools = &ts
		}
		var ws *responses.ToolTypeWebSearchDef
		for _, t := range *r.Tools {
			if t.ToolType() == responses.ToolTypeWebSearch {
				ws = t.(*responses.ToolTypeWebSearchDef)
			}
		}
		if ws == nil {
			ws = &responses.ToolTypeWebSearchDef{
				Type: responses.ToolTypeWebSearch,
			}
			*r.Tools = append(*r.Tools, ws)
		}
		ws.SearchContextSize = toWebSearchContextSize(e)
	}
}

func toWebSearchContextSize(e WebSearchContextSize) *responses.ToolTypleWebSearchSearchContextSize {
	out := responses.ToolTypleWebSearchSearchContextSize(e)
	return &out
}
