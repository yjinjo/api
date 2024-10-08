// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.6.1
// source: spaceone/api/identity/v2/role_binding.proto

package v2

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
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	RoleBinding_Create_FullMethodName     = "/spaceone.api.identity.v2.RoleBinding/create"
	RoleBinding_UpdateRole_FullMethodName = "/spaceone.api.identity.v2.RoleBinding/update_role"
	RoleBinding_Delete_FullMethodName     = "/spaceone.api.identity.v2.RoleBinding/delete"
	RoleBinding_Get_FullMethodName        = "/spaceone.api.identity.v2.RoleBinding/get"
	RoleBinding_List_FullMethodName       = "/spaceone.api.identity.v2.RoleBinding/list"
	RoleBinding_Stat_FullMethodName       = "/spaceone.api.identity.v2.RoleBinding/stat"
)

// RoleBindingClient is the client API for RoleBinding service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RoleBindingClient interface {
	Create(ctx context.Context, in *CreateRoleBindingRequest, opts ...grpc.CallOption) (*RoleBindingInfo, error)
	UpdateRole(ctx context.Context, in *UpdateRoleBindingRequest, opts ...grpc.CallOption) (*RoleBindingInfo, error)
	Delete(ctx context.Context, in *RoleBindingRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	Get(ctx context.Context, in *RoleBindingRequest, opts ...grpc.CallOption) (*RoleBindingInfo, error)
	List(ctx context.Context, in *RoleBindingSearchQuery, opts ...grpc.CallOption) (*RoleBindingsInfo, error)
	Stat(ctx context.Context, in *RoleBindingStatQuery, opts ...grpc.CallOption) (*_struct.Struct, error)
}

type roleBindingClient struct {
	cc grpc.ClientConnInterface
}

func NewRoleBindingClient(cc grpc.ClientConnInterface) RoleBindingClient {
	return &roleBindingClient{cc}
}

func (c *roleBindingClient) Create(ctx context.Context, in *CreateRoleBindingRequest, opts ...grpc.CallOption) (*RoleBindingInfo, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RoleBindingInfo)
	err := c.cc.Invoke(ctx, RoleBinding_Create_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleBindingClient) UpdateRole(ctx context.Context, in *UpdateRoleBindingRequest, opts ...grpc.CallOption) (*RoleBindingInfo, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RoleBindingInfo)
	err := c.cc.Invoke(ctx, RoleBinding_UpdateRole_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleBindingClient) Delete(ctx context.Context, in *RoleBindingRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, RoleBinding_Delete_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleBindingClient) Get(ctx context.Context, in *RoleBindingRequest, opts ...grpc.CallOption) (*RoleBindingInfo, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RoleBindingInfo)
	err := c.cc.Invoke(ctx, RoleBinding_Get_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleBindingClient) List(ctx context.Context, in *RoleBindingSearchQuery, opts ...grpc.CallOption) (*RoleBindingsInfo, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RoleBindingsInfo)
	err := c.cc.Invoke(ctx, RoleBinding_List_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleBindingClient) Stat(ctx context.Context, in *RoleBindingStatQuery, opts ...grpc.CallOption) (*_struct.Struct, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(_struct.Struct)
	err := c.cc.Invoke(ctx, RoleBinding_Stat_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RoleBindingServer is the server API for RoleBinding service.
// All implementations must embed UnimplementedRoleBindingServer
// for forward compatibility.
type RoleBindingServer interface {
	Create(context.Context, *CreateRoleBindingRequest) (*RoleBindingInfo, error)
	UpdateRole(context.Context, *UpdateRoleBindingRequest) (*RoleBindingInfo, error)
	Delete(context.Context, *RoleBindingRequest) (*empty.Empty, error)
	Get(context.Context, *RoleBindingRequest) (*RoleBindingInfo, error)
	List(context.Context, *RoleBindingSearchQuery) (*RoleBindingsInfo, error)
	Stat(context.Context, *RoleBindingStatQuery) (*_struct.Struct, error)
	mustEmbedUnimplementedRoleBindingServer()
}

// UnimplementedRoleBindingServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedRoleBindingServer struct{}

func (UnimplementedRoleBindingServer) Create(context.Context, *CreateRoleBindingRequest) (*RoleBindingInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedRoleBindingServer) UpdateRole(context.Context, *UpdateRoleBindingRequest) (*RoleBindingInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRole not implemented")
}
func (UnimplementedRoleBindingServer) Delete(context.Context, *RoleBindingRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedRoleBindingServer) Get(context.Context, *RoleBindingRequest) (*RoleBindingInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedRoleBindingServer) List(context.Context, *RoleBindingSearchQuery) (*RoleBindingsInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedRoleBindingServer) Stat(context.Context, *RoleBindingStatQuery) (*_struct.Struct, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stat not implemented")
}
func (UnimplementedRoleBindingServer) mustEmbedUnimplementedRoleBindingServer() {}
func (UnimplementedRoleBindingServer) testEmbeddedByValue()                     {}

// UnsafeRoleBindingServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RoleBindingServer will
// result in compilation errors.
type UnsafeRoleBindingServer interface {
	mustEmbedUnimplementedRoleBindingServer()
}

func RegisterRoleBindingServer(s grpc.ServiceRegistrar, srv RoleBindingServer) {
	// If the following call pancis, it indicates UnimplementedRoleBindingServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&RoleBinding_ServiceDesc, srv)
}

func _RoleBinding_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRoleBindingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleBindingServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RoleBinding_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleBindingServer).Create(ctx, req.(*CreateRoleBindingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoleBinding_UpdateRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRoleBindingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleBindingServer).UpdateRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RoleBinding_UpdateRole_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleBindingServer).UpdateRole(ctx, req.(*UpdateRoleBindingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoleBinding_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoleBindingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleBindingServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RoleBinding_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleBindingServer).Delete(ctx, req.(*RoleBindingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoleBinding_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoleBindingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleBindingServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RoleBinding_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleBindingServer).Get(ctx, req.(*RoleBindingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoleBinding_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoleBindingSearchQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleBindingServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RoleBinding_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleBindingServer).List(ctx, req.(*RoleBindingSearchQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoleBinding_Stat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoleBindingStatQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleBindingServer).Stat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RoleBinding_Stat_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleBindingServer).Stat(ctx, req.(*RoleBindingStatQuery))
	}
	return interceptor(ctx, in, info, handler)
}

// RoleBinding_ServiceDesc is the grpc.ServiceDesc for RoleBinding service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RoleBinding_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "spaceone.api.identity.v2.RoleBinding",
	HandlerType: (*RoleBindingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "create",
			Handler:    _RoleBinding_Create_Handler,
		},
		{
			MethodName: "update_role",
			Handler:    _RoleBinding_UpdateRole_Handler,
		},
		{
			MethodName: "delete",
			Handler:    _RoleBinding_Delete_Handler,
		},
		{
			MethodName: "get",
			Handler:    _RoleBinding_Get_Handler,
		},
		{
			MethodName: "list",
			Handler:    _RoleBinding_List_Handler,
		},
		{
			MethodName: "stat",
			Handler:    _RoleBinding_Stat_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "spaceone/api/identity/v2/role_binding.proto",
}
