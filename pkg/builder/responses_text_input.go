package builder

import (
	"github.com/stefan79/openai-driver/pkg/openai"
	"github.com/stefan79/openai-driver/pkg/openai/responses/request"
)

func responsesWithTextInput(text string) ResponsesOption {
	return func(r *request.RequestDef) {
		if r.Input == nil {
			r.Input = make([]request.InputDef, 0)
		}
		r.Input = append(r.Input, request.InputDef{
			Role: openai.RoleUser,
			Content: []request.BaseInputContent{
				request.InputContentTypeTextDef{
					Text: text,
					Type: request.InputContentTypeImage,
				},
			},
		})
	}
}
