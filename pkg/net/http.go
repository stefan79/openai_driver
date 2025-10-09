package net

import (
	"bytes"
	"context"
	"driver/pkg/config"
	"fmt"
	"net/http"
)

type defaultHTTPClient struct {
	client  *http.Client
	apiKey  string
	baseUrl string
}

func NewHTTPClient(cfg *config.Config) HTTPClient {
	return &defaultHTTPClient{
		client: &http.Client{
			Timeout: cfg.TimeOut,
		},
		apiKey:  cfg.OpenaiApiKey,
		baseUrl: cfg.BaseUrl,
	}
}

func (c *defaultHTTPClient) Do(ctx context.Context, req *OpenAIRequest) (*OpenAIResponse, error) {
	httpReq, err := mapOpenAIToHttp(req, c.baseUrl, c.apiKey)
	if err != nil {
		return nil, err
	}
	resp, err := c.client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error code: %d", resp.StatusCode)
	}
	return nil, fmt.Errorf("not implemented")
}

func mapOpenAIToHttp(req *OpenAIRequest, baseUrl string, apiKey string) (*http.Request, error) {
	httpReq, err := http.NewRequest(req.Method, baseUrl+req.Path, bytes.NewReader(req.Body))
	if err != nil {
		return nil, err
	}
	httpReq.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiKey))
	return httpReq, nil
}
