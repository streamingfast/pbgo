// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v4.24.4
// source: sf/firehose/v1/firehose.proto

package pbfirehose

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

type ForkStep int32

const (
	ForkStep_STEP_UNKNOWN ForkStep = 0
	// Block is new head block of the chain, that is linear with the previous block
	ForkStep_STEP_NEW ForkStep = 1
	// Block is now forked and should be undone, it's not the head block of the chain anymore
	ForkStep_STEP_UNDO ForkStep = 2
	// Block is now irreversible and can be committed to (finality is chain specific, see chain documentation for more details)
	ForkStep_STEP_IRREVERSIBLE ForkStep = 4
)

// Enum value maps for ForkStep.
var (
	ForkStep_name = map[int32]string{
		0: "STEP_UNKNOWN",
		1: "STEP_NEW",
		2: "STEP_UNDO",
		4: "STEP_IRREVERSIBLE",
	}
	ForkStep_value = map[string]int32{
		"STEP_UNKNOWN":      0,
		"STEP_NEW":          1,
		"STEP_UNDO":         2,
		"STEP_IRREVERSIBLE": 4,
	}
)

func (x ForkStep) Enum() *ForkStep {
	p := new(ForkStep)
	*p = x
	return p
}

func (x ForkStep) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ForkStep) Descriptor() protoreflect.EnumDescriptor {
	return file_sf_firehose_v1_firehose_proto_enumTypes[0].Descriptor()
}

func (ForkStep) Type() protoreflect.EnumType {
	return &file_sf_firehose_v1_firehose_proto_enumTypes[0]
}

func (x ForkStep) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ForkStep.Descriptor instead.
func (ForkStep) EnumDescriptor() ([]byte, []int) {
	return file_sf_firehose_v1_firehose_proto_rawDescGZIP(), []int{0}
}

