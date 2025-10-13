package builder

import (
        "testing"
)

func TestFileRegistryUpload(t *testing.T) {
        registry := NewFileRegistry()
        data := []byte("content")
        upload, err := registry.Upload(
                registry.Purpose("fine-tune"),
                registry.LocalFile("test.txt", data),
        )
        if err != nil {
                t.Fatalf("expected no error, got %v", err)
        }
        if upload.Purpose != "fine-tune" {
                t.Fatalf("expected purpose to be fine-tune, got %s", upload.Purpose)
        }
        if upload.FileName != "test.txt" {
                t.Fatalf("expected filename to be test.txt, got %s", upload.FileName)
        }
        if upload.File == nil {
                t.Fatalf("expected file to be set")
        }
        if string(upload.File.Data) != string(data) {
                t.Fatalf("expected file data to match")
        }
        if upload.File.MimeType == "" {
                t.Fatalf("expected mime type to be set")
        }
}

func TestFileRegistryUploadValidation(t *testing.T) {
        registry := NewFileRegistry()
        if _, err := registry.Upload(registry.LocalFile("test.txt", []byte("content"))); err == nil {
                t.Fatalf("expected error for missing purpose")
        }
        if _, err := registry.Upload(registry.Purpose("purpose")); err == nil {
                t.Fatalf("expected error for missing file data")
        }
        if _, err := registry.Upload(registry.Purpose("purpose"), registry.LocalFile("", []byte("content"))); err == nil {
                t.Fatalf("expected error for empty filename")
        }
        if _, err := registry.Upload(registry.Purpose("purpose"), registry.LocalFile("file.txt", []byte{})); err == nil {
                t.Fatalf("expected error for empty data")
        }
}
