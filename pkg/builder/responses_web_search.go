package builder

import (
	"fmt"

	"github.com/stefan79/openai-driver/pkg/openai/responses/request"
)

type WebSearchContextSize string

const (
	WebSearchEffortLow    WebSearchContextSize     = "low"
	WebSearchEffortMedium WebSearchContextSize     = "medium"
	EffortHigh            ResponsesReasoningEffort = "high"
)

func responsesWithWebSearch() ResponsesOption {
	return func(r *request.RequestDef) {
		if r.Tools == nil {
			ts := make([]request.BaseToolType, 0)
			r.Tools = &ts
		}
		var ws *request.ToolTypeWebSearchDef
		for _, t := range *r.Tools {
			if t.ToolType() == request.ToolTypeWebSearch {
				ws = t.(*request.ToolTypeWebSearchDef)
			}
		}
		if ws == nil {
			ws = &request.ToolTypeWebSearchDef{
				Type: request.ToolTypeWebSearch,
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
	return func(r *request.RequestDef) {
		if r.Tools == nil {
			ts := make([]request.BaseToolType, 0)
			r.Tools = &ts
		}
		var ws *request.ToolTypeWebSearchDef
		for _, t := range *r.Tools {
			if t.ToolType() == request.ToolTypeWebSearch {
				ws = t.(*request.ToolTypeWebSearchDef)
			}
		}
		if ws == nil {
			ws = &request.ToolTypeWebSearchDef{
				Type: request.ToolTypeWebSearch,
			}
			*r.Tools = append(*r.Tools, ws)
		}
		ws.SearchContextSize = toWebSearchContextSize(e)
	}
}

func toWebSearchContextSize(e WebSearchContextSize) *request.ToolTypleWebSearchSearchContextSize {
	out := request.ToolTypleWebSearchSearchContextSize(e)
	return &out
}
