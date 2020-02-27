// Code generated by protoc-gen-go. DO NOT EDIT.
// source: dfuse/graphql/v1/graphql.proto

package v1

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_struct "github.com/golang/protobuf/ptypes/struct"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Request struct {
	Query                string          `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	Variables            *_struct.Struct `protobuf:"bytes,2,opt,name=variables,proto3" json:"variables,omitempty"`
	OperationName        string          `protobuf:"bytes,3,opt,name=operationName,proto3" json:"operationName,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d3632bb23d33cb5, []int{0}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetQuery() string {
	if m != nil {
		return m.Query
	}
	return ""
}

func (m *Request) GetVariables() *_struct.Struct {
	if m != nil {
		return m.Variables
	}
	return nil
}

func (m *Request) GetOperationName() string {
	if m != nil {
		return m.OperationName
	}
	return ""
}

type Response struct {
	Data                 string   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Errors               []*Error `protobuf:"bytes,2,rep,name=errors,proto3" json:"errors,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d3632bb23d33cb5, []int{1}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func (m *Response) GetErrors() []*Error {
	if m != nil {
		return m.Errors
	}
	return nil
}

// GraphQL Error
type Error struct {
	// Description of the error intended for the developer.
	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	// The source location for the error.
	Locations []*SourceLocation `protobuf:"bytes,2,rep,name=locations,proto3" json:"locations,omitempty"`
	// Path to the `null` value justified by this error.
	Path *_struct.ListValue `protobuf:"bytes,3,opt,name=path,proto3" json:"path,omitempty"`
	// Free-form extensions (starts with a map)
	Extensions           *_struct.Struct `protobuf:"bytes,4,opt,name=extensions,proto3" json:"extensions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Error) Reset()         { *m = Error{} }
func (m *Error) String() string { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()    {}
func (*Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d3632bb23d33cb5, []int{2}
}

func (m *Error) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Error.Unmarshal(m, b)
}
func (m *Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Error.Marshal(b, m, deterministic)
}
func (m *Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Error.Merge(m, src)
}
func (m *Error) XXX_Size() int {
	return xxx_messageInfo_Error.Size(m)
}
func (m *Error) XXX_DiscardUnknown() {
	xxx_messageInfo_Error.DiscardUnknown(m)
}

var xxx_messageInfo_Error proto.InternalMessageInfo

func (m *Error) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *Error) GetLocations() []*SourceLocation {
	if m != nil {
		return m.Locations
	}
	return nil
}

func (m *Error) GetPath() *_struct.ListValue {
	if m != nil {
		return m.Path
	}
	return nil
}

func (m *Error) GetExtensions() *_struct.Struct {
	if m != nil {
		return m.Extensions
	}
	return nil
}

// The source location of an error.
type SourceLocation struct {
	// The line the error occurred at.
	Line int32 `protobuf:"varint,1,opt,name=line,proto3" json:"line,omitempty"`
	// The column the error occurred at.
	Column               int32    `protobuf:"varint,2,opt,name=column,proto3" json:"column,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SourceLocation) Reset()         { *m = SourceLocation{} }
func (m *SourceLocation) String() string { return proto.CompactTextString(m) }
func (*SourceLocation) ProtoMessage()    {}
func (*SourceLocation) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d3632bb23d33cb5, []int{3}
}

func (m *SourceLocation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SourceLocation.Unmarshal(m, b)
}
func (m *SourceLocation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SourceLocation.Marshal(b, m, deterministic)
}
func (m *SourceLocation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SourceLocation.Merge(m, src)
}
func (m *SourceLocation) XXX_Size() int {
	return xxx_messageInfo_SourceLocation.Size(m)
}
func (m *SourceLocation) XXX_DiscardUnknown() {
	xxx_messageInfo_SourceLocation.DiscardUnknown(m)
}

var xxx_messageInfo_SourceLocation proto.InternalMessageInfo

func (m *SourceLocation) GetLine() int32 {
	if m != nil {
		return m.Line
	}
	return 0
}

func (m *SourceLocation) GetColumn() int32 {
	if m != nil {
		return m.Column
	}
	return 0
}

type BlockCursor struct {
	Ver                  int32    `protobuf:"varint,1,opt,name=ver,proto3" json:"ver,omitempty"`
	BlockNum             uint64   `protobuf:"varint,2,opt,name=blockNum,proto3" json:"blockNum,omitempty"`
	BlockId              string   `protobuf:"bytes,3,opt,name=blockId,proto3" json:"blockId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BlockCursor) Reset()         { *m = BlockCursor{} }
func (m *BlockCursor) String() string { return proto.CompactTextString(m) }
func (*BlockCursor) ProtoMessage()    {}
func (*BlockCursor) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d3632bb23d33cb5, []int{4}
}

