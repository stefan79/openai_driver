package builder

import (
	"testing"

	"github.com/stefan79/openai-driver/pkg/openai/responses"
)

func TestResponsesReasoningEffort(t *testing.T) {
	req := &responses.ResponsesRequest{}
	responsesWithReasoningEffort(ResponsesReasoningEffort("high"))(req)

	if *req.Reasoning.Effort != responses.ReasoningEffort("high") {
		t.Errorf("Expected effort to be high, got %v", req.Reasoning.Effort)
	}
}

func TestResponsesReasoningSummary(t *testing.T) {
	req := &responses.ResponsesRequest{}
	responsesWithReasoningSummary(ResponsesReasoningSummary("detailed"))(req)

	if *req.Reasoning.Summary != responses.ReasoningSummary("detailed") {
		t.Errorf("Expected summary to be detailed, got %v", req.Reasoning.Summary)
	}
}
