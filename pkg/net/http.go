package net

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"

	"github.com/stefan79/openai-driver/pkg/output"
)

type defaultHTTPClient struct {
	client    *http.Client
	apiKey    string
	baseUrl   string
	outputter output.Outputter
}

func NewHTTPClient(o output.Outputter, openApiKey, baseUrl string, proxy *string) (HTTPClient, error) {
	client := defaultHTTPClient{
		client:    &http.Client{},
		apiKey:    openApiKey,
		baseUrl:   baseUrl,
		outputter: o,
	}
	if proxy != nil {
		url, err := url.Parse(*proxy)
		if err != nil {
			return nil, err
		}
		o.VMessage(fmt.Sprintf("Using Proxy: %s/n", *proxy))
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
	c.outputter.VVMessage(fmt.Sprintf("Request: %s %s", httpReq.Method, httpReq.URL))
	resp, err := c.client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			// Log the error if closing the body fails
			log.Printf("error closing response body: %v", err)
		}
	}()
	c.outputter.VVMessage(fmt.Sprintf("Response: %d %s", resp.StatusCode, resp.Status))
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

func mapOpenAIToHttp(req *OpenAIRequest, baseUrl, apiKey string) (*http.Request, error) {
	httpReq, err := http.NewRequest(req.Method, baseUrl+req.Path, bytes.NewReader(req.Body))
	if err != nil {
		return nil, err
	}
	httpReq.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiKey))
	httpReq.Header.Set("Content-Type", "application/json")
	return httpReq, nil
}
