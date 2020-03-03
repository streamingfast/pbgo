package v1

import (
	context "context"
	"time"

	"github.com/eoscanada/dgrpc"
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
