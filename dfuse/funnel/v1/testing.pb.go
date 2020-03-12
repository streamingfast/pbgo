package funnel

import (
	"context"
	"io"

	grpc "google.golang.org/grpc"
	deos "github.com/dfuse-io/pbgo/dfuse/codecs/deos"
)

type TestFunnelClient struct {
	elements []*deos.Block
}

func NewTestFunnelClient(elements []*deos.Block) *TestFunnelClient {
	return &TestFunnelClient{elements: elements}
}

func (c *TestFunnelClient) StreamBlocks(ctx context.Context, in *StreamBlockRequest, opts ...grpc.CallOption) (Funnel_StreamBlocksClient, error) {
	return &TestStreamBlocksClient{elements: c.elements}, nil
}

type TestStreamBlocksClient struct {
	elements []*deos.Block
	idx      int

	grpc.ClientStream
}

func (c *TestStreamBlocksClient) Recv() (*deos.Block, error) {
	defer func() { c.idx++ }()

	if c.idx >= len(c.elements) {
		return nil, io.EOF
	}

	return c.elements[c.idx], nil
}
