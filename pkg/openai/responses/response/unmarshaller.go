package response

import (
	"encoding/json"

	"github.com/stefan79/openai-driver/pkg/openai/responses/response/output"
	"github.com/stefan79/openai-driver/pkg/util"
)

var outputRegistry = util.NewRegistry[output.BaseType]()

func init() {
	outputRegistry.Register("message", func() output.BaseType {
		return &output.TypeMessageDef{}
	})
	outputRegistry.Register("file_search_call", func() output.BaseType {
		return &output.TypeFileSearchCallDef{}
	})
	outputRegistry.Register("web_search_call", func() output.BaseType {
		return &output.TypeWebSearchCallDef{}
	})
	outputRegistry.Register("reasoning", func() output.BaseType {
		return &output.TypeReasoningDef{}
	})
}

func (r *ResponseDef) UnmarshalJSON(data []byte) error {
	// Unmarshal into a map first
	var raw map[string]json.RawMessage
	errHandler := util.NewErrorHandler()
	errHandler.Push("root", json.Unmarshal(data, &raw))

	// If we couldn't parse the initial JSON, return early
	if err := errHandler.Error(); err != nil {
		return err
	}

	// Unmarshal all fields, collecting any errors
	util.MapField(raw, "background", &r.Background, errHandler)
	util.MapField(raw, "conversation", &r.Conversation, errHandler)
	util.MapField(raw, "created_at", &r.CreatedAt, errHandler)
	util.MapField(raw, "error", &r.Error, errHandler)
	util.MapField(raw, "id", &r.Id, errHandler)
	util.MapField(raw, "incomplete_details", &r.IncompleteDetails, errHandler)
	util.MapField(raw, "instructions", &r.Instructions, errHandler)
	util.MapField(raw, "max_output_tokens", &r.MaxOutputTokens, errHandler)
	util.MapField(raw, "max_tool_calls", &r.MaxToolCalls, errHandler)
	util.MapField(raw, "meta_data", &r.MetaData, errHandler)
	util.MapField(raw, "model", &r.Model, errHandler)
	util.MapField(raw, "object", &r.Object, errHandler)

	if output, ok := raw["output"]; ok && len(output) > 0 {
		var err error
		r.Output, err = util.UnmarshalSliceWithRegistry(output, outputRegistry, "type")
		errHandler.Push("output", err)
	}

	// Return the first error that occurred, or nil if no errors
	return errHandler.Error()
}
