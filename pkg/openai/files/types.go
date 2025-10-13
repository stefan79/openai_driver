package files

type UploadRequest struct {
	Purpose  string
	FileName string
	FileData []byte
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
