package cli

import (
	"context"
	"errors"
	"fmt"
	"mime"
	"path/filepath"
	"testing"

	"github.com/stefan79/openai-driver/pkg/builder"
	"github.com/stefan79/openai-driver/pkg/client"
	"github.com/stefan79/openai-driver/pkg/openai"
	openaifiles "github.com/stefan79/openai-driver/pkg/openai/files"
	"github.com/stefan79/openai-driver/pkg/openai/responses/resp"
)

func TestFilesUploadCommand(t *testing.T) {
	registry := &mockFileRegistry{}
	client := &mockFileClient{}
	expected := &openaifiles.File{Id: "file-123"}
	client.uploadFunc = func(ctx context.Context, options ...builder.FilesOption) (*openaifiles.File, error) {
		req, err := builder.BuildFilesUploadRequest(options...)
		if err != nil {
			t.Fatalf("unexpected error applying options: %v", err)
		}
		if req.Purpose != openaifiles.PurposeFineTune {
			t.Fatalf("expected purpose fine-tune, got %s", req.Purpose)
		}
		if req.FileName != "test.txt" {
			t.Fatalf("expected filename test.txt, got %s", req.FileName)
		}
		if req.File == nil || string(req.File.Data) != "data" {
			t.Fatalf("expected file data to be applied")
		}
		return expected, nil
	}
	err := FilesUploadCommand(context.Background(), client, registry, &FilesUploadOptions{
		Purpose:  "fine-tune",
		FileName: "test.txt",
		FileData: []byte("data"),
	})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if !registry.uploadCalled {
		t.Fatalf("expected upload to be called")
	}
}

func TestFilesListCommand(t *testing.T) {
	client := &mockFileClient{}
	var receivedPurpose *openaifiles.Purpose
	client.listFunc = func(ctx context.Context, purpose *openaifiles.Purpose) (*openaifiles.ListResponse, error) {
		receivedPurpose = purpose
		return &openaifiles.ListResponse{}, nil
	}
	purpose := "fine-tune"
	err := FilesListCommand(context.Background(), client, &FilesListOptions{Purpose: &purpose})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if receivedPurpose == nil || *receivedPurpose != openaifiles.PurposeFineTune {
		t.Fatalf("expected purpose to be passed")
	}
}

func TestFilesRetrieveCommand(t *testing.T) {
	client := &mockFileClient{}
	client.retrieveFunc = func(ctx context.Context, fileID string) (*openaifiles.File, error) {
		if fileID != "file-123" {
			t.Fatalf("unexpected file id %s", fileID)
		}
		return &openaifiles.File{}, nil
	}
	err := FilesRetrieveCommand(context.Background(), client, &FilesRetrieveOptions{FileID: "file-123"})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}

func TestFilesDeleteCommand(t *testing.T) {
	client := &mockFileClient{}
	client.deleteFunc = func(ctx context.Context, fileID string) (*openaifiles.DeleteResponse, error) {
		if fileID != "file-123" {
			t.Fatalf("unexpected file id %s", fileID)
		}
		return &openaifiles.DeleteResponse{Deleted: true}, nil
	}
	err := FilesDeleteCommand(context.Background(), client, &FilesDeleteOptions{FileID: "file-123"})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}

type mockFileRegistry struct {
	uploadCalled bool
}

func (m *mockFileRegistry) Purpose(purpose builder.FilesPurpose) builder.FilesOption {
	return func(req *openaifiles.UploadRequest) error {
		m.uploadCalled = true
		if purpose == "" {
			return fmt.Errorf("purpose is required")
		}
		req.Purpose = purpose.ToOpenAIPurpose()
		return nil
	}
}

func (m *mockFileRegistry) LocalFile(name string, data []byte) builder.FilesOption {
	return func(req *openaifiles.UploadRequest) error {
		m.uploadCalled = true
		req.FileName = name
		var mimeType string
		if ext := filepath.Ext(name); ext != "" {
			mimeType = mime.TypeByExtension(ext)
		}
		base64 := openai.NewBase64Bytes(data, mimeType)
		req.File = base64
		return nil
	}
}

func (m *mockFileRegistry) Upload(options ...builder.FilesOption) (*openaifiles.UploadRequest, error) {
	m.uploadCalled = true
	return builder.BuildFilesUploadRequest(options...)
}

var _ builder.FilesRegistry = (*mockFileRegistry)(nil)

type mockFileClient struct {
	uploadFunc   func(ctx context.Context, options ...builder.FilesOption) (*openaifiles.File, error)
	listFunc     func(ctx context.Context, purpose *openaifiles.Purpose) (*openaifiles.ListResponse, error)
	retrieveFunc func(ctx context.Context, fileID string) (*openaifiles.File, error)
	deleteFunc   func(ctx context.Context, fileID string) (*openaifiles.DeleteResponse, error)
}

func (m *mockFileClient) UploadFile(ctx context.Context, options ...builder.FilesOption) (*openaifiles.File, error) {
	if m.uploadFunc == nil {
		return nil, errors.New("upload not implemented")
	}
	return m.uploadFunc(ctx, options...)
}

func (m *mockFileClient) ListFiles(ctx context.Context, purpose *openaifiles.Purpose) (*openaifiles.ListResponse, error) {
	if m.listFunc == nil {
		return nil, errors.New("list not implemented")
	}
	return m.listFunc(ctx, purpose)
}

func (m *mockFileClient) RetrieveFile(ctx context.Context, fileID string) (*openaifiles.File, error) {
	if m.retrieveFunc == nil {
		return nil, errors.New("retrieve not implemented")
	}
	return m.retrieveFunc(ctx, fileID)
}

func (m *mockFileClient) DeleteFile(ctx context.Context, fileID string) (*openaifiles.DeleteResponse, error) {
	if m.deleteFunc == nil {
		return nil, errors.New("delete not implemented")
	}
	return m.deleteFunc(ctx, fileID)
}

func (m *mockFileClient) DownloadFile(ctx context.Context, fileID string) ([]byte, error) {
	return nil, errors.New("download not implemented")
}

func (m *mockFileClient) Create(ctx context.Context, model string, options ...builder.ResponsesOption) (*resp.ResponseDef, error) {
	return nil, errors.New("not implemented")
}

func (m *mockFileClient) CreateStream(ctx context.Context, model string, options ...builder.ResponsesOption) (client.StreamReader, error) {
	return nil, errors.New("not implemented")
}

func (m *mockFileClient) Retrieve(ctx context.Context, responseId string) (*resp.ResponseDef, error) {
	return nil, errors.New("not implemented")
}

func (m *mockFileClient) Cancel(ctx context.Context, responseId string) error {
	return errors.New("not implemented")
}

var _ client.Client = (*mockFileClient)(nil)
