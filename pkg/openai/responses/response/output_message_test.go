package response

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

func TestOutputMessage_Kitchensink(t *testing.T) {

	path := "testdata/messages_kitchensink.json"

	absPath, err := filepath.Abs(path)
	if err != nil {
		t.Fatal(err)
	}

	d, err := os.ReadFile(absPath)
	if err != nil {

		t.Fatal(err)
	}

	var rsp ResponseDef
	err = json.Unmarshal(d, &rsp)
	if err != nil {
		t.Fatal(err)
	}

	if len(rsp.Output) != 1 {
		t.Errorf("Expected 1 output, got %d", len(rsp.Output))
	}

	if rsp.Output[0].OutputType() != OutputTypeMessage {
		t.Errorf("Expected output type message, got %s", rsp.Output[0].OutputType())
	}

	message := rsp.Output[0].(*OutputTypeMessageDef)

	if len(message.Content) != 2 {
		t.Errorf("Expected 2 content, got %d", len(message.Content))
	}

	if message.Content[0].ContentType() != OutputTypeMessageContentTypeText {
		t.Errorf("Expected content type text, got %s", message.Content[0].ContentType())
	}

	text := message.Content[0].(*OutputTypeMessageContentTypeOutputTextDef)

	if text.Text != "Hello, world!" {
		t.Errorf("Expected text Hello, world!, got %s", text.Text)
	}

	if message.Content[1].ContentType() != OutputTypeMessageContentTypeRefusal {
		t.Errorf("Expected content type refusal, got %s", message.Content[1].ContentType())
	}

	refusal := message.Content[1].(*OutputTypeMessageContentTypeRefusalDef)

	if refusal.Refusal != "I wont do this!" {
		t.Errorf("Expected refusal I wont do this!, got %s", refusal.Refusal)
	}

}
