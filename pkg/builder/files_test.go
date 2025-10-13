package builder

import (
	"testing"
)

func TestFileRegistryUpload(t *testing.T) {
	registry := NewFileRegistry()
	data := []byte("content")
	upload, err := registry.Upload("fine-tune", "test.txt", data)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if upload.Purpose != "fine-tune" {
		t.Fatalf("expected purpose to be fine-tune, got %s", upload.Purpose)
	}
	if upload.FileName != "test.txt" {
		t.Fatalf("expected filename to be test.txt, got %s", upload.FileName)
	}
	if string(upload.FileData) != string(data) {
		t.Fatalf("expected file data to match")
	}
}

func TestFileRegistryUploadValidation(t *testing.T) {
	registry := NewFileRegistry()
	if _, err := registry.Upload("", "test.txt", []byte("content")); err == nil {
		t.Fatalf("expected error for empty purpose")
	}
	if _, err := registry.Upload("purpose", "", []byte("content")); err == nil {
		t.Fatalf("expected error for empty filename")
	}
	if _, err := registry.Upload("purpose", "file.txt", []byte{}); err == nil {
		t.Fatalf("expected error for empty data")
	}
}
