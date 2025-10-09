package net

import (
	"bytes"
	"context"
	"driver/pkg/config"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type defaultHTTPClient struct {
	client  *http.Client
	apiKey  string
	baseUrl string
}

func NewHTTPClient(cfg *config.Config) (HTTPClient, error) {
	client := defaultHTTPClient{
		client: &http.Client{
			Timeout: cfg.TimeOut,
		},
		apiKey:  cfg.OpenaiApiKey,
		baseUrl: cfg.BaseUrl,
	}
	if cfg.Proxy != nil {
		url, err := url.Parse(*cfg.Proxy)
		if err != nil {
			return nil, err
		}
		fmt.Printf("Using proxy: %s\n", url)
		client.client.Transport = &http.Transport{
			Proxy: http.ProxyURL(url),
		}
	}
	return &client, nil
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
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error code: %d", resp.StatusCode)
	}

	headers := make(map[string]string, len(resp.Header))
	for k, v := range resp.Header {
		if len(v) > 0 {
			headers[k] = v[0] // Take the first value if multiple exist
		}
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return &OpenAIResponse{
		Headers:      headers,
		Body:         body,
		ResponseCode: resp.StatusCode,
	}, nil
}

func mapOpenAIToHttp(req *OpenAIRequest, baseUrl string, apiKey string) (*http.Request, error) {
	httpReq, err := http.NewRequest(req.Method, baseUrl+req.Path, bytes.NewReader(req.Body))
	if err != nil {
		return nil, err
	}
	httpReq.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiKey))
	httpReq.Header.Set("Content-Type", "application/json")
	return httpReq, nil
}
