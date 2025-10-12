package builder

import (
	"mime"
	"path/filepath"

	"github.com/stefan79/openai-driver/pkg/openai"
	"github.com/stefan79/openai-driver/pkg/openai/responses"
)

func responsesWithFileInput(name string, data []byte) ResponsesOption {

	return func(r *responses.ResponsesRequest) {
		mimeType := mime.TypeByExtension(filepath.Ext(name))
		if r.Input == nil {
			r.Input = []responses.Input{}
		}
		r.Input = append(r.Input, responses.Input{
			Role: responses.RoleUser,
			Content: []responses.InputContent{
				responses.FileInputContent{
					FileName: &name,
					FileData: openai.NewBase64Bytes(data, mimeType),
					Type:     responses.InputContentTypeFile,
				},
			},
		})
	}
}
