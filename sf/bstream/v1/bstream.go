package pbbstream

import (
	"fmt"
	"time"
)

func (b *Block) Time() time.Time {
	if b == nil {
		return time.Time{}
	}
	if err := b.Timestamp.CheckValid(); err != nil {
		panic(fmt.Errorf("invalid timestamp: %w", err))
	}

	return b.Timestamp.AsTime()
}

func (b *Block) LIBNum() uint64 {
	if b == nil {
		return 0
	}

	return b.LibNum
}

type basicBlockRef struct {
	id  string
	num uint64
}

func (e *basicBlockRef) Num() uint64 { return e.num }
func (e *basicBlockRef) ID() string  { return e.id }
func (e *basicBlockRef) String() string {
	if e == nil || (e.id == "" && e.num == 0) {
		return "Block <nil>"
	}

	return fmt.Sprintf("#%d (%s)", e.Num(), e.ID())
}

func (b *Block) AsRef() *basicBlockRef {
	if b == nil {
		return &basicBlockRef{"", 0}
	}

	return &basicBlockRef{b.Id, b.Number}
}

func (b *Block) PreviousRef() *basicBlockRef {
	if b == nil || b.PreviousNum == 0 || b.PreviousId == "" {
		return &basicBlockRef{"", 0}
	}
	return &basicBlockRef{b.PreviousId, b.PreviousNum}
}

func (b *Block) PreviousID() string {
	return b.PreviousRef().ID()
}

func (b *Block) GetFirehoseBlockID() string     { return b.Id }
func (b *Block) GetFirehoseBlockNumber() uint64 { return b.Number }

// GetFirehoseBlockParentID implements firecore.Block.
func (b *Block) GetFirehoseBlockParentID() string {
	return b.PreviousId
}

// GetFirehoseBlockParentNumber implements firecore.Block.
func (b *Block) GetFirehoseBlockParentNumber() uint64 {
	return b.PreviousNum
}

// GetFirehoseBlockTime implements firecore.Block.
func (b *Block) GetFirehoseBlockTime() time.Time {
	return b.Timestamp.AsTime()
}

// GetFirehoseBlockVersion implements firecore.Block.
func (b *Block) GetFirehoseBlockVersion() int32 {
	panic("implement me, not sure this is the right value to return")
}
