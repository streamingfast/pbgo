// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.24.4
// source: sf/merger/v1/merger.proto

package pbmerger

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// MergerClient is the client API for Merger service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MergerClient interface {
	PreMergedBlocks(ctx context.Context, in *Request, opts ...grpc.CallOption) (Merger_PreMergedBlocksClient, error)
}

type mergerClient struct {
	cc grpc.ClientConnInterface
}

func NewMergerClient(cc grpc.ClientConnInterface) MergerClient {
	return &mergerClient{cc}
}

func (c *mergerClient) PreMergedBlocks(ctx context.Context, in *Request, opts ...grpc.CallOption) (Merger_PreMergedBlocksClient, error) {
	stream, err := c.cc.NewStream(ctx, &Merger_ServiceDesc.Streams[0], "/sf.merger.v1.Merger/PreMergedBlocks", opts...)
	if err != nil {
		return nil, err
	}
	x := &mergerPreMergedBlocksClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Merger_PreMergedBlocksClient interface {
	Recv() (*Response, error)
	grpc.ClientStream
}

type mergerPreMergedBlocksClient struct {
	grpc.ClientStream
}

func (x *mergerPreMergedBlocksClient) Recv() (*Response, error) {
	m := new(Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MergerServer is the server API for Merger service.
// All implementations should embed UnimplementedMergerServer
// for forward compatibility
type MergerServer interface {
	PreMergedBlocks(*Request, Merger_PreMergedBlocksServer) error
}

// UnimplementedMergerServer should be embedded to have forward compatible implementations.
type UnimplementedMergerServer struct {
}

func (UnimplementedMergerServer) PreMergedBlocks(*Request, Merger_PreMergedBlocksServer) error {
	return status.Errorf(codes.Unimplemented, "method PreMergedBlocks not implemented")
}

// UnsafeMergerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MergerServer will
// result in compilation errors.
type UnsafeMergerServer interface {
	mustEmbedUnimplementedMergerServer()
}

func RegisterMergerServer(s grpc.ServiceRegistrar, srv MergerServer) {
	s.RegisterService(&Merger_ServiceDesc, srv)
}

func _Merger_PreMergedBlocks_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Request)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MergerServer).PreMergedBlocks(m, &mergerPreMergedBlocksServer{stream})
}

type Merger_PreMergedBlocksServer interface {
	Send(*Response) error
	grpc.ServerStream
}

type mergerPreMergedBlocksServer struct {
	grpc.ServerStream
}

func (x *mergerPreMergedBlocksServer) Send(m *Response) error {
	return x.ServerStream.SendMsg(m)
}

// Merger_ServiceDesc is the grpc.ServiceDesc for Merger service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Merger_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sf.merger.v1.Merger",
	HandlerType: (*MergerServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "PreMergedBlocks",
			Handler:       _Merger_PreMergedBlocks_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "sf/merger/v1/merger.proto",
}
