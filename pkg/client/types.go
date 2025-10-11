package client

import (
	"context"

	"driver/pkg/builder"
	"driver/pkg/openai/responses/resp"
)

type StreamEvent struct {
}

// StreamReader represents a streaming response reader
type StreamReader interface {
	Next() (*StreamEvent, error)
	Close() error
}

type TypedResponse[T any] struct {
	*resp.ResponseDef
	ParsedOutput T
}

type Client interface {
	Create(ctx context.Context, model string, options ...builder.ResponsesOption) (*resp.ResponseDef, error)
	CreateStream(ctx context.Context, model string, options ...builder.ResponsesOption) (*StreamReader, error)
	Retrieve(ctx context.Context, responseId string) (*resp.ResponseDef, error)
	Cancel(ctx context.Context, responseId string) error
}

type TypedClient[T any] interface {
	Create(ctx context.Context, model string, options ...builder.ResponsesOption) (*TypedResponse[T], error)
	Retrieve(ctx context.Context, responseId string) (*TypedResponse[T], error)
}
