package builder

import (
	"testing"

	"github.com/stefan79/openai-driver/pkg/openai/responses/request"
)

func TestResponsesReasoningEffort(t *testing.T) {
	req := &request.RequestDef{}
	responsesWithReasoningEffort(ResponsesReasoningEffort("high"))(req)

	if *req.Reasoning.Effort != request.ReasoningEffort("high") {
		t.Errorf("Expected effort to be high, got %v", req.Reasoning.Effort)
	}
}

func TestResponsesReasoningSummary(t *testing.T) {
	req := &request.RequestDef{}
	responsesWithReasoningSummary(ResponsesReasoningSummary("detailed"))(req)

	if *req.Reasoning.Summary != request.ReasoningSummary("detailed") {
		t.Errorf("Expected summary to be detailed, got %v", req.Reasoning.Summary)
	}
}
