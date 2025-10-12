package output

import (
	"encoding/json"
	"fmt"

	"github.com/stefan79/openai-driver/pkg/util"
)

var outputMessageAnnotationsRegistry = util.NewRegistry[BaseTypeMessageContentTypeOutputTextAnnotations]()
var outputMessageRegistry = util.NewRegistry[BaseTypeMessageContent]()
var outputWebSearchCallRegistry = util.NewRegistry[BaseTypeWebSearchCallAction]()

func init() {
	outputMessageRegistry.Register("output_text", func() BaseTypeMessageContent {
		return &TypeMessageContentTypeOutputTextDef{}
	})
	outputMessageRegistry.Register("output_refusal", func() BaseTypeMessageContent {
		return &TypeMessageContentTypeRefusalDef{}
	})
	outputWebSearchCallRegistry.Register("search", func() BaseTypeWebSearchCallAction {
		return &TypeWebSearchCallActionTypeSearchDef{}
	})
	outputWebSearchCallRegistry.Register("open_page", func() BaseTypeWebSearchCallAction {
		return &TypeWebSearchCallActionTypeOpenPageDef{}
	})
	outputWebSearchCallRegistry.Register("find", func() BaseTypeWebSearchCallAction {
		return &TypeWebSearchCallActionTypeFindDef{}
	})
	outputMessageAnnotationsRegistry.Register("file_citation", func() BaseTypeMessageContentTypeOutputTextAnnotations {
		return &TypeMessageContentTypeOutputTextAnnotationsFileCitationDef{}
	})
	outputMessageAnnotationsRegistry.Register("url_citation", func() BaseTypeMessageContentTypeOutputTextAnnotations {
		return &OutputContentMessageTextAnnotationUrlCitation{}
	})
	outputMessageAnnotationsRegistry.Register("container_file_citation", func() BaseTypeMessageContentTypeOutputTextAnnotations {
		return &OutputContentMessageTextAnnotationContainerFilerCitation{}
	})
	outputMessageAnnotationsRegistry.Register("file_path", func() BaseTypeMessageContentTypeOutputTextAnnotations {
		return &OutputContentMessageTextAnnotationFilePath{}
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

func (o *TypeWebSearchCallDef) UnmarshalJSON(data []byte) error {
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

	if action, ok := raw["action"]; ok && len(action) > 0 {
		var err error
		o.Action, err = util.UnmarshalWithRegistry(action, outputWebSearchCallRegistry, "type")
		errHandler.Push("content", err)
	}

	return errHandler.Error()
}

func (o *TypeMessageContentTypeOutputTextDef) UnmarshalJSON(data []byte) error {

	fmt.Println("Will unmarshal some Annotaitons using the custom handler!")

	var raw map[string]json.RawMessage
	errHandler := util.NewErrorHandler()
	errHandler.Push("content", json.Unmarshal(data, &raw))

	// If we couldn't parse the initial JSON, return early
	if err := errHandler.Error(); err != nil {
		return err
	}

	util.MapField(raw, "type", &o.Type, errHandler)
	util.MapField(raw, "text", &o.Text, errHandler)
	util.MapField(raw, "logprobs", &o.LogProbs, errHandler)

	if annotations, ok := raw["annotations"]; ok && len(annotations) > 0 {
		var err error
		o.Annotations, err = util.UnmarshalSliceWithRegistry(annotations, outputMessageAnnotationsRegistry, "type")
		errHandler.Push("content", err)
	}

	return errHandler.Error()
}
