package net

import (
	"context"

	"github.com/stefan79/openai-driver/pkg/openai"
)

type OpenAIRequest struct {
	Path       string
	Method     string
	Body       []byte
	Headers    map[string]string
	FormValues map[string]string
	FormFiles  []FormFile
}

type OpenAIResponse struct {
	Body         []byte
	Headers      map[string]string
	ResponseCode int
}

type HTTPClient interface {
	Do(ctx context.Context, req *OpenAIRequest) (*OpenAIResponse, error)
}

type FormFile struct {
	FieldName string
	FileName  string
	File      *openai.Base64Bytes
}
