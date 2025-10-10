package responses

import "driver/pkg/openai/responses"

func WithTextInput(text string) ResponseOption {
	return func(r *responses.ResponsesRequest) {
		if r.Input == nil {
			r.Input = make([]responses.Input, 0)
		}
		r.Input = append(r.Input, responses.Input{
			Role: responses.RoleUser,
			Content: []responses.InputContent{
				responses.TextInputContent{
					Text: text,
					Type: responses.InputContentTypeText,
				},
			},
		})
	}
}
