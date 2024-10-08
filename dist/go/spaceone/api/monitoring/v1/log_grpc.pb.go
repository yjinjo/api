// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.6.1
// source: spaceone/api/monitoring/v1/log.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Log_List_FullMethodName = "/spaceone.api.monitoring.v1.Log/list"
)

// LogClient is the client API for Log service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LogClient interface {
	List(ctx context.Context, in *LogRequest, opts ...grpc.CallOption) (*LogDataInfo, error)
}

type logClient struct {
	cc grpc.ClientConnInterface
}

func NewLogClient(cc grpc.ClientConnInterface) LogClient {
	return &logClient{cc}
}

func (c *logClient) List(ctx context.Context, in *LogRequest, opts ...grpc.CallOption) (*LogDataInfo, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(LogDataInfo)
	err := c.cc.Invoke(ctx, Log_List_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LogServer is the server API for Log service.
// All implementations must embed UnimplementedLogServer
// for forward compatibility.
type LogServer interface {
	List(context.Context, *LogRequest) (*LogDataInfo, error)
	mustEmbedUnimplementedLogServer()
}

// UnimplementedLogServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedLogServer struct{}

func (UnimplementedLogServer) List(context.Context, *LogRequest) (*LogDataInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedLogServer) mustEmbedUnimplementedLogServer() {}
func (UnimplementedLogServer) testEmbeddedByValue()             {}

// UnsafeLogServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LogServer will
// result in compilation errors.
type UnsafeLogServer interface {
	mustEmbedUnimplementedLogServer()
}

func RegisterLogServer(s grpc.ServiceRegistrar, srv LogServer) {
	// If the following call pancis, it indicates UnimplementedLogServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Log_ServiceDesc, srv)
}

func _Log_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Log_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogServer).List(ctx, req.(*LogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Log_ServiceDesc is the grpc.ServiceDesc for Log service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Log_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "spaceone.api.monitoring.v1.Log",
	HandlerType: (*LogServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "list",
			Handler:    _Log_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "spaceone/api/monitoring/v1/log.proto",
}
