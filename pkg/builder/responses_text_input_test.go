package builder

import (
	"testing"

	"github.com/stefan79/openai-driver/pkg/openai/responses"
)

func TestResponsesTextInput(t *testing.T) {
	req := &responses.ResponsesRequest{}
	responsesWithTextInput("test")(req)

	if len(req.Input) != 1 {
		t.Errorf("Expected 1 input, got %d", len(req.Input))
	}

	if req.Input[0].Role != responses.RoleUser {
		t.Errorf("Expected role to be user, got %s", req.Input[0].Role)
	}

	if len(req.Input[0].Content) != 1 {
		t.Errorf("Expected 1 content, got %d", len(req.Input[0].Content))
	}

	if req.Input[0].Content[0].ContentType() != responses.InputContentTypeText {
		t.Errorf("Expected type to be text, got %s", req.Input[0].Content[0].ContentType())
	}

	c := req.Input[0].Content[0].(responses.TextInputContent)

	if c.Text != "test" {
		t.Errorf("Expected text to be test, got %s", c.Text)
	}
}
