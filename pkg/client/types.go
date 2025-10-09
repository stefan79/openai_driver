package client

import (
	"context"

	builderResp "driver/pkg/builder/response"
	openAIResp "driver/pkg/openai/response"
)

type StreamEvent struct {
}

// StreamReader represents a streaming response reader
type StreamReader interface {
	Next() (*StreamEvent, error)
	Close() error
}

type TypedResponse[T any] struct {
	*openAIResp.ResponsesResponse
	ParsedOutput T
}

type Client interface {
	Create(ctx context.Context, options ...builderResp.ResponseOption) (*openAIResp.ResponsesResponse, error)
	CreateStream(ctx context.Context, options ...builderResp.ResponseOption) (*StreamReader, error)
	Retrieve(ctx context.Context, responseId string) (*openAIResp.ResponsesResponse, error)
	Cancel(ctx context.Context, responseId string) error
}

type TypedClient[T any] interface {
	Create(ctx context.Context, options ...builderResp.ResponseOption) (*TypedResponse[T], error)
	Retrieve(ctx context.Context, responseId string) (*TypedResponse[T], error)
}
