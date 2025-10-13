package builder

import (
	"fmt"

	"github.com/stefan79/openai-driver/pkg/openai/files"
)

type FilesOption func(*files.UploadRequest) error

type FilesRegistry interface {
	Purpose(purpose string) FilesOption
	LocalFile(name string, data []byte) FilesOption
	Upload(options ...FilesOption) (*files.UploadRequest, error)
}

type filesRegistry struct{}

func NewFilesRegistry() FilesRegistry {
	return &filesRegistry{}
}

func (f *filesRegistry) Purpose(purpose string) FilesOption {
	return fileWithPurpose(purpose)
}

func (f *filesRegistry) LocalFile(name string, data []byte) FilesOption {
	return fileWithLocalFile(name, data)
}

func (f *filesRegistry) Upload(options ...FilesOption) (*files.UploadRequest, error) {
	return BuildFilesUploadRequest(options...)
}

func BuildFilesUploadRequest(options ...FilesOption) (*files.UploadRequest, error) {
	req := &files.UploadRequest{}
	if err := ApplyFilesOptions(req, options...); err != nil {
		return nil, err
	}
	if err := ValidateFilesUploadRequest(req); err != nil {
		return nil, err
	}
	return req, nil
}

func ApplyFilesOptions(req *files.UploadRequest, options ...FilesOption) error {
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

func ValidateFilesUploadRequest(req *files.UploadRequest) error {
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