func (m *BlockCursor) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlockCursor.Unmarshal(m, b)
}
func (m *BlockCursor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlockCursor.Marshal(b, m, deterministic)
}
func (m *BlockCursor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockCursor.Merge(m, src)
}
func (m *BlockCursor) XXX_Size() int {
	return xxx_messageInfo_BlockCursor.Size(m)
}
func (m *BlockCursor) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockCursor.DiscardUnknown(m)
}

var xxx_messageInfo_BlockCursor proto.InternalMessageInfo

func (m *BlockCursor) GetVer() int32 {
	if m != nil {
		return m.Ver
	}
	return 0
}

func (m *BlockCursor) GetBlockNum() uint64 {
	if m != nil {
		return m.BlockNum
	}
	return 0
}

func (m *BlockCursor) GetBlockId() string {
	if m != nil {
		return m.BlockId
	}
	return ""
}

type TransactionCursor struct {
	Ver                  int32    `protobuf:"varint,1,opt,name=ver,proto3" json:"ver,omitempty"`
	TransactionIndex     uint32   `protobuf:"varint,2,opt,name=transactionIndex,proto3" json:"transactionIndex,omitempty"`
	TransactionHash      string   `protobuf:"bytes,3,opt,name=transactionHash,proto3" json:"transactionHash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TransactionCursor) Reset()         { *m = TransactionCursor{} }
func (m *TransactionCursor) String() string { return proto.CompactTextString(m) }
func (*TransactionCursor) ProtoMessage()    {}
func (*TransactionCursor) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d3632bb23d33cb5, []int{5}
}

func (m *TransactionCursor) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransactionCursor.Unmarshal(m, b)
}
func (m *TransactionCursor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransactionCursor.Marshal(b, m, deterministic)
}
func (m *TransactionCursor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransactionCursor.Merge(m, src)
}
func (m *TransactionCursor) XXX_Size() int {
	return xxx_messageInfo_TransactionCursor.Size(m)
}
func (m *TransactionCursor) XXX_DiscardUnknown() {
	xxx_messageInfo_TransactionCursor.DiscardUnknown(m)
}

var xxx_messageInfo_TransactionCursor proto.InternalMessageInfo

func (m *TransactionCursor) GetVer() int32 {
	if m != nil {
		return m.Ver
	}
	return 0
}

func (m *TransactionCursor) GetTransactionIndex() uint32 {
	if m != nil {
		return m.TransactionIndex
	}
	return 0
}

func (m *TransactionCursor) GetTransactionHash() string {
	if m != nil {
		return m.TransactionHash
	}
	return ""
}

func init() {
	proto.RegisterType((*Request)(nil), "dfuse.graphql.v1.Request")
	proto.RegisterType((*Response)(nil), "dfuse.graphql.v1.Response")
	proto.RegisterType((*Error)(nil), "dfuse.graphql.v1.Error")
	proto.RegisterType((*SourceLocation)(nil), "dfuse.graphql.v1.SourceLocation")
	proto.RegisterType((*BlockCursor)(nil), "dfuse.graphql.v1.BlockCursor")
	proto.RegisterType((*TransactionCursor)(nil), "dfuse.graphql.v1.TransactionCursor")
}

func init() { proto.RegisterFile("dfuse/graphql/v1/graphql.proto", fileDescriptor_0d3632bb23d33cb5) }

var fileDescriptor_0d3632bb23d33cb5 = []byte{
	// 471 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x53, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0xc5, 0xe4, 0xab, 0x99, 0xa8, 0x10, 0x56, 0x08, 0x82, 0x85, 0x50, 0x64, 0x71, 0x08, 0x20,
	0x6c, 0x1a, 0x84, 0xb8, 0x20, 0x0e, 0x85, 0x0a, 0x2a, 0x45, 0xad, 0xd8, 0x02, 0x07, 0x6e, 0x6b,
	0x67, 0xea, 0x58, 0xd8, 0x5e, 0x67, 0x3f, 0xa2, 0x20, 0xf8, 0x85, 0xfc, 0x2a, 0xe4, 0xb1, 0x4d,
	0x9a, 0x04, 0x7a, 0x9b, 0x79, 0xf3, 0xf6, 0xbd, 0xf9, 0xb0, 0xe1, 0xd1, 0xfc, 0xd2, 0x6a, 0x0c,
	0x62, 0x25, 0x8a, 0xc5, 0x32, 0x0d, 0x56, 0x47, 0x4d, 0xe8, 0x17, 0x4a, 0x1a, 0xc9, 0x86, 0x54,
	0xf7, 0x1b, 0x70, 0x75, 0xe4, 0x3e, 0x8c, 0xa5, 0x8c, 0x53, 0x0c, 0xa8, 0x1e, 0xda, 0xcb, 0x40,
	0x1b, 0x65, 0x23, 0x53, 0xf1, 0xbd, 0x5f, 0xd0, 0xe3, 0xb8, 0xb4, 0xa8, 0x0d, 0xbb, 0x0b, 0x9d,
	0xa5, 0x45, 0xf5, 0x63, 0xe4, 0x8c, 0x9d, 0x49, 0x9f, 0x57, 0x09, 0x7b, 0x05, 0xfd, 0x95, 0x50,
	0x89, 0x08, 0x53, 0xd4, 0xa3, 0x9b, 0x63, 0x67, 0x32, 0x98, 0xde, 0xf7, 0x2b, 0x49, 0xbf, 0x91,
	0xf4, 0x2f, 0x48, 0x92, 0x6f, 0x98, 0xec, 0x31, 0x1c, 0xca, 0x02, 0x95, 0x30, 0x89, 0xcc, 0xcf,
	0x44, 0x86, 0xa3, 0x16, 0x89, 0x6e, 0x83, 0xde, 0x39, 0x1c, 0x70, 0xd4, 0x85, 0xcc, 0x35, 0x32,
	0x06, 0xed, 0xb9, 0x30, 0xa2, 0x76, 0xa7, 0x98, 0x05, 0xd0, 0x45, 0xa5, 0xa4, 0x2a, 0x9d, 0x5b,
	0xe4, 0xbc, 0x3b, 0x9e, 0x7f, 0x52, 0xd6, 0x79, 0x4d, 0xf3, 0x7e, 0x3b, 0xd0, 0x21, 0x84, 0x8d,
	0xa0, 0x97, 0xa1, 0xd6, 0x22, 0xc6, 0x5a, 0xb1, 0x49, 0xd9, 0x5b, 0xe8, 0xa7, 0x32, 0xa2, 0x26,
	0x1a, 0xdd, 0xf1, 0xbe, 0xee, 0x85, 0xb4, 0x2a, 0xc2, 0x59, 0x4d, 0xe4, 0x9b, 0x27, 0xcc, 0x87,
	0x76, 0x21, 0xcc, 0x82, 0x26, 0x1a, 0x4c, 0xdd, 0xbd, 0x65, 0xcc, 0x12, 0x6d, 0xbe, 0x8a, 0xd4,
	0x22, 0x27, 0x1e, 0x7b, 0x0d, 0x80, 0x6b, 0x83, 0xb9, 0x26, 0xc3, 0xf6, 0xf5, 0x2b, 0xbc, 0x42,
	0xf5, 0xde, 0xc0, 0xad, 0xed, 0x2e, 0xca, 0x1d, 0xa5, 0x49, 0x5e, 0x4d, 0xd4, 0xe1, 0x14, 0xb3,
	0x7b, 0xd0, 0x8d, 0x64, 0x6a, 0xb3, 0x9c, 0xae, 0xd3, 0xe1, 0x75, 0xe6, 0x7d, 0x81, 0xc1, 0x71,
	0x2a, 0xa3, 0xef, 0xef, 0xac, 0xd2, 0x52, 0xb1, 0x21, 0xb4, 0x56, 0xa8, 0xea, 0x97, 0x65, 0xc8,
	0x5c, 0x38, 0x08, 0x4b, 0xc2, 0x99, 0xcd, 0xe8, 0x69, 0x9b, 0xff, 0xcd, 0xcb, 0xed, 0x51, 0x7c,
	0x3a, 0xaf, 0x0f, 0xd7, 0xa4, 0xde, 0x4f, 0xb8, 0xf3, 0x59, 0x89, 0x5c, 0x8b, 0xa8, 0xec, 0xe8,
	0xbf, 0xe2, 0x4f, 0x61, 0x68, 0x36, 0xb4, 0xd3, 0x7c, 0x8e, 0x6b, 0x32, 0x39, 0xe4, 0x7b, 0x38,
	0x9b, 0xc0, 0xed, 0x2b, 0xd8, 0x47, 0xa1, 0x17, 0xb5, 0xe9, 0x2e, 0x3c, 0x3d, 0x87, 0xde, 0x87,
	0xf2, 0x44, 0x9f, 0x66, 0xec, 0x3d, 0xf4, 0x4e, 0xd6, 0x18, 0x59, 0x83, 0xec, 0xc1, 0xfe, 0xf5,
	0xea, 0x6f, 0xda, 0x75, 0xff, 0x55, 0xaa, 0x3e, 0x38, 0xef, 0xc6, 0x0b, 0xe7, 0xf8, 0xd9, 0xb7,
	0x27, 0x71, 0x62, 0x16, 0x36, 0xf4, 0x23, 0x99, 0x05, 0xc4, 0x7d, 0x9e, 0xc8, 0xa0, 0x08, 0x63,
	0x19, 0xec, 0xfe, 0x69, 0x61, 0x97, 0x8e, 0xf5, 0xf2, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x6c,
	0xa7, 0xe9, 0x30, 0x84, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GraphQLClient is the client API for GraphQL service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GraphQLClient interface {
	Execute(ctx context.Context, in *Request, opts ...grpc.CallOption) (GraphQL_ExecuteClient, error)
}

type graphQLClient struct {
	cc *grpc.ClientConn
}

func NewGraphQLClient(cc *grpc.ClientConn) GraphQLClient {
	return &graphQLClient{cc}
}

func (c *graphQLClient) Execute(ctx context.Context, in *Request, opts ...grpc.CallOption) (GraphQL_ExecuteClient, error) {
	stream, err := c.cc.NewStream(ctx, &_GraphQL_serviceDesc.Streams[0], "/dfuse.graphql.v1.GraphQL/Execute", opts...)
	if err != nil {
		return nil, err
	}
	x := &graphQLExecuteClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type GraphQL_ExecuteClient interface {
	Recv() (*Response, error)
	grpc.ClientStream
}

type graphQLExecuteClient struct {
	grpc.ClientStream
}

func (x *graphQLExecuteClient) Recv() (*Response, error) {
	m := new(Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GraphQLServer is the server API for GraphQL service.
type GraphQLServer interface {
	Execute(*Request, GraphQL_ExecuteServer) error
}

// UnimplementedGraphQLServer can be embedded to have forward compatible implementations.
type UnimplementedGraphQLServer struct {
}

func (*UnimplementedGraphQLServer) Execute(req *Request, srv GraphQL_ExecuteServer) error {
	return status.Errorf(codes.Unimplemented, "method Execute not implemented")
}

func RegisterGraphQLServer(s *grpc.Server, srv GraphQLServer) {
	s.RegisterService(&_GraphQL_serviceDesc, srv)
}

func _GraphQL_Execute_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Request)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GraphQLServer).Execute(m, &graphQLExecuteServer{stream})
}

type GraphQL_ExecuteServer interface {
	Send(*Response) error
	grpc.ServerStream
}

type graphQLExecuteServer struct {
	grpc.ServerStream
}

func (x *graphQLExecuteServer) Send(m *Response) error {
	return x.ServerStream.SendMsg(m)
}

var _GraphQL_serviceDesc = grpc.ServiceDesc{
	ServiceName: "dfuse.graphql.v1.GraphQL",
	HandlerType: (*GraphQLServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Execute",
			Handler:       _GraphQL_Execute_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "dfuse/graphql/v1/graphql.proto",
}
