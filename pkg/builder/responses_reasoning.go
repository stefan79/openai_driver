package builder

import (
	"fmt"

	"github.com/stefan79/openai-driver/pkg/openai/responses/request"
)

type (
	ResponsesReasoningEffort  string
	ResponsesReasoningSummary string
)

const (
	ResponsesReasoningMinimalEffort ResponsesReasoningEffort = "minimal"
	ResponsesReasoningLowEffort     ResponsesReasoningEffort = "low"
	ResponsesReasoningMediumEffort  ResponsesReasoningEffort = "medium"
	ResponsesReasoningHighEffort    ResponsesReasoningEffort = "high"

	ResponsesReasoningAutoSummary     ResponsesReasoningSummary = "auto"
	ResponsesReasoningConciseSummary  ResponsesReasoningSummary = "concise"
	ResponsesReasoningDetailedSummary ResponsesReasoningSummary = "detailed"
)

func responsesWithReasoningEffort(e ResponsesReasoningEffort) ResponsesOption {
	return func(r *request.RequestDef) {
		if r.Reasoning == nil {
			r.Reasoning = &request.Reasoning{}
		}
		r.Reasoning.Effort = e.toReasoningEffort()
	}
}

func responsesWithReasoningSummary(s ResponsesReasoningSummary) ResponsesOption {
	return func(r *request.RequestDef) {
		if r.Reasoning == nil {
			r.Reasoning = &request.Reasoning{}
		}
		r.Reasoning.Summary = s.toReasoningSummary()
	}
}

func ParseReasoningEffort(i string) (ResponsesReasoningEffort, error) {
	switch i {
	case "minimal":
		return ResponsesReasoningMinimalEffort, nil
	case "low":
		return ResponsesReasoningLowEffort, nil
	case "medium":
		return ResponsesReasoningMediumEffort, nil
	case "high":
		return ResponsesReasoningHighEffort, nil
	default:
		return "", fmt.Errorf("invalid effort: %s", i)
	}
}

func ParseReasoningSummary(i string) (ResponsesReasoningSummary, error) {
	switch i {
	case "auto":
		return ResponsesReasoningAutoSummary, nil
	case "concise":
		return ResponsesReasoningConciseSummary, nil
	case "detailed":
		return ResponsesReasoningDetailedSummary, nil
	default:
		return "", fmt.Errorf("invalid summary: %s", i)
	}
}

func (e ResponsesReasoningEffort) toReasoningEffort() *request.ReasoningEffort {
	out := request.ReasoningEffort(e)
	return &out
}

func (s ResponsesReasoningSummary) toReasoningSummary() *request.ReasoningSummary {
	out := request.ReasoningSummary(s)
	return &out
}
