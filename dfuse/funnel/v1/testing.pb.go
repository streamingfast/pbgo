package funnel

import (
	"context"
	"io"

	grpc "google.golang.org/grpc"
	deos "github.com/dfuse-io/pbgo/dfuse/codecs/deos"
)

type TestFunnelClient struct {
	elements []interface{}
}

func NewTestFunnelClient(elements []interface{}) *TestFunnelClient {
	return &TestFunnelClient{elements: elements}
}

func (c *TestFunnelClient) StreamBlocks(ctx context.Context, in *FunnelRequest, opts ...grpc.CallOption) (Funnel_StreamBlocksClient, error) {
	return &TestStreamBlocksClient{elements: c.elements}, nil
}

type TestStreamBlocksClient struct {
	elements []interface{}
	idx      int

	grpc.ClientStream
}

func (c *TestStreamBlocksClient) Recv() (*deos.Block, error) {
	defer func() { c.idx++ }()

	if c.idx >= len(c.elements) {
		return nil, io.EOF
	}

	switch out := c.elements[c.idx].(type) {
	case *deos.Block:
		return out, nil
	case error:
		return nil, out
	default:
		panic("unsupported typed")
	}
}
