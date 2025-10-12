package builder

import (
	"mime"
	"path/filepath"

	"github.com/stefan79/openai-driver/pkg/openai"
	"github.com/stefan79/openai-driver/pkg/openai/responses/request"
)

func responsesWithFileInput(name string, data []byte) ResponsesOption {

	return func(r *request.RequestDef) {
		mimeType := mime.TypeByExtension(filepath.Ext(name))
		if r.Input == nil {
			r.Input = []request.InputDef{}
		}
		r.Input = append(r.Input, request.InputDef{
			Role: openai.RoleUser,
			Content: []request.BaseInputContent{
				request.InputContentTypeFileDef{
					FileName: &name,
					FileData: openai.NewBase64Bytes(data, mimeType),
					Type:     request.InputContentTypeFile,
				},
			},
		})
	}
}
