package client

import (
	"context"
	builderResp "driver/pkg/builder/response"
	"driver/pkg/config"
	"driver/pkg/net"
	openAIResp "driver/pkg/openai/response"
)

type defaultClient struct {
	httpClient net.HTTPClient
}

func NewClient(cfg *config.Config) *defaultClient {
	return &defaultClient{
		httpClient: net.NewHTTPClient(cfg),
	}
}

func (c *defaultClient) Create(ctx context.Context, options ...builderResp.ResponseOption) (*openAIResp.ResponsesResponse, error) {
	//Add a mapper function which creates the openAI Request from the ResponseRequest

	requestBody := openAIResp.ResponsesRequest{}
	for _, option := range options {
		option(&requestBody)
	}
	request := net.OpenAIRequest{
		Path: "/v1/chat/completions",
	}

	_, err := c.httpClient.Do(ctx, &request)
	if err != nil {
		return nil, err
	}

	return &openAIResp.ResponsesResponse{}, nil

}
