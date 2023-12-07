// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.21.5
// source: sf/fluxdb/v1/fluxdb.proto

package pbfluxdb

import (
	v1 "github.com/streamingfast/bstream/pb/sf/bstream/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// TabletIndex is the binary message used to save a Tablet index
// value inside persisten storage.
type TabletIndex struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SquelchedCount uint64              `protobuf:"varint,1,opt,name=squelched_count,json=squelchedCount,proto3" json:"squelched_count,omitempty"`
	Entries        []*TabletIndexEntry `protobuf:"bytes,2,rep,name=entries,proto3" json:"entries,omitempty"`
}

func (x *TabletIndex) Reset() {
	*x = TabletIndex{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sf_fluxdb_v1_fluxdb_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TabletIndex) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TabletIndex) ProtoMessage() {}

func (x *TabletIndex) ProtoReflect() protoreflect.Message {
	mi := &file_sf_fluxdb_v1_fluxdb_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TabletIndex.ProtoReflect.Descriptor instead.
func (*TabletIndex) Descriptor() ([]byte, []int) {
	return file_sf_fluxdb_v1_fluxdb_proto_rawDescGZIP(), []int{0}
}

func (x *TabletIndex) GetSquelchedCount() uint64 {
	if x != nil {
		return x.SquelchedCount
	}
	return 0
}

func (x *TabletIndex) GetEntries() []*TabletIndexEntry {
	if x != nil {
		return x.Entries
	}
	return nil
}

// TabletIndexEntry represents a single entry within a Tablet index
// instance.
type TabletIndexEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PrimaryKey []byte `protobuf:"bytes,1,opt,name=primary_key,json=primaryKey,proto3" json:"primary_key,omitempty"`
	Height     uint64 `protobuf:"varint,2,opt,name=height,proto3" json:"height,omitempty"`
}

func (x *TabletIndexEntry) Reset() {
	*x = TabletIndexEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sf_fluxdb_v1_fluxdb_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TabletIndexEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TabletIndexEntry) ProtoMessage() {}

