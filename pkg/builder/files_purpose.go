package builder

import (
	"fmt"

	openaifiles "github.com/stefan79/openai-driver/pkg/openai/files"
)

type FilesPurpose string

const (
	FilesPurposeFineTune   FilesPurpose = FilesPurpose(openaifiles.PurposeFineTune)
	FilesPurposeAssistants FilesPurpose = FilesPurpose(openaifiles.PurposeAssistants)
	FilesPurposeBatch      FilesPurpose = FilesPurpose(openaifiles.PurposeBatch)
	FilesPurposeUserData   FilesPurpose = FilesPurpose(openaifiles.PurposeUserData)
	FilesPurposeVision     FilesPurpose = FilesPurpose(openaifiles.PurposeVision)
	FilesPurposeEvals      FilesPurpose = FilesPurpose(openaifiles.PurposeEvals)
)

func ParseFilesPurpose(input string) (FilesPurpose, error) {
	switch input {
	case string(FilesPurposeFineTune):
		return FilesPurposeFineTune, nil
	case string(FilesPurposeAssistants):
		return FilesPurposeAssistants, nil
	case string(FilesPurposeBatch):
		return FilesPurposeBatch, nil
	case string(FilesPurposeUserData):
		return FilesPurposeUserData, nil
	case string(FilesPurposeVision):
		return FilesPurposeVision, nil
	case string(FilesPurposeEvals):
		return FilesPurposeEvals, nil
	default:
		return "", fmt.Errorf("invalid purpose: %s", input)
	}
}

func (p FilesPurpose) ToOpenAIPurpose() openaifiles.Purpose {
	return openaifiles.Purpose(p)
}

func fileWithPurpose(purpose FilesPurpose) FilesOption {
	return func(req *openaifiles.UploadRequest) error {
		if purpose == "" {
			return fmt.Errorf("purpose is required")
		}
		req.Purpose = purpose.ToOpenAIPurpose()
		return nil
	}
}
