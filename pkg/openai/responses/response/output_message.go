package response

import (
	"encoding/json"
	"fmt"

	"github.com/stefan79/openai-driver/pkg/openai"
	"github.com/stefan79/openai-driver/pkg/util"
)

var outputMessageRegistry = util.NewRegistry[BaseOutputTypeMessageContent]()
var outputMessageAnnotationsRegistry = util.NewRegistry[BaseOutputTypeMessageContentTypeOutputTextAnnotations]()

func init() {
	outputMessageRegistry.Register("output_text", func() BaseOutputTypeMessageContent {
		return &OutputTypeMessageContentTypeOutputTextDef{}
	})
	outputMessageRegistry.Register("output_refusal", func() BaseOutputTypeMessageContent {
		return &OutputTypeMessageContentTypeRefusalDef{}
	})

	outputMessageAnnotationsRegistry.Register("file_citation", func() BaseOutputTypeMessageContentTypeOutputTextAnnotations {
		return &OutputTypeMessageContentTypeOutputTextAnnotationsFileCitationDef{}
	})
	outputMessageAnnotationsRegistry.Register("url_citation", func() BaseOutputTypeMessageContentTypeOutputTextAnnotations {
		return &OutputTypeMessageContentTypeOutputTextMessageTextAnnotationUrlCitation{}
	})
	outputMessageAnnotationsRegistry.Register("container_file_citation", func() BaseOutputTypeMessageContentTypeOutputTextAnnotations {
		return &OutputContentMessageTextAnnotationContainerFilerCitation{}
	})
	outputMessageAnnotationsRegistry.Register("file_path", func() BaseOutputTypeMessageContentTypeOutputTextAnnotations {
		return &OutputContentMessageTextAnnotationFilePath{}
	})
}

// Used at /response/output[@type=message]/content/type
type OutputTypeMessageContentTypeEnum string

const (
	OutputTypeMessageContentTypeText    OutputTypeMessageContentTypeEnum = "output_text"
	OutputTypeMessageContentTypeRefusal OutputTypeMessageContentTypeEnum = "output_refusal"
)

// Used at /response/output[@type=message]
type OutputTypeMessageDef struct {
	Type    OutputTypeMessageContentTypeEnum `json:"type"`
	Id      string                           `json:"id"`
	Status  OutputStatusEnum                 `json:"status"`
	Role    openai.Role                      `json:"role"`
	Content []BaseOutputTypeMessageContent   `json:"content"`
}

func (o *OutputTypeMessageDef) OutputType() OutputTypeEnum {
	return OutputTypeMessage
}