func (x *TabletIndexEntry) ProtoReflect() protoreflect.Message {
	mi := &file_sf_fluxdb_v1_fluxdb_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TabletIndexEntry.ProtoReflect.Descriptor instead.
func (*TabletIndexEntry) Descriptor() ([]byte, []int) {
	return file_sf_fluxdb_v1_fluxdb_proto_rawDescGZIP(), []int{1}
}

func (x *TabletIndexEntry) GetPrimaryKey() []byte {
	if x != nil {
		return x.PrimaryKey
	}
	return nil
}

func (x *TabletIndexEntry) GetHeight() uint64 {
	if x != nil {
		return x.Height
	}
	return 0
}

// Checkpoint represents where FluxDB did its last write
// information to the database. It mainly contains the latest
// height value written as well as the block reference where it
// was.
type Checkpoint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Height uint64       `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
	Block  *v1.BlockRef `protobuf:"bytes,2,opt,name=block,proto3" json:"block,omitempty"`
}

func (x *Checkpoint) Reset() {
	*x = Checkpoint{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sf_fluxdb_v1_fluxdb_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Checkpoint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Checkpoint) ProtoMessage() {}

func (x *Checkpoint) ProtoReflect() protoreflect.Message {
	mi := &file_sf_fluxdb_v1_fluxdb_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Checkpoint.ProtoReflect.Descriptor instead.
func (*Checkpoint) Descriptor() ([]byte, []int) {
	return file_sf_fluxdb_v1_fluxdb_proto_rawDescGZIP(), []int{2}
}

func (x *Checkpoint) GetHeight() uint64 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *Checkpoint) GetBlock() *v1.BlockRef {
	if x != nil {
		return x.Block
	}
	return nil
}

// WriteRequest represents all the necessary data to write to database
// for a given height.
//
// This is used by the reprocessing sharder to store write requests
// for a given specific shard over a bounded sequence of heights.
type WriteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SingletEntries []*WriteEntry `protobuf:"bytes,1,rep,name=singlet_entries,json=singletEntries,proto3" json:"singlet_entries,omitempty"`
	TabletRows     []*WriteEntry `protobuf:"bytes,2,rep,name=tablet_rows,json=tabletRows,proto3" json:"tablet_rows,omitempty"`
	Height         uint64        `protobuf:"varint,3,opt,name=height,proto3" json:"height,omitempty"`
	Block          *v1.BlockRef  `protobuf:"bytes,4,opt,name=block,proto3" json:"block,omitempty"`
}

func (x *WriteRequest) Reset() {
	*x = WriteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sf_fluxdb_v1_fluxdb_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WriteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WriteRequest) ProtoMessage() {}

func (x *WriteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sf_fluxdb_v1_fluxdb_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WriteRequest.ProtoReflect.Descriptor instead.
func (*WriteRequest) Descriptor() ([]byte, []int) {
	return file_sf_fluxdb_v1_fluxdb_proto_rawDescGZIP(), []int{3}
}

func (x *WriteRequest) GetSingletEntries() []*WriteEntry {
	if x != nil {
		return x.SingletEntries
	}
	return nil
}

func (x *WriteRequest) GetTabletRows() []*WriteEntry {
	if x != nil {
		return x.TabletRows
	}
	return nil
}

func (x *WriteRequest) GetHeight() uint64 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *WriteRequest) GetBlock() *v1.BlockRef {
	if x != nil {
		return x.Block
	}
	return nil
}

type WriteEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   []byte `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	Value []byte `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *WriteEntry) Reset() {
	*x = WriteEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sf_fluxdb_v1_fluxdb_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WriteEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WriteEntry) ProtoMessage() {}

func (x *WriteEntry) ProtoReflect() protoreflect.Message {
	mi := &file_sf_fluxdb_v1_fluxdb_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WriteEntry.ProtoReflect.Descriptor instead.
func (*WriteEntry) Descriptor() ([]byte, []int) {
	return file_sf_fluxdb_v1_fluxdb_proto_rawDescGZIP(), []int{4}
}

func (x *WriteEntry) GetKey() []byte {
	if x != nil {
		return x.Key
	}
	return nil
}

func (x *WriteEntry) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

var File_sf_fluxdb_v1_fluxdb_proto protoreflect.FileDescriptor

var file_sf_fluxdb_v1_fluxdb_proto_rawDesc = []byte{
	0x0a, 0x19, 0x73, 0x66, 0x2f, 0x66, 0x6c, 0x75, 0x78, 0x64, 0x62, 0x2f, 0x76, 0x31, 0x2f, 0x66,
	0x6c, 0x75, 0x78, 0x64, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x73, 0x66, 0x2e,
	0x66, 0x6c, 0x75, 0x78, 0x64, 0x62, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x73, 0x66, 0x2f, 0x62, 0x73,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x70, 0x0a, 0x0b, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x74,
	0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x27, 0x0a, 0x0f, 0x73, 0x71, 0x75, 0x65, 0x6c, 0x63, 0x68,
	0x65, 0x64, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0e,
	0x73, 0x71, 0x75, 0x65, 0x6c, 0x63, 0x68, 0x65, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x38,
	0x0a, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1e, 0x2e, 0x73, 0x66, 0x2e, 0x66, 0x6c, 0x75, 0x78, 0x64, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x54,
	0x61, 0x62, 0x6c, 0x65, 0x74, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x07, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x22, 0x4b, 0x0a, 0x10, 0x54, 0x61, 0x62, 0x6c,
	0x65, 0x74, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x1f, 0x0a, 0x0b,
	0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x0a, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x4b, 0x65, 0x79, 0x12, 0x16, 0x0a,
	0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x68,
	0x65, 0x69, 0x67, 0x68, 0x74, 0x22, 0x53, 0x0a, 0x0a, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x2d, 0x0a, 0x05, 0x62,
	0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73, 0x66, 0x2e,
	0x62, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b,
	0x52, 0x65, 0x66, 0x52, 0x05, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x22, 0xd3, 0x01, 0x0a, 0x0c, 0x57,
	0x72, 0x69, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x41, 0x0a, 0x0f, 0x73,
	0x69, 0x6e, 0x67, 0x6c, 0x65, 0x74, 0x5f, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x73, 0x66, 0x2e, 0x66, 0x6c, 0x75, 0x78, 0x64, 0x62,
	0x2e, 0x76, 0x31, 0x2e, 0x57, 0x72, 0x69, 0x74, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0e,
	0x73, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x12, 0x39,
	0x0a, 0x0b, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x74, 0x5f, 0x72, 0x6f, 0x77, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x73, 0x66, 0x2e, 0x66, 0x6c, 0x75, 0x78, 0x64, 0x62, 0x2e,
	0x76, 0x31, 0x2e, 0x57, 0x72, 0x69, 0x74, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0a, 0x74,
	0x61, 0x62, 0x6c, 0x65, 0x74, 0x52, 0x6f, 0x77, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x69,
	0x67, 0x68, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68,
	0x74, 0x12, 0x2d, 0x0a, 0x05, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x17, 0x2e, 0x73, 0x66, 0x2e, 0x62, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x76, 0x31,
	0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x66, 0x52, 0x05, 0x62, 0x6c, 0x6f, 0x63, 0x6b,
	0x22, 0x34, 0x0a, 0x0a, 0x57, 0x72, 0x69, 0x74, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x35, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x66, 0x61,
	0x73, 0x74, 0x2f, 0x70, 0x62, 0x67, 0x6f, 0x2f, 0x73, 0x66, 0x2f, 0x66, 0x6c, 0x75, 0x78, 0x64,
	0x62, 0x2f, 0x76, 0x31, 0x3b, 0x70, 0x62, 0x66, 0x6c, 0x75, 0x78, 0x64, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sf_fluxdb_v1_fluxdb_proto_rawDescOnce sync.Once
	file_sf_fluxdb_v1_fluxdb_proto_rawDescData = file_sf_fluxdb_v1_fluxdb_proto_rawDesc
)

func file_sf_fluxdb_v1_fluxdb_proto_rawDescGZIP() []byte {
	file_sf_fluxdb_v1_fluxdb_proto_rawDescOnce.Do(func() {
		file_sf_fluxdb_v1_fluxdb_proto_rawDescData = protoimpl.X.CompressGZIP(file_sf_fluxdb_v1_fluxdb_proto_rawDescData)
	})
	return file_sf_fluxdb_v1_fluxdb_proto_rawDescData
}

var file_sf_fluxdb_v1_fluxdb_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_sf_fluxdb_v1_fluxdb_proto_goTypes = []interface{}{
	(*TabletIndex)(nil),      // 0: sf.fluxdb.v1.TabletIndex
	(*TabletIndexEntry)(nil), // 1: sf.fluxdb.v1.TabletIndexEntry
	(*Checkpoint)(nil),       // 2: sf.fluxdb.v1.Checkpoint
	(*WriteRequest)(nil),     // 3: sf.fluxdb.v1.WriteRequest
	(*WriteEntry)(nil),       // 4: sf.fluxdb.v1.WriteEntry
	(*v1.BlockRef)(nil),      // 5: sf.bstream.v1.BlockRef
}
var file_sf_fluxdb_v1_fluxdb_proto_depIdxs = []int32{
	1, // 0: sf.fluxdb.v1.TabletIndex.entries:type_name -> sf.fluxdb.v1.TabletIndexEntry
	5, // 1: sf.fluxdb.v1.Checkpoint.block:type_name -> sf.bstream.v1.BlockRef
	4, // 2: sf.fluxdb.v1.WriteRequest.singlet_entries:type_name -> sf.fluxdb.v1.WriteEntry
	4, // 3: sf.fluxdb.v1.WriteRequest.tablet_rows:type_name -> sf.fluxdb.v1.WriteEntry
	5, // 4: sf.fluxdb.v1.WriteRequest.block:type_name -> sf.bstream.v1.BlockRef
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_sf_fluxdb_v1_fluxdb_proto_init() }
func file_sf_fluxdb_v1_fluxdb_proto_init() {
	if File_sf_fluxdb_v1_fluxdb_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sf_fluxdb_v1_fluxdb_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TabletIndex); i {
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
		file_sf_fluxdb_v1_fluxdb_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TabletIndexEntry); i {
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
		file_sf_fluxdb_v1_fluxdb_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Checkpoint); i {
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
		file_sf_fluxdb_v1_fluxdb_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WriteRequest); i {
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
		file_sf_fluxdb_v1_fluxdb_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WriteEntry); i {
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
			RawDescriptor: file_sf_fluxdb_v1_fluxdb_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_sf_fluxdb_v1_fluxdb_proto_goTypes,
		DependencyIndexes: file_sf_fluxdb_v1_fluxdb_proto_depIdxs,
		MessageInfos:      file_sf_fluxdb_v1_fluxdb_proto_msgTypes,
	}.Build()
	File_sf_fluxdb_v1_fluxdb_proto = out.File
	file_sf_fluxdb_v1_fluxdb_proto_rawDesc = nil
	file_sf_fluxdb_v1_fluxdb_proto_goTypes = nil
	file_sf_fluxdb_v1_fluxdb_proto_depIdxs = nil
}
