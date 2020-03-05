package v1

import (
	"context"
	"io"

	grpc "google.golang.org/grpc"
)

type TestRouterClient struct {
	elements []interface{}
}

func NewTestRouterClient(elements []interface{}) *TestRouterClient {
	return &TestRouterClient{elements: elements}
}

func (c *TestRouterClient) StreamMatches(ctx context.Context, in *RouterRequest, opts ...grpc.CallOption) (Router_StreamMatchesClient, error) {
	return &TestStreamMatchesClient{elements: c.elements}, nil
}

type TestStreamMatchesClient struct {
	elements []interface{}
	idx      int

	grpc.ClientStream
}

func (c *TestStreamMatchesClient) Recv() (*SearchMatch, error) {
	defer func() { c.idx++ }()

	if c.idx >= len(c.elements) {
		return nil, io.EOF
	}

	switch out := c.elements[c.idx].(type) {
	case *SearchMatch:
		return out, nil
	case error:
		return nil, out
	default:
		panic("unsupported typed")
	}
}
