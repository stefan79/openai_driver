package files

import "github.com/stefan79/openai-driver/pkg/openai"

type UploadRequest struct {
	Purpose  string              `json:"purpose"`
	FileName string              `json:"file_name"`
	File     *openai.Base64Bytes `json:"file"`
}

type File struct {
	Id            string                 `json:"id"`
	Object        string                 `json:"object"`
	Bytes         int                    `json:"bytes"`
	CreatedAt     int64                  `json:"created_at"`
	Filename      string                 `json:"filename"`
	Purpose       string                 `json:"purpose"`
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
