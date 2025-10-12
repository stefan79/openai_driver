package builder

import (
	"testing"

	"github.com/stefan79/openai-driver/pkg/openai"
	"github.com/stefan79/openai-driver/pkg/openai/responses/request"
)

func TestResponsesTextInput(t *testing.T) {
	req := &request.RequestDef{}
	responsesWithTextInput("test")(req)

	if len(req.Input) != 1 {
		t.Errorf("Expected 1 input, got %d", len(req.Input))
	}

	if req.Input[0].Role != openai.RoleUser {
		t.Errorf("Expected role to be user, got %s", req.Input[0].Role)
	}

	if len(req.Input[0].Content) != 1 {
		t.Errorf("Expected 1 content, got %d", len(req.Input[0].Content))
	}

	if req.Input[0].Content[0].ContentType() != request.InputContentTypeText {
		t.Errorf("Expected type to be text, got %s", req.Input[0].Content[0].ContentType())
	}

	c := req.Input[0].Content[0].(request.TextInputContent)

	if c.Text != "test" {
		t.Errorf("Expected text to be test, got %s", c.Text)
	}
}
