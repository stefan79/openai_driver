package builder

import (
	"fmt"

	"github.com/stefan79/openai-driver/pkg/openai/files"
)

func fileWithPurpose(purpose string) FilesOption {
	return func(req *files.UploadRequest) error {
		if purpose == "" {
			return fmt.Errorf("purpose is required")
		}
		req.Purpose = purpose
		return nil
	}
}
