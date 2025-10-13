package request

import "github.com/stefan79/openai-driver/pkg/openai"

type InputContentTypeEnum string

const (
	InputContentTypeText  InputContentTypeEnum = "input_text"
	InputContentTypeFile  InputContentTypeEnum = "input_file"
	InputContentTypeImage InputContentTypeEnum = "input_image"
	InputContentTypeAudio InputContentTypeEnum = "input_audio"
)

type InputContentTypeImageImageDetailEnum string

const (
	InputContentTypeImageImageDetailAuto InputContentTypeImageImageDetailEnum = "auto"
	InputContentTypeImageImageDetailHigh InputContentTypeImageImageDetailEnum = "high"
	InputContentTypeImageImageDetailLow  InputContentTypeImageImageDetailEnum = "low"
)

type InputContentTypeAudioAudioInputFormatEnum string

const (
	InputContentTypeAudioAudioInputFormatMP3 InputContentTypeAudioAudioInputFormatEnum = "mp3"
	InputContentTypeAudioAudioInputFormatWAV InputContentTypeAudioAudioInputFormatEnum = "wav"
)

// Used at /request/input
type InputDef struct {
	Role    openai.Role        `json:"role"`
	Content []BaseInputContent `json:"content"`
}

type BaseInputContent interface {
	ContentType() InputContentTypeEnum
}

// TextContent represents text input
type InputContentTypeTextDef struct {
	Type InputContentTypeEnum `json:"type"`
	Text string               `json:"text"`
}

func (t InputContentTypeTextDef) ContentType() InputContentTypeEnum {
	return InputContentTypeText
}

// FileContent represents file input
type InputContentTypeFileDef struct {
	Type     InputContentTypeEnum `json:"type"`
	FileData *openai.Base64Bytes  `json:"file_data,omitempty"`
	FileId   *string              `json:"file_id,omitempty"`
	FileUrl  *string              `json:"file_url,omitempty"`
	FileName *string              `json:"filename,omitempty"`
}

func (t InputContentTypeFileDef) ContentType() InputContentTypeEnum {
	return InputContentTypeFile
}

// ImageContent represents image input
type InputContentTypeImageDef struct {
	Type        InputContentTypeEnum                 `json:"type"`
	ImageDetail InputContentTypeImageImageDetailEnum `json:"detail"`
	FileId      *string                              `json:"file_id,omitempty"`
	ImageUrl    *string                              `json:"image_url,omitempty"`
}

func (t InputContentTypeImageDef) ContentType() InputContentTypeEnum {
	return InputContentTypeImage
}

type InputContentTypeAudioAudioInputDef struct {
	Data   []byte                                    `json:"data"`
	Format InputContentTypeAudioAudioInputFormatEnum `json:"format"`
}

// AudioContent represents audio input
type InputContentTypeAudioDef struct {
	Type       InputContentTypeEnum               `json:"type"`
	AudioInput InputContentTypeAudioAudioInputDef `json:"audio_input"`
}

func (t InputContentTypeAudioDef) ContentType() InputContentTypeEnum {
	return InputContentTypeAudio
}