// For historical segments, forks are not passed
type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Controls where the stream of blocks will start.
	//
	// The stream will start **inclusively** at the requested block num.
	//
	// When not provided, starts at first streamable block of the chain. Not all
	// chain starts at the same block number, so you might get an higher block than
	// requested when using default value of 0.
	//
	// Can be negative, will be resolved relative to the chain head block, assuming
	// a chain at head block #100, then using `-50` as the value will start at block
	// #50. If it resolves before first streamable block of chain, we assume start
	// of chain.
	//
	// If `start_cursor` is passed, this value is ignored and the stream instead starts
	// immediately after the Block pointed by the opaque `start_cursor` value.
	StartBlockNum int64 `protobuf:"varint,1,opt,name=start_block_num,json=startBlockNum,proto3" json:"start_block_num,omitempty"`
	// Controls where the stream of blocks will start which will be immediately after
	// the Block pointed by this opaque cursor.
	//
	// Obtain this value from a previously received `Response.cursor`.
	//
	// This value takes precedence over `start_block_num`.
	StartCursor string `protobuf:"bytes,13,opt,name=start_cursor,json=startCursor,proto3" json:"start_cursor,omitempty"`
	// When non-zero, controls where the stream of blocks will stop.
	//
	// The stream will close **after** that block has passed so the boundary is
	// **inclusive**.
	StopBlockNum uint64 `protobuf:"varint,5,opt,name=stop_block_num,json=stopBlockNum,proto3" json:"stop_block_num,omitempty"`
	// Filter the steps you want to see. If not specified, defaults to all steps.
	//
	// Most common steps will be [STEP_IRREVERSIBLE], or [STEP_NEW, STEP_UNDO, STEP_IRREVERSIBLE].
	ForkSteps []ForkStep `protobuf:"varint,8,rep,packed,name=fork_steps,json=forkSteps,proto3,enum=sf.firehose.v1.ForkStep" json:"fork_steps,omitempty"`
	// feature moved to transforms
	//
	// Deprecated: Do not use.
	IncludeFilterExpr string `protobuf:"bytes,10,opt,name=include_filter_expr,json=includeFilterExpr,proto3" json:"include_filter_expr,omitempty"`
	// Deprecated: Do not use.
	ExcludeFilterExpr string `protobuf:"bytes,11,opt,name=exclude_filter_expr,json=excludeFilterExpr,proto3" json:"exclude_filter_expr,omitempty"`
	//- EOS "handoffs:3"
	//- EOS "lib"
	//- EOS "confirms:3"
	//- ETH "confirms:200"
	//- ETH "confirms:7"
	//- SOL "commmitement:finalized"
	//- SOL "confirms:200"
	IrreversibilityCondition string       `protobuf:"bytes,17,opt,name=irreversibility_condition,json=irreversibilityCondition,proto3" json:"irreversibility_condition,omitempty"`
	Transforms               []*anypb.Any `protobuf:"bytes,18,rep,name=transforms,proto3" json:"transforms,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sf_firehose_v1_firehose_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_sf_firehose_v1_firehose_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_sf_firehose_v1_firehose_proto_rawDescGZIP(), []int{0}
}

func (x *Request) GetStartBlockNum() int64 {
	if x != nil {
		return x.StartBlockNum
	}
	return 0
}

func (x *Request) GetStartCursor() string {
	if x != nil {
		return x.StartCursor
	}
	return ""
}

func (x *Request) GetStopBlockNum() uint64 {
	if x != nil {
		return x.StopBlockNum
	}
	return 0
}

func (x *Request) GetForkSteps() []ForkStep {
	if x != nil {
		return x.ForkSteps
	}
	return nil
}

// Deprecated: Do not use.
func (x *Request) GetIncludeFilterExpr() string {
	if x != nil {
		return x.IncludeFilterExpr
	}
	return ""
}

// Deprecated: Do not use.
func (x *Request) GetExcludeFilterExpr() string {
	if x != nil {
		return x.ExcludeFilterExpr
	}
	return ""
}

func (x *Request) GetIrreversibilityCondition() string {
	if x != nil {
		return x.IrreversibilityCondition
	}
	return ""
}

func (x *Request) GetTransforms() []*anypb.Any {
	if x != nil {
		return x.Transforms
	}
	return nil
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Chain specific block payload, one of:
	// - sf.eosio.codec.v1.Block
	// - sf.ethereum.codec.v1.Block
	// - sf.near.codec.v1.Block
	// - sf.solana.codec.v1.Block
	Block  *anypb.Any `protobuf:"bytes,1,opt,name=block,proto3" json:"block,omitempty"`
	Step   ForkStep   `protobuf:"varint,6,opt,name=step,proto3,enum=sf.firehose.v1.ForkStep" json:"step,omitempty"`
	Cursor string     `protobuf:"bytes,10,opt,name=cursor,proto3" json:"cursor,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sf_firehose_v1_firehose_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_sf_firehose_v1_firehose_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_sf_firehose_v1_firehose_proto_rawDescGZIP(), []int{1}
}

func (x *Response) GetBlock() *anypb.Any {
	if x != nil {
		return x.Block
	}
	return nil
}

func (x *Response) GetStep() ForkStep {
	if x != nil {
		return x.Step
	}
	return ForkStep_STEP_UNKNOWN
}

func (x *Response) GetCursor() string {
	if x != nil {
		return x.Cursor
	}
	return ""
}

var File_sf_firehose_v1_firehose_proto protoreflect.FileDescriptor

