// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.0
// source: sf/search/v1/search.proto

package pbsearch

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type RouterRequest_Mode int32

const (
	RouterRequest_STREAMING RouterRequest_Mode = 0
	RouterRequest_PAGINATED RouterRequest_Mode = 1
)

// Enum value maps for RouterRequest_Mode.
var (
	RouterRequest_Mode_name = map[int32]string{
		0: "STREAMING",
		1: "PAGINATED",
	}
	RouterRequest_Mode_value = map[string]int32{
		"STREAMING": 0,
		"PAGINATED": 1,
	}
)

func (x RouterRequest_Mode) Enum() *RouterRequest_Mode {
	p := new(RouterRequest_Mode)
	*p = x
	return p
}

func (x RouterRequest_Mode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RouterRequest_Mode) Descriptor() protoreflect.EnumDescriptor {
	return file_sf_search_v1_search_proto_enumTypes[0].Descriptor()
}

func (RouterRequest_Mode) Type() protoreflect.EnumType {
	return &file_sf_search_v1_search_proto_enumTypes[0]
}

func (x RouterRequest_Mode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RouterRequest_Mode.Descriptor instead.
func (RouterRequest_Mode) EnumDescriptor() ([]byte, []int) {
	return file_sf_search_v1_search_proto_rawDescGZIP(), []int{3, 0}
}

type ForkResolveRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Query           string      `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	ForkedBlockRefs []*BlockRef `protobuf:"bytes,2,rep,name=forkedBlockRefs,proto3" json:"forkedBlockRefs,omitempty"`
}

func (x *ForkResolveRequest) Reset() {
	*x = ForkResolveRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sf_search_v1_search_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ForkResolveRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ForkResolveRequest) ProtoMessage() {}

func (x *ForkResolveRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sf_search_v1_search_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ForkResolveRequest.ProtoReflect.Descriptor instead.
func (*ForkResolveRequest) Descriptor() ([]byte, []int) {
	return file_sf_search_v1_search_proto_rawDescGZIP(), []int{0}
}

func (x *ForkResolveRequest) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

func (x *ForkResolveRequest) GetForkedBlockRefs() []*BlockRef {
	if x != nil {
		return x.ForkedBlockRefs
	}
	return nil
}

type BlockRef struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlockNum uint64 `protobuf:"varint,1,opt,name=blockNum,proto3" json:"blockNum,omitempty"`
	BlockID  string `protobuf:"bytes,2,opt,name=blockID,proto3" json:"blockID,omitempty"`
}

func (x *BlockRef) Reset() {
	*x = BlockRef{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sf_search_v1_search_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockRef) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockRef) ProtoMessage() {}

func (x *BlockRef) ProtoReflect() protoreflect.Message {
	mi := &file_sf_search_v1_search_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockRef.ProtoReflect.Descriptor instead.
func (*BlockRef) Descriptor() ([]byte, []int) {
	return file_sf_search_v1_search_proto_rawDescGZIP(), []int{1}
}

func (x *BlockRef) GetBlockNum() uint64 {
	if x != nil {
		return x.BlockNum
	}
	return 0
}

func (x *BlockRef) GetBlockID() string {
	if x != nil {
		return x.BlockID
	}
	return ""
}

type BackendRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Query        string `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	LowBlockNum  uint64 `protobuf:"varint,2,opt,name=lowBlockNum,proto3" json:"lowBlockNum,omitempty"`   // Inclusive
	HighBlockNum uint64 `protobuf:"varint,3,opt,name=highBlockNum,proto3" json:"highBlockNum,omitempty"` // Inclusive
	Descending   bool   `protobuf:"varint,4,opt,name=descending,proto3" json:"descending,omitempty"`
	// More specific to the live / reversible backend
	WithReversible       bool   `protobuf:"varint,5,opt,name=withReversible,proto3" json:"withReversible,omitempty"`         // Fails on Archive
	LiveMarkerInterval   uint64 `protobuf:"varint,6,opt,name=liveMarkerInterval,proto3" json:"liveMarkerInterval,omitempty"` // in blocks
	StopAtVirtualHead    bool   `protobuf:"varint,7,opt,name=stopAtVirtualHead,proto3" json:"stopAtVirtualHead,omitempty"`
	NavigateFromBlockID  string `protobuf:"bytes,8,opt,name=navigateFromBlockID,proto3" json:"navigateFromBlockID,omitempty"`    // When a cursor was provided, which had a HEAD block ID
	NavigateFromBlockNum uint64 `protobuf:"varint,9,opt,name=navigateFromBlockNum,proto3" json:"navigateFromBlockNum,omitempty"` // When a cursor was provided, which had a HEAD block ID
}

