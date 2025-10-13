package client

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/stefan79/openai-driver/pkg/openai/responses/response"
	"github.com/stefan79/openai-driver/pkg/output"

	"github.com/stefan79/openai-driver/pkg/openai/responses/request"

	"github.com/stefan79/openai-driver/pkg/builder"
	"github.com/stefan79/openai-driver/pkg/net"
)

type defaultClient struct {
	httpClient net.HTTPClient
}

func NewClient(o output.Outputter, openApiKey, baseUrl string, proxy *string) (Client, error) {
	httpClient, err := net.NewHTTPClient(o, openApiKey, baseUrl, proxy)
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
