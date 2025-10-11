package responses

import "driver/pkg/openai/responses"

type (
	Effort  byte
	Summary byte
)

const (
	EffortMinimal Effort = iota
	EffortLow
	EffortMedium
	EffortHigh
)

const (
	SummaryNone Summary = iota
	SummaryAuto
	SummaryConcise
	SummaryDetailed
)

func WithReasoning(e Effort, s Summary) ResponseOption {
	return func(r *responses.ResponsesRequest) {
		var rEff responses.ReasoningEffort
		var rSum responses.ReasoningSummary
		switch e {
		case EffortMinimal:
			rEff = responses.ReasoningEffortMinimal
		case EffortLow:
			rEff = responses.ReasoningEffortLow
		case EffortMedium:
			rEff = responses.ReasoningEffortMedium
		case EffortHigh:
			rEff = responses.ReasoningEffortHigh
		}
		if s == SummaryNone {
			r.Reasoning = &responses.Reasoning{
				Effort: &rEff,
			}
			return
		}
		switch s {
		case SummaryAuto:
			rSum = responses.ReasoningSummaryAuto
		case SummaryConcise:
			rSum = responses.ReasoningSummaryConcise
		case SummaryDetailed:
			rSum = responses.ReasoningSummaryDetailed
		}
		r.Reasoning = &responses.Reasoning{
			Effort:  &rEff,
			Summary: &rSum,
		}
	}
}
