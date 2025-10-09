package output

import "driver/pkg/openai/response"

// Used at /response/output[@type=message]/content/type
type TypeMessageContentTypeEnum string

const (
	TypeMessageContentTypeText    TypeMessageContentTypeEnum = "output_text"
	TypeMessageContentTypeRefusal TypeMessageContentTypeEnum = "output_refusal"
)

// Used at /response/output[@type=message]
type TypeMessageDef struct {
	Type    TypeEnum                 `json:"type"`
	Id      string                   `json:"id"`
	Status  StatusEnum               `json:"status"`
	Role    response.Role            `json:"role"`
	Content []BaseTypeMessageContent `json:"content"`
}

func (o TypeMessageDef) outputType() TypeEnum {
	return TypeMessage
}

// Used at Used at /response/output[@type=message]/content
type BaseTypeMessageContent interface {
	contentType() TypeMessageContentTypeEnum
}

// Used at /response/output[@type=message]/content[@type=refusal]
type TypeMessageContentTypeRefusalDef struct {
	Type    TypeMessageContentTypeEnum `json:"type"`
	Refusal string                     `json:"refusal"`
}

func (t TypeMessageContentTypeRefusalDef) contentType() TypeMessageContentTypeEnum {
	return TypeMessageContentTypeRefusal
}

// Used at /response/output[@type=message]/content[@type=output_text]
type TypeMessageContentTypeOutputTextDef struct {
	Type        TypeMessageContentTypeEnum                        `json:"type"`
	Annotations []BaseTypeMessageContentTypeOutputTextAnnotations `json:"annotations,omitempty"`
	Text        string                                            `json:"text"`
	LogProbs    []TypeMessageContentTypeOutputTextLogprobsDef     `json:"logprobs,omitempty"`
}

func (t TypeMessageContentTypeOutputTextDef) contentType() TypeMessageContentTypeEnum {
	return TypeMessageContentTypeText
}

// Used at Used at /response/output[@type=message]/content[@type=output_text]/annotations
type BaseTypeMessageContentTypeOutputTextAnnotations interface {
	annotationType() TypeMessageContentTypeOutputTextAnnotationsTypeEnum
}

// Used at /response/output[@type=message]/content[@type=output_text]/annotations/type
type TypeMessageContentTypeOutputTextAnnotationsTypeEnum string

const (
	TypeMessageContentTypeOutputTextAnnotationsTypeFileCitation          TypeMessageContentTypeOutputTextAnnotationsTypeEnum = "file_citation"
	TypeMessageContentTypeOutputTextAnnotationsTypeUrlCitation           TypeMessageContentTypeOutputTextAnnotationsTypeEnum = "url_citation"
	TypeMessageContentTypeOutputTextAnnotationsTypeContainerFileCitation TypeMessageContentTypeOutputTextAnnotationsTypeEnum = "container_file_citation"
	TypeMessageContentTypeOutputTextAnnotationsTypeFilePath              TypeMessageContentTypeOutputTextAnnotationsTypeEnum = "file_path"
)

// Used at /response/output[@type=message]/content[@type=output_text]/annotations[@type=file_citation]]
type TypeMessageContentTypeOutputTextAnnotationsFileCitationDef struct {
	Type     TypeMessageContentTypeOutputTextAnnotationsTypeEnum `json:"type"`
	FileId   string                                              `json:"file_id"`
	FileName string                                              `json:"file_name"`
	Index    int                                                 `json:"file_index"`
}

func (t TypeMessageContentTypeOutputTextAnnotationsFileCitationDef) annotationType() TypeMessageContentTypeOutputTextAnnotationsTypeEnum {
	return TypeMessageContentTypeOutputTextAnnotationsTypeFileCitation
}

// Used at /response/output[@type=message]/content[@type=output_text]/annotations[@type=url_citation]]
type OutputContentMessageTextAnnotationUrlCitation struct {
	Type       TypeMessageContentTypeOutputTextAnnotationsTypeEnum `json:"type"`
	EndIndex   int                                                 `json:"end_index"`
	StartIndex int                                                 `json:"start_index"`
	Title      string                                              `json:"title"`
	Url        string                                              `json:"url"`
}

func (t OutputContentMessageTextAnnotationUrlCitation) annotationType() TypeMessageContentTypeOutputTextAnnotationsTypeEnum {
	return TypeMessageContentTypeOutputTextAnnotationsTypeUrlCitation
}

// Used at /response/output[@type=message]/content[@type=output_text]/annotations[@type=container_file_citation]]
type OutputContentMessageTextAnnotationContainerFilerCitation struct {
	Type        TypeMessageContentTypeOutputTextAnnotationsTypeEnum `json:"type"`
	ContainerId string                                              `json:"container_id"`
	EndIndex    int                                                 `json:"end_index"`
	StartIndex  int                                                 `json:"start_index"`
	FileId      string                                              `json:"file_id"`
	FileName    string                                              `json:"filename"`
}

func (t OutputContentMessageTextAnnotationContainerFilerCitation) annotationType() TypeMessageContentTypeOutputTextAnnotationsTypeEnum {
	return TypeMessageContentTypeOutputTextAnnotationsTypeContainerFileCitation
}

// Used at /response/output[@type=message]/content[@type=output_text]/annotations[@type=container_file_citation]]
type OutputContentMessageTextAnnotationFilePath struct {
	Type   TypeMessageContentTypeOutputTextAnnotationsTypeEnum `json:"type"`
	FileId string                                              `json:"file_id"`
	Index  int                                                 `json:"file_index"`
}

func (t OutputContentMessageTextAnnotationFilePath) annotationType() TypeMessageContentTypeOutputTextAnnotationsTypeEnum {
	return TypeMessageContentTypeOutputTextAnnotationsTypeFilePath
}

// Used at /response/output[@type=message]/content[@type=output_text]/logprobs/*
type BaseTypeMessageContentTypeOutputTextLogprobsDef struct {
	Bytes   []byte `json:"bytes"`
	LogProb int    `json:"logprob"`
	Token   string `json:"token"`
}

// Used at /response/output[@type=message]/content[@type=output_text]/logprobs
type TypeMessageContentTypeOutputTextLogprobsDef struct {
	BaseTypeMessageContentTypeOutputTextLogprobsDef
	TopLogProbs []BaseTypeMessageContentTypeOutputTextLogprobsDef `json:"top_logprobs"`
}
