package builder

import (
	"testing"

	"github.com/stefan79/openai-driver/pkg/openai/responses"
)

func TestResponsesWebSearch(t *testing.T) {
	req := &responses.ResponsesRequest{}
	responsesWithWebSearch()(req)

	if req.Tools == nil {
		t.Errorf("Expected tools to be set")
	}

	if len(*req.Tools) != 1 {
		t.Errorf("Expected 1 tool, got %d", len(*req.Tools))
	}

	if (*req.Tools)[0].ToolType() != responses.ToolTypeWebSearch {
		t.Errorf("Expected tool type to be web search, got %s", (*req.Tools)[0].ToolType())
	}
}

func TestResponsesWebSearchContextSize(t *testing.T) {
	req := &responses.ResponsesRequest{}
	size, err := ParseWebSearchContextSize("medium")
	if err != nil {
		t.Fatal(err)
	}
	responsesWithWebSearchContextSize(size)(req)

	if req.Tools == nil {
		t.Errorf("Expected tools to be set")
	}

	if len(*req.Tools) != 1 {
		t.Errorf("Expected 1 tool, got %d", len(*req.Tools))
	}

	if (*req.Tools)[0].ToolType() != responses.ToolTypeWebSearch {
		t.Errorf("Expected tool type to be web search, got %s", (*req.Tools)[0].ToolType())
	}

	ws := (*req.Tools)[0].(*responses.ToolTypeWebSearchDef)

	if ws.SearchContextSize == nil {
		t.Errorf("Expected search context size to be set")
	}

	if *ws.SearchContextSize != responses.ToolTypleWebSearchSearchContextSizeMedium {
		t.Errorf("Expected search context size to be medium, got %v", *ws.SearchContextSize)
	}
}
