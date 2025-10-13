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

	if len(text.Annotations) != 2 {
		t.Errorf("Expected 2 annotations, got %d", len(text.Annotations))
	}

	if text.Annotations[0].AnnotationType() != OutputTypeMessageContentTypeOutputTextAnnotationsTypeFileCitation {
		t.Errorf("Expected annotation type file_citation, got %s", text.Annotations[0].AnnotationType())
	}

	fileCitation := text.Annotations[0].(*OutputTypeMessageContentTypeOutputTextAnnotationsFileCitationDef)

	if fileCitation.FileId != "1" {
		t.Errorf("Expected file_id 1, got %s", fileCitation.FileId)
	}

	if fileCitation.FileName != "file1.txt" {
		t.Errorf("Expected file_name file1.txt, got %s", fileCitation.FileName)
	}

	if text.Annotations[1].AnnotationType() != OutputTypeMessageContentTypeOutputTextAnnotationsTypeUrlCitation {
		t.Errorf("Expected annotation type url_citation, got %s", text.Annotations[1].AnnotationType())
	}

	urlCitation := text.Annotations[1].(*OutputTypeMessageContentTypeOutputTextMessageTextAnnotationUrlCitation)

	if urlCitation.Url != "https://example.com" {
		t.Errorf("Expected url https://example.com, got %s", urlCitation.Url)
	}

	if urlCitation.Title != "Example" {
		t.Errorf("Expected title Example, got %s", urlCitation.Title)
	}

	refusal := message.Content[1].(*OutputTypeMessageContentTypeRefusalDef)

	if refusal.Refusal != "I wont do this!" {
		t.Errorf("Expected refusal I wont do this!, got %s", refusal.Refusal)
	}

}
