package deth

import (
	"encoding/hex"
	"encoding/json"
	fmt "fmt"
	"math/big"
	"time"

	"github.com/eoscanada/jsonpb"
	"github.com/gogo/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
)

// TODO: We should probably memoize all fields that requires computation
//       like ID() and likes.

func (b *Block) ID() string {
	return hex.EncodeToString(b.Hash)
}

func (b *Block) Num() uint64 {
	return b.Number
}

func (b *Block) Time() (time.Time, error) {
	timestamp, err := ptypes.Timestamp(b.Header.Timestamp)
	if err != nil {
		return time.Time{}, fmt.Errorf("unable to turn google proto Timestamp into time.Time: %s", err)
	}

	return timestamp, nil
}

func (b *Block) MustTime() time.Time {
	timestamp, err := b.Time()
	if err != nil {
		panic(err)
	}

	return timestamp
}

func (b *Block) PreviousID() string {
	return hex.EncodeToString(b.Header.ParentHash)
}

// FIXME: This logic at some point is hard-coded and will need to be re-visited in regard
//        of the fork logic.
func (b *Block) LIBNum() uint64 {
	if b.Number <= 200 {
		return 1
	}

	return b.Number - 200
}

// func (b *Block) AsRef() bstream.BlockRef {
// 	return bstream.NewBlockRef(b.ID(), b.Number)
// }

func NewBigInt(in int64) *BigInt {
	return BigIntFromNative(big.NewInt(in))
}

func BigIntFromNative(in *big.Int) *BigInt {
	return &BigInt{Bytes: in.Bytes()}
}

func BigIntFromBytes(in []byte) *BigInt {
	return &BigInt{Bytes: in}
}

func (m *BigInt) Uint64() uint64 {
	i := new(big.Int).SetBytes(m.Bytes)
	return i.Uint64()
}

func (m *BigInt) Native() *big.Int {
	z := new(big.Int)
	z.SetBytes(m.Bytes)
	return z
}

func (m *BigInt) MarshalJSON() ([]byte, error) {
	return []byte(`"` + hex.EncodeToString(m.Bytes) + `"`), nil
}

func (m *BigInt) UnmarshalJSON(in []byte) (err error) {
	var s string
	err = json.Unmarshal(in, &s)
	if err != nil {
		return
	}

	m.Bytes, err = hex.DecodeString(s)
	return
}

func (m *BigInt) MarshalJSONPB(marshaler *jsonpb.Marshaler) ([]byte, error) {
	return m.MarshalJSON()
}

func (m *BigInt) UnmarshalJSONPB(unmarshaler *jsonpb.Unmarshaler, in []byte) (err error) {
	return m.UnmarshalJSON(in)
}

func BlockToBuffer(block *Block) ([]byte, error) {
	buf, err := proto.Marshal(block)
	if err != nil {
		return nil, err
	}
	return buf, nil
}

func MustBlockToBuffer(block *Block) []byte {
	buf, err := BlockToBuffer(block)
	if err != nil {
		panic(err)
	}
	return buf
}

func (trace *TransactionTrace) PopulateStateReverted() {
	// Calls are ordered by execution index. So the algo is quite simple.
	// We loop through the flat calls, at each call, if the parent is present
	// and reverted, the current call is reverted. Otherwise, if the current call
	// is failed, the state is reverted. In all other cases, we simply continue
	// our iteration loop.
	//
	// This works because we see the parent before its children, and since we
	// trickle down the state reverted value down the children, checking the parent
	// of a call will always tell us if the whole chain of parent/child should
	// be reverted
	//
	calls := trace.Calls
	for _, call := range trace.Calls {
		var parent *Call
		if call.ParentIndex > 0 {
			parent = calls[call.ParentIndex-1]
		}

		call.StateReverted = (parent != nil && parent.StateReverted) || call.StatusFailed
	}

	return
}
