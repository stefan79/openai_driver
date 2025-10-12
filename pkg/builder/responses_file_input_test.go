package builder

import (
	"testing"

	"github.com/stefan79/openai-driver/pkg/openai"
	"github.com/stefan79/openai-driver/pkg/openai/responses/request"
)

func TestResponsesFileInput_Default(t *testing.T) {
	req := &request.RequestDef{}
	responsesWithFileInput("test.txt", []byte("test"))(req)

	if len(req.Input) != 1 {
		t.Errorf("Expected 1 input, got %d", len(req.Input))
	}

	if req.Input[0].Role != openai.RoleUser {
		t.Errorf("Expected role to be user, got %s", req.Input[0].Role)
	}

	if len(req.Input[0].Content) != 1 {
		t.Errorf("Expected 1 content, got %d", len(req.Input[0].Content))
	}

	if req.Input[0].Content[0].ContentType() != request.InputContentTypeFile {
		t.Errorf("Expected type to be file, got %s", req.Input[0].Content[0].ContentType())
	}

	f := req.Input[0].Content[0].(request.InputContentTypeFileDef)

	if f.FileName == nil {
		t.Errorf("Expected file name to be set")
	}

	if *f.FileName != "test.txt" {
		t.Errorf("Expected file name to be test.txt, got %s", *f.FileName)
	}

	if f.FileData == nil {
		t.Errorf("Expected file data to be set")
	}

}