var file_sf_firehose_v1_firehose_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x73, 0x66, 0x2f, 0x66, 0x69, 0x72, 0x65, 0x68, 0x6f, 0x73, 0x65, 0x2f, 0x76, 0x31,
	0x2f, 0x66, 0x69, 0x72, 0x65, 0x68, 0x6f, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0e, 0x73, 0x66, 0x2e, 0x66, 0x69, 0x72, 0x65, 0x68, 0x6f, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x1a,
	0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9a, 0x03, 0x0a, 0x07, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x26, 0x0a, 0x0f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f,
	0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0d, 0x73, 0x74, 0x61, 0x72, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x4e, 0x75, 0x6d, 0x12, 0x21,
	0x0a, 0x0c, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x18, 0x0d,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x74, 0x61, 0x72, 0x74, 0x43, 0x75, 0x72, 0x73, 0x6f,
	0x72, 0x12, 0x24, 0x0a, 0x0e, 0x73, 0x74, 0x6f, 0x70, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f,
	0x6e, 0x75, 0x6d, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0c, 0x73, 0x74, 0x6f, 0x70, 0x42,
	0x6c, 0x6f, 0x63, 0x6b, 0x4e, 0x75, 0x6d, 0x12, 0x37, 0x0a, 0x0a, 0x66, 0x6f, 0x72, 0x6b, 0x5f,
	0x73, 0x74, 0x65, 0x70, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x73, 0x66,
	0x2e, 0x66, 0x69, 0x72, 0x65, 0x68, 0x6f, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x6f, 0x72,
	0x6b, 0x53, 0x74, 0x65, 0x70, 0x52, 0x09, 0x66, 0x6f, 0x72, 0x6b, 0x53, 0x74, 0x65, 0x70, 0x73,
	0x12, 0x32, 0x0a, 0x13, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x5f, 0x66, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x5f, 0x65, 0x78, 0x70, 0x72, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x42, 0x02, 0x18,
	0x01, 0x52, 0x11, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x45, 0x78, 0x70, 0x72, 0x12, 0x32, 0x0a, 0x13, 0x65, 0x78, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x5f,
	0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x5f, 0x65, 0x78, 0x70, 0x72, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x02, 0x18, 0x01, 0x52, 0x11, 0x65, 0x78, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x46, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x45, 0x78, 0x70, 0x72, 0x12, 0x3b, 0x0a, 0x19, 0x69, 0x72, 0x72, 0x65,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x5f, 0x63, 0x6f, 0x6e, 0x64,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x52, 0x18, 0x69, 0x72, 0x72,
	0x65, 0x76, 0x65, 0x72, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x43, 0x6f, 0x6e, 0x64,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x34, 0x0a, 0x0a, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f,
	0x72, 0x6d, 0x73, 0x18, 0x12, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52,
	0x0a, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x73, 0x4a, 0x04, 0x08, 0x0f, 0x10,
	0x10, 0x4a, 0x04, 0x08, 0x10, 0x10, 0x11, 0x22, 0x7c, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x05, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x05, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x12,
	0x2c, 0x0a, 0x04, 0x73, 0x74, 0x65, 0x70, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e,
	0x73, 0x66, 0x2e, 0x66, 0x69, 0x72, 0x65, 0x68, 0x6f, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x46,
	0x6f, 0x72, 0x6b, 0x53, 0x74, 0x65, 0x70, 0x52, 0x04, 0x73, 0x74, 0x65, 0x70, 0x12, 0x16, 0x0a,
	0x06, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63,
	0x75, 0x72, 0x73, 0x6f, 0x72, 0x2a, 0x5c, 0x0a, 0x08, 0x46, 0x6f, 0x72, 0x6b, 0x53, 0x74, 0x65,
	0x70, 0x12, 0x10, 0x0a, 0x0c, 0x53, 0x54, 0x45, 0x50, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57,
	0x4e, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x53, 0x54, 0x45, 0x50, 0x5f, 0x4e, 0x45, 0x57, 0x10,
	0x01, 0x12, 0x0d, 0x0a, 0x09, 0x53, 0x54, 0x45, 0x50, 0x5f, 0x55, 0x4e, 0x44, 0x4f, 0x10, 0x02,
	0x12, 0x15, 0x0a, 0x11, 0x53, 0x54, 0x45, 0x50, 0x5f, 0x49, 0x52, 0x52, 0x45, 0x56, 0x45, 0x52,
	0x53, 0x49, 0x42, 0x4c, 0x45, 0x10, 0x04, 0x22, 0x04, 0x08, 0x03, 0x10, 0x03, 0x22, 0x04, 0x08,
	0x05, 0x10, 0x05, 0x32, 0x47, 0x0a, 0x06, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x3d, 0x0a,
	0x06, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x12, 0x17, 0x2e, 0x73, 0x66, 0x2e, 0x66, 0x69, 0x72,
	0x65, 0x68, 0x6f, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x18, 0x2e, 0x73, 0x66, 0x2e, 0x66, 0x69, 0x72, 0x65, 0x68, 0x6f, 0x73, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x42, 0x39, 0x5a, 0x37,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x69, 0x6e, 0x67, 0x66, 0x61, 0x73, 0x74, 0x2f, 0x70, 0x62, 0x67, 0x6f, 0x2f, 0x73, 0x66,
	0x2f, 0x66, 0x69, 0x72, 0x65, 0x68, 0x6f, 0x73, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x70, 0x62, 0x66,
	0x69, 0x72, 0x65, 0x68, 0x6f, 0x73, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sf_firehose_v1_firehose_proto_rawDescOnce sync.Once
	file_sf_firehose_v1_firehose_proto_rawDescData = file_sf_firehose_v1_firehose_proto_rawDesc
)

