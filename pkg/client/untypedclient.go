package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/stefan79/openai-driver/pkg/openai/responses/response"

	openAIFiles "github.com/stefan79/openai-driver/pkg/openai/files"
	"github.com/stefan79/openai-driver/pkg/openai/responses/request"

	"github.com/stefan79/openai-driver/pkg/builder"
	"github.com/stefan79/openai-driver/pkg/net"
)

type defaultClient struct {
	httpClient net.HTTPClient
}

func NewClient(openApiKey, baseUrl string, proxy *string) (Client, error) {
	httpClient, err := net.NewHTTPClient(openApiKey, baseUrl, proxy)
	if err != nil {
		return nil, err
	}
	return &defaultClient{
		httpClient: httpClient,
	}, nil
}

func (c *defaultClient) Create(ctx context.Context, model string, options ...builder.ResponsesOption) (*response.ResponseDef, error) {
	// Add a mapper function which creates the openAI Request from the ResponseRequest

	responsesRequest := request.RequestDef{
		Model: &model,
	}
	for _, option := range options {
		option(&responsesRequest)
	}
	httpRequestBody, err := json.Marshal(responsesRequest)
	if err != nil {
		return nil, err
	}
	request := net.OpenAIRequest{
		Path:   "/v1/responses",
		Method: "POST",
		Body:   httpRequestBody,
	}

	httpResponse, err := c.httpClient.Do(ctx, &request)
	if err != nil {
		return nil, err
	}

	response := &response.ResponseDef{}
	err = json.Unmarshal(httpResponse.Body, response)
	if err != nil {
		fmt.Printf("Unmarshalling error: %e\n", err)
		return nil, err
	}

	return response, nil
}

func (c *defaultClient) CreateStream(ctx context.Context, model string, options ...builder.ResponsesOption) (StreamReader, error) {
	return nil, fmt.Errorf("not implemented")
}

func (c *defaultClient) Retrieve(ctx context.Context, responseId string) (*response.ResponseDef, error) {
	return nil, fmt.Errorf("not implemented")
}

func (c *defaultClient) Cancel(ctx context.Context, responseId string) error {
	return fmt.Errorf("not implemented")
}

func (c *defaultClient) UploadFile(ctx context.Context, options ...builder.FilesOption) (*openAIFiles.File, error) {
	req := &openAIFiles.UploadRequest{}
	if err := builder.ApplyFilesOptions(req, options...); err != nil {
		return nil, err
	}
	if err := builder.ValidateFilesUploadRequest(req); err != nil {
		return nil, err
	}
	request := net.OpenAIRequest{
		Path:   "/v1/files",
		Method: "POST",
		FormValues: map[string]string{
			"purpose": string(req.Purpose),
		},
		FormFiles: []net.FormFile{
			{
				FieldName: "file",
				FileName:  req.FileName,
				File:      req.File,
			},
		},
	}
	httpResponse, err := c.httpClient.Do(ctx, &request)
	if err != nil {
		return nil, err
	}
	file := &openAIFiles.File{}
	if err := json.Unmarshal(httpResponse.Body, file); err != nil {
		return nil, err
	}
	return file, nil
}

func (c *defaultClient) ListFiles(ctx context.Context, purpose *openAIFiles.Purpose) (*openAIFiles.ListResponse, error) {
	path := "/v1/files"
	if purpose != nil {
		values := url.Values{}
		values.Set("purpose", string(*purpose))
		path = fmt.Sprintf("%s?%s", path, values.Encode())
	}
	request := net.OpenAIRequest{
		Path:   path,
		Method: "GET",
	}
	httpResponse, err := c.httpClient.Do(ctx, &request)
	if err != nil {
		return nil, err
	}
	list := &openAIFiles.ListResponse{}
	if err := json.Unmarshal(httpResponse.Body, list); err != nil {
		return nil, err
	}
	return list, nil
}

func (c *defaultClient) RetrieveFile(ctx context.Context, fileID string) (*openAIFiles.File, error) {
	request := net.OpenAIRequest{
		Path:   fmt.Sprintf("/v1/files/%s", fileID),
		Method: "GET",
	}
	httpResponse, err := c.httpClient.Do(ctx, &request)
	if err != nil {
		return nil, err
	}
	file := &openAIFiles.File{}
	if err := json.Unmarshal(httpResponse.Body, file); err != nil {
		return nil, err
	}
	return file, nil
}

func (c *defaultClient) DeleteFile(ctx context.Context, fileID string) (*openAIFiles.DeleteResponse, error) {
	request := net.OpenAIRequest{
		Path:   fmt.Sprintf("/v1/files/%s", fileID),
		Method: "DELETE",
	}
	httpResponse, err := c.httpClient.Do(ctx, &request)
	if err != nil {
		return nil, err
	}
	resp := &openAIFiles.DeleteResponse{}
	if err := json.Unmarshal(httpResponse.Body, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *defaultClient) DownloadFile(ctx context.Context, fileID string) ([]byte, error) {
	request := net.OpenAIRequest{
		Path:   fmt.Sprintf("/v1/files/%s/content", fileID),
		Method: "GET",
	}
	httpResponse, err := c.httpClient.Do(ctx, &request)
	if err != nil {
		return nil, err
	}
	return httpResponse.Body, nil
}
