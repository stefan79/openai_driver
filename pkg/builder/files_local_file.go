package builder

import (
	"fmt"
	"mime"
	"path/filepath"

	"github.com/stefan79/openai-driver/pkg/openai"
	"github.com/stefan79/openai-driver/pkg/openai/files"
)

func fileWithLocalFile(name string, data []byte) FilesOption {
	return func(req *files.UploadRequest) error {
		if name == "" {
			return fmt.Errorf("file name is required")
		}
		if len(data) == 0 {
			return fmt.Errorf("file data is required")
		}
		var mimeType string
		if ext := filepath.Ext(name); ext != "" {
			mimeType = mime.TypeByExtension(ext)
		}
		base64Data := openai.NewBase64Bytes(data, mimeType)
		req.FileName = name
		req.File = base64Data
		return nil
	}
}
