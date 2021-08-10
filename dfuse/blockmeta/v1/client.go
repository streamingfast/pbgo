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

package pbblockmeta

import (
	context "context"
	fmt "fmt"
	"time"

	"github.com/streamingfast/dgrpc"
	"github.com/golang/protobuf/ptypes"
	tspb "github.com/golang/protobuf/ptypes/timestamp"
	grpc "google.golang.org/grpc"
)

type Client struct {
	conn *grpc.ClientConn
}

func NewClient(addr string) (*Client, error) {
	conn, err := dgrpc.NewInternalClient(addr)
	if err != nil {
		return nil, err
	}

	return &Client{
		conn: conn,
	}, nil
}

func (b *Client) timeToIDClient() TimeToIDClient {
	return NewTimeToIDClient(b.conn)
}

func (b *Client) blockIDClient() BlockIDClient {
	return NewBlockIDClient(b.conn)
}

func (b *Client) ChainDiscriminatorClient() ChainDiscriminatorClient {
	return NewChainDiscriminatorClient(b.conn)
}

func (b *Client) BlockNumToID(ctx context.Context, blockNum uint64) (*BlockIDResponse, error) {
	return b.blockIDClient().NumToID(ctx, &NumToIDRequest{BlockNum: blockNum})
}

func (b *Client) LIBID(ctx context.Context) (*BlockIDResponse, error) {
	return b.blockIDClient().LIBID(ctx, &LIBRequest{})
}

func (b *Client) BlockAt(ctx context.Context, t time.Time) (*BlockResponse, error) {
	return b.timeToIDClient().At(ctx, &TimeRequest{
		Time: TimestampProto(t),
	})
}

func (b *Client) BlockBefore(ctx context.Context, t time.Time, inclusive bool) (*BlockResponse, error) {
	return b.timeToIDClient().Before(ctx, &RelativeTimeRequest{
		Time:      TimestampProto(t),
		Inclusive: inclusive,
	})
}

func (b *Client) BlockAfter(ctx context.Context, t time.Time, inclusive bool) (*BlockResponse, error) {
	return b.timeToIDClient().After(ctx, &RelativeTimeRequest{
		Time:      TimestampProto(t),
		Inclusive: inclusive,
	})
}

// TODO: This will eventually move to our own libs.. we'll need that everywhere.
func Timestamp(ts *tspb.Timestamp) time.Time {
	t, _ := ptypes.Timestamp(ts)
	return t
}

func TimestampProto(t time.Time) *tspb.Timestamp {
	out, _ := ptypes.TimestampProto(t)
	return out
}

// StartBlockResolver will try to return same blocknum as target, with
// irreversible ID set to that block itself. If the requested block is
// NOT considered irreversible, it returns an error and you will have to
// use another method.
func StartBlockResolver(cli BlockIDClient) func(ctx context.Context, targetBlockNum uint64) (uint64, string, error) {
	return func(ctx context.Context, targetBlockNum uint64) (uint64, string, error) {
		if targetBlockNum <= 2 {
			return targetBlockNum, "", nil
		}
		idResp, err := cli.NumToID(ctx, &NumToIDRequest{BlockNum: targetBlockNum}, grpc.WaitForReady(false))
		if err != nil {
			return 0, "", err
		}

		if idResp.Irreversible {
			return targetBlockNum, idResp.Id, nil
		}
		return 0, "", fmt.Errorf("could not find irreversible id for targetBlockNum:%d", targetBlockNum)
	}
}
