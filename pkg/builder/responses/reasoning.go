package responses

import (
	"driver/pkg/openai/responses"
	"fmt"
)

type (
	Effort  string
	Summary string
)

const (
	MinimalEffort Effort = "minimal"
	LowEffort     Effort = "low"
	MediumEffort  Effort = "medium"
	HighEffort    Effort = "high"

	AutoSummary     Summary = "auto"
	ConciseSummary  Summary = "concise"
	DetailedSummary Summary = "detailed"
)

func ParseReasoningEffort(i string) (Effort, error) {
	switch i {
	case "minimal":
		return MinimalEffort, nil
	case "low":
		return LowEffort, nil
	case "medium":
		return MediumEffort, nil
	case "high":
		return HighEffort, nil
	default:
		return "", fmt.Errorf("invalid effort: %s", i)
	}
}

func ParseReasoningSummary(i string) (Summary, error) {
	switch i {
	case "auto":
		return AutoSummary, nil
	case "concise":
		return ConciseSummary, nil
	case "detailed":
		return DetailedSummary, nil
	default:
		return "", fmt.Errorf("invalid summary: %s", i)
	}
}

func (e Effort) ToReasoningEffort() *responses.ReasoningEffort {
	out := responses.ReasoningEffort(e)
	return &out
}

func (s Summary) ToReasoningSummary() *responses.ReasoningSummary {
	out := responses.ReasoningSummary(s)
	return &out
}

func WithReasoningEffort(e Effort) ResponseOption {
	return func(r *responses.ResponsesRequest) {
		if r.Reasoning == nil {
			r.Reasoning = &responses.Reasoning{}
		}
		r.Reasoning.Effort = e.ToReasoningEffort()
	}
}

func WithReasoningSummary(s Summary) ResponseOption {
	return func(r *responses.ResponsesRequest) {
		if r.Reasoning == nil {
			r.Reasoning = &responses.Reasoning{}
		}
		r.Reasoning.Summary = s.ToReasoningSummary()
	}
}
