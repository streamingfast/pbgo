// Copyright 2019 dfuse Platform Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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