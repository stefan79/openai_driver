package builder

import "driver/pkg/openai/responses"

func responsesWithTextInput(text string) ResponsesOption {
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
