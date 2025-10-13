package builder

import (
	"fmt"

	openaifiles "github.com/stefan79/openai-driver/pkg/openai/files"
)

type FileRegistry interface {
	Upload(purpose, fileName string, data []byte) (*openaifiles.UploadRequest, error)
}

type fileRegistry struct{}

func NewFileRegistry() FileRegistry {
	return &fileRegistry{}
}

func (f *fileRegistry) Upload(purpose, fileName string, data []byte) (*openaifiles.UploadRequest, error) {
	if purpose == "" {
		return nil, fmt.Errorf("purpose is required")
	}
	if fileName == "" {
		return nil, fmt.Errorf("file name is required")
	}
	if len(data) == 0 {
		return nil, fmt.Errorf("file data is required")
	}
	return &openaifiles.UploadRequest{
		Purpose:  purpose,
		FileName: fileName,
		FileData: data,
	}, nil
}
