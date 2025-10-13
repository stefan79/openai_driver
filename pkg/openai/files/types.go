package files

import "github.com/stefan79/openai-driver/pkg/openai"

type Purpose string

const (
	PurposeFineTune   Purpose = "fine-tune"
	PurposeAssistants Purpose = "assistants"
	PurposeBatch      Purpose = "batch"
	PurposeUserData   Purpose = "user_data"
	PurposeVision     Purpose = "vision"
	PurposeEvals      Purpose = "evals"
)

type UploadRequest struct {
	Purpose  Purpose             `json:"purpose"`
	FileName string              `json:"file_name"`
	File     *openai.Base64Bytes `json:"file"`
}

type File struct {
	Id            string                 `json:"id"`
	Object        string                 `json:"object"`
	Bytes         int                    `json:"bytes"`
	CreatedAt     int64                  `json:"created_at"`
	Filename      string                 `json:"filename"`
	Purpose       Purpose                `json:"purpose"`
	Status        string                 `json:"status"`
	StatusDetails map[string]interface{} `json:"status_details,omitempty"`
}

type ListResponse struct {
	Object string `json:"object"`
	Data   []File `json:"data"`
}

type DeleteResponse struct {
	Id      string `json:"id"`
	Object  string `json:"object"`
	Deleted bool   `json:"deleted"`
}
