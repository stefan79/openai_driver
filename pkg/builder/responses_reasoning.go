package builder

import (
	"fmt"

	"github.com/stefan79/openai-driver/pkg/openai/responses"
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
	return func(r *responses.ResponsesRequest) {
		if r.Reasoning == nil {
			r.Reasoning = &responses.Reasoning{}
		}
		r.Reasoning.Effort = e.toReasoningEffort()
	}
}

func responsesWithReasoningSummary(s ResponsesReasoningSummary) ResponsesOption {
	return func(r *responses.ResponsesRequest) {
		if r.Reasoning == nil {
			r.Reasoning = &responses.Reasoning{}
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

func (e ResponsesReasoningEffort) toReasoningEffort() *responses.ReasoningEffort {
	out := responses.ReasoningEffort(e)
	return &out
}

func (s ResponsesReasoningSummary) toReasoningSummary() *responses.ReasoningSummary {
	out := responses.ReasoningSummary(s)
	return &out
}
