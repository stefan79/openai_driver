package net

import (
	"context"
)

type OpenAIRequest struct {
	Path   string
	Method string
	Body   []byte
}

type OpenAIResponse struct {
	Body         []byte
	Headers      map[string]string
	ResponseCode int
}

type HTTPClient interface {
	Do(ctx context.Context, req *OpenAIRequest) (*OpenAIResponse, error)
}
