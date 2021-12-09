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

package pbsearch

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