func (o *OutputTypeMessageDef) UnmarshalJSON(data []byte) error {
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

// Used at Used at /response/output[@type=message]/content
type BaseOutputTypeMessageContent interface {
	ContentType() OutputTypeMessageContentTypeEnum
}

// Used at /response/output[@type=message]/content[@type=refusal]
type OutputTypeMessageContentTypeRefusalDef struct {
	Type    OutputTypeMessageContentTypeEnum `json:"type"`
	Refusal string                           `json:"refusal"`
}

func (t *OutputTypeMessageContentTypeRefusalDef) ContentType() OutputTypeMessageContentTypeEnum {
	return OutputTypeMessageContentTypeRefusal
}

// Used at /response/output[@type=message]/content[@type=output_text]
type OutputTypeMessageContentTypeOutputTextDef struct {
	Type        OutputTypeMessageContentTypeEnum                        `json:"type"`
	Annotations []BaseOutputTypeMessageContentTypeOutputTextAnnotations `json:"annotations,omitempty"`
	Text        string                                                  `json:"text"`
	LogProbs    []OutputTypeMessageContentTypeOutputTextLogprobsDef     `json:"logprobs,omitempty"`
}

func (t *OutputTypeMessageContentTypeOutputTextDef) ContentType() OutputTypeMessageContentTypeEnum {
	return OutputTypeMessageContentTypeText
}

func (o *OutputTypeMessageContentTypeOutputTextDef) UnmarshalJSON(data []byte) error {

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

// Used at Used at /response/output[@type=message]/content[@type=output_text]/annotations
type BaseOutputTypeMessageContentTypeOutputTextAnnotations interface {
	AnnotationType() OutputTypeMessageContentTypeOutputTextAnnotationsTypeEnum
}

// Used at /response/output[@type=message]/content[@type=output_text]/annotations/type
type OutputTypeMessageContentTypeOutputTextAnnotationsTypeEnum string

const (
	OutputTypeMessageContentTypeOutputTextAnnotationsTypeFileCitation          OutputTypeMessageContentTypeOutputTextAnnotationsTypeEnum = "file_citation"
	OutputTypeMessageContentTypeOutputTextAnnotationsTypeUrlCitation           OutputTypeMessageContentTypeOutputTextAnnotationsTypeEnum = "url_citation"
	OutputTypeMessageContentTypeOutputTextAnnotationsTypeContainerFileCitation OutputTypeMessageContentTypeOutputTextAnnotationsTypeEnum = "container_file_citation"
	OutputTypeMessageContentTypeOutputTextAnnotationsTypeFilePath              OutputTypeMessageContentTypeOutputTextAnnotationsTypeEnum = "file_path"
)

// Used at /response/output[@type=message]/content[@type=output_text]/annotations[@type=file_citation]]
type OutputTypeMessageContentTypeOutputTextAnnotationsFileCitationDef struct {
	Type     OutputTypeMessageContentTypeOutputTextAnnotationsTypeEnum `json:"type"`
	FileId   string                                                    `json:"file_id"`
	FileName string                                                    `json:"file_name"`
	Index    int                                                       `json:"file_index"`
}

func (t *OutputTypeMessageContentTypeOutputTextAnnotationsFileCitationDef) AnnotationType() OutputTypeMessageContentTypeOutputTextAnnotationsTypeEnum {
	return OutputTypeMessageContentTypeOutputTextAnnotationsTypeFileCitation
}

// Used at /response/output[@type=message]/content[@type=output_text]/annotations[@type=url_citation]]
type OutputTypeMessageContentTypeOutputTextMessageTextAnnotationUrlCitation struct {
	Type       OutputTypeMessageContentTypeOutputTextAnnotationsTypeEnum `json:"type"`
	EndIndex   int                                                       `json:"end_index"`
	StartIndex int                                                       `json:"start_index"`
	Title      string                                                    `json:"title"`
	Url        string                                                    `json:"url"`
}

func (t OutputTypeMessageContentTypeOutputTextMessageTextAnnotationUrlCitation) AnnotationType() OutputTypeMessageContentTypeOutputTextAnnotationsTypeEnum {
	return OutputTypeMessageContentTypeOutputTextAnnotationsTypeUrlCitation
}

// Used at /response/output[@type=message]/content[@type=output_text]/annotations[@type=container_file_citation]]
type OutputContentMessageTextAnnotationContainerFilerCitation struct {
	Type        OutputTypeMessageContentTypeOutputTextAnnotationsTypeEnum `json:"type"`
	ContainerId string                                                    `json:"container_id"`
	EndIndex    int                                                       `json:"end_index"`
	StartIndex  int                                                       `json:"start_index"`
	FileId      string                                                    `json:"file_id"`
	FileName    string                                                    `json:"filename"`
}

func (t *OutputContentMessageTextAnnotationContainerFilerCitation) AnnotationType() OutputTypeMessageContentTypeOutputTextAnnotationsTypeEnum {
	return OutputTypeMessageContentTypeOutputTextAnnotationsTypeContainerFileCitation
}

// Used at /response/output[@type=message]/content[@type=output_text]/annotations[@type=container_file_citation]]
type OutputContentMessageTextAnnotationFilePath struct {
	Type   OutputTypeMessageContentTypeOutputTextAnnotationsTypeEnum `json:"type"`
	FileId string                                                    `json:"file_id"`
	Index  int                                                       `json:"file_index"`
}

func (t OutputContentMessageTextAnnotationFilePath) AnnotationType() OutputTypeMessageContentTypeOutputTextAnnotationsTypeEnum {
	return OutputTypeMessageContentTypeOutputTextAnnotationsTypeFilePath
}

// Used at /response/output[@type=message]/content[@type=output_text]/logprobs/*
type BaseOutputTypeMessageContentTypeOutputTextLogprobsDef struct {
	Bytes   []byte `json:"bytes"`
	LogProb int    `json:"logprob"`
	Token   string `json:"token"`
}

// Used at /response/output[@type=message]/content[@type=output_text]/logprobs
type OutputTypeMessageContentTypeOutputTextLogprobsDef struct {
	BaseOutputTypeMessageContentTypeOutputTextLogprobsDef
	TopLogProbs []BaseOutputTypeMessageContentTypeOutputTextLogprobsDef `json:"top_logprobs"`
}
