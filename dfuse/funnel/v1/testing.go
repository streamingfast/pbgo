package funnel

import (
	"context"
	"io"

	"google.golang.org/grpc"
)

type TestFunnelClient struct {
	elements []*StreamBlockResponse
}

func NewTestFunnelClient(elements []*StreamBlockResponse) *TestFunnelClient {
	return &TestFunnelClient{elements: elements}
}

func (c *TestFunnelClient) StreamBlocks(ctx context.Context, in *StreamBlockRequest, opts ...grpc.CallOption) (Funnel_StreamBlocksClient, error) {
	return &TestStreamBlocksClient{elements: c.elements}, nil
}

type TestStreamBlocksClient struct {
	elements []*StreamBlockResponse
	idx      int

	grpc.ClientStream
}

func (c *TestStreamBlocksClient) Recv() (*StreamBlockResponse, error) {
	defer func() { c.idx++ }()

	if c.idx >= len(c.elements) {
		return nil, io.EOF
	}

	return c.elements[c.idx], nil
}