func (x *BackendRequest) Reset() {
	*x = BackendRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sf_search_v1_search_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BackendRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BackendRequest) ProtoMessage() {}

func (x *BackendRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sf_search_v1_search_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BackendRequest.ProtoReflect.Descriptor instead.
func (*BackendRequest) Descriptor() ([]byte, []int) {
	return file_sf_search_v1_search_proto_rawDescGZIP(), []int{2}
}

func (x *BackendRequest) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

func (x *BackendRequest) GetLowBlockNum() uint64 {
	if x != nil {
		return x.LowBlockNum
	}
	return 0
}

func (x *BackendRequest) GetHighBlockNum() uint64 {
	if x != nil {
		return x.HighBlockNum
	}
	return 0
}

func (x *BackendRequest) GetDescending() bool {
	if x != nil {
		return x.Descending
	}
	return false
}

func (x *BackendRequest) GetWithReversible() bool {
	if x != nil {
		return x.WithReversible
	}
	return false
}

func (x *BackendRequest) GetLiveMarkerInterval() uint64 {
	if x != nil {
		return x.LiveMarkerInterval
	}
	return 0
}

func (x *BackendRequest) GetStopAtVirtualHead() bool {
	if x != nil {
		return x.StopAtVirtualHead
	}
	return false
}

func (x *BackendRequest) GetNavigateFromBlockID() string {
	if x != nil {
		return x.NavigateFromBlockID
	}
	return ""
}

func (x *BackendRequest) GetNavigateFromBlockNum() uint64 {
	if x != nil {
		return x.NavigateFromBlockNum
	}
	return 0
}

type RouterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Query              string             `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	LowBlockNum        int64              `protobuf:"varint,2,opt,name=lowBlockNum,proto3" json:"lowBlockNum,omitempty"`   // Inclusive
	HighBlockNum       int64              `protobuf:"varint,3,opt,name=highBlockNum,proto3" json:"highBlockNum,omitempty"` // Inclusive
	LowBlockUnbounded  bool               `protobuf:"varint,13,opt,name=lowBlockUnbounded,proto3" json:"lowBlockUnbounded,omitempty"`
	HighBlockUnbounded bool               `protobuf:"varint,14,opt,name=highBlockUnbounded,proto3" json:"highBlockUnbounded,omitempty"`
	Descending         bool               `protobuf:"varint,4,opt,name=descending,proto3" json:"descending,omitempty"`
	Cursor             string             `protobuf:"bytes,5,opt,name=cursor,proto3" json:"cursor,omitempty"`
	Limit              int64              `protobuf:"varint,6,opt,name=limit,proto3" json:"limit,omitempty"`
	WithReversible     bool               `protobuf:"varint,7,opt,name=withReversible,proto3" json:"withReversible,omitempty"` // Fails on Archive
	Mode               RouterRequest_Mode `protobuf:"varint,8,opt,name=mode,proto3,enum=sf.search.v1.RouterRequest_Mode" json:"mode,omitempty"`
	// Legacy boundaries, overrides `lowBlockNum` and `highBlockNum`.
	UseLegacyBoundaries bool   `protobuf:"varint,9,opt,name=useLegacyBoundaries,proto3" json:"useLegacyBoundaries,omitempty"`
	StartBlock          uint64 `protobuf:"varint,10,opt,name=startBlock,proto3" json:"startBlock,omitempty"`
	BlockCount          uint64 `protobuf:"varint,11,opt,name=blockCount,proto3" json:"blockCount,omitempty"`
	LiveMarkerInterval  uint64 `protobuf:"varint,12,opt,name=liveMarkerInterval,proto3" json:"liveMarkerInterval,omitempty"` // in blocks
}

func (x *RouterRequest) Reset() {
	*x = RouterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sf_search_v1_search_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RouterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RouterRequest) ProtoMessage() {}

func (x *RouterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sf_search_v1_search_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RouterRequest.ProtoReflect.Descriptor instead.
func (*RouterRequest) Descriptor() ([]byte, []int) {
	return file_sf_search_v1_search_proto_rawDescGZIP(), []int{3}
}

func (x *RouterRequest) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

func (x *RouterRequest) GetLowBlockNum() int64 {
	if x != nil {
		return x.LowBlockNum
	}
	return 0
}

func (x *RouterRequest) GetHighBlockNum() int64 {
	if x != nil {
		return x.HighBlockNum
	}
	return 0
}

func (x *RouterRequest) GetLowBlockUnbounded() bool {
	if x != nil {
		return x.LowBlockUnbounded
	}
	return false
}

func (x *RouterRequest) GetHighBlockUnbounded() bool {
	if x != nil {
		return x.HighBlockUnbounded
	}
	return false
}

func (x *RouterRequest) GetDescending() bool {
	if x != nil {
		return x.Descending
	}
	return false
}

func (x *RouterRequest) GetCursor() string {
	if x != nil {
		return x.Cursor
	}
	return ""
}

func (x *RouterRequest) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *RouterRequest) GetWithReversible() bool {
	if x != nil {
		return x.WithReversible
	}
	return false
}

func (x *RouterRequest) GetMode() RouterRequest_Mode {
	if x != nil {
		return x.Mode
	}
	return RouterRequest_STREAMING
}

func (x *RouterRequest) GetUseLegacyBoundaries() bool {
	if x != nil {
		return x.UseLegacyBoundaries
	}
	return false
}

func (x *RouterRequest) GetStartBlock() uint64 {
	if x != nil {
		return x.StartBlock
	}
	return 0
}

func (x *RouterRequest) GetBlockCount() uint64 {
	if x != nil {
		return x.BlockCount
	}
	return 0
}

func (x *RouterRequest) GetLiveMarkerInterval() uint64 {
	if x != nil {
		return x.LiveMarkerInterval
	}
	return 0
}

type SearchMatch struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TrxIdPrefix string `protobuf:"bytes,1,opt,name=trxIdPrefix,proto3" json:"trxIdPrefix,omitempty"`
	BlockNum    uint64 `protobuf:"varint,2,opt,name=blockNum,proto3" json:"blockNum,omitempty"`
	Index       uint64 `protobuf:"varint,3,opt,name=index,proto3" json:"index,omitempty"`
	Cursor      string `protobuf:"bytes,4,opt,name=cursor,proto3" json:"cursor,omitempty"`
	// Holds chain specific Protobuf message, this usually contains informations about how the
	// document are indexed (for example at the action level on EOSIO). Check out chain specific
	// repository for exact possible message.
	ChainSpecific *anypb.Any `protobuf:"bytes,13,opt,name=chainSpecific,proto3" json:"chainSpecific,omitempty"`
	Undo          bool       `protobuf:"varint,25,opt,name=undo,proto3" json:"undo,omitempty"`
	IrrBlockNum   uint64     `protobuf:"varint,26,opt,name=irrBlockNum,proto3" json:"irrBlockNum,omitempty"`
}

func (x *SearchMatch) Reset() {
	*x = SearchMatch{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sf_search_v1_search_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchMatch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchMatch) ProtoMessage() {}

func (x *SearchMatch) ProtoReflect() protoreflect.Message {
	mi := &file_sf_search_v1_search_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchMatch.ProtoReflect.Descriptor instead.
func (*SearchMatch) Descriptor() ([]byte, []int) {
	return file_sf_search_v1_search_proto_rawDescGZIP(), []int{4}
}

func (x *SearchMatch) GetTrxIdPrefix() string {
	if x != nil {
		return x.TrxIdPrefix
	}
	return ""
}

func (x *SearchMatch) GetBlockNum() uint64 {
	if x != nil {
		return x.BlockNum
	}
	return 0
}

func (x *SearchMatch) GetIndex() uint64 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *SearchMatch) GetCursor() string {
	if x != nil {
		return x.Cursor
	}
	return ""
}

func (x *SearchMatch) GetChainSpecific() *anypb.Any {
	if x != nil {
		return x.ChainSpecific
	}
	return nil
}

func (x *SearchMatch) GetUndo() bool {
	if x != nil {
		return x.Undo
	}
	return false
}

func (x *SearchMatch) GetIrrBlockNum() uint64 {
	if x != nil {
		return x.IrrBlockNum
	}
	return 0
}

var File_sf_search_v1_search_proto protoreflect.FileDescriptor

var file_sf_search_v1_search_proto_rawDesc = []byte{
	0x0a, 0x19, 0x73, 0x66, 0x2f, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2f, 0x76, 0x31, 0x2f, 0x73,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x73, 0x66, 0x2e,
	0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2e, 0x76, 0x31, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6c, 0x0a, 0x12, 0x46, 0x6f, 0x72, 0x6b, 0x52, 0x65, 0x73, 0x6f,
	0x6c, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x71, 0x75,
	0x65, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79,
	0x12, 0x40, 0x0a, 0x0f, 0x66, 0x6f, 0x72, 0x6b, 0x65, 0x64, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52,
	0x65, 0x66, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x73, 0x66, 0x2e, 0x73,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65,
	0x66, 0x52, 0x0f, 0x66, 0x6f, 0x72, 0x6b, 0x65, 0x64, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65,
	0x66, 0x73, 0x22, 0x40, 0x0a, 0x08, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x66, 0x12, 0x1a,
	0x0a, 0x08, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x4e, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x08, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x4e, 0x75, 0x6d, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x6c,
	0x6f, 0x63, 0x6b, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x62, 0x6c, 0x6f,
	0x63, 0x6b, 0x49, 0x44, 0x22, 0xf8, 0x02, 0x0a, 0x0e, 0x42, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x20, 0x0a,
	0x0b, 0x6c, 0x6f, 0x77, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x4e, 0x75, 0x6d, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x0b, 0x6c, 0x6f, 0x77, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x4e, 0x75, 0x6d, 0x12,
	0x22, 0x0a, 0x0c, 0x68, 0x69, 0x67, 0x68, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x4e, 0x75, 0x6d, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0c, 0x68, 0x69, 0x67, 0x68, 0x42, 0x6c, 0x6f, 0x63, 0x6b,
	0x4e, 0x75, 0x6d, 0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x65, 0x73, 0x63, 0x65, 0x6e, 0x64, 0x69, 0x6e,
	0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x64, 0x65, 0x73, 0x63, 0x65, 0x6e, 0x64,
	0x69, 0x6e, 0x67, 0x12, 0x26, 0x0a, 0x0e, 0x77, 0x69, 0x74, 0x68, 0x52, 0x65, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x62, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x77, 0x69, 0x74,
	0x68, 0x52, 0x65, 0x76, 0x65, 0x72, 0x73, 0x69, 0x62, 0x6c, 0x65, 0x12, 0x2e, 0x0a, 0x12, 0x6c,
	0x69, 0x76, 0x65, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x72, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61,
	0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x12, 0x6c, 0x69, 0x76, 0x65, 0x4d, 0x61, 0x72,
	0x6b, 0x65, 0x72, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x12, 0x2c, 0x0a, 0x11, 0x73,
	0x74, 0x6f, 0x70, 0x41, 0x74, 0x56, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x48, 0x65, 0x61, 0x64,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x11, 0x73, 0x74, 0x6f, 0x70, 0x41, 0x74, 0x56, 0x69,
	0x72, 0x74, 0x75, 0x61, 0x6c, 0x48, 0x65, 0x61, 0x64, 0x12, 0x30, 0x0a, 0x13, 0x6e, 0x61, 0x76,
	0x69, 0x67, 0x61, 0x74, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x44,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x6e, 0x61, 0x76, 0x69, 0x67, 0x61, 0x74, 0x65,
	0x46, 0x72, 0x6f, 0x6d, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x44, 0x12, 0x32, 0x0a, 0x14, 0x6e,
	0x61, 0x76, 0x69, 0x67, 0x61, 0x74, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x42, 0x6c, 0x6f, 0x63, 0x6b,
	0x4e, 0x75, 0x6d, 0x18, 0x09, 0x20, 0x01, 0x28, 0x04, 0x52, 0x14, 0x6e, 0x61, 0x76, 0x69, 0x67,
	0x61, 0x74, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x4e, 0x75, 0x6d, 0x22,
	0xbd, 0x04, 0x0a, 0x0d, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x20, 0x0a, 0x0b, 0x6c, 0x6f, 0x77, 0x42, 0x6c,
	0x6f, 0x63, 0x6b, 0x4e, 0x75, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x6c, 0x6f,
	0x77, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x4e, 0x75, 0x6d, 0x12, 0x22, 0x0a, 0x0c, 0x68, 0x69, 0x67,
	0x68, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x4e, 0x75, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0c, 0x68, 0x69, 0x67, 0x68, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x4e, 0x75, 0x6d, 0x12, 0x2c, 0x0a,
	0x11, 0x6c, 0x6f, 0x77, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x55, 0x6e, 0x62, 0x6f, 0x75, 0x6e, 0x64,
	0x65, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x08, 0x52, 0x11, 0x6c, 0x6f, 0x77, 0x42, 0x6c, 0x6f,
	0x63, 0x6b, 0x55, 0x6e, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x65, 0x64, 0x12, 0x2e, 0x0a, 0x12, 0x68,
	0x69, 0x67, 0x68, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x55, 0x6e, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x65,
	0x64, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x08, 0x52, 0x12, 0x68, 0x69, 0x67, 0x68, 0x42, 0x6c, 0x6f,
	0x63, 0x6b, 0x55, 0x6e, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x65, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x64,
	0x65, 0x73, 0x63, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0a, 0x64, 0x65, 0x73, 0x63, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x16, 0x0a, 0x06, 0x63,
	0x75, 0x72, 0x73, 0x6f, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x75, 0x72,
	0x73, 0x6f, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x26, 0x0a, 0x0e, 0x77, 0x69, 0x74,
	0x68, 0x52, 0x65, 0x76, 0x65, 0x72, 0x73, 0x69, 0x62, 0x6c, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0e, 0x77, 0x69, 0x74, 0x68, 0x52, 0x65, 0x76, 0x65, 0x72, 0x73, 0x69, 0x62, 0x6c,
	0x65, 0x12, 0x34, 0x0a, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x20, 0x2e, 0x73, 0x66, 0x2e, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x52,
	0x6f, 0x75, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x4d, 0x6f, 0x64,
	0x65, 0x52, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x12, 0x30, 0x0a, 0x13, 0x75, 0x73, 0x65, 0x4c, 0x65,
	0x67, 0x61, 0x63, 0x79, 0x42, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x72, 0x69, 0x65, 0x73, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x13, 0x75, 0x73, 0x65, 0x4c, 0x65, 0x67, 0x61, 0x63, 0x79, 0x42,
	0x6f, 0x75, 0x6e, 0x64, 0x61, 0x72, 0x69, 0x65, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x1e, 0x0a, 0x0a, 0x62, 0x6c, 0x6f,
	0x63, 0x6b, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x62,
	0x6c, 0x6f, 0x63, 0x6b, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2e, 0x0a, 0x12, 0x6c, 0x69, 0x76,
	0x65, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x72, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x04, 0x52, 0x12, 0x6c, 0x69, 0x76, 0x65, 0x4d, 0x61, 0x72, 0x6b, 0x65,
	0x72, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x22, 0x24, 0x0a, 0x04, 0x4d, 0x6f, 0x64,
	0x65, 0x12, 0x0d, 0x0a, 0x09, 0x53, 0x54, 0x52, 0x45, 0x41, 0x4d, 0x49, 0x4e, 0x47, 0x10, 0x00,
	0x12, 0x0d, 0x0a, 0x09, 0x50, 0x41, 0x47, 0x49, 0x4e, 0x41, 0x54, 0x45, 0x44, 0x10, 0x01, 0x22,
	0xf7, 0x01, 0x0a, 0x0b, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x12,
	0x20, 0x0a, 0x0b, 0x74, 0x72, 0x78, 0x49, 0x64, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x74, 0x72, 0x78, 0x49, 0x64, 0x50, 0x72, 0x65, 0x66, 0x69,
	0x78, 0x12, 0x1a, 0x0a, 0x08, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x4e, 0x75, 0x6d, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x08, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x4e, 0x75, 0x6d, 0x12, 0x14, 0x0a,
	0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x69, 0x6e,
	0x64, 0x65, 0x78, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x12, 0x3a, 0x0a, 0x0d, 0x63,
	0x68, 0x61, 0x69, 0x6e, 0x53, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x18, 0x0d, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x0d, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x53,
	0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x6e, 0x64, 0x6f, 0x18,
	0x19, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x75, 0x6e, 0x64, 0x6f, 0x12, 0x20, 0x0a, 0x0b, 0x69,
	0x72, 0x72, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x4e, 0x75, 0x6d, 0x18, 0x1a, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x0b, 0x69, 0x72, 0x72, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x4e, 0x75, 0x6d, 0x4a, 0x04, 0x08,
	0x0a, 0x10, 0x0c, 0x4a, 0x04, 0x08, 0x0e, 0x10, 0x19, 0x32, 0x55, 0x0a, 0x07, 0x42, 0x61, 0x63,
	0x6b, 0x65, 0x6e, 0x64, 0x12, 0x4a, 0x0a, 0x0d, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x4d, 0x61,
	0x74, 0x63, 0x68, 0x65, 0x73, 0x12, 0x1c, 0x2e, 0x73, 0x66, 0x2e, 0x73, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x73, 0x66, 0x2e, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x30, 0x01,
	0x32, 0x53, 0x0a, 0x06, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x12, 0x49, 0x0a, 0x0d, 0x53, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x12, 0x1b, 0x2e, 0x73, 0x66,
	0x2e, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x73, 0x66, 0x2e, 0x73, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4d, 0x61,
	0x74, 0x63, 0x68, 0x30, 0x01, 0x32, 0x62, 0x0a, 0x0c, 0x46, 0x6f, 0x72, 0x6b, 0x52, 0x65, 0x73,
	0x6f, 0x6c, 0x76, 0x65, 0x72, 0x12, 0x52, 0x0a, 0x11, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x55,
	0x6e, 0x64, 0x6f, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x12, 0x20, 0x2e, 0x73, 0x66, 0x2e,
	0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x6f, 0x72, 0x6b, 0x52, 0x65,
	0x73, 0x6f, 0x6c, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x73,
	0x66, 0x2e, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x30, 0x01, 0x42, 0x35, 0x5a, 0x33, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e,
	0x67, 0x66, 0x61, 0x73, 0x74, 0x2f, 0x70, 0x62, 0x67, 0x6f, 0x2f, 0x73, 0x66, 0x2f, 0x73, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x2f, 0x76, 0x31, 0x3b, 0x70, 0x62, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sf_search_v1_search_proto_rawDescOnce sync.Once
	file_sf_search_v1_search_proto_rawDescData = file_sf_search_v1_search_proto_rawDesc
)

func file_sf_search_v1_search_proto_rawDescGZIP() []byte {
	file_sf_search_v1_search_proto_rawDescOnce.Do(func() {
		file_sf_search_v1_search_proto_rawDescData = protoimpl.X.CompressGZIP(file_sf_search_v1_search_proto_rawDescData)
	})
	return file_sf_search_v1_search_proto_rawDescData
}

var file_sf_search_v1_search_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_sf_search_v1_search_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_sf_search_v1_search_proto_goTypes = []interface{}{
	(RouterRequest_Mode)(0),    // 0: sf.search.v1.RouterRequest.Mode
	(*ForkResolveRequest)(nil), // 1: sf.search.v1.ForkResolveRequest
	(*BlockRef)(nil),           // 2: sf.search.v1.BlockRef
	(*BackendRequest)(nil),     // 3: sf.search.v1.BackendRequest
	(*RouterRequest)(nil),      // 4: sf.search.v1.RouterRequest
	(*SearchMatch)(nil),        // 5: sf.search.v1.SearchMatch
	(*anypb.Any)(nil),          // 6: google.protobuf.Any
}
var file_sf_search_v1_search_proto_depIdxs = []int32{
	2, // 0: sf.search.v1.ForkResolveRequest.forkedBlockRefs:type_name -> sf.search.v1.BlockRef
	0, // 1: sf.search.v1.RouterRequest.mode:type_name -> sf.search.v1.RouterRequest.Mode
	6, // 2: sf.search.v1.SearchMatch.chainSpecific:type_name -> google.protobuf.Any
	3, // 3: sf.search.v1.Backend.StreamMatches:input_type -> sf.search.v1.BackendRequest
	4, // 4: sf.search.v1.Router.StreamMatches:input_type -> sf.search.v1.RouterRequest
	1, // 5: sf.search.v1.ForkResolver.StreamUndoMatches:input_type -> sf.search.v1.ForkResolveRequest
	5, // 6: sf.search.v1.Backend.StreamMatches:output_type -> sf.search.v1.SearchMatch
	5, // 7: sf.search.v1.Router.StreamMatches:output_type -> sf.search.v1.SearchMatch
	5, // 8: sf.search.v1.ForkResolver.StreamUndoMatches:output_type -> sf.search.v1.SearchMatch
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_sf_search_v1_search_proto_init() }
func file_sf_search_v1_search_proto_init() {
	if File_sf_search_v1_search_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sf_search_v1_search_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ForkResolveRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_sf_search_v1_search_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockRef); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_sf_search_v1_search_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BackendRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_sf_search_v1_search_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RouterRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_sf_search_v1_search_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchMatch); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_sf_search_v1_search_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   3,
		},
		GoTypes:           file_sf_search_v1_search_proto_goTypes,
		DependencyIndexes: file_sf_search_v1_search_proto_depIdxs,
		EnumInfos:         file_sf_search_v1_search_proto_enumTypes,
		MessageInfos:      file_sf_search_v1_search_proto_msgTypes,
	}.Build()
	File_sf_search_v1_search_proto = out.File
	file_sf_search_v1_search_proto_rawDesc = nil
	file_sf_search_v1_search_proto_goTypes = nil
	file_sf_search_v1_search_proto_depIdxs = nil
}
