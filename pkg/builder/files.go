package builder

import (
	"fmt"

	"github.com/stefan79/openai-driver/pkg/openai"
	openaifiles "github.com/stefan79/openai-driver/pkg/openai/files"
)

type FileUploadOption func(*openaifiles.UploadRequest) error

type FileRegistry interface {
	Purpose(purpose string) FileUploadOption
	LocalFile(name string, data []byte) FileUploadOption
	Upload(options ...FileUploadOption) (*openaifiles.UploadRequest, error)
}

type fileRegistry struct{}

func NewFileRegistry() FileRegistry {
	return &fileRegistry{}
}

func (f *fileRegistry) Purpose(purpose string) FileUploadOption {
	return func(req *openaifiles.UploadRequest) error {
		if purpose == "" {
			return fmt.Errorf("purpose is required")
		}
		req.Purpose = purpose
		return nil
	}
}

func (f *fileRegistry) LocalFile(name string, data []byte) FileUploadOption {
	return func(req *openaifiles.UploadRequest) error {
		if name == "" {
			return fmt.Errorf("file name is required")
		}
		if len(data) == 0 {
			return fmt.Errorf("file data is required")
		}
		base64Data := openai.NewBase64Bytes(data, "")
		base64Data.SetMimeTypeFromFilename(name)
		req.FileName = name
		req.File = base64Data
		return nil
	}
}

func (f *fileRegistry) Upload(options ...FileUploadOption) (*openaifiles.UploadRequest, error) {
	return BuildUploadRequest(options...)
}

func BuildUploadRequest(options ...FileUploadOption) (*openaifiles.UploadRequest, error) {
	req := &openaifiles.UploadRequest{}
	if err := ApplyFileUploadOptions(req, options...); err != nil {
		return nil, err
	}
	if err := validateUploadRequest(req); err != nil {
		return nil, err
	}
	return req, nil
}

func ApplyFileUploadOptions(req *openaifiles.UploadRequest, options ...FileUploadOption) error {
	for _, option := range options {
		if option == nil {
			continue
		}
		if err := option(req); err != nil {
			return err
		}
	}
	return nil
}

func validateUploadRequest(req *openaifiles.UploadRequest) error {
	if req.Purpose == "" {
		return fmt.Errorf("purpose is required")
	}
	if req.FileName == "" {
		return fmt.Errorf("file name is required")
	}
	if req.File == nil || len(req.File.Data) == 0 {
		return fmt.Errorf("file data is required")
	}
	return nil
}