func file_sf_firehose_v1_firehose_proto_rawDescGZIP() []byte {
	file_sf_firehose_v1_firehose_proto_rawDescOnce.Do(func() {
		file_sf_firehose_v1_firehose_proto_rawDescData = protoimpl.X.CompressGZIP(file_sf_firehose_v1_firehose_proto_rawDescData)
	})
	return file_sf_firehose_v1_firehose_proto_rawDescData
}

var file_sf_firehose_v1_firehose_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_sf_firehose_v1_firehose_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_sf_firehose_v1_firehose_proto_goTypes = []interface{}{
	(ForkStep)(0),     // 0: sf.firehose.v1.ForkStep
	(*Request)(nil),   // 1: sf.firehose.v1.Request
	(*Response)(nil),  // 2: sf.firehose.v1.Response
	(*anypb.Any)(nil), // 3: google.protobuf.Any
}
var file_sf_firehose_v1_firehose_proto_depIdxs = []int32{
	0, // 0: sf.firehose.v1.Request.fork_steps:type_name -> sf.firehose.v1.ForkStep
	3, // 1: sf.firehose.v1.Request.transforms:type_name -> google.protobuf.Any
	3, // 2: sf.firehose.v1.Response.block:type_name -> google.protobuf.Any
	0, // 3: sf.firehose.v1.Response.step:type_name -> sf.firehose.v1.ForkStep
	1, // 4: sf.firehose.v1.Stream.Blocks:input_type -> sf.firehose.v1.Request
	2, // 5: sf.firehose.v1.Stream.Blocks:output_type -> sf.firehose.v1.Response
	5, // [5:6] is the sub-list for method output_type
	4, // [4:5] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_sf_firehose_v1_firehose_proto_init() }
func file_sf_firehose_v1_firehose_proto_init() {
	if File_sf_firehose_v1_firehose_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sf_firehose_v1_firehose_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_sf_firehose_v1_firehose_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
			RawDescriptor: file_sf_firehose_v1_firehose_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_sf_firehose_v1_firehose_proto_goTypes,
		DependencyIndexes: file_sf_firehose_v1_firehose_proto_depIdxs,
		EnumInfos:         file_sf_firehose_v1_firehose_proto_enumTypes,
		MessageInfos:      file_sf_firehose_v1_firehose_proto_msgTypes,
	}.Build()
	File_sf_firehose_v1_firehose_proto = out.File
	file_sf_firehose_v1_firehose_proto_rawDesc = nil
	file_sf_firehose_v1_firehose_proto_goTypes = nil
	file_sf_firehose_v1_firehose_proto_depIdxs = nil
}
