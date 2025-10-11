package responses

import (
	"driver/pkg/openai"
	"driver/pkg/openai/responses"
	"mime"
	"path/filepath"
)

func WithFileInput(name string, data []byte) ResponseOption {

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
