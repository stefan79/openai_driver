package response

import "driver/pkg/openai/response"

func WithTextInput(text string) ResponseOption {
	return func(r *response.ResponsesRequest) {
		if r.Input == nil {
			r.Input = make([]response.Input, 0)
		}
		r.Input = append(r.Input, response.Input{
			Role:    response.RoleUser,
			Content: []response.Content{response.TextContent{Text: text}},
		})
	}
}
