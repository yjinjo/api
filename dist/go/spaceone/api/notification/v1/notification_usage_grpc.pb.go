// A NotificationUsage is a resource indicating daily Protocol usage. The limit set by the resource Quota is applied based on the NotificationUsage.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.6.1
// source: spaceone/api/notification/v1/notification_usage.proto

package v1

import (
	context "context"
	_struct "github.com/golang/protobuf/ptypes/struct"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	NotificationUsage_List_FullMethodName = "/spaceone.api.notification.v1.NotificationUsage/list"
	NotificationUsage_Stat_FullMethodName = "/spaceone.api.notification.v1.NotificationUsage/stat"
)

// NotificationUsageClient is the client API for NotificationUsage service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NotificationUsageClient interface {
	// Gets a list of all NotificationUsages. You can use a query to get a filtered list of Notification Usages.
	List(ctx context.Context, in *NotificationUsageQuery, opts ...grpc.CallOption) (*NotificationUsagesInfo, error)
	Stat(ctx context.Context, in *NotificationUsageStatQuery, opts ...grpc.CallOption) (*_struct.Struct, error)
}

type notificationUsageClient struct {
	cc grpc.ClientConnInterface
}

func NewNotificationUsageClient(cc grpc.ClientConnInterface) NotificationUsageClient {
	return &notificationUsageClient{cc}
}

func (c *notificationUsageClient) List(ctx context.Context, in *NotificationUsageQuery, opts ...grpc.CallOption) (*NotificationUsagesInfo, error) {
	out := new(NotificationUsagesInfo)
	err := c.cc.Invoke(ctx, NotificationUsage_List_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationUsageClient) Stat(ctx context.Context, in *NotificationUsageStatQuery, opts ...grpc.CallOption) (*_struct.Struct, error) {
	out := new(_struct.Struct)
	err := c.cc.Invoke(ctx, NotificationUsage_Stat_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NotificationUsageServer is the server API for NotificationUsage service.
// All implementations must embed UnimplementedNotificationUsageServer
// for forward compatibility
type NotificationUsageServer interface {
	// Gets a list of all NotificationUsages. You can use a query to get a filtered list of Notification Usages.
	List(context.Context, *NotificationUsageQuery) (*NotificationUsagesInfo, error)
	Stat(context.Context, *NotificationUsageStatQuery) (*_struct.Struct, error)
	mustEmbedUnimplementedNotificationUsageServer()
}

// UnimplementedNotificationUsageServer must be embedded to have forward compatible implementations.
type UnimplementedNotificationUsageServer struct {
}

func (UnimplementedNotificationUsageServer) List(context.Context, *NotificationUsageQuery) (*NotificationUsagesInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedNotificationUsageServer) Stat(context.Context, *NotificationUsageStatQuery) (*_struct.Struct, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stat not implemented")
}
func (UnimplementedNotificationUsageServer) mustEmbedUnimplementedNotificationUsageServer() {}

// UnsafeNotificationUsageServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NotificationUsageServer will
// result in compilation errors.
type UnsafeNotificationUsageServer interface {
	mustEmbedUnimplementedNotificationUsageServer()
}

func RegisterNotificationUsageServer(s grpc.ServiceRegistrar, srv NotificationUsageServer) {
	s.RegisterService(&NotificationUsage_ServiceDesc, srv)
}

func _NotificationUsage_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NotificationUsageQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationUsageServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NotificationUsage_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationUsageServer).List(ctx, req.(*NotificationUsageQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _NotificationUsage_Stat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NotificationUsageStatQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationUsageServer).Stat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NotificationUsage_Stat_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationUsageServer).Stat(ctx, req.(*NotificationUsageStatQuery))
	}
	return interceptor(ctx, in, info, handler)
}

// NotificationUsage_ServiceDesc is the grpc.ServiceDesc for NotificationUsage service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NotificationUsage_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "spaceone.api.notification.v1.NotificationUsage",
	HandlerType: (*NotificationUsageServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "list",
			Handler:    _NotificationUsage_List_Handler,
		},
		{
			MethodName: "stat",
			Handler:    _NotificationUsage_Stat_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "spaceone/api/notification/v1/notification_usage.proto",
}
