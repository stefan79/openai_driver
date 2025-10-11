package output

import (
	"driver/pkg/util"
	"encoding/json"
)

var outputMessageRegistry = util.NewRegistry[BaseTypeMessageContent]()

func init() {
	outputMessageRegistry.Register("output_text", func() BaseTypeMessageContent {
		return &TypeMessageContentTypeOutputTextDef{}
	})
	outputMessageRegistry.Register("output_refusal", func() BaseTypeMessageContent {
		return &TypeMessageContentTypeRefusalDef{}
	})

}

func (o *TypeMessageDef) UnmarshalJSON(data []byte) error {
	var raw map[string]json.RawMessage
	errHandler := util.NewErrorHandler()
	errHandler.Push("content", json.Unmarshal(data, &raw))

	// If we couldn't parse the initial JSON, return early
	if err := errHandler.Error(); err != nil {
		return err
	}

	util.MapField(raw, "type", &o.Type, errHandler)
	util.MapField(raw, "id", &o.Id, errHandler)
	util.MapField(raw, "status", &o.Status, errHandler)
	util.MapField(raw, "role", &o.Role, errHandler)

	if content, ok := raw["content"]; ok && len(content) > 0 {
		var err error
		o.Content, err = util.UnmarshalSliceWithRegistry(content, outputMessageRegistry, "type")
		errHandler.Push("content", err)
	}

	return errHandler.Error()
}
