package resp

// TODO: Instructions can be a complex object as well: https://platform.openai.com/docs/api-reference/responses/object
// Used at /response
type ResponseDef struct {
	Background        bool                  `json:"background"`
	Conversation      ConversationDef       `json:"conversation"`
	CreatedAt         int64                 `json:"created_at"`
	Error             *ErrorDef             `json:"error,omitempty"`
	Id                string                `json:"id"`
	IncompleteDetails *IncompleteDetailsDef `json:"incomplete_details,omitempty"`
	Instructions      *string               `json:"instructions,omitempty"`
	MaxOutputTokens   *int                  `json:"max_output_tokens,omitempty"`
	MaxToolCalls      *int                  `json:"max_tool_calls,omitempty"`
	MetaData          *map[string]string    `json:"meta_data,omitempty"`
	Model             *string               `json:"model,omitempty"`
	Object            string                `json:"object"`
}

// Used At /response/conversation
type ConversationDef struct {
	Id string `json:"id"`
}

type ErrorDef struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type IncompleteDetailsDef struct {
	Reason string `json:"reason"`
}
