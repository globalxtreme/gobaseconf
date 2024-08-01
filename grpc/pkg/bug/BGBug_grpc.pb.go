// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.3
// source: BGBug.proto

package bug

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

const (
	BugService_Log_FullMethodName = "/bug.BugService/Log"
)

// BugServiceClient is the client API for BugService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BugServiceClient interface {
	Log(ctx context.Context, in *LogRequest, opts ...grpc.CallOption) (*BGResponse, error)
}

type bugServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBugServiceClient(cc grpc.ClientConnInterface) BugServiceClient {
	return &bugServiceClient{cc}
}

func (c *bugServiceClient) Log(ctx context.Context, in *LogRequest, opts ...grpc.CallOption) (*BGResponse, error) {
	out := new(BGResponse)
	err := c.cc.Invoke(ctx, BugService_Log_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BugServiceServer is the server API for BugService service.
// All implementations must embed UnimplementedBugServiceServer
// for forward compatibility
type BugServiceServer interface {
	Log(context.Context, *LogRequest) (*BGResponse, error)
	mustEmbedUnimplementedBugServiceServer()
}

// UnimplementedBugServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBugServiceServer struct {
}

func (UnimplementedBugServiceServer) Log(context.Context, *LogRequest) (*BGResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Log not implemented")
}
func (UnimplementedBugServiceServer) mustEmbedUnimplementedBugServiceServer() {}

// UnsafeBugServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BugServiceServer will
// result in compilation errors.
type UnsafeBugServiceServer interface {
	mustEmbedUnimplementedBugServiceServer()
}

func RegisterBugServiceServer(s grpc.ServiceRegistrar, srv BugServiceServer) {
	s.RegisterService(&BugService_ServiceDesc, srv)
}

func _BugService_Log_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BugServiceServer).Log(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BugService_Log_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BugServiceServer).Log(ctx, req.(*LogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BugService_ServiceDesc is the grpc.ServiceDesc for BugService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BugService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "bug.BugService",
	HandlerType: (*BugServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Log",
			Handler:    _BugService_Log_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "BGBug.proto",
}
