// A Notification is a service that delivers event data generated in Cloudforet to a Project or User.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.6.1
// source: spaceone/api/notification/v1/notification.proto

package v1

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
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
	Notification_Create_FullMethodName  = "/spaceone.api.notification.v1.Notification/create"
	Notification_Push_FullMethodName    = "/spaceone.api.notification.v1.Notification/push"
	Notification_Delete_FullMethodName  = "/spaceone.api.notification.v1.Notification/delete"
	Notification_SetRead_FullMethodName = "/spaceone.api.notification.v1.Notification/set_read"
	Notification_Get_FullMethodName     = "/spaceone.api.notification.v1.Notification/get"
	Notification_List_FullMethodName    = "/spaceone.api.notification.v1.Notification/list"
	Notification_Stat_FullMethodName    = "/spaceone.api.notification.v1.Notification/stat"
)

// NotificationClient is the client API for Notification service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NotificationClient interface {
	// Creates a new Notification. When a Notification is created, it is delivered to a UserChannel or a ProjectChannel depending on the parameter `resource_type`. If a Notification is delivered to a UserChannel, the `resource_type` is `identity.User`, and if a Notification is delivered to a ProjectChannel, the `resource_type` is `identity.Project`.
	Create(ctx context.Context, in *CreateNotificationRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// Manually raises a Notification. A Notification is raised with a message to be sent using a valid specific Protocol, and data used for a specific Protocol such as a phone number.
	Push(ctx context.Context, in *PushNotificationRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// Deletes multiple Notifications. You must specify `notifications` of the list of Notifications to delete.
	Delete(ctx context.Context, in *NotificationRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// Marks a Notification as read. When a Notification is raised and if the Notification has been acknowledged, it can be marked as read with the method.
	SetRead(ctx context.Context, in *SetReadNotificationRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// Gets a specific Notification. Prints detailed information about the Notification, including not only the message contents(`title`, `description`) but also related data such as created time and urgency.
	Get(ctx context.Context, in *GetNotificationRequest, opts ...grpc.CallOption) (*NotificationInfo, error)
	// Gets a list of all Notifications. You can use a query to get a filtered list of Notifications.
	List(ctx context.Context, in *NotificationQuery, opts ...grpc.CallOption) (*NotificationsInfo, error)
	Stat(ctx context.Context, in *NotificationStatQuery, opts ...grpc.CallOption) (*_struct.Struct, error)
}

type notificationClient struct {
	cc grpc.ClientConnInterface
}

func NewNotificationClient(cc grpc.ClientConnInterface) NotificationClient {
	return &notificationClient{cc}
}

func (c *notificationClient) Create(ctx context.Context, in *CreateNotificationRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, Notification_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationClient) Push(ctx context.Context, in *PushNotificationRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, Notification_Push_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationClient) Delete(ctx context.Context, in *NotificationRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, Notification_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationClient) SetRead(ctx context.Context, in *SetReadNotificationRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, Notification_SetRead_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationClient) Get(ctx context.Context, in *GetNotificationRequest, opts ...grpc.CallOption) (*NotificationInfo, error) {
	out := new(NotificationInfo)
	err := c.cc.Invoke(ctx, Notification_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationClient) List(ctx context.Context, in *NotificationQuery, opts ...grpc.CallOption) (*NotificationsInfo, error) {
	out := new(NotificationsInfo)
	err := c.cc.Invoke(ctx, Notification_List_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationClient) Stat(ctx context.Context, in *NotificationStatQuery, opts ...grpc.CallOption) (*_struct.Struct, error) {
	out := new(_struct.Struct)
	err := c.cc.Invoke(ctx, Notification_Stat_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NotificationServer is the server API for Notification service.
// All implementations must embed UnimplementedNotificationServer
// for forward compatibility
type NotificationServer interface {
	// Creates a new Notification. When a Notification is created, it is delivered to a UserChannel or a ProjectChannel depending on the parameter `resource_type`. If a Notification is delivered to a UserChannel, the `resource_type` is `identity.User`, and if a Notification is delivered to a ProjectChannel, the `resource_type` is `identity.Project`.
	Create(context.Context, *CreateNotificationRequest) (*empty.Empty, error)
	// Manually raises a Notification. A Notification is raised with a message to be sent using a valid specific Protocol, and data used for a specific Protocol such as a phone number.
	Push(context.Context, *PushNotificationRequest) (*empty.Empty, error)
	// Deletes multiple Notifications. You must specify `notifications` of the list of Notifications to delete.
	Delete(context.Context, *NotificationRequest) (*empty.Empty, error)
	// Marks a Notification as read. When a Notification is raised and if the Notification has been acknowledged, it can be marked as read with the method.
	SetRead(context.Context, *SetReadNotificationRequest) (*empty.Empty, error)
	// Gets a specific Notification. Prints detailed information about the Notification, including not only the message contents(`title`, `description`) but also related data such as created time and urgency.
	Get(context.Context, *GetNotificationRequest) (*NotificationInfo, error)
	// Gets a list of all Notifications. You can use a query to get a filtered list of Notifications.
	List(context.Context, *NotificationQuery) (*NotificationsInfo, error)
	Stat(context.Context, *NotificationStatQuery) (*_struct.Struct, error)
	mustEmbedUnimplementedNotificationServer()
}

// UnimplementedNotificationServer must be embedded to have forward compatible implementations.
type UnimplementedNotificationServer struct {
}

func (UnimplementedNotificationServer) Create(context.Context, *CreateNotificationRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedNotificationServer) Push(context.Context, *PushNotificationRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Push not implemented")
}
func (UnimplementedNotificationServer) Delete(context.Context, *NotificationRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedNotificationServer) SetRead(context.Context, *SetReadNotificationRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetRead not implemented")
}
func (UnimplementedNotificationServer) Get(context.Context, *GetNotificationRequest) (*NotificationInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedNotificationServer) List(context.Context, *NotificationQuery) (*NotificationsInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedNotificationServer) Stat(context.Context, *NotificationStatQuery) (*_struct.Struct, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stat not implemented")
}
func (UnimplementedNotificationServer) mustEmbedUnimplementedNotificationServer() {}

// UnsafeNotificationServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NotificationServer will
// result in compilation errors.
type UnsafeNotificationServer interface {
	mustEmbedUnimplementedNotificationServer()
}

func RegisterNotificationServer(s grpc.ServiceRegistrar, srv NotificationServer) {
	s.RegisterService(&Notification_ServiceDesc, srv)
}

func _Notification_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateNotificationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Notification_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationServer).Create(ctx, req.(*CreateNotificationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Notification_Push_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PushNotificationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationServer).Push(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Notification_Push_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationServer).Push(ctx, req.(*PushNotificationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Notification_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NotificationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Notification_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationServer).Delete(ctx, req.(*NotificationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Notification_SetRead_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetReadNotificationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationServer).SetRead(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Notification_SetRead_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationServer).SetRead(ctx, req.(*SetReadNotificationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Notification_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNotificationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Notification_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationServer).Get(ctx, req.(*GetNotificationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Notification_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NotificationQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Notification_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationServer).List(ctx, req.(*NotificationQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _Notification_Stat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NotificationStatQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationServer).Stat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Notification_Stat_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationServer).Stat(ctx, req.(*NotificationStatQuery))
	}
	return interceptor(ctx, in, info, handler)
}

// Notification_ServiceDesc is the grpc.ServiceDesc for Notification service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Notification_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "spaceone.api.notification.v1.Notification",
	HandlerType: (*NotificationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "create",
			Handler:    _Notification_Create_Handler,
		},
		{
			MethodName: "push",
			Handler:    _Notification_Push_Handler,
		},
		{
			MethodName: "delete",
			Handler:    _Notification_Delete_Handler,
		},
		{
			MethodName: "set_read",
			Handler:    _Notification_SetRead_Handler,
		},
		{
			MethodName: "get",
			Handler:    _Notification_Get_Handler,
		},
		{
			MethodName: "list",
			Handler:    _Notification_List_Handler,
		},
		{
			MethodName: "stat",
			Handler:    _Notification_Stat_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "spaceone/api/notification/v1/notification.proto",
}
