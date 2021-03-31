// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package protocol

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// WatchServiceClient is the client API for WatchService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WatchServiceClient interface {
	Watch(ctx context.Context, in *MatchPodCondition, opts ...grpc.CallOption) (WatchService_WatchClient, error)
}

type watchServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewWatchServiceClient(cc grpc.ClientConnInterface) WatchServiceClient {
	return &watchServiceClient{cc}
}

func (c *watchServiceClient) Watch(ctx context.Context, in *MatchPodCondition, opts ...grpc.CallOption) (WatchService_WatchClient, error) {
	stream, err := c.cc.NewStream(ctx, &_WatchService_serviceDesc.Streams[0], "/WatchService/Watch", opts...)
	if err != nil {
		return nil, err
	}
	x := &watchServiceWatchClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type WatchService_WatchClient interface {
	Recv() (*MatchPodResponse, error)
	grpc.ClientStream
}

type watchServiceWatchClient struct {
	grpc.ClientStream
}

func (x *watchServiceWatchClient) Recv() (*MatchPodResponse, error) {
	m := new(MatchPodResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// WatchServiceServer is the server API for WatchService service.
// All implementations must embed UnimplementedWatchServiceServer
// for forward compatibility
type WatchServiceServer interface {
	Watch(*MatchPodCondition, WatchService_WatchServer) error
	mustEmbedUnimplementedWatchServiceServer()
}

// UnimplementedWatchServiceServer must be embedded to have forward compatible implementations.
type UnimplementedWatchServiceServer struct {
}

func (UnimplementedWatchServiceServer) Watch(*MatchPodCondition, WatchService_WatchServer) error {
	return status.Errorf(codes.Unimplemented, "method Watch not implemented")
}
func (UnimplementedWatchServiceServer) mustEmbedUnimplementedWatchServiceServer() {}

// UnsafeWatchServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WatchServiceServer will
// result in compilation errors.
type UnsafeWatchServiceServer interface {
	mustEmbedUnimplementedWatchServiceServer()
}

func RegisterWatchServiceServer(s grpc.ServiceRegistrar, srv WatchServiceServer) {
	s.RegisterService(&_WatchService_serviceDesc, srv)
}

func _WatchService_Watch_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(MatchPodCondition)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(WatchServiceServer).Watch(m, &watchServiceWatchServer{stream})
}

type WatchService_WatchServer interface {
	Send(*MatchPodResponse) error
	grpc.ServerStream
}

type watchServiceWatchServer struct {
	grpc.ServerStream
}

func (x *watchServiceWatchServer) Send(m *MatchPodResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _WatchService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "WatchService",
	HandlerType: (*WatchServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Watch",
			Handler:       _WatchService_Watch_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "watcher.proto",
}
