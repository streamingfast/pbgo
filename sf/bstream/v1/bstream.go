package pbbstream

import (
	"time"
)

// GetFirehoseBlockID implements firecore.Block.
func (b *Block) GetFirehoseBlockID() string {
	return b.Id
}

// GetFirehoseBlockNumber implements firecore.Block.
func (b *Block) GetFirehoseBlockNumber() uint64 {
	return b.Number
}

// GetFirehoseBlockParentID implements firecore.Block.
func (b *Block) GetFirehoseBlockParentID() string {
	return b.PreviousId
}

// GetFirehoseBlockParentNumber implements firecore.Block.
func (b *Block) GetFirehoseBlockParentNumber() uint64 {
	panic("implement me")
}

// GetFirehoseBlockTime implements firecore.Block.
func (b *Block) GetFirehoseBlockTime() time.Time {
	return b.Timestamp.AsTime()
}

// GetFirehoseBlockVersion implements firecore.Block.
func (b *Block) GetFirehoseBlockVersion() int32 {
	panic("implement me, not sure this is the right value to return")
	return b.PayloadVersion
}